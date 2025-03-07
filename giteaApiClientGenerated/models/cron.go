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

// Cron Cron represents a Cron task
//
// swagger:model Cron
type Cron struct {

	// exec times
	ExecTimes int64 `json:"exec_times,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// next
	// Format: date-time
	Next strfmt.DateTime `json:"next,omitempty"`

	// prev
	// Format: date-time
	Prev strfmt.DateTime `json:"prev,omitempty"`

	// schedule
	Schedule string `json:"schedule,omitempty"`
}

// Validate validates this cron
func (m *Cron) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrev(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cron) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("next", "body", "date-time", m.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Cron) validatePrev(formats strfmt.Registry) error {
	if swag.IsZero(m.Prev) { // not required
		return nil
	}

	if err := validate.FormatOf("prev", "body", "date-time", m.Prev.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cron based on context it is used
func (m *Cron) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Cron) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cron) UnmarshalBinary(b []byte) error {
	var res Cron
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
