package main

import (
	"context"
	"log"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/utils"
	hertzZerolog "github.com/hertz-contrib/logger/zerolog"
	consul "github.com/kitex-contrib/registry-consul"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	h := initHTTPServer()
	hlog.SetLogger(hertzZerolog.New(hertzZerolog.WithTimestamp()))

	r, err := consul.NewConsulResolver(CONSUL_SERVER_ADDR)
	if err != nil {
		log.Println("Problem adding Consul Resolver")
		panic(err)
	}

	rc := utils.NewThriftMessageCodec()

	h.POST("/:service/:method", LoggerMiddleware(), func(c context.Context, ctx *app.RequestContext) {
		serviceName := ctx.Param("service") // see https://www.cloudwego.io/docs/hertz/tutorials/basic-feature/route/
		methodName := cases.Title(language.English, cases.NoLower).String(ctx.Param("method"))

		req, res, err := FillRequestGetResponse(serviceName, methodName, ctx)

		if err != nil {
			log.Println("Problem filling request struct")
			ctx.AbortWithError(consts.StatusBadRequest, err)
			panic(err)
		}
		log.Println(req)

		reqBuf, err := rc.Encode(methodName, thrift.CALL, 1, req)
		if err != nil {
			log.Println("Problem encoding request struct to thrift")
			panic(err)
		}

		rpcClient, err := genericclient.NewClient(serviceName, generic.BinaryThriftGeneric(), client.WithResolver(r), client.WithRPCTimeout(time.Second*10))
		if err != nil {
			log.Println("Problem creating new generic client")
			panic(err)
		}

		resBuf, err := rpcClient.GenericCall(context.Background(), methodName, reqBuf)
		if err != nil {
			log.Println("Problem with generic call")
			panic(err)
		}
		_, _, err = rc.Decode(resBuf.([]byte), res)
		if err != nil {
			log.Println("Problem decoding thrift binary")
			panic(err)
		}

		log.Println(res)
		ctx.JSON(consts.StatusOK, res)
	})

	register(h)
	h.Spin()
}
