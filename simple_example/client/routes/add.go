package routes

import (
	"context"
	"log"
	"simpleExample/kitex_gen/api"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/utils"
)

func Add(h *server.Hertz, genericCli genericclient.Client) {
	add := h.Group("/add")
	add.POST("/post", func(c context.Context, ctx *app.RequestContext) {
	rc := utils.NewThriftMessageCodec()
	var req api.AddRequest
	var res api.AddResponse
	err := ctx.BindAndValidate(&req)
	if err != nil {
		panic(err)
	}

	log.Println(req.First)
	log.Println(req.Second)

	reqBuf, err := rc.Encode("Add", thrift.CALL, 1, &api.AddRequest{First: req.First, Second: req.Second})
	if err != nil {
		panic(err)
	}

	resBuf, err := genericCli.GenericCall(context.Background(), "Add", reqBuf)

	_, _, err = rc.Decode(resBuf.([]byte), &res)
	if err != nil {
		panic(err)
	}

	log.Println(res)
	ctx.JSON(consts.StatusOK, res)
	})
}