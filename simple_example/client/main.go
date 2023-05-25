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

	"github.com/cloudwego/kitex/client"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"

	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
)

type JSONRawBody struct {
	RawBody string `raw_body:""`
}

func main() {
	// client, err := ex.NewClient("example", client.WithHostPorts("127.0.0.2:8888"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	h := server.Default()

	h.POST("/handle", func(c context.Context, ctx *app.RequestContext) {
		var arg JSONRawBody
		err := ctx.BindAndValidate(&arg)
		if err != nil {
			panic(err)
		}

		p, err := generic.NewThriftFileProvider("../ex.thrift")
		if err != nil {
			panic(err)
		}

		g, err := generic.JSONThriftGeneric(p)
		if err != nil {
			panic(err)
		}

		cli, err := genericclient.NewClient("example", g, client.WithHostPorts("127.0.0.2:8888"))
		if err != nil {
			panic(err)
		}

		log.Println(arg.RawBody)

		resp, err := cli.GenericCall(c, "Add", arg.RawBody)

		// var nums api.AddRequest

		// addReq := &nums
		// addResp, err := client.Add(context.Background(), addReq)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Println(addResp)

		// ctx.String(200, addResp.String())

		log.Println(resp)
		ctx.JSON(200, resp)
	})

	h.Spin()
}
