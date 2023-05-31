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

	h.POST("/handle", func(c context.Context, ctx *app.RequestContext) {
		var arg JSONRawBody
		err := ctx.BindAndValidate(&arg)
		if err != nil {
			panic(err)
		}

		genericCli, err := genericclient.NewClient("example", generic.BinaryThriftGeneric(), client.WithHostPorts("127.0.0.1:8080"))
		if err != nil {
			panic(err)
		}

		rc := utils.NewThriftMessageCodec()
		buf, err := rc.Encode("Add", thrift.CALL, 1, &api.SimpleExampleAddArgs{Req: &api.AddRequest{First: 1, Second: 2}})
		if err != nil {
			panic(err)
		}

		resp, err := genericCli.GenericCall(context.Background(), "Add", buf)

		result := &api.AddResponse{}
		_, _, err = rc.Decode(resp.([]byte), result)
		if err != nil {
			panic(err)
		}

		// log.Println(result)
		ctx.JSON(200, result)
	})

	h.Spin()
}
