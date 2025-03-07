// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// CommitStatusState CommitStatusState holds the state of a CommitStatus
// It can be "pending", "success", "error" and "failure"
//
// swagger:model CommitStatusState
type CommitStatusState string

// Validate validates this commit status state
func (m CommitStatusState) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this commit status state based on context it is used
func (m CommitStatusState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
