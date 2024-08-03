// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"giteaApiClientGenerated/models"
)

// DeleteRepoVariableReader is a Reader for the DeleteRepoVariable structure.
type DeleteRepoVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRepoVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRepoVariableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewDeleteRepoVariableCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteRepoVariableNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRepoVariableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRepoVariableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /repos/{owner}/{repo}/actions/variables/{variablename}] deleteRepoVariable", response, response.Code())
	}
}

// NewDeleteRepoVariableOK creates a DeleteRepoVariableOK with default headers values
func NewDeleteRepoVariableOK() *DeleteRepoVariableOK {
	return &DeleteRepoVariableOK{}
}

/*
DeleteRepoVariableOK describes a response with status code 200, with default header values.

ActionVariable
*/
type DeleteRepoVariableOK struct {
	Payload *models.ActionVariable
}

// IsSuccess returns true when this delete repo variable o k response has a 2xx status code
func (o *DeleteRepoVariableOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete repo variable o k response has a 3xx status code
func (o *DeleteRepoVariableOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete repo variable o k response has a 4xx status code
func (o *DeleteRepoVariableOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete repo variable o k response has a 5xx status code
func (o *DeleteRepoVariableOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete repo variable o k response a status code equal to that given
func (o *DeleteRepoVariableOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete repo variable o k response
func (o *DeleteRepoVariableOK) Code() int {
	return 200
}

func (o *DeleteRepoVariableOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/actions/variables/{variablename}][%d] deleteRepoVariableOK %s", 200, payload)
}

func (o *DeleteRepoVariableOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/actions/variables/{variablename}][%d] deleteRepoVariableOK %s", 200, payload)
}

func (o *DeleteRepoVariableOK) GetPayload() *models.ActionVariable {
	return o.Payload
}

func (o *DeleteRepoVariableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionVariable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRepoVariableCreated creates a DeleteRepoVariableCreated with default headers values
func NewDeleteRepoVariableCreated() *DeleteRepoVariableCreated {
	return &DeleteRepoVariableCreated{}
}

/*
DeleteRepoVariableCreated describes a response with status code 201, with default header values.

response when deleting a variable
*/
type DeleteRepoVariableCreated struct {
}

// IsSuccess returns true when this delete repo variable created response has a 2xx status code
func (o *DeleteRepoVariableCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete repo variable created response has a 3xx status code
func (o *DeleteRepoVariableCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete repo variable created response has a 4xx status code
func (o *DeleteRepoVariableCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete repo variable created response has a 5xx status code
func (o *DeleteRepoVariableCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this delete repo variable created response a status code equal to that given
func (o *DeleteRepoVariableCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the delete repo variable created response
func (o *DeleteRepoVariableCreated) Code() int {
	return 201
}

func (o *DeleteRepoVariableCreated) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/actions/variables/{variablename}][%d] deleteRepoVariableCreated", 201)
}

func (o *DeleteRepoVariableCreated) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/actions/variables/{variablename}][%d] deleteRepoVariableCreated", 201)
}

func (o *DeleteRepoVariableCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRepoVariableNoContent creates a DeleteRepoVariableNoContent with default headers values
func NewDeleteRepoVariableNoContent() *DeleteRepoVariableNoContent {
	return &DeleteRepoVariableNoContent{}
}

/*
DeleteRepoVariableNoContent describes a response with status code 204, with default header values.

response when deleting a variable
*/
type DeleteRepoVariableNoContent struct {
}

// IsSuccess returns true when this delete repo variable no content response has a 2xx status code
func (o *DeleteRepoVariableNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete repo variable no content response has a 3xx status code
func (o *DeleteRepoVariableNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete repo variable no content response has a 4xx status code
func (o *DeleteRepoVariableNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete repo variable no content response has a 5xx status code
func (o *DeleteRepoVariableNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete repo variable no content response a status code equal to that given
func (o *DeleteRepoVariableNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete repo variable no content response
func (o *DeleteRepoVariableNoContent) Code() int {
	return 204
}

func (o *DeleteRepoVariableNoContent) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/actions/variables/{variablename}][%d] deleteRepoVariableNoContent", 204)
}

func (o *DeleteRepoVariableNoContent) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/actions/variables/{variablename}][%d] deleteRepoVariableNoContent", 204)
}

func (o *DeleteRepoVariableNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRepoVariableBadRequest creates a DeleteRepoVariableBadRequest with default headers values
func NewDeleteRepoVariableBadRequest() *DeleteRepoVariableBadRequest {
	return &DeleteRepoVariableBadRequest{}
}

/*
DeleteRepoVariableBadRequest describes a response with status code 400, with default header values.

APIError is error format response
*/
type DeleteRepoVariableBadRequest struct {
	Message string
	URL     string
}

// IsSuccess returns true when this delete repo variable bad request response has a 2xx status code
func (o *DeleteRepoVariableBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete repo variable bad request response has a 3xx status code
func (o *DeleteRepoVariableBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete repo variable bad request response has a 4xx status code
func (o *DeleteRepoVariableBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete repo variable bad request response has a 5xx status code
func (o *DeleteRepoVariableBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete repo variable bad request response a status code equal to that given
func (o *DeleteRepoVariableBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete repo variable bad request response
func (o *DeleteRepoVariableBadRequest) Code() int {
	return 400
}

func (o *DeleteRepoVariableBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/actions/variables/{variablename}][%d] deleteRepoVariableBadRequest", 400)
}

func (o *DeleteRepoVariableBadRequest) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/actions/variables/{variablename}][%d] deleteRepoVariableBadRequest", 400)
}

func (o *DeleteRepoVariableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRepoVariableNotFound creates a DeleteRepoVariableNotFound with default headers values
func NewDeleteRepoVariableNotFound() *DeleteRepoVariableNotFound {
	return &DeleteRepoVariableNotFound{}
}

/*
DeleteRepoVariableNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type DeleteRepoVariableNotFound struct {
}

// IsSuccess returns true when this delete repo variable not found response has a 2xx status code
func (o *DeleteRepoVariableNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete repo variable not found response has a 3xx status code
func (o *DeleteRepoVariableNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete repo variable not found response has a 4xx status code
func (o *DeleteRepoVariableNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete repo variable not found response has a 5xx status code
func (o *DeleteRepoVariableNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete repo variable not found response a status code equal to that given
func (o *DeleteRepoVariableNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete repo variable not found response
func (o *DeleteRepoVariableNotFound) Code() int {
	return 404
}

func (o *DeleteRepoVariableNotFound) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/actions/variables/{variablename}][%d] deleteRepoVariableNotFound", 404)
}

func (o *DeleteRepoVariableNotFound) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/actions/variables/{variablename}][%d] deleteRepoVariableNotFound", 404)
}

func (o *DeleteRepoVariableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
