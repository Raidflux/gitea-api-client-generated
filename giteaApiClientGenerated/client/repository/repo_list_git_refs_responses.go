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

// RepoListGitRefsReader is a Reader for the RepoListGitRefs structure.
type RepoListGitRefsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoListGitRefsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoListGitRefsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoListGitRefsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/git/refs/{ref}] repoListGitRefs", response, response.Code())
	}
}

// NewRepoListGitRefsOK creates a RepoListGitRefsOK with default headers values
func NewRepoListGitRefsOK() *RepoListGitRefsOK {
	return &RepoListGitRefsOK{}
}

/*
RepoListGitRefsOK describes a response with status code 200, with default header values.

ReferenceList
*/
type RepoListGitRefsOK struct {
	Payload []*models.Reference
}

// IsSuccess returns true when this repo list git refs o k response has a 2xx status code
func (o *RepoListGitRefsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo list git refs o k response has a 3xx status code
func (o *RepoListGitRefsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo list git refs o k response has a 4xx status code
func (o *RepoListGitRefsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo list git refs o k response has a 5xx status code
func (o *RepoListGitRefsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo list git refs o k response a status code equal to that given
func (o *RepoListGitRefsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo list git refs o k response
func (o *RepoListGitRefsOK) Code() int {
	return 200
}

func (o *RepoListGitRefsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/refs/{ref}][%d] repoListGitRefsOK %s", 200, payload)
}

func (o *RepoListGitRefsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/refs/{ref}][%d] repoListGitRefsOK %s", 200, payload)
}

func (o *RepoListGitRefsOK) GetPayload() []*models.Reference {
	return o.Payload
}

func (o *RepoListGitRefsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoListGitRefsNotFound creates a RepoListGitRefsNotFound with default headers values
func NewRepoListGitRefsNotFound() *RepoListGitRefsNotFound {
	return &RepoListGitRefsNotFound{}
}

/*
RepoListGitRefsNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoListGitRefsNotFound struct {
}

// IsSuccess returns true when this repo list git refs not found response has a 2xx status code
func (o *RepoListGitRefsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo list git refs not found response has a 3xx status code
func (o *RepoListGitRefsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo list git refs not found response has a 4xx status code
func (o *RepoListGitRefsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo list git refs not found response has a 5xx status code
func (o *RepoListGitRefsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo list git refs not found response a status code equal to that given
func (o *RepoListGitRefsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo list git refs not found response
func (o *RepoListGitRefsNotFound) Code() int {
	return 404
}

func (o *RepoListGitRefsNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/refs/{ref}][%d] repoListGitRefsNotFound", 404)
}

func (o *RepoListGitRefsNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/refs/{ref}][%d] repoListGitRefsNotFound", 404)
}

func (o *RepoListGitRefsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
