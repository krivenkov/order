// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateOrderHandlerFunc turns a function with the right signature into a create order handler
type CreateOrderHandlerFunc func(CreateOrderParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateOrderHandlerFunc) Handle(params CreateOrderParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateOrderHandler interface for that can handle valid create order params
type CreateOrderHandler interface {
	Handle(CreateOrderParams, interface{}) middleware.Responder
}

// NewCreateOrder creates a new http.Handler for the create order operation
func NewCreateOrder(ctx *middleware.Context, handler CreateOrderHandler) *CreateOrder {
	return &CreateOrder{Context: ctx, Handler: handler}
}

/*
	CreateOrder swagger:route POST / order createOrder

Create new order
*/
type CreateOrder struct {
	Context *middleware.Context
	Handler CreateOrderHandler
}

func (o *CreateOrder) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateOrderParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
