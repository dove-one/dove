package main

import (
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	user "github.com/dove-one/dove/server/common/kitex_gen/user/userservice"
	"log"
)

func main() {
	svr := user.NewServer(&UserServiceImpl{}, server.WithServiceAddr(utils.NewNetAddr("tcp", "127.0.0.1:8881")))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
