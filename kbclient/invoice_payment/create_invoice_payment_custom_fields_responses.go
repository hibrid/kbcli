// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

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

// CreateInvoicePaymentCustomFieldsReader is a Reader for the CreateInvoicePaymentCustomFields structure.
type CreateInvoicePaymentCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInvoicePaymentCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateInvoicePaymentCustomFieldsCreated()
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

// NewCreateInvoicePaymentCustomFieldsCreated creates a CreateInvoicePaymentCustomFieldsCreated with default headers values
func NewCreateInvoicePaymentCustomFieldsCreated() *CreateInvoicePaymentCustomFieldsCreated {
	return &CreateInvoicePaymentCustomFieldsCreated{}
}

/* CreateInvoicePaymentCustomFieldsCreated describes a response with status code 201, with default header values.

Custom field created successfully
*/
type CreateInvoicePaymentCustomFieldsCreated struct {
	Payload      []*kbmodel.CustomField
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the create invoice payment custom fields created response
func (o *CreateInvoicePaymentCustomFieldsCreated) Code() int {
	return 201
}

// IsSuccess returns true when this create invoice payment custom fields created response has a 2xx status code
func (o *CreateInvoicePaymentCustomFieldsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create invoice payment custom fields created response has a 3xx status code
func (o *CreateInvoicePaymentCustomFieldsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create invoice payment custom fields created response has a 4xx status code
func (o *CreateInvoicePaymentCustomFieldsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create invoice payment custom fields created response has a 5xx status code
func (o *CreateInvoicePaymentCustomFieldsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create invoice payment custom fields created response a status code equal to that given
func (o *CreateInvoicePaymentCustomFieldsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateInvoicePaymentCustomFieldsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/customFields][%d] createInvoicePaymentCustomFieldsCreated  %+v", 201, o.Payload)
}

func (o *CreateInvoicePaymentCustomFieldsCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/customFields][%d] createInvoicePaymentCustomFieldsCreated  %+v", 201, o.Payload)
}

func (o *CreateInvoicePaymentCustomFieldsCreated) GetPayload() []*kbmodel.CustomField {
	return o.Payload
}

func (o *CreateInvoicePaymentCustomFieldsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInvoicePaymentCustomFieldsBadRequest creates a CreateInvoicePaymentCustomFieldsBadRequest with default headers values
func NewCreateInvoicePaymentCustomFieldsBadRequest() *CreateInvoicePaymentCustomFieldsBadRequest {
	return &CreateInvoicePaymentCustomFieldsBadRequest{}
}

/* CreateInvoicePaymentCustomFieldsBadRequest describes a response with status code 400, with default header values.

Invalid payment id supplied
*/
type CreateInvoicePaymentCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the create invoice payment custom fields bad request response
func (o *CreateInvoicePaymentCustomFieldsBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this create invoice payment custom fields bad request response has a 2xx status code
func (o *CreateInvoicePaymentCustomFieldsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create invoice payment custom fields bad request response has a 3xx status code
func (o *CreateInvoicePaymentCustomFieldsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create invoice payment custom fields bad request response has a 4xx status code
func (o *CreateInvoicePaymentCustomFieldsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create invoice payment custom fields bad request response has a 5xx status code
func (o *CreateInvoicePaymentCustomFieldsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create invoice payment custom fields bad request response a status code equal to that given
func (o *CreateInvoicePaymentCustomFieldsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateInvoicePaymentCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/customFields][%d] createInvoicePaymentCustomFieldsBadRequest ", 400)
}

func (o *CreateInvoicePaymentCustomFieldsBadRequest) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/customFields][%d] createInvoicePaymentCustomFieldsBadRequest ", 400)
}

func (o *CreateInvoicePaymentCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
