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

// func addThriftFile(serviceName string, serviceVersion string, serviceMap map[string]map[string]string, idlmClient idlm.Client) (err error) {
// 	thriftFileName := serviceMap[serviceName][serviceVersion]
// 	file, err := os.Create("./thrift_files/" + thriftFileName)
// 	if err != nil {
// 		hlog.Error("Problem creating new thrift file: " + thriftFileName)
// 		panic(err)
// 	}
// 	content, err := idlmClient.GetThriftFile(context.Background(), serviceName, serviceVersion)
// 	if err != nil {
// 		hlog.Error("Problem getting thrift file")
// 		panic(err)
// 	}
// 	size, err := file.WriteString(content)
// 	defer file.Close()
// 	log.Printf("Downloaded a file %s with size %d", thriftFileName, size)
// 	return err
// }

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8888"))
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
			hlog.Error("Problem unmarshalling serviceMap")
			panic(err)
		}
		log.Println(serviceMap)

		idlmClient, err := idlm.NewClient("idlmanagement", client.WithResolver(r), client.WithRPCTimeout(time.Second*3))
		if err != nil {
			ctx.JSON(consts.StatusInternalServerError, err.Error())
			return
		}

		for serviceName, innerMap := range serviceMap { // download the individual thrift files from the IDL management service using RPC
			for serviceVersion, thriftFileName := range innerMap {
				log.Println("serviceName: " + serviceName)
				log.Println("serviceVersion: " + serviceVersion)
				log.Println("thriftFileName: " + thriftFileName)
				file, err := os.Create("./thrift_files/" + thriftFileName)
				if err != nil {
					hlog.Error("Problem creating new thrift file: " + thriftFileName)
					panic(err)
				}
				content, err := idlmClient.GetThriftFile(context.Background(), thriftFileName)
				if err != nil {
					hlog.Error("Problem getting thrift file")
					panic(err)
				}
				size, err := file.WriteString(content)
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

		thriftFileDir := "./thrift_files/" + serviceMap[serviceName][serviceVersion]

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
