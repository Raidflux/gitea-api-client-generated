// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RepoTestHookReader is a Reader for the RepoTestHook structure.
type RepoTestHookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoTestHookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRepoTestHookNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoTestHookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /repos/{owner}/{repo}/hooks/{id}/tests] repoTestHook", response, response.Code())
	}
}

// NewRepoTestHookNoContent creates a RepoTestHookNoContent with default headers values
func NewRepoTestHookNoContent() *RepoTestHookNoContent {
	return &RepoTestHookNoContent{}
}

/*
RepoTestHookNoContent describes a response with status code 204, with default header values.

APIEmpty is an empty response
*/
type RepoTestHookNoContent struct {
}

// IsSuccess returns true when this repo test hook no content response has a 2xx status code
func (o *RepoTestHookNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo test hook no content response has a 3xx status code
func (o *RepoTestHookNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo test hook no content response has a 4xx status code
func (o *RepoTestHookNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo test hook no content response has a 5xx status code
func (o *RepoTestHookNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this repo test hook no content response a status code equal to that given
func (o *RepoTestHookNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the repo test hook no content response
func (o *RepoTestHookNoContent) Code() int {
	return 204
}

func (o *RepoTestHookNoContent) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/hooks/{id}/tests][%d] repoTestHookNoContent", 204)
}

func (o *RepoTestHookNoContent) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/hooks/{id}/tests][%d] repoTestHookNoContent", 204)
}

func (o *RepoTestHookNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoTestHookNotFound creates a RepoTestHookNotFound with default headers values
func NewRepoTestHookNotFound() *RepoTestHookNotFound {
	return &RepoTestHookNotFound{}
}

/*
RepoTestHookNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoTestHookNotFound struct {
}

// IsSuccess returns true when this repo test hook not found response has a 2xx status code
func (o *RepoTestHookNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo test hook not found response has a 3xx status code
func (o *RepoTestHookNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo test hook not found response has a 4xx status code
func (o *RepoTestHookNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo test hook not found response has a 5xx status code
func (o *RepoTestHookNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo test hook not found response a status code equal to that given
func (o *RepoTestHookNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo test hook not found response
func (o *RepoTestHookNotFound) Code() int {
	return 404
}

func (o *RepoTestHookNotFound) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/hooks/{id}/tests][%d] repoTestHookNotFound", 404)
}

func (o *RepoTestHookNotFound) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/hooks/{id}/tests][%d] repoTestHookNotFound", 404)
}

func (o *RepoTestHookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
