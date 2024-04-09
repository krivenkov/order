// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/krivenkov/order/internal/server/http/models"
)

// GetOrderOKCode is the HTTP code returned for type GetOrderOK
const GetOrderOKCode int = 200

/*
GetOrderOK OK

swagger:response getOrderOK
*/
type GetOrderOK struct {

	/*
	  In: Body
	*/
	Payload *models.GetOrderResponse `json:"body,omitempty"`
}

// NewGetOrderOK creates GetOrderOK with default headers values
func NewGetOrderOK() *GetOrderOK {

	return &GetOrderOK{}
}

// WithPayload adds the payload to the get order o k response
func (o *GetOrderOK) WithPayload(payload *models.GetOrderResponse) *GetOrderOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get order o k response
func (o *GetOrderOK) SetPayload(payload *models.GetOrderResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOrderOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetOrderUnauthorizedCode is the HTTP code returned for type GetOrderUnauthorized
const GetOrderUnauthorizedCode int = 401

/*
GetOrderUnauthorized Unauthorized

swagger:response getOrderUnauthorized
*/
type GetOrderUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetOrderUnauthorized creates GetOrderUnauthorized with default headers values
func NewGetOrderUnauthorized() *GetOrderUnauthorized {

	return &GetOrderUnauthorized{}
}

// WithPayload adds the payload to the get order unauthorized response
func (o *GetOrderUnauthorized) WithPayload(payload *models.Error) *GetOrderUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get order unauthorized response
func (o *GetOrderUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOrderUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetOrderForbiddenCode is the HTTP code returned for type GetOrderForbidden
const GetOrderForbiddenCode int = 403

/*
GetOrderForbidden Forbidden

swagger:response getOrderForbidden
*/
type GetOrderForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetOrderForbidden creates GetOrderForbidden with default headers values
func NewGetOrderForbidden() *GetOrderForbidden {

	return &GetOrderForbidden{}
}

// WithPayload adds the payload to the get order forbidden response
func (o *GetOrderForbidden) WithPayload(payload *models.Error) *GetOrderForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get order forbidden response
func (o *GetOrderForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOrderForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetOrderNotFoundCode is the HTTP code returned for type GetOrderNotFound
const GetOrderNotFoundCode int = 404

/*
GetOrderNotFound Not Found

swagger:response getOrderNotFound
*/
type GetOrderNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetOrderNotFound creates GetOrderNotFound with default headers values
func NewGetOrderNotFound() *GetOrderNotFound {

	return &GetOrderNotFound{}
}

// WithPayload adds the payload to the get order not found response
func (o *GetOrderNotFound) WithPayload(payload *models.Error) *GetOrderNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get order not found response
func (o *GetOrderNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOrderNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetOrderInternalServerErrorCode is the HTTP code returned for type GetOrderInternalServerError
const GetOrderInternalServerErrorCode int = 500

/*
GetOrderInternalServerError Internal Server Error

swagger:response getOrderInternalServerError
*/
type GetOrderInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetOrderInternalServerError creates GetOrderInternalServerError with default headers values
func NewGetOrderInternalServerError() *GetOrderInternalServerError {

	return &GetOrderInternalServerError{}
}

// WithPayload adds the payload to the get order internal server error response
func (o *GetOrderInternalServerError) WithPayload(payload *models.Error) *GetOrderInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get order internal server error response
func (o *GetOrderInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOrderInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}