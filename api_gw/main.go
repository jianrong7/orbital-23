package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	hclient "github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	kclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/cors"
	hertzZerolog "github.com/hertz-contrib/logger/zerolog"
	hconsul "github.com/hertz-contrib/registry/consul"
	jsoniter "github.com/json-iterator/go"
	kconsul "github.com/kitex-contrib/registry-consul"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)
func genErrResp(ctx *app.RequestContext, statusCode int, err error) {
	ctx.JSON(statusCode, utils.H{
		"error": utils.H{
			"code": statusCode,
			"message": err.Error(),
		},
	})
}
func genSucResp(ctx *app.RequestContext, res interface{}) {
	ctx.JSON(consts.StatusOK, utils.H{
		"data": res,
	})
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

	r, err := kconsul.NewConsulResolver("172.31.26.48:8500")
	if err != nil {
		hlog.Error("Problem adding Consul Resolver (Kitex)")
		panic(err)
	}

	serviceMap := make(map[string]map[string]string)

	// serviceMap is a nested map, the outer layer key is the serviceName,
	// and the inner layer key is the serviceVersion, where the value is the name of the thrift file
	/*
		JSON Mapping as follows:
		{
			"serviceName1" :
			{
				"versionNumber1" : "thriftFileName1",
				"versionNumber2" : "thriftFileName2"
			},
			"serviceName2" :
			{
				"versionNumber1" : "thriftFileName3"
			}
		}

		****** All thrift file names must be unique ******
	*/

	// TODO: add basicauth to this post request path, to prevent malicious attacks on the API Gateway
	h.POST("/idlmanagement/update", LoggerMiddleware(), func(c context.Context, ctx *app.RequestContext) {
		serviceMap = make(map[string]map[string]string)         // reallocate serviceMap to clear all entries
		err = jsoniter.Unmarshal(ctx.GetRawData(), &serviceMap) // receive the new serviceMap from IDL management service in JSON body
		if err != nil {
			genErrResp(ctx, consts.StatusBadRequest, err)
			hlog.Error("Problem unmarshalling serviceMap")
			panic(err)
		}

		// build a consul client
		consulConfig := consulapi.DefaultConfig()
		consulConfig.Address = "172.31.26.48:8500"
		consulClient, err := consulapi.NewClient(consulConfig)
		if err != nil {
			log.Fatal(err)
			return
		}
		// build a consul resolver with the consul client
		hr := hconsul.NewConsulResolver(consulClient)

		// // build a hertz client with the consul resolver
		cli, err := hclient.NewClient()
		if err != nil {
			panic(err)
		}
		cli.Use(sd.Discovery(hr))

		for serviceName, innerMap := range serviceMap { // download the individual thrift files from the IDL management service using RPC
			for serviceVersion, thriftFileName := range innerMap {
				file, err := os.Create("./thrift_files/" + thriftFileName)
				if err != nil {
					hlog.Error("Problem creating new thrift file: " + thriftFileName)
					panic(err)
				}
				address := "http://idlmanagement/getthriftfile/" + thriftFileName
				var content []byte
				status, urlbody, err := cli.Get(context.Background(), content, address, config.WithSD(true))
				// log.Println(urlbody)
				// log.Println(content)
				log.Printf("Status: %d \n", status)
				if err != nil {
					hlog.Error(err)
				}
				if err != nil {
					hlog.Error("Problem getting thrift file")
					panic(err)
				}
				size, err := file.Write(urlbody)
				defer file.Close()
				log.Printf("Downloaded a file %s with size %d", thriftFileName, size)
				if err != nil {
					hlog.Error("Problem adding thrift file: " + serviceName + " " + serviceVersion)
					panic(err)
				}
			}
		}

		log.Println("Updated services")
		ctx.JSON(consts.StatusOK, &serviceMap)
	})

	h.POST("/:service/:version/:method", LoggerMiddleware(), func(c context.Context, ctx *app.RequestContext) {
		serviceName := ctx.Param("service") // see https://www.cloudwego.io/docs/hertz/tutorials/basic-feature/route/
		serviceVersion := ctx.Param("version")
		methodName := cases.Title(language.English, cases.NoLower).String(ctx.Param("method"))
		log.Println(serviceName, serviceVersion, methodName)

		thriftFileDir := "./thrift_files/" + serviceMap[serviceName][serviceVersion]

		// Process RPC Call

		p, err := generic.NewThriftFileProvider(thriftFileDir)
		if err != nil {
			hlog.Error(fmt.Errorf("missing thrift file for service %s, version %s", serviceName, serviceVersion))
			genErrResp(ctx, consts.StatusNotFound, fmt.Errorf("missing thrift file for service %s, version %s", serviceName, serviceVersion))
			return
		}

		g, err := generic.JSONThriftGeneric(p)
		if err != nil {
			hlog.Error("Problem creating new JSONThriftGeneric")
			genErrResp(ctx, consts.StatusInternalServerError, err)
			return
		}

		rpcClient, err := genericclient.NewClient(serviceName+serviceVersion, g, kclient.WithResolver(r), kclient.WithRPCTimeout(time.Second*3))
		if err != nil {
			hlog.Error("Problem creating new generic client")
			genErrResp(ctx, consts.StatusInternalServerError, err)
			return
		}
		
		res, err := rpcClient.GenericCall(context.Background(), methodName, string(ctx.GetRawData()))
		if err != nil {
			hlog.Error(err)
			genErrResp(ctx, consts.StatusNotFound, err)
			return
		}

		err = jsoniter.UnmarshalFromString(res.(string), &res)
		if err != nil {
			hlog.Error("Problem with unmarshalling")
			genErrResp(ctx, consts.StatusInternalServerError, err)
			return
		}

		genSucResp(ctx, res)
	})

	register(h)
	h.Spin()
}
