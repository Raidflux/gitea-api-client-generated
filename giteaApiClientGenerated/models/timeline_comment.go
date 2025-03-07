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

// TimelineComment TimelineComment represents a timeline comment (comment of any type) on a commit or issue
//
// swagger:model TimelineComment
type TimelineComment struct {

	// body
	Body string `json:"body,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created_at,omitempty"`

	// HTML URL
	HTMLURL string `json:"html_url,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// issue URL
	IssueURL string `json:"issue_url,omitempty"`

	// new ref
	NewRef string `json:"new_ref,omitempty"`

	// new title
	NewTitle string `json:"new_title,omitempty"`

	// old project ID
	OldProjectID int64 `json:"old_project_id,omitempty"`

	// old ref
	OldRef string `json:"old_ref,omitempty"`

	// old title
	OldTitle string `json:"old_title,omitempty"`

	// p r URL
	PRURL string `json:"pull_request_url,omitempty"`

	// project ID
	ProjectID int64 `json:"project_id,omitempty"`

	// ref action
	RefAction string `json:"ref_action,omitempty"`

	// commit SHA where issue/PR was referenced
	RefCommitSHA string `json:"ref_commit_sha,omitempty"`

	// whether the assignees were removed or added
	RemovedAssignee bool `json:"removed_assignee,omitempty"`

	// review ID
	ReviewID int64 `json:"review_id,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated
	// Format: date-time
	Updated strfmt.DateTime `json:"updated_at,omitempty"`

	// assignee
	Assignee *User `json:"assignee,omitempty"`

	// assignee team
	AssigneeTeam *Team `json:"assignee_team,omitempty"`

	// dependent issue
	DependentIssue *Issue `json:"dependent_issue,omitempty"`

	// label
	Label *Label `json:"label,omitempty"`

	// milestone
	Milestone *Milestone `json:"milestone,omitempty"`

	// old milestone
	OldMilestone *Milestone `json:"old_milestone,omitempty"`

	// ref comment
	RefComment *Comment `json:"ref_comment,omitempty"`

	// ref issue
	RefIssue *Issue `json:"ref_issue,omitempty"`

	// resolve doer
	ResolveDoer *User `json:"resolve_doer,omitempty"`

	// tracked time
	TrackedTime *TrackedTime `json:"tracked_time,omitempty"`

	// user
	User *User `json:"user,omitempty"`
}

// Validate validates this timeline comment
func (m *TimelineComment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssigneeTeam(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDependentIssue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMilestone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOldMilestone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefIssue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolveDoer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackedTime(formats); err != nil {
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

func (m *TimelineComment) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TimelineComment) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TimelineComment) validateAssignee(formats strfmt.Registry) error {
	if swag.IsZero(m.Assignee) { // not required
		return nil
	}

	if m.Assignee != nil {
		if err := m.Assignee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assignee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assignee")
			}
			return err
		}
	}

	return nil
}

func (m *TimelineComment) validateAssigneeTeam(formats strfmt.Registry) error {
	if swag.IsZero(m.AssigneeTeam) { // not required
		return nil
	}

	if m.AssigneeTeam != nil {
		if err := m.AssigneeTeam.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assignee_team")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assignee_team")
			}
			return err
		}
	}

	return nil
}

func (m *TimelineComment) validateDependentIssue(formats strfmt.Registry) error {
	if swag.IsZero(m.DependentIssue) { // not required
		return nil
	}

	if m.DependentIssue != nil {
		if err := m.DependentIssue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dependent_issue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dependent_issue")
			}
			return err
		}
	}

	return nil
}

func (m *TimelineComment) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if m.Label != nil {
		if err := m.Label.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("label")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("label")
			}
			return err
		}
	}

	return nil
}

func (m *TimelineComment) validateMilestone(formats strfmt.Registry) error {
	if swag.IsZero(m.Milestone) { // not required
		return nil
	}

	if m.Milestone != nil {
		if err := m.Milestone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("milestone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("milestone")
			}
			return err
		}
	}

	return nil
}

func (m *TimelineComment) validateOldMilestone(formats strfmt.Registry) error {
	if swag.IsZero(m.OldMilestone) { // not required
		return nil
	}

	if m.OldMilestone != nil {
		if err := m.OldMilestone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("old_milestone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("old_milestone")
			}
			return err
		}
	}

	return nil
}

func (m *TimelineComment) validateRefComment(formats strfmt.Registry) error {
	if swag.IsZero(m.RefComment) { // not required
		return nil
	}

	if m.RefComment != nil {
		if err := m.RefComment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ref_comment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ref_comment")
			}
			return err
		}
	}

	return nil
}

