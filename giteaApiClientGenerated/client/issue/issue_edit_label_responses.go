// Code generated by go-swagger; DO NOT EDIT.

package issue

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

// IssueEditLabelReader is a Reader for the IssueEditLabel structure.
type IssueEditLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueEditLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIssueEditLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewIssueEditLabelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewIssueEditLabelUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /repos/{owner}/{repo}/labels/{id}] issueEditLabel", response, response.Code())
	}
}

// NewIssueEditLabelOK creates a IssueEditLabelOK with default headers values
func NewIssueEditLabelOK() *IssueEditLabelOK {
	return &IssueEditLabelOK{}
}

/*
IssueEditLabelOK describes a response with status code 200, with default header values.

Label
*/
type IssueEditLabelOK struct {
	Payload *models.Label
}

// IsSuccess returns true when this issue edit label o k response has a 2xx status code
func (o *IssueEditLabelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this issue edit label o k response has a 3xx status code
func (o *IssueEditLabelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue edit label o k response has a 4xx status code
func (o *IssueEditLabelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this issue edit label o k response has a 5xx status code
func (o *IssueEditLabelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this issue edit label o k response a status code equal to that given
func (o *IssueEditLabelOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the issue edit label o k response
func (o *IssueEditLabelOK) Code() int {
	return 200
}

func (o *IssueEditLabelOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/labels/{id}][%d] issueEditLabelOK %s", 200, payload)
}

func (o *IssueEditLabelOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/labels/{id}][%d] issueEditLabelOK %s", 200, payload)
}

func (o *IssueEditLabelOK) GetPayload() *models.Label {
	return o.Payload
}

func (o *IssueEditLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Label)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIssueEditLabelNotFound creates a IssueEditLabelNotFound with default headers values
func NewIssueEditLabelNotFound() *IssueEditLabelNotFound {
	return &IssueEditLabelNotFound{}
}

/*
IssueEditLabelNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type IssueEditLabelNotFound struct {
}

// IsSuccess returns true when this issue edit label not found response has a 2xx status code
func (o *IssueEditLabelNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue edit label not found response has a 3xx status code
func (o *IssueEditLabelNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue edit label not found response has a 4xx status code
func (o *IssueEditLabelNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue edit label not found response has a 5xx status code
func (o *IssueEditLabelNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this issue edit label not found response a status code equal to that given
func (o *IssueEditLabelNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the issue edit label not found response
func (o *IssueEditLabelNotFound) Code() int {
	return 404
}

func (o *IssueEditLabelNotFound) Error() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/labels/{id}][%d] issueEditLabelNotFound", 404)
}

func (o *IssueEditLabelNotFound) String() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/labels/{id}][%d] issueEditLabelNotFound", 404)
}

func (o *IssueEditLabelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIssueEditLabelUnprocessableEntity creates a IssueEditLabelUnprocessableEntity with default headers values
func NewIssueEditLabelUnprocessableEntity() *IssueEditLabelUnprocessableEntity {
	return &IssueEditLabelUnprocessableEntity{}
}

/*
IssueEditLabelUnprocessableEntity describes a response with status code 422, with default header values.

APIValidationError is error format response related to input validation
*/
type IssueEditLabelUnprocessableEntity struct {
	Message string
	URL     string
}

// IsSuccess returns true when this issue edit label unprocessable entity response has a 2xx status code
func (o *IssueEditLabelUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue edit label unprocessable entity response has a 3xx status code
func (o *IssueEditLabelUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue edit label unprocessable entity response has a 4xx status code
func (o *IssueEditLabelUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue edit label unprocessable entity response has a 5xx status code
func (o *IssueEditLabelUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this issue edit label unprocessable entity response a status code equal to that given
func (o *IssueEditLabelUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the issue edit label unprocessable entity response
func (o *IssueEditLabelUnprocessableEntity) Code() int {
	return 422
}

func (o *IssueEditLabelUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/labels/{id}][%d] issueEditLabelUnprocessableEntity", 422)
}

func (o *IssueEditLabelUnprocessableEntity) String() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/labels/{id}][%d] issueEditLabelUnprocessableEntity", 422)
}

func (o *IssueEditLabelUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
