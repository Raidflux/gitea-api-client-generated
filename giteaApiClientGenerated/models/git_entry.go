// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GitEntry GitEntry represents a git tree
//
// swagger:model GitEntry
type GitEntry struct {

	// mode
	Mode string `json:"mode,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// s h a
	SHA string `json:"sha,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// URL
	URL string `json:"url,omitempty"`
}

// Validate validates this git entry
func (m *GitEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this git entry based on context it is used
func (m *GitEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitEntry) UnmarshalBinary(b []byte) error {
	var res GitEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
