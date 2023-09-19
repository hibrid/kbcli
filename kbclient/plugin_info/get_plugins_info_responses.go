// Code generated by go-swagger; DO NOT EDIT.

package plugin_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"

	"github.com/killbill/kbcli/v3/kbmodel"
)

// GetPluginsInfoReader is a Reader for the GetPluginsInfo structure.
type GetPluginsInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPluginsInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPluginsInfoOK()
		result.HttpResponse = response
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		errorResult := kbcommon.NewKillbillError(response.Code())
		if err := consumer.Consume(response.Body(), &errorResult); err != nil && err != io.EOF {
			return nil, err
		}
		return nil, errorResult
	}
}

// NewGetPluginsInfoOK creates a GetPluginsInfoOK with default headers values
func NewGetPluginsInfoOK() *GetPluginsInfoOK {
	return &GetPluginsInfoOK{}
}

/* GetPluginsInfoOK describes a response with status code 200, with default header values.

successful operation
*/
type GetPluginsInfoOK struct {
	Payload      []*kbmodel.PluginInfo
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get plugins info o k response
func (o *GetPluginsInfoOK) Code() int {
	return 200
}

// IsSuccess returns true when this get plugins info o k response has a 2xx status code
func (o *GetPluginsInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get plugins info o k response has a 3xx status code
func (o *GetPluginsInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get plugins info o k response has a 4xx status code
func (o *GetPluginsInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get plugins info o k response has a 5xx status code
func (o *GetPluginsInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get plugins info o k response a status code equal to that given
func (o *GetPluginsInfoOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPluginsInfoOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/pluginsInfo][%d] getPluginsInfoOK  %+v", 200, o.Payload)
}

func (o *GetPluginsInfoOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/pluginsInfo][%d] getPluginsInfoOK  %+v", 200, o.Payload)
}

func (o *GetPluginsInfoOK) GetPayload() []*kbmodel.PluginInfo {
	return o.Payload
}

func (o *GetPluginsInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
