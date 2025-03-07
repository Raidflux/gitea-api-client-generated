// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MergePullRequestForm MergePullRequestForm form for merging Pull Request
//
// swagger:model MergePullRequestForm
type MergePullRequestForm struct {

	// delete branch after merge
	DeleteBranchAfterMerge bool `json:"delete_branch_after_merge,omitempty"`

	// do
	// Required: true
	// Enum: ["merge","rebase","rebase-merge","squash","fast-forward-only","manually-merged"]
	Do *string `json:"Do"`

	// force merge
	ForceMerge bool `json:"force_merge,omitempty"`

	// head commit ID
	HeadCommitID string `json:"head_commit_id,omitempty"`

	// merge commit ID
	MergeCommitID string `json:"MergeCommitID,omitempty"`

	// merge message field
	MergeMessageField string `json:"MergeMessageField,omitempty"`

	// merge title field
	MergeTitleField string `json:"MergeTitleField,omitempty"`

	// merge when checks succeed
	MergeWhenChecksSucceed bool `json:"merge_when_checks_succeed,omitempty"`
}

// Validate validates this merge pull request form
func (m *MergePullRequestForm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var mergePullRequestFormTypeDoPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["merge","rebase","rebase-merge","squash","fast-forward-only","manually-merged"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mergePullRequestFormTypeDoPropEnum = append(mergePullRequestFormTypeDoPropEnum, v)
	}
}

const (

	// MergePullRequestFormDoMerge captures enum value "merge"
	MergePullRequestFormDoMerge string = "merge"

	// MergePullRequestFormDoRebase captures enum value "rebase"
	MergePullRequestFormDoRebase string = "rebase"

	// MergePullRequestFormDoRebaseDashMerge captures enum value "rebase-merge"
	MergePullRequestFormDoRebaseDashMerge string = "rebase-merge"

	// MergePullRequestFormDoSquash captures enum value "squash"
	MergePullRequestFormDoSquash string = "squash"

	// MergePullRequestFormDoFastDashForwardDashOnly captures enum value "fast-forward-only"
	MergePullRequestFormDoFastDashForwardDashOnly string = "fast-forward-only"

	// MergePullRequestFormDoManuallyDashMerged captures enum value "manually-merged"
	MergePullRequestFormDoManuallyDashMerged string = "manually-merged"
)

// prop value enum
func (m *MergePullRequestForm) validateDoEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mergePullRequestFormTypeDoPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MergePullRequestForm) validateDo(formats strfmt.Registry) error {

	if err := validate.Required("Do", "body", m.Do); err != nil {
		return err
	}

	// value enum
	if err := m.validateDoEnum("Do", "body", *m.Do); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this merge pull request form based on context it is used
func (m *MergePullRequestForm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MergePullRequestForm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MergePullRequestForm) UnmarshalBinary(b []byte) error {
	var res MergePullRequestForm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
