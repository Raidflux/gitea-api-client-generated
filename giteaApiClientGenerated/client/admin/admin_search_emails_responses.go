// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// AdminSearchEmailsReader is a Reader for the AdminSearchEmails structure.
type AdminSearchEmailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminSearchEmailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminSearchEmailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAdminSearchEmailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /admin/emails/search] adminSearchEmails", response, response.Code())
	}
}

// NewAdminSearchEmailsOK creates a AdminSearchEmailsOK with default headers values
func NewAdminSearchEmailsOK() *AdminSearchEmailsOK {
	return &AdminSearchEmailsOK{}
}

/*
AdminSearchEmailsOK describes a response with status code 200, with default header values.

EmailList
*/
type AdminSearchEmailsOK struct {
	Payload []*models.Email
}

// IsSuccess returns true when this admin search emails o k response has a 2xx status code
func (o *AdminSearchEmailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin search emails o k response has a 3xx status code
func (o *AdminSearchEmailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin search emails o k response has a 4xx status code
func (o *AdminSearchEmailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin search emails o k response has a 5xx status code
func (o *AdminSearchEmailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin search emails o k response a status code equal to that given
func (o *AdminSearchEmailsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the admin search emails o k response
func (o *AdminSearchEmailsOK) Code() int {
	return 200
}

func (o *AdminSearchEmailsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/emails/search][%d] adminSearchEmailsOK %s", 200, payload)
}

func (o *AdminSearchEmailsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/emails/search][%d] adminSearchEmailsOK %s", 200, payload)
}

func (o *AdminSearchEmailsOK) GetPayload() []*models.Email {
	return o.Payload
}

func (o *AdminSearchEmailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminSearchEmailsForbidden creates a AdminSearchEmailsForbidden with default headers values
func NewAdminSearchEmailsForbidden() *AdminSearchEmailsForbidden {
	return &AdminSearchEmailsForbidden{}
}

/*
AdminSearchEmailsForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type AdminSearchEmailsForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this admin search emails forbidden response has a 2xx status code
func (o *AdminSearchEmailsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin search emails forbidden response has a 3xx status code
func (o *AdminSearchEmailsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin search emails forbidden response has a 4xx status code
func (o *AdminSearchEmailsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin search emails forbidden response has a 5xx status code
func (o *AdminSearchEmailsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin search emails forbidden response a status code equal to that given
func (o *AdminSearchEmailsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the admin search emails forbidden response
func (o *AdminSearchEmailsForbidden) Code() int {
	return 403
}

func (o *AdminSearchEmailsForbidden) Error() string {
	return fmt.Sprintf("[GET /admin/emails/search][%d] adminSearchEmailsForbidden", 403)
}

func (o *AdminSearchEmailsForbidden) String() string {
	return fmt.Sprintf("[GET /admin/emails/search][%d] adminSearchEmailsForbidden", 403)
}

func (o *AdminSearchEmailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
