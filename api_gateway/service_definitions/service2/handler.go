package main

import (
	"context"
	"log"

	s2v1 "api_gw/service_definitions/kitex_gen/service2v1"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/kitex/pkg/utils"
)

// Service2Impl implements the last service interface defined in the IDL.
type Service2Impl struct{}

// Mul implements the Service2Impl interface.
func (g *Service2Impl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	log.Println("GenericCall:", request)
	rc := utils.NewThriftMessageCodec()
	reqBuf := request.([]byte)
	switch method {
	case "Mul":
		var req s2v1.MulRequest
			_, seqId, err := rc.Decode(reqBuf, &req)
		if err != nil {
			panic(err)
		}
		result := &s2v1.MulResponse{Sum: req.First * req.Second}
		respBuf, err := rc.Encode(method, thrift.REPLY, seqId, result)
		if err != nil {
			panic(err)
		}
	
		return respBuf, nil
	}
	return
}
