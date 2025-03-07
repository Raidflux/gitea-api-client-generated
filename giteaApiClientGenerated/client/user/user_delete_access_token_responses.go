// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UserDeleteAccessTokenReader is a Reader for the UserDeleteAccessToken structure.
type UserDeleteAccessTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserDeleteAccessTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUserDeleteAccessTokenNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUserDeleteAccessTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserDeleteAccessTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUserDeleteAccessTokenUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /users/{username}/tokens/{token}] userDeleteAccessToken", response, response.Code())
	}
}

// NewUserDeleteAccessTokenNoContent creates a UserDeleteAccessTokenNoContent with default headers values
func NewUserDeleteAccessTokenNoContent() *UserDeleteAccessTokenNoContent {
	return &UserDeleteAccessTokenNoContent{}
}

/*
UserDeleteAccessTokenNoContent describes a response with status code 204, with default header values.

APIEmpty is an empty response
*/
type UserDeleteAccessTokenNoContent struct {
}

// IsSuccess returns true when this user delete access token no content response has a 2xx status code
func (o *UserDeleteAccessTokenNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user delete access token no content response has a 3xx status code
func (o *UserDeleteAccessTokenNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user delete access token no content response has a 4xx status code
func (o *UserDeleteAccessTokenNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this user delete access token no content response has a 5xx status code
func (o *UserDeleteAccessTokenNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this user delete access token no content response a status code equal to that given
func (o *UserDeleteAccessTokenNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the user delete access token no content response
func (o *UserDeleteAccessTokenNoContent) Code() int {
	return 204
}

func (o *UserDeleteAccessTokenNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/{username}/tokens/{token}][%d] userDeleteAccessTokenNoContent", 204)
}

func (o *UserDeleteAccessTokenNoContent) String() string {
	return fmt.Sprintf("[DELETE /users/{username}/tokens/{token}][%d] userDeleteAccessTokenNoContent", 204)
}

func (o *UserDeleteAccessTokenNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserDeleteAccessTokenForbidden creates a UserDeleteAccessTokenForbidden with default headers values
func NewUserDeleteAccessTokenForbidden() *UserDeleteAccessTokenForbidden {
	return &UserDeleteAccessTokenForbidden{}
}

/*
UserDeleteAccessTokenForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type UserDeleteAccessTokenForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this user delete access token forbidden response has a 2xx status code
func (o *UserDeleteAccessTokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user delete access token forbidden response has a 3xx status code
func (o *UserDeleteAccessTokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user delete access token forbidden response has a 4xx status code
func (o *UserDeleteAccessTokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user delete access token forbidden response has a 5xx status code
func (o *UserDeleteAccessTokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user delete access token forbidden response a status code equal to that given
func (o *UserDeleteAccessTokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the user delete access token forbidden response
func (o *UserDeleteAccessTokenForbidden) Code() int {
	return 403
}

func (o *UserDeleteAccessTokenForbidden) Error() string {
	return fmt.Sprintf("[DELETE /users/{username}/tokens/{token}][%d] userDeleteAccessTokenForbidden", 403)
}

func (o *UserDeleteAccessTokenForbidden) String() string {
	return fmt.Sprintf("[DELETE /users/{username}/tokens/{token}][%d] userDeleteAccessTokenForbidden", 403)
}

func (o *UserDeleteAccessTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUserDeleteAccessTokenNotFound creates a UserDeleteAccessTokenNotFound with default headers values
func NewUserDeleteAccessTokenNotFound() *UserDeleteAccessTokenNotFound {
	return &UserDeleteAccessTokenNotFound{}
}

/*
UserDeleteAccessTokenNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type UserDeleteAccessTokenNotFound struct {
}

// IsSuccess returns true when this user delete access token not found response has a 2xx status code
func (o *UserDeleteAccessTokenNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user delete access token not found response has a 3xx status code
func (o *UserDeleteAccessTokenNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user delete access token not found response has a 4xx status code
func (o *UserDeleteAccessTokenNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user delete access token not found response has a 5xx status code
func (o *UserDeleteAccessTokenNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user delete access token not found response a status code equal to that given
func (o *UserDeleteAccessTokenNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the user delete access token not found response
func (o *UserDeleteAccessTokenNotFound) Code() int {
	return 404
}

func (o *UserDeleteAccessTokenNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{username}/tokens/{token}][%d] userDeleteAccessTokenNotFound", 404)
}

func (o *UserDeleteAccessTokenNotFound) String() string {
	return fmt.Sprintf("[DELETE /users/{username}/tokens/{token}][%d] userDeleteAccessTokenNotFound", 404)
}

func (o *UserDeleteAccessTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserDeleteAccessTokenUnprocessableEntity creates a UserDeleteAccessTokenUnprocessableEntity with default headers values
func NewUserDeleteAccessTokenUnprocessableEntity() *UserDeleteAccessTokenUnprocessableEntity {
	return &UserDeleteAccessTokenUnprocessableEntity{}
}

/*
UserDeleteAccessTokenUnprocessableEntity describes a response with status code 422, with default header values.

APIError is error format response
*/
type UserDeleteAccessTokenUnprocessableEntity struct {
	Message string
	URL     string
}

// IsSuccess returns true when this user delete access token unprocessable entity response has a 2xx status code
func (o *UserDeleteAccessTokenUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user delete access token unprocessable entity response has a 3xx status code
func (o *UserDeleteAccessTokenUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user delete access token unprocessable entity response has a 4xx status code
func (o *UserDeleteAccessTokenUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this user delete access token unprocessable entity response has a 5xx status code
func (o *UserDeleteAccessTokenUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this user delete access token unprocessable entity response a status code equal to that given
func (o *UserDeleteAccessTokenUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the user delete access token unprocessable entity response
func (o *UserDeleteAccessTokenUnprocessableEntity) Code() int {
	return 422
}

func (o *UserDeleteAccessTokenUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /users/{username}/tokens/{token}][%d] userDeleteAccessTokenUnprocessableEntity", 422)
}

func (o *UserDeleteAccessTokenUnprocessableEntity) String() string {
	return fmt.Sprintf("[DELETE /users/{username}/tokens/{token}][%d] userDeleteAccessTokenUnprocessableEntity", 422)
}

func (o *UserDeleteAccessTokenUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
