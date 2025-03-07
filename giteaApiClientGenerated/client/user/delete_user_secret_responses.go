// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteUserSecretReader is a Reader for the DeleteUserSecret structure.
type DeleteUserSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUserSecretNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUserSecretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUserSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /user/actions/secrets/{secretname}] deleteUserSecret", response, response.Code())
	}
}

// NewDeleteUserSecretNoContent creates a DeleteUserSecretNoContent with default headers values
func NewDeleteUserSecretNoContent() *DeleteUserSecretNoContent {
	return &DeleteUserSecretNoContent{}
}

/*
DeleteUserSecretNoContent describes a response with status code 204, with default header values.

delete one secret of the user
*/
type DeleteUserSecretNoContent struct {
}

// IsSuccess returns true when this delete user secret no content response has a 2xx status code
func (o *DeleteUserSecretNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete user secret no content response has a 3xx status code
func (o *DeleteUserSecretNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user secret no content response has a 4xx status code
func (o *DeleteUserSecretNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user secret no content response has a 5xx status code
func (o *DeleteUserSecretNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user secret no content response a status code equal to that given
func (o *DeleteUserSecretNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete user secret no content response
func (o *DeleteUserSecretNoContent) Code() int {
	return 204
}

func (o *DeleteUserSecretNoContent) Error() string {
	return fmt.Sprintf("[DELETE /user/actions/secrets/{secretname}][%d] deleteUserSecretNoContent", 204)
}

func (o *DeleteUserSecretNoContent) String() string {
	return fmt.Sprintf("[DELETE /user/actions/secrets/{secretname}][%d] deleteUserSecretNoContent", 204)
}

func (o *DeleteUserSecretNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserSecretBadRequest creates a DeleteUserSecretBadRequest with default headers values
func NewDeleteUserSecretBadRequest() *DeleteUserSecretBadRequest {
	return &DeleteUserSecretBadRequest{}
}

/*
DeleteUserSecretBadRequest describes a response with status code 400, with default header values.

APIError is error format response
*/
type DeleteUserSecretBadRequest struct {
	Message string
	URL     string
}

// IsSuccess returns true when this delete user secret bad request response has a 2xx status code
func (o *DeleteUserSecretBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user secret bad request response has a 3xx status code
func (o *DeleteUserSecretBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user secret bad request response has a 4xx status code
func (o *DeleteUserSecretBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user secret bad request response has a 5xx status code
func (o *DeleteUserSecretBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user secret bad request response a status code equal to that given
func (o *DeleteUserSecretBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete user secret bad request response
func (o *DeleteUserSecretBadRequest) Code() int {
	return 400
}

func (o *DeleteUserSecretBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /user/actions/secrets/{secretname}][%d] deleteUserSecretBadRequest", 400)
}

func (o *DeleteUserSecretBadRequest) String() string {
	return fmt.Sprintf("[DELETE /user/actions/secrets/{secretname}][%d] deleteUserSecretBadRequest", 400)
}

func (o *DeleteUserSecretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteUserSecretNotFound creates a DeleteUserSecretNotFound with default headers values
func NewDeleteUserSecretNotFound() *DeleteUserSecretNotFound {
	return &DeleteUserSecretNotFound{}
}

/*
DeleteUserSecretNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type DeleteUserSecretNotFound struct {
}

// IsSuccess returns true when this delete user secret not found response has a 2xx status code
func (o *DeleteUserSecretNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user secret not found response has a 3xx status code
func (o *DeleteUserSecretNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user secret not found response has a 4xx status code
func (o *DeleteUserSecretNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user secret not found response has a 5xx status code
func (o *DeleteUserSecretNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user secret not found response a status code equal to that given
func (o *DeleteUserSecretNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete user secret not found response
func (o *DeleteUserSecretNotFound) Code() int {
	return 404
}

func (o *DeleteUserSecretNotFound) Error() string {
	return fmt.Sprintf("[DELETE /user/actions/secrets/{secretname}][%d] deleteUserSecretNotFound", 404)
}

func (o *DeleteUserSecretNotFound) String() string {
	return fmt.Sprintf("[DELETE /user/actions/secrets/{secretname}][%d] deleteUserSecretNotFound", 404)
}

func (o *DeleteUserSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
