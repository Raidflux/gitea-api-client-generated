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
	"github.com/go-openapi/validate"
)

// Issue Issue represents an issue in a repository
//
// swagger:model Issue
type Issue struct {

	// assignees
	Assignees []*User `json:"assignees"`

	// attachments
	Attachments []*Attachment `json:"assets"`

	// body
	Body string `json:"body,omitempty"`

	// closed
	// Format: date-time
	Closed strfmt.DateTime `json:"closed_at,omitempty"`

	// comments
	Comments int64 `json:"comments,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created_at,omitempty"`

	// deadline
	// Format: date-time
	Deadline strfmt.DateTime `json:"due_date,omitempty"`

	// HTML URL
	HTMLURL string `json:"html_url,omitempty"`

	// ID
	ID int64 `json:"id,omitempty"`

	// index
	Index int64 `json:"number,omitempty"`

	// is locked
	IsLocked bool `json:"is_locked,omitempty"`

	// labels
	Labels []*Label `json:"labels"`

	// original author
	OriginalAuthor string `json:"original_author,omitempty"`

	// original author ID
	OriginalAuthorID int64 `json:"original_author_id,omitempty"`

	// pin order
	PinOrder int64 `json:"pin_order,omitempty"`

	// ref
	Ref string `json:"ref,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// URL
	URL string `json:"url,omitempty"`

	// updated
	// Format: date-time
	Updated strfmt.DateTime `json:"updated_at,omitempty"`

	// assignee
	Assignee *User `json:"assignee,omitempty"`

	// milestone
	Milestone *Milestone `json:"milestone,omitempty"`

	// pull request
	PullRequest *PullRequestMeta `json:"pull_request,omitempty"`

	// repository
	Repository *RepositoryMeta `json:"repository,omitempty"`

	// state
	State StateType `json:"state,omitempty"`

	// user
	User *User `json:"user,omitempty"`
}

// Validate validates this issue
func (m *Issue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignees(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttachments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClosed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeadline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMilestone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePullRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
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

func (m *Issue) validateAssignees(formats strfmt.Registry) error {
	if swag.IsZero(m.Assignees) { // not required
		return nil
	}

	for i := 0; i < len(m.Assignees); i++ {
		if swag.IsZero(m.Assignees[i]) { // not required
			continue
		}

		if m.Assignees[i] != nil {
			if err := m.Assignees[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assignees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Issue) validateAttachments(formats strfmt.Registry) error {
	if swag.IsZero(m.Attachments) { // not required
		return nil
	}

	for i := 0; i < len(m.Attachments); i++ {
		if swag.IsZero(m.Attachments[i]) { // not required
			continue
		}

		if m.Attachments[i] != nil {
			if err := m.Attachments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Issue) validateClosed(formats strfmt.Registry) error {
	if swag.IsZero(m.Closed) { // not required
		return nil
	}

	if err := validate.FormatOf("closed_at", "body", "date-time", m.Closed.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Issue) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Issue) validateDeadline(formats strfmt.Registry) error {
	if swag.IsZero(m.Deadline) { // not required
		return nil
	}

	if err := validate.FormatOf("due_date", "body", "date-time", m.Deadline.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Issue) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if m.Labels[i] != nil {
			if err := m.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Issue) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Issue) validateAssignee(formats strfmt.Registry) error {
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

func (m *Issue) validateMilestone(formats strfmt.Registry) error {
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

func (m *Issue) validatePullRequest(formats strfmt.Registry) error {
	if swag.IsZero(m.PullRequest) { // not required
		return nil
	}

	if m.PullRequest != nil {
		if err := m.PullRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pull_request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pull_request")
			}
			return err
		}
	}

	return nil
}

func (m *Issue) validateRepository(formats strfmt.Registry) error {
	if swag.IsZero(m.Repository) { // not required
		return nil
	}

	if m.Repository != nil {
		if err := m.Repository.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("repository")
			}
			return err
		}
	}

	return nil
}

func (m *Issue) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("state")
		}
		return err
	}

	return nil
}

func (m *Issue) validateUser(formats strfmt.Registry) error {
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

// ContextValidate validate this issue based on the context it is used
func (m *Issue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAssignees(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAttachments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAssignee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMilestone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePullRequest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRepository(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
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

func (m *Issue) contextValidateAssignees(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Assignees); i++ {

		if m.Assignees[i] != nil {

			if swag.IsZero(m.Assignees[i]) { // not required
				return nil
			}

			if err := m.Assignees[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assignees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Issue) contextValidateAttachments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attachments); i++ {

		if m.Attachments[i] != nil {

			if swag.IsZero(m.Attachments[i]) { // not required
				return nil
			}

			if err := m.Attachments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Issue) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Labels); i++ {

		if m.Labels[i] != nil {

			if swag.IsZero(m.Labels[i]) { // not required
				return nil
			}

			if err := m.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Issue) contextValidateAssignee(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Issue) contextValidateMilestone(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Issue) contextValidatePullRequest(ctx context.Context, formats strfmt.Registry) error {

	if m.PullRequest != nil {

		if swag.IsZero(m.PullRequest) { // not required
			return nil
		}

		if err := m.PullRequest.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pull_request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pull_request")
			}
			return err
		}
	}

	return nil
}

func (m *Issue) contextValidateRepository(ctx context.Context, formats strfmt.Registry) error {

	if m.Repository != nil {

		if swag.IsZero(m.Repository) { // not required
			return nil
		}

		if err := m.Repository.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("repository")
			}
			return err
		}
	}

	return nil
}

func (m *Issue) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("state")
		}
		return err
	}

	return nil
}

func (m *Issue) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Issue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Issue) UnmarshalBinary(b []byte) error {
	var res Issue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
