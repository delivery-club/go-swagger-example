// Code generated by go-swagger; DO NOT EDIT.

package hello

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// HelloWorldHandlerFunc turns a function with the right signature into a hello world handler
type HelloWorldHandlerFunc func(HelloWorldParams) middleware.Responder

// Handle executing the request and returning a response
func (fn HelloWorldHandlerFunc) Handle(params HelloWorldParams) middleware.Responder {
	return fn(params)
}

// HelloWorldHandler interface for that can handle valid hello world params
type HelloWorldHandler interface {
	Handle(HelloWorldParams) middleware.Responder
}

// NewHelloWorld creates a new http.Handler for the hello world operation
func NewHelloWorld(ctx *middleware.Context, handler HelloWorldHandler) *HelloWorld {
	return &HelloWorld{Context: ctx, Handler: handler}
}

/*HelloWorld swagger:route GET /hello Hello World helloWorld

Example Route

Some description

*/
type HelloWorld struct {
	Context *middleware.Context
	Handler HelloWorldHandler
}

func (o *HelloWorld) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewHelloWorldParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}