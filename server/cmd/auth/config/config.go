/**
 * @Author: zzy
 * @Description:
 * @File: config.go
 * @Version: 1.0.0
 * @Date: 2023/5/6 10:30
 */

package config

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Key  string `mapstructure:"key" json:"key"`
}

type ServerConfig struct {
	Name string `mapstructure:"name" json:"name"`
	Host string `mapstructure:"host" json:"host"`
}
