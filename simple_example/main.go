package main

import (
	"log"
	"net"
	api "simpleExample/kitex_gen/api/simpleexample"

	"github.com/cloudwego/kitex/server"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.2:8888")
	svr := api.NewServer(new(ExampleImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
