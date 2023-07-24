package main

import (
	"context"
	"errors"
	"log"

	s1v2 "api_gw/service_definitions/kitex_gen/service1v2"

	jsoniter "github.com/json-iterator/go"
)

type Service1Impl struct{}

// GenericCall implements generic.Service.
func (g *Service1Impl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	log.Println("GenericCall from handler1:", request)
	// JSON Validity Check
	reqBuf := request.(string) // type assertion
	isValid := jsoniter.Valid([]byte(reqBuf))
	if !isValid {
		return nil, errors.New("Invalid JSON request: " + reqBuf)
	}
	switch method {
	case "Add":
		var req s1v2.AddRequest
		err = jsoniter.UnmarshalFromString(reqBuf, &req)
		if err != nil {
			panic(err)
		}
		respBuf := &s1v2.AddResponse{Sum: req.First + req.Second}
		res, err := jsoniter.MarshalToString(respBuf)
		if err != nil {
			panic(err)
		}
		return res, nil

	case "Sub":
		var req s1v2.SubRequest
		err = jsoniter.UnmarshalFromString(reqBuf, &req)
		if err != nil {
			panic(err)
		}
		respBuf := &s1v2.SubResponse{Diff: req.First - req.Second}
		res, err := jsoniter.MarshalToString(respBuf)
		if err != nil {
			panic(err)
		}
		return res, nil

		default:
			return nil, errors.New("Unknown method: " + method)
	}
}
