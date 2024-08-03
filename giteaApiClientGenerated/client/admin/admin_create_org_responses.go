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

	"github.com/Raidflux/gitea-api-client-generated/giteaApiClientGenerated/models"
)

// AdminCreateOrgReader is a Reader for the AdminCreateOrg structure.
type AdminCreateOrgReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminCreateOrgReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAdminCreateOrgCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAdminCreateOrgForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewAdminCreateOrgUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /admin/users/{username}/orgs] adminCreateOrg", response, response.Code())
	}
}

// NewAdminCreateOrgCreated creates a AdminCreateOrgCreated with default headers values
func NewAdminCreateOrgCreated() *AdminCreateOrgCreated {
	return &AdminCreateOrgCreated{}
}

/*
AdminCreateOrgCreated describes a response with status code 201, with default header values.

Organization
*/
type AdminCreateOrgCreated struct {
	Payload *models.Organization
}

// IsSuccess returns true when this admin create org created response has a 2xx status code
func (o *AdminCreateOrgCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin create org created response has a 3xx status code
func (o *AdminCreateOrgCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin create org created response has a 4xx status code
func (o *AdminCreateOrgCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin create org created response has a 5xx status code
func (o *AdminCreateOrgCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this admin create org created response a status code equal to that given
func (o *AdminCreateOrgCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the admin create org created response
func (o *AdminCreateOrgCreated) Code() int {
	return 201
}

func (o *AdminCreateOrgCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /admin/users/{username}/orgs][%d] adminCreateOrgCreated %s", 201, payload)
}

func (o *AdminCreateOrgCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /admin/users/{username}/orgs][%d] adminCreateOrgCreated %s", 201, payload)
}

func (o *AdminCreateOrgCreated) GetPayload() *models.Organization {
	return o.Payload
}

func (o *AdminCreateOrgCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Organization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminCreateOrgForbidden creates a AdminCreateOrgForbidden with default headers values
func NewAdminCreateOrgForbidden() *AdminCreateOrgForbidden {
	return &AdminCreateOrgForbidden{}
}

/*
AdminCreateOrgForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type AdminCreateOrgForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this admin create org forbidden response has a 2xx status code
func (o *AdminCreateOrgForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin create org forbidden response has a 3xx status code
func (o *AdminCreateOrgForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin create org forbidden response has a 4xx status code
func (o *AdminCreateOrgForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin create org forbidden response has a 5xx status code
func (o *AdminCreateOrgForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin create org forbidden response a status code equal to that given
func (o *AdminCreateOrgForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the admin create org forbidden response
func (o *AdminCreateOrgForbidden) Code() int {
	return 403
}

func (o *AdminCreateOrgForbidden) Error() string {
	return fmt.Sprintf("[POST /admin/users/{username}/orgs][%d] adminCreateOrgForbidden", 403)
}

func (o *AdminCreateOrgForbidden) String() string {
	return fmt.Sprintf("[POST /admin/users/{username}/orgs][%d] adminCreateOrgForbidden", 403)
}

func (o *AdminCreateOrgForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAdminCreateOrgUnprocessableEntity creates a AdminCreateOrgUnprocessableEntity with default headers values
func NewAdminCreateOrgUnprocessableEntity() *AdminCreateOrgUnprocessableEntity {
	return &AdminCreateOrgUnprocessableEntity{}
}

/*
AdminCreateOrgUnprocessableEntity describes a response with status code 422, with default header values.

APIValidationError is error format response related to input validation
*/
type AdminCreateOrgUnprocessableEntity struct {
	Message string
	URL     string
}

// IsSuccess returns true when this admin create org unprocessable entity response has a 2xx status code
func (o *AdminCreateOrgUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin create org unprocessable entity response has a 3xx status code
func (o *AdminCreateOrgUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin create org unprocessable entity response has a 4xx status code
func (o *AdminCreateOrgUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin create org unprocessable entity response has a 5xx status code
func (o *AdminCreateOrgUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this admin create org unprocessable entity response a status code equal to that given
func (o *AdminCreateOrgUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the admin create org unprocessable entity response
func (o *AdminCreateOrgUnprocessableEntity) Code() int {
	return 422
}

func (o *AdminCreateOrgUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /admin/users/{username}/orgs][%d] adminCreateOrgUnprocessableEntity", 422)
}

func (o *AdminCreateOrgUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /admin/users/{username}/orgs][%d] adminCreateOrgUnprocessableEntity", 422)
}

func (o *AdminCreateOrgUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
