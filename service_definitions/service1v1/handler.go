package main

import (
	"context"
	"log"

	s1v1 "api_gw/service_definitions/kitex_gen/service1v1"

	jsoniter "github.com/json-iterator/go"
)

type Service1Impl struct{}

// GenericCall implements generic.Service.
func (g *Service1Impl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	log.Println("GenericCall from handler1:", request)
	reqBuf := request.(string)
	switch method {
	case "Add":
		var req s1v1.AddRequest
		err = jsoniter.UnmarshalFromString(reqBuf, &req)
		if err != nil {
			panic(err)
		}
		respBuf := &s1v1.AddResponse{Sum: req.First + req.Second}
		res, err := jsoniter.MarshalToString(respBuf)
		if err != nil {
			panic(err)
		}
		return res, nil
	}
	return

}
