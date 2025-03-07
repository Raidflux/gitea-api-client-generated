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

// AdminCronListReader is a Reader for the AdminCronList structure.
type AdminCronListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminCronListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminCronListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAdminCronListForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /admin/cron] adminCronList", response, response.Code())
	}
}

// NewAdminCronListOK creates a AdminCronListOK with default headers values
func NewAdminCronListOK() *AdminCronListOK {
	return &AdminCronListOK{}
}

/*
AdminCronListOK describes a response with status code 200, with default header values.

CronList
*/
type AdminCronListOK struct {
	Payload []*models.Cron
}

// IsSuccess returns true when this admin cron list o k response has a 2xx status code
func (o *AdminCronListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin cron list o k response has a 3xx status code
func (o *AdminCronListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin cron list o k response has a 4xx status code
func (o *AdminCronListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin cron list o k response has a 5xx status code
func (o *AdminCronListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin cron list o k response a status code equal to that given
func (o *AdminCronListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the admin cron list o k response
func (o *AdminCronListOK) Code() int {
	return 200
}

func (o *AdminCronListOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/cron][%d] adminCronListOK %s", 200, payload)
}

func (o *AdminCronListOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /admin/cron][%d] adminCronListOK %s", 200, payload)
}

func (o *AdminCronListOK) GetPayload() []*models.Cron {
	return o.Payload
}

func (o *AdminCronListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminCronListForbidden creates a AdminCronListForbidden with default headers values
func NewAdminCronListForbidden() *AdminCronListForbidden {
	return &AdminCronListForbidden{}
}

/*
AdminCronListForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type AdminCronListForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this admin cron list forbidden response has a 2xx status code
func (o *AdminCronListForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin cron list forbidden response has a 3xx status code
func (o *AdminCronListForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin cron list forbidden response has a 4xx status code
func (o *AdminCronListForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin cron list forbidden response has a 5xx status code
func (o *AdminCronListForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin cron list forbidden response a status code equal to that given
func (o *AdminCronListForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the admin cron list forbidden response
func (o *AdminCronListForbidden) Code() int {
	return 403
}

func (o *AdminCronListForbidden) Error() string {
	return fmt.Sprintf("[GET /admin/cron][%d] adminCronListForbidden", 403)
}

func (o *AdminCronListForbidden) String() string {
	return fmt.Sprintf("[GET /admin/cron][%d] adminCronListForbidden", 403)
}

func (o *AdminCronListForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
