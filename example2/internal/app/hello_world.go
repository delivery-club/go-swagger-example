package app

import (
	apiHello "github.com/delivery-club/go-swagger-example/example2/internal/generated/restapi/operations/hello"
	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) HelloWorldHandler(params apiHello.HelloWorldParams) middleware.Responder {
	return middleware.NotImplemented("operation hello HelloWorld has not yet been implemented")
}
