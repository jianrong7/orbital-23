package main

import (
	"net"

	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/server/genericserver"
)

func main() {
	g2 := generic.BinaryThriftGeneric()
	addr2, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8081")
	svr2 := genericserver.NewServer(&Service2Impl{}, g2, server.WithServiceAddr(addr2))
	err2 := svr2.Run()
	if err2 != nil {
		panic(err2)
	}
}
