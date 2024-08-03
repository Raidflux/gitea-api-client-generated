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

// AdminGetAllOrgsReader is a Reader for the AdminGetAllOrgs structure.
type AdminGetAllOrgsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminGetAllOrgsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminGetAllOrgsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAdminGetAllOrgsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /admin/orgs] adminGetAllOrgs", response, response.Code())
	}
}

// NewAdminGetAllOrgsOK creates a AdminGetAllOrgsOK with default headers values
func NewAdminGetAllOrgsOK() *AdminGetAllOrgsOK {
	return &AdminGetAllOrgsOK{}
}

/*
AdminGetAllOrgsOK describes a response with status code 200, with default header values.

OrganizationList
*/
type AdminGetAllOrgsOK struct {
	Payload []*models.Organization
}

// IsSuccess returns true when this admin get all orgs o k response has a 2xx status code
func (o *AdminGetAllOrgsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin get all orgs o k response has a 3xx status code
func (o *AdminGetAllOrgsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin get all orgs o k response has a 4xx status code
func (o *AdminGetAllOrgsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin get all orgs o k response has a 5xx status code
func (o *AdminGetAllOrgsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin get all orgs o k response a status code equal to that given
func (o *AdminGetAllOrgsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the admin get all orgs o k response
func (o *AdminGetAllOrgsOK) Code() int {
	return 200
}

func (o *AdminGetAllOrgsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/orgs][%d] adminGetAllOrgsOK %s", 200, payload)
}

func (o *AdminGetAllOrgsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/orgs][%d] adminGetAllOrgsOK %s", 200, payload)
}

func (o *AdminGetAllOrgsOK) GetPayload() []*models.Organization {
	return o.Payload
}

func (o *AdminGetAllOrgsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminGetAllOrgsForbidden creates a AdminGetAllOrgsForbidden with default headers values
func NewAdminGetAllOrgsForbidden() *AdminGetAllOrgsForbidden {
	return &AdminGetAllOrgsForbidden{}
}

/*
AdminGetAllOrgsForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type AdminGetAllOrgsForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this admin get all orgs forbidden response has a 2xx status code
func (o *AdminGetAllOrgsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin get all orgs forbidden response has a 3xx status code
func (o *AdminGetAllOrgsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin get all orgs forbidden response has a 4xx status code
func (o *AdminGetAllOrgsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin get all orgs forbidden response has a 5xx status code
func (o *AdminGetAllOrgsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin get all orgs forbidden response a status code equal to that given
func (o *AdminGetAllOrgsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the admin get all orgs forbidden response
func (o *AdminGetAllOrgsForbidden) Code() int {
	return 403
}

func (o *AdminGetAllOrgsForbidden) Error() string {
	return fmt.Sprintf("[GET /admin/orgs][%d] adminGetAllOrgsForbidden", 403)
}

func (o *AdminGetAllOrgsForbidden) String() string {
	return fmt.Sprintf("[GET /admin/orgs][%d] adminGetAllOrgsForbidden", 403)
}

func (o *AdminGetAllOrgsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
