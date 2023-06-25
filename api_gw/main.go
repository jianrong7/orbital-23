package main

import (
	"context"
	"log"
	"os"
	"time"

	idlm "api_gw/service_definitions/kitex_gen/idlmanagement/idlmanagement"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	hertzZerolog "github.com/hertz-contrib/logger/zerolog"
	jsoniter "github.com/json-iterator/go"
	consul "github.com/kitex-contrib/registry-consul"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	h := initHTTPServer()
	hlog.SetLogger(hertzZerolog.New(hertzZerolog.WithTimestamp()))

	r, err := consul.NewConsulResolver(CONSUL_SERVER_ADDR)
	if err != nil {
		hlog.Error("Problem adding Consul Resolver")
		panic(err)
	}

	var versionNumber string
	serviceMap := make(map[string]string)

	h.POST("/:service/:method", LoggerMiddleware(), func(c context.Context, ctx *app.RequestContext) {
		serviceName := ctx.Param("service") // see https://www.cloudwego.io/docs/hertz/tutorials/basic-feature/route/
		methodName := cases.Title(language.English, cases.NoLower).String(ctx.Param("method"))

		// check version number with IDL management service
		idlmClient, err := idlm.NewClient("idlmanagement", client.WithResolver(r), client.WithRPCTimeout(time.Second*3))
		if err != nil {
			hlog.Error("Problem creating new idlmanagement client")
			panic(err)
		}
		idlmVersionNumber, _ := idlmClient.CheckVersion(context.Background())

		if versionNumber != idlmVersionNumber || serviceMap[serviceName] == "" {
			serviceMap[serviceName], err = addThriftFile(serviceName, idlmClient)
			if serviceMap[serviceName] == "" {
				ctx.JSON(consts.StatusBadRequest, err.Error())
			}
			if err != nil {
				hlog.Error("Problem adding thrift file")
				panic(err)
			}
			versionNumber = idlmVersionNumber
		}

		log.Println(versionNumber)
		thriftFileDir := "./thrift_files/" + serviceMap[serviceName]

		p, err := generic.NewThriftFileProvider(thriftFileDir)
		if err != nil {
			hlog.Error("Problem adding new thrift file provider")
			panic(err)
		}

		g, err := generic.JSONThriftGeneric(p)
		if err != nil {
			hlog.Error("Problem creating new JSONThriftGeneric")
			panic(err)
		}

		rpcClient, err := genericclient.NewClient(serviceName, g, client.WithResolver(r), client.WithRPCTimeout(time.Second*3))
		if err != nil {
			hlog.Error("Problem creating new generic client")
			panic(err)
		}

		res, err := rpcClient.GenericCall(context.Background(), methodName, string(ctx.GetRawData()))
		if err != nil {
			hlog.Error("Problem with generic call")
			panic(err)
		}

		err = jsoniter.UnmarshalFromString(res.(string), &res)
		if err != nil {
			hlog.Error("Problem with unmarshalling")
			panic(err)
		}
		
		log.Println(res)
		ctx.JSON(consts.StatusOK, res)
	})

	register(h)
	h.Spin()
}

func addThriftFile(serviceName string, idlmClient idlm.Client) (thriftFileName string, err error) {
	thriftFileName, err = idlmClient.GetServiceThriftFileName(context.Background(), serviceName)
	if err != nil {
		return "", err
	}
	file, err := os.Create("./thrift_files/" + thriftFileName)
	if err != nil {
		hlog.Error("Problem creating new thrift file: " + thriftFileName)
		panic(err)
	}
	content, err := idlmClient.GetThriftFile(context.Background(), serviceName)
	if err != nil {
		hlog.Error("Problem getting thrift file")
		panic(err)
	}
	size, err := file.WriteString(content)
	defer file.Close()
	log.Printf("Downloaded a file %s with size %d", thriftFileName, size)
	return thriftFileName, err
}
