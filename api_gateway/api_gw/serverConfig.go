package main

import (
	"net"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func initHTTPServer() *server.Hertz {
	// config := consulapi.DefaultConfig()
	// config.Address = "13.229.205.99:8500"
	// consulClient, err := consulapi.NewClient(config)
	// if err != nil {
	// 	log.Fatal(err)
	// 	panic(err)
	// }

	addr := net.JoinHostPort("0.0.0.0", "8080")
	// r := consul.NewConsulRegister(consulClient)
	return server.Default(
		server.WithHostPorts(addr),
		// server.WithHostPorts(addr),
		// server.WithRegistry(r, &registry.Info{
		// 	ServiceName: "hertz.test.demo",
		// 	Addr: utils.NewNetAddr("tcp", addr),
		// 	Weight: 10,
		// 	Tags: nil,
		// }),
	)
}
