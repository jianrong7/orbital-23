// Code generated by Kitex v0.5.2. DO NOT EDIT.

package service2

import (
	service2v1 "api_gw/service_definitions/kitex_gen/service2v1"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Mul(ctx context.Context, req *service2v1.MulRequest, callOptions ...callopt.Option) (r *service2v1.MulResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kService2Client{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kService2Client struct {
	*kClient
}

func (p *kService2Client) Mul(ctx context.Context, req *service2v1.MulRequest, callOptions ...callopt.Option) (r *service2v1.MulResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Mul(ctx, req)
}