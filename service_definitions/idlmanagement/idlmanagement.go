package main

import (
	"bytes"
	"context"
	"log"
	"net"
	"net/http"
	"os"

	idlm "api_gw/service_definitions/kitex_gen/idlmanagement/idlmanagement"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	jsoniter "github.com/json-iterator/go"
	consul "github.com/kitex-contrib/registry-consul"
)

func updateAPIGateway(content []byte) {
	_, err := http.Post("http://127.0.0.1:8888/idlmanagement/update", "application/json",
		bytes.NewBuffer(content))
	if err != nil {
		log.Println("Problem sending POST request update")
		panic(err)
	}
}

// GLOBAL VARIABLE
var serviceMap = make(map[string]map[string]string)

func main() {
	// update the API Gateway that the IDL Management service has changes with HTTP Get request
	// API Gateway will call the relevant functions to update itself with RPC

	r, err := consul.NewConsulRegister("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}

	svr := idlm.NewServer(
		new(IDLManagementImpl),
		server.WithRegistry(r),
		server.WithRegistryInfo(&registry.Info{ServiceName: "idlmanagement", Weight: 1}),
		server.WithServiceAddr(&net.TCPAddr{Port: 9999}),
	)

	err = svr.Run()
	if err != nil {
		log.Fatal(err)
	}

	// // read in the new service mapping from the serviceMap.json file, reallocating a new map
	// serviceMap = make(map[string]map[string]string)
	// content, err := os.ReadFile("serviceMap.json")
	// if err != nil {
	// 	log.Println("Problem reading serverConfig.json")
	// 	panic(err)
	// }

	// err = jsoniter.Unmarshal(content, &serviceMap)
	// if err != nil {
	// 	log.Println("Problem unmarshalling config")
	// 	panic(err)
	// }

	// updateAPIGateway(content)
}

type IDLManagementImpl struct{}

// GetThriftFile implements the IDLManagementImpl interface.
func (s *IDLManagementImpl) GetThriftFile(ctx context.Context, fileName string) (resp string, err error) {
	log.Println(fileName)
	log.Println(serviceMap)
	log.Println("xx")
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Println("Problem reading " + fileName)
		panic(err)
	}
	return string(content), err
}

// watcher, err := fsnotify.NewWatcher()
// if err != nil {
// 	log.Fatal(err)
// }
// defer watcher.Close()

// // Start listening for events.
// go func() {
// 	for {
// 		select {
// 		case event, ok := <-watcher.Events:
// 			if !ok {
// 				return
// 			}
// 			log.Println("event:", event)
// 			if event.Name == "./serviceMap.json" {
// 				// update the API Gateway via HTTP POST request
// 				updateAPIGateway(content)
// 			}
// 		case err, ok := <-watcher.Errors:
// 			if !ok {
// 				return
// 			}
// 			log.Println("error:", err)
// 		}
// 	}
// }()

// // Add a path.
// err = watcher.Add("./")
// if err != nil {
// 	log.Fatal(err)
// }

// // Block main goroutine forever.
// <-make(chan struct{})
