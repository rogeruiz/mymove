// Code generated by go-swagger; DO NOT EDIT.

package payment_request

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/transcom/mymove/pkg/gen/supportmessages"
)

// RecalculatePaymentRequestReader is a Reader for the RecalculatePaymentRequest structure.
type RecalculatePaymentRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RecalculatePaymentRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRecalculatePaymentRequestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRecalculatePaymentRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRecalculatePaymentRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRecalculatePaymentRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRecalculatePaymentRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRecalculatePaymentRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewRecalculatePaymentRequestPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRecalculatePaymentRequestUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRecalculatePaymentRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRecalculatePaymentRequestCreated creates a RecalculatePaymentRequestCreated with default headers values
func NewRecalculatePaymentRequestCreated() *RecalculatePaymentRequestCreated {
	return &RecalculatePaymentRequestCreated{}
}

/* RecalculatePaymentRequestCreated describes a response with status code 201, with default header values.

The new payment request with recalculated pricing.
*/
type RecalculatePaymentRequestCreated struct {
	Payload *supportmessages.PaymentRequest
}

func (o *RecalculatePaymentRequestCreated) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/recalculate][%d] recalculatePaymentRequestCreated  %+v", 201, o.Payload)
}
func (o *RecalculatePaymentRequestCreated) GetPayload() *supportmessages.PaymentRequest {
	return o.Payload
}

func (o *RecalculatePaymentRequestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.PaymentRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRecalculatePaymentRequestBadRequest creates a RecalculatePaymentRequestBadRequest with default headers values
func NewRecalculatePaymentRequestBadRequest() *RecalculatePaymentRequestBadRequest {
	return &RecalculatePaymentRequestBadRequest{}
}

/* RecalculatePaymentRequestBadRequest describes a response with status code 400, with default header values.

The request payload is invalid.
*/
type RecalculatePaymentRequestBadRequest struct {
	Payload *supportmessages.ClientError
}

func (o *RecalculatePaymentRequestBadRequest) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/recalculate][%d] recalculatePaymentRequestBadRequest  %+v", 400, o.Payload)
}
func (o *RecalculatePaymentRequestBadRequest) GetPayload() *supportmessages.ClientError {
	return o.Payload
}

func (o *RecalculatePaymentRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRecalculatePaymentRequestUnauthorized creates a RecalculatePaymentRequestUnauthorized with default headers values
func NewRecalculatePaymentRequestUnauthorized() *RecalculatePaymentRequestUnauthorized {
	return &RecalculatePaymentRequestUnauthorized{}
}

/* RecalculatePaymentRequestUnauthorized describes a response with status code 401, with default header values.

The request was denied.
*/
type RecalculatePaymentRequestUnauthorized struct {
	Payload *supportmessages.ClientError
}

func (o *RecalculatePaymentRequestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/recalculate][%d] recalculatePaymentRequestUnauthorized  %+v", 401, o.Payload)
}
func (o *RecalculatePaymentRequestUnauthorized) GetPayload() *supportmessages.ClientError {
	return o.Payload
}

func (o *RecalculatePaymentRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRecalculatePaymentRequestForbidden creates a RecalculatePaymentRequestForbidden with default headers values
func NewRecalculatePaymentRequestForbidden() *RecalculatePaymentRequestForbidden {
	return &RecalculatePaymentRequestForbidden{}
}

/* RecalculatePaymentRequestForbidden describes a response with status code 403, with default header values.

The request was denied.
*/
type RecalculatePaymentRequestForbidden struct {
	Payload *supportmessages.ClientError
}

func (o *RecalculatePaymentRequestForbidden) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/recalculate][%d] recalculatePaymentRequestForbidden  %+v", 403, o.Payload)
}
func (o *RecalculatePaymentRequestForbidden) GetPayload() *supportmessages.ClientError {
	return o.Payload
}

