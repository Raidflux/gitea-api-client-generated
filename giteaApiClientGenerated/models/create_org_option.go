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

// CreateOrgOption CreateOrgOption options for creating an organization
//
// swagger:model CreateOrgOption
type CreateOrgOption struct {

	// description
	Description string `json:"description,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// full name
	FullName string `json:"full_name,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// repo admin change team access
	RepoAdminChangeTeamAccess bool `json:"repo_admin_change_team_access,omitempty"`

	// user name
	// Required: true
	UserName *string `json:"username"`

	// possible values are `public` (default), `limited` or `private`
	// Enum: ["public","limited","private"]
	Visibility string `json:"visibility,omitempty"`

	// website
	Website string `json:"website,omitempty"`
}

// Validate validates this create org option
func (m *CreateOrgOption) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVisibility(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateOrgOption) validateUserName(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.UserName); err != nil {
		return err
	}

	return nil
}

var createOrgOptionTypeVisibilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["public","limited","private"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createOrgOptionTypeVisibilityPropEnum = append(createOrgOptionTypeVisibilityPropEnum, v)
	}
}

const (

	// CreateOrgOptionVisibilityPublic captures enum value "public"
	CreateOrgOptionVisibilityPublic string = "public"

	// CreateOrgOptionVisibilityLimited captures enum value "limited"
	CreateOrgOptionVisibilityLimited string = "limited"

	// CreateOrgOptionVisibilityPrivate captures enum value "private"
	CreateOrgOptionVisibilityPrivate string = "private"
)

// prop value enum
func (m *CreateOrgOption) validateVisibilityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createOrgOptionTypeVisibilityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateOrgOption) validateVisibility(formats strfmt.Registry) error {
	if swag.IsZero(m.Visibility) { // not required
		return nil
	}

	// value enum
	if err := m.validateVisibilityEnum("visibility", "body", m.Visibility); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create org option based on context it is used
func (m *CreateOrgOption) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateOrgOption) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateOrgOption) UnmarshalBinary(b []byte) error {
	var res CreateOrgOption
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
