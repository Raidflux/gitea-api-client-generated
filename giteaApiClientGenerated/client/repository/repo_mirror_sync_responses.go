// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RepoMirrorSyncReader is a Reader for the RepoMirrorSync structure.
type RepoMirrorSyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoMirrorSyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoMirrorSyncOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRepoMirrorSyncForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRepoMirrorSyncNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /repos/{owner}/{repo}/mirror-sync] repoMirrorSync", response, response.Code())
	}
}

// NewRepoMirrorSyncOK creates a RepoMirrorSyncOK with default headers values
func NewRepoMirrorSyncOK() *RepoMirrorSyncOK {
	return &RepoMirrorSyncOK{}
}

/*
RepoMirrorSyncOK describes a response with status code 200, with default header values.

APIEmpty is an empty response
*/
type RepoMirrorSyncOK struct {
}

// IsSuccess returns true when this repo mirror sync o k response has a 2xx status code
func (o *RepoMirrorSyncOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo mirror sync o k response has a 3xx status code
func (o *RepoMirrorSyncOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo mirror sync o k response has a 4xx status code
func (o *RepoMirrorSyncOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo mirror sync o k response has a 5xx status code
func (o *RepoMirrorSyncOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo mirror sync o k response a status code equal to that given
func (o *RepoMirrorSyncOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo mirror sync o k response
func (o *RepoMirrorSyncOK) Code() int {
	return 200
}

func (o *RepoMirrorSyncOK) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/mirror-sync][%d] repoMirrorSyncOK", 200)
}

func (o *RepoMirrorSyncOK) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/mirror-sync][%d] repoMirrorSyncOK", 200)
}

func (o *RepoMirrorSyncOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoMirrorSyncForbidden creates a RepoMirrorSyncForbidden with default headers values
func NewRepoMirrorSyncForbidden() *RepoMirrorSyncForbidden {
	return &RepoMirrorSyncForbidden{}
}

/*
RepoMirrorSyncForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type RepoMirrorSyncForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo mirror sync forbidden response has a 2xx status code
func (o *RepoMirrorSyncForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo mirror sync forbidden response has a 3xx status code
func (o *RepoMirrorSyncForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo mirror sync forbidden response has a 4xx status code
func (o *RepoMirrorSyncForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo mirror sync forbidden response has a 5xx status code
func (o *RepoMirrorSyncForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this repo mirror sync forbidden response a status code equal to that given
func (o *RepoMirrorSyncForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the repo mirror sync forbidden response
func (o *RepoMirrorSyncForbidden) Code() int {
	return 403
}

func (o *RepoMirrorSyncForbidden) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/mirror-sync][%d] repoMirrorSyncForbidden", 403)
}

func (o *RepoMirrorSyncForbidden) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/mirror-sync][%d] repoMirrorSyncForbidden", 403)
}

func (o *RepoMirrorSyncForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRepoMirrorSyncNotFound creates a RepoMirrorSyncNotFound with default headers values
func NewRepoMirrorSyncNotFound() *RepoMirrorSyncNotFound {
	return &RepoMirrorSyncNotFound{}
}

/*
RepoMirrorSyncNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoMirrorSyncNotFound struct {
}

// IsSuccess returns true when this repo mirror sync not found response has a 2xx status code
func (o *RepoMirrorSyncNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo mirror sync not found response has a 3xx status code
func (o *RepoMirrorSyncNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo mirror sync not found response has a 4xx status code
func (o *RepoMirrorSyncNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo mirror sync not found response has a 5xx status code
func (o *RepoMirrorSyncNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo mirror sync not found response a status code equal to that given
func (o *RepoMirrorSyncNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo mirror sync not found response
func (o *RepoMirrorSyncNotFound) Code() int {
	return 404
}

func (o *RepoMirrorSyncNotFound) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/mirror-sync][%d] repoMirrorSyncNotFound", 404)
}

func (o *RepoMirrorSyncNotFound) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/mirror-sync][%d] repoMirrorSyncNotFound", 404)
}

func (o *RepoMirrorSyncNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
