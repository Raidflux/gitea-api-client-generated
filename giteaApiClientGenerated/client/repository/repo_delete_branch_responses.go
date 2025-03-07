// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RepoDeleteBranchReader is a Reader for the RepoDeleteBranch structure.
type RepoDeleteBranchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoDeleteBranchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRepoDeleteBranchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRepoDeleteBranchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRepoDeleteBranchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 423:
		result := NewRepoDeleteBranchLocked()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /repos/{owner}/{repo}/branches/{branch}] repoDeleteBranch", response, response.Code())
	}
}

// NewRepoDeleteBranchNoContent creates a RepoDeleteBranchNoContent with default headers values
func NewRepoDeleteBranchNoContent() *RepoDeleteBranchNoContent {
	return &RepoDeleteBranchNoContent{}
}

/*
RepoDeleteBranchNoContent describes a response with status code 204, with default header values.

APIEmpty is an empty response
*/
type RepoDeleteBranchNoContent struct {
}

// IsSuccess returns true when this repo delete branch no content response has a 2xx status code
func (o *RepoDeleteBranchNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo delete branch no content response has a 3xx status code
func (o *RepoDeleteBranchNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo delete branch no content response has a 4xx status code
func (o *RepoDeleteBranchNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo delete branch no content response has a 5xx status code
func (o *RepoDeleteBranchNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this repo delete branch no content response a status code equal to that given
func (o *RepoDeleteBranchNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the repo delete branch no content response
func (o *RepoDeleteBranchNoContent) Code() int {
	return 204
}

func (o *RepoDeleteBranchNoContent) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/branches/{branch}][%d] repoDeleteBranchNoContent", 204)
}

func (o *RepoDeleteBranchNoContent) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/branches/{branch}][%d] repoDeleteBranchNoContent", 204)
}

func (o *RepoDeleteBranchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoDeleteBranchForbidden creates a RepoDeleteBranchForbidden with default headers values
func NewRepoDeleteBranchForbidden() *RepoDeleteBranchForbidden {
	return &RepoDeleteBranchForbidden{}
}

/*
RepoDeleteBranchForbidden describes a response with status code 403, with default header values.

APIError is error format response
*/
type RepoDeleteBranchForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo delete branch forbidden response has a 2xx status code
func (o *RepoDeleteBranchForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo delete branch forbidden response has a 3xx status code
func (o *RepoDeleteBranchForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo delete branch forbidden response has a 4xx status code
func (o *RepoDeleteBranchForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo delete branch forbidden response has a 5xx status code
func (o *RepoDeleteBranchForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this repo delete branch forbidden response a status code equal to that given
func (o *RepoDeleteBranchForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the repo delete branch forbidden response
func (o *RepoDeleteBranchForbidden) Code() int {
	return 403
}

func (o *RepoDeleteBranchForbidden) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/branches/{branch}][%d] repoDeleteBranchForbidden", 403)
}

func (o *RepoDeleteBranchForbidden) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/branches/{branch}][%d] repoDeleteBranchForbidden", 403)
}

func (o *RepoDeleteBranchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRepoDeleteBranchNotFound creates a RepoDeleteBranchNotFound with default headers values
func NewRepoDeleteBranchNotFound() *RepoDeleteBranchNotFound {
	return &RepoDeleteBranchNotFound{}
}

/*
RepoDeleteBranchNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoDeleteBranchNotFound struct {
}

// IsSuccess returns true when this repo delete branch not found response has a 2xx status code
func (o *RepoDeleteBranchNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo delete branch not found response has a 3xx status code
func (o *RepoDeleteBranchNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo delete branch not found response has a 4xx status code
func (o *RepoDeleteBranchNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo delete branch not found response has a 5xx status code
func (o *RepoDeleteBranchNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo delete branch not found response a status code equal to that given
func (o *RepoDeleteBranchNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo delete branch not found response
func (o *RepoDeleteBranchNotFound) Code() int {
	return 404
}

func (o *RepoDeleteBranchNotFound) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/branches/{branch}][%d] repoDeleteBranchNotFound", 404)
}

func (o *RepoDeleteBranchNotFound) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/branches/{branch}][%d] repoDeleteBranchNotFound", 404)
}

func (o *RepoDeleteBranchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoDeleteBranchLocked creates a RepoDeleteBranchLocked with default headers values
func NewRepoDeleteBranchLocked() *RepoDeleteBranchLocked {
	return &RepoDeleteBranchLocked{}
}

/*
RepoDeleteBranchLocked describes a response with status code 423, with default header values.

APIRepoArchivedError is an error that is raised when an archived repo should be modified
*/
type RepoDeleteBranchLocked struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo delete branch locked response has a 2xx status code
func (o *RepoDeleteBranchLocked) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo delete branch locked response has a 3xx status code
func (o *RepoDeleteBranchLocked) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo delete branch locked response has a 4xx status code
func (o *RepoDeleteBranchLocked) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo delete branch locked response has a 5xx status code
func (o *RepoDeleteBranchLocked) IsServerError() bool {
	return false
}

// IsCode returns true when this repo delete branch locked response a status code equal to that given
func (o *RepoDeleteBranchLocked) IsCode(code int) bool {
	return code == 423
}

// Code gets the status code for the repo delete branch locked response
func (o *RepoDeleteBranchLocked) Code() int {
	return 423
}

func (o *RepoDeleteBranchLocked) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/branches/{branch}][%d] repoDeleteBranchLocked", 423)
}

func (o *RepoDeleteBranchLocked) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/branches/{branch}][%d] repoDeleteBranchLocked", 423)
}

func (o *RepoDeleteBranchLocked) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
