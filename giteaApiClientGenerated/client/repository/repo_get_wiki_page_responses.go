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

	"giteaApiClientGenerated/models"
)

// RepoGetWikiPageReader is a Reader for the RepoGetWikiPage structure.
type RepoGetWikiPageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoGetWikiPageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoGetWikiPageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoGetWikiPageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/wiki/page/{pageName}] repoGetWikiPage", response, response.Code())
	}
}

// NewRepoGetWikiPageOK creates a RepoGetWikiPageOK with default headers values
func NewRepoGetWikiPageOK() *RepoGetWikiPageOK {
	return &RepoGetWikiPageOK{}
}

/*
RepoGetWikiPageOK describes a response with status code 200, with default header values.

WikiPage
*/
type RepoGetWikiPageOK struct {
	Payload *models.WikiPage
}

// IsSuccess returns true when this repo get wiki page o k response has a 2xx status code
func (o *RepoGetWikiPageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo get wiki page o k response has a 3xx status code
func (o *RepoGetWikiPageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get wiki page o k response has a 4xx status code
func (o *RepoGetWikiPageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo get wiki page o k response has a 5xx status code
func (o *RepoGetWikiPageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get wiki page o k response a status code equal to that given
func (o *RepoGetWikiPageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo get wiki page o k response
func (o *RepoGetWikiPageOK) Code() int {
	return 200
}

func (o *RepoGetWikiPageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoGetWikiPageOK %s", 200, payload)
}

func (o *RepoGetWikiPageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoGetWikiPageOK %s", 200, payload)
}

func (o *RepoGetWikiPageOK) GetPayload() *models.WikiPage {
	return o.Payload
}

func (o *RepoGetWikiPageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WikiPage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoGetWikiPageNotFound creates a RepoGetWikiPageNotFound with default headers values
func NewRepoGetWikiPageNotFound() *RepoGetWikiPageNotFound {
	return &RepoGetWikiPageNotFound{}
}

/*
RepoGetWikiPageNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoGetWikiPageNotFound struct {
}

// IsSuccess returns true when this repo get wiki page not found response has a 2xx status code
func (o *RepoGetWikiPageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo get wiki page not found response has a 3xx status code
func (o *RepoGetWikiPageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get wiki page not found response has a 4xx status code
func (o *RepoGetWikiPageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo get wiki page not found response has a 5xx status code
func (o *RepoGetWikiPageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get wiki page not found response a status code equal to that given
func (o *RepoGetWikiPageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo get wiki page not found response
func (o *RepoGetWikiPageNotFound) Code() int {
	return 404
}

func (o *RepoGetWikiPageNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoGetWikiPageNotFound", 404)
}

func (o *RepoGetWikiPageNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/wiki/page/{pageName}][%d] repoGetWikiPageNotFound", 404)
}

func (o *RepoGetWikiPageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
