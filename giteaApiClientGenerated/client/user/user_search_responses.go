// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/Raidflux/gitea-api-client-generated/giteaApiClientGenerated/models"
)

// UserSearchReader is a Reader for the UserSearch structure.
type UserSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /users/search] userSearch", response, response.Code())
	}
}

// NewUserSearchOK creates a UserSearchOK with default headers values
func NewUserSearchOK() *UserSearchOK {
	return &UserSearchOK{}
}

/*
UserSearchOK describes a response with status code 200, with default header values.

SearchResults of a successful search
*/
type UserSearchOK struct {
	Payload *UserSearchOKBody
}

// IsSuccess returns true when this user search o k response has a 2xx status code
func (o *UserSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user search o k response has a 3xx status code
func (o *UserSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user search o k response has a 4xx status code
func (o *UserSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user search o k response has a 5xx status code
func (o *UserSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user search o k response a status code equal to that given
func (o *UserSearchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user search o k response
func (o *UserSearchOK) Code() int {
	return 200
}

func (o *UserSearchOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/search][%d] userSearchOK %s", 200, payload)
}

func (o *UserSearchOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/search][%d] userSearchOK %s", 200, payload)
}

func (o *UserSearchOK) GetPayload() *UserSearchOKBody {
	return o.Payload
}

func (o *UserSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UserSearchOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UserSearchOKBody user search o k body
swagger:model UserSearchOKBody
*/
type UserSearchOKBody struct {

	// data
	Data []*models.User `json:"data"`

	// ok
	Ok bool `json:"ok,omitempty"`
}

// Validate validates this user search o k body
func (o *UserSearchOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UserSearchOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userSearchOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userSearchOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this user search o k body based on the context it is used
func (o *UserSearchOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UserSearchOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if swag.IsZero(o.Data[i]) { // not required
				return nil
			}

			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userSearchOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userSearchOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UserSearchOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserSearchOKBody) UnmarshalBinary(b []byte) error {
	var res UserSearchOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
