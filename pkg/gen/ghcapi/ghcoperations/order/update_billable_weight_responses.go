// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// UpdateBillableWeightOKCode is the HTTP code returned for type UpdateBillableWeightOK
const UpdateBillableWeightOKCode int = 200

/*UpdateBillableWeightOK updated Order

swagger:response updateBillableWeightOK
*/
type UpdateBillableWeightOK struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.Order `json:"body,omitempty"`
}

// NewUpdateBillableWeightOK creates UpdateBillableWeightOK with default headers values
func NewUpdateBillableWeightOK() *UpdateBillableWeightOK {

	return &UpdateBillableWeightOK{}
}

// WithPayload adds the payload to the update billable weight o k response
func (o *UpdateBillableWeightOK) WithPayload(payload *ghcmessages.Order) *UpdateBillableWeightOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update billable weight o k response
func (o *UpdateBillableWeightOK) SetPayload(payload *ghcmessages.Order) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBillableWeightOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateBillableWeightForbiddenCode is the HTTP code returned for type UpdateBillableWeightForbidden
const UpdateBillableWeightForbiddenCode int = 403

/*UpdateBillableWeightForbidden The request was denied

swagger:response updateBillableWeightForbidden
*/
type UpdateBillableWeightForbidden struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.Error `json:"body,omitempty"`
}

// NewUpdateBillableWeightForbidden creates UpdateBillableWeightForbidden with default headers values
func NewUpdateBillableWeightForbidden() *UpdateBillableWeightForbidden {

	return &UpdateBillableWeightForbidden{}
}

// WithPayload adds the payload to the update billable weight forbidden response
func (o *UpdateBillableWeightForbidden) WithPayload(payload *ghcmessages.Error) *UpdateBillableWeightForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update billable weight forbidden response
func (o *UpdateBillableWeightForbidden) SetPayload(payload *ghcmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBillableWeightForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateBillableWeightNotFoundCode is the HTTP code returned for type UpdateBillableWeightNotFound
const UpdateBillableWeightNotFoundCode int = 404

/*UpdateBillableWeightNotFound The requested resource wasn't found

swagger:response updateBillableWeightNotFound
*/
type UpdateBillableWeightNotFound struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.Error `json:"body,omitempty"`
}

// NewUpdateBillableWeightNotFound creates UpdateBillableWeightNotFound with default headers values
func NewUpdateBillableWeightNotFound() *UpdateBillableWeightNotFound {

	return &UpdateBillableWeightNotFound{}
}

// WithPayload adds the payload to the update billable weight not found response
func (o *UpdateBillableWeightNotFound) WithPayload(payload *ghcmessages.Error) *UpdateBillableWeightNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update billable weight not found response
func (o *UpdateBillableWeightNotFound) SetPayload(payload *ghcmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBillableWeightNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateBillableWeightPreconditionFailedCode is the HTTP code returned for type UpdateBillableWeightPreconditionFailed
const UpdateBillableWeightPreconditionFailedCode int = 412

/*UpdateBillableWeightPreconditionFailed Precondition failed

swagger:response updateBillableWeightPreconditionFailed
*/
type UpdateBillableWeightPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.Error `json:"body,omitempty"`
}

// NewUpdateBillableWeightPreconditionFailed creates UpdateBillableWeightPreconditionFailed with default headers values
func NewUpdateBillableWeightPreconditionFailed() *UpdateBillableWeightPreconditionFailed {

	return &UpdateBillableWeightPreconditionFailed{}
}

// WithPayload adds the payload to the update billable weight precondition failed response
func (o *UpdateBillableWeightPreconditionFailed) WithPayload(payload *ghcmessages.Error) *UpdateBillableWeightPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update billable weight precondition failed response
func (o *UpdateBillableWeightPreconditionFailed) SetPayload(payload *ghcmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBillableWeightPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateBillableWeightUnprocessableEntityCode is the HTTP code returned for type UpdateBillableWeightUnprocessableEntity
const UpdateBillableWeightUnprocessableEntityCode int = 422

/*UpdateBillableWeightUnprocessableEntity The payload was unprocessable.

swagger:response updateBillableWeightUnprocessableEntity
*/
type UpdateBillableWeightUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.ValidationError `json:"body,omitempty"`
}

// NewUpdateBillableWeightUnprocessableEntity creates UpdateBillableWeightUnprocessableEntity with default headers values
func NewUpdateBillableWeightUnprocessableEntity() *UpdateBillableWeightUnprocessableEntity {

	return &UpdateBillableWeightUnprocessableEntity{}
}

// WithPayload adds the payload to the update billable weight unprocessable entity response
func (o *UpdateBillableWeightUnprocessableEntity) WithPayload(payload *ghcmessages.ValidationError) *UpdateBillableWeightUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update billable weight unprocessable entity response
func (o *UpdateBillableWeightUnprocessableEntity) SetPayload(payload *ghcmessages.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBillableWeightUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateBillableWeightInternalServerErrorCode is the HTTP code returned for type UpdateBillableWeightInternalServerError
const UpdateBillableWeightInternalServerErrorCode int = 500

/*UpdateBillableWeightInternalServerError A server error occurred

swagger:response updateBillableWeightInternalServerError
*/
type UpdateBillableWeightInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.Error `json:"body,omitempty"`
}

// NewUpdateBillableWeightInternalServerError creates UpdateBillableWeightInternalServerError with default headers values
func NewUpdateBillableWeightInternalServerError() *UpdateBillableWeightInternalServerError {

	return &UpdateBillableWeightInternalServerError{}
}

// WithPayload adds the payload to the update billable weight internal server error response
func (o *UpdateBillableWeightInternalServerError) WithPayload(payload *ghcmessages.Error) *UpdateBillableWeightInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update billable weight internal server error response
func (o *UpdateBillableWeightInternalServerError) SetPayload(payload *ghcmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateBillableWeightInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}