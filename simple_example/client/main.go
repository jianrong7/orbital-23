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

	"gabrielexample/gabriel_example/kitex_gen/api/example"

	"github.com/cloudwego/kitex/client"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

type Args struct {
	Query      string   `query:"query"`
	QuerySlice []string `query:"q"`
	Path       string   `path:"path"`
	Header     string   `header:"header"`
	Form       string   `form:"form"`
	Json       string   `json:"json"`
	Vd         int      `query:"vd" vd:"$==0||$==1"`
}

func main() {
	client, err := example.NewClient("example", client.WithHostPorts("127.0.0.2:8888"))
	if err != nil {
		log.Fatal(err)
	}
	// for {
	// 	req := &api.Request{Message: "my request"}
	// 	resp, err := client.Echo(context.Background(), req)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(resp)
	// 	time.Sleep(time.Second)
	// 	addReq := &api.AddRequest{First: 20, Second: 20}
	// 	addResp, err := client.Add(context.Background(), addReq)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(addResp)
	// 	time.Sleep(time.Second)
	// }

	h := server.Default()

	h.POST("/handle", func(c context.Context, ctx *app.RequestContext) {
		var arg Args
		err := ctx.BindAndValidate(&arg)
		if err != nil {
			panic(err)
		}
		log.Println(arg)
	})

	h.Spin()
}
