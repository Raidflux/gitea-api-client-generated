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

// RepoListAllGitRefsReader is a Reader for the RepoListAllGitRefs structure.
type RepoListAllGitRefsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoListAllGitRefsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoListAllGitRefsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoListAllGitRefsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/git/refs] repoListAllGitRefs", response, response.Code())
	}
}

// NewRepoListAllGitRefsOK creates a RepoListAllGitRefsOK with default headers values
func NewRepoListAllGitRefsOK() *RepoListAllGitRefsOK {
	return &RepoListAllGitRefsOK{}
}

/*
RepoListAllGitRefsOK describes a response with status code 200, with default header values.

ReferenceList
*/
type RepoListAllGitRefsOK struct {
	Payload []*models.Reference
}

// IsSuccess returns true when this repo list all git refs o k response has a 2xx status code
func (o *RepoListAllGitRefsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo list all git refs o k response has a 3xx status code
func (o *RepoListAllGitRefsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo list all git refs o k response has a 4xx status code
func (o *RepoListAllGitRefsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo list all git refs o k response has a 5xx status code
func (o *RepoListAllGitRefsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo list all git refs o k response a status code equal to that given
func (o *RepoListAllGitRefsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo list all git refs o k response
func (o *RepoListAllGitRefsOK) Code() int {
	return 200
}

func (o *RepoListAllGitRefsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/refs][%d] repoListAllGitRefsOK %s", 200, payload)
}

func (o *RepoListAllGitRefsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/refs][%d] repoListAllGitRefsOK %s", 200, payload)
}

func (o *RepoListAllGitRefsOK) GetPayload() []*models.Reference {
	return o.Payload
}

func (o *RepoListAllGitRefsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoListAllGitRefsNotFound creates a RepoListAllGitRefsNotFound with default headers values
func NewRepoListAllGitRefsNotFound() *RepoListAllGitRefsNotFound {
	return &RepoListAllGitRefsNotFound{}
}

/*
RepoListAllGitRefsNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoListAllGitRefsNotFound struct {
}

// IsSuccess returns true when this repo list all git refs not found response has a 2xx status code
func (o *RepoListAllGitRefsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo list all git refs not found response has a 3xx status code
func (o *RepoListAllGitRefsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo list all git refs not found response has a 4xx status code
func (o *RepoListAllGitRefsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo list all git refs not found response has a 5xx status code
func (o *RepoListAllGitRefsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo list all git refs not found response a status code equal to that given
func (o *RepoListAllGitRefsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo list all git refs not found response
func (o *RepoListAllGitRefsNotFound) Code() int {
	return 404
}

func (o *RepoListAllGitRefsNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/refs][%d] repoListAllGitRefsNotFound", 404)
}

func (o *RepoListAllGitRefsNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/refs][%d] repoListAllGitRefsNotFound", 404)
}

func (o *RepoListAllGitRefsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
