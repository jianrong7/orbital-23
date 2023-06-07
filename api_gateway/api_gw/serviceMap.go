package main

import (
	sl "api_gw/api_gw/serviceslib"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/hertz/pkg/app"
)

var serviceMap = map[string]func(string, *app.RequestContext) (thrift.TStruct, thrift.TStruct, error){
	"service1v1": sl.Service1v1,
	"service2v1": sl.Service2v1,
}

func FillRequestGetResponse(serviceName string, methodName string, ctx *app.RequestContext) (requestStruct thrift.TStruct, responseStruct thrift.TStruct, err error) {
	return serviceMap[serviceName](methodName, ctx)
}


