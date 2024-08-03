// Code generated by go-swagger; DO NOT EDIT.

package user

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

// UserCreateOAuth2ApplicationReader is a Reader for the UserCreateOAuth2Application structure.
type UserCreateOAuth2ApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCreateOAuth2ApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUserCreateOAuth2ApplicationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserCreateOAuth2ApplicationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /user/applications/oauth2] userCreateOAuth2Application", response, response.Code())
	}
}

// NewUserCreateOAuth2ApplicationCreated creates a UserCreateOAuth2ApplicationCreated with default headers values
func NewUserCreateOAuth2ApplicationCreated() *UserCreateOAuth2ApplicationCreated {
	return &UserCreateOAuth2ApplicationCreated{}
}

/*
UserCreateOAuth2ApplicationCreated describes a response with status code 201, with default header values.

OAuth2Application
*/
type UserCreateOAuth2ApplicationCreated struct {
	Payload *models.OAuth2Application
}

// IsSuccess returns true when this user create o auth2 application created response has a 2xx status code
func (o *UserCreateOAuth2ApplicationCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user create o auth2 application created response has a 3xx status code
func (o *UserCreateOAuth2ApplicationCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user create o auth2 application created response has a 4xx status code
func (o *UserCreateOAuth2ApplicationCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this user create o auth2 application created response has a 5xx status code
func (o *UserCreateOAuth2ApplicationCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this user create o auth2 application created response a status code equal to that given
func (o *UserCreateOAuth2ApplicationCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the user create o auth2 application created response
func (o *UserCreateOAuth2ApplicationCreated) Code() int {
	return 201
}

func (o *UserCreateOAuth2ApplicationCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/applications/oauth2][%d] userCreateOAuth2ApplicationCreated %s", 201, payload)
}

func (o *UserCreateOAuth2ApplicationCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /user/applications/oauth2][%d] userCreateOAuth2ApplicationCreated %s", 201, payload)
}

func (o *UserCreateOAuth2ApplicationCreated) GetPayload() *models.OAuth2Application {
	return o.Payload
}

func (o *UserCreateOAuth2ApplicationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuth2Application)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserCreateOAuth2ApplicationBadRequest creates a UserCreateOAuth2ApplicationBadRequest with default headers values
func NewUserCreateOAuth2ApplicationBadRequest() *UserCreateOAuth2ApplicationBadRequest {
	return &UserCreateOAuth2ApplicationBadRequest{}
}

/*
UserCreateOAuth2ApplicationBadRequest describes a response with status code 400, with default header values.

APIError is error format response
*/
type UserCreateOAuth2ApplicationBadRequest struct {
	Message string
	URL     string
}

// IsSuccess returns true when this user create o auth2 application bad request response has a 2xx status code
func (o *UserCreateOAuth2ApplicationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user create o auth2 application bad request response has a 3xx status code
func (o *UserCreateOAuth2ApplicationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user create o auth2 application bad request response has a 4xx status code
func (o *UserCreateOAuth2ApplicationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this user create o auth2 application bad request response has a 5xx status code
func (o *UserCreateOAuth2ApplicationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this user create o auth2 application bad request response a status code equal to that given
func (o *UserCreateOAuth2ApplicationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the user create o auth2 application bad request response
func (o *UserCreateOAuth2ApplicationBadRequest) Code() int {
	return 400
}

func (o *UserCreateOAuth2ApplicationBadRequest) Error() string {
	return fmt.Sprintf("[POST /user/applications/oauth2][%d] userCreateOAuth2ApplicationBadRequest", 400)
}

func (o *UserCreateOAuth2ApplicationBadRequest) String() string {
	return fmt.Sprintf("[POST /user/applications/oauth2][%d] userCreateOAuth2ApplicationBadRequest", 400)
}

func (o *UserCreateOAuth2ApplicationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
