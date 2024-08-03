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

// IssueEditIssueDeadlineReader is a Reader for the IssueEditIssueDeadline structure.
type IssueEditIssueDeadlineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueEditIssueDeadlineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIssueEditIssueDeadlineCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewIssueEditIssueDeadlineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIssueEditIssueDeadlineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /repos/{owner}/{repo}/issues/{index}/deadline] issueEditIssueDeadline", response, response.Code())
	}
}

// NewIssueEditIssueDeadlineCreated creates a IssueEditIssueDeadlineCreated with default headers values
func NewIssueEditIssueDeadlineCreated() *IssueEditIssueDeadlineCreated {
	return &IssueEditIssueDeadlineCreated{}
}

/*
IssueEditIssueDeadlineCreated describes a response with status code 201, with default header values.

IssueDeadline
*/
type IssueEditIssueDeadlineCreated struct {
	Payload *models.IssueDeadline
}

// IsSuccess returns true when this issue edit issue deadline created response has a 2xx status code
func (o *IssueEditIssueDeadlineCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this issue edit issue deadline created response has a 3xx status code
func (o *IssueEditIssueDeadlineCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue edit issue deadline created response has a 4xx status code
func (o *IssueEditIssueDeadlineCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this issue edit issue deadline created response has a 5xx status code
func (o *IssueEditIssueDeadlineCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this issue edit issue deadline created response a status code equal to that given
func (o *IssueEditIssueDeadlineCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the issue edit issue deadline created response
func (o *IssueEditIssueDeadlineCreated) Code() int {
	return 201
}

func (o *IssueEditIssueDeadlineCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/issues/{index}/deadline][%d] issueEditIssueDeadlineCreated %s", 201, payload)
}

func (o *IssueEditIssueDeadlineCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/issues/{index}/deadline][%d] issueEditIssueDeadlineCreated %s", 201, payload)
}

func (o *IssueEditIssueDeadlineCreated) GetPayload() *models.IssueDeadline {
	return o.Payload
}

func (o *IssueEditIssueDeadlineCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IssueDeadline)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIssueEditIssueDeadlineForbidden creates a IssueEditIssueDeadlineForbidden with default headers values
func NewIssueEditIssueDeadlineForbidden() *IssueEditIssueDeadlineForbidden {
	return &IssueEditIssueDeadlineForbidden{}
}

/*
IssueEditIssueDeadlineForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type IssueEditIssueDeadlineForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this issue edit issue deadline forbidden response has a 2xx status code
func (o *IssueEditIssueDeadlineForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue edit issue deadline forbidden response has a 3xx status code
func (o *IssueEditIssueDeadlineForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue edit issue deadline forbidden response has a 4xx status code
func (o *IssueEditIssueDeadlineForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue edit issue deadline forbidden response has a 5xx status code
func (o *IssueEditIssueDeadlineForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this issue edit issue deadline forbidden response a status code equal to that given
func (o *IssueEditIssueDeadlineForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the issue edit issue deadline forbidden response
func (o *IssueEditIssueDeadlineForbidden) Code() int {
	return 403
}

func (o *IssueEditIssueDeadlineForbidden) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/issues/{index}/deadline][%d] issueEditIssueDeadlineForbidden", 403)
}

func (o *IssueEditIssueDeadlineForbidden) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/issues/{index}/deadline][%d] issueEditIssueDeadlineForbidden", 403)
}

func (o *IssueEditIssueDeadlineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewIssueEditIssueDeadlineNotFound creates a IssueEditIssueDeadlineNotFound with default headers values
func NewIssueEditIssueDeadlineNotFound() *IssueEditIssueDeadlineNotFound {
	return &IssueEditIssueDeadlineNotFound{}
}

/*
IssueEditIssueDeadlineNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type IssueEditIssueDeadlineNotFound struct {
}

// IsSuccess returns true when this issue edit issue deadline not found response has a 2xx status code
func (o *IssueEditIssueDeadlineNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue edit issue deadline not found response has a 3xx status code
func (o *IssueEditIssueDeadlineNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue edit issue deadline not found response has a 4xx status code
func (o *IssueEditIssueDeadlineNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue edit issue deadline not found response has a 5xx status code
func (o *IssueEditIssueDeadlineNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this issue edit issue deadline not found response a status code equal to that given
func (o *IssueEditIssueDeadlineNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the issue edit issue deadline not found response
func (o *IssueEditIssueDeadlineNotFound) Code() int {
	return 404
}

func (o *IssueEditIssueDeadlineNotFound) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/issues/{index}/deadline][%d] issueEditIssueDeadlineNotFound", 404)
}

func (o *IssueEditIssueDeadlineNotFound) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/issues/{index}/deadline][%d] issueEditIssueDeadlineNotFound", 404)
}

func (o *IssueEditIssueDeadlineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
