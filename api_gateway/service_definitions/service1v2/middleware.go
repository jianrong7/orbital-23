package main

import (
	"context"
	"log"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

func PrintService1Server1(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
			log.Println("Service1v2-Server1-8100")
			err := next(ctx, request, response)
			return err
	}
}

func PrintService1Server2(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
			log.Println("Service1v2-Server2-8101")
			err := next(ctx, request, response)
			return err
	}
}
