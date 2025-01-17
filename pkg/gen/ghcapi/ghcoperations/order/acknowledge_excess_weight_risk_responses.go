// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/ghcmessages"
)

// AcknowledgeExcessWeightRiskOKCode is the HTTP code returned for type AcknowledgeExcessWeightRiskOK
const AcknowledgeExcessWeightRiskOKCode int = 200

/*AcknowledgeExcessWeightRiskOK updated Move

swagger:response acknowledgeExcessWeightRiskOK
*/
type AcknowledgeExcessWeightRiskOK struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.Move `json:"body,omitempty"`
}

// NewAcknowledgeExcessWeightRiskOK creates AcknowledgeExcessWeightRiskOK with default headers values
func NewAcknowledgeExcessWeightRiskOK() *AcknowledgeExcessWeightRiskOK {

	return &AcknowledgeExcessWeightRiskOK{}
}

// WithPayload adds the payload to the acknowledge excess weight risk o k response
func (o *AcknowledgeExcessWeightRiskOK) WithPayload(payload *ghcmessages.Move) *AcknowledgeExcessWeightRiskOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the acknowledge excess weight risk o k response
func (o *AcknowledgeExcessWeightRiskOK) SetPayload(payload *ghcmessages.Move) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AcknowledgeExcessWeightRiskOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AcknowledgeExcessWeightRiskForbiddenCode is the HTTP code returned for type AcknowledgeExcessWeightRiskForbidden
const AcknowledgeExcessWeightRiskForbiddenCode int = 403

/*AcknowledgeExcessWeightRiskForbidden The request was denied

swagger:response acknowledgeExcessWeightRiskForbidden
*/
type AcknowledgeExcessWeightRiskForbidden struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.Error `json:"body,omitempty"`
}

// NewAcknowledgeExcessWeightRiskForbidden creates AcknowledgeExcessWeightRiskForbidden with default headers values
func NewAcknowledgeExcessWeightRiskForbidden() *AcknowledgeExcessWeightRiskForbidden {

	return &AcknowledgeExcessWeightRiskForbidden{}
}

// WithPayload adds the payload to the acknowledge excess weight risk forbidden response
func (o *AcknowledgeExcessWeightRiskForbidden) WithPayload(payload *ghcmessages.Error) *AcknowledgeExcessWeightRiskForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the acknowledge excess weight risk forbidden response
func (o *AcknowledgeExcessWeightRiskForbidden) SetPayload(payload *ghcmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AcknowledgeExcessWeightRiskForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AcknowledgeExcessWeightRiskNotFoundCode is the HTTP code returned for type AcknowledgeExcessWeightRiskNotFound
const AcknowledgeExcessWeightRiskNotFoundCode int = 404

/*AcknowledgeExcessWeightRiskNotFound The requested resource wasn't found

swagger:response acknowledgeExcessWeightRiskNotFound
*/
type AcknowledgeExcessWeightRiskNotFound struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.Error `json:"body,omitempty"`
}

// NewAcknowledgeExcessWeightRiskNotFound creates AcknowledgeExcessWeightRiskNotFound with default headers values
func NewAcknowledgeExcessWeightRiskNotFound() *AcknowledgeExcessWeightRiskNotFound {

	return &AcknowledgeExcessWeightRiskNotFound{}
}

// WithPayload adds the payload to the acknowledge excess weight risk not found response
func (o *AcknowledgeExcessWeightRiskNotFound) WithPayload(payload *ghcmessages.Error) *AcknowledgeExcessWeightRiskNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the acknowledge excess weight risk not found response
func (o *AcknowledgeExcessWeightRiskNotFound) SetPayload(payload *ghcmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AcknowledgeExcessWeightRiskNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AcknowledgeExcessWeightRiskPreconditionFailedCode is the HTTP code returned for type AcknowledgeExcessWeightRiskPreconditionFailed
const AcknowledgeExcessWeightRiskPreconditionFailedCode int = 412

/*AcknowledgeExcessWeightRiskPreconditionFailed Precondition failed

swagger:response acknowledgeExcessWeightRiskPreconditionFailed
*/
type AcknowledgeExcessWeightRiskPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.Error `json:"body,omitempty"`
}

// NewAcknowledgeExcessWeightRiskPreconditionFailed creates AcknowledgeExcessWeightRiskPreconditionFailed with default headers values
func NewAcknowledgeExcessWeightRiskPreconditionFailed() *AcknowledgeExcessWeightRiskPreconditionFailed {

	return &AcknowledgeExcessWeightRiskPreconditionFailed{}
}

// WithPayload adds the payload to the acknowledge excess weight risk precondition failed response
func (o *AcknowledgeExcessWeightRiskPreconditionFailed) WithPayload(payload *ghcmessages.Error) *AcknowledgeExcessWeightRiskPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the acknowledge excess weight risk precondition failed response
func (o *AcknowledgeExcessWeightRiskPreconditionFailed) SetPayload(payload *ghcmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AcknowledgeExcessWeightRiskPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AcknowledgeExcessWeightRiskUnprocessableEntityCode is the HTTP code returned for type AcknowledgeExcessWeightRiskUnprocessableEntity
const AcknowledgeExcessWeightRiskUnprocessableEntityCode int = 422

/*AcknowledgeExcessWeightRiskUnprocessableEntity The payload was unprocessable.

swagger:response acknowledgeExcessWeightRiskUnprocessableEntity
*/
type AcknowledgeExcessWeightRiskUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.ValidationError `json:"body,omitempty"`
}

// NewAcknowledgeExcessWeightRiskUnprocessableEntity creates AcknowledgeExcessWeightRiskUnprocessableEntity with default headers values
func NewAcknowledgeExcessWeightRiskUnprocessableEntity() *AcknowledgeExcessWeightRiskUnprocessableEntity {

	return &AcknowledgeExcessWeightRiskUnprocessableEntity{}
}

// WithPayload adds the payload to the acknowledge excess weight risk unprocessable entity response
func (o *AcknowledgeExcessWeightRiskUnprocessableEntity) WithPayload(payload *ghcmessages.ValidationError) *AcknowledgeExcessWeightRiskUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the acknowledge excess weight risk unprocessable entity response
func (o *AcknowledgeExcessWeightRiskUnprocessableEntity) SetPayload(payload *ghcmessages.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AcknowledgeExcessWeightRiskUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AcknowledgeExcessWeightRiskInternalServerErrorCode is the HTTP code returned for type AcknowledgeExcessWeightRiskInternalServerError
const AcknowledgeExcessWeightRiskInternalServerErrorCode int = 500

/*AcknowledgeExcessWeightRiskInternalServerError A server error occurred

swagger:response acknowledgeExcessWeightRiskInternalServerError
*/
type AcknowledgeExcessWeightRiskInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *ghcmessages.Error `json:"body,omitempty"`
}

// NewAcknowledgeExcessWeightRiskInternalServerError creates AcknowledgeExcessWeightRiskInternalServerError with default headers values
func NewAcknowledgeExcessWeightRiskInternalServerError() *AcknowledgeExcessWeightRiskInternalServerError {

	return &AcknowledgeExcessWeightRiskInternalServerError{}
}

// WithPayload adds the payload to the acknowledge excess weight risk internal server error response
func (o *AcknowledgeExcessWeightRiskInternalServerError) WithPayload(payload *ghcmessages.Error) *AcknowledgeExcessWeightRiskInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the acknowledge excess weight risk internal server error response
func (o *AcknowledgeExcessWeightRiskInternalServerError) SetPayload(payload *ghcmessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AcknowledgeExcessWeightRiskInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
