/**
 * @Author: zzy
 * @Description:
 * @File: config.go
 * @Version: 1.0.0
 * @Date: 2023/5/6 10:34
 */

package initialize

import (
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/dove-one/dove/server/cmd/user/config"
	"github.com/dove-one/dove/server/common/consts"
	"github.com/dove-one/dove/server/common/tools"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	"net"
	"strconv"
)

// InitConfig to init consul config server
func InitConfig() {
	v := viper.New()
	v.SetConfigFile(consts.UserConfigPath)
	if err := v.ReadInConfig(); err != nil {
		hlog.Fatalf("read viper config failed: %s", err.Error())
	}
	if err := v.Unmarshal(&config.GlobalConsulConfig); err != nil {
		hlog.Fatalf("unmarshal err failed: %s", err.Error())
	}
	hlog.Infof("Config Info: %v", config.GlobalConsulConfig)

	cfg := api.DefaultConfig()
	cfg.Address = net.JoinHostPort(
		config.GlobalConsulConfig.Host,
		strconv.Itoa(config.GlobalConsulConfig.Port))

	consulClient, err := api.NewClient(cfg)
	if err != nil {
		hlog.Fatalf("new consul client failed: %s", err.Error())
	}
	content, _, err := consulClient.KV().Get(config.GlobalConsulConfig.Key, nil)
	if err != nil {
		hlog.Fatalf("consul kv failed: %s", err.Error())
	}

	err = sonic.Unmarshal(content.Value, &config.GlobalServerConfig)
	if err != nil {
		hlog.Fatalf("sonic unmarshal config failed: %s", err.Error())
	}

	if config.GlobalServerConfig.Host == "" {
		config.GlobalServerConfig.Host, err = tools.GetLocalIPv4Address()
		if err != nil {
			hlog.Fatalf("get localIpv4Addr failed:%s", err.Error())
		}
	}
}
