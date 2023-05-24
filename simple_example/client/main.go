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
	"encoding/json"
	"log"
	"simpleExample/kitex_gen/api"

	ex "simpleExample/kitex_gen/api/simpleexample"

	"github.com/cloudwego/kitex/client"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/binding"
)

type TestBind struct {
	A string `raw_body:""`
}

func main() {
	client, err := ex.NewClient("example", client.WithHostPorts("127.0.0.2:8888"))
	if err != nil {
		log.Fatal(err)
	}

	binding.UseStdJSONUnmarshaler()

	h := server.Default()

	h.POST("/handle", func(c context.Context, ctx *app.RequestContext) {
		var arg TestBind
		err := ctx.BindAndValidate(&arg)
		if err != nil {
			panic(err)
		}

		var nums api.AddRequest
		err = json.Unmarshal([]byte(arg.A), &nums)
		if err != nil {
			log.Fatal(err)
		}

		addReq := &nums
		addResp, err := client.Add(context.Background(), addReq)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(addResp)

		ctx.String(200, addResp.String())
	})

	h.Spin()
}
