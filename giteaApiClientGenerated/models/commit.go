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

// Commit Commit contains information generated from a Git commit.
//
// swagger:model Commit
type Commit struct {

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// files
	Files []*CommitAffectedFiles `json:"files"`

	// HTML URL
	HTMLURL string `json:"html_url,omitempty"`

	// parents
	Parents []*CommitMeta `json:"parents"`

	// s h a
	SHA string `json:"sha,omitempty"`

	// URL
	URL string `json:"url,omitempty"`

	// author
	Author *User `json:"author,omitempty"`

	// commit
	Commit *RepoCommit `json:"commit,omitempty"`

	// committer
	Committer *User `json:"committer,omitempty"`

	// stats
	Stats *CommitStats `json:"stats,omitempty"`
}

// Validate validates this commit
func (m *Commit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommitter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Commit) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Commit) validateFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.Files) { // not required
		return nil
	}

	for i := 0; i < len(m.Files); i++ {
		if swag.IsZero(m.Files[i]) { // not required
			continue
		}

		if m.Files[i] != nil {
			if err := m.Files[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Commit) validateParents(formats strfmt.Registry) error {
	if swag.IsZero(m.Parents) { // not required
		return nil
	}

	for i := 0; i < len(m.Parents); i++ {
		if swag.IsZero(m.Parents[i]) { // not required
			continue
		}

		if m.Parents[i] != nil {
			if err := m.Parents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Commit) validateAuthor(formats strfmt.Registry) error {
	if swag.IsZero(m.Author) { // not required
		return nil
	}

	if m.Author != nil {
		if err := m.Author.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("author")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("author")
			}
			return err
		}
	}

	return nil
}

func (m *Commit) validateCommit(formats strfmt.Registry) error {
	if swag.IsZero(m.Commit) { // not required
		return nil
	}

	if m.Commit != nil {
		if err := m.Commit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commit")
			}
			return err
		}
	}

	return nil
}

func (m *Commit) validateCommitter(formats strfmt.Registry) error {
	if swag.IsZero(m.Committer) { // not required
		return nil
	}

	if m.Committer != nil {
		if err := m.Committer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("committer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("committer")
			}
			return err
		}
	}

	return nil
}

func (m *Commit) validateStats(formats strfmt.Registry) error {
	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if m.Stats != nil {
		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this commit based on the context it is used
func (m *Commit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuthor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCommit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCommitter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Commit) contextValidateFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Files); i++ {

		if m.Files[i] != nil {

			if swag.IsZero(m.Files[i]) { // not required
				return nil
			}

			if err := m.Files[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Commit) contextValidateParents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Parents); i++ {

		if m.Parents[i] != nil {

			if swag.IsZero(m.Parents[i]) { // not required
				return nil
			}

			if err := m.Parents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Commit) contextValidateAuthor(ctx context.Context, formats strfmt.Registry) error {

	if m.Author != nil {

		if swag.IsZero(m.Author) { // not required
			return nil
		}

		if err := m.Author.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("author")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("author")
			}
			return err
		}
	}

	return nil
}

func (m *Commit) contextValidateCommit(ctx context.Context, formats strfmt.Registry) error {

	if m.Commit != nil {

		if swag.IsZero(m.Commit) { // not required
			return nil
		}

		if err := m.Commit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commit")
			}
			return err
		}
	}

	return nil
}

func (m *Commit) contextValidateCommitter(ctx context.Context, formats strfmt.Registry) error {

	if m.Committer != nil {

		if swag.IsZero(m.Committer) { // not required
			return nil
		}

		if err := m.Committer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("committer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("committer")
			}
			return err
		}
	}

	return nil
}

func (m *Commit) contextValidateStats(ctx context.Context, formats strfmt.Registry) error {

	if m.Stats != nil {

		if swag.IsZero(m.Stats) { // not required
			return nil
		}

		if err := m.Stats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Commit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Commit) UnmarshalBinary(b []byte) error {
	var res Commit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
