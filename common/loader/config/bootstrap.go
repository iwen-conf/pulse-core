package config

import (
	"fmt"
	"os"
	"strconv"
)

// BootstrapConfig 结构体用于映射本地的 bootstrap.toml 文件
type BootstrapConfig struct {
	Nacos struct {
		IpAddr    string `mapstructure:"ip_addr" toml:"ip_addr"`
		Port      uint64 `mapstructure:"port" toml:"port"`
		Namespace string `mapstructure:"namespace" toml:"namespace"`
		DataID    string `mapstructure:"data_id" toml:"data_id"`
		Group     string `mapstructure:"group" toml:"group"`
	} `mapstructure:"nacos" toml:"nacos"`
}

// LoadConfigFromEnv 从环境变量加载Nacos配置信息
func LoadConfigFromEnv() (*BootstrapConfig, error) {
	var bootstrapCfg BootstrapConfig
	// 从环境变量读取Nacos配置
	nacosHost := os.Getenv("NACOS_HOST")
	if nacosHost == "" {
		return nil, fmt.Errorf("NACOS_HOST 环境变量未设置 无法启动")
	}
	nacosPortStr := os.Getenv("NACOS_PORT")
	if nacosPortStr == "" {
		return nil, fmt.Errorf("NACOS_PORT 环境变量未设置 无法启动")
	}
	nacosPort, err := strconv.ParseUint(nacosPortStr, 10, 64)

	if err != nil {
		return nil, fmt.Errorf("无效的NACOS_PORT环境变量: %w", err)
	}

	nacosNamespace := os.Getenv("NACOS_NAMESPACE")
	if nacosNamespace == "" {
		return nil, fmt.Errorf("NACOS_NAMESPACE 环境变量未设置 无法启动")
	}
	nacosDataID := os.Getenv("NACOS_DATA_ID")
	if nacosDataID == "" {
		return nil, fmt.Errorf("NACOS_DATA_ID 环境变量未设置 无法启动")
	}
	nacosGroup := os.Getenv("NACOS_GROUP")
	if nacosGroup == "" {
		return nil, fmt.Errorf("NACOS_GROUP 环境变量未设置 无法启动")
	}

	bootstrapCfg.Nacos.IpAddr = nacosHost
	bootstrapCfg.Nacos.Port = nacosPort
	bootstrapCfg.Nacos.Namespace = nacosNamespace
	bootstrapCfg.Nacos.DataID = nacosDataID
	bootstrapCfg.Nacos.Group = nacosGroup

	fmt.Println("通过NEV获取的NacOS的Config为：", bootstrapCfg)

	return &bootstrapCfg, nil
}
