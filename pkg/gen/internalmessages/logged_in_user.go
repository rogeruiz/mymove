// Code generated by go-swagger; DO NOT EDIT.

package internalmessages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LoggedInUser logged in user
//
// swagger:model LoggedInUser
type LoggedInUser struct {

	// email
	// Read Only: true
	// Pattern: ^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$
	Email string `json:"email,omitempty"`

	// first name
	// Read Only: true
	FirstName string `json:"first_name,omitempty"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// office user
	OfficeUser *OfficeUser `json:"office_user,omitempty"`

	// roles
	Roles []*Role `json:"roles"`

	// service member
	ServiceMember *ServiceMemberPayload `json:"service_member,omitempty"`
}

// Validate validates this logged in user
func (m *LoggedInUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOfficeUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceMember(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoggedInUser) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.Pattern("email", "body", string(m.Email), `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`); err != nil {
		return err
	}

	return nil
}

func (m *LoggedInUser) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LoggedInUser) validateOfficeUser(formats strfmt.Registry) error {

	if swag.IsZero(m.OfficeUser) { // not required
		return nil
	}

	if m.OfficeUser != nil {
		if err := m.OfficeUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("office_user")
			}
			return err
		}
	}

	return nil
}

func (m *LoggedInUser) validateRoles(formats strfmt.Registry) error {

	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	for i := 0; i < len(m.Roles); i++ {
		if swag.IsZero(m.Roles[i]) { // not required
			continue
		}

		if m.Roles[i] != nil {
			if err := m.Roles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LoggedInUser) validateServiceMember(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceMember) { // not required
		return nil
	}

	if m.ServiceMember != nil {
		if err := m.ServiceMember.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_member")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoggedInUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoggedInUser) UnmarshalBinary(b []byte) error {
	var res LoggedInUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}