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

// RepoGetBranchReader is a Reader for the RepoGetBranch structure.
type RepoGetBranchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoGetBranchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoGetBranchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoGetBranchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/branches/{branch}] repoGetBranch", response, response.Code())
	}
}

// NewRepoGetBranchOK creates a RepoGetBranchOK with default headers values
func NewRepoGetBranchOK() *RepoGetBranchOK {
	return &RepoGetBranchOK{}
}

/*
RepoGetBranchOK describes a response with status code 200, with default header values.

Branch
*/
type RepoGetBranchOK struct {
	Payload *models.Branch
}

// IsSuccess returns true when this repo get branch o k response has a 2xx status code
func (o *RepoGetBranchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo get branch o k response has a 3xx status code
func (o *RepoGetBranchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get branch o k response has a 4xx status code
func (o *RepoGetBranchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo get branch o k response has a 5xx status code
func (o *RepoGetBranchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get branch o k response a status code equal to that given
func (o *RepoGetBranchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo get branch o k response
func (o *RepoGetBranchOK) Code() int {
	return 200
}

func (o *RepoGetBranchOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/branches/{branch}][%d] repoGetBranchOK %s", 200, payload)
}

func (o *RepoGetBranchOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/branches/{branch}][%d] repoGetBranchOK %s", 200, payload)
}

func (o *RepoGetBranchOK) GetPayload() *models.Branch {
	return o.Payload
}

func (o *RepoGetBranchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Branch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoGetBranchNotFound creates a RepoGetBranchNotFound with default headers values
func NewRepoGetBranchNotFound() *RepoGetBranchNotFound {
	return &RepoGetBranchNotFound{}
}

/*
RepoGetBranchNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoGetBranchNotFound struct {
}

// IsSuccess returns true when this repo get branch not found response has a 2xx status code
func (o *RepoGetBranchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo get branch not found response has a 3xx status code
func (o *RepoGetBranchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get branch not found response has a 4xx status code
func (o *RepoGetBranchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo get branch not found response has a 5xx status code
func (o *RepoGetBranchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get branch not found response a status code equal to that given
func (o *RepoGetBranchNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo get branch not found response
func (o *RepoGetBranchNotFound) Code() int {
	return 404
}

func (o *RepoGetBranchNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/branches/{branch}][%d] repoGetBranchNotFound", 404)
}

func (o *RepoGetBranchNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/branches/{branch}][%d] repoGetBranchNotFound", 404)
}

func (o *RepoGetBranchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
