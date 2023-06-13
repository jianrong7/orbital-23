package main

import (
	"log"
	"net"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/kitex/pkg/utils"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
)

func initHTTPServer() *server.Hertz {
	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	consulClient, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	addr := net.JoinHostPort("127.0.0.1", "8888")
	r := consul.NewConsulRegister(consulClient)
	return server.Default(
		server.WithHostPorts(addr),
		server.WithRegistry(r, &registry.Info{
			ServiceName: "hertz.test.demo",
			Addr: utils.NewNetAddr("tcp", addr),
			Weight: 10,
			Tags: nil,
		}),
	)
}
