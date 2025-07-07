package temporal

import (
	"fmt"

	"go.temporal.io/sdk/client"
)

// TemporalConfig Temporal配置接口
type TemporalConfig interface {
	GetHostPort() string
	GetNamespace() string
}

// NewTemporalClient 创建并返回一个 Temporal 客户端
func NewTemporalClient(cfg TemporalConfig) (client.Client, error) {
	// 1. 配置 Temporal 客户端连接参数并建立连接
	c, err := client.Dial(client.Options{
		HostPort:  cfg.GetHostPort(),  // Temporal 服务器地址和端口
		Namespace: cfg.GetNamespace(), // Temporal 命名空间
	})
	if err != nil {
		return nil, fmt.Errorf("创建 Temporal 客户端失败: %w", err)
	}

	// 2. 返回 Temporal 客户端实例
	return c, nil
}
