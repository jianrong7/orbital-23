package main

import (
	"log"
	api "simpleExample/kitex_gen/api/simpleexample"
)

func main() {
	svr := api.NewServer(new(SimpleExampleImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
