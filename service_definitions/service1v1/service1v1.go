/*
service1v1 is the code for the RPC service service1v1.

It is built using the Kitex RPC framework.
https://github.com/cloudwego/kitex

It also utilises the Hashicorp Consul registry
https://www.hashicorp.com/products/consul

Usage:

	./service1v1 [Consul Private Address]
*/

package main

import (
	"log"
	"net"
	"os"
	"sync"

	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/server/genericserver"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	consul_address := os.Args[1] // taking the Consul server address as a command-line argument to make deployment easier
	r, err := consul.NewConsulRegister(consul_address)
	if err != nil {
		log.Fatal(err)
	}
	// https://www.cloudwego.io/docs/kitex/tutorials/advanced-feature/generic-call/#4-json-mapping-generic-call
	p, err := generic.NewThriftFileProvider("./service1v1.thrift")
	if err != nil {
		panic(err)
	}
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	svr0 := genericserver.NewServer(
		&Service1Impl{},
		g,
		server.WithMiddleware(PrintService1Server1),
		server.WithRegistry(r),
		server.WithRegistryInfo(&registry.Info{ServiceName: "service1v1", Weight: 1}), // can experiment with changing the weight, higher weight means higher likelihood of being used.
		server.WithServiceAddr(&net.TCPAddr{Port: 8080}),
	)
	svr1 := genericserver.NewServer(
		&Service1Impl{},
		g,
		server.WithMiddleware(PrintService1Server2),
		server.WithRegistry(r),
		server.WithRegistryInfo(&registry.Info{ServiceName: "service1v1", Weight: 1}),
		server.WithServiceAddr(&net.TCPAddr{Port: 8081}),
	)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := svr0.Run(); err != nil {
			log.Println("server0 stopped with error:", err)
		} else {
			log.Println("server0 stopped")
		}
	}()

	go func() {
		defer wg.Done()
		if err := svr1.Run(); err != nil {
			log.Println("server1 stopped with error:", err)
		} else {
			log.Println("server1 stopped")
		}
	}()
	wg.Wait()
}
