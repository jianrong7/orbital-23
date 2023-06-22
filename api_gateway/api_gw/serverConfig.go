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
	config.Address = "13.229.205.99:8500"
	consulClient, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	addr := net.JoinHostPort("0.0.0.0", "8888")
	r := consul.NewConsulRegister(consulClient)
	return server.Default(
		server.WithHostPorts(addr),
		server.WithRegistry(r, &registry.Info{
			ServiceName: "hertz.test.demo",
			Addr:        utils.NewNetAddr("tcp", net.JoinHostPort("18.139.209.232", "8888")),
			Weight:      10,
			Tags:        nil,
		}),
	)
}
