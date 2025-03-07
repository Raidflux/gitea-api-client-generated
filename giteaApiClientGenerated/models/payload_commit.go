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

// PayloadCommit PayloadCommit represents a commit
//
// swagger:model PayloadCommit
type PayloadCommit struct {

	// added
	Added []string `json:"added"`

	// sha1 hash of the commit
	ID string `json:"id,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// modified
	Modified []string `json:"modified"`

	// removed
	Removed []string `json:"removed"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// URL
	URL string `json:"url,omitempty"`

	// author
	Author *PayloadUser `json:"author,omitempty"`

	// committer
	Committer *PayloadUser `json:"committer,omitempty"`

	// verification
	Verification *PayloadCommitVerification `json:"verification,omitempty"`
}

// Validate validates this payload commit
func (m *PayloadCommit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommitter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PayloadCommit) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PayloadCommit) validateAuthor(formats strfmt.Registry) error {
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

func (m *PayloadCommit) validateCommitter(formats strfmt.Registry) error {
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

func (m *PayloadCommit) validateVerification(formats strfmt.Registry) error {
	if swag.IsZero(m.Verification) { // not required
		return nil
	}

	if m.Verification != nil {
		if err := m.Verification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("verification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("verification")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this payload commit based on the context it is used
func (m *PayloadCommit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCommitter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVerification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PayloadCommit) contextValidateAuthor(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PayloadCommit) contextValidateCommitter(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PayloadCommit) contextValidateVerification(ctx context.Context, formats strfmt.Registry) error {

	if m.Verification != nil {

		if swag.IsZero(m.Verification) { // not required
			return nil
		}

		if err := m.Verification.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("verification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("verification")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PayloadCommit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PayloadCommit) UnmarshalBinary(b []byte) error {
	var res PayloadCommit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
