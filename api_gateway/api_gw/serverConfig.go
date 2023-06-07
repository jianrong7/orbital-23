package main

import (
	"log"
	"os"

	"github.com/cloudwego/hertz/pkg/app/server"
	jsoniter "github.com/json-iterator/go"
)

type sc struct {
	WithHostPorts                    string `json:"WithHostPorts"`
	WithMaxRequestBodySize           int    `json:"WithMaxRequestBodySize"`
	WithHandleMethodNotAllowed       bool   `json:"WithHandleMethodNotAllowed"`
	WithDisablePreParseMultipartForm bool   `json:"WithDisablePreParseMultipartForm"`
	WithBasePath                     string `json:"WithBasePath"`
}

func initHTTPServer() *server.Hertz { // reads serverConfig.json file and loads configuration
	content, err := os.ReadFile("serverConfig.json")
	if err != nil {
		log.Println("Problem reading serverConfig.json")
		panic(err)
	}

	var newConfig sc
	err = jsoniter.Unmarshal(content, &newConfig)
	if err != nil {
		log.Println("Problem unmarshalling config")
		panic(err)
	}

	return server.New(server.WithHostPorts(newConfig.WithHostPorts),
		server.WithMaxRequestBodySize(newConfig.WithMaxRequestBodySize),
		server.WithHandleMethodNotAllowed(newConfig.WithHandleMethodNotAllowed),
		server.WithDisablePreParseMultipartForm(newConfig.WithDisablePreParseMultipartForm),
		server.WithBasePath(newConfig.WithBasePath))
}
