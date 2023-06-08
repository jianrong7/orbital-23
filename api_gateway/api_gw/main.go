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

	s1v1 "api_gw/service_definitions/kitex_gen/service1v1"
)

func main() {
	h := initHTTPServer()

	service1Cli, err := genericclient.NewClient("service1v1", generic.BinaryThriftGeneric(), client.WithHostPorts("127.0.0.1:8080"))
	if err != nil {
		panic(err)
	}

	rc := utils.NewThriftMessageCodec()

	h.POST("/:service/:method", func(c context.Context, ctx *app.RequestContext) {
		serviceName := ctx.Param("service") // see https://www.cloudwego.io/docs/hertz/tutorials/basic-feature/route/
		methodName := ctx.Param("method")

		_, _, err := FillRequestGetResponse(serviceName, methodName, ctx)
		
		if err != nil {
			log.Println("Problem filling request struct")
			panic(err)
		}
		var req s1v1.AddRequest
		var res s1v1.AddResponse
		err = ctx.BindAndValidate(&req)
		if err != nil {
			panic(err)
		}
		log.Println(strings.Title(methodName), thrift.CALL, 1, &req)
		reqBuf, err := rc.Encode(strings.Title(methodName), thrift.CALL, 1, &req)
		if err != nil {
			panic(err)
		}

		resBuf, err := service1Cli.GenericCall(context.Background(), methodName, reqBuf)

		_, _, err = rc.Decode(resBuf.([]byte), &res)
		if err != nil {
			panic(err)
		}
		
		log.Println(res)
		ctx.JSON(consts.StatusOK, res)
	})

	register(h)
	h.Spin()
}
