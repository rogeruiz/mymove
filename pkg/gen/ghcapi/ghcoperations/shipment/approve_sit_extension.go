// Code generated by go-swagger; DO NOT EDIT.

package shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ApproveSitExtensionHandlerFunc turns a function with the right signature into a approve sit extension handler
type ApproveSitExtensionHandlerFunc func(ApproveSitExtensionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ApproveSitExtensionHandlerFunc) Handle(params ApproveSitExtensionParams) middleware.Responder {
	return fn(params)
}

// ApproveSitExtensionHandler interface for that can handle valid approve sit extension params
type ApproveSitExtensionHandler interface {
	Handle(ApproveSitExtensionParams) middleware.Responder
}

// NewApproveSitExtension creates a new http.Handler for the approve sit extension operation
func NewApproveSitExtension(ctx *middleware.Context, handler ApproveSitExtensionHandler) *ApproveSitExtension {
	return &ApproveSitExtension{Context: ctx, Handler: handler}
}

/* ApproveSitExtension swagger:route PATCH /shipments/{shipmentID}/sit-extensions/{sitExtensionID}/approve shipment sitExtension approveSitExtension

Approves a SIT extension

Approves a SIT extension

*/
type ApproveSitExtension struct {
	Context *middleware.Context
	Handler ApproveSitExtensionHandler
}

func (o *ApproveSitExtension) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewApproveSitExtensionParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}