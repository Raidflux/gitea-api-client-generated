// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IssueConfigContactLink issue config contact link
//
// swagger:model IssueConfigContactLink
type IssueConfigContactLink struct {

	// about
	About string `json:"about,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// URL
	URL string `json:"url,omitempty"`
}

// Validate validates this issue config contact link
func (m *IssueConfigContactLink) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this issue config contact link based on context it is used
func (m *IssueConfigContactLink) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IssueConfigContactLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssueConfigContactLink) UnmarshalBinary(b []byte) error {
	var res IssueConfigContactLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
