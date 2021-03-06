// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/rai-project/dlframework/httpapi/models"
)

// UpdateHandlerFunc turns a function with the right signature into a update handler
type UpdateHandlerFunc func(UpdateParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateHandlerFunc) Handle(params UpdateParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// UpdateHandler interface for that can handle valid update params
type UpdateHandler interface {
	Handle(UpdateParams, *models.User) middleware.Responder
}

// NewUpdate creates a new http.Handler for the update operation
func NewUpdate(ctx *middleware.Context, handler UpdateHandler) *Update {
	return &Update{Context: ctx, Handler: handler}
}

/*Update swagger:route POST /auth/update Authentication update

Update User Info on MLModelScope platform

*/
type Update struct {
	Context *middleware.Context
	Handler UpdateHandler
}

func (o *Update) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
