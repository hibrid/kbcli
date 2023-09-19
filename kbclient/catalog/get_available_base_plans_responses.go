// Code generated by go-swagger; DO NOT EDIT.

package catalog

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

// GetAvailableBasePlansReader is a Reader for the GetAvailableBasePlans structure.
type GetAvailableBasePlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAvailableBasePlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAvailableBasePlansOK()
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

// NewGetAvailableBasePlansOK creates a GetAvailableBasePlansOK with default headers values
func NewGetAvailableBasePlansOK() *GetAvailableBasePlansOK {
	return &GetAvailableBasePlansOK{}
}

/* GetAvailableBasePlansOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAvailableBasePlansOK struct {
	Payload      []*kbmodel.PlanDetail
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get available base plans o k response
func (o *GetAvailableBasePlansOK) Code() int {
	return 200
}

// IsSuccess returns true when this get available base plans o k response has a 2xx status code
func (o *GetAvailableBasePlansOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get available base plans o k response has a 3xx status code
func (o *GetAvailableBasePlansOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get available base plans o k response has a 4xx status code
func (o *GetAvailableBasePlansOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get available base plans o k response has a 5xx status code
func (o *GetAvailableBasePlansOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get available base plans o k response a status code equal to that given
func (o *GetAvailableBasePlansOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAvailableBasePlansOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/catalog/availableBasePlans][%d] getAvailableBasePlansOK  %+v", 200, o.Payload)
}

func (o *GetAvailableBasePlansOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/catalog/availableBasePlans][%d] getAvailableBasePlansOK  %+v", 200, o.Payload)
}

func (o *GetAvailableBasePlansOK) GetPayload() []*kbmodel.PlanDetail {
	return o.Payload
}

func (o *GetAvailableBasePlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
