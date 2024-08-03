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

// IssueGetIssueAttachmentReader is a Reader for the IssueGetIssueAttachment structure.
type IssueGetIssueAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IssueGetIssueAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIssueGetIssueAttachmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewIssueGetIssueAttachmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/issues/{index}/assets/{attachment_id}] issueGetIssueAttachment", response, response.Code())
	}
}

// NewIssueGetIssueAttachmentOK creates a IssueGetIssueAttachmentOK with default headers values
func NewIssueGetIssueAttachmentOK() *IssueGetIssueAttachmentOK {
	return &IssueGetIssueAttachmentOK{}
}

/*
IssueGetIssueAttachmentOK describes a response with status code 200, with default header values.

Attachment
*/
type IssueGetIssueAttachmentOK struct {
	Payload *models.Attachment
}

// IsSuccess returns true when this issue get issue attachment o k response has a 2xx status code
func (o *IssueGetIssueAttachmentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this issue get issue attachment o k response has a 3xx status code
func (o *IssueGetIssueAttachmentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue get issue attachment o k response has a 4xx status code
func (o *IssueGetIssueAttachmentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this issue get issue attachment o k response has a 5xx status code
func (o *IssueGetIssueAttachmentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this issue get issue attachment o k response a status code equal to that given
func (o *IssueGetIssueAttachmentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the issue get issue attachment o k response
func (o *IssueGetIssueAttachmentOK) Code() int {
	return 200
}

func (o *IssueGetIssueAttachmentOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/{index}/assets/{attachment_id}][%d] issueGetIssueAttachmentOK %s", 200, payload)
}

func (o *IssueGetIssueAttachmentOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/{index}/assets/{attachment_id}][%d] issueGetIssueAttachmentOK %s", 200, payload)
}

func (o *IssueGetIssueAttachmentOK) GetPayload() *models.Attachment {
	return o.Payload
}

func (o *IssueGetIssueAttachmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Attachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIssueGetIssueAttachmentNotFound creates a IssueGetIssueAttachmentNotFound with default headers values
func NewIssueGetIssueAttachmentNotFound() *IssueGetIssueAttachmentNotFound {
	return &IssueGetIssueAttachmentNotFound{}
}

/*
IssueGetIssueAttachmentNotFound describes a response with status code 404, with default header values.

APIError is error format response
*/
type IssueGetIssueAttachmentNotFound struct {
	Message string
	URL     string
}

// IsSuccess returns true when this issue get issue attachment not found response has a 2xx status code
func (o *IssueGetIssueAttachmentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this issue get issue attachment not found response has a 3xx status code
func (o *IssueGetIssueAttachmentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this issue get issue attachment not found response has a 4xx status code
func (o *IssueGetIssueAttachmentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this issue get issue attachment not found response has a 5xx status code
func (o *IssueGetIssueAttachmentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this issue get issue attachment not found response a status code equal to that given
func (o *IssueGetIssueAttachmentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the issue get issue attachment not found response
func (o *IssueGetIssueAttachmentNotFound) Code() int {
	return 404
}

func (o *IssueGetIssueAttachmentNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/{index}/assets/{attachment_id}][%d] issueGetIssueAttachmentNotFound", 404)
}

func (o *IssueGetIssueAttachmentNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issues/{index}/assets/{attachment_id}][%d] issueGetIssueAttachmentNotFound", 404)
}

func (o *IssueGetIssueAttachmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
