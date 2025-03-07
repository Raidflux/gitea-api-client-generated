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

// TrackedTime TrackedTime worked time for an issue / pr
//
// swagger:model TrackedTime
type TrackedTime struct {

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// deprecated (only for backwards compatibility)
	IssueID int64 `json:"issue_id,omitempty"`

	// Time in seconds
	Time int64 `json:"time,omitempty"`

	// deprecated (only for backwards compatibility)
	UserID int64 `json:"user_id,omitempty"`

	// user name
	UserName string `json:"user_name,omitempty"`

	// issue
	Issue *Issue `json:"issue,omitempty"`
}

// Validate validates this tracked time
func (m *TrackedTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrackedTime) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TrackedTime) validateIssue(formats strfmt.Registry) error {
	if swag.IsZero(m.Issue) { // not required
		return nil
	}

	if m.Issue != nil {
		if err := m.Issue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("issue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("issue")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tracked time based on the context it is used
func (m *TrackedTime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIssue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrackedTime) contextValidateIssue(ctx context.Context, formats strfmt.Registry) error {

	if m.Issue != nil {

		if swag.IsZero(m.Issue) { // not required
			return nil
		}

		if err := m.Issue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("issue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("issue")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrackedTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrackedTime) UnmarshalBinary(b []byte) error {
	var res TrackedTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
