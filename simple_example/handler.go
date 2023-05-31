package main

import (
	"context"
	"simpleExample/kitex_gen/api"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/kitex/pkg/utils"
)

// ExampleImpl implements the last service interface defined in the IDL.
type ExampleImpl struct{}

func (g *ExampleImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	rc := utils.NewThriftMessageCodec()
	var req api.AddRequest
	reqBuf := request.([]byte)
	_, seqId, err := rc.Decode(reqBuf, &req)
	if err != nil {
		panic(err)
	}

	result := &api.AddResponse{Sum: req.First + req.Second}
	respBuf, err := rc.Encode(method, thrift.REPLY, seqId, result)

	return respBuf, nil
}

// // Echo implements the ExampleImpl interface.
// func (s *ExampleImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
// 	return &api.Response{Message: req.Message}, nil
// }

// // Add implements the ExampleImpl interface.
// func (s *ExampleImpl) Add(ctx context.Context, req *api.AddRequest) (resp *api.AddResponse, err error) {
// 	resp = &api.AddResponse{Sum: req.First + req.Second}
// 	return
// }
