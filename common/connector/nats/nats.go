package nats

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

// NATSConfig NATS配置接口
type NATSConfig interface {
	GetURL() string
}

// NewNATSConn 创建并返回一个 NATS 连接
func NewNATSConn(cfg NATSConfig) (*nats.Conn, error) {
	// 1. 连接到 NATS 服务器
	nc, err := nats.Connect(cfg.GetURL())
	if err != nil {
		return nil, fmt.Errorf("连接 NATS 失败: %w", err)
	}

	// 2. 返回 NATS 连接实例
	return nc, nil
}
