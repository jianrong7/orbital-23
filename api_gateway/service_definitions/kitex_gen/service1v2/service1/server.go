// Code generated by Kitex v0.5.2. DO NOT EDIT.
package service1

import (
	service1v2 "api_gw/service_definitions/kitex_gen/service1v2"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler service1v2.Service1, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
