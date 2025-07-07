package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// DatabaseConfig 数据库配置接口
type DatabaseConfig interface {
	GetHost() string
	GetPort() int
	GetUser() string
	GetPassword() string
	GetDBName() string
	GetSSLMode() string
}

// NewDBConn 创建数据库连接
func NewDBConn(cfg DatabaseConfig) (*sql.DB, func(), error) {
	// 1. 构建数据库连接字符串
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.GetHost(), cfg.GetPort(), cfg.GetUser(), cfg.GetPassword(), cfg.GetDBName(), cfg.GetSSLMode())

	// 2. 打开数据库连接
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, nil, fmt.Errorf("打开数据库失败: %w", err)
	}

	// 3. 测试数据库连接是否可用
	pingErr := db.Ping()
	if pingErr != nil {
		// 3.1. 连接测试失败时，关闭数据库连接并返回错误
		db.Close() // 尝试关闭连接，即使失败也忽略
		return nil, nil, fmt.Errorf("数据库连接测试失败: %w", pingErr)
	}

	// 4. 定义清理函数，用于关闭数据库连接
	cleanup := func() {
		if err := db.Close(); err != nil {
			log.Printf("错误: 关闭数据库连接失败: %v", err)
		}
	}

	// 5. 返回数据库连接、清理函数和nil错误
	return db, cleanup, nil
}
