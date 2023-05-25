package main

import (
	"context"
	"log"
	api "simpleExample/kitex_gen/api"
)

// ExampleImpl implements the last service interface defined in the IDL.
type ExampleImpl struct{}

func (g *ExampleImpl) GenericCall(c context.Context, method string, request interface{}) (response interface{}, err error) {
	// use JSON parsing library to assert request
	m := request.(string) // string type assertion
	log.Printf("Recv: %v\n", m)
	return "{\"Msg\": \"world\"}", nilo 
}

// Echo implements the ExampleImpl interface.
func (s *ExampleImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	return &api.Response{Message: req.Message}, nil
}

// Add implements the ExampleImpl interface.
func (s *ExampleImpl) Add(ctx context.Context, req *api.AddRequest) (resp *api.AddResponse, err error) {
	resp = &api.AddResponse{Sum: req.First + req.Second}
	return
}
