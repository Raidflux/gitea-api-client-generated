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

// RepoListGitHooksReader is a Reader for the RepoListGitHooks structure.
type RepoListGitHooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoListGitHooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoListGitHooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoListGitHooksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/hooks/git] repoListGitHooks", response, response.Code())
	}
}

// NewRepoListGitHooksOK creates a RepoListGitHooksOK with default headers values
func NewRepoListGitHooksOK() *RepoListGitHooksOK {
	return &RepoListGitHooksOK{}
}

/*
RepoListGitHooksOK describes a response with status code 200, with default header values.

GitHookList
*/
type RepoListGitHooksOK struct {
	Payload []*models.GitHook
}

// IsSuccess returns true when this repo list git hooks o k response has a 2xx status code
func (o *RepoListGitHooksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo list git hooks o k response has a 3xx status code
func (o *RepoListGitHooksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo list git hooks o k response has a 4xx status code
func (o *RepoListGitHooksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo list git hooks o k response has a 5xx status code
func (o *RepoListGitHooksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo list git hooks o k response a status code equal to that given
func (o *RepoListGitHooksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo list git hooks o k response
func (o *RepoListGitHooksOK) Code() int {
	return 200
}

func (o *RepoListGitHooksOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/hooks/git][%d] repoListGitHooksOK %s", 200, payload)
}

func (o *RepoListGitHooksOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/hooks/git][%d] repoListGitHooksOK %s", 200, payload)
}

func (o *RepoListGitHooksOK) GetPayload() []*models.GitHook {
	return o.Payload
}

func (o *RepoListGitHooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoListGitHooksNotFound creates a RepoListGitHooksNotFound with default headers values
func NewRepoListGitHooksNotFound() *RepoListGitHooksNotFound {
	return &RepoListGitHooksNotFound{}
}

/*
RepoListGitHooksNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoListGitHooksNotFound struct {
}

// IsSuccess returns true when this repo list git hooks not found response has a 2xx status code
func (o *RepoListGitHooksNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo list git hooks not found response has a 3xx status code
func (o *RepoListGitHooksNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo list git hooks not found response has a 4xx status code
func (o *RepoListGitHooksNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo list git hooks not found response has a 5xx status code
func (o *RepoListGitHooksNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo list git hooks not found response a status code equal to that given
func (o *RepoListGitHooksNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo list git hooks not found response
func (o *RepoListGitHooksNotFound) Code() int {
	return 404
}

func (o *RepoListGitHooksNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/hooks/git][%d] repoListGitHooksNotFound", 404)
}

func (o *RepoListGitHooksNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/hooks/git][%d] repoListGitHooksNotFound", 404)
}

func (o *RepoListGitHooksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
