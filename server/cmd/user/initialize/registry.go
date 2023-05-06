/**
 * @Author: zzy
 * @Description:
 * @File: registry.go
 * @Version: 1.0.0
 * @Date: 2023/5/6 10:35
 */

package initialize

import (
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/dove-one/dove/server/cmd/user/config"
	"github.com/dove-one/dove/server/common/consts"
	"github.com/hashicorp/consul/api"
	consul "github.com/kitex-contrib/registry-consul"
	"net"
	"strconv"
)

// InitRegistry to init consul
func InitRegistry(Port int) (registry.Registry, *registry.Info) {
	r, err := consul.NewConsulRegister(net.JoinHostPort(
		config.GlobalConsulConfig.Host,
		strconv.Itoa(config.GlobalConsulConfig.Port)),
		consul.WithCheck(&api.AgentServiceCheck{
			Interval:                       consts.ConsulCheckInterval,
			Timeout:                        consts.ConsulCheckTimeout,
			DeregisterCriticalServiceAfter: consts.ConsulCheckDeregisterCriticalServiceAfter,
		}))
	if err != nil {
		klog.Fatalf("new consul register failed: %s", err.Error())
	}

	// Using snowflake to generate service name.
	sf, err := snowflake.NewNode(2)
	if err != nil {
		klog.Fatalf("generate service name failed: %s", err.Error())
	}
	info := &registry.Info{
		ServiceName: config.GlobalServerConfig.Name,
		Addr:        utils.NewNetAddr(consts.TCP, net.JoinHostPort(config.GlobalServerConfig.Host, strconv.Itoa(Port))),
		Tags: map[string]string{
			"ID": sf.Generate().Base36(),
		},
	}
	return r, info
}
