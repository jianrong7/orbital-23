// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"context"
	"log"
	"simpleExample/kitex_gen/api"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/utils"
)

type JSONRawBody struct {
	RawBody string `raw_body:""`
}

func main() {
	h := server.Default()

	genericCli, err := genericclient.NewClient("example", generic.BinaryThriftGeneric(), client.WithHostPorts("127.0.0.1:8080"))
	if err != nil {
		panic(err)
	}

	rc := utils.NewThriftMessageCodec()

	h.POST("/handle", func(c context.Context, ctx *app.RequestContext) {
		var req api.AddRequest
		var res api.AddResponse
		err := ctx.BindAndValidate(&req)
		if err != nil {
			panic(err)
		}

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
		ctx.JSON(200, res)
	})

	h.Spin()
}
