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

// SITStatus s i t status
//
// swagger:model SITStatus
type SITStatus struct {

	// days in s i t
	// Minimum: 0
	DaysInSIT *int64 `json:"daysInSIT,omitempty"`

	// location
	// Enum: [ORIGIN DESTINATION]
	Location interface{} `json:"location,omitempty"`

	// past s i t service items
	PastSITServiceItems MTOServiceItems `json:"pastSITServiceItems,omitempty"`

	// sit departure date
	// Format: date-time
	SitDepartureDate *strfmt.DateTime `json:"sitDepartureDate,omitempty"`

	// sit entry date
	// Format: date-time
	SitEntryDate strfmt.DateTime `json:"sitEntryDate,omitempty"`

	// total days remaining
	// Minimum: 0
	TotalDaysRemaining *int64 `json:"totalDaysRemaining,omitempty"`

	// total s i t days used
	// Minimum: 0
	TotalSITDaysUsed *int64 `json:"totalSITDaysUsed,omitempty"`
}

// Validate validates this s i t status
func (m *SITStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDaysInSIT(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePastSITServiceItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSitDepartureDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSitEntryDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalDaysRemaining(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalSITDaysUsed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SITStatus) validateDaysInSIT(formats strfmt.Registry) error {
	if swag.IsZero(m.DaysInSIT) { // not required
		return nil
	}

	if err := validate.MinimumInt("daysInSIT", "body", *m.DaysInSIT, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *SITStatus) validatePastSITServiceItems(formats strfmt.Registry) error {
	if swag.IsZero(m.PastSITServiceItems) { // not required
		return nil
	}

	if err := m.PastSITServiceItems.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("pastSITServiceItems")
		}
		return err
	}

	return nil
}

func (m *SITStatus) validateSitDepartureDate(formats strfmt.Registry) error {
	if swag.IsZero(m.SitDepartureDate) { // not required
		return nil
	}

	if err := validate.FormatOf("sitDepartureDate", "body", "date-time", m.SitDepartureDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SITStatus) validateSitEntryDate(formats strfmt.Registry) error {
	if swag.IsZero(m.SitEntryDate) { // not required
		return nil
	}

	if err := validate.FormatOf("sitEntryDate", "body", "date-time", m.SitEntryDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SITStatus) validateTotalDaysRemaining(formats strfmt.Registry) error {
	if swag.IsZero(m.TotalDaysRemaining) { // not required
		return nil
	}

	if err := validate.MinimumInt("totalDaysRemaining", "body", *m.TotalDaysRemaining, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *SITStatus) validateTotalSITDaysUsed(formats strfmt.Registry) error {
	if swag.IsZero(m.TotalSITDaysUsed) { // not required
		return nil
	}

	if err := validate.MinimumInt("totalSITDaysUsed", "body", *m.TotalSITDaysUsed, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this s i t status based on the context it is used
func (m *SITStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePastSITServiceItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SITStatus) contextValidatePastSITServiceItems(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PastSITServiceItems.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("pastSITServiceItems")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SITStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SITStatus) UnmarshalBinary(b []byte) error {
	var res SITStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}