// Code generated by go-swagger; DO NOT EDIT.

package account

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

// PayAllInvoicesReader is a Reader for the PayAllInvoices structure.
type PayAllInvoicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PayAllInvoicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewPayAllInvoicesCreated()
		result.HttpResponse = response
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewPayAllInvoicesNoContent()
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

// NewPayAllInvoicesCreated creates a PayAllInvoicesCreated with default headers values
func NewPayAllInvoicesCreated() *PayAllInvoicesCreated {
	return &PayAllInvoicesCreated{}
}

/* PayAllInvoicesCreated describes a response with status code 201, with default header values.

Successful operation
*/
type PayAllInvoicesCreated struct {
	Payload      []*kbmodel.Invoice
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the pay all invoices created response
func (o *PayAllInvoicesCreated) Code() int {
	return 201
}

// IsSuccess returns true when this pay all invoices created response has a 2xx status code
func (o *PayAllInvoicesCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pay all invoices created response has a 3xx status code
func (o *PayAllInvoicesCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pay all invoices created response has a 4xx status code
func (o *PayAllInvoicesCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this pay all invoices created response has a 5xx status code
func (o *PayAllInvoicesCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this pay all invoices created response a status code equal to that given
func (o *PayAllInvoicesCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PayAllInvoicesCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/invoicePayments][%d] payAllInvoicesCreated  %+v", 201, o.Payload)
}

func (o *PayAllInvoicesCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/invoicePayments][%d] payAllInvoicesCreated  %+v", 201, o.Payload)
}

func (o *PayAllInvoicesCreated) GetPayload() []*kbmodel.Invoice {
	return o.Payload
}

func (o *PayAllInvoicesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPayAllInvoicesNoContent creates a PayAllInvoicesNoContent with default headers values
func NewPayAllInvoicesNoContent() *PayAllInvoicesNoContent {
	return &PayAllInvoicesNoContent{}
}

/* PayAllInvoicesNoContent describes a response with status code 204, with default header values.

Nothing to pay
*/
type PayAllInvoicesNoContent struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the pay all invoices no content response
func (o *PayAllInvoicesNoContent) Code() int {
	return 204
}

// IsSuccess returns true when this pay all invoices no content response has a 2xx status code
func (o *PayAllInvoicesNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pay all invoices no content response has a 3xx status code
func (o *PayAllInvoicesNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pay all invoices no content response has a 4xx status code
func (o *PayAllInvoicesNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this pay all invoices no content response has a 5xx status code
func (o *PayAllInvoicesNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this pay all invoices no content response a status code equal to that given
func (o *PayAllInvoicesNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *PayAllInvoicesNoContent) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/invoicePayments][%d] payAllInvoicesNoContent ", 204)
}

func (o *PayAllInvoicesNoContent) String() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/invoicePayments][%d] payAllInvoicesNoContent ", 204)
}

func (o *PayAllInvoicesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPayAllInvoicesNotFound creates a PayAllInvoicesNotFound with default headers values
func NewPayAllInvoicesNotFound() *PayAllInvoicesNotFound {
	return &PayAllInvoicesNotFound{}
}

/* PayAllInvoicesNotFound describes a response with status code 404, with default header values.

Invalid account id supplied
*/
type PayAllInvoicesNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the pay all invoices not found response
func (o *PayAllInvoicesNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this pay all invoices not found response has a 2xx status code
func (o *PayAllInvoicesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pay all invoices not found response has a 3xx status code
func (o *PayAllInvoicesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pay all invoices not found response has a 4xx status code
func (o *PayAllInvoicesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pay all invoices not found response has a 5xx status code
func (o *PayAllInvoicesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pay all invoices not found response a status code equal to that given
func (o *PayAllInvoicesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PayAllInvoicesNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/invoicePayments][%d] payAllInvoicesNotFound ", 404)
}

func (o *PayAllInvoicesNotFound) String() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/invoicePayments][%d] payAllInvoicesNotFound ", 404)
}

func (o *PayAllInvoicesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
