// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateFileOptions UpdateFileOptions options for updating files
// Note: `author` and `committer` are optional (if only one is given, it will be used for the other, otherwise the authenticated user will be used)
//
// swagger:model UpdateFileOptions
type UpdateFileOptions struct {

	// branch (optional) to base this file from. if not given, the default branch is used
	BranchName string `json:"branch,omitempty"`

	// content must be base64 encoded
	// Required: true
	ContentBase64 *string `json:"content"`

	// from_path (optional) is the path of the original file which will be moved/renamed to the path in the URL
	FromPath string `json:"from_path,omitempty"`

	// message (optional) for the commit of this file. if not supplied, a default message will be used
	Message string `json:"message,omitempty"`

	// new_branch (optional) will make a new branch from `branch` before creating the file
	NewBranchName string `json:"new_branch,omitempty"`

	// sha is the SHA for the file that already exists
	// Required: true
	SHA *string `json:"sha"`

	// Add a Signed-off-by trailer by the committer at the end of the commit log message.
	Signoff bool `json:"signoff,omitempty"`

	// author
	Author *Identity `json:"author,omitempty"`

	// committer
	Committer *Identity `json:"committer,omitempty"`

	// dates
	Dates *CommitDateOptions `json:"dates,omitempty"`
}

// Validate validates this update file options
func (m *UpdateFileOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentBase64(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSHA(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommitter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateFileOptions) validateContentBase64(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.ContentBase64); err != nil {
		return err
	}

	return nil
}

func (m *UpdateFileOptions) validateSHA(formats strfmt.Registry) error {

	if err := validate.Required("sha", "body", m.SHA); err != nil {
		return err
	}

	return nil
}

func (m *UpdateFileOptions) validateAuthor(formats strfmt.Registry) error {
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

func (m *UpdateFileOptions) validateCommitter(formats strfmt.Registry) error {
	if swag.IsZero(m.Committer) { // not required
		return nil
	}

	if m.Committer != nil {
		if err := m.Committer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("committer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("committer")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateFileOptions) validateDates(formats strfmt.Registry) error {
	if swag.IsZero(m.Dates) { // not required
		return nil
	}

	if m.Dates != nil {
		if err := m.Dates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dates")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update file options based on the context it is used
func (m *UpdateFileOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCommitter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateFileOptions) contextValidateAuthor(ctx context.Context, formats strfmt.Registry) error {

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

func (m *UpdateFileOptions) contextValidateCommitter(ctx context.Context, formats strfmt.Registry) error {

	if m.Committer != nil {

		if swag.IsZero(m.Committer) { // not required
			return nil
		}

		if err := m.Committer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("committer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("committer")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateFileOptions) contextValidateDates(ctx context.Context, formats strfmt.Registry) error {

	if m.Dates != nil {

		if swag.IsZero(m.Dates) { // not required
			return nil
		}

		if err := m.Dates.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dates")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateFileOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateFileOptions) UnmarshalBinary(b []byte) error {
	var res UpdateFileOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
