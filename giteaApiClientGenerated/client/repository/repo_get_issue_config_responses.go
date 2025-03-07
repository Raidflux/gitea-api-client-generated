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

// RepoGetIssueConfigReader is a Reader for the RepoGetIssueConfig structure.
type RepoGetIssueConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoGetIssueConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoGetIssueConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoGetIssueConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/issue_config] repoGetIssueConfig", response, response.Code())
	}
}

// NewRepoGetIssueConfigOK creates a RepoGetIssueConfigOK with default headers values
func NewRepoGetIssueConfigOK() *RepoGetIssueConfigOK {
	return &RepoGetIssueConfigOK{}
}

/*
RepoGetIssueConfigOK describes a response with status code 200, with default header values.

RepoIssueConfig
*/
type RepoGetIssueConfigOK struct {
	Payload *models.IssueConfig
}

// IsSuccess returns true when this repo get issue config o k response has a 2xx status code
func (o *RepoGetIssueConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo get issue config o k response has a 3xx status code
func (o *RepoGetIssueConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get issue config o k response has a 4xx status code
func (o *RepoGetIssueConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo get issue config o k response has a 5xx status code
func (o *RepoGetIssueConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get issue config o k response a status code equal to that given
func (o *RepoGetIssueConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo get issue config o k response
func (o *RepoGetIssueConfigOK) Code() int {
	return 200
}

func (o *RepoGetIssueConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issue_config][%d] repoGetIssueConfigOK %s", 200, payload)
}

func (o *RepoGetIssueConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issue_config][%d] repoGetIssueConfigOK %s", 200, payload)
}

func (o *RepoGetIssueConfigOK) GetPayload() *models.IssueConfig {
	return o.Payload
}

func (o *RepoGetIssueConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IssueConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoGetIssueConfigNotFound creates a RepoGetIssueConfigNotFound with default headers values
func NewRepoGetIssueConfigNotFound() *RepoGetIssueConfigNotFound {
	return &RepoGetIssueConfigNotFound{}
}

/*
RepoGetIssueConfigNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoGetIssueConfigNotFound struct {
}

// IsSuccess returns true when this repo get issue config not found response has a 2xx status code
func (o *RepoGetIssueConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo get issue config not found response has a 3xx status code
func (o *RepoGetIssueConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get issue config not found response has a 4xx status code
func (o *RepoGetIssueConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo get issue config not found response has a 5xx status code
func (o *RepoGetIssueConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get issue config not found response a status code equal to that given
func (o *RepoGetIssueConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo get issue config not found response
func (o *RepoGetIssueConfigNotFound) Code() int {
	return 404
}

func (o *RepoGetIssueConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issue_config][%d] repoGetIssueConfigNotFound", 404)
}

func (o *RepoGetIssueConfigNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/issue_config][%d] repoGetIssueConfigNotFound", 404)
}

func (o *RepoGetIssueConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
