// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ModelAgentsHandlerFunc turns a function with the right signature into a model agents handler
type ModelAgentsHandlerFunc func(ModelAgentsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ModelAgentsHandlerFunc) Handle(params ModelAgentsParams) middleware.Responder {
	return fn(params)
}

// ModelAgentsHandler interface for that can handle valid model agents params
type ModelAgentsHandler interface {
	Handle(ModelAgentsParams) middleware.Responder
}

// NewModelAgents creates a new http.Handler for the model agents operation
func NewModelAgents(ctx *middleware.Context, handler ModelAgentsHandler) *ModelAgents {
	return &ModelAgents{Context: ctx, Handler: handler}
}

/*ModelAgents swagger:route GET /v1/registry/models/agent Registry modelAgents

ModelAgents model agents API

*/
type ModelAgents struct {
	Context *middleware.Context
	Handler ModelAgentsHandler
}

func (o *ModelAgents) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewModelAgentsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}