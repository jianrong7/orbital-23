/*
idlmanagement is the code for the IDL Management Service.

It is built using the Hertz HTTP framework.
https://github.com/cloudwego/hertz

It also utilises the Hashicorp Consul registry
https://www.hashicorp.com/products/consul

Usage:

	./idlmanagement [Consul Private Address] [Own Address] [API Gateway Public Address]
*/

package main

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/fsnotify/fsnotify"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
	jsoniter "github.com/json-iterator/go"
)

func readFileUpdateAPIGateway() {
	// read in the new service mapping from the serviceMap.json file, reallocating a new map
	serviceMap := make(map[string]map[string]string)
	content, err := os.ReadFile("serviceMap.json")
	if err != nil {
		log.Println("Problem reading serverConfig.json")
		panic(err)
	}
	err = jsoniter.Unmarshal(content, &serviceMap)
	if err != nil {
		log.Println("Problem unmarshalling config")
		panic(err)
	}
	api_gw_addr := os.Args[3]
	address := "http://" + api_gw_addr + "/idlmanagement/update" // api gateway address as command-line argument 3
	_, err = http.Post(address, "application/json",
		bytes.NewBuffer(content))
	if err != nil {
		log.Println("Problem sending POST request update")
		panic(err)
	}
}

func runFileWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				// if the file serviceMap.json is modified, it will generate a Write operation, which will trigger the update
				if event.Name == "serviceMap.json" && event.Has(fsnotify.Write) {
					log.Println("serviceMap.json modified")
					go readFileUpdateAPIGateway()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add a path.
	err = watcher.Add("./")
	if err != nil {
		log.Fatal(err)
	}

	// Block function goroutine forever.
	<-make(chan struct{})
}

func main() {
	// build a consul client
	config := consulapi.DefaultConfig()
	consul_address := os.Args[1] // taking the Consul server address as command-line argument 1 to make deployment easier
	config.Address = consul_address
	consulClient, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
		return
	}
	// build a consul register with the consul client
	r := consul.NewConsulRegister(consulClient)
	// run Hertz with the consul register
	if err != nil {
		log.Fatal(err)
	}
	own_address := os.Args[2] // own address as command-line argument 2
	h := server.Default(
		server.WithHostPorts("0.0.0.0:9999"), // use "0.0.0.0:9999" to listen to all addresses on port 9999
		server.WithRegistry(r, &registry.Info{
			ServiceName: "idlmanagement",
			Addr:        utils.NewNetAddr("tcp", own_address),
			Weight:      10,
			Tags:        nil,
		}),
	)

	h.OnRun = append(h.OnRun, func(ctx context.Context) error {
		go func() {
			// to allow the idlmanagement http server to start up properly before it sends the update request
			time.Sleep(10 * time.Second) // hacky time delay -- no idea how to implement asynchronous flag to trigger
			readFileUpdateAPIGateway()
		}()
		return nil
	})

	h.OnRun = append(h.OnRun, func(ctx context.Context) error {
		go runFileWatcher()
		return nil
	})

	// Future improvement: Protect with authentication middleware
	h.GET("/getthriftfile/:name", func(c context.Context, ctx *app.RequestContext) {
		thriftFileName := ctx.Param("name")
		content, err := os.ReadFile(thriftFileName)
		if err != nil {
			log.Println("Problem reading " + thriftFileName)
			panic(err)
		}
		_, err = ctx.Write(content)
		if err != nil {
			log.Println("Problem writing to app requestcontext " + thriftFileName)
			panic(err)
		}
		ctx.SetStatusCode(consts.StatusOK)
	})
	h.Spin()
}
