package main

import (
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/dove-one/dove/server/cmd/user/config"
	"github.com/dove-one/dove/server/cmd/user/initialize"
	"github.com/dove-one/dove/server/cmd/user/model"
	u "github.com/dove-one/dove/server/cmd/user/model/user"
	"github.com/dove-one/dove/server/common/consts"
	userservice "github.com/dove-one/dove/server/common/kitex_gen/user/userservice"
	"log"
	"net"
	"strconv"
)

func main() {
	initialize.InitConfig()
	db := initialize.InitDB()
	model.InitTable(db)
	IP, Port := initialize.InitFlag()
	r, info := initialize.InitRegistry(Port)
	addr := utils.NewNetAddr(consts.TCP, net.JoinHostPort(IP, strconv.Itoa(Port)))

	svr := userservice.NewServer(&UserServiceImpl{
		MysqlManagerUser: u.NewUserManager(db, config.GlobalServerConfig.MysqlInfo.Salt),
	},
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
	)
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
