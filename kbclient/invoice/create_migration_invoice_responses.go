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

// CreateMigrationInvoiceReader is a Reader for the CreateMigrationInvoice structure.
type CreateMigrationInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMigrationInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateMigrationInvoiceCreated()
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

// NewCreateMigrationInvoiceCreated creates a CreateMigrationInvoiceCreated with default headers values
func NewCreateMigrationInvoiceCreated() *CreateMigrationInvoiceCreated {
	return &CreateMigrationInvoiceCreated{}
}

/* CreateMigrationInvoiceCreated describes a response with status code 201, with default header values.

Created migration invoice successfully
*/
type CreateMigrationInvoiceCreated struct {
	Payload      *kbmodel.Invoice
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the create migration invoice created response
func (o *CreateMigrationInvoiceCreated) Code() int {
	return 201
}

// IsSuccess returns true when this create migration invoice created response has a 2xx status code
func (o *CreateMigrationInvoiceCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create migration invoice created response has a 3xx status code
func (o *CreateMigrationInvoiceCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create migration invoice created response has a 4xx status code
func (o *CreateMigrationInvoiceCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create migration invoice created response has a 5xx status code
func (o *CreateMigrationInvoiceCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create migration invoice created response a status code equal to that given
func (o *CreateMigrationInvoiceCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateMigrationInvoiceCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/migration/{accountId}][%d] createMigrationInvoiceCreated  %+v", 201, o.Payload)
}

func (o *CreateMigrationInvoiceCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/migration/{accountId}][%d] createMigrationInvoiceCreated  %+v", 201, o.Payload)
}

func (o *CreateMigrationInvoiceCreated) GetPayload() *kbmodel.Invoice {
	return o.Payload
}

func (o *CreateMigrationInvoiceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Invoice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMigrationInvoiceBadRequest creates a CreateMigrationInvoiceBadRequest with default headers values
func NewCreateMigrationInvoiceBadRequest() *CreateMigrationInvoiceBadRequest {
	return &CreateMigrationInvoiceBadRequest{}
}

/* CreateMigrationInvoiceBadRequest describes a response with status code 400, with default header values.

Invalid account id or target datetime supplied
*/
type CreateMigrationInvoiceBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the create migration invoice bad request response
func (o *CreateMigrationInvoiceBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this create migration invoice bad request response has a 2xx status code
func (o *CreateMigrationInvoiceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create migration invoice bad request response has a 3xx status code
func (o *CreateMigrationInvoiceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create migration invoice bad request response has a 4xx status code
func (o *CreateMigrationInvoiceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create migration invoice bad request response has a 5xx status code
func (o *CreateMigrationInvoiceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create migration invoice bad request response a status code equal to that given
func (o *CreateMigrationInvoiceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateMigrationInvoiceBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/migration/{accountId}][%d] createMigrationInvoiceBadRequest ", 400)
}

func (o *CreateMigrationInvoiceBadRequest) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/migration/{accountId}][%d] createMigrationInvoiceBadRequest ", 400)
}

func (o *CreateMigrationInvoiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
