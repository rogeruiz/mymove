// Code generated by go-swagger; DO NOT EDIT.

package primemessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PPMShipment A personally procured move is a type of shipment that a service members moves themselves.
//
// swagger:model PPMShipment
type PPMShipment struct {

	// actual move date
	// Format: date
	ActualMoveDate *strfmt.Date `json:"actualMoveDate"`

	// The amount request for an advance, or null if no advance is requested
	//
	Advance *int64 `json:"advance,omitempty"`

	// Indicates whether an advance has been requested for the PPM shipment.
	//
	AdvanceRequested *bool `json:"advanceRequested,omitempty"`

	// approved at
	// Format: date-time
	ApprovedAt *strfmt.DateTime `json:"approvedAt"`

	// created at
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// deleted at
	// Format: date-time
	DeletedAt *strfmt.DateTime `json:"deletedAt"`

	// ZIP
	// Example: 90210
	// Pattern: ^(\d{5})$
	DestinationPostalCode string `json:"destinationPostalCode,omitempty"`

	// A hash unique to this shipment that should be used as the "If-Match" header for any updates.
	// Read Only: true
	ETag string `json:"eTag,omitempty"`

	// estimated incentive
	EstimatedIncentive *int64 `json:"estimatedIncentive"`

	// estimated weight
	// Example: 4200
	EstimatedWeight *int64 `json:"estimatedWeight"`

	// Date the customer expects to move.
	//
	// Format: date
	ExpectedDepartureDate strfmt.Date `json:"expectedDepartureDate,omitempty"`

	// Indicates whether PPM shipment has pro gear.
	//
	HasProGear *bool `json:"hasProGear"`

	// id
	// Example: 1f2270c7-7166-40ae-981e-b200ebdf3054
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The net weight of the shipment once it has been weight
	//
	// Example: 4300
	NetWeight *int64 `json:"netWeight"`

	// ZIP
	//
	// zip code
	// Example: 90210
	// Pattern: ^(\d{5})$
	PickupPostalCode string `json:"pickupPostalCode,omitempty"`

	// pro gear weight
	ProGearWeight *int64 `json:"proGearWeight"`

	// reviewed at
	// Format: date-time
	ReviewedAt *strfmt.DateTime `json:"reviewedAt"`

	// ZIP
	// Example: 90210
	// Pattern: ^(\d{5})$
	SecondaryDestinationPostalCode *string `json:"secondaryDestinationPostalCode"`

	// ZIP
	// Example: 90210
	// Pattern: ^(\d{5})$
	SecondaryPickupPostalCode *string `json:"secondaryPickupPostalCode"`

	// shipment Id
	// Example: 1f2270c7-7166-40ae-981e-b200ebdf3054
	// Read Only: true
	// Format: uuid
	ShipmentID strfmt.UUID `json:"shipmentId,omitempty"`

	// sit expected
	SitExpected bool `json:"sitExpected,omitempty"`

	// spouse pro gear weight
	SpouseProGearWeight *int64 `json:"spouseProGearWeight"`

	// status
	Status PPMShipmentStatus `json:"status,omitempty"`

	// submitted at
	// Format: date-time
	SubmittedAt *strfmt.DateTime `json:"submittedAt"`

	// updated at
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this p p m shipment
func (m *PPMShipment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActualMoveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApprovedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationPostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpectedDepartureDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePickupPostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReviewedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryDestinationPostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryPickupPostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmittedAt(formats); err != nil {
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

func (m *PPMShipment) validateActualMoveDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ActualMoveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("actualMoveDate", "body", "date", m.ActualMoveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PPMShipment) validateApprovedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ApprovedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("approvedAt", "body", "date-time", m.ApprovedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PPMShipment) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PPMShipment) validateDeletedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.DeletedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("deletedAt", "body", "date-time", m.DeletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PPMShipment) validateDestinationPostalCode(formats strfmt.Registry) error {
	if swag.IsZero(m.DestinationPostalCode) { // not required
		return nil
	}

	if err := validate.Pattern("destinationPostalCode", "body", m.DestinationPostalCode, `^(\d{5})$`); err != nil {
		return err
	}

	return nil
}

func (m *PPMShipment) validateExpectedDepartureDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpectedDepartureDate) { // not required
		return nil
	}

	if err := validate.FormatOf("expectedDepartureDate", "body", "date", m.ExpectedDepartureDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PPMShipment) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PPMShipment) validatePickupPostalCode(formats strfmt.Registry) error {
	if swag.IsZero(m.PickupPostalCode) { // not required
		return nil
	}

	if err := validate.Pattern("pickupPostalCode", "body", m.PickupPostalCode, `^(\d{5})$`); err != nil {
		return err
	}

	return nil
}

func (m *PPMShipment) validateReviewedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ReviewedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("reviewedAt", "body", "date-time", m.ReviewedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PPMShipment) validateSecondaryDestinationPostalCode(formats strfmt.Registry) error {
	if swag.IsZero(m.SecondaryDestinationPostalCode) { // not required
		return nil
	}

	if err := validate.Pattern("secondaryDestinationPostalCode", "body", *m.SecondaryDestinationPostalCode, `^(\d{5})$`); err != nil {
		return err
	}

	return nil
}

func (m *PPMShipment) validateSecondaryPickupPostalCode(formats strfmt.Registry) error {
	if swag.IsZero(m.SecondaryPickupPostalCode) { // not required
		return nil
	}

	if err := validate.Pattern("secondaryPickupPostalCode", "body", *m.SecondaryPickupPostalCode, `^(\d{5})$`); err != nil {
		return err
	}

	return nil
}

func (m *PPMShipment) validateShipmentID(formats strfmt.Registry) error {
	if swag.IsZero(m.ShipmentID) { // not required
		return nil
	}

	if err := validate.FormatOf("shipmentId", "body", "uuid", m.ShipmentID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PPMShipment) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *PPMShipment) validateSubmittedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.SubmittedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("submittedAt", "body", "date-time", m.SubmittedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PPMShipment) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p p m shipment based on the context it is used
func (m *PPMShipment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateETag(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShipmentID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PPMShipment) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdAt", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *PPMShipment) contextValidateETag(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "eTag", "body", string(m.ETag)); err != nil {
		return err
	}

	return nil
}

func (m *PPMShipment) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *PPMShipment) contextValidateShipmentID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "shipmentId", "body", strfmt.UUID(m.ShipmentID)); err != nil {
		return err
	}

	return nil
}

func (m *PPMShipment) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *PPMShipment) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updatedAt", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PPMShipment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PPMShipment) UnmarshalBinary(b []byte) error {
	var res PPMShipment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}