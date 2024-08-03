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

	"giteaApiClientGenerated/models"
)

// NotifyReadThreadReader is a Reader for the NotifyReadThread structure.
type NotifyReadThreadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NotifyReadThreadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 205:
		result := NewNotifyReadThreadResetContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewNotifyReadThreadForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewNotifyReadThreadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /notifications/threads/{id}] notifyReadThread", response, response.Code())
	}
}

// NewNotifyReadThreadResetContent creates a NotifyReadThreadResetContent with default headers values
func NewNotifyReadThreadResetContent() *NotifyReadThreadResetContent {
	return &NotifyReadThreadResetContent{}
}

/*
NotifyReadThreadResetContent describes a response with status code 205, with default header values.

NotificationThread
*/
type NotifyReadThreadResetContent struct {
	Payload *models.NotificationThread
}

// IsSuccess returns true when this notify read thread reset content response has a 2xx status code
func (o *NotifyReadThreadResetContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this notify read thread reset content response has a 3xx status code
func (o *NotifyReadThreadResetContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notify read thread reset content response has a 4xx status code
func (o *NotifyReadThreadResetContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this notify read thread reset content response has a 5xx status code
func (o *NotifyReadThreadResetContent) IsServerError() bool {
	return false
}

// IsCode returns true when this notify read thread reset content response a status code equal to that given
func (o *NotifyReadThreadResetContent) IsCode(code int) bool {
	return code == 205
}

// Code gets the status code for the notify read thread reset content response
func (o *NotifyReadThreadResetContent) Code() int {
	return 205
}

func (o *NotifyReadThreadResetContent) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /notifications/threads/{id}][%d] notifyReadThreadResetContent %s", 205, payload)
}

func (o *NotifyReadThreadResetContent) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /notifications/threads/{id}][%d] notifyReadThreadResetContent %s", 205, payload)
}

func (o *NotifyReadThreadResetContent) GetPayload() *models.NotificationThread {
	return o.Payload
}

func (o *NotifyReadThreadResetContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotificationThread)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNotifyReadThreadForbidden creates a NotifyReadThreadForbidden with default headers values
func NewNotifyReadThreadForbidden() *NotifyReadThreadForbidden {
	return &NotifyReadThreadForbidden{}
}

/*
NotifyReadThreadForbidden describes a response with status code 403, with default header values.

APIForbiddenError is a forbidden error response
*/
type NotifyReadThreadForbidden struct {
	Message string
	URL     string
}

// IsSuccess returns true when this notify read thread forbidden response has a 2xx status code
func (o *NotifyReadThreadForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notify read thread forbidden response has a 3xx status code
func (o *NotifyReadThreadForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notify read thread forbidden response has a 4xx status code
func (o *NotifyReadThreadForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this notify read thread forbidden response has a 5xx status code
func (o *NotifyReadThreadForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this notify read thread forbidden response a status code equal to that given
func (o *NotifyReadThreadForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the notify read thread forbidden response
func (o *NotifyReadThreadForbidden) Code() int {
	return 403
}

func (o *NotifyReadThreadForbidden) Error() string {
	return fmt.Sprintf("[PATCH /notifications/threads/{id}][%d] notifyReadThreadForbidden", 403)
}

func (o *NotifyReadThreadForbidden) String() string {
	return fmt.Sprintf("[PATCH /notifications/threads/{id}][%d] notifyReadThreadForbidden", 403)
}

func (o *NotifyReadThreadForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewNotifyReadThreadNotFound creates a NotifyReadThreadNotFound with default headers values
func NewNotifyReadThreadNotFound() *NotifyReadThreadNotFound {
	return &NotifyReadThreadNotFound{}
}

/*
NotifyReadThreadNotFound describes a response with status code 404, with default header values.

APINotFound is a not found empty response
*/
type NotifyReadThreadNotFound struct {
}

// IsSuccess returns true when this notify read thread not found response has a 2xx status code
func (o *NotifyReadThreadNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this notify read thread not found response has a 3xx status code
func (o *NotifyReadThreadNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this notify read thread not found response has a 4xx status code
func (o *NotifyReadThreadNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this notify read thread not found response has a 5xx status code
func (o *NotifyReadThreadNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this notify read thread not found response a status code equal to that given
func (o *NotifyReadThreadNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the notify read thread not found response
func (o *NotifyReadThreadNotFound) Code() int {
	return 404
}

func (o *NotifyReadThreadNotFound) Error() string {
	return fmt.Sprintf("[PATCH /notifications/threads/{id}][%d] notifyReadThreadNotFound", 404)
}

func (o *NotifyReadThreadNotFound) String() string {
	return fmt.Sprintf("[PATCH /notifications/threads/{id}][%d] notifyReadThreadNotFound", 404)
}

func (o *NotifyReadThreadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
