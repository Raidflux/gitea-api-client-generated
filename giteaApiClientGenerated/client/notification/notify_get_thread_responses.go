// Code generated by go-swagger; DO NOT EDIT.

package notification

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

// NotifyGetThreadReader is a Reader for the NotifyGetThread structure.
type NotifyGetThreadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotifyGetThreadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNotifyGetThreadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewNotifyGetThreadForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNotifyGetThreadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /notifications/threads/{id}] notifyGetThread", response, response.Code())
	}
}

// NewNotifyGetThreadOK creates a NotifyGetThreadOK with default headers values
func NewNotifyGetThreadOK() *NotifyGetThreadOK {
	return &NotifyGetThreadOK{}
}

/*
NotifyGetThreadOK describes a response with status code 200, with default header values.

NotificationThread
*/
type NotifyGetThreadOK struct {
	Payload *models.NotificationThread
}

// IsSuccess returns true when this notify get thread o k response has a 2xx status code
func (o *NotifyGetThreadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this notify get thread o k response has a 3xx status code
func (o *NotifyGetThreadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notify get thread o k response has a 4xx status code
func (o *NotifyGetThreadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this notify get thread o k response has a 5xx status code
func (o *NotifyGetThreadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this notify get thread o k response a status code equal to that given
func (o *NotifyGetThreadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the notify get thread o k response
func (o *NotifyGetThreadOK) Code() int {
	return 200
}

func (o *NotifyGetThreadOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /notifications/threads/{id}][%d] notifyGetThreadOK %s", 200, payload)
}

func (o *NotifyGetThreadOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /notifications/threads/{id}][%d] notifyGetThreadOK %s", 200, payload)
}

func (o *NotifyGetThreadOK) GetPayload() *models.NotificationThread {
	return o.Payload
}

func (o *NotifyGetThreadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotificationThread)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotifyGetThreadForbidden creates a NotifyGetThreadForbidden with default headers values
func NewNotifyGetThreadForbidden() *NotifyGetThreadForbidden {
	return &NotifyGetThreadForbidden{}
}

/*
NotifyGetThreadForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type NotifyGetThreadForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this notify get thread forbidden response has a 2xx status code
func (o *NotifyGetThreadForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notify get thread forbidden response has a 3xx status code
func (o *NotifyGetThreadForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notify get thread forbidden response has a 4xx status code
func (o *NotifyGetThreadForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this notify get thread forbidden response has a 5xx status code
func (o *NotifyGetThreadForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this notify get thread forbidden response a status code equal to that given
func (o *NotifyGetThreadForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the notify get thread forbidden response
func (o *NotifyGetThreadForbidden) Code() int {
	return 403
}

func (o *NotifyGetThreadForbidden) Error() string {
	return fmt.Sprintf("[GET /notifications/threads/{id}][%d] notifyGetThreadForbidden", 403)
}

func (o *NotifyGetThreadForbidden) String() string {
	return fmt.Sprintf("[GET /notifications/threads/{id}][%d] notifyGetThreadForbidden", 403)
}

func (o *NotifyGetThreadForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewNotifyGetThreadNotFound creates a NotifyGetThreadNotFound with default headers values
func NewNotifyGetThreadNotFound() *NotifyGetThreadNotFound {
	return &NotifyGetThreadNotFound{}
}

/*
NotifyGetThreadNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type NotifyGetThreadNotFound struct {
}

// IsSuccess returns true when this notify get thread not found response has a 2xx status code
func (o *NotifyGetThreadNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notify get thread not found response has a 3xx status code
func (o *NotifyGetThreadNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notify get thread not found response has a 4xx status code
func (o *NotifyGetThreadNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this notify get thread not found response has a 5xx status code
func (o *NotifyGetThreadNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this notify get thread not found response a status code equal to that given
func (o *NotifyGetThreadNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the notify get thread not found response
func (o *NotifyGetThreadNotFound) Code() int {
	return 404
}

func (o *NotifyGetThreadNotFound) Error() string {
	return fmt.Sprintf("[GET /notifications/threads/{id}][%d] notifyGetThreadNotFound", 404)
}

func (o *NotifyGetThreadNotFound) String() string {
	return fmt.Sprintf("[GET /notifications/threads/{id}][%d] notifyGetThreadNotFound", 404)
}

func (o *NotifyGetThreadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
