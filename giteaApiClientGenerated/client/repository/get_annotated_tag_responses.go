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

// GetAnnotatedTagReader is a Reader for the GetAnnotatedTag structure.
type GetAnnotatedTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnnotatedTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnnotatedTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAnnotatedTagBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAnnotatedTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /repos/{owner}/{repo}/git/tags/{sha}] GetAnnotatedTag", response, response.Code())
	}
}

// NewGetAnnotatedTagOK creates a GetAnnotatedTagOK with default headers values
func NewGetAnnotatedTagOK() *GetAnnotatedTagOK {
	return &GetAnnotatedTagOK{}
}

/*
GetAnnotatedTagOK describes a response with status code 200, with default header values.

AnnotatedTag
*/
type GetAnnotatedTagOK struct {
	Payload *models.AnnotatedTag
}

// IsSuccess returns true when this get annotated tag o k response has a 2xx status code
func (o *GetAnnotatedTagOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get annotated tag o k response has a 3xx status code
func (o *GetAnnotatedTagOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get annotated tag o k response has a 4xx status code
func (o *GetAnnotatedTagOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get annotated tag o k response has a 5xx status code
func (o *GetAnnotatedTagOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get annotated tag o k response a status code equal to that given
func (o *GetAnnotatedTagOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get annotated tag o k response
func (o *GetAnnotatedTagOK) Code() int {
	return 200
}

func (o *GetAnnotatedTagOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/tags/{sha}][%d] getAnnotatedTagOK %s", 200, payload)
}

func (o *GetAnnotatedTagOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/tags/{sha}][%d] getAnnotatedTagOK %s", 200, payload)
}

func (o *GetAnnotatedTagOK) GetPayload() *models.AnnotatedTag {
	return o.Payload
}

func (o *GetAnnotatedTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AnnotatedTag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnnotatedTagBadRequest creates a GetAnnotatedTagBadRequest with default headers values
func NewGetAnnotatedTagBadRequest() *GetAnnotatedTagBadRequest {
	return &GetAnnotatedTagBadRequest{}
}

/*
GetAnnotatedTagBadRequest describes a response with status code 400, with default header values.

APIError is error format response
*/
type GetAnnotatedTagBadRequest struct {
	Message string
	URL     string
}

// IsSuccess returns true when this get annotated tag bad request response has a 2xx status code
func (o *GetAnnotatedTagBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get annotated tag bad request response has a 3xx status code
func (o *GetAnnotatedTagBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get annotated tag bad request response has a 4xx status code
func (o *GetAnnotatedTagBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get annotated tag bad request response has a 5xx status code
func (o *GetAnnotatedTagBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get annotated tag bad request response a status code equal to that given
func (o *GetAnnotatedTagBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get annotated tag bad request response
func (o *GetAnnotatedTagBadRequest) Code() int {
	return 400
}

func (o *GetAnnotatedTagBadRequest) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/tags/{sha}][%d] getAnnotatedTagBadRequest", 400)
}

func (o *GetAnnotatedTagBadRequest) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/tags/{sha}][%d] getAnnotatedTagBadRequest", 400)
}

func (o *GetAnnotatedTagBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAnnotatedTagNotFound creates a GetAnnotatedTagNotFound with default headers values
func NewGetAnnotatedTagNotFound() *GetAnnotatedTagNotFound {
	return &GetAnnotatedTagNotFound{}
}

/*
GetAnnotatedTagNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type GetAnnotatedTagNotFound struct {
}

// IsSuccess returns true when this get annotated tag not found response has a 2xx status code
func (o *GetAnnotatedTagNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get annotated tag not found response has a 3xx status code
func (o *GetAnnotatedTagNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get annotated tag not found response has a 4xx status code
func (o *GetAnnotatedTagNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get annotated tag not found response has a 5xx status code
func (o *GetAnnotatedTagNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get annotated tag not found response a status code equal to that given
func (o *GetAnnotatedTagNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get annotated tag not found response
func (o *GetAnnotatedTagNotFound) Code() int {
	return 404
}

func (o *GetAnnotatedTagNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/tags/{sha}][%d] getAnnotatedTagNotFound", 404)
}

func (o *GetAnnotatedTagNotFound) String() string {
	return fmt.Sprintf("[GET /repos/{owner}/{repo}/git/tags/{sha}][%d] getAnnotatedTagNotFound", 404)
}

func (o *GetAnnotatedTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
