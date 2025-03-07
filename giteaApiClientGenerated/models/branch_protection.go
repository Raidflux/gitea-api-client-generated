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

// BranchProtection BranchProtection represents a branch protection for a repository
//
// swagger:model BranchProtection
type BranchProtection struct {

	// approvals whitelist teams
	ApprovalsWhitelistTeams []string `json:"approvals_whitelist_teams"`

	// approvals whitelist usernames
	ApprovalsWhitelistUsernames []string `json:"approvals_whitelist_username"`

	// block on official review requests
	BlockOnOfficialReviewRequests bool `json:"block_on_official_review_requests,omitempty"`

	// block on outdated branch
	BlockOnOutdatedBranch bool `json:"block_on_outdated_branch,omitempty"`

	// block on rejected reviews
	BlockOnRejectedReviews bool `json:"block_on_rejected_reviews,omitempty"`

	// Deprecated: true
	BranchName string `json:"branch_name,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created_at,omitempty"`

	// dismiss stale approvals
	DismissStaleApprovals bool `json:"dismiss_stale_approvals,omitempty"`

	// enable approvals whitelist
	EnableApprovalsWhitelist bool `json:"enable_approvals_whitelist,omitempty"`

	// enable force push
	EnableForcePush bool `json:"enable_force_push,omitempty"`

	// enable force push allowlist
	EnableForcePushAllowlist bool `json:"enable_force_push_allowlist,omitempty"`

	// enable merge whitelist
	EnableMergeWhitelist bool `json:"enable_merge_whitelist,omitempty"`

	// enable push
	EnablePush bool `json:"enable_push,omitempty"`

	// enable push whitelist
	EnablePushWhitelist bool `json:"enable_push_whitelist,omitempty"`

	// enable status check
	EnableStatusCheck bool `json:"enable_status_check,omitempty"`

	// force push allowlist deploy keys
	ForcePushAllowlistDeployKeys bool `json:"force_push_allowlist_deploy_keys,omitempty"`

	// force push allowlist teams
	ForcePushAllowlistTeams []string `json:"force_push_allowlist_teams"`

	// force push allowlist usernames
	ForcePushAllowlistUsernames []string `json:"force_push_allowlist_usernames"`

	// ignore stale approvals
	IgnoreStaleApprovals bool `json:"ignore_stale_approvals,omitempty"`

	// merge whitelist teams
	MergeWhitelistTeams []string `json:"merge_whitelist_teams"`

	// merge whitelist usernames
	MergeWhitelistUsernames []string `json:"merge_whitelist_usernames"`

	// protected file patterns
	ProtectedFilePatterns string `json:"protected_file_patterns,omitempty"`

	// push whitelist deploy keys
	PushWhitelistDeployKeys bool `json:"push_whitelist_deploy_keys,omitempty"`

	// push whitelist teams
	PushWhitelistTeams []string `json:"push_whitelist_teams"`

	// push whitelist usernames
	PushWhitelistUsernames []string `json:"push_whitelist_usernames"`

	// require signed commits
	RequireSignedCommits bool `json:"require_signed_commits,omitempty"`

	// required approvals
	RequiredApprovals int64 `json:"required_approvals,omitempty"`

	// rule name
	RuleName string `json:"rule_name,omitempty"`

	// status check contexts
	StatusCheckContexts []string `json:"status_check_contexts"`

	// unprotected file patterns
	UnprotectedFilePatterns string `json:"unprotected_file_patterns,omitempty"`

	// updated
	// Format: date-time
	Updated strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this branch protection
func (m *BranchProtection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BranchProtection) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BranchProtection) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this branch protection based on context it is used
func (m *BranchProtection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BranchProtection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BranchProtection) UnmarshalBinary(b []byte) error {
	var res BranchProtection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
