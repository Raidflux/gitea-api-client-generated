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

// RepoGetIssueTemplatesReader is a Reader for the RepoGetIssueTemplates structure.
type RepoGetIssueTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoGetIssueTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoGetIssueTemplatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoGetIssueTemplatesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/issue_templates] repoGetIssueTemplates", response, response.Code())
	}
}

// NewRepoGetIssueTemplatesOK creates a RepoGetIssueTemplatesOK with default headers values
func NewRepoGetIssueTemplatesOK() *RepoGetIssueTemplatesOK {
	return &RepoGetIssueTemplatesOK{}
}

/*
RepoGetIssueTemplatesOK describes a response with status code 200, with default header values.

IssueTemplates
*/
type RepoGetIssueTemplatesOK struct {
	Payload []*models.IssueTemplate
}

// IsSuccess returns true when this repo get issue templates o k response has a 2xx status code
func (o *RepoGetIssueTemplatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo get issue templates o k response has a 3xx status code
func (o *RepoGetIssueTemplatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get issue templates o k response has a 4xx status code
func (o *RepoGetIssueTemplatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo get issue templates o k response has a 5xx status code
func (o *RepoGetIssueTemplatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get issue templates o k response a status code equal to that given
func (o *RepoGetIssueTemplatesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo get issue templates o k response
func (o *RepoGetIssueTemplatesOK) Code() int {
	return 200
}

func (o *RepoGetIssueTemplatesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issue_templates][%d] repoGetIssueTemplatesOK %s", 200, payload)
}

func (o *RepoGetIssueTemplatesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issue_templates][%d] repoGetIssueTemplatesOK %s", 200, payload)
}

func (o *RepoGetIssueTemplatesOK) GetPayload() []*models.IssueTemplate {
	return o.Payload
}

func (o *RepoGetIssueTemplatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoGetIssueTemplatesNotFound creates a RepoGetIssueTemplatesNotFound with default headers values
func NewRepoGetIssueTemplatesNotFound() *RepoGetIssueTemplatesNotFound {
	return &RepoGetIssueTemplatesNotFound{}
}

/*
RepoGetIssueTemplatesNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoGetIssueTemplatesNotFound struct {
}

// IsSuccess returns true when this repo get issue templates not found response has a 2xx status code
func (o *RepoGetIssueTemplatesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo get issue templates not found response has a 3xx status code
func (o *RepoGetIssueTemplatesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get issue templates not found response has a 4xx status code
func (o *RepoGetIssueTemplatesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo get issue templates not found response has a 5xx status code
func (o *RepoGetIssueTemplatesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get issue templates not found response a status code equal to that given
func (o *RepoGetIssueTemplatesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo get issue templates not found response
func (o *RepoGetIssueTemplatesNotFound) Code() int {
	return 404
}

func (o *RepoGetIssueTemplatesNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issue_templates][%d] repoGetIssueTemplatesNotFound", 404)
}

func (o *RepoGetIssueTemplatesNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issue_templates][%d] repoGetIssueTemplatesNotFound", 404)
}

func (o *RepoGetIssueTemplatesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
