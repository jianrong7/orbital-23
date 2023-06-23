package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

// const CONSUL_SERVER_ADDR = "172.31.28.216:8500"
const CONSUL_SERVER_ADDR = "127.0.0.1:8500"
// const API_GW_ADDR = "172.31.22.36:8888"
const API_GW_ADDR = "127.0.0.1:8888"

func initHTTPServer() *server.Hertz {
	// config := consulapi.DefaultConfig()
	// config.Address = "127.0.0.1:8500"
	// consulClient, err := consulapi.NewClient(config)
	// if err != nil {
	// 	log.Fatal(err)
	// 	panic(err)
	// }

	// addr := net.JoinHostPort("127.0.0.1", "8888")
	// r := consul.NewConsulRegister(consulClient)
	return server.Default(
		server.WithHostPorts(API_GW_ADDR),
	// 	server.WithRegistry(r, &registry.Info{
	// 		ServiceName: "hertz.test.demo",
	// 		Addr:        utils.NewNetAddr("tcp", net.JoinHostPort("127.0.0.1", "8888")),
	// 		Weight:      10,
	// 		Tags:        nil,
	// 	}),
	)
}
