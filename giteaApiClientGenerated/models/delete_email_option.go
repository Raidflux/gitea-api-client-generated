// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteEmailOption DeleteEmailOption options when deleting email addresses
//
// swagger:model DeleteEmailOption
type DeleteEmailOption struct {

	// email addresses to delete
	Emails []string `json:"emails"`
}

// Validate validates this delete email option
func (m *DeleteEmailOption) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete email option based on context it is used
func (m *DeleteEmailOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteEmailOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteEmailOption) UnmarshalBinary(b []byte) error {
	var res DeleteEmailOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
