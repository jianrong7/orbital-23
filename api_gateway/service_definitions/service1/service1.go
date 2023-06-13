package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/server/genericserver"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulRegister("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}

	g := generic.BinaryThriftGeneric()
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	svr := genericserver.NewServer(
		&Service1Impl{},
		g,
		server.WithRegistry(r),
		server.WithRegistryInfo(&registry.Info{ServiceName: "service1", Weight: 1}),
		server.WithServiceAddr(addr))

	err = svr.Run()
	if err != nil {
		panic(err)
	}
}
