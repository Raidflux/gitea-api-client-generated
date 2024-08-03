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

// RepoGetContentsReader is a Reader for the RepoGetContents structure.
type RepoGetContentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoGetContentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoGetContentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoGetContentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/contents/{filepath}] repoGetContents", response, response.Code())
	}
}

// NewRepoGetContentsOK creates a RepoGetContentsOK with default headers values
func NewRepoGetContentsOK() *RepoGetContentsOK {
	return &RepoGetContentsOK{}
}

/*
RepoGetContentsOK describes a response with status code 200, with default header values.

ContentsResponse
*/
type RepoGetContentsOK struct {
	Payload *models.ContentsResponse
}

// IsSuccess returns true when this repo get contents o k response has a 2xx status code
func (o *RepoGetContentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo get contents o k response has a 3xx status code
func (o *RepoGetContentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get contents o k response has a 4xx status code
func (o *RepoGetContentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo get contents o k response has a 5xx status code
func (o *RepoGetContentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get contents o k response a status code equal to that given
func (o *RepoGetContentsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo get contents o k response
func (o *RepoGetContentsOK) Code() int {
	return 200
}

func (o *RepoGetContentsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/contents/{filepath}][%d] repoGetContentsOK %s", 200, payload)
}

func (o *RepoGetContentsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/contents/{filepath}][%d] repoGetContentsOK %s", 200, payload)
}

func (o *RepoGetContentsOK) GetPayload() *models.ContentsResponse {
	return o.Payload
}

func (o *RepoGetContentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContentsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoGetContentsNotFound creates a RepoGetContentsNotFound with default headers values
func NewRepoGetContentsNotFound() *RepoGetContentsNotFound {
	return &RepoGetContentsNotFound{}
}

/*
RepoGetContentsNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoGetContentsNotFound struct {
}

// IsSuccess returns true when this repo get contents not found response has a 2xx status code
func (o *RepoGetContentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo get contents not found response has a 3xx status code
func (o *RepoGetContentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get contents not found response has a 4xx status code
func (o *RepoGetContentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo get contents not found response has a 5xx status code
func (o *RepoGetContentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get contents not found response a status code equal to that given
func (o *RepoGetContentsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo get contents not found response
func (o *RepoGetContentsNotFound) Code() int {
	return 404
}

func (o *RepoGetContentsNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/contents/{filepath}][%d] repoGetContentsNotFound", 404)
}

func (o *RepoGetContentsNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/contents/{filepath}][%d] repoGetContentsNotFound", 404)
}

func (o *RepoGetContentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
