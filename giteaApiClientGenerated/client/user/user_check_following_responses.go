// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UserCheckFollowingReader is a Reader for the UserCheckFollowing structure.
type UserCheckFollowingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCheckFollowingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUserCheckFollowingNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUserCheckFollowingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /users/{username}/following/{target}] userCheckFollowing", response, response.Code())
	}
}

// NewUserCheckFollowingNoContent creates a UserCheckFollowingNoContent with default headers values
func NewUserCheckFollowingNoContent() *UserCheckFollowingNoContent {
	return &UserCheckFollowingNoContent{}
}

/*
UserCheckFollowingNoContent describes a response with status code 204, with default header values.

APIEmpty is an empty response
*/
type UserCheckFollowingNoContent struct {
}

// IsSuccess returns true when this user check following no content response has a 2xx status code
func (o *UserCheckFollowingNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user check following no content response has a 3xx status code
func (o *UserCheckFollowingNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user check following no content response has a 4xx status code
func (o *UserCheckFollowingNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this user check following no content response has a 5xx status code
func (o *UserCheckFollowingNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this user check following no content response a status code equal to that given
func (o *UserCheckFollowingNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the user check following no content response
func (o *UserCheckFollowingNoContent) Code() int {
	return 204
}

func (o *UserCheckFollowingNoContent) Error() string {
	return fmt.Sprintf("[GET /users/{username}/following/{target}][%d] userCheckFollowingNoContent", 204)
}

func (o *UserCheckFollowingNoContent) String() string {
	return fmt.Sprintf("[GET /users/{username}/following/{target}][%d] userCheckFollowingNoContent", 204)
}

func (o *UserCheckFollowingNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserCheckFollowingNotFound creates a UserCheckFollowingNotFound with default headers values
func NewUserCheckFollowingNotFound() *UserCheckFollowingNotFound {
	return &UserCheckFollowingNotFound{}
}

/*
UserCheckFollowingNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type UserCheckFollowingNotFound struct {
}

// IsSuccess returns true when this user check following not found response has a 2xx status code
func (o *UserCheckFollowingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user check following not found response has a 3xx status code
func (o *UserCheckFollowingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user check following not found response has a 4xx status code
func (o *UserCheckFollowingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user check following not found response has a 5xx status code
func (o *UserCheckFollowingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user check following not found response a status code equal to that given
func (o *UserCheckFollowingNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the user check following not found response
func (o *UserCheckFollowingNotFound) Code() int {
	return 404
}

func (o *UserCheckFollowingNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{username}/following/{target}][%d] userCheckFollowingNotFound", 404)
}

func (o *UserCheckFollowingNotFound) String() string {
	return fmt.Sprintf("[GET /users/{username}/following/{target}][%d] userCheckFollowingNotFound", 404)
}

func (o *UserCheckFollowingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
