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

// PullReviewComment PullReviewComment represents a comment on a pull request review
//
// swagger:model PullReviewComment
type PullReviewComment struct {

	// body
	Body string `json:"body,omitempty"`

	// commit ID
	CommitID string `json:"commit_id,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created_at,omitempty"`

	// diff hunk
	DiffHunk string `json:"diff_hunk,omitempty"`

	// HTML pull URL
	HTMLPullURL string `json:"pull_request_url,omitempty"`

	// HTML URL
	HTMLURL string `json:"html_url,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// line num
	LineNum uint64 `json:"position,omitempty"`

	// old line num
	OldLineNum uint64 `json:"original_position,omitempty"`

	// orig commit ID
	OrigCommitID string `json:"original_commit_id,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// review ID
	ReviewID int64 `json:"pull_request_review_id,omitempty"`

	// updated
	// Format: date-time
	Updated strfmt.DateTime `json:"updated_at,omitempty"`

	// resolver
	Resolver *User `json:"resolver,omitempty"`

	// user
	User *User `json:"user,omitempty"`
}

// Validate validates this pull review comment
func (m *PullReviewComment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolver(formats); err != nil {
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

func (m *PullReviewComment) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PullReviewComment) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PullReviewComment) validateResolver(formats strfmt.Registry) error {
	if swag.IsZero(m.Resolver) { // not required
		return nil
	}

	if m.Resolver != nil {
		if err := m.Resolver.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resolver")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resolver")
			}
			return err
		}
	}

	return nil
}

func (m *PullReviewComment) validateUser(formats strfmt.Registry) error {
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

// ContextValidate validate this pull review comment based on the context it is used
func (m *PullReviewComment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResolver(ctx, formats); err != nil {
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

func (m *PullReviewComment) contextValidateResolver(ctx context.Context, formats strfmt.Registry) error {

	if m.Resolver != nil {

		if swag.IsZero(m.Resolver) { // not required
			return nil
		}

		if err := m.Resolver.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resolver")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resolver")
			}
			return err
		}
	}

	return nil
}

func (m *PullReviewComment) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PullReviewComment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PullReviewComment) UnmarshalBinary(b []byte) error {
	var res PullReviewComment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
