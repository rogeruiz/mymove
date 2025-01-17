// Code generated by go-swagger; DO NOT EDIT.

package move

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateMoveHandlerFunc turns a function with the right signature into a update move handler
type UpdateMoveHandlerFunc func(UpdateMoveParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateMoveHandlerFunc) Handle(params UpdateMoveParams) middleware.Responder {
	return fn(params)
}

// UpdateMoveHandler interface for that can handle valid update move params
type UpdateMoveHandler interface {
	Handle(UpdateMoveParams) middleware.Responder
}

// NewUpdateMove creates a new http.Handler for the update move operation
func NewUpdateMove(ctx *middleware.Context, handler UpdateMoveHandler) *UpdateMove {
	return &UpdateMove{Context: ctx, Handler: handler}
}

/* UpdateMove swagger:route PATCH /moves/{moveID} move updateMove

Disables or re-enables a move

Allows the user to change the `show` field on the selected field to either `True` or `False`. A "shown" move will appear to all users as normal, a "hidden" move will not be returned or editable using any other endpoint (besides those in the Support API), and thus effectively deactivated.


*/
type UpdateMove struct {
	Context *middleware.Context
	Handler UpdateMoveHandler
}

func (o *UpdateMove) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateMoveParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
