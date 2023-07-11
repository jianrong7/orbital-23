package main

import (
	idlmanagement "api_gw/service_definitions/kitex_gen/idlmanagement/idlmanagement"
	"log"
)

func main() {
	svr := idlmanagement.NewServer(new(IDLManagementImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
