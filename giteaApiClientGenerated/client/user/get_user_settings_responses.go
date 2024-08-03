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

// GetUserSettingsReader is a Reader for the GetUserSettings structure.
type GetUserSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /user/settings] getUserSettings", response, response.Code())
	}
}

// NewGetUserSettingsOK creates a GetUserSettingsOK with default headers values
func NewGetUserSettingsOK() *GetUserSettingsOK {
	return &GetUserSettingsOK{}
}

/*
GetUserSettingsOK describes a response with status code 200, with default header values.

UserSettings
*/
type GetUserSettingsOK struct {
	Payload []*models.UserSettings
}

// IsSuccess returns true when this get user settings o k response has a 2xx status code
func (o *GetUserSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user settings o k response has a 3xx status code
func (o *GetUserSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user settings o k response has a 4xx status code
func (o *GetUserSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user settings o k response has a 5xx status code
func (o *GetUserSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user settings o k response a status code equal to that given
func (o *GetUserSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get user settings o k response
func (o *GetUserSettingsOK) Code() int {
	return 200
}

func (o *GetUserSettingsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/settings][%d] getUserSettingsOK %s", 200, payload)
}

func (o *GetUserSettingsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user/settings][%d] getUserSettingsOK %s", 200, payload)
}

func (o *GetUserSettingsOK) GetPayload() []*models.UserSettings {
	return o.Payload
}

func (o *GetUserSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
