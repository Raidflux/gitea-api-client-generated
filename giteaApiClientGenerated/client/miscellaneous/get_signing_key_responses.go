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

// GetSigningKeyReader is a Reader for the GetSigningKey structure.
type GetSigningKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSigningKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSigningKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /signing-key.gpg] getSigningKey", response, response.Code())
	}
}

// NewGetSigningKeyOK creates a GetSigningKeyOK with default headers values
func NewGetSigningKeyOK() *GetSigningKeyOK {
	return &GetSigningKeyOK{}
}

/*
GetSigningKeyOK describes a response with status code 200, with default header values.

GPG armored public key
*/
type GetSigningKeyOK struct {
	Payload string
}

// IsSuccess returns true when this get signing key o k response has a 2xx status code
func (o *GetSigningKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get signing key o k response has a 3xx status code
func (o *GetSigningKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get signing key o k response has a 4xx status code
func (o *GetSigningKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get signing key o k response has a 5xx status code
func (o *GetSigningKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get signing key o k response a status code equal to that given
func (o *GetSigningKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get signing key o k response
func (o *GetSigningKeyOK) Code() int {
	return 200
}

func (o *GetSigningKeyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /signing-key.gpg][%d] getSigningKeyOK %s", 200, payload)
}

func (o *GetSigningKeyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /signing-key.gpg][%d] getSigningKeyOK %s", 200, payload)
}

func (o *GetSigningKeyOK) GetPayload() string {
	return o.Payload
}

func (o *GetSigningKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
