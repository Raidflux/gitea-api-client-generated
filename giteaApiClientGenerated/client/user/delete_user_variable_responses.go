// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteUserVariableReader is a Reader for the DeleteUserVariable structure.
type DeleteUserVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDeleteUserVariableCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteUserVariableNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUserVariableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUserVariableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /user/actions/variables/{variablename}] deleteUserVariable", response, response.Code())
	}
}

// NewDeleteUserVariableCreated creates a DeleteUserVariableCreated with default headers values
func NewDeleteUserVariableCreated() *DeleteUserVariableCreated {
	return &DeleteUserVariableCreated{}
}

/*
DeleteUserVariableCreated describes a response with status code 201, with default header values.

response when deleting a variable
*/
type DeleteUserVariableCreated struct {
}

// IsSuccess returns true when this delete user variable created response has a 2xx status code
func (o *DeleteUserVariableCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete user variable created response has a 3xx status code
func (o *DeleteUserVariableCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user variable created response has a 4xx status code
func (o *DeleteUserVariableCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user variable created response has a 5xx status code
func (o *DeleteUserVariableCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user variable created response a status code equal to that given
func (o *DeleteUserVariableCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the delete user variable created response
func (o *DeleteUserVariableCreated) Code() int {
	return 201
}

func (o *DeleteUserVariableCreated) Error() string {
	return fmt.Sprintf("[DELETE /user/actions/variables/{variablename}][%d] deleteUserVariableCreated", 201)
}

func (o *DeleteUserVariableCreated) String() string {
	return fmt.Sprintf("[DELETE /user/actions/variables/{variablename}][%d] deleteUserVariableCreated", 201)
}

func (o *DeleteUserVariableCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserVariableNoContent creates a DeleteUserVariableNoContent with default headers values
func NewDeleteUserVariableNoContent() *DeleteUserVariableNoContent {
	return &DeleteUserVariableNoContent{}
}

/*
DeleteUserVariableNoContent describes a response with status code 204, with default header values.

response when deleting a variable
*/
type DeleteUserVariableNoContent struct {
}

// IsSuccess returns true when this delete user variable no content response has a 2xx status code
func (o *DeleteUserVariableNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete user variable no content response has a 3xx status code
func (o *DeleteUserVariableNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user variable no content response has a 4xx status code
func (o *DeleteUserVariableNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user variable no content response has a 5xx status code
func (o *DeleteUserVariableNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user variable no content response a status code equal to that given
func (o *DeleteUserVariableNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete user variable no content response
func (o *DeleteUserVariableNoContent) Code() int {
	return 204
}

func (o *DeleteUserVariableNoContent) Error() string {
	return fmt.Sprintf("[DELETE /user/actions/variables/{variablename}][%d] deleteUserVariableNoContent", 204)
}

func (o *DeleteUserVariableNoContent) String() string {
	return fmt.Sprintf("[DELETE /user/actions/variables/{variablename}][%d] deleteUserVariableNoContent", 204)
}

func (o *DeleteUserVariableNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserVariableBadRequest creates a DeleteUserVariableBadRequest with default headers values
func NewDeleteUserVariableBadRequest() *DeleteUserVariableBadRequest {
	return &DeleteUserVariableBadRequest{}
}

/*
DeleteUserVariableBadRequest describes a response with status code 400, with default header values.

APIError is error format response
*/
type DeleteUserVariableBadRequest struct {
	Message string
	URL     string
}

// IsSuccess returns true when this delete user variable bad request response has a 2xx status code
func (o *DeleteUserVariableBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user variable bad request response has a 3xx status code
func (o *DeleteUserVariableBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user variable bad request response has a 4xx status code
func (o *DeleteUserVariableBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user variable bad request response has a 5xx status code
func (o *DeleteUserVariableBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user variable bad request response a status code equal to that given
func (o *DeleteUserVariableBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete user variable bad request response
func (o *DeleteUserVariableBadRequest) Code() int {
	return 400
}

func (o *DeleteUserVariableBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /user/actions/variables/{variablename}][%d] deleteUserVariableBadRequest", 400)
}

func (o *DeleteUserVariableBadRequest) String() string {
	return fmt.Sprintf("[DELETE /user/actions/variables/{variablename}][%d] deleteUserVariableBadRequest", 400)
}

func (o *DeleteUserVariableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteUserVariableNotFound creates a DeleteUserVariableNotFound with default headers values
func NewDeleteUserVariableNotFound() *DeleteUserVariableNotFound {
	return &DeleteUserVariableNotFound{}
}

/*
DeleteUserVariableNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type DeleteUserVariableNotFound struct {
}

// IsSuccess returns true when this delete user variable not found response has a 2xx status code
func (o *DeleteUserVariableNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user variable not found response has a 3xx status code
func (o *DeleteUserVariableNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user variable not found response has a 4xx status code
func (o *DeleteUserVariableNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user variable not found response has a 5xx status code
func (o *DeleteUserVariableNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user variable not found response a status code equal to that given
func (o *DeleteUserVariableNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete user variable not found response
func (o *DeleteUserVariableNotFound) Code() int {
	return 404
}

func (o *DeleteUserVariableNotFound) Error() string {
	return fmt.Sprintf("[DELETE /user/actions/variables/{variablename}][%d] deleteUserVariableNotFound", 404)
}

func (o *DeleteUserVariableNotFound) String() string {
	return fmt.Sprintf("[DELETE /user/actions/variables/{variablename}][%d] deleteUserVariableNotFound", 404)
}

func (o *DeleteUserVariableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
