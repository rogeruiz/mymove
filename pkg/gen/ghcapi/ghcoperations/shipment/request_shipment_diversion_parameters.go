// Code generated by go-swagger; DO NOT EDIT.

package shipment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewRequestShipmentDiversionParams creates a new RequestShipmentDiversionParams object
// no default values defined in spec.
func NewRequestShipmentDiversionParams() RequestShipmentDiversionParams {

	return RequestShipmentDiversionParams{}
}

// RequestShipmentDiversionParams contains all the bound params for the request shipment diversion operation
// typically these are obtained from a http.Request
//
// swagger:parameters requestShipmentDiversion
type RequestShipmentDiversionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: header
	*/
	IfMatch string
	/*ID of the shipment
	  Required: true
	  In: path
	*/
	ShipmentID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRequestShipmentDiversionParams() beforehand.
func (o *RequestShipmentDiversionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindIfMatch(r.Header[http.CanonicalHeaderKey("If-Match")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rShipmentID, rhkShipmentID, _ := route.Params.GetOK("shipmentID")
	if err := o.bindShipmentID(rShipmentID, rhkShipmentID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIfMatch binds and validates parameter IfMatch from header.
func (o *RequestShipmentDiversionParams) bindIfMatch(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("If-Match", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("If-Match", "header", raw); err != nil {
		return err
	}

	o.IfMatch = raw

	return nil
}

// bindShipmentID binds and validates parameter ShipmentID from path.
func (o *RequestShipmentDiversionParams) bindShipmentID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("shipmentID", "path", "strfmt.UUID", raw)
	}
	o.ShipmentID = *(value.(*strfmt.UUID))

	if err := o.validateShipmentID(formats); err != nil {
		return err
	}

	return nil
}

// validateShipmentID carries on validations for parameter ShipmentID
func (o *RequestShipmentDiversionParams) validateShipmentID(formats strfmt.Registry) error {

	if err := validate.FormatOf("shipmentID", "path", "uuid", o.ShipmentID.String(), formats); err != nil {
		return err
	}
	return nil
}