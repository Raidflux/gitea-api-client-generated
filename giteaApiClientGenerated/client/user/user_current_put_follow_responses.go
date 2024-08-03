// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UserCurrentPutFollowReader is a Reader for the UserCurrentPutFollow structure.
type UserCurrentPutFollowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCurrentPutFollowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUserCurrentPutFollowNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUserCurrentPutFollowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserCurrentPutFollowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /user/following/{username}] userCurrentPutFollow", response, response.Code())
	}
}

// NewUserCurrentPutFollowNoContent creates a UserCurrentPutFollowNoContent with default headers values
func NewUserCurrentPutFollowNoContent() *UserCurrentPutFollowNoContent {
	return &UserCurrentPutFollowNoContent{}
}

/*
UserCurrentPutFollowNoContent describes a response with status code 204, with default header values.

APIEmpty is an empty response
*/
type UserCurrentPutFollowNoContent struct {
}

// IsSuccess returns true when this user current put follow no content response has a 2xx status code
func (o *UserCurrentPutFollowNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user current put follow no content response has a 3xx status code
func (o *UserCurrentPutFollowNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user current put follow no content response has a 4xx status code
func (o *UserCurrentPutFollowNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this user current put follow no content response has a 5xx status code
func (o *UserCurrentPutFollowNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this user current put follow no content response a status code equal to that given
func (o *UserCurrentPutFollowNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the user current put follow no content response
func (o *UserCurrentPutFollowNoContent) Code() int {
	return 204
}

func (o *UserCurrentPutFollowNoContent) Error() string {
	return fmt.Sprintf("[PUT /user/following/{username}][%d] userCurrentPutFollowNoContent", 204)
}

func (o *UserCurrentPutFollowNoContent) String() string {
	return fmt.Sprintf("[PUT /user/following/{username}][%d] userCurrentPutFollowNoContent", 204)
}

func (o *UserCurrentPutFollowNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserCurrentPutFollowForbidden creates a UserCurrentPutFollowForbidden with default headers values
func NewUserCurrentPutFollowForbidden() *UserCurrentPutFollowForbidden {
	return &UserCurrentPutFollowForbidden{}
}

/*
UserCurrentPutFollowForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type UserCurrentPutFollowForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this user current put follow forbidden response has a 2xx status code
func (o *UserCurrentPutFollowForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user current put follow forbidden response has a 3xx status code
func (o *UserCurrentPutFollowForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user current put follow forbidden response has a 4xx status code
func (o *UserCurrentPutFollowForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user current put follow forbidden response has a 5xx status code
func (o *UserCurrentPutFollowForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user current put follow forbidden response a status code equal to that given
func (o *UserCurrentPutFollowForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the user current put follow forbidden response
func (o *UserCurrentPutFollowForbidden) Code() int {
	return 403
}

func (o *UserCurrentPutFollowForbidden) Error() string {
	return fmt.Sprintf("[PUT /user/following/{username}][%d] userCurrentPutFollowForbidden", 403)
}

func (o *UserCurrentPutFollowForbidden) String() string {
	return fmt.Sprintf("[PUT /user/following/{username}][%d] userCurrentPutFollowForbidden", 403)
}

func (o *UserCurrentPutFollowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header message
	hdrMessage := response.GetHeader("message")

	if hdrMessage != "" {
		o.Message = hdrMessage
	}

	// hydrates response header url
	hdrURL := response.GetHeader("url")

	if hdrURL != "" {
		o.URL = hdrURL
	}

	return nil
}

// NewUserCurrentPutFollowNotFound creates a UserCurrentPutFollowNotFound with default headers values
func NewUserCurrentPutFollowNotFound() *UserCurrentPutFollowNotFound {
	return &UserCurrentPutFollowNotFound{}
}

/*
UserCurrentPutFollowNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type UserCurrentPutFollowNotFound struct {
}

// IsSuccess returns true when this user current put follow not found response has a 2xx status code
func (o *UserCurrentPutFollowNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user current put follow not found response has a 3xx status code
func (o *UserCurrentPutFollowNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user current put follow not found response has a 4xx status code
func (o *UserCurrentPutFollowNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user current put follow not found response has a 5xx status code
func (o *UserCurrentPutFollowNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user current put follow not found response a status code equal to that given
func (o *UserCurrentPutFollowNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the user current put follow not found response
func (o *UserCurrentPutFollowNotFound) Code() int {
	return 404
}

func (o *UserCurrentPutFollowNotFound) Error() string {
	return fmt.Sprintf("[PUT /user/following/{username}][%d] userCurrentPutFollowNotFound", 404)
}

func (o *UserCurrentPutFollowNotFound) String() string {
	return fmt.Sprintf("[PUT /user/following/{username}][%d] userCurrentPutFollowNotFound", 404)
}

func (o *UserCurrentPutFollowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
