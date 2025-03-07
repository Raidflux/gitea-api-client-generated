// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Release Release represents a repository release
//
// swagger:model Release
type Release struct {

	// attachments
	Attachments []*Attachment `json:"assets"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// HTML URL
	HTMLURL string `json:"html_url,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// is draft
	IsDraft bool `json:"draft,omitempty"`

	// is prerelease
	IsPrerelease bool `json:"prerelease,omitempty"`

	// note
	Note string `json:"body,omitempty"`

	// published at
	// Format: date-time
	PublishedAt strfmt.DateTime `json:"published_at,omitempty"`

	// tag name
	TagName string `json:"tag_name,omitempty"`

	// tar URL
	TarURL string `json:"tarball_url,omitempty"`

	// target
	Target string `json:"target_commitish,omitempty"`

	// title
	Title string `json:"name,omitempty"`

	// URL
	URL string `json:"url,omitempty"`

	// upload URL
	UploadURL string `json:"upload_url,omitempty"`

	// zip URL
	ZipURL string `json:"zipball_url,omitempty"`

	// author
	Author *User `json:"author,omitempty"`
}

// Validate validates this release
func (m *Release) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Release) validateAttachments(formats strfmt.Registry) error {
	if swag.IsZero(m.Attachments) { // not required
		return nil
	}

	for i := 0; i < len(m.Attachments); i++ {
		if swag.IsZero(m.Attachments[i]) { // not required
			continue
		}

		if m.Attachments[i] != nil {
			if err := m.Attachments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Release) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Release) validatePublishedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.PublishedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("published_at", "body", "date-time", m.PublishedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Release) validateAuthor(formats strfmt.Registry) error {
	if swag.IsZero(m.Author) { // not required
		return nil
	}

	if m.Author != nil {
		if err := m.Author.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("author")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("author")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this release based on the context it is used
func (m *Release) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttachments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuthor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Release) contextValidateAttachments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attachments); i++ {

		if m.Attachments[i] != nil {

			if swag.IsZero(m.Attachments[i]) { // not required
				return nil
			}

			if err := m.Attachments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Release) contextValidateAuthor(ctx context.Context, formats strfmt.Registry) error {

	if m.Author != nil {

		if swag.IsZero(m.Author) { // not required
			return nil
		}

		if err := m.Author.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("author")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("author")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Release) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Release) UnmarshalBinary(b []byte) error {
	var res Release
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
