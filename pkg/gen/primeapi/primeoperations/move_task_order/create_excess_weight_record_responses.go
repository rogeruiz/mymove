// Code generated by go-swagger; DO NOT EDIT.

package move_task_order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/transcom/mymove/pkg/gen/primemessages"
)

// CreateExcessWeightRecordCreatedCode is the HTTP code returned for type CreateExcessWeightRecordCreated
const CreateExcessWeightRecordCreatedCode int = 201

/*CreateExcessWeightRecordCreated Successfully uploaded the excess weight record file.

swagger:response createExcessWeightRecordCreated
*/
type CreateExcessWeightRecordCreated struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ExcessWeightRecord `json:"body,omitempty"`
}

// NewCreateExcessWeightRecordCreated creates CreateExcessWeightRecordCreated with default headers values
func NewCreateExcessWeightRecordCreated() *CreateExcessWeightRecordCreated {

	return &CreateExcessWeightRecordCreated{}
}

// WithPayload adds the payload to the create excess weight record created response
func (o *CreateExcessWeightRecordCreated) WithPayload(payload *primemessages.ExcessWeightRecord) *CreateExcessWeightRecordCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create excess weight record created response
func (o *CreateExcessWeightRecordCreated) SetPayload(payload *primemessages.ExcessWeightRecord) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateExcessWeightRecordCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateExcessWeightRecordUnauthorizedCode is the HTTP code returned for type CreateExcessWeightRecordUnauthorized
const CreateExcessWeightRecordUnauthorizedCode int = 401

/*CreateExcessWeightRecordUnauthorized The request was denied.

swagger:response createExcessWeightRecordUnauthorized
*/
type CreateExcessWeightRecordUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ClientError `json:"body,omitempty"`
}

// NewCreateExcessWeightRecordUnauthorized creates CreateExcessWeightRecordUnauthorized with default headers values
func NewCreateExcessWeightRecordUnauthorized() *CreateExcessWeightRecordUnauthorized {

	return &CreateExcessWeightRecordUnauthorized{}
}

// WithPayload adds the payload to the create excess weight record unauthorized response
func (o *CreateExcessWeightRecordUnauthorized) WithPayload(payload *primemessages.ClientError) *CreateExcessWeightRecordUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create excess weight record unauthorized response
func (o *CreateExcessWeightRecordUnauthorized) SetPayload(payload *primemessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateExcessWeightRecordUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateExcessWeightRecordForbiddenCode is the HTTP code returned for type CreateExcessWeightRecordForbidden
const CreateExcessWeightRecordForbiddenCode int = 403

/*CreateExcessWeightRecordForbidden The request was denied.

swagger:response createExcessWeightRecordForbidden
*/
type CreateExcessWeightRecordForbidden struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ClientError `json:"body,omitempty"`
}

// NewCreateExcessWeightRecordForbidden creates CreateExcessWeightRecordForbidden with default headers values
func NewCreateExcessWeightRecordForbidden() *CreateExcessWeightRecordForbidden {

	return &CreateExcessWeightRecordForbidden{}
}

// WithPayload adds the payload to the create excess weight record forbidden response
func (o *CreateExcessWeightRecordForbidden) WithPayload(payload *primemessages.ClientError) *CreateExcessWeightRecordForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create excess weight record forbidden response
func (o *CreateExcessWeightRecordForbidden) SetPayload(payload *primemessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateExcessWeightRecordForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateExcessWeightRecordNotFoundCode is the HTTP code returned for type CreateExcessWeightRecordNotFound
const CreateExcessWeightRecordNotFoundCode int = 404

/*CreateExcessWeightRecordNotFound The requested resource wasn't found.

swagger:response createExcessWeightRecordNotFound
*/
type CreateExcessWeightRecordNotFound struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ClientError `json:"body,omitempty"`
}

// NewCreateExcessWeightRecordNotFound creates CreateExcessWeightRecordNotFound with default headers values
func NewCreateExcessWeightRecordNotFound() *CreateExcessWeightRecordNotFound {

	return &CreateExcessWeightRecordNotFound{}
}

// WithPayload adds the payload to the create excess weight record not found response
func (o *CreateExcessWeightRecordNotFound) WithPayload(payload *primemessages.ClientError) *CreateExcessWeightRecordNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create excess weight record not found response
func (o *CreateExcessWeightRecordNotFound) SetPayload(payload *primemessages.ClientError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateExcessWeightRecordNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateExcessWeightRecordUnprocessableEntityCode is the HTTP code returned for type CreateExcessWeightRecordUnprocessableEntity
const CreateExcessWeightRecordUnprocessableEntityCode int = 422

/*CreateExcessWeightRecordUnprocessableEntity The request was unprocessable, likely due to bad input from the requester.

swagger:response createExcessWeightRecordUnprocessableEntity
*/
type CreateExcessWeightRecordUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *primemessages.ValidationError `json:"body,omitempty"`
}

// NewCreateExcessWeightRecordUnprocessableEntity creates CreateExcessWeightRecordUnprocessableEntity with default headers values
func NewCreateExcessWeightRecordUnprocessableEntity() *CreateExcessWeightRecordUnprocessableEntity {

	return &CreateExcessWeightRecordUnprocessableEntity{}
}

// WithPayload adds the payload to the create excess weight record unprocessable entity response
func (o *CreateExcessWeightRecordUnprocessableEntity) WithPayload(payload *primemessages.ValidationError) *CreateExcessWeightRecordUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create excess weight record unprocessable entity response
func (o *CreateExcessWeightRecordUnprocessableEntity) SetPayload(payload *primemessages.ValidationError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateExcessWeightRecordUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateExcessWeightRecordInternalServerErrorCode is the HTTP code returned for type CreateExcessWeightRecordInternalServerError
const CreateExcessWeightRecordInternalServerErrorCode int = 500

/*CreateExcessWeightRecordInternalServerError A server error occurred.

swagger:response createExcessWeightRecordInternalServerError
*/
type CreateExcessWeightRecordInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *primemessages.Error `json:"body,omitempty"`
}

// NewCreateExcessWeightRecordInternalServerError creates CreateExcessWeightRecordInternalServerError with default headers values
func NewCreateExcessWeightRecordInternalServerError() *CreateExcessWeightRecordInternalServerError {

	return &CreateExcessWeightRecordInternalServerError{}
}

// WithPayload adds the payload to the create excess weight record internal server error response
func (o *CreateExcessWeightRecordInternalServerError) WithPayload(payload *primemessages.Error) *CreateExcessWeightRecordInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create excess weight record internal server error response
func (o *CreateExcessWeightRecordInternalServerError) SetPayload(payload *primemessages.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateExcessWeightRecordInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}