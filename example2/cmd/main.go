package main

import (
	"log"

	"github.com/delivery-club/go-swagger-example/example2/internal/config"
	"github.com/delivery-club/go-swagger-example/example2/internal/generated/restapi"
	"github.com/delivery-club/go-swagger-example/example2/internal/generated/restapi/operations"

	apiHello "github.com/delivery-club/go-swagger-example/example2/internal/generated/restapi/operations/hello"

	"github.com/go-openapi/loads"

	"github.com/delivery-club/go-swagger-example/example2/internal/app"
)

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	srv := app.New()
	api := operations.NewExample2API(swaggerSpec)

	api.HelloHelloWorldHandler = apiHello.HelloWorldHandlerFunc(srv.HelloWorldHandler)
	api.HelloHelloWorldFullHandler = apiHello.HelloWorldFullHandlerFunc(srv.HelloWorldFullHandler)
	api.ServerShutdown = srv.OnShutdown
	server := restapi.NewServer(api)
	defer server.Shutdown()

	cfg, err := config.InitConfig("example2")
	if err != nil {
		log.Fatalln(err)
	}

	server.ConfigureAPI()

	server.Port = cfg.HTTPBindPort
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
