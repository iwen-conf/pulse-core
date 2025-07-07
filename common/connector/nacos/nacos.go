package nacos

import (
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

// NacosConfig Nacos配置接口
type NacosConfig interface {
	GetIpAddr() string
	GetPort() uint64
	GetNamespace() string
}

// NewConfigClient 创建 Nacos 配置客户端
func NewConfigClient(cfg NacosConfig) (config_client.IConfigClient, error) {
	// 1. 配置 Nacos 服务器信息
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(cfg.GetIpAddr(), cfg.GetPort()),
	}

	// 2. 配置 Nacos 客户端参数
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(cfg.GetNamespace()),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogLevel("warn"),
	)

	// 3. 创建 Nacos 配置客户端实例
	client, err := clients.NewConfigClient(vo.NacosClientParam{ClientConfig: &cc, ServerConfigs: sc})
	if err != nil {
		return nil, fmt.Errorf("创建 Nacos 配置客户端失败: %w", err)
	}

	// 4. 返回配置客户端实例
	return client, nil
}

// NewNamingClient 创建 Nacos 服务发现客户端
func NewNamingClient(cfg NacosConfig) (naming_client.INamingClient, error) {
	// 1. 配置 Nacos 服务器信息
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(cfg.GetIpAddr(), cfg.GetPort()),
	}

	// 2. 配置 Nacos 客户端参数
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(cfg.GetNamespace()),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogLevel("warn"),
	)

	// 3. 创建 Nacos 服务发现客户端实例
	client, err := clients.NewNamingClient(vo.NacosClientParam{ClientConfig: &cc, ServerConfigs: sc})
	if err != nil {
		return nil, fmt.Errorf("创建 Nacos 服务发现客户端失败: %w", err)
	}

	// 4. 返回服务发现客户端实例
	return client, nil
}

// RegisterServiceInstance 向 Nacos 注册服务实例
func RegisterServiceInstance(client naming_client.INamingClient, serviceName string, ip string, port uint64) (bool, error) {
	// 1. 构建服务实例注册参数并向 Nacos 注册服务实例
	return client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          ip,                                    // 服务实例IP地址
		Port:        port,                                  // 服务实例端口
		ServiceName: serviceName,                           // 服务名称
		Weight:      10,                                    // 服务权重
		Enable:      true,                                  // 是否启用
		Healthy:     true,                                  // 健康状态
		Ephemeral:   true,                                  // 是否为临时实例
		Metadata:    map[string]string{"protocol": "grpc"}, // 元数据信息
	})
}
