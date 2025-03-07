// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FileLinksResponse FileLinksResponse contains the links for a repo's file
//
// swagger:model FileLinksResponse
type FileLinksResponse struct {

	// git URL
	GitURL string `json:"git,omitempty"`

	// HTML URL
	HTMLURL string `json:"html,omitempty"`

	// self
	Self string `json:"self,omitempty"`
}

// Validate validates this file links response
func (m *FileLinksResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this file links response based on context it is used
func (m *FileLinksResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FileLinksResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileLinksResponse) UnmarshalBinary(b []byte) error {
	var res FileLinksResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
