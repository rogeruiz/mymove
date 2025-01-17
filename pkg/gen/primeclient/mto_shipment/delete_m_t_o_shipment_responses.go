// Code generated by go-swagger; DO NOT EDIT.

package mto_shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/transcom/mymove/pkg/gen/primemessages"
)

// DeleteMTOShipmentReader is a Reader for the DeleteMTOShipment structure.
type DeleteMTOShipmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMTOShipmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteMTOShipmentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteMTOShipmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteMTOShipmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteMTOShipmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteMTOShipmentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteMTOShipmentUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteMTOShipmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteMTOShipmentNoContent creates a DeleteMTOShipmentNoContent with default headers values
func NewDeleteMTOShipmentNoContent() *DeleteMTOShipmentNoContent {
	return &DeleteMTOShipmentNoContent{}
}

/* DeleteMTOShipmentNoContent describes a response with status code 204, with default header values.

Successfully deleted the MTO shipment.
*/
type DeleteMTOShipmentNoContent struct {
}

func (o *DeleteMTOShipmentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /mto-shipments/{mtoShipmentID}][%d] deleteMTOShipmentNoContent ", 204)
}

func (o *DeleteMTOShipmentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMTOShipmentBadRequest creates a DeleteMTOShipmentBadRequest with default headers values
func NewDeleteMTOShipmentBadRequest() *DeleteMTOShipmentBadRequest {
	return &DeleteMTOShipmentBadRequest{}
}

/* DeleteMTOShipmentBadRequest describes a response with status code 400, with default header values.

The request payload is invalid.
*/
type DeleteMTOShipmentBadRequest struct {
	Payload *primemessages.ClientError
}

func (o *DeleteMTOShipmentBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /mto-shipments/{mtoShipmentID}][%d] deleteMTOShipmentBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteMTOShipmentBadRequest) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *DeleteMTOShipmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMTOShipmentForbidden creates a DeleteMTOShipmentForbidden with default headers values
func NewDeleteMTOShipmentForbidden() *DeleteMTOShipmentForbidden {
	return &DeleteMTOShipmentForbidden{}
}

/* DeleteMTOShipmentForbidden describes a response with status code 403, with default header values.

The request was denied.
*/
type DeleteMTOShipmentForbidden struct {
	Payload *primemessages.ClientError
}

func (o *DeleteMTOShipmentForbidden) Error() string {
	return fmt.Sprintf("[DELETE /mto-shipments/{mtoShipmentID}][%d] deleteMTOShipmentForbidden  %+v", 403, o.Payload)
}
func (o *DeleteMTOShipmentForbidden) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *DeleteMTOShipmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMTOShipmentNotFound creates a DeleteMTOShipmentNotFound with default headers values
func NewDeleteMTOShipmentNotFound() *DeleteMTOShipmentNotFound {
	return &DeleteMTOShipmentNotFound{}
}

/* DeleteMTOShipmentNotFound describes a response with status code 404, with default header values.

The requested resource wasn't found.
*/
type DeleteMTOShipmentNotFound struct {
	Payload *primemessages.ClientError
}

func (o *DeleteMTOShipmentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /mto-shipments/{mtoShipmentID}][%d] deleteMTOShipmentNotFound  %+v", 404, o.Payload)
}
func (o *DeleteMTOShipmentNotFound) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *DeleteMTOShipmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMTOShipmentConflict creates a DeleteMTOShipmentConflict with default headers values
func NewDeleteMTOShipmentConflict() *DeleteMTOShipmentConflict {
	return &DeleteMTOShipmentConflict{}
}

/* DeleteMTOShipmentConflict describes a response with status code 409, with default header values.

The request could not be processed because of conflict in the current state of the resource.
*/
type DeleteMTOShipmentConflict struct {
	Payload *primemessages.ClientError
}

func (o *DeleteMTOShipmentConflict) Error() string {
	return fmt.Sprintf("[DELETE /mto-shipments/{mtoShipmentID}][%d] deleteMTOShipmentConflict  %+v", 409, o.Payload)
}
func (o *DeleteMTOShipmentConflict) GetPayload() *primemessages.ClientError {
	return o.Payload
}

func (o *DeleteMTOShipmentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ClientError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMTOShipmentUnprocessableEntity creates a DeleteMTOShipmentUnprocessableEntity with default headers values
func NewDeleteMTOShipmentUnprocessableEntity() *DeleteMTOShipmentUnprocessableEntity {
	return &DeleteMTOShipmentUnprocessableEntity{}
}

/* DeleteMTOShipmentUnprocessableEntity describes a response with status code 422, with default header values.

The request was unprocessable, likely due to bad input from the requester.
*/
type DeleteMTOShipmentUnprocessableEntity struct {
	Payload *primemessages.ValidationError
}

func (o *DeleteMTOShipmentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /mto-shipments/{mtoShipmentID}][%d] deleteMTOShipmentUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *DeleteMTOShipmentUnprocessableEntity) GetPayload() *primemessages.ValidationError {
	return o.Payload
}

func (o *DeleteMTOShipmentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMTOShipmentInternalServerError creates a DeleteMTOShipmentInternalServerError with default headers values
func NewDeleteMTOShipmentInternalServerError() *DeleteMTOShipmentInternalServerError {
	return &DeleteMTOShipmentInternalServerError{}
}

/* DeleteMTOShipmentInternalServerError describes a response with status code 500, with default header values.

A server error occurred.
*/
type DeleteMTOShipmentInternalServerError struct {
	Payload *primemessages.Error
}

func (o *DeleteMTOShipmentInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /mto-shipments/{mtoShipmentID}][%d] deleteMTOShipmentInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteMTOShipmentInternalServerError) GetPayload() *primemessages.Error {
	return o.Payload
}

func (o *DeleteMTOShipmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(primemessages.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
