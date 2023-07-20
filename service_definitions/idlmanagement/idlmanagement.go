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

func delayReadFileUpdateAPIGateway() { // to allow the idlmanagement http server to start up properly before it sends the update request
	time.Sleep(5 * time.Second)
	readFileUpdateAPIGateway()
}

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
	_, err = http.Post("http://54.179.129.9:8888/idlmanagement/update", "application/json",
		bytes.NewBuffer(content))
	if err != nil {
		log.Println("Problem sending POST request update")
		panic(err)
	}
}

// func getLocalIPv4Address() (string, error) {
// 	addr, err := net.InterfaceAddrs()
// 	if err != nil {
// 		return "", err
// 	}
// 	for _, addr := range addr {
// 		ipNet, isIpNet := addr.(*net.IPNet)
// 		if isIpNet && !ipNet.IP.IsLoopback() {
// 			ipv4 := ipNet.IP.To4()
// 			if ipv4 != nil {
// 				return ipv4.String(), nil
// 			}
// 		}
// 	}
// 	return "", fmt.Errorf("not found ipv4 address")
// }

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
	config.Address = "54.255.152.160:8500"
	consulClient, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
		return
	}
	// build a consul register with the consul client
	r := consul.NewConsulRegister(consulClient)
	// run Hertz with the consul register
	// localIP, err := getLocalIPv4Address()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	addr := "172.31.30.231:9999"
	h := server.Default(
		server.WithHostPorts(addr),
		server.WithRegistry(r, &registry.Info{
			ServiceName: "idlmanagement",
			Addr:        utils.NewNetAddr("tcp", addr),
			Weight:      10,
			Tags:        nil,
		}),
	)

	h.OnRun = append(h.OnRun, func(ctx context.Context) error {
		log.Println("Running the first hook")
		go delayReadFileUpdateAPIGateway()
		return nil
	})

	h.OnRun = append(h.OnRun, func(ctx context.Context) error {
		log.Println("Running the second hook")
		go runFileWatcher()
		return nil
	})

	// TODO: Protect with basicauth
	h.GET("/getthriftfile/:name", func(c context.Context, ctx *app.RequestContext) {
		thriftFileName := ctx.Param("name")
		log.Println(thriftFileName)
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

	// go runHTTPServer()

}
