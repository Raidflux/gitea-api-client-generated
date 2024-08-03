// Code generated by go-swagger; DO NOT EDIT.

package user

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

// GetUserVariableReader is a Reader for the GetUserVariable structure.
type GetUserVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserVariableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserVariableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserVariableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /user/actions/variables/{variablename}] getUserVariable", response, response.Code())
	}
}

// NewGetUserVariableOK creates a GetUserVariableOK with default headers values
func NewGetUserVariableOK() *GetUserVariableOK {
	return &GetUserVariableOK{}
}

/*
GetUserVariableOK describes a response with status code 200, with default header values.

ActionVariable
*/
type GetUserVariableOK struct {
	Payload *models.ActionVariable
}

// IsSuccess returns true when this get user variable o k response has a 2xx status code
func (o *GetUserVariableOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user variable o k response has a 3xx status code
func (o *GetUserVariableOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user variable o k response has a 4xx status code
func (o *GetUserVariableOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user variable o k response has a 5xx status code
func (o *GetUserVariableOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user variable o k response a status code equal to that given
func (o *GetUserVariableOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get user variable o k response
func (o *GetUserVariableOK) Code() int {
	return 200
}

func (o *GetUserVariableOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/actions/variables/{variablename}][%d] getUserVariableOK %s", 200, payload)
}

func (o *GetUserVariableOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/actions/variables/{variablename}][%d] getUserVariableOK %s", 200, payload)
}

func (o *GetUserVariableOK) GetPayload() *models.ActionVariable {
	return o.Payload
}

func (o *GetUserVariableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionVariable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserVariableBadRequest creates a GetUserVariableBadRequest with default headers values
func NewGetUserVariableBadRequest() *GetUserVariableBadRequest {
	return &GetUserVariableBadRequest{}
}

/*
GetUserVariableBadRequest describes a response with status code 400, with default header values.

APIError is error format response
*/
type GetUserVariableBadRequest struct {
	Message string
	URL     string
}

// IsSuccess returns true when this get user variable bad request response has a 2xx status code
func (o *GetUserVariableBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user variable bad request response has a 3xx status code
func (o *GetUserVariableBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user variable bad request response has a 4xx status code
func (o *GetUserVariableBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user variable bad request response has a 5xx status code
func (o *GetUserVariableBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get user variable bad request response a status code equal to that given
func (o *GetUserVariableBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get user variable bad request response
func (o *GetUserVariableBadRequest) Code() int {
	return 400
}

func (o *GetUserVariableBadRequest) Error() string {
	return fmt.Sprintf("[GET /user/actions/variables/{variablename}][%d] getUserVariableBadRequest", 400)
}

func (o *GetUserVariableBadRequest) String() string {
	return fmt.Sprintf("[GET /user/actions/variables/{variablename}][%d] getUserVariableBadRequest", 400)
}

func (o *GetUserVariableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUserVariableNotFound creates a GetUserVariableNotFound with default headers values
func NewGetUserVariableNotFound() *GetUserVariableNotFound {
	return &GetUserVariableNotFound{}
}

/*
GetUserVariableNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type GetUserVariableNotFound struct {
}

// IsSuccess returns true when this get user variable not found response has a 2xx status code
func (o *GetUserVariableNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user variable not found response has a 3xx status code
func (o *GetUserVariableNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user variable not found response has a 4xx status code
func (o *GetUserVariableNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user variable not found response has a 5xx status code
func (o *GetUserVariableNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get user variable not found response a status code equal to that given
func (o *GetUserVariableNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get user variable not found response
func (o *GetUserVariableNotFound) Code() int {
	return 404
}

func (o *GetUserVariableNotFound) Error() string {
	return fmt.Sprintf("[GET /user/actions/variables/{variablename}][%d] getUserVariableNotFound", 404)
}

func (o *GetUserVariableNotFound) String() string {
	return fmt.Sprintf("[GET /user/actions/variables/{variablename}][%d] getUserVariableNotFound", 404)
}

func (o *GetUserVariableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
