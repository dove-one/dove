// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.New(
		server.WithHostPorts(":8880"))

	register(h)
	h.Spin()
}
