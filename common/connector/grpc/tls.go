package connector

import (
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// TLSConfig TLS配置结构
type TLSConfig struct {
	Enabled            bool   `json:"enabled" mapstructure:"enabled"`
	CertFile           string `json:"cert_file" mapstructure:"cert_file"`
	KeyFile            string `json:"key_file" mapstructure:"key_file"`
	CAFile             string `json:"ca_file" mapstructure:"ca_file"`
	ServerName         string `json:"server_name" mapstructure:"server_name"`
	InsecureSkipVerify bool   `json:"insecure_skip_verify" mapstructure:"insecure_skip_verify"`
}

// DefaultTLSConfig 返回默认的TLS配置（禁用TLS）
func DefaultTLSConfig() TLSConfig {
	return TLSConfig{
		Enabled:            false,
		InsecureSkipVerify: false,
	}
}

// CreateTLSDialOption 根据TLS配置创建gRPC拨号选项
func CreateTLSDialOption(config TLSConfig) (grpc.DialOption, error) {
	if !config.Enabled {
		return nil, fmt.Errorf("TLS未启用")
	}

	tlsConfig := &tls.Config{
		ServerName:         config.ServerName,
		InsecureSkipVerify: config.InsecureSkipVerify,
	}

	// 如果提供了证书文件，加载客户端证书
	if config.CertFile != "" && config.KeyFile != "" {
		cert, err := tls.LoadX509KeyPair(config.CertFile, config.KeyFile)
		if err != nil {
			return nil, fmt.Errorf("加载客户端证书失败: %w", err)
		}
		tlsConfig.Certificates = []tls.Certificate{cert}
	}

	// 如果提供了CA文件，可以在这里加载根证书
	// 这里简化处理，实际使用时可能需要加载CA证书池

	creds := credentials.NewTLS(tlsConfig)
	return grpc.WithTransportCredentials(creds), nil
}

// ValidateTLSConfig 验证TLS配置的有效性
func ValidateTLSConfig(config TLSConfig) error {
	if !config.Enabled {
		return nil // TLS未启用时不需要验证
	}

	// 如果启用了TLS但没有提供服务器名称，且没有跳过验证，则报错
	if config.ServerName == "" && !config.InsecureSkipVerify {
		return fmt.Errorf("启用TLS时必须提供服务器名称或启用InsecureSkipVerify")
	}

	// 如果提供了证书文件，必须同时提供密钥文件
	if (config.CertFile != "" && config.KeyFile == "") || (config.CertFile == "" && config.KeyFile != "") {
		return fmt.Errorf("证书文件和密钥文件必须同时提供")
	}

	return nil
}