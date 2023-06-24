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

	svr0 := genericserver.NewServer(
		&Service1Impl{},
		g,
		server.WithMiddleware(PrintService1Server0),
		server.WithRegistry(r),
		server.WithRegistryInfo(&registry.Info{ServiceName: "service1v2", Weight: 1}), // can experiment with changing the weight, higher weight means higher likelihood of being used.
		server.WithServiceAddr(&net.TCPAddr{Port: 8100}),
	)
	svr1 := genericserver.NewServer(
		&Service1Impl{},
		g,
		server.WithMiddleware(PrintService1Server1),
		server.WithRegistry(r),
		server.WithRegistryInfo(&registry.Info{ServiceName: "service1v2", Weight: 1}),
		server.WithServiceAddr(&net.TCPAddr{Port: 8101}),
	)
	svr2 := genericserver.NewServer(
		&Service1Impl{},
		g,
		server.WithMiddleware(PrintService1Server2),
		server.WithRegistry(r),
		server.WithRegistryInfo(&registry.Info{ServiceName: "service1v2", Weight: 2}), // note that this server has a higher weight than the other two.
		server.WithServiceAddr(&net.TCPAddr{Port: 8102}),
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
			log.Println("server2 stopped with error:", err)
		} else {
			log.Println("server2 stopped")
		}
	}()
	wg.Wait()
}