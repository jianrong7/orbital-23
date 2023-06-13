package main

import (
	"context"
	"log"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

func PrintService1Server1(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
			log.Println("Service1-Server1-8080")
			err := next(ctx, request, response)
			return err
	}
}

func PrintService1Server2(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
			log.Println("Service1-Server2-8081")
			err := next(ctx, request, response)
			return err
	}
}