func (m *TimelineComment) validateRefIssue(formats strfmt.Registry) error {
	if swag.IsZero(m.RefIssue) { // not required
		return nil
	}

	if m.RefIssue != nil {
		if err := m.RefIssue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ref_issue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ref_issue")
			}
			return err
		}
	}

	return nil
}

func (m *TimelineComment) validateResolveDoer(formats strfmt.Registry) error {
	if swag.IsZero(m.ResolveDoer) { // not required
		return nil
	}

	if m.ResolveDoer != nil {
		if err := m.ResolveDoer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resolve_doer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resolve_doer")
			}
			return err
		}
	}

	return nil
}

func (m *TimelineComment) validateTrackedTime(formats strfmt.Registry) error {
	if swag.IsZero(m.TrackedTime) { // not required
		return nil
	}

	if m.TrackedTime != nil {
		if err := m.TrackedTime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tracked_time")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tracked_time")
			}
			return err
		}
	}

	return nil
}

func (m *TimelineComment) validateUser(formats strfmt.Registry) error {
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

// ContextValidate validate this timeline comment based on the context it is used
func (m *TimelineComment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssignee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAssigneeTeam(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDependentIssue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMilestone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOldMilestone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRefComment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRefIssue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResolveDoer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrackedTime(ctx, formats); err != nil {
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

func (m *TimelineComment) contextValidateAssignee(ctx context.Context, formats strfmt.Registry) error {

	if m.Assignee != nil {

		if swag.IsZero(m.Assignee) { // not required
			return nil
		}

		if err := m.Assignee.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assignee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assignee")
			}
			return err
		}
	}

	return nil
}

func (m *TimelineComment) contextValidateAssigneeTeam(ctx context.Context, formats strfmt.Registry) error {

	if m.AssigneeTeam != nil {

		if swag.IsZero(m.AssigneeTeam) { // not required
			return nil
		}

		if err := m.AssigneeTeam.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assignee_team")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assignee_team")
			}
			return err
		}
	}

	return nil
}

func (m *TimelineComment) contextValidateDependentIssue(ctx context.Context, formats strfmt.Registry) error {

	if m.DependentIssue != nil {

		if swag.IsZero(m.DependentIssue) { // not required
			return nil
		}

		if err := m.DependentIssue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dependent_issue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dependent_issue")
			}
			return err
		}
	}

	return nil
}

func (m *TimelineComment) contextValidateLabel(ctx context.Context, formats strfmt.Registry) error {

	if m.Label != nil {

		if swag.IsZero(m.Label) { // not required
			return nil
		}

		if err := m.Label.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("label")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("label")
			}
			return err
		}
	}

	return nil
}

func (m *TimelineComment) contextValidateMilestone(ctx context.Context, formats strfmt.Registry) error {

	if m.Milestone != nil {

		if swag.IsZero(m.Milestone) { // not required
			return nil
		}

		if err := m.Milestone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("milestone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("milestone")
			}
			return err
		}
	}

	return nil
}

func (m *TimelineComment) contextValidateOldMilestone(ctx context.Context, formats strfmt.Registry) error {

	if m.OldMilestone != nil {

		if swag.IsZero(m.OldMilestone) { // not required
			return nil
		}

		if err := m.OldMilestone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("old_milestone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("old_milestone")
			}
			return err
		}
	}

	return nil
}

func (m *TimelineComment) contextValidateRefComment(ctx context.Context, formats strfmt.Registry) error {

	if m.RefComment != nil {

		if swag.IsZero(m.RefComment) { // not required
			return nil
		}

		if err := m.RefComment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ref_comment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ref_comment")
			}
			return err
		}
	}

	return nil
}

func (m *TimelineComment) contextValidateRefIssue(ctx context.Context, formats strfmt.Registry) error {

	if m.RefIssue != nil {

		if swag.IsZero(m.RefIssue) { // not required
			return nil
		}

		if err := m.RefIssue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ref_issue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ref_issue")
			}
			return err
		}
	}

	return nil
}

func (m *TimelineComment) contextValidateResolveDoer(ctx context.Context, formats strfmt.Registry) error {

	if m.ResolveDoer != nil {

		if swag.IsZero(m.ResolveDoer) { // not required
			return nil
		}

		if err := m.ResolveDoer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resolve_doer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resolve_doer")
			}
			return err
		}
	}

	return nil
}

func (m *TimelineComment) contextValidateTrackedTime(ctx context.Context, formats strfmt.Registry) error {

	if m.TrackedTime != nil {

		if swag.IsZero(m.TrackedTime) { // not required
			return nil
		}

		if err := m.TrackedTime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tracked_time")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tracked_time")
			}
			return err
		}
	}

	return nil
}

func (m *TimelineComment) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TimelineComment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimelineComment) UnmarshalBinary(b []byte) error {
	var res TimelineComment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
