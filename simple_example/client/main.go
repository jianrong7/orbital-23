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
	"fmt"
	"simpleExample/client/routes"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
)

type JSONRawBody struct {
	RawBody string `raw_body:""`
}
func ServerMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		fmt.Println("server middleware")
		c.Next(ctx)
	}
}

func main() {
	h := server.Default()
	h.Use(ServerMiddleware())
	
	genericCli, err := genericclient.NewClient("example", generic.BinaryThriftGeneric(), client.WithHostPorts("127.0.0.1:8080"))
	if err != nil {
		panic(err)
	}

	routes.Add(h, genericCli)

	h.Spin()
}
