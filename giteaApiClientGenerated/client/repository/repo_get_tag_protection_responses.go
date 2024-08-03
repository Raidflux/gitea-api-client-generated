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

// RepoGetTagProtectionReader is a Reader for the RepoGetTagProtection structure.
type RepoGetTagProtectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoGetTagProtectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoGetTagProtectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRepoGetTagProtectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/tag_protections/{id}] repoGetTagProtection", response, response.Code())
	}
}

// NewRepoGetTagProtectionOK creates a RepoGetTagProtectionOK with default headers values
func NewRepoGetTagProtectionOK() *RepoGetTagProtectionOK {
	return &RepoGetTagProtectionOK{}
}

/*
RepoGetTagProtectionOK describes a response with status code 200, with default header values.

TagProtection
*/
type RepoGetTagProtectionOK struct {
	Payload *models.TagProtection
}

// IsSuccess returns true when this repo get tag protection o k response has a 2xx status code
func (o *RepoGetTagProtectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo get tag protection o k response has a 3xx status code
func (o *RepoGetTagProtectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get tag protection o k response has a 4xx status code
func (o *RepoGetTagProtectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo get tag protection o k response has a 5xx status code
func (o *RepoGetTagProtectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get tag protection o k response a status code equal to that given
func (o *RepoGetTagProtectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo get tag protection o k response
func (o *RepoGetTagProtectionOK) Code() int {
	return 200
}

func (o *RepoGetTagProtectionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/tag_protections/{id}][%d] repoGetTagProtectionOK %s", 200, payload)
}

func (o *RepoGetTagProtectionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/tag_protections/{id}][%d] repoGetTagProtectionOK %s", 200, payload)
}

func (o *RepoGetTagProtectionOK) GetPayload() *models.TagProtection {
	return o.Payload
}

func (o *RepoGetTagProtectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagProtection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepoGetTagProtectionNotFound creates a RepoGetTagProtectionNotFound with default headers values
func NewRepoGetTagProtectionNotFound() *RepoGetTagProtectionNotFound {
	return &RepoGetTagProtectionNotFound{}
}

/*
RepoGetTagProtectionNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoGetTagProtectionNotFound struct {
}

// IsSuccess returns true when this repo get tag protection not found response has a 2xx status code
func (o *RepoGetTagProtectionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo get tag protection not found response has a 3xx status code
func (o *RepoGetTagProtectionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo get tag protection not found response has a 4xx status code
func (o *RepoGetTagProtectionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo get tag protection not found response has a 5xx status code
func (o *RepoGetTagProtectionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo get tag protection not found response a status code equal to that given
func (o *RepoGetTagProtectionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo get tag protection not found response
func (o *RepoGetTagProtectionNotFound) Code() int {
	return 404
}

func (o *RepoGetTagProtectionNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/tag_protections/{id}][%d] repoGetTagProtectionNotFound", 404)
}

func (o *RepoGetTagProtectionNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/tag_protections/{id}][%d] repoGetTagProtectionNotFound", 404)
}

func (o *RepoGetTagProtectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
