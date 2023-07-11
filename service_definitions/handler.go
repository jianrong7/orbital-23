package main

import (
	service2v1 "api_gw/service_definitions/kitex_gen/service2v1"
	"context"
)

// Service2Impl implements the last service interface defined in the IDL.
type Service2Impl struct{}

// Mul implements the Service2Impl interface.
func (s *Service2Impl) Mul(ctx context.Context, req *service2v1.MulRequest) (resp *service2v1.MulResponse, err error) {
	// TODO: Your code here...
	return
}
