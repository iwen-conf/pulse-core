package connector

import (
	"fmt"
	"time"

	"google.golang.org/grpc"
)

// ClientConfig 客户端配置
type ClientConfig struct {
	Host               string        `json:"host" mapstructure:"host"`
	Port               int           `json:"port" mapstructure:"port"`
	Timeout            time.Duration `json:"timeout" mapstructure:"timeout"`
	Debug              bool          `json:"debug" mapstructure:"debug"`
	HealthCheckEnabled bool          `json:"health_check_enabled" mapstructure:"health_check_enabled"`
	HealthCheckInterval time.Duration `json:"health_check_interval" mapstructure:"health_check_interval"`
	TLS                TLSConfig     `json:"tls" mapstructure:"tls"`
}

// DefaultClientConfig 返回默认的客户端配置
func DefaultClientConfig() ClientConfig {
	return ClientConfig{
		Host:                "localhost",
		Port:                9090,
		Timeout:             10 * time.Second,
		Debug:               false,
		HealthCheckEnabled:  false,
		HealthCheckInterval: 30 * time.Second,
		TLS:                 DefaultTLSConfig(),
	}
}

// NewGRPCClient 创建新的gRPC客户端连接
func NewGRPCClient(config ClientConfig) (*grpc.ClientConn, func(), error) {
	// 验证配置
	if err := ValidateClientConfig(config); err != nil {
		return nil, nil, fmt.Errorf("客户端配置无效: %w", err)
	}

	// 构建目标地址
	target := fmt.Sprintf("%s:%d", config.Host, config.Port)

	// 创建连接管理器选项
	var opts []ManagerOption

	// 添加TLS选项
	if config.TLS.Enabled {
		opts = append(opts, WithTLS(config.TLS))
	}

	// 添加健康检查选项
	if config.HealthCheckEnabled {
		opts = append(opts, WithHealthCheck(true, config.HealthCheckInterval))
	}

	// 创建连接管理器
	manager, err := NewManager(target, config.Timeout, config.Debug, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("创建连接管理器失败: %w", err)
	}

	// 返回连接和清理函数
	cleanup := func() {
		if err := manager.Close(); err != nil {
			if config.Debug {
				fmt.Printf("关闭连接时出错: %v\n", err)
			}
		}
	}

	return manager.GetConn(), cleanup, nil
}

// NewSimpleGRPCClient 创建简单的gRPC客户端连接（不带健康检查和TLS）
func NewSimpleGRPCClient(host string, port int) (*grpc.ClientConn, func(), error) {
	config := ClientConfig{
		Host:                host,
		Port:                port,
		Timeout:             10 * time.Second,
		Debug:               false,
		HealthCheckEnabled:  false,
		TLS:                 DefaultTLSConfig(),
	}

	return NewGRPCClient(config)
}

// ValidateClientConfig 验证客户端配置
func ValidateClientConfig(config ClientConfig) error {
	if config.Host == "" {
		return fmt.Errorf("主机地址不能为空")
	}

	if config.Port <= 0 || config.Port > 65535 {
		return fmt.Errorf("端口号必须在1-65535之间")
	}

	if config.Timeout <= 0 {
		return fmt.Errorf("超时时间必须大于0")
	}

	if config.HealthCheckEnabled && config.HealthCheckInterval <= 0 {
		return fmt.Errorf("启用健康检查时，检查间隔必须大于0")
	}

	// 验证TLS配置
	if err := ValidateTLSConfig(config.TLS); err != nil {
		return fmt.Errorf("TLS配置无效: %w", err)
	}

	return nil
}