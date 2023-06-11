package main

import (
	"context"
	"log"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/utils"
	consulapi "github.com/hashicorp/consul/api"
)

func main() {
	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	consulClient, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
		return
	}

	service1Cli, err := genericclient.NewClient("service1v1", generic.BinaryThriftGeneric(), client.WithHostPorts("127.0.0.1:8080"))
	if err != nil {
		panic(err)
	}

	service2Cli, err := genericclient.NewClient("service2v1", generic.BinaryThriftGeneric(), client.WithHostPorts("127.0.0.1:8081"))
	if err != nil {
		panic(err)
	}

	rc := utils.NewThriftMessageCodec()

	h.POST("/:service/:method", func(c context.Context, ctx *app.RequestContext) {
		serviceName := ctx.Param("service") // see https://www.cloudwego.io/docs/hertz/tutorials/basic-feature/route/
		methodName := strings.Title(ctx.Param("method"))

		req, res, err := FillRequestGetResponse(serviceName, methodName, ctx)
		
		if err != nil {
			log.Println("Problem filling request struct")
			panic(err)
		}
		
		reqBuf, err := rc.Encode(methodName, thrift.CALL, 1, req)
		if err != nil {
			panic(err)
		}

		if serviceName == "service1v1" {
			resBuf, err := service1Cli.GenericCall(context.Background(), methodName, reqBuf)
			if err != nil {
				panic(err)
			}
			_, _, err = rc.Decode(resBuf.([]byte), res)
			if err != nil {
				panic(err)
			}
		} else if serviceName == "service2v1" {
			resBuf, err := service2Cli.GenericCall(context.Background(), methodName, reqBuf)
			if err != nil {
				panic(err)
			}
			_, _, err = rc.Decode(resBuf.([]byte), res)
			if err != nil {
				panic(err)
			}
		}

		
		log.Println(res)
		ctx.JSON(consts.StatusOK, res)
	})

	register(h)
	h.Spin()
}
