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

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// AdjustInvoiceItemReader is a Reader for the AdjustInvoiceItem structure.
type AdjustInvoiceItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdjustInvoiceItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewAdjustInvoiceItemCreated()
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

// NewAdjustInvoiceItemCreated creates a AdjustInvoiceItemCreated with default headers values
func NewAdjustInvoiceItemCreated() *AdjustInvoiceItemCreated {
	return &AdjustInvoiceItemCreated{}
}

/*AdjustInvoiceItemCreated handles this case with default header values.

Created adjustment Successfully
*/
type AdjustInvoiceItemCreated struct {
	Payload *kbmodel.Invoice

	HttpResponse runtime.ClientResponse
}

func (o *AdjustInvoiceItemCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/{invoiceId}][%d] adjustInvoiceItemCreated  %+v", 201, o.Payload)
}

func (o *AdjustInvoiceItemCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Invoice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdjustInvoiceItemBadRequest creates a AdjustInvoiceItemBadRequest with default headers values
func NewAdjustInvoiceItemBadRequest() *AdjustInvoiceItemBadRequest {
	return &AdjustInvoiceItemBadRequest{}
}

/*AdjustInvoiceItemBadRequest handles this case with default header values.

Invalid account id, invoice id or invoice item id supplied
*/
type AdjustInvoiceItemBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *AdjustInvoiceItemBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/{invoiceId}][%d] adjustInvoiceItemBadRequest ", 400)
}

func (o *AdjustInvoiceItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAdjustInvoiceItemNotFound creates a AdjustInvoiceItemNotFound with default headers values
func NewAdjustInvoiceItemNotFound() *AdjustInvoiceItemNotFound {
	return &AdjustInvoiceItemNotFound{}
}

/*AdjustInvoiceItemNotFound handles this case with default header values.

Invoice not found
*/
type AdjustInvoiceItemNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *AdjustInvoiceItemNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/{invoiceId}][%d] adjustInvoiceItemNotFound ", 404)
}

func (o *AdjustInvoiceItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
