package main

import (
	"log"
	"net"
	"sync"

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

	// log.Println(net.InterfaceAddrs())

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
