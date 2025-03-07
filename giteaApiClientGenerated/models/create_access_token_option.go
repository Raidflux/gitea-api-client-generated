// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateAccessTokenOption CreateAccessTokenOption options when create access token
//
// swagger:model CreateAccessTokenOption
type CreateAccessTokenOption struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// scopes
	Scopes []string `json:"scopes"`
}

// Validate validates this create access token option
func (m *CreateAccessTokenOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAccessTokenOption) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create access token option based on context it is used
func (m *CreateAccessTokenOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateAccessTokenOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAccessTokenOption) UnmarshalBinary(b []byte) error {
	var res CreateAccessTokenOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
