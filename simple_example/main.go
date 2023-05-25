package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/server/genericserver"
)

func main() {
	// addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.2:8888")
	// svr := api.NewServer(new(ExampleImpl), server.WithServiceAddr(addr))

	p, err := generic.NewThriftFileProvider("./ex.thrift")
	if err != nil {
		panic(err)
	}
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.2:8888")
	svr := genericserver.NewServer(new(ExampleImpl), g, server.WithServiceAddr(addr))
	if err != nil {
		panic(err)
	}
	err = svr.Run()
	if err != nil {
		panic(err)
	}

	if err != nil {
		log.Println(err.Error())
	}
}
