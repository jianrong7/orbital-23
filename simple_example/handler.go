package main

import (
	"context"
	api "simpleExample/kitex_gen/api"
)

// SimpleExampleImpl implements the last service interface defined in the IDL.
type SimpleExampleImpl struct{}

// Echo implements the SimpleExampleImpl interface.
func (s *SimpleExampleImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	return
}

// Add implements the SimpleExampleImpl interface.
func (s *SimpleExampleImpl) Add(ctx context.Context, req *api.AddRequest) (resp *api.AddResponse, err error) {
	// TODO: Your code here...
	return
}