func (o *RecalculatePaymentRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRecalculatePaymentRequestNotFound creates a RecalculatePaymentRequestNotFound with default headers values
func NewRecalculatePaymentRequestNotFound() *RecalculatePaymentRequestNotFound {
	return &RecalculatePaymentRequestNotFound{}
}

/* RecalculatePaymentRequestNotFound describes a response with status code 404, with default header values.

The requested resource wasn't found.
*/
type RecalculatePaymentRequestNotFound struct {
	Payload *supportmessages.ClientError
}

func (o *RecalculatePaymentRequestNotFound) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/recalculate][%d] recalculatePaymentRequestNotFound  %+v", 404, o.Payload)
}
func (o *RecalculatePaymentRequestNotFound) GetPayload() *supportmessages.ClientError {
	return o.Payload
}

func (o *RecalculatePaymentRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRecalculatePaymentRequestConflict creates a RecalculatePaymentRequestConflict with default headers values
func NewRecalculatePaymentRequestConflict() *RecalculatePaymentRequestConflict {
	return &RecalculatePaymentRequestConflict{}
}

/* RecalculatePaymentRequestConflict describes a response with status code 409, with default header values.

There was a conflict with the request.
*/
type RecalculatePaymentRequestConflict struct {
	Payload *supportmessages.ClientError
}

func (o *RecalculatePaymentRequestConflict) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/recalculate][%d] recalculatePaymentRequestConflict  %+v", 409, o.Payload)
}
func (o *RecalculatePaymentRequestConflict) GetPayload() *supportmessages.ClientError {
	return o.Payload
}

func (o *RecalculatePaymentRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRecalculatePaymentRequestPreconditionFailed creates a RecalculatePaymentRequestPreconditionFailed with default headers values
func NewRecalculatePaymentRequestPreconditionFailed() *RecalculatePaymentRequestPreconditionFailed {
	return &RecalculatePaymentRequestPreconditionFailed{}
}

/* RecalculatePaymentRequestPreconditionFailed describes a response with status code 412, with default header values.

Precondition failed, likely due to a stale eTag (If-Match). Fetch the request again to get the updated eTag value.
*/
type RecalculatePaymentRequestPreconditionFailed struct {
	Payload *supportmessages.ClientError
}

func (o *RecalculatePaymentRequestPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/recalculate][%d] recalculatePaymentRequestPreconditionFailed  %+v", 412, o.Payload)
}
func (o *RecalculatePaymentRequestPreconditionFailed) GetPayload() *supportmessages.ClientError {
	return o.Payload
}

func (o *RecalculatePaymentRequestPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRecalculatePaymentRequestUnprocessableEntity creates a RecalculatePaymentRequestUnprocessableEntity with default headers values
func NewRecalculatePaymentRequestUnprocessableEntity() *RecalculatePaymentRequestUnprocessableEntity {
	return &RecalculatePaymentRequestUnprocessableEntity{}
}

/* RecalculatePaymentRequestUnprocessableEntity describes a response with status code 422, with default header values.

The payload was unprocessable.
*/
type RecalculatePaymentRequestUnprocessableEntity struct {
	Payload *supportmessages.ValidationError
}

func (o *RecalculatePaymentRequestUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/recalculate][%d] recalculatePaymentRequestUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *RecalculatePaymentRequestUnprocessableEntity) GetPayload() *supportmessages.ValidationError {
	return o.Payload
}

func (o *RecalculatePaymentRequestUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRecalculatePaymentRequestInternalServerError creates a RecalculatePaymentRequestInternalServerError with default headers values
func NewRecalculatePaymentRequestInternalServerError() *RecalculatePaymentRequestInternalServerError {
	return &RecalculatePaymentRequestInternalServerError{}
}

/* RecalculatePaymentRequestInternalServerError describes a response with status code 500, with default header values.

A server error occurred.
*/
type RecalculatePaymentRequestInternalServerError struct {
	Payload *supportmessages.Error
}

func (o *RecalculatePaymentRequestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /payment-requests/{paymentRequestID}/recalculate][%d] recalculatePaymentRequestInternalServerError  %+v", 500, o.Payload)
}
func (o *RecalculatePaymentRequestInternalServerError) GetPayload() *supportmessages.Error {
	return o.Payload
}

func (o *RecalculatePaymentRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(supportmessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
