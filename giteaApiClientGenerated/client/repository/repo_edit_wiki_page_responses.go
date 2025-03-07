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

// RepoEditWikiPageReader is a Reader for the RepoEditWikiPage structure.
type RepoEditWikiPageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoEditWikiPageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoEditWikiPageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRepoEditWikiPageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRepoEditWikiPageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRepoEditWikiPageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 423:
		result := NewRepoEditWikiPageLocked()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /repos/{owner}/{repo}/wiki/page/{pageName}] repoEditWikiPage", response, response.Code())
	}
}

// NewRepoEditWikiPageOK creates a RepoEditWikiPageOK with default headers values
func NewRepoEditWikiPageOK() *RepoEditWikiPageOK {
	return &RepoEditWikiPageOK{}
}

/*
RepoEditWikiPageOK describes a response with status code 200, with default header values.

WikiPage
*/
type RepoEditWikiPageOK struct {
	Payload *models.WikiPage
}

// IsSuccess returns true when this repo edit wiki page o k response has a 2xx status code
func (o *RepoEditWikiPageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo edit wiki page o k response has a 3xx status code
func (o *RepoEditWikiPageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo edit wiki page o k response has a 4xx status code
func (o *RepoEditWikiPageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo edit wiki page o k response has a 5xx status code
func (o *RepoEditWikiPageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo edit wiki page o k response a status code equal to that given
func (o *RepoEditWikiPageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo edit wiki page o k response
func (o *RepoEditWikiPageOK) Code() int {
	return 200
}

func (o *RepoEditWikiPageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoEditWikiPageOK %s", 200, payload)
}

func (o *RepoEditWikiPageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoEditWikiPageOK %s", 200, payload)
}

func (o *RepoEditWikiPageOK) GetPayload() *models.WikiPage {
	return o.Payload
}

func (o *RepoEditWikiPageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WikiPage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoEditWikiPageBadRequest creates a RepoEditWikiPageBadRequest with default headers values
func NewRepoEditWikiPageBadRequest() *RepoEditWikiPageBadRequest {
	return &RepoEditWikiPageBadRequest{}
}

/*
RepoEditWikiPageBadRequest describes a response with status code 400, with default header values.

APIError is error format response
*/
type RepoEditWikiPageBadRequest struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo edit wiki page bad request response has a 2xx status code
func (o *RepoEditWikiPageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo edit wiki page bad request response has a 3xx status code
func (o *RepoEditWikiPageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo edit wiki page bad request response has a 4xx status code
func (o *RepoEditWikiPageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo edit wiki page bad request response has a 5xx status code
func (o *RepoEditWikiPageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this repo edit wiki page bad request response a status code equal to that given
func (o *RepoEditWikiPageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the repo edit wiki page bad request response
func (o *RepoEditWikiPageBadRequest) Code() int {
	return 400
}

func (o *RepoEditWikiPageBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoEditWikiPageBadRequest", 400)
}

func (o *RepoEditWikiPageBadRequest) String() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoEditWikiPageBadRequest", 400)
}

func (o *RepoEditWikiPageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRepoEditWikiPageForbidden creates a RepoEditWikiPageForbidden with default headers values
func NewRepoEditWikiPageForbidden() *RepoEditWikiPageForbidden {
	return &RepoEditWikiPageForbidden{}
}

/*
RepoEditWikiPageForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type RepoEditWikiPageForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo edit wiki page forbidden response has a 2xx status code
func (o *RepoEditWikiPageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo edit wiki page forbidden response has a 3xx status code
func (o *RepoEditWikiPageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo edit wiki page forbidden response has a 4xx status code
func (o *RepoEditWikiPageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo edit wiki page forbidden response has a 5xx status code
func (o *RepoEditWikiPageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this repo edit wiki page forbidden response a status code equal to that given
func (o *RepoEditWikiPageForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the repo edit wiki page forbidden response
func (o *RepoEditWikiPageForbidden) Code() int {
	return 403
}

func (o *RepoEditWikiPageForbidden) Error() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoEditWikiPageForbidden", 403)
}

func (o *RepoEditWikiPageForbidden) String() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoEditWikiPageForbidden", 403)
}

func (o *RepoEditWikiPageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRepoEditWikiPageNotFound creates a RepoEditWikiPageNotFound with default headers values
func NewRepoEditWikiPageNotFound() *RepoEditWikiPageNotFound {
	return &RepoEditWikiPageNotFound{}
}

/*
RepoEditWikiPageNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoEditWikiPageNotFound struct {
}

// IsSuccess returns true when this repo edit wiki page not found response has a 2xx status code
func (o *RepoEditWikiPageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo edit wiki page not found response has a 3xx status code
func (o *RepoEditWikiPageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo edit wiki page not found response has a 4xx status code
func (o *RepoEditWikiPageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo edit wiki page not found response has a 5xx status code
func (o *RepoEditWikiPageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo edit wiki page not found response a status code equal to that given
func (o *RepoEditWikiPageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo edit wiki page not found response
func (o *RepoEditWikiPageNotFound) Code() int {
	return 404
}

func (o *RepoEditWikiPageNotFound) Error() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoEditWikiPageNotFound", 404)
}

func (o *RepoEditWikiPageNotFound) String() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoEditWikiPageNotFound", 404)
}

func (o *RepoEditWikiPageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoEditWikiPageLocked creates a RepoEditWikiPageLocked with default headers values
func NewRepoEditWikiPageLocked() *RepoEditWikiPageLocked {
	return &RepoEditWikiPageLocked{}
}

/*
RepoEditWikiPageLocked describes a response with status code 423, with default header values.

APIRepoArchivedError is an error that is raised when an archived repo should be modified
*/
type RepoEditWikiPageLocked struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo edit wiki page locked response has a 2xx status code
func (o *RepoEditWikiPageLocked) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo edit wiki page locked response has a 3xx status code
func (o *RepoEditWikiPageLocked) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo edit wiki page locked response has a 4xx status code
func (o *RepoEditWikiPageLocked) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo edit wiki page locked response has a 5xx status code
func (o *RepoEditWikiPageLocked) IsServerError() bool {
	return false
}

// IsCode returns true when this repo edit wiki page locked response a status code equal to that given
func (o *RepoEditWikiPageLocked) IsCode(code int) bool {
	return code == 423
}

// Code gets the status code for the repo edit wiki page locked response
func (o *RepoEditWikiPageLocked) Code() int {
	return 423
}

func (o *RepoEditWikiPageLocked) Error() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoEditWikiPageLocked", 423)
}

func (o *RepoEditWikiPageLocked) String() string {
	return fmt.Sprintf("[PATCH /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoEditWikiPageLocked", 423)
}

func (o *RepoEditWikiPageLocked) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
