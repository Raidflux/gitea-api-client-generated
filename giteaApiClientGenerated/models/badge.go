// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Badge Badge represents a user badge
//
// swagger:model Badge
type Badge struct {

	// description
	Description string `json:"description,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// image URL
	ImageURL string `json:"image_url,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`
}

// Validate validates this badge
func (m *Badge) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this badge based on context it is used
func (m *Badge) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Badge) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Badge) UnmarshalBinary(b []byte) error {
	var res Badge
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
