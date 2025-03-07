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

// Comment Comment represents a comment on a commit or issue
//
// swagger:model Comment
type Comment struct {

	// attachments
	Attachments []*Attachment `json:"assets"`

	// body
	Body string `json:"body,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created_at,omitempty"`

	// HTML URL
	HTMLURL string `json:"html_url,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// issue URL
	IssueURL string `json:"issue_url,omitempty"`

	// original author
	OriginalAuthor string `json:"original_author,omitempty"`

	// original author ID
	OriginalAuthorID int64 `json:"original_author_id,omitempty"`

	// p r URL
	PRURL string `json:"pull_request_url,omitempty"`

	// updated
	// Format: date-time
	Updated strfmt.DateTime `json:"updated_at,omitempty"`

	// user
	User *User `json:"user,omitempty"`
}

// Validate validates this comment
func (m *Comment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Comment) validateAttachments(formats strfmt.Registry) error {
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

func (m *Comment) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Comment) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Comment) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this comment based on the context it is used
func (m *Comment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttachments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Comment) contextValidateAttachments(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Comment) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {

		if swag.IsZero(m.User) { // not required
			return nil
		}

		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Comment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Comment) UnmarshalBinary(b []byte) error {
	var res Comment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
