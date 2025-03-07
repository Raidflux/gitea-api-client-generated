// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/Raidflux/gitea-api-client-generated/giteaApiClientGenerated/models"
)

// RepoGetAllCommitsReader is a Reader for the RepoGetAllCommits structure.
type RepoGetAllCommitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoGetAllCommitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoGetAllCommitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoGetAllCommitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRepoGetAllCommitsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/commits] repoGetAllCommits", response, response.Code())
	}
}

// NewRepoGetAllCommitsOK creates a RepoGetAllCommitsOK with default headers values
func NewRepoGetAllCommitsOK() *RepoGetAllCommitsOK {
	return &RepoGetAllCommitsOK{}
}

/*
RepoGetAllCommitsOK describes a response with status code 200, with default header values.

CommitList
*/
type RepoGetAllCommitsOK struct {

	/* True if there is another page
	 */
	XHasMore bool

	/* The current page

	   Format: int64
	*/
	XPage int64

	/* Total number of pages

	   Format: int64
	*/
	XPageCount int64

	/* Commits per page

	   Format: int64
	*/
	XPerPage int64

	/* Total commit count

	   Format: int64
	*/
	XTotal int64

	Payload []*models.Commit
}

// IsSuccess returns true when this repo get all commits o k response has a 2xx status code
func (o *RepoGetAllCommitsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo get all commits o k response has a 3xx status code
func (o *RepoGetAllCommitsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get all commits o k response has a 4xx status code
func (o *RepoGetAllCommitsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo get all commits o k response has a 5xx status code
func (o *RepoGetAllCommitsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get all commits o k response a status code equal to that given
func (o *RepoGetAllCommitsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo get all commits o k response
func (o *RepoGetAllCommitsOK) Code() int {
	return 200
}

func (o *RepoGetAllCommitsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/commits][%d] repoGetAllCommitsOK %s", 200, payload)
}

func (o *RepoGetAllCommitsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/commits][%d] repoGetAllCommitsOK %s", 200, payload)
}

func (o *RepoGetAllCommitsOK) GetPayload() []*models.Commit {
	return o.Payload
}

func (o *RepoGetAllCommitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-HasMore
	hdrXHasMore := response.GetHeader("X-HasMore")

	if hdrXHasMore != "" {
		valxHasMore, err := swag.ConvertBool(hdrXHasMore)
		if err != nil {
			return errors.InvalidType("X-HasMore", "header", "bool", hdrXHasMore)
		}
		o.XHasMore = valxHasMore
	}

	// hydrates response header X-Page
	hdrXPage := response.GetHeader("X-Page")

	if hdrXPage != "" {
		valxPage, err := swag.ConvertInt64(hdrXPage)
		if err != nil {
			return errors.InvalidType("X-Page", "header", "int64", hdrXPage)
		}
		o.XPage = valxPage
	}

	// hydrates response header X-PageCount
	hdrXPageCount := response.GetHeader("X-PageCount")

	if hdrXPageCount != "" {
		valxPageCount, err := swag.ConvertInt64(hdrXPageCount)
		if err != nil {
			return errors.InvalidType("X-PageCount", "header", "int64", hdrXPageCount)
		}
		o.XPageCount = valxPageCount
	}

	// hydrates response header X-PerPage
	hdrXPerPage := response.GetHeader("X-PerPage")

	if hdrXPerPage != "" {
		valxPerPage, err := swag.ConvertInt64(hdrXPerPage)
		if err != nil {
			return errors.InvalidType("X-PerPage", "header", "int64", hdrXPerPage)
		}
		o.XPerPage = valxPerPage
	}

	// hydrates response header X-Total
	hdrXTotal := response.GetHeader("X-Total")

	if hdrXTotal != "" {
		valxTotal, err := swag.ConvertInt64(hdrXTotal)
		if err != nil {
			return errors.InvalidType("X-Total", "header", "int64", hdrXTotal)
		}
		o.XTotal = valxTotal
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoGetAllCommitsNotFound creates a RepoGetAllCommitsNotFound with default headers values
func NewRepoGetAllCommitsNotFound() *RepoGetAllCommitsNotFound {
	return &RepoGetAllCommitsNotFound{}
}

/*
RepoGetAllCommitsNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoGetAllCommitsNotFound struct {
}

// IsSuccess returns true when this repo get all commits not found response has a 2xx status code
func (o *RepoGetAllCommitsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo get all commits not found response has a 3xx status code
func (o *RepoGetAllCommitsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get all commits not found response has a 4xx status code
func (o *RepoGetAllCommitsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo get all commits not found response has a 5xx status code
func (o *RepoGetAllCommitsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get all commits not found response a status code equal to that given
func (o *RepoGetAllCommitsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo get all commits not found response
func (o *RepoGetAllCommitsNotFound) Code() int {
	return 404
}

func (o *RepoGetAllCommitsNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/commits][%d] repoGetAllCommitsNotFound", 404)
}

func (o *RepoGetAllCommitsNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/commits][%d] repoGetAllCommitsNotFound", 404)
}

func (o *RepoGetAllCommitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoGetAllCommitsConflict creates a RepoGetAllCommitsConflict with default headers values
func NewRepoGetAllCommitsConflict() *RepoGetAllCommitsConflict {
	return &RepoGetAllCommitsConflict{}
}

/*
RepoGetAllCommitsConflict describes a response with status code 409, with default header values.

EmptyRepository
*/
type RepoGetAllCommitsConflict struct {
	Payload *models.APIError
}

// IsSuccess returns true when this repo get all commits conflict response has a 2xx status code
func (o *RepoGetAllCommitsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo get all commits conflict response has a 3xx status code
func (o *RepoGetAllCommitsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get all commits conflict response has a 4xx status code
func (o *RepoGetAllCommitsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo get all commits conflict response has a 5xx status code
func (o *RepoGetAllCommitsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get all commits conflict response a status code equal to that given
func (o *RepoGetAllCommitsConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the repo get all commits conflict response
func (o *RepoGetAllCommitsConflict) Code() int {
	return 409
}

func (o *RepoGetAllCommitsConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/commits][%d] repoGetAllCommitsConflict %s", 409, payload)
}

func (o *RepoGetAllCommitsConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/commits][%d] repoGetAllCommitsConflict %s", 409, payload)
}

func (o *RepoGetAllCommitsConflict) GetPayload() *models.APIError {
	return o.Payload
}

func (o *RepoGetAllCommitsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
