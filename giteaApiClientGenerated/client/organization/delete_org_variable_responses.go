// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Raidflux/gitea-api-client-generated/giteaApiClientGenerated/models"
)

// DeleteOrgVariableReader is a Reader for the DeleteOrgVariable structure.
type DeleteOrgVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrgVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOrgVariableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewDeleteOrgVariableCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteOrgVariableNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOrgVariableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOrgVariableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /orgs/{org}/actions/variables/{variablename}] deleteOrgVariable", response, response.Code())
	}
}

// NewDeleteOrgVariableOK creates a DeleteOrgVariableOK with default headers values
func NewDeleteOrgVariableOK() *DeleteOrgVariableOK {
	return &DeleteOrgVariableOK{}
}

/*
DeleteOrgVariableOK describes a response with status code 200, with default header values.

ActionVariable
*/
type DeleteOrgVariableOK struct {
	Payload *models.ActionVariable
}

// IsSuccess returns true when this delete org variable o k response has a 2xx status code
func (o *DeleteOrgVariableOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete org variable o k response has a 3xx status code
func (o *DeleteOrgVariableOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete org variable o k response has a 4xx status code
func (o *DeleteOrgVariableOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete org variable o k response has a 5xx status code
func (o *DeleteOrgVariableOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete org variable o k response a status code equal to that given
func (o *DeleteOrgVariableOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete org variable o k response
func (o *DeleteOrgVariableOK) Code() int {
	return 200
}

func (o *DeleteOrgVariableOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /orgs/{org}/actions/variables/{variablename}][%d] deleteOrgVariableOK %s", 200, payload)
}

func (o *DeleteOrgVariableOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /orgs/{org}/actions/variables/{variablename}][%d] deleteOrgVariableOK %s", 200, payload)
}

func (o *DeleteOrgVariableOK) GetPayload() *models.ActionVariable {
	return o.Payload
}

func (o *DeleteOrgVariableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionVariable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrgVariableCreated creates a DeleteOrgVariableCreated with default headers values
func NewDeleteOrgVariableCreated() *DeleteOrgVariableCreated {
	return &DeleteOrgVariableCreated{}
}

/*
DeleteOrgVariableCreated describes a response with status code 201, with default header values.

response when deleting a variable
*/
type DeleteOrgVariableCreated struct {
}

// IsSuccess returns true when this delete org variable created response has a 2xx status code
func (o *DeleteOrgVariableCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete org variable created response has a 3xx status code
func (o *DeleteOrgVariableCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete org variable created response has a 4xx status code
func (o *DeleteOrgVariableCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete org variable created response has a 5xx status code
func (o *DeleteOrgVariableCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this delete org variable created response a status code equal to that given
func (o *DeleteOrgVariableCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the delete org variable created response
func (o *DeleteOrgVariableCreated) Code() int {
	return 201
}

func (o *DeleteOrgVariableCreated) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/actions/variables/{variablename}][%d] deleteOrgVariableCreated", 201)
}

func (o *DeleteOrgVariableCreated) String() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/actions/variables/{variablename}][%d] deleteOrgVariableCreated", 201)
}

func (o *DeleteOrgVariableCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrgVariableNoContent creates a DeleteOrgVariableNoContent with default headers values
func NewDeleteOrgVariableNoContent() *DeleteOrgVariableNoContent {
	return &DeleteOrgVariableNoContent{}
}

/*
DeleteOrgVariableNoContent describes a response with status code 204, with default header values.

response when deleting a variable
*/
type DeleteOrgVariableNoContent struct {
}

// IsSuccess returns true when this delete org variable no content response has a 2xx status code
func (o *DeleteOrgVariableNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete org variable no content response has a 3xx status code
func (o *DeleteOrgVariableNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete org variable no content response has a 4xx status code
func (o *DeleteOrgVariableNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete org variable no content response has a 5xx status code
func (o *DeleteOrgVariableNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete org variable no content response a status code equal to that given
func (o *DeleteOrgVariableNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete org variable no content response
func (o *DeleteOrgVariableNoContent) Code() int {
	return 204
}

func (o *DeleteOrgVariableNoContent) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/actions/variables/{variablename}][%d] deleteOrgVariableNoContent", 204)
}

func (o *DeleteOrgVariableNoContent) String() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/actions/variables/{variablename}][%d] deleteOrgVariableNoContent", 204)
}

func (o *DeleteOrgVariableNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrgVariableBadRequest creates a DeleteOrgVariableBadRequest with default headers values
func NewDeleteOrgVariableBadRequest() *DeleteOrgVariableBadRequest {
	return &DeleteOrgVariableBadRequest{}
}

/*
DeleteOrgVariableBadRequest describes a response with status code 400, with default header values.

APIError is error format response
*/
type DeleteOrgVariableBadRequest struct {
	Message string
	URL     string
}

// IsSuccess returns true when this delete org variable bad request response has a 2xx status code
func (o *DeleteOrgVariableBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete org variable bad request response has a 3xx status code
func (o *DeleteOrgVariableBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete org variable bad request response has a 4xx status code
func (o *DeleteOrgVariableBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete org variable bad request response has a 5xx status code
func (o *DeleteOrgVariableBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete org variable bad request response a status code equal to that given
func (o *DeleteOrgVariableBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete org variable bad request response
func (o *DeleteOrgVariableBadRequest) Code() int {
	return 400
}

func (o *DeleteOrgVariableBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/actions/variables/{variablename}][%d] deleteOrgVariableBadRequest", 400)
}

func (o *DeleteOrgVariableBadRequest) String() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/actions/variables/{variablename}][%d] deleteOrgVariableBadRequest", 400)
}

func (o *DeleteOrgVariableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteOrgVariableNotFound creates a DeleteOrgVariableNotFound with default headers values
func NewDeleteOrgVariableNotFound() *DeleteOrgVariableNotFound {
	return &DeleteOrgVariableNotFound{}
}

/*
DeleteOrgVariableNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type DeleteOrgVariableNotFound struct {
}

// IsSuccess returns true when this delete org variable not found response has a 2xx status code
func (o *DeleteOrgVariableNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete org variable not found response has a 3xx status code
func (o *DeleteOrgVariableNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete org variable not found response has a 4xx status code
func (o *DeleteOrgVariableNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete org variable not found response has a 5xx status code
func (o *DeleteOrgVariableNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete org variable not found response a status code equal to that given
func (o *DeleteOrgVariableNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete org variable not found response
func (o *DeleteOrgVariableNotFound) Code() int {
	return 404
}

func (o *DeleteOrgVariableNotFound) Error() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/actions/variables/{variablename}][%d] deleteOrgVariableNotFound", 404)
}

func (o *DeleteOrgVariableNotFound) String() string {
	return fmt.Sprintf("[DELETE /orgs/{org}/actions/variables/{variablename}][%d] deleteOrgVariableNotFound", 404)
}

func (o *DeleteOrgVariableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
