// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RepoDeleteWikiPageReader is a Reader for the RepoDeleteWikiPage structure.
type RepoDeleteWikiPageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoDeleteWikiPageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRepoDeleteWikiPageNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRepoDeleteWikiPageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRepoDeleteWikiPageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 423:
		result := NewRepoDeleteWikiPageLocked()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /repos/{owner}/{repo}/wiki/page/{pageName}] repoDeleteWikiPage", response, response.Code())
	}
}

// NewRepoDeleteWikiPageNoContent creates a RepoDeleteWikiPageNoContent with default headers values
func NewRepoDeleteWikiPageNoContent() *RepoDeleteWikiPageNoContent {
	return &RepoDeleteWikiPageNoContent{}
}

/*
RepoDeleteWikiPageNoContent describes a response with status code 204, with default header values.

APIEmpty is an empty response
*/
type RepoDeleteWikiPageNoContent struct {
}

// IsSuccess returns true when this repo delete wiki page no content response has a 2xx status code
func (o *RepoDeleteWikiPageNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo delete wiki page no content response has a 3xx status code
func (o *RepoDeleteWikiPageNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo delete wiki page no content response has a 4xx status code
func (o *RepoDeleteWikiPageNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo delete wiki page no content response has a 5xx status code
func (o *RepoDeleteWikiPageNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this repo delete wiki page no content response a status code equal to that given
func (o *RepoDeleteWikiPageNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the repo delete wiki page no content response
func (o *RepoDeleteWikiPageNoContent) Code() int {
	return 204
}

func (o *RepoDeleteWikiPageNoContent) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoDeleteWikiPageNoContent", 204)
}

func (o *RepoDeleteWikiPageNoContent) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoDeleteWikiPageNoContent", 204)
}

func (o *RepoDeleteWikiPageNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoDeleteWikiPageForbidden creates a RepoDeleteWikiPageForbidden with default headers values
func NewRepoDeleteWikiPageForbidden() *RepoDeleteWikiPageForbidden {
	return &RepoDeleteWikiPageForbidden{}
}

/*
RepoDeleteWikiPageForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type RepoDeleteWikiPageForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo delete wiki page forbidden response has a 2xx status code
func (o *RepoDeleteWikiPageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo delete wiki page forbidden response has a 3xx status code
func (o *RepoDeleteWikiPageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo delete wiki page forbidden response has a 4xx status code
func (o *RepoDeleteWikiPageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo delete wiki page forbidden response has a 5xx status code
func (o *RepoDeleteWikiPageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this repo delete wiki page forbidden response a status code equal to that given
func (o *RepoDeleteWikiPageForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the repo delete wiki page forbidden response
func (o *RepoDeleteWikiPageForbidden) Code() int {
	return 403
}

func (o *RepoDeleteWikiPageForbidden) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoDeleteWikiPageForbidden", 403)
}

func (o *RepoDeleteWikiPageForbidden) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoDeleteWikiPageForbidden", 403)
}

func (o *RepoDeleteWikiPageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRepoDeleteWikiPageNotFound creates a RepoDeleteWikiPageNotFound with default headers values
func NewRepoDeleteWikiPageNotFound() *RepoDeleteWikiPageNotFound {
	return &RepoDeleteWikiPageNotFound{}
}

/*
RepoDeleteWikiPageNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoDeleteWikiPageNotFound struct {
}

// IsSuccess returns true when this repo delete wiki page not found response has a 2xx status code
func (o *RepoDeleteWikiPageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo delete wiki page not found response has a 3xx status code
func (o *RepoDeleteWikiPageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo delete wiki page not found response has a 4xx status code
func (o *RepoDeleteWikiPageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo delete wiki page not found response has a 5xx status code
func (o *RepoDeleteWikiPageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo delete wiki page not found response a status code equal to that given
func (o *RepoDeleteWikiPageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo delete wiki page not found response
func (o *RepoDeleteWikiPageNotFound) Code() int {
	return 404
}

func (o *RepoDeleteWikiPageNotFound) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoDeleteWikiPageNotFound", 404)
}

func (o *RepoDeleteWikiPageNotFound) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoDeleteWikiPageNotFound", 404)
}

func (o *RepoDeleteWikiPageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoDeleteWikiPageLocked creates a RepoDeleteWikiPageLocked with default headers values
func NewRepoDeleteWikiPageLocked() *RepoDeleteWikiPageLocked {
	return &RepoDeleteWikiPageLocked{}
}

/*
RepoDeleteWikiPageLocked describes a response with status code 423, with default header values.

APIRepoArchivedError is an error that is raised when an archived repo should be modified
*/
type RepoDeleteWikiPageLocked struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo delete wiki page locked response has a 2xx status code
func (o *RepoDeleteWikiPageLocked) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo delete wiki page locked response has a 3xx status code
func (o *RepoDeleteWikiPageLocked) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo delete wiki page locked response has a 4xx status code
func (o *RepoDeleteWikiPageLocked) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo delete wiki page locked response has a 5xx status code
func (o *RepoDeleteWikiPageLocked) IsServerError() bool {
	return false
}

// IsCode returns true when this repo delete wiki page locked response a status code equal to that given
func (o *RepoDeleteWikiPageLocked) IsCode(code int) bool {
	return code == 423
}

// Code gets the status code for the repo delete wiki page locked response
func (o *RepoDeleteWikiPageLocked) Code() int {
	return 423
}

func (o *RepoDeleteWikiPageLocked) Error() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoDeleteWikiPageLocked", 423)
}

func (o *RepoDeleteWikiPageLocked) String() string {
	return fmt.Sprintf("[DELETE /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoDeleteWikiPageLocked", 423)
}

func (o *RepoDeleteWikiPageLocked) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
