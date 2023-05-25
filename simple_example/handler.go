package main

import (
	"context"
	"log"
	api "simpleExample/kitex_gen/api"

	jsoniter "github.com/json-iterator/go"
)

// ExampleImpl implements the last service interface defined in the IDL.
type ExampleImpl struct{}

func (g *ExampleImpl) GenericCall(c context.Context, method string, request interface{}) (response interface{}, err error) {
	// use JSON parsing library to assert request
	m := request.(string) // string type assertion
	log.Println(m)
	switch method {
	case "Add":
		var nums api.AddRequest
		err = jsoniter.Unmarshal([]byte(m), &nums)
		if err != nil {
			log.Fatal(err)
		}
		response = api.AddResponse{Sum: nums.First + nums.Second}
		return jsoniter.MarshalToString(response)
	}
	return
}

// // Echo implements the ExampleImpl interface.
// func (s *ExampleImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
// 	return &api.Response{Message: req.Message}, nil
// }

// // Add implements the ExampleImpl interface.
// func (s *ExampleImpl) Add(ctx context.Context, req *api.AddRequest) (resp *api.AddResponse, err error) {
// 	resp = &api.AddResponse{Sum: req.First + req.Second}
// 	return
// }
