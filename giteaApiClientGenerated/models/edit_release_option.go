// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EditReleaseOption EditReleaseOption options when editing a release
//
// swagger:model EditReleaseOption
type EditReleaseOption struct {

	// is draft
	IsDraft bool `json:"draft,omitempty"`

	// is prerelease
	IsPrerelease bool `json:"prerelease,omitempty"`

	// note
	Note string `json:"body,omitempty"`

	// tag name
	TagName string `json:"tag_name,omitempty"`

	// target
	Target string `json:"target_commitish,omitempty"`

	// title
	Title string `json:"name,omitempty"`
}

// Validate validates this edit release option
func (m *EditReleaseOption) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this edit release option based on context it is used
func (m *EditReleaseOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EditReleaseOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EditReleaseOption) UnmarshalBinary(b []byte) error {
	var res EditReleaseOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
