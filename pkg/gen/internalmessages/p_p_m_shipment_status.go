// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// PPMShipmentStatus p p m shipment status
//
// swagger:model PPMShipmentStatus
type PPMShipmentStatus string

func NewPPMShipmentStatus(value PPMShipmentStatus) *PPMShipmentStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated PPMShipmentStatus.
func (m PPMShipmentStatus) Pointer() *PPMShipmentStatus {
	return &m
}

const (

	// PPMShipmentStatusDRAFT captures enum value "DRAFT"
	PPMShipmentStatusDRAFT PPMShipmentStatus = "DRAFT"

	// PPMShipmentStatusSUBMITTED captures enum value "SUBMITTED"
	PPMShipmentStatusSUBMITTED PPMShipmentStatus = "SUBMITTED"

	// PPMShipmentStatusWAITINGONCUSTOMER captures enum value "WAITING_ON_CUSTOMER"
	PPMShipmentStatusWAITINGONCUSTOMER PPMShipmentStatus = "WAITING_ON_CUSTOMER"

	// PPMShipmentStatusNEEDSADVANCEAPPROVAL captures enum value "NEEDS_ADVANCE_APPROVAL"
	PPMShipmentStatusNEEDSADVANCEAPPROVAL PPMShipmentStatus = "NEEDS_ADVANCE_APPROVAL"

	// PPMShipmentStatusNEEDSPAYMENTAPPROVAL captures enum value "NEEDS_PAYMENT_APPROVAL"
	PPMShipmentStatusNEEDSPAYMENTAPPROVAL PPMShipmentStatus = "NEEDS_PAYMENT_APPROVAL"

	// PPMShipmentStatusPAYMENTAPPROVED captures enum value "PAYMENT_APPROVED"
	PPMShipmentStatusPAYMENTAPPROVED PPMShipmentStatus = "PAYMENT_APPROVED"
)

// for schema
var pPMShipmentStatusEnum []interface{}

func init() {
	var res []PPMShipmentStatus
	if err := json.Unmarshal([]byte(`["DRAFT","SUBMITTED","WAITING_ON_CUSTOMER","NEEDS_ADVANCE_APPROVAL","NEEDS_PAYMENT_APPROVAL","PAYMENT_APPROVED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pPMShipmentStatusEnum = append(pPMShipmentStatusEnum, v)
	}
}

func (m PPMShipmentStatus) validatePPMShipmentStatusEnum(path, location string, value PPMShipmentStatus) error {
	if err := validate.EnumCase(path, location, value, pPMShipmentStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this p p m shipment status
func (m PPMShipmentStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePPMShipmentStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this p p m shipment status based on the context it is used
func (m PPMShipmentStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := validate.ReadOnly(ctx, "", "body", PPMShipmentStatus(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
