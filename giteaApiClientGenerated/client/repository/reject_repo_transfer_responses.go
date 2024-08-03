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

// RejectRepoTransferReader is a Reader for the RejectRepoTransfer structure.
type RejectRepoTransferReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RejectRepoTransferReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRejectRepoTransferOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewRejectRepoTransferForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRejectRepoTransferNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /repos/{owner}/{repo}/transfer/reject] rejectRepoTransfer", response, response.Code())
	}
}

// NewRejectRepoTransferOK creates a RejectRepoTransferOK with default headers values
func NewRejectRepoTransferOK() *RejectRepoTransferOK {
	return &RejectRepoTransferOK{}
}

/*
RejectRepoTransferOK describes a response with status code 200, with default header values.

Repository
*/
type RejectRepoTransferOK struct {
	Payload *models.Repository
}

// IsSuccess returns true when this reject repo transfer o k response has a 2xx status code
func (o *RejectRepoTransferOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this reject repo transfer o k response has a 3xx status code
func (o *RejectRepoTransferOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reject repo transfer o k response has a 4xx status code
func (o *RejectRepoTransferOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this reject repo transfer o k response has a 5xx status code
func (o *RejectRepoTransferOK) IsServerError() bool {
	return false
}

// IsCode returns true when this reject repo transfer o k response a status code equal to that given
func (o *RejectRepoTransferOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the reject repo transfer o k response
func (o *RejectRepoTransferOK) Code() int {
	return 200
}

func (o *RejectRepoTransferOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/transfer/reject][%d] rejectRepoTransferOK %s", 200, payload)
}

func (o *RejectRepoTransferOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/transfer/reject][%d] rejectRepoTransferOK %s", 200, payload)
}

func (o *RejectRepoTransferOK) GetPayload() *models.Repository {
	return o.Payload
}

func (o *RejectRepoTransferOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Repository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRejectRepoTransferForbidden creates a RejectRepoTransferForbidden with default headers values
func NewRejectRepoTransferForbidden() *RejectRepoTransferForbidden {
	return &RejectRepoTransferForbidden{}
}

/*
RejectRepoTransferForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type RejectRepoTransferForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this reject repo transfer forbidden response has a 2xx status code
func (o *RejectRepoTransferForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this reject repo transfer forbidden response has a 3xx status code
func (o *RejectRepoTransferForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reject repo transfer forbidden response has a 4xx status code
func (o *RejectRepoTransferForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this reject repo transfer forbidden response has a 5xx status code
func (o *RejectRepoTransferForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this reject repo transfer forbidden response a status code equal to that given
func (o *RejectRepoTransferForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the reject repo transfer forbidden response
func (o *RejectRepoTransferForbidden) Code() int {
	return 403
}

func (o *RejectRepoTransferForbidden) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/transfer/reject][%d] rejectRepoTransferForbidden", 403)
}

func (o *RejectRepoTransferForbidden) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/transfer/reject][%d] rejectRepoTransferForbidden", 403)
}

func (o *RejectRepoTransferForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRejectRepoTransferNotFound creates a RejectRepoTransferNotFound with default headers values
func NewRejectRepoTransferNotFound() *RejectRepoTransferNotFound {
	return &RejectRepoTransferNotFound{}
}

/*
RejectRepoTransferNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RejectRepoTransferNotFound struct {
}

// IsSuccess returns true when this reject repo transfer not found response has a 2xx status code
func (o *RejectRepoTransferNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this reject repo transfer not found response has a 3xx status code
func (o *RejectRepoTransferNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reject repo transfer not found response has a 4xx status code
func (o *RejectRepoTransferNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this reject repo transfer not found response has a 5xx status code
func (o *RejectRepoTransferNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this reject repo transfer not found response a status code equal to that given
func (o *RejectRepoTransferNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the reject repo transfer not found response
func (o *RejectRepoTransferNotFound) Code() int {
	return 404
}

func (o *RejectRepoTransferNotFound) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/transfer/reject][%d] rejectRepoTransferNotFound", 404)
}

func (o *RejectRepoTransferNotFound) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/transfer/reject][%d] rejectRepoTransferNotFound", 404)
}

func (o *RejectRepoTransferNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
