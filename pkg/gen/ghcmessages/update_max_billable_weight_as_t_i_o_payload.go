// Code generated by go-swagger; DO NOT EDIT.

package ghcmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateMaxBillableWeightAsTIOPayload update max billable weight as t i o payload
//
// swagger:model UpdateMaxBillableWeightAsTIOPayload
type UpdateMaxBillableWeightAsTIOPayload struct {

	// unit is in lbs
	// Example: 2000
	// Required: true
	// Minimum: 1
	AuthorizedWeight *int64 `json:"authorizedWeight"`

	// TIO remarks for updating the max billable weight
	// Example: Increasing max billable weight
	// Required: true
	// Min Length: 1
	TioRemarks *string `json:"tioRemarks"`
}

// Validate validates this update max billable weight as t i o payload
func (m *UpdateMaxBillableWeightAsTIOPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorizedWeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTioRemarks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateMaxBillableWeightAsTIOPayload) validateAuthorizedWeight(formats strfmt.Registry) error {

	if err := validate.Required("authorizedWeight", "body", m.AuthorizedWeight); err != nil {
		return err
	}

	if err := validate.MinimumInt("authorizedWeight", "body", *m.AuthorizedWeight, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *UpdateMaxBillableWeightAsTIOPayload) validateTioRemarks(formats strfmt.Registry) error {

	if err := validate.Required("tioRemarks", "body", m.TioRemarks); err != nil {
		return err
	}

	if err := validate.MinLength("tioRemarks", "body", *m.TioRemarks, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update max billable weight as t i o payload based on context it is used
func (m *UpdateMaxBillableWeightAsTIOPayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateMaxBillableWeightAsTIOPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateMaxBillableWeightAsTIOPayload) UnmarshalBinary(b []byte) error {
	var res UpdateMaxBillableWeightAsTIOPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}