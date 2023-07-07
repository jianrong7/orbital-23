package main

import (
	"bytes"
	"context"
	"errors"
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

func main() {
	/*
		JSON Mapping as follows:
		{
			"serviceName1" :
			{
				"versionNumber1" : "thriftFileName1",
				"versionNumber2" : "thriftFileName2"
			}
			"serviceName2" :
			{
				"versionNumber1" : "thriftFileName3"
			}
		}

		****** All thrift file names must be unique ******
	*/

	// read in the service mapping from the serviceMap.json file
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

	// update the API Gateway that the IDL Management service has changes with HTTP Get request
	// API Gateway will call the relevant functions to update itself with RPC

	if err != nil {
		log.Fatal(err)
	}

	_, err = http.Post("http://127.0.0.1:8888/idlmanagement/start", "application/json",
		bytes.NewBuffer(content))

	if err != nil {
		log.Println("Problem sending Post request")
		panic(err)
	}

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
	// TODO: auto update watching the folder, will implement soon
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
	// 			if event.Has(fsnotify.Write) {

	// 				log.Println("modified file:", event.Name)
	// 			} else if event.Has(fsnotify.Create) {
	// 				log.Println("new file: ", event.Name)
	// 			} else if event.Has(fsnotify.Remove) {
	// 				log.Println("removed file: ", event.Name)
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
	// err = watcher.Add("/tmp")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Block main goroutine forever.
	// <-make(chan struct{})
}

type IDLManagementImpl struct{}

// CheckVersion implements the IDLManagementImpl interface.
func (s *IDLManagementImpl) CheckVersion(ctx context.Context) (resp string, err error) {
	return VERSIONNUMBER, nil
}

func (s *IDLManagementImpl) GetServiceThriftFileName(ctx context.Context, serviceName string) (resp string, err error) {
	fileName := serviceMap[serviceName]
	if fileName == "" {
		err = errors.New("No such service found:" + serviceName)
		return "", err
	}
	return fileName, nil
}

// GetThriftFile implements the IDLManagementImpl interface.
func (s *IDLManagementImpl) GetThriftFile(ctx context.Context, serviceName string) (resp string, err error) {
	fileName := serviceMap[serviceName]
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Println("Problem reading " + fileName)
		panic(err)
	}
	return string(content), err
}
