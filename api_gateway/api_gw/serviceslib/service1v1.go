package services

import (
	s1v1 "api_gw/service_definitions/kitex_gen/service1v1"
	"errors"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/hertz/pkg/app"
	jsoniter "github.com/json-iterator/go"
)

func Service1v1(methodName string, ctx *app.RequestContext) (requestStruct thrift.TStruct, responseStruct thrift.TStruct, err error) {
	switch methodName {
	case "Add":
		var req s1v1.AddRequest
		err = jsoniter.Unmarshal(ctx.GetRawData(), &req)
		var res s1v1.AddResponse
		return &req, &res, err
	}
	return nil, nil, errors.New("No method found associated with service")
}
