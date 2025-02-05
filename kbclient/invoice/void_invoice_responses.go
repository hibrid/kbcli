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
)

// VoidInvoiceReader is a Reader for the VoidInvoice structure.
type VoidInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VoidInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewVoidInvoiceNoContent()
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

// NewVoidInvoiceNoContent creates a VoidInvoiceNoContent with default headers values
func NewVoidInvoiceNoContent() *VoidInvoiceNoContent {
	return &VoidInvoiceNoContent{}
}

/* VoidInvoiceNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type VoidInvoiceNoContent struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the void invoice no content response
func (o *VoidInvoiceNoContent) Code() int {
	return 204
}

// IsSuccess returns true when this void invoice no content response has a 2xx status code
func (o *VoidInvoiceNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this void invoice no content response has a 3xx status code
func (o *VoidInvoiceNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this void invoice no content response has a 4xx status code
func (o *VoidInvoiceNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this void invoice no content response has a 5xx status code
func (o *VoidInvoiceNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this void invoice no content response a status code equal to that given
func (o *VoidInvoiceNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *VoidInvoiceNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoices/{invoiceId}/voidInvoice][%d] voidInvoiceNoContent ", 204)
}

func (o *VoidInvoiceNoContent) String() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoices/{invoiceId}/voidInvoice][%d] voidInvoiceNoContent ", 204)
}

func (o *VoidInvoiceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVoidInvoiceBadRequest creates a VoidInvoiceBadRequest with default headers values
func NewVoidInvoiceBadRequest() *VoidInvoiceBadRequest {
	return &VoidInvoiceBadRequest{}
}

/* VoidInvoiceBadRequest describes a response with status code 400, with default header values.

Invalid invoice id supplied
*/
type VoidInvoiceBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the void invoice bad request response
func (o *VoidInvoiceBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this void invoice bad request response has a 2xx status code
func (o *VoidInvoiceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this void invoice bad request response has a 3xx status code
func (o *VoidInvoiceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this void invoice bad request response has a 4xx status code
func (o *VoidInvoiceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this void invoice bad request response has a 5xx status code
func (o *VoidInvoiceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this void invoice bad request response a status code equal to that given
func (o *VoidInvoiceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *VoidInvoiceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoices/{invoiceId}/voidInvoice][%d] voidInvoiceBadRequest ", 400)
}

func (o *VoidInvoiceBadRequest) String() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoices/{invoiceId}/voidInvoice][%d] voidInvoiceBadRequest ", 400)
}

func (o *VoidInvoiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVoidInvoiceNotFound creates a VoidInvoiceNotFound with default headers values
func NewVoidInvoiceNotFound() *VoidInvoiceNotFound {
	return &VoidInvoiceNotFound{}
}

/* VoidInvoiceNotFound describes a response with status code 404, with default header values.

Invoice not found
*/
type VoidInvoiceNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the void invoice not found response
func (o *VoidInvoiceNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this void invoice not found response has a 2xx status code
func (o *VoidInvoiceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this void invoice not found response has a 3xx status code
func (o *VoidInvoiceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this void invoice not found response has a 4xx status code
func (o *VoidInvoiceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this void invoice not found response has a 5xx status code
func (o *VoidInvoiceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this void invoice not found response a status code equal to that given
func (o *VoidInvoiceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *VoidInvoiceNotFound) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoices/{invoiceId}/voidInvoice][%d] voidInvoiceNotFound ", 404)
}

func (o *VoidInvoiceNotFound) String() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoices/{invoiceId}/voidInvoice][%d] voidInvoiceNotFound ", 404)
}

func (o *VoidInvoiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
