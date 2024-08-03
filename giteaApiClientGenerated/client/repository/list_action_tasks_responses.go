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

	"github.com/Raidflux/gitea-api-client-generated/giteaApiClientGenerated/models"
)

// ListActionTasksReader is a Reader for the ListActionTasks structure.
type ListActionTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListActionTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListActionTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListActionTasksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListActionTasksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListActionTasksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListActionTasksConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewListActionTasksUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/actions/tasks] ListActionTasks", response, response.Code())
	}
}

// NewListActionTasksOK creates a ListActionTasksOK with default headers values
func NewListActionTasksOK() *ListActionTasksOK {
	return &ListActionTasksOK{}
}

/*
ListActionTasksOK describes a response with status code 200, with default header values.

TasksList
*/
type ListActionTasksOK struct {
	Payload *models.ActionTaskResponse
}

// IsSuccess returns true when this list action tasks o k response has a 2xx status code
func (o *ListActionTasksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list action tasks o k response has a 3xx status code
func (o *ListActionTasksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list action tasks o k response has a 4xx status code
func (o *ListActionTasksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list action tasks o k response has a 5xx status code
func (o *ListActionTasksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list action tasks o k response a status code equal to that given
func (o *ListActionTasksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list action tasks o k response
func (o *ListActionTasksOK) Code() int {
	return 200
}

func (o *ListActionTasksOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/actions/tasks][%d] listActionTasksOK %s", 200, payload)
}

func (o *ListActionTasksOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/actions/tasks][%d] listActionTasksOK %s", 200, payload)
}

func (o *ListActionTasksOK) GetPayload() *models.ActionTaskResponse {
	return o.Payload
}

func (o *ListActionTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionTaskResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListActionTasksBadRequest creates a ListActionTasksBadRequest with default headers values
func NewListActionTasksBadRequest() *ListActionTasksBadRequest {
	return &ListActionTasksBadRequest{}
}

/*
ListActionTasksBadRequest describes a response with status code 400, with default header values.

APIError is error format response
*/
type ListActionTasksBadRequest struct {
	Message string
	URL     string
}

// IsSuccess returns true when this list action tasks bad request response has a 2xx status code
func (o *ListActionTasksBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list action tasks bad request response has a 3xx status code
func (o *ListActionTasksBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list action tasks bad request response has a 4xx status code
func (o *ListActionTasksBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list action tasks bad request response has a 5xx status code
func (o *ListActionTasksBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list action tasks bad request response a status code equal to that given
func (o *ListActionTasksBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list action tasks bad request response
func (o *ListActionTasksBadRequest) Code() int {
	return 400
}

func (o *ListActionTasksBadRequest) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/actions/tasks][%d] listActionTasksBadRequest", 400)
}

func (o *ListActionTasksBadRequest) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/actions/tasks][%d] listActionTasksBadRequest", 400)
}

func (o *ListActionTasksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListActionTasksForbidden creates a ListActionTasksForbidden with default headers values
func NewListActionTasksForbidden() *ListActionTasksForbidden {
	return &ListActionTasksForbidden{}
}

/*
ListActionTasksForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type ListActionTasksForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this list action tasks forbidden response has a 2xx status code
func (o *ListActionTasksForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list action tasks forbidden response has a 3xx status code
func (o *ListActionTasksForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list action tasks forbidden response has a 4xx status code
func (o *ListActionTasksForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list action tasks forbidden response has a 5xx status code
func (o *ListActionTasksForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list action tasks forbidden response a status code equal to that given
func (o *ListActionTasksForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list action tasks forbidden response
func (o *ListActionTasksForbidden) Code() int {
	return 403
}

func (o *ListActionTasksForbidden) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/actions/tasks][%d] listActionTasksForbidden", 403)
}

func (o *ListActionTasksForbidden) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/actions/tasks][%d] listActionTasksForbidden", 403)
}

func (o *ListActionTasksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListActionTasksNotFound creates a ListActionTasksNotFound with default headers values
func NewListActionTasksNotFound() *ListActionTasksNotFound {
	return &ListActionTasksNotFound{}
}

/*
ListActionTasksNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type ListActionTasksNotFound struct {
}

// IsSuccess returns true when this list action tasks not found response has a 2xx status code
func (o *ListActionTasksNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list action tasks not found response has a 3xx status code
func (o *ListActionTasksNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list action tasks not found response has a 4xx status code
func (o *ListActionTasksNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list action tasks not found response has a 5xx status code
func (o *ListActionTasksNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list action tasks not found response a status code equal to that given
func (o *ListActionTasksNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list action tasks not found response
func (o *ListActionTasksNotFound) Code() int {
	return 404
}

func (o *ListActionTasksNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/actions/tasks][%d] listActionTasksNotFound", 404)
}

func (o *ListActionTasksNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/actions/tasks][%d] listActionTasksNotFound", 404)
}

func (o *ListActionTasksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListActionTasksConflict creates a ListActionTasksConflict with default headers values
func NewListActionTasksConflict() *ListActionTasksConflict {
	return &ListActionTasksConflict{}
}

/*
ListActionTasksConflict describes a response with status code 409, with default header values.

APIConflict is a conflict empty response
*/
type ListActionTasksConflict struct {
}

// IsSuccess returns true when this list action tasks conflict response has a 2xx status code
func (o *ListActionTasksConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list action tasks conflict response has a 3xx status code
func (o *ListActionTasksConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list action tasks conflict response has a 4xx status code
func (o *ListActionTasksConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this list action tasks conflict response has a 5xx status code
func (o *ListActionTasksConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this list action tasks conflict response a status code equal to that given
func (o *ListActionTasksConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the list action tasks conflict response
func (o *ListActionTasksConflict) Code() int {
	return 409
}

func (o *ListActionTasksConflict) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/actions/tasks][%d] listActionTasksConflict", 409)
}

func (o *ListActionTasksConflict) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/actions/tasks][%d] listActionTasksConflict", 409)
}

func (o *ListActionTasksConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListActionTasksUnprocessableEntity creates a ListActionTasksUnprocessableEntity with default headers values
func NewListActionTasksUnprocessableEntity() *ListActionTasksUnprocessableEntity {
	return &ListActionTasksUnprocessableEntity{}
}

/*
ListActionTasksUnprocessableEntity describes a response with status code 422, with default header values.

APIValidationError is error format response related to input validation
*/
type ListActionTasksUnprocessableEntity struct {
	Message string
	URL     string
}

// IsSuccess returns true when this list action tasks unprocessable entity response has a 2xx status code
func (o *ListActionTasksUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list action tasks unprocessable entity response has a 3xx status code
func (o *ListActionTasksUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list action tasks unprocessable entity response has a 4xx status code
func (o *ListActionTasksUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this list action tasks unprocessable entity response has a 5xx status code
func (o *ListActionTasksUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this list action tasks unprocessable entity response a status code equal to that given
func (o *ListActionTasksUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the list action tasks unprocessable entity response
func (o *ListActionTasksUnprocessableEntity) Code() int {
	return 422
}

func (o *ListActionTasksUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/actions/tasks][%d] listActionTasksUnprocessableEntity", 422)
}

func (o *ListActionTasksUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/actions/tasks][%d] listActionTasksUnprocessableEntity", 422)
}

func (o *ListActionTasksUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
