package app

import (
	apiHello "github.com/delivery-club/go-swagger-example/example2/internal/generated/restapi/operations/hello"
	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) HelloWorldFullHandler(params apiHello.HelloWorldFullParams) middleware.Responder {
	return middleware.NotImplemented("operation hello HelloWorldFull has not yet been implemented")
}
