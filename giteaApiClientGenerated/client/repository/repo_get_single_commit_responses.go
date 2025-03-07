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

// RepoGetSingleCommitReader is a Reader for the RepoGetSingleCommit structure.
type RepoGetSingleCommitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoGetSingleCommitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoGetSingleCommitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoGetSingleCommitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRepoGetSingleCommitUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/git/commits/{sha}] repoGetSingleCommit", response, response.Code())
	}
}

// NewRepoGetSingleCommitOK creates a RepoGetSingleCommitOK with default headers values
func NewRepoGetSingleCommitOK() *RepoGetSingleCommitOK {
	return &RepoGetSingleCommitOK{}
}

/*
RepoGetSingleCommitOK describes a response with status code 200, with default header values.

Commit
*/
type RepoGetSingleCommitOK struct {
	Payload *models.Commit
}

// IsSuccess returns true when this repo get single commit o k response has a 2xx status code
func (o *RepoGetSingleCommitOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo get single commit o k response has a 3xx status code
func (o *RepoGetSingleCommitOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get single commit o k response has a 4xx status code
func (o *RepoGetSingleCommitOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo get single commit o k response has a 5xx status code
func (o *RepoGetSingleCommitOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get single commit o k response a status code equal to that given
func (o *RepoGetSingleCommitOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo get single commit o k response
func (o *RepoGetSingleCommitOK) Code() int {
	return 200
}

func (o *RepoGetSingleCommitOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/commits/{sha}][%d] repoGetSingleCommitOK %s", 200, payload)
}

func (o *RepoGetSingleCommitOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/commits/{sha}][%d] repoGetSingleCommitOK %s", 200, payload)
}

func (o *RepoGetSingleCommitOK) GetPayload() *models.Commit {
	return o.Payload
}

func (o *RepoGetSingleCommitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Commit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoGetSingleCommitNotFound creates a RepoGetSingleCommitNotFound with default headers values
func NewRepoGetSingleCommitNotFound() *RepoGetSingleCommitNotFound {
	return &RepoGetSingleCommitNotFound{}
}

/*
RepoGetSingleCommitNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoGetSingleCommitNotFound struct {
}

// IsSuccess returns true when this repo get single commit not found response has a 2xx status code
func (o *RepoGetSingleCommitNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo get single commit not found response has a 3xx status code
func (o *RepoGetSingleCommitNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get single commit not found response has a 4xx status code
func (o *RepoGetSingleCommitNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo get single commit not found response has a 5xx status code
func (o *RepoGetSingleCommitNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get single commit not found response a status code equal to that given
func (o *RepoGetSingleCommitNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo get single commit not found response
func (o *RepoGetSingleCommitNotFound) Code() int {
	return 404
}

func (o *RepoGetSingleCommitNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/commits/{sha}][%d] repoGetSingleCommitNotFound", 404)
}

func (o *RepoGetSingleCommitNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/commits/{sha}][%d] repoGetSingleCommitNotFound", 404)
}

func (o *RepoGetSingleCommitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoGetSingleCommitUnprocessableEntity creates a RepoGetSingleCommitUnprocessableEntity with default headers values
func NewRepoGetSingleCommitUnprocessableEntity() *RepoGetSingleCommitUnprocessableEntity {
	return &RepoGetSingleCommitUnprocessableEntity{}
}

/*
RepoGetSingleCommitUnprocessableEntity describes a response with status code 422, with default header values.

APIValidationError is error format response related to input validation
*/
type RepoGetSingleCommitUnprocessableEntity struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo get single commit unprocessable entity response has a 2xx status code
func (o *RepoGetSingleCommitUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo get single commit unprocessable entity response has a 3xx status code
func (o *RepoGetSingleCommitUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get single commit unprocessable entity response has a 4xx status code
func (o *RepoGetSingleCommitUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo get single commit unprocessable entity response has a 5xx status code
func (o *RepoGetSingleCommitUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get single commit unprocessable entity response a status code equal to that given
func (o *RepoGetSingleCommitUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the repo get single commit unprocessable entity response
func (o *RepoGetSingleCommitUnprocessableEntity) Code() int {
	return 422
}

func (o *RepoGetSingleCommitUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/commits/{sha}][%d] repoGetSingleCommitUnprocessableEntity", 422)
}

func (o *RepoGetSingleCommitUnprocessableEntity) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/commits/{sha}][%d] repoGetSingleCommitUnprocessableEntity", 422)
}

func (o *RepoGetSingleCommitUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
