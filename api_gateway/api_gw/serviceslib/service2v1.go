package services

import (
	s2v1 "api_gw/service_definitions/kitex_gen/service2v1"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/hertz/pkg/app"
	jsoniter "github.com/json-iterator/go"
)

func Service2v1(methodName string, ctx *app.RequestContext) (requestStruct thrift.TStruct, responseStruct thrift.TStruct, err error) {
	switch methodName {
	case "Mul":
		var req s2v1.MulRequest
		err = jsoniter.Unmarshal(ctx.GetRawData(), &req)
		// perform checks on unmarshalled data

		var res s2v1.MulResponse
		return &req, &res, err
	}
	return
}
