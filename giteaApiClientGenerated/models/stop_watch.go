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

// StopWatch StopWatch represent a running stopwatch
//
// swagger:model StopWatch
type StopWatch struct {

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// duration
	Duration string `json:"duration,omitempty"`

	// issue index
	IssueIndex int64 `json:"issue_index,omitempty"`

	// issue title
	IssueTitle string `json:"issue_title,omitempty"`

	// repo name
	RepoName string `json:"repo_name,omitempty"`

	// repo owner name
	RepoOwnerName string `json:"repo_owner_name,omitempty"`

	// seconds
	Seconds int64 `json:"seconds,omitempty"`
}

// Validate validates this stop watch
func (m *StopWatch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StopWatch) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this stop watch based on context it is used
func (m *StopWatch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StopWatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StopWatch) UnmarshalBinary(b []byte) error {
	var res StopWatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
