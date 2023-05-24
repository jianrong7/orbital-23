package main

import (
	"context"
	api "simpleExample/kitex_gen/api"
)

// ExampleImpl implements the last service interface defined in the IDL.
type ExampleImpl struct{}

// Echo implements the ExampleImpl interface.
func (s *ExampleImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	return &api.Response{Message: req.Message}, nil
}

// Add implements the ExampleImpl interface.
func (s *ExampleImpl) Add(ctx context.Context, req *api.AddRequest) (resp *api.AddResponse, err error) {
	resp = &api.AddResponse{Sum: req.First + req.Second}
	return
}
