package config

import (
	"github.com/dove-one/dove/server/common/kitex_gen/auth/authservice"
	"github.com/dove-one/dove/server/common/kitex_gen/user/userservice"
)

var (
	GlobalServerConfig ServerConfig
	GlobalConsulConfig ConsulConfig

	GlobalUserClient userservice.Client
	GlobalAuthClient authservice.Client
)
