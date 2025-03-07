// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IssueConfigValidation issue config validation
//
// swagger:model IssueConfigValidation
type IssueConfigValidation struct {

	// message
	Message string `json:"message,omitempty"`

	// valid
	Valid bool `json:"valid,omitempty"`
}

// Validate validates this issue config validation
func (m *IssueConfigValidation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this issue config validation based on context it is used
func (m *IssueConfigValidation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IssueConfigValidation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssueConfigValidation) UnmarshalBinary(b []byte) error {
	var res IssueConfigValidation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
