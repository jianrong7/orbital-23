package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"

	idlm "api_gw/service_definitions/kitex_gen/idlmanagement/idlmanagement"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
)

const VERSIONNUMBER = "1.1"

var serviceMap = map[string]string{
	"service1v1": "service1v1.thrift",
	"service2v1": "service2v1.thrift",
	"service1v2": "service1v2.thrift",
}

func main() {
	r, err := consul.NewConsulRegister("172.31.28.216:8500")
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
