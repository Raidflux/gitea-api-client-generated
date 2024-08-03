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
)

// FilesResponse FilesResponse contains information about multiple files from a repo
//
// swagger:model FilesResponse
type FilesResponse struct {

	// files
	Files []*ContentsResponse `json:"files"`

	// commit
	Commit *FileCommitResponse `json:"commit,omitempty"`

	// verification
	Verification *PayloadCommitVerification `json:"verification,omitempty"`
}

// Validate validates this files response
func (m *FilesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommit(formats); err != nil {
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

func (m *FilesResponse) validateFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.Files) { // not required
		return nil
	}

	for i := 0; i < len(m.Files); i++ {
		if swag.IsZero(m.Files[i]) { // not required
			continue
		}

		if m.Files[i] != nil {
			if err := m.Files[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FilesResponse) validateCommit(formats strfmt.Registry) error {
	if swag.IsZero(m.Commit) { // not required
		return nil
	}

	if m.Commit != nil {
		if err := m.Commit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commit")
			}
			return err
		}
	}

	return nil
}

func (m *FilesResponse) validateVerification(formats strfmt.Registry) error {
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

// ContextValidate validate this files response based on the context it is used
func (m *FilesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCommit(ctx, formats); err != nil {
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

func (m *FilesResponse) contextValidateFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Files); i++ {

		if m.Files[i] != nil {

			if swag.IsZero(m.Files[i]) { // not required
				return nil
			}

			if err := m.Files[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FilesResponse) contextValidateCommit(ctx context.Context, formats strfmt.Registry) error {

	if m.Commit != nil {

		if swag.IsZero(m.Commit) { // not required
			return nil
		}

		if err := m.Commit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commit")
			}
			return err
		}
	}

	return nil
}

func (m *FilesResponse) contextValidateVerification(ctx context.Context, formats strfmt.Registry) error {

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
func (m *FilesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilesResponse) UnmarshalBinary(b []byte) error {
	var res FilesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
