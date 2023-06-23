package services

import (
	// example "api_gw/service_definitions/kitex_gen/example"
	// Include your generated files from the kitex_gen folder

	"errors"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/hertz/pkg/app"
	// jsoniter "github.com/json-iterator/go"
)

func ServiceExample(methodName string, ctx *app.RequestContext) (requestStruct thrift.TStruct, responseStruct thrift.TStruct, err error) {
	switch methodName {
	case "Method1":
		// var req example.RequestStruct
		// err = jsoniter.Unmarshal(ctx.GetRawData(), &req)
		// var res example.ResponseStruct
		// return &req, &res, err

	case "Method2":
		// var req example.RequestStruct
		// err = jsoniter.Unmarshal(ctx.GetRawData(), &req)
		// var res example.ResponseStruct
		// return &req, &res, err
	}

	return nil, nil, errors.New("No method found associated with service")
}
