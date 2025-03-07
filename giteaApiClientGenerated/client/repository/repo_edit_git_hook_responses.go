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

// RepoEditGitHookReader is a Reader for the RepoEditGitHook structure.
type RepoEditGitHookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoEditGitHookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoEditGitHookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoEditGitHookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /repos/{owner}/{repo}/hooks/git/{id}] repoEditGitHook", response, response.Code())
	}
}

// NewRepoEditGitHookOK creates a RepoEditGitHookOK with default headers values
func NewRepoEditGitHookOK() *RepoEditGitHookOK {
	return &RepoEditGitHookOK{}
}

/*
RepoEditGitHookOK describes a response with status code 200, with default header values.

GitHook
*/
type RepoEditGitHookOK struct {
	Payload *models.GitHook
}

// IsSuccess returns true when this repo edit git hook o k response has a 2xx status code
func (o *RepoEditGitHookOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo edit git hook o k response has a 3xx status code
func (o *RepoEditGitHookOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo edit git hook o k response has a 4xx status code
func (o *RepoEditGitHookOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo edit git hook o k response has a 5xx status code
func (o *RepoEditGitHookOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo edit git hook o k response a status code equal to that given
func (o *RepoEditGitHookOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo edit git hook o k response
func (o *RepoEditGitHookOK) Code() int {
	return 200
}

func (o *RepoEditGitHookOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/hooks/git/{id}][%d] repoEditGitHookOK %s", 200, payload)
}

func (o *RepoEditGitHookOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/hooks/git/{id}][%d] repoEditGitHookOK %s", 200, payload)
}

func (o *RepoEditGitHookOK) GetPayload() *models.GitHook {
	return o.Payload
}

func (o *RepoEditGitHookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitHook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoEditGitHookNotFound creates a RepoEditGitHookNotFound with default headers values
func NewRepoEditGitHookNotFound() *RepoEditGitHookNotFound {
	return &RepoEditGitHookNotFound{}
}

/*
RepoEditGitHookNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoEditGitHookNotFound struct {
}

// IsSuccess returns true when this repo edit git hook not found response has a 2xx status code
func (o *RepoEditGitHookNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo edit git hook not found response has a 3xx status code
func (o *RepoEditGitHookNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo edit git hook not found response has a 4xx status code
func (o *RepoEditGitHookNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo edit git hook not found response has a 5xx status code
func (o *RepoEditGitHookNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo edit git hook not found response a status code equal to that given
func (o *RepoEditGitHookNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo edit git hook not found response
func (o *RepoEditGitHookNotFound) Code() int {
	return 404
}

func (o *RepoEditGitHookNotFound) Error() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/hooks/git/{id}][%d] repoEditGitHookNotFound", 404)
}

func (o *RepoEditGitHookNotFound) String() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/hooks/git/{id}][%d] repoEditGitHookNotFound", 404)
}

func (o *RepoEditGitHookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
