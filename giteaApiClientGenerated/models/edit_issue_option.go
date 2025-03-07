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

// EditIssueOption EditIssueOption options for editing an issue
//
// swagger:model EditIssueOption
type EditIssueOption struct {

	// deprecated
	Assignee string `json:"assignee,omitempty"`

	// assignees
	Assignees []string `json:"assignees"`

	// body
	Body string `json:"body,omitempty"`

	// deadline
	// Format: date-time
	Deadline strfmt.DateTime `json:"due_date,omitempty"`

	// milestone
	Milestone int64 `json:"milestone,omitempty"`

	// ref
	Ref string `json:"ref,omitempty"`

	// remove deadline
	RemoveDeadline bool `json:"unset_due_date,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this edit issue option
func (m *EditIssueOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeadline(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EditIssueOption) validateDeadline(formats strfmt.Registry) error {
	if swag.IsZero(m.Deadline) { // not required
		return nil
	}

	if err := validate.FormatOf("due_date", "body", "date-time", m.Deadline.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this edit issue option based on context it is used
func (m *EditIssueOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EditIssueOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EditIssueOption) UnmarshalBinary(b []byte) error {
	var res EditIssueOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
