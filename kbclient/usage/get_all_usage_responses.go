// Code generated by go-swagger; DO NOT EDIT.

package usage

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

// GetAllUsageReader is a Reader for the GetAllUsage structure.
type GetAllUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllUsageOK()
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

// NewGetAllUsageOK creates a GetAllUsageOK with default headers values
func NewGetAllUsageOK() *GetAllUsageOK {
	return &GetAllUsageOK{}
}

/* GetAllUsageOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAllUsageOK struct {
	Payload      *kbmodel.RolledUpUsage
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get all usage o k response
func (o *GetAllUsageOK) Code() int {
	return 200
}

// IsSuccess returns true when this get all usage o k response has a 2xx status code
func (o *GetAllUsageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all usage o k response has a 3xx status code
func (o *GetAllUsageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all usage o k response has a 4xx status code
func (o *GetAllUsageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all usage o k response has a 5xx status code
func (o *GetAllUsageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all usage o k response a status code equal to that given
func (o *GetAllUsageOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAllUsageOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/usages/{subscriptionId}][%d] getAllUsageOK  %+v", 200, o.Payload)
}

func (o *GetAllUsageOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/usages/{subscriptionId}][%d] getAllUsageOK  %+v", 200, o.Payload)
}

func (o *GetAllUsageOK) GetPayload() *kbmodel.RolledUpUsage {
	return o.Payload
}

func (o *GetAllUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.RolledUpUsage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllUsageBadRequest creates a GetAllUsageBadRequest with default headers values
func NewGetAllUsageBadRequest() *GetAllUsageBadRequest {
	return &GetAllUsageBadRequest{}
}

/* GetAllUsageBadRequest describes a response with status code 400, with default header values.

Missing start date or end date
*/
type GetAllUsageBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get all usage bad request response
func (o *GetAllUsageBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this get all usage bad request response has a 2xx status code
func (o *GetAllUsageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all usage bad request response has a 3xx status code
func (o *GetAllUsageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all usage bad request response has a 4xx status code
func (o *GetAllUsageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all usage bad request response has a 5xx status code
func (o *GetAllUsageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get all usage bad request response a status code equal to that given
func (o *GetAllUsageBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAllUsageBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/usages/{subscriptionId}][%d] getAllUsageBadRequest ", 400)
}

func (o *GetAllUsageBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/usages/{subscriptionId}][%d] getAllUsageBadRequest ", 400)
}

func (o *GetAllUsageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
