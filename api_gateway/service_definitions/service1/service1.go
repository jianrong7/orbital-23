package main

import (
	"net"

	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/server/genericserver"
)

func main() {
	g := generic.BinaryThriftGeneric()
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	svr := genericserver.NewServer(&Service1Impl{}, g, server.WithServiceAddr(addr))
	err := svr.Run()
	if err != nil {
		panic(err)
	}
}
