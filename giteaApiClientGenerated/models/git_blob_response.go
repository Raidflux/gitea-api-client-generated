// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GitBlobResponse GitBlobResponse represents a git blob
//
// swagger:model GitBlobResponse
type GitBlobResponse struct {

	// content
	Content string `json:"content,omitempty"`

	// encoding
	Encoding string `json:"encoding,omitempty"`

	// s h a
	SHA string `json:"sha,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// URL
	URL string `json:"url,omitempty"`
}

// Validate validates this git blob response
func (m *GitBlobResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this git blob response based on context it is used
func (m *GitBlobResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitBlobResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitBlobResponse) UnmarshalBinary(b []byte) error {
	var res GitBlobResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
