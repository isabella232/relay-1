// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WorkflowRunOperation workflow run operation
//
// swagger:model WorkflowRunOperation
type WorkflowRunOperation struct {

	// Cancel a workflow run
	Cancel bool `json:"cancel,omitempty"`
}

// Validate validates this workflow run operation
func (m *WorkflowRunOperation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowRunOperation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowRunOperation) UnmarshalBinary(b []byte) error {
	var res WorkflowRunOperation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
