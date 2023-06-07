package main

import (
	"context"
	"log"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/utils"
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

		req, res, err := FillRequestGetResponse(serviceName, methodName, ctx)
		if err != nil {
			log.Println("Problem filling request struct")
			panic(err)
		}

		// err := ctx.BindAndValidate(&req)
		// if err != nil {
		// 	panic(err)
		// }

		reqBuf, err := rc.Encode("Add", thrift.CALL, 1, req)
		if err != nil {
			panic(err)
		}

		resBuf, err := service1Cli.GenericCall(context.Background(), "Add", reqBuf)

		_, _, err = rc.Decode(resBuf.([]byte), res)
		if err != nil {
			panic(err)
		}

		log.Println(res)
		ctx.JSON(consts.StatusOK, res)
	})

	register(h)
	h.Spin()
}
