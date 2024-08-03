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

// ListForksReader is a Reader for the ListForks structure.
type ListForksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListForksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListForksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewListForksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/forks] listForks", response, response.Code())
	}
}

// NewListForksOK creates a ListForksOK with default headers values
func NewListForksOK() *ListForksOK {
	return &ListForksOK{}
}

/*
ListForksOK describes a response with status code 200, with default header values.

RepositoryList
*/
type ListForksOK struct {
	Payload []*models.Repository
}

// IsSuccess returns true when this list forks o k response has a 2xx status code
func (o *ListForksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list forks o k response has a 3xx status code
func (o *ListForksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list forks o k response has a 4xx status code
func (o *ListForksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list forks o k response has a 5xx status code
func (o *ListForksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list forks o k response a status code equal to that given
func (o *ListForksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list forks o k response
func (o *ListForksOK) Code() int {
	return 200
}

func (o *ListForksOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/forks][%d] listForksOK %s", 200, payload)
}

func (o *ListForksOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/forks][%d] listForksOK %s", 200, payload)
}

func (o *ListForksOK) GetPayload() []*models.Repository {
	return o.Payload
}

func (o *ListForksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListForksNotFound creates a ListForksNotFound with default headers values
func NewListForksNotFound() *ListForksNotFound {
	return &ListForksNotFound{}
}

/*
ListForksNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type ListForksNotFound struct {
}

// IsSuccess returns true when this list forks not found response has a 2xx status code
func (o *ListForksNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list forks not found response has a 3xx status code
func (o *ListForksNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list forks not found response has a 4xx status code
func (o *ListForksNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list forks not found response has a 5xx status code
func (o *ListForksNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list forks not found response a status code equal to that given
func (o *ListForksNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list forks not found response
func (o *ListForksNotFound) Code() int {
	return 404
}

func (o *ListForksNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/forks][%d] listForksNotFound", 404)
}

func (o *ListForksNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/forks][%d] listForksNotFound", 404)
}

func (o *ListForksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
