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

// SitExtension sit extension
//
// swagger:model SitExtension
type SitExtension struct {

	// approved days
	ApprovedDays int64 `json:"approvedDays,omitempty"`

	// contractor remarks
	ContractorRemarks string `json:"contractorRemarks,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// decision date
	// Format: date-time
	DecisionDate strfmt.DateTime `json:"decisionDate,omitempty"`

	// e tag
	ETag string `json:"eTag,omitempty"`

	// id
	// Example: 1f2270c7-7166-40ae-981e-b200ebdf3054
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// mto shipment ID
	// Example: 1f2270c7-7166-40ae-981e-b200ebdf3054
	// Format: uuid
	MtoShipmentID strfmt.UUID `json:"mtoShipmentID,omitempty"`

	// office remarks
	OfficeRemarks string `json:"officeRemarks,omitempty"`

	// request reason
	// Enum: [SERIOUS_ILLNESS_MEMBER SERIOUS_ILLNESS_DEPENDENT IMPENDING_ASSIGNEMENT DIRECTED_TEMPORARY_DUTY NONAVAILABILITY_OF_CIVILIAN_HOUSING AWAITING_COMPLETION_OF_RESIDENCE OTHER]
	RequestReason interface{} `json:"requestReason,omitempty"`

	// requested days
	RequestedDays int64 `json:"requestedDays,omitempty"`

	// status
	// Enum: [PENDING APPROVED DENIED]
	Status interface{} `json:"status,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this sit extension
func (m *SitExtension) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDecisionDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtoShipmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SitExtension) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SitExtension) validateDecisionDate(formats strfmt.Registry) error {
	if swag.IsZero(m.DecisionDate) { // not required
		return nil
	}

	if err := validate.FormatOf("decisionDate", "body", "date-time", m.DecisionDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SitExtension) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SitExtension) validateMtoShipmentID(formats strfmt.Registry) error {
	if swag.IsZero(m.MtoShipmentID) { // not required
		return nil
	}

	if err := validate.FormatOf("mtoShipmentID", "body", "uuid", m.MtoShipmentID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SitExtension) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this sit extension based on context it is used
func (m *SitExtension) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SitExtension) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SitExtension) UnmarshalBinary(b []byte) error {
	var res SitExtension
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}