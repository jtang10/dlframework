// Code generated by go-swagger; DO NOT EDIT.

package predict

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CloseHandlerFunc turns a function with the right signature into a close handler
type CloseHandlerFunc func(CloseParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CloseHandlerFunc) Handle(params CloseParams) middleware.Responder {
	return fn(params)
}

// CloseHandler interface for that can handle valid close params
type CloseHandler interface {
	Handle(CloseParams) middleware.Responder
}

// NewClose creates a new http.Handler for the close operation
func NewClose(ctx *middleware.Context, handler CloseHandler) *Close {
	return &Close{Context: ctx, Handler: handler}
}

/*Close swagger:route POST /v1/predict/close Predict close

Close a predictor clear it's memory.

*/
type Close struct {
	Context *middleware.Context
	Handler CloseHandler
}

func (o *Close) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCloseParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}