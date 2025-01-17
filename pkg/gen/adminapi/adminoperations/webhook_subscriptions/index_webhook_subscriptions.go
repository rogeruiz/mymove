// Code generated by go-swagger; DO NOT EDIT.

package webhook_subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// IndexWebhookSubscriptionsHandlerFunc turns a function with the right signature into a index webhook subscriptions handler
type IndexWebhookSubscriptionsHandlerFunc func(IndexWebhookSubscriptionsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn IndexWebhookSubscriptionsHandlerFunc) Handle(params IndexWebhookSubscriptionsParams) middleware.Responder {
	return fn(params)
}

// IndexWebhookSubscriptionsHandler interface for that can handle valid index webhook subscriptions params
type IndexWebhookSubscriptionsHandler interface {
	Handle(IndexWebhookSubscriptionsParams) middleware.Responder
}

// NewIndexWebhookSubscriptions creates a new http.Handler for the index webhook subscriptions operation
func NewIndexWebhookSubscriptions(ctx *middleware.Context, handler IndexWebhookSubscriptionsHandler) *IndexWebhookSubscriptions {
	return &IndexWebhookSubscriptions{Context: ctx, Handler: handler}
}

/* IndexWebhookSubscriptions swagger:route GET /webhook_subscriptions webhook_subscriptions indexWebhookSubscriptions

Lists webhook subscriptions

Returns a list of webhook subscriptions

*/
type IndexWebhookSubscriptions struct {
	Context *middleware.Context
	Handler IndexWebhookSubscriptionsHandler
}

func (o *IndexWebhookSubscriptions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewIndexWebhookSubscriptionsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
