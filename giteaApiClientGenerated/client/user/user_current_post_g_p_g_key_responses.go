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

// UserCurrentPostGPGKeyReader is a Reader for the UserCurrentPostGPGKey structure.
type UserCurrentPostGPGKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCurrentPostGPGKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUserCurrentPostGPGKeyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUserCurrentPostGPGKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUserCurrentPostGPGKeyUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /user/gpg_keys] userCurrentPostGPGKey", response, response.Code())
	}
}

// NewUserCurrentPostGPGKeyCreated creates a UserCurrentPostGPGKeyCreated with default headers values
func NewUserCurrentPostGPGKeyCreated() *UserCurrentPostGPGKeyCreated {
	return &UserCurrentPostGPGKeyCreated{}
}

/*
UserCurrentPostGPGKeyCreated describes a response with status code 201, with default header values.

GPGKey
*/
type UserCurrentPostGPGKeyCreated struct {
	Payload *models.GPGKey
}

// IsSuccess returns true when this user current post g p g key created response has a 2xx status code
func (o *UserCurrentPostGPGKeyCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user current post g p g key created response has a 3xx status code
func (o *UserCurrentPostGPGKeyCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user current post g p g key created response has a 4xx status code
func (o *UserCurrentPostGPGKeyCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this user current post g p g key created response has a 5xx status code
func (o *UserCurrentPostGPGKeyCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this user current post g p g key created response a status code equal to that given
func (o *UserCurrentPostGPGKeyCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the user current post g p g key created response
func (o *UserCurrentPostGPGKeyCreated) Code() int {
	return 201
}

func (o *UserCurrentPostGPGKeyCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/gpg_keys][%d] userCurrentPostGPGKeyCreated %s", 201, payload)
}

func (o *UserCurrentPostGPGKeyCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/gpg_keys][%d] userCurrentPostGPGKeyCreated %s", 201, payload)
}

func (o *UserCurrentPostGPGKeyCreated) GetPayload() *models.GPGKey {
	return o.Payload
}

func (o *UserCurrentPostGPGKeyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GPGKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserCurrentPostGPGKeyNotFound creates a UserCurrentPostGPGKeyNotFound with default headers values
func NewUserCurrentPostGPGKeyNotFound() *UserCurrentPostGPGKeyNotFound {
	return &UserCurrentPostGPGKeyNotFound{}
}

/*
UserCurrentPostGPGKeyNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type UserCurrentPostGPGKeyNotFound struct {
}

// IsSuccess returns true when this user current post g p g key not found response has a 2xx status code
func (o *UserCurrentPostGPGKeyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user current post g p g key not found response has a 3xx status code
func (o *UserCurrentPostGPGKeyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user current post g p g key not found response has a 4xx status code
func (o *UserCurrentPostGPGKeyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user current post g p g key not found response has a 5xx status code
func (o *UserCurrentPostGPGKeyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user current post g p g key not found response a status code equal to that given
func (o *UserCurrentPostGPGKeyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the user current post g p g key not found response
func (o *UserCurrentPostGPGKeyNotFound) Code() int {
	return 404
}

func (o *UserCurrentPostGPGKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /user/gpg_keys][%d] userCurrentPostGPGKeyNotFound", 404)
}

func (o *UserCurrentPostGPGKeyNotFound) String() string {
	return fmt.Sprintf("[POST /user/gpg_keys][%d] userCurrentPostGPGKeyNotFound", 404)
}

func (o *UserCurrentPostGPGKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserCurrentPostGPGKeyUnprocessableEntity creates a UserCurrentPostGPGKeyUnprocessableEntity with default headers values
func NewUserCurrentPostGPGKeyUnprocessableEntity() *UserCurrentPostGPGKeyUnprocessableEntity {
	return &UserCurrentPostGPGKeyUnprocessableEntity{}
}

/*
UserCurrentPostGPGKeyUnprocessableEntity describes a response with status code 422, with default header values.

APIValidationError is error format response related to input validation
*/
type UserCurrentPostGPGKeyUnprocessableEntity struct {
	Message string
	URL     string
}

// IsSuccess returns true when this user current post g p g key unprocessable entity response has a 2xx status code
func (o *UserCurrentPostGPGKeyUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user current post g p g key unprocessable entity response has a 3xx status code
func (o *UserCurrentPostGPGKeyUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user current post g p g key unprocessable entity response has a 4xx status code
func (o *UserCurrentPostGPGKeyUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this user current post g p g key unprocessable entity response has a 5xx status code
func (o *UserCurrentPostGPGKeyUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this user current post g p g key unprocessable entity response a status code equal to that given
func (o *UserCurrentPostGPGKeyUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the user current post g p g key unprocessable entity response
func (o *UserCurrentPostGPGKeyUnprocessableEntity) Code() int {
	return 422
}

func (o *UserCurrentPostGPGKeyUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /user/gpg_keys][%d] userCurrentPostGPGKeyUnprocessableEntity", 422)
}

func (o *UserCurrentPostGPGKeyUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /user/gpg_keys][%d] userCurrentPostGPGKeyUnprocessableEntity", 422)
}

func (o *UserCurrentPostGPGKeyUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
