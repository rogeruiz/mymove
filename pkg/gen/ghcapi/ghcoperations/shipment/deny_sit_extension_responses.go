// Code generated by go-swagger; DO NOT EDIT.

package shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// DenySitExtensionOKCode is the HTTP code returned for type DenySitExtensionOK
const DenySitExtensionOKCode int = 200

/*DenySitExtensionOK Successfully denied a SIT extension

swagger:response denySitExtensionOK
*/
type DenySitExtensionOK struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.MTOShipment `json:"body,omitempty"`
}

// NewDenySitExtensionOK creates DenySitExtensionOK with default headers values
func NewDenySitExtensionOK() *DenySitExtensionOK {

	return &DenySitExtensionOK{}
}

// WithPayload adds the payload to the deny sit extension o k response
func (o *DenySitExtensionOK) WithPayload(payload *ghcmessages.MTOShipment) *DenySitExtensionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the deny sit extension o k response
func (o *DenySitExtensionOK) SetPayload(payload *ghcmessages.MTOShipment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DenySitExtensionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DenySitExtensionForbiddenCode is the HTTP code returned for type DenySitExtensionForbidden
const DenySitExtensionForbiddenCode int = 403

/*DenySitExtensionForbidden The request was denied

swagger:response denySitExtensionForbidden
*/
type DenySitExtensionForbidden struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.Error `json:"body,omitempty"`
}

// NewDenySitExtensionForbidden creates DenySitExtensionForbidden with default headers values
func NewDenySitExtensionForbidden() *DenySitExtensionForbidden {

	return &DenySitExtensionForbidden{}
}

// WithPayload adds the payload to the deny sit extension forbidden response
func (o *DenySitExtensionForbidden) WithPayload(payload *ghcmessages.Error) *DenySitExtensionForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the deny sit extension forbidden response
func (o *DenySitExtensionForbidden) SetPayload(payload *ghcmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DenySitExtensionForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DenySitExtensionNotFoundCode is the HTTP code returned for type DenySitExtensionNotFound
const DenySitExtensionNotFoundCode int = 404

/*DenySitExtensionNotFound The requested resource wasn't found

swagger:response denySitExtensionNotFound
*/
type DenySitExtensionNotFound struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.Error `json:"body,omitempty"`
}

// NewDenySitExtensionNotFound creates DenySitExtensionNotFound with default headers values
func NewDenySitExtensionNotFound() *DenySitExtensionNotFound {

	return &DenySitExtensionNotFound{}
}

// WithPayload adds the payload to the deny sit extension not found response
func (o *DenySitExtensionNotFound) WithPayload(payload *ghcmessages.Error) *DenySitExtensionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the deny sit extension not found response
func (o *DenySitExtensionNotFound) SetPayload(payload *ghcmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DenySitExtensionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DenySitExtensionConflictCode is the HTTP code returned for type DenySitExtensionConflict
const DenySitExtensionConflictCode int = 409

/*DenySitExtensionConflict Conflict error

swagger:response denySitExtensionConflict
*/
type DenySitExtensionConflict struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.Error `json:"body,omitempty"`
}

// NewDenySitExtensionConflict creates DenySitExtensionConflict with default headers values
func NewDenySitExtensionConflict() *DenySitExtensionConflict {

	return &DenySitExtensionConflict{}
}

// WithPayload adds the payload to the deny sit extension conflict response
func (o *DenySitExtensionConflict) WithPayload(payload *ghcmessages.Error) *DenySitExtensionConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the deny sit extension conflict response
func (o *DenySitExtensionConflict) SetPayload(payload *ghcmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DenySitExtensionConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DenySitExtensionPreconditionFailedCode is the HTTP code returned for type DenySitExtensionPreconditionFailed
const DenySitExtensionPreconditionFailedCode int = 412

/*DenySitExtensionPreconditionFailed Precondition failed

swagger:response denySitExtensionPreconditionFailed
*/
type DenySitExtensionPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.Error `json:"body,omitempty"`
}

// NewDenySitExtensionPreconditionFailed creates DenySitExtensionPreconditionFailed with default headers values
func NewDenySitExtensionPreconditionFailed() *DenySitExtensionPreconditionFailed {

	return &DenySitExtensionPreconditionFailed{}
}

// WithPayload adds the payload to the deny sit extension precondition failed response
func (o *DenySitExtensionPreconditionFailed) WithPayload(payload *ghcmessages.Error) *DenySitExtensionPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the deny sit extension precondition failed response
func (o *DenySitExtensionPreconditionFailed) SetPayload(payload *ghcmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DenySitExtensionPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DenySitExtensionUnprocessableEntityCode is the HTTP code returned for type DenySitExtensionUnprocessableEntity
const DenySitExtensionUnprocessableEntityCode int = 422

/*DenySitExtensionUnprocessableEntity The payload was unprocessable.

swagger:response denySitExtensionUnprocessableEntity
*/
type DenySitExtensionUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.ValidationError `json:"body,omitempty"`
}

// NewDenySitExtensionUnprocessableEntity creates DenySitExtensionUnprocessableEntity with default headers values
func NewDenySitExtensionUnprocessableEntity() *DenySitExtensionUnprocessableEntity {

	return &DenySitExtensionUnprocessableEntity{}
}

// WithPayload adds the payload to the deny sit extension unprocessable entity response
func (o *DenySitExtensionUnprocessableEntity) WithPayload(payload *ghcmessages.ValidationError) *DenySitExtensionUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the deny sit extension unprocessable entity response
func (o *DenySitExtensionUnprocessableEntity) SetPayload(payload *ghcmessages.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DenySitExtensionUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DenySitExtensionInternalServerErrorCode is the HTTP code returned for type DenySitExtensionInternalServerError
const DenySitExtensionInternalServerErrorCode int = 500

/*DenySitExtensionInternalServerError A server error occurred

swagger:response denySitExtensionInternalServerError
*/
type DenySitExtensionInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.Error `json:"body,omitempty"`
}

// NewDenySitExtensionInternalServerError creates DenySitExtensionInternalServerError with default headers values
func NewDenySitExtensionInternalServerError() *DenySitExtensionInternalServerError {

	return &DenySitExtensionInternalServerError{}
}

// WithPayload adds the payload to the deny sit extension internal server error response
func (o *DenySitExtensionInternalServerError) WithPayload(payload *ghcmessages.Error) *DenySitExtensionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the deny sit extension internal server error response
func (o *DenySitExtensionInternalServerError) SetPayload(payload *ghcmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DenySitExtensionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}