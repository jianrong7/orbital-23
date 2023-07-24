package main

import (
	"context"
	"errors"
	"log"

	s2v1 "api_gw/service_definitions/kitex_gen/service2v1"

	jsoniter "github.com/json-iterator/go"
)

// Service2Impl implements the last service interface defined in the IDL.
type Service2Impl struct{}

// Mul implements the Service2Impl interface.
func (g *Service2Impl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	log.Println("GenericCall from handler2:", request)
	// JSON Validity Check
	reqBuf := request.(string) // type assertion
	isValid := jsoniter.Valid([]byte(reqBuf))
	if !isValid {
		return nil, errors.New("Invalid JSON request: " + reqBuf)
	}
	switch method {
	case "Mul":
		var req s2v1.MulRequest
		err = jsoniter.UnmarshalFromString(reqBuf, &req)
		if err != nil {
			panic(err)
		}
		// Generate response to valid request
		respBuf := &s2v1.MulResponse{Product: req.First * req.Second}
		res, err := jsoniter.MarshalToString(respBuf)
		if err != nil {
			panic(err)
		}
		return res, nil

		default:
			return nil, errors.New("Unknown method: " + method)
	}
}
