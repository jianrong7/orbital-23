// Code generated by Kitex v0.5.2. DO NOT EDIT.

package service2

import (
	service2v1 "api_gw/service_definitions/kitex_gen/service2v1"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return service2ServiceInfo
}

var service2ServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "Service2"
	handlerType := (*service2v1.Service2)(nil)
	methods := map[string]kitex.MethodInfo{
		"Mul": kitex.NewMethodInfo(mulHandler, newService2MulArgs, newService2MulResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "service2v1",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.2",
		Extra:           extra,
	}
	return svcInfo
}

func mulHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*service2v1.Service2MulArgs)
	realResult := result.(*service2v1.Service2MulResult)
	success, err := handler.(service2v1.Service2).Mul(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newService2MulArgs() interface{} {
	return service2v1.NewService2MulArgs()
}

func newService2MulResult() interface{} {
	return service2v1.NewService2MulResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Mul(ctx context.Context, req *service2v1.MulRequest) (r *service2v1.MulResponse, err error) {
	var _args service2v1.Service2MulArgs
	_args.Req = req
	var _result service2v1.Service2MulResult
	if err = p.c.Call(ctx, "Mul", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
