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

// IssueReplaceLabelsReader is a Reader for the IssueReplaceLabels structure.
type IssueReplaceLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueReplaceLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIssueReplaceLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewIssueReplaceLabelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIssueReplaceLabelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /repos/{owner}/{repo}/issues/{index}/labels] issueReplaceLabels", response, response.Code())
	}
}

// NewIssueReplaceLabelsOK creates a IssueReplaceLabelsOK with default headers values
func NewIssueReplaceLabelsOK() *IssueReplaceLabelsOK {
	return &IssueReplaceLabelsOK{}
}

/*
IssueReplaceLabelsOK describes a response with status code 200, with default header values.

LabelList
*/
type IssueReplaceLabelsOK struct {
	Payload []*models.Label
}

// IsSuccess returns true when this issue replace labels o k response has a 2xx status code
func (o *IssueReplaceLabelsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this issue replace labels o k response has a 3xx status code
func (o *IssueReplaceLabelsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue replace labels o k response has a 4xx status code
func (o *IssueReplaceLabelsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this issue replace labels o k response has a 5xx status code
func (o *IssueReplaceLabelsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this issue replace labels o k response a status code equal to that given
func (o *IssueReplaceLabelsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the issue replace labels o k response
func (o *IssueReplaceLabelsOK) Code() int {
	return 200
}

func (o *IssueReplaceLabelsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /repos/{owner}/{repo}/issues/{index}/labels][%d] issueReplaceLabelsOK %s", 200, payload)
}

func (o *IssueReplaceLabelsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /repos/{owner}/{repo}/issues/{index}/labels][%d] issueReplaceLabelsOK %s", 200, payload)
}

func (o *IssueReplaceLabelsOK) GetPayload() []*models.Label {
	return o.Payload
}

func (o *IssueReplaceLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIssueReplaceLabelsForbidden creates a IssueReplaceLabelsForbidden with default headers values
func NewIssueReplaceLabelsForbidden() *IssueReplaceLabelsForbidden {
	return &IssueReplaceLabelsForbidden{}
}

/*
IssueReplaceLabelsForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type IssueReplaceLabelsForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this issue replace labels forbidden response has a 2xx status code
func (o *IssueReplaceLabelsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue replace labels forbidden response has a 3xx status code
func (o *IssueReplaceLabelsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue replace labels forbidden response has a 4xx status code
func (o *IssueReplaceLabelsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue replace labels forbidden response has a 5xx status code
func (o *IssueReplaceLabelsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this issue replace labels forbidden response a status code equal to that given
func (o *IssueReplaceLabelsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the issue replace labels forbidden response
func (o *IssueReplaceLabelsForbidden) Code() int {
	return 403
}

func (o *IssueReplaceLabelsForbidden) Error() string {
	return fmt.Sprintf("[PUT /repos/{owner}/{repo}/issues/{index}/labels][%d] issueReplaceLabelsForbidden", 403)
}

func (o *IssueReplaceLabelsForbidden) String() string {
	return fmt.Sprintf("[PUT /repos/{owner}/{repo}/issues/{index}/labels][%d] issueReplaceLabelsForbidden", 403)
}

func (o *IssueReplaceLabelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewIssueReplaceLabelsNotFound creates a IssueReplaceLabelsNotFound with default headers values
func NewIssueReplaceLabelsNotFound() *IssueReplaceLabelsNotFound {
	return &IssueReplaceLabelsNotFound{}
}

/*
IssueReplaceLabelsNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type IssueReplaceLabelsNotFound struct {
}

// IsSuccess returns true when this issue replace labels not found response has a 2xx status code
func (o *IssueReplaceLabelsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue replace labels not found response has a 3xx status code
func (o *IssueReplaceLabelsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue replace labels not found response has a 4xx status code
func (o *IssueReplaceLabelsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue replace labels not found response has a 5xx status code
func (o *IssueReplaceLabelsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this issue replace labels not found response a status code equal to that given
func (o *IssueReplaceLabelsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the issue replace labels not found response
func (o *IssueReplaceLabelsNotFound) Code() int {
	return 404
}

func (o *IssueReplaceLabelsNotFound) Error() string {
	return fmt.Sprintf("[PUT /repos/{owner}/{repo}/issues/{index}/labels][%d] issueReplaceLabelsNotFound", 404)
}

func (o *IssueReplaceLabelsNotFound) String() string {
	return fmt.Sprintf("[PUT /repos/{owner}/{repo}/issues/{index}/labels][%d] issueReplaceLabelsNotFound", 404)
}

func (o *IssueReplaceLabelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
