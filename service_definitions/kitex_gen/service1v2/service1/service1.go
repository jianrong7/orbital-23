// Code generated by Kitex v0.6.1. DO NOT EDIT.

package service1

import (
	service1v2 "api_gw/service_definitions/kitex_gen/service1v2"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return service1ServiceInfo
}

var service1ServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "Service1"
	handlerType := (*service1v2.Service1)(nil)
	methods := map[string]kitex.MethodInfo{
		"Add": kitex.NewMethodInfo(addHandler, newService1AddArgs, newService1AddResult, false),
		"Sub": kitex.NewMethodInfo(subHandler, newService1SubArgs, newService1SubResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "service1v2",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func addHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*service1v2.Service1AddArgs)
	realResult := result.(*service1v2.Service1AddResult)
	success, err := handler.(service1v2.Service1).Add(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newService1AddArgs() interface{} {
	return service1v2.NewService1AddArgs()
}

func newService1AddResult() interface{} {
	return service1v2.NewService1AddResult()
}

func subHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*service1v2.Service1SubArgs)
	realResult := result.(*service1v2.Service1SubResult)
	success, err := handler.(service1v2.Service1).Sub(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newService1SubArgs() interface{} {
	return service1v2.NewService1SubArgs()
}

func newService1SubResult() interface{} {
	return service1v2.NewService1SubResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Add(ctx context.Context, req *service1v2.AddRequest) (r *service1v2.AddResponse, err error) {
	var _args service1v2.Service1AddArgs
	_args.Req = req
	var _result service1v2.Service1AddResult
	if err = p.c.Call(ctx, "Add", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Sub(ctx context.Context, req *service1v2.SubRequest) (r *service1v2.SubResponse, err error) {
	var _args service1v2.Service1SubArgs
	_args.Req = req
	var _result service1v2.Service1SubResult
	if err = p.c.Call(ctx, "Sub", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
