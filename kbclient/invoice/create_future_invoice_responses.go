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

// CreateFutureInvoiceReader is a Reader for the CreateFutureInvoice structure.
type CreateFutureInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFutureInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateFutureInvoiceCreated()
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

// NewCreateFutureInvoiceCreated creates a CreateFutureInvoiceCreated with default headers values
func NewCreateFutureInvoiceCreated() *CreateFutureInvoiceCreated {
	return &CreateFutureInvoiceCreated{}
}

/* CreateFutureInvoiceCreated describes a response with status code 201, with default header values.

Created invoice successfully
*/
type CreateFutureInvoiceCreated struct {
	Payload      *kbmodel.Invoice
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the create future invoice created response
func (o *CreateFutureInvoiceCreated) Code() int {
	return 201
}

// IsSuccess returns true when this create future invoice created response has a 2xx status code
func (o *CreateFutureInvoiceCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create future invoice created response has a 3xx status code
func (o *CreateFutureInvoiceCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create future invoice created response has a 4xx status code
func (o *CreateFutureInvoiceCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create future invoice created response has a 5xx status code
func (o *CreateFutureInvoiceCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create future invoice created response a status code equal to that given
func (o *CreateFutureInvoiceCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateFutureInvoiceCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices][%d] createFutureInvoiceCreated  %+v", 201, o.Payload)
}

func (o *CreateFutureInvoiceCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices][%d] createFutureInvoiceCreated  %+v", 201, o.Payload)
}

func (o *CreateFutureInvoiceCreated) GetPayload() *kbmodel.Invoice {
	return o.Payload
}

func (o *CreateFutureInvoiceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Invoice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFutureInvoiceBadRequest creates a CreateFutureInvoiceBadRequest with default headers values
func NewCreateFutureInvoiceBadRequest() *CreateFutureInvoiceBadRequest {
	return &CreateFutureInvoiceBadRequest{}
}

/* CreateFutureInvoiceBadRequest describes a response with status code 400, with default header values.

Invalid account id or target datetime supplied
*/
type CreateFutureInvoiceBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the create future invoice bad request response
func (o *CreateFutureInvoiceBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this create future invoice bad request response has a 2xx status code
func (o *CreateFutureInvoiceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create future invoice bad request response has a 3xx status code
func (o *CreateFutureInvoiceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create future invoice bad request response has a 4xx status code
func (o *CreateFutureInvoiceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create future invoice bad request response has a 5xx status code
func (o *CreateFutureInvoiceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create future invoice bad request response a status code equal to that given
func (o *CreateFutureInvoiceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateFutureInvoiceBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices][%d] createFutureInvoiceBadRequest ", 400)
}

func (o *CreateFutureInvoiceBadRequest) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices][%d] createFutureInvoiceBadRequest ", 400)
}

func (o *CreateFutureInvoiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
