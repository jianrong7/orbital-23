package main

import (
	"context"
	"log"

	s1v2 "api_gw/service_definitions/kitex_gen/service1v2"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/kitex/pkg/utils"
)

type Service1Impl struct{}

// GenericCall implements generic.Service.
func (g *Service1Impl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	log.Println("GenericCall from handler1:", request)
	rc := utils.NewThriftMessageCodec()
	reqBuf := request.([]byte)
	switch method {
	case "Add":
		var req s1v2.AddRequest
		_, seqId, err := rc.Decode(reqBuf, &req)
		if err != nil {
			panic(err)
		}
		result := &s1v2.AddResponse{Sum: req.First + req.Second}
		respBuf, err := rc.Encode(method, thrift.REPLY, seqId, result)
		if err != nil {
			panic(err)
		}

		return respBuf, nil

	case "Sub":
		var req s1v2.SubRequest
		_, seqId, err := rc.Decode(reqBuf, &req)
		if err != nil {
			panic(err)
		}
		result := &s1v2.SubResponse{Diff: req.First - req.Second}
		respBuf, err := rc.Encode(method, thrift.REPLY, seqId, result)
		if err != nil {
			panic(err)
		}

		return respBuf, nil
	}
	return
}
