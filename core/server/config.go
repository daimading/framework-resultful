package server

import "testProject/core/config"

var Config = struct {
	IP       string `json:"ip" yaml:"ip"`
	Port     int    `json:"port" yaml:"port"`
	CertPath string `json:"cert_path" yaml:"cert_path"`
	KeyPath  string `json:"key_path" yaml:"key_path"`
}{
	IP:   "127.0.0.1",
	Port: 8080,
}

//注册服务配置
func init() {
	config.RegisterConfig("server", &Config)
}
