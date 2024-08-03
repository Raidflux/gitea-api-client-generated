// Code generated by go-swagger; DO NOT EDIT.

package settings

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

// GetGeneralAttachmentSettingsReader is a Reader for the GetGeneralAttachmentSettings structure.
type GetGeneralAttachmentSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGeneralAttachmentSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGeneralAttachmentSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /settings/attachment] getGeneralAttachmentSettings", response, response.Code())
	}
}

// NewGetGeneralAttachmentSettingsOK creates a GetGeneralAttachmentSettingsOK with default headers values
func NewGetGeneralAttachmentSettingsOK() *GetGeneralAttachmentSettingsOK {
	return &GetGeneralAttachmentSettingsOK{}
}

/*
GetGeneralAttachmentSettingsOK describes a response with status code 200, with default header values.

GeneralAttachmentSettings
*/
type GetGeneralAttachmentSettingsOK struct {
	Payload *models.GeneralAttachmentSettings
}

// IsSuccess returns true when this get general attachment settings o k response has a 2xx status code
func (o *GetGeneralAttachmentSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get general attachment settings o k response has a 3xx status code
func (o *GetGeneralAttachmentSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get general attachment settings o k response has a 4xx status code
func (o *GetGeneralAttachmentSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get general attachment settings o k response has a 5xx status code
func (o *GetGeneralAttachmentSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get general attachment settings o k response a status code equal to that given
func (o *GetGeneralAttachmentSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get general attachment settings o k response
func (o *GetGeneralAttachmentSettingsOK) Code() int {
	return 200
}

func (o *GetGeneralAttachmentSettingsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /settings/attachment][%d] getGeneralAttachmentSettingsOK %s", 200, payload)
}

func (o *GetGeneralAttachmentSettingsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /settings/attachment][%d] getGeneralAttachmentSettingsOK %s", 200, payload)
}

func (o *GetGeneralAttachmentSettingsOK) GetPayload() *models.GeneralAttachmentSettings {
	return o.Payload
}

func (o *GetGeneralAttachmentSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralAttachmentSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
