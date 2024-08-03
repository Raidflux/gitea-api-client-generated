// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RepoPushMirrorSyncReader is a Reader for the RepoPushMirrorSync structure.
type RepoPushMirrorSyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepoPushMirrorSyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRepoPushMirrorSyncOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRepoPushMirrorSyncBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRepoPushMirrorSyncForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRepoPushMirrorSyncNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /repos/{owner}/{repo}/push_mirrors-sync] repoPushMirrorSync", response, response.Code())
	}
}

// NewRepoPushMirrorSyncOK creates a RepoPushMirrorSyncOK with default headers values
func NewRepoPushMirrorSyncOK() *RepoPushMirrorSyncOK {
	return &RepoPushMirrorSyncOK{}
}

/*
RepoPushMirrorSyncOK describes a response with status code 200, with default header values.

APIEmpty is an empty response
*/
type RepoPushMirrorSyncOK struct {
}

// IsSuccess returns true when this repo push mirror sync o k response has a 2xx status code
func (o *RepoPushMirrorSyncOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this repo push mirror sync o k response has a 3xx status code
func (o *RepoPushMirrorSyncOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo push mirror sync o k response has a 4xx status code
func (o *RepoPushMirrorSyncOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this repo push mirror sync o k response has a 5xx status code
func (o *RepoPushMirrorSyncOK) IsServerError() bool {
	return false
}

// IsCode returns true when this repo push mirror sync o k response a status code equal to that given
func (o *RepoPushMirrorSyncOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the repo push mirror sync o k response
func (o *RepoPushMirrorSyncOK) Code() int {
	return 200
}

func (o *RepoPushMirrorSyncOK) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/push_mirrors-sync][%d] repoPushMirrorSyncOK", 200)
}

func (o *RepoPushMirrorSyncOK) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/push_mirrors-sync][%d] repoPushMirrorSyncOK", 200)
}

func (o *RepoPushMirrorSyncOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRepoPushMirrorSyncBadRequest creates a RepoPushMirrorSyncBadRequest with default headers values
func NewRepoPushMirrorSyncBadRequest() *RepoPushMirrorSyncBadRequest {
	return &RepoPushMirrorSyncBadRequest{}
}

/*
RepoPushMirrorSyncBadRequest describes a response with status code 400, with default header values.

APIError is error format response
*/
type RepoPushMirrorSyncBadRequest struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo push mirror sync bad request response has a 2xx status code
func (o *RepoPushMirrorSyncBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo push mirror sync bad request response has a 3xx status code
func (o *RepoPushMirrorSyncBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo push mirror sync bad request response has a 4xx status code
func (o *RepoPushMirrorSyncBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo push mirror sync bad request response has a 5xx status code
func (o *RepoPushMirrorSyncBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this repo push mirror sync bad request response a status code equal to that given
func (o *RepoPushMirrorSyncBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the repo push mirror sync bad request response
func (o *RepoPushMirrorSyncBadRequest) Code() int {
	return 400
}

func (o *RepoPushMirrorSyncBadRequest) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/push_mirrors-sync][%d] repoPushMirrorSyncBadRequest", 400)
}

func (o *RepoPushMirrorSyncBadRequest) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/push_mirrors-sync][%d] repoPushMirrorSyncBadRequest", 400)
}

func (o *RepoPushMirrorSyncBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRepoPushMirrorSyncForbidden creates a RepoPushMirrorSyncForbidden with default headers values
func NewRepoPushMirrorSyncForbidden() *RepoPushMirrorSyncForbidden {
	return &RepoPushMirrorSyncForbidden{}
}

/*
RepoPushMirrorSyncForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type RepoPushMirrorSyncForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this repo push mirror sync forbidden response has a 2xx status code
func (o *RepoPushMirrorSyncForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo push mirror sync forbidden response has a 3xx status code
func (o *RepoPushMirrorSyncForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo push mirror sync forbidden response has a 4xx status code
func (o *RepoPushMirrorSyncForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo push mirror sync forbidden response has a 5xx status code
func (o *RepoPushMirrorSyncForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this repo push mirror sync forbidden response a status code equal to that given
func (o *RepoPushMirrorSyncForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the repo push mirror sync forbidden response
func (o *RepoPushMirrorSyncForbidden) Code() int {
	return 403
}

func (o *RepoPushMirrorSyncForbidden) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/push_mirrors-sync][%d] repoPushMirrorSyncForbidden", 403)
}

func (o *RepoPushMirrorSyncForbidden) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/push_mirrors-sync][%d] repoPushMirrorSyncForbidden", 403)
}

func (o *RepoPushMirrorSyncForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRepoPushMirrorSyncNotFound creates a RepoPushMirrorSyncNotFound with default headers values
func NewRepoPushMirrorSyncNotFound() *RepoPushMirrorSyncNotFound {
	return &RepoPushMirrorSyncNotFound{}
}

/*
RepoPushMirrorSyncNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type RepoPushMirrorSyncNotFound struct {
}

// IsSuccess returns true when this repo push mirror sync not found response has a 2xx status code
func (o *RepoPushMirrorSyncNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this repo push mirror sync not found response has a 3xx status code
func (o *RepoPushMirrorSyncNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this repo push mirror sync not found response has a 4xx status code
func (o *RepoPushMirrorSyncNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this repo push mirror sync not found response has a 5xx status code
func (o *RepoPushMirrorSyncNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this repo push mirror sync not found response a status code equal to that given
func (o *RepoPushMirrorSyncNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the repo push mirror sync not found response
func (o *RepoPushMirrorSyncNotFound) Code() int {
	return 404
}

func (o *RepoPushMirrorSyncNotFound) Error() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/push_mirrors-sync][%d] repoPushMirrorSyncNotFound", 404)
}

func (o *RepoPushMirrorSyncNotFound) String() string {
	return fmt.Sprintf("[POST /repos/{owner}/{repo}/push_mirrors-sync][%d] repoPushMirrorSyncNotFound", 404)
}

func (o *RepoPushMirrorSyncNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
