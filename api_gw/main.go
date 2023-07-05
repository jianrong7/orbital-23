package main

import (
	"context"
	"log"
	"os"
	"time"

	idlm "api_gw/service_definitions/kitex_gen/idlmanagement/idlmanagement"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/hertz-contrib/cors"
	hertzZerolog "github.com/hertz-contrib/logger/zerolog"
	jsoniter "github.com/json-iterator/go"
	consul "github.com/kitex-contrib/registry-consul"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

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

func main() {
	h := server.Default(server.WithHostPorts("0.0.0.0:8888"))
	h.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		
	}))
	hlog.SetLogger(hertzZerolog.New(hertzZerolog.WithTimestamp()))

	r, err := consul.NewConsulResolver("127.0.0.1:8500")
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
			ctx.JSON(consts.StatusBadRequest, err.Error())
			return
		}

		idlmVersionNumber, _ := idlmClient.CheckVersion(context.Background())

		// if there is a version update on the idlmanagement service, overwrite existing files / download new files
		// if service is not found in the local map, download it
		// if service is not found on the idlmanagement service, throw an error
		// if service is found, set the thriftFileDir
		if versionNumber != idlmVersionNumber || serviceMap[serviceName] == "" {
			serviceMap[serviceName], err = addThriftFile(serviceName, idlmClient)
			if serviceMap[serviceName] == "" {
				hlog.Error("Problem adding thrift file")
				ctx.JSON(consts.StatusBadRequest, err.Error())
				return
			}
			if err != nil {
				hlog.Error("Problem adding thrift file")
				ctx.JSON(consts.StatusBadRequest, err.Error())
				return
			}
			versionNumber = idlmVersionNumber
		}

		thriftFileDir := "./thrift_files/" + serviceMap[serviceName]

		// Process RPC Call

		p, err := generic.NewThriftFileProvider(thriftFileDir)
		if err != nil {
			hlog.Error("Problem adding new thrift file provider")
			ctx.JSON(consts.StatusBadRequest, err.Error())
			return
		}

		g, err := generic.JSONThriftGeneric(p)
		if err != nil {
			hlog.Error("Problem creating new JSONThriftGeneric")
			ctx.JSON(consts.StatusBadRequest, err.Error())
			return
		}

		rpcClient, err := genericclient.NewClient(serviceName, g, client.WithResolver(r), client.WithRPCTimeout(time.Second*3))
		if err != nil {
			hlog.Error("Problem creating new generic client")
			ctx.JSON(consts.StatusInternalServerError, err.Error())
			return
		}

		res, err := rpcClient.GenericCall(context.Background(), methodName, string(ctx.GetRawData()))
		if err != nil {
			hlog.Error("Problem with generic call")
			ctx.JSON(consts.StatusInternalServerError, err.Error())
			return
		}

		err = jsoniter.UnmarshalFromString(res.(string), &res)
		if err != nil {
			hlog.Error("Problem with unmarshalling")
			ctx.JSON(consts.StatusInternalServerError, err.Error())
			return
		}

		// no errors, set result in RequestContext
		log.Println(res)
		ctx.JSON(consts.StatusOK, res)
	})

	register(h)
	h.Spin()
}
