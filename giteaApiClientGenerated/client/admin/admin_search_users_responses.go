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

// AdminSearchUsersReader is a Reader for the AdminSearchUsers structure.
type AdminSearchUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminSearchUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminSearchUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAdminSearchUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /admin/users] adminSearchUsers", response, response.Code())
	}
}

// NewAdminSearchUsersOK creates a AdminSearchUsersOK with default headers values
func NewAdminSearchUsersOK() *AdminSearchUsersOK {
	return &AdminSearchUsersOK{}
}

/*
AdminSearchUsersOK describes a response with status code 200, with default header values.

UserList
*/
type AdminSearchUsersOK struct {
	Payload []*models.User
}

// IsSuccess returns true when this admin search users o k response has a 2xx status code
func (o *AdminSearchUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin search users o k response has a 3xx status code
func (o *AdminSearchUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin search users o k response has a 4xx status code
func (o *AdminSearchUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin search users o k response has a 5xx status code
func (o *AdminSearchUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin search users o k response a status code equal to that given
func (o *AdminSearchUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the admin search users o k response
func (o *AdminSearchUsersOK) Code() int {
	return 200
}

func (o *AdminSearchUsersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/users][%d] adminSearchUsersOK %s", 200, payload)
}

func (o *AdminSearchUsersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/users][%d] adminSearchUsersOK %s", 200, payload)
}

func (o *AdminSearchUsersOK) GetPayload() []*models.User {
	return o.Payload
}

func (o *AdminSearchUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminSearchUsersForbidden creates a AdminSearchUsersForbidden with default headers values
func NewAdminSearchUsersForbidden() *AdminSearchUsersForbidden {
	return &AdminSearchUsersForbidden{}
}

/*
AdminSearchUsersForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type AdminSearchUsersForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this admin search users forbidden response has a 2xx status code
func (o *AdminSearchUsersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin search users forbidden response has a 3xx status code
func (o *AdminSearchUsersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin search users forbidden response has a 4xx status code
func (o *AdminSearchUsersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin search users forbidden response has a 5xx status code
func (o *AdminSearchUsersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin search users forbidden response a status code equal to that given
func (o *AdminSearchUsersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the admin search users forbidden response
func (o *AdminSearchUsersForbidden) Code() int {
	return 403
}

func (o *AdminSearchUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /admin/users][%d] adminSearchUsersForbidden", 403)
}

func (o *AdminSearchUsersForbidden) String() string {
	return fmt.Sprintf("[GET /admin/users][%d] adminSearchUsersForbidden", 403)
}

func (o *AdminSearchUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
