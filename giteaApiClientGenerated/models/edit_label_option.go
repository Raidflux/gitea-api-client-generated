// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EditLabelOption EditLabelOption options for editing a label
//
// swagger:model EditLabelOption
type EditLabelOption struct {

	// color
	// Example: #00aabb
	Color string `json:"color,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// exclusive
	// Example: false
	Exclusive bool `json:"exclusive,omitempty"`

	// is archived
	// Example: false
	IsArchived bool `json:"is_archived,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this edit label option
func (m *EditLabelOption) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this edit label option based on context it is used
func (m *EditLabelOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EditLabelOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EditLabelOption) UnmarshalBinary(b []byte) error {
	var res EditLabelOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
