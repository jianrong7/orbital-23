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
	r, err := consul.NewConsulRegister(os.Args[1]) // consul address as command-line argument
	if err != nil {
		log.Fatal(err)
	}

	p, err := generic.NewThriftFileProvider("./service2v1.thrift")
	if err != nil {
		panic(err)
	}
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	svr0 := genericserver.NewServer(
		&Service2Impl{},
		g,
		server.WithMiddleware(PrintService2Server0),
		server.WithRegistry(r),
		server.WithRegistryInfo(&registry.Info{ServiceName: "service2v1", Weight: 1}),
		server.WithServiceAddr(&net.TCPAddr{Port: 8090}),
	)
	svr1 := genericserver.NewServer(
		&Service2Impl{},
		g,
		server.WithMiddleware(PrintService2Server1),
		server.WithRegistry(r),
		server.WithRegistryInfo(&registry.Info{ServiceName: "service2v1", Weight: 1}),
		server.WithServiceAddr(&net.TCPAddr{Port: 8091}),
	)
	svr2 := genericserver.NewServer(
		&Service2Impl{},
		g,
		server.WithMiddleware(PrintService2Server2),
		server.WithRegistry(r),
		server.WithRegistryInfo(&registry.Info{ServiceName: "service2v1", Weight: 1}),
		server.WithServiceAddr(&net.TCPAddr{Port: 8092}),
	)

	var wg sync.WaitGroup
	wg.Add(3)
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

	go func() {
		defer wg.Done()
		if err := svr2.Run(); err != nil {
			log.Println("server1 stopped with error:", err)
		} else {
			log.Println("server1 stopped")
		}
	}()
	wg.Wait()
}
