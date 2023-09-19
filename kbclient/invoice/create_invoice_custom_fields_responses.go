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

// CreateInvoiceCustomFieldsReader is a Reader for the CreateInvoiceCustomFields structure.
type CreateInvoiceCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInvoiceCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateInvoiceCustomFieldsCreated()
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

// NewCreateInvoiceCustomFieldsCreated creates a CreateInvoiceCustomFieldsCreated with default headers values
func NewCreateInvoiceCustomFieldsCreated() *CreateInvoiceCustomFieldsCreated {
	return &CreateInvoiceCustomFieldsCreated{}
}

/* CreateInvoiceCustomFieldsCreated describes a response with status code 201, with default header values.

Custom field created successfully
*/
type CreateInvoiceCustomFieldsCreated struct {
	Payload      []*kbmodel.CustomField
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the create invoice custom fields created response
func (o *CreateInvoiceCustomFieldsCreated) Code() int {
	return 201
}

// IsSuccess returns true when this create invoice custom fields created response has a 2xx status code
func (o *CreateInvoiceCustomFieldsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create invoice custom fields created response has a 3xx status code
func (o *CreateInvoiceCustomFieldsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create invoice custom fields created response has a 4xx status code
func (o *CreateInvoiceCustomFieldsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create invoice custom fields created response has a 5xx status code
func (o *CreateInvoiceCustomFieldsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create invoice custom fields created response a status code equal to that given
func (o *CreateInvoiceCustomFieldsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateInvoiceCustomFieldsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/{invoiceId}/customFields][%d] createInvoiceCustomFieldsCreated  %+v", 201, o.Payload)
}

func (o *CreateInvoiceCustomFieldsCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/{invoiceId}/customFields][%d] createInvoiceCustomFieldsCreated  %+v", 201, o.Payload)
}

func (o *CreateInvoiceCustomFieldsCreated) GetPayload() []*kbmodel.CustomField {
	return o.Payload
}

func (o *CreateInvoiceCustomFieldsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInvoiceCustomFieldsBadRequest creates a CreateInvoiceCustomFieldsBadRequest with default headers values
func NewCreateInvoiceCustomFieldsBadRequest() *CreateInvoiceCustomFieldsBadRequest {
	return &CreateInvoiceCustomFieldsBadRequest{}
}

/* CreateInvoiceCustomFieldsBadRequest describes a response with status code 400, with default header values.

Invalid invoice id supplied
*/
type CreateInvoiceCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the create invoice custom fields bad request response
func (o *CreateInvoiceCustomFieldsBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this create invoice custom fields bad request response has a 2xx status code
func (o *CreateInvoiceCustomFieldsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create invoice custom fields bad request response has a 3xx status code
func (o *CreateInvoiceCustomFieldsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create invoice custom fields bad request response has a 4xx status code
func (o *CreateInvoiceCustomFieldsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create invoice custom fields bad request response has a 5xx status code
func (o *CreateInvoiceCustomFieldsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create invoice custom fields bad request response a status code equal to that given
func (o *CreateInvoiceCustomFieldsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateInvoiceCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/{invoiceId}/customFields][%d] createInvoiceCustomFieldsBadRequest ", 400)
}

func (o *CreateInvoiceCustomFieldsBadRequest) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/{invoiceId}/customFields][%d] createInvoiceCustomFieldsBadRequest ", 400)
}

func (o *CreateInvoiceCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
