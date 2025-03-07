// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GeneralAttachmentSettings GeneralAttachmentSettings contains global Attachment settings exposed by API
//
// swagger:model GeneralAttachmentSettings
type GeneralAttachmentSettings struct {

	// allowed types
	AllowedTypes string `json:"allowed_types,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// max files
	MaxFiles int64 `json:"max_files,omitempty"`

	// max size
	MaxSize int64 `json:"max_size,omitempty"`
}

// Validate validates this general attachment settings
func (m *GeneralAttachmentSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this general attachment settings based on context it is used
func (m *GeneralAttachmentSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GeneralAttachmentSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GeneralAttachmentSettings) UnmarshalBinary(b []byte) error {
	var res GeneralAttachmentSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
