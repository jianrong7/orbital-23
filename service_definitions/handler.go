package main

import (
	service1v2 "api_gw/service_definitions/kitex_gen/service1v2"
	"context"
)

// Service1Impl implements the last service interface defined in the IDL.
type Service1Impl struct{}

// Add implements the Service1Impl interface.
func (s *Service1Impl) Add(ctx context.Context, req *service1v2.AddRequest) (resp *service1v2.AddResponse, err error) {
	// TODO: Your code here...
	return
}

// Sub implements the Service1Impl interface.
func (s *Service1Impl) Sub(ctx context.Context, req *service1v2.SubRequest) (resp *service1v2.SubResponse, err error) {
	// TODO: Your code here...
	return
}
