// Code generated by go-swagger; DO NOT EDIT.

package miscellaneous

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RenderMarkdownReader is a Reader for the RenderMarkdown structure.
type RenderMarkdownReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenderMarkdownReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRenderMarkdownOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewRenderMarkdownUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /markdown] renderMarkdown", response, response.Code())
	}
}

// NewRenderMarkdownOK creates a RenderMarkdownOK with default headers values
func NewRenderMarkdownOK() *RenderMarkdownOK {
	return &RenderMarkdownOK{}
}

/*
RenderMarkdownOK describes a response with status code 200, with default header values.

MarkdownRender is a rendered markdown document
*/
type RenderMarkdownOK struct {
	Payload string
}

// IsSuccess returns true when this render markdown o k response has a 2xx status code
func (o *RenderMarkdownOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this render markdown o k response has a 3xx status code
func (o *RenderMarkdownOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this render markdown o k response has a 4xx status code
func (o *RenderMarkdownOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this render markdown o k response has a 5xx status code
func (o *RenderMarkdownOK) IsServerError() bool {
	return false
}

// IsCode returns true when this render markdown o k response a status code equal to that given
func (o *RenderMarkdownOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the render markdown o k response
func (o *RenderMarkdownOK) Code() int {
	return 200
}

func (o *RenderMarkdownOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /markdown][%d] renderMarkdownOK %s", 200, payload)
}

func (o *RenderMarkdownOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /markdown][%d] renderMarkdownOK %s", 200, payload)
}

func (o *RenderMarkdownOK) GetPayload() string {
	return o.Payload
}

func (o *RenderMarkdownOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenderMarkdownUnprocessableEntity creates a RenderMarkdownUnprocessableEntity with default headers values
func NewRenderMarkdownUnprocessableEntity() *RenderMarkdownUnprocessableEntity {
	return &RenderMarkdownUnprocessableEntity{}
}

/*
RenderMarkdownUnprocessableEntity describes a response with status code 422, with default header values.

APIValidationError is error format response related to input validation
*/
type RenderMarkdownUnprocessableEntity struct {
	Message string
	URL     string
}

// IsSuccess returns true when this render markdown unprocessable entity response has a 2xx status code
func (o *RenderMarkdownUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this render markdown unprocessable entity response has a 3xx status code
func (o *RenderMarkdownUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this render markdown unprocessable entity response has a 4xx status code
func (o *RenderMarkdownUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this render markdown unprocessable entity response has a 5xx status code
func (o *RenderMarkdownUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this render markdown unprocessable entity response a status code equal to that given
func (o *RenderMarkdownUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the render markdown unprocessable entity response
func (o *RenderMarkdownUnprocessableEntity) Code() int {
	return 422
}

func (o *RenderMarkdownUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /markdown][%d] renderMarkdownUnprocessableEntity", 422)
}

func (o *RenderMarkdownUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /markdown][%d] renderMarkdownUnprocessableEntity", 422)
}

func (o *RenderMarkdownUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
