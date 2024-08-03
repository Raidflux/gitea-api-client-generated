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

// ListLabelTemplatesReader is a Reader for the ListLabelTemplates structure.
type ListLabelTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListLabelTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListLabelTemplatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /label/templates] listLabelTemplates", response, response.Code())
	}
}

// NewListLabelTemplatesOK creates a ListLabelTemplatesOK with default headers values
func NewListLabelTemplatesOK() *ListLabelTemplatesOK {
	return &ListLabelTemplatesOK{}
}

/*
ListLabelTemplatesOK describes a response with status code 200, with default header values.

LabelTemplateList
*/
type ListLabelTemplatesOK struct {
	Payload []string
}

// IsSuccess returns true when this list label templates o k response has a 2xx status code
func (o *ListLabelTemplatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list label templates o k response has a 3xx status code
func (o *ListLabelTemplatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list label templates o k response has a 4xx status code
func (o *ListLabelTemplatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list label templates o k response has a 5xx status code
func (o *ListLabelTemplatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list label templates o k response a status code equal to that given
func (o *ListLabelTemplatesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list label templates o k response
func (o *ListLabelTemplatesOK) Code() int {
	return 200
}

func (o *ListLabelTemplatesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /label/templates][%d] listLabelTemplatesOK %s", 200, payload)
}

func (o *ListLabelTemplatesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /label/templates][%d] listLabelTemplatesOK %s", 200, payload)
}

func (o *ListLabelTemplatesOK) GetPayload() []string {
	return o.Payload
}

func (o *ListLabelTemplatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
