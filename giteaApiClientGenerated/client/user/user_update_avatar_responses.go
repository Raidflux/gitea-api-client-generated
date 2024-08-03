// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UserUpdateAvatarReader is a Reader for the UserUpdateAvatar structure.
type UserUpdateAvatarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserUpdateAvatarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUserUpdateAvatarNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /user/avatar] userUpdateAvatar", response, response.Code())
	}
}

// NewUserUpdateAvatarNoContent creates a UserUpdateAvatarNoContent with default headers values
func NewUserUpdateAvatarNoContent() *UserUpdateAvatarNoContent {
	return &UserUpdateAvatarNoContent{}
}

/*
UserUpdateAvatarNoContent describes a response with status code 204, with default header values.

APIEmpty is an empty response
*/
type UserUpdateAvatarNoContent struct {
}

// IsSuccess returns true when this user update avatar no content response has a 2xx status code
func (o *UserUpdateAvatarNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user update avatar no content response has a 3xx status code
func (o *UserUpdateAvatarNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user update avatar no content response has a 4xx status code
func (o *UserUpdateAvatarNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this user update avatar no content response has a 5xx status code
func (o *UserUpdateAvatarNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this user update avatar no content response a status code equal to that given
func (o *UserUpdateAvatarNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the user update avatar no content response
func (o *UserUpdateAvatarNoContent) Code() int {
	return 204
}

func (o *UserUpdateAvatarNoContent) Error() string {
	return fmt.Sprintf("[POST /user/avatar][%d] userUpdateAvatarNoContent", 204)
}

func (o *UserUpdateAvatarNoContent) String() string {
	return fmt.Sprintf("[POST /user/avatar][%d] userUpdateAvatarNoContent", 204)
}

func (o *UserUpdateAvatarNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
