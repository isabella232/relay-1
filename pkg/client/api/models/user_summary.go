// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserSummary user summary
//
// swagger:model UserSummary
type UserSummary struct {
	UserIdentifier

	// User email
	Email string `json:"email,omitempty"`

	// User name
	// Required: true
	Name *string `json:"name"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *UserSummary) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 UserIdentifier
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.UserIdentifier = aO0

	// AO1
	var dataAO1 struct {
		Email string `json:"email,omitempty"`

		Name *string `json:"name"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Email = dataAO1.Email

	m.Name = dataAO1.Name

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m UserSummary) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.UserIdentifier)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		Email string `json:"email,omitempty"`

		Name *string `json:"name"`
	}

	dataAO1.Email = m.Email

	dataAO1.Name = m.Name

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this user summary
func (m *UserSummary) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with UserIdentifier
	if err := m.UserIdentifier.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserSummary) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserSummary) UnmarshalBinary(b []byte) error {
	var res UserSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
