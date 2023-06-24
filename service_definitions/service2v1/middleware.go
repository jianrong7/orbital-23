package main

import (
	"context"
	"log"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

func PrintService2Server0(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
			log.Println("Service2-Server0-8090")
			err := next(ctx, request, response)
			return err
	}
}

func PrintService2Server1(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
			log.Println("Service2-Server2-8091")
			err := next(ctx, request, response)
			return err
	}
}
func PrintService2Server2(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
			log.Println("Service2-Server2-8092")
			err := next(ctx, request, response)
			return err
	}
}
