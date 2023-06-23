package main

import (
	service1v2 "api_gw/service_definitions/kitex_gen/service1v2/service1"
	"log"
)

func main() {
	svr := service1v2.NewServer(new(Service1Impl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
