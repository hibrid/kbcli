// Code generated by go-swagger; DO NOT EDIT.

package invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
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

/*VoidInvoiceNoContent handles this case with default header values.

Successful operation
*/
type VoidInvoiceNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *VoidInvoiceNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoices/{invoiceId}/voidInvoice][%d] voidInvoiceNoContent ", 204)
}

func (o *VoidInvoiceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVoidInvoiceBadRequest creates a VoidInvoiceBadRequest with default headers values
func NewVoidInvoiceBadRequest() *VoidInvoiceBadRequest {
	return &VoidInvoiceBadRequest{}
}

/*VoidInvoiceBadRequest handles this case with default header values.

Invalid invoice id supplied
*/
type VoidInvoiceBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *VoidInvoiceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoices/{invoiceId}/voidInvoice][%d] voidInvoiceBadRequest ", 400)
}

func (o *VoidInvoiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVoidInvoiceNotFound creates a VoidInvoiceNotFound with default headers values
func NewVoidInvoiceNotFound() *VoidInvoiceNotFound {
	return &VoidInvoiceNotFound{}
}

/*VoidInvoiceNotFound handles this case with default header values.

Invoice not found
*/
type VoidInvoiceNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *VoidInvoiceNotFound) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/invoices/{invoiceId}/voidInvoice][%d] voidInvoiceNotFound ", 404)
}

func (o *VoidInvoiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
