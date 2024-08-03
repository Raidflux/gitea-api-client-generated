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

	"giteaApiClientGenerated/models"
)

// IssueRemoveIssueDependenciesReader is a Reader for the IssueRemoveIssueDependencies structure.
type IssueRemoveIssueDependenciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueRemoveIssueDependenciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIssueRemoveIssueDependenciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewIssueRemoveIssueDependenciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 423:
		result := NewIssueRemoveIssueDependenciesLocked()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /repos/{owner}/{repo}/issues/{index}/dependencies] issueRemoveIssueDependencies", response, response.Code())
	}
}

// NewIssueRemoveIssueDependenciesOK creates a IssueRemoveIssueDependenciesOK with default headers values
func NewIssueRemoveIssueDependenciesOK() *IssueRemoveIssueDependenciesOK {
	return &IssueRemoveIssueDependenciesOK{}
}

/*
IssueRemoveIssueDependenciesOK describes a response with status code 200, with default header values.

Issue
*/
type IssueRemoveIssueDependenciesOK struct {
	Payload *models.Issue
}

// IsSuccess returns true when this issue remove issue dependencies o k response has a 2xx status code
func (o *IssueRemoveIssueDependenciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this issue remove issue dependencies o k response has a 3xx status code
func (o *IssueRemoveIssueDependenciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue remove issue dependencies o k response has a 4xx status code
func (o *IssueRemoveIssueDependenciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this issue remove issue dependencies o k response has a 5xx status code
func (o *IssueRemoveIssueDependenciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this issue remove issue dependencies o k response a status code equal to that given
func (o *IssueRemoveIssueDependenciesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the issue remove issue dependencies o k response
func (o *IssueRemoveIssueDependenciesOK) Code() int {
	return 200
}

func (o *IssueRemoveIssueDependenciesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/{index}/dependencies][%d] issueRemoveIssueDependenciesOK %s", 200, payload)
}

func (o *IssueRemoveIssueDependenciesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/{index}/dependencies][%d] issueRemoveIssueDependenciesOK %s", 200, payload)
}

func (o *IssueRemoveIssueDependenciesOK) GetPayload() *models.Issue {
	return o.Payload
}

func (o *IssueRemoveIssueDependenciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Issue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIssueRemoveIssueDependenciesNotFound creates a IssueRemoveIssueDependenciesNotFound with default headers values
func NewIssueRemoveIssueDependenciesNotFound() *IssueRemoveIssueDependenciesNotFound {
	return &IssueRemoveIssueDependenciesNotFound{}
}

/*
IssueRemoveIssueDependenciesNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type IssueRemoveIssueDependenciesNotFound struct {
}

// IsSuccess returns true when this issue remove issue dependencies not found response has a 2xx status code
func (o *IssueRemoveIssueDependenciesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue remove issue dependencies not found response has a 3xx status code
func (o *IssueRemoveIssueDependenciesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue remove issue dependencies not found response has a 4xx status code
func (o *IssueRemoveIssueDependenciesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue remove issue dependencies not found response has a 5xx status code
func (o *IssueRemoveIssueDependenciesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this issue remove issue dependencies not found response a status code equal to that given
func (o *IssueRemoveIssueDependenciesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the issue remove issue dependencies not found response
func (o *IssueRemoveIssueDependenciesNotFound) Code() int {
	return 404
}

func (o *IssueRemoveIssueDependenciesNotFound) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/{index}/dependencies][%d] issueRemoveIssueDependenciesNotFound", 404)
}

func (o *IssueRemoveIssueDependenciesNotFound) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/{index}/dependencies][%d] issueRemoveIssueDependenciesNotFound", 404)
}

func (o *IssueRemoveIssueDependenciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIssueRemoveIssueDependenciesLocked creates a IssueRemoveIssueDependenciesLocked with default headers values
func NewIssueRemoveIssueDependenciesLocked() *IssueRemoveIssueDependenciesLocked {
	return &IssueRemoveIssueDependenciesLocked{}
}

/*
IssueRemoveIssueDependenciesLocked describes a response with status code 423, with default header values.

APIRepoArchivedError is an error that is raised when an archived repo should be modified
*/
type IssueRemoveIssueDependenciesLocked struct {
	Message string
	URL     string
}

// IsSuccess returns true when this issue remove issue dependencies locked response has a 2xx status code
func (o *IssueRemoveIssueDependenciesLocked) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue remove issue dependencies locked response has a 3xx status code
func (o *IssueRemoveIssueDependenciesLocked) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue remove issue dependencies locked response has a 4xx status code
func (o *IssueRemoveIssueDependenciesLocked) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue remove issue dependencies locked response has a 5xx status code
func (o *IssueRemoveIssueDependenciesLocked) IsServerError() bool {
	return false
}

// IsCode returns true when this issue remove issue dependencies locked response a status code equal to that given
func (o *IssueRemoveIssueDependenciesLocked) IsCode(code int) bool {
	return code == 423
}

// Code gets the status code for the issue remove issue dependencies locked response
func (o *IssueRemoveIssueDependenciesLocked) Code() int {
	return 423
}

func (o *IssueRemoveIssueDependenciesLocked) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/{index}/dependencies][%d] issueRemoveIssueDependenciesLocked", 423)
}

func (o *IssueRemoveIssueDependenciesLocked) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/issues/{index}/dependencies][%d] issueRemoveIssueDependenciesLocked", 423)
}

func (o *IssueRemoveIssueDependenciesLocked) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
