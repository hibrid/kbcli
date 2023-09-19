// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

// GetInvoicesGroupReader is a Reader for the GetInvoicesGroup structure.
type GetInvoicesGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoicesGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoicesGroupOK()
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

// NewGetInvoicesGroupOK creates a GetInvoicesGroupOK with default headers values
func NewGetInvoicesGroupOK() *GetInvoicesGroupOK {
	return &GetInvoicesGroupOK{}
}

/* GetInvoicesGroupOK describes a response with status code 200, with default header values.

successful operation
*/
type GetInvoicesGroupOK struct {
	Payload      []*kbmodel.Invoice
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get invoices group o k response
func (o *GetInvoicesGroupOK) Code() int {
	return 200
}

// IsSuccess returns true when this get invoices group o k response has a 2xx status code
func (o *GetInvoicesGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get invoices group o k response has a 3xx status code
func (o *GetInvoicesGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices group o k response has a 4xx status code
func (o *GetInvoicesGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get invoices group o k response has a 5xx status code
func (o *GetInvoicesGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices group o k response a status code equal to that given
func (o *GetInvoicesGroupOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetInvoicesGroupOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{groupId}/group][%d] getInvoicesGroupOK  %+v", 200, o.Payload)
}

func (o *GetInvoicesGroupOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{groupId}/group][%d] getInvoicesGroupOK  %+v", 200, o.Payload)
}

func (o *GetInvoicesGroupOK) GetPayload() []*kbmodel.Invoice {
	return o.Payload
}

func (o *GetInvoicesGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicesGroupBadRequest creates a GetInvoicesGroupBadRequest with default headers values
func NewGetInvoicesGroupBadRequest() *GetInvoicesGroupBadRequest {
	return &GetInvoicesGroupBadRequest{}
}

/* GetInvoicesGroupBadRequest describes a response with status code 400, with default header values.

Invalid group id supplied
*/
type GetInvoicesGroupBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get invoices group bad request response
func (o *GetInvoicesGroupBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this get invoices group bad request response has a 2xx status code
func (o *GetInvoicesGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoices group bad request response has a 3xx status code
func (o *GetInvoicesGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices group bad request response has a 4xx status code
func (o *GetInvoicesGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoices group bad request response has a 5xx status code
func (o *GetInvoicesGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices group bad request response a status code equal to that given
func (o *GetInvoicesGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetInvoicesGroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{groupId}/group][%d] getInvoicesGroupBadRequest ", 400)
}

func (o *GetInvoicesGroupBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{groupId}/group][%d] getInvoicesGroupBadRequest ", 400)
}

func (o *GetInvoicesGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
