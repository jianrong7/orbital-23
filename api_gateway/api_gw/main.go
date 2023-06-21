package main

import (
	"context"
	"log"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/utils"
	consul "github.com/kitex-contrib/registry-consul"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	h := initHTTPServer()

	r, err := consul.NewConsulResolver("13.229.205.99:8500")
	if err != nil {
		log.Fatal(err)
	}

	rc := utils.NewThriftMessageCodec()

	h.POST("/:service/:method", func(c context.Context, ctx *app.RequestContext) {
		serviceName := ctx.Param("service") // see https://www.cloudwego.io/docs/hertz/tutorials/basic-feature/route/
		methodName := cases.Title(language.English, cases.NoLower).String(ctx.Param("method"))

		req, res, err := FillRequestGetResponse(serviceName, methodName, ctx)

		if err != nil {
			log.Println("Problem filling request struct")
			panic(err)
		}

		reqBuf, err := rc.Encode(methodName, thrift.CALL, 1, req)
		if err != nil {
			panic(err)
		}

		rpcClient, err := genericclient.NewClient(serviceName, generic.BinaryThriftGeneric(), client.WithResolver(r), client.WithRPCTimeout(time.Second*3))
		if err != nil {
			panic(err)
		}

		resBuf, err := rpcClient.GenericCall(context.Background(), methodName, reqBuf)
		if err != nil {
			panic(err)
		}
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
