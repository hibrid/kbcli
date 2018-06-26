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

// CreateTaxItemsReader is a Reader for the CreateTaxItems structure.
type CreateTaxItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTaxItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateTaxItemsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateTaxItemsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateTaxItemsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		errorResult := kbcommon.NewKillbillError(response.Code())
		if err := consumer.Consume(response.Body(), &errorResult); err != nil && err != io.EOF {
			return nil, err
		}
		return nil, errorResult
	}
}

// NewCreateTaxItemsCreated creates a CreateTaxItemsCreated with default headers values
func NewCreateTaxItemsCreated() *CreateTaxItemsCreated {
	return &CreateTaxItemsCreated{}
}

/*CreateTaxItemsCreated handles this case with default header values.

Create tax items successfully
*/
type CreateTaxItemsCreated struct {
	Payload []*kbmodel.InvoiceItem
}

func (o *CreateTaxItemsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/taxes/{accountId}][%d] createTaxItemsCreated  %+v", 201, o.Payload)
}

func (o *CreateTaxItemsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTaxItemsBadRequest creates a CreateTaxItemsBadRequest with default headers values
func NewCreateTaxItemsBadRequest() *CreateTaxItemsBadRequest {
	return &CreateTaxItemsBadRequest{}
}

/*CreateTaxItemsBadRequest handles this case with default header values.

Invalid account id supplied
*/
type CreateTaxItemsBadRequest struct {
}

func (o *CreateTaxItemsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/taxes/{accountId}][%d] createTaxItemsBadRequest ", 400)
}

func (o *CreateTaxItemsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateTaxItemsNotFound creates a CreateTaxItemsNotFound with default headers values
func NewCreateTaxItemsNotFound() *CreateTaxItemsNotFound {
	return &CreateTaxItemsNotFound{}
}

/*CreateTaxItemsNotFound handles this case with default header values.

Account not found
*/
type CreateTaxItemsNotFound struct {
}

func (o *CreateTaxItemsNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/taxes/{accountId}][%d] createTaxItemsNotFound ", 404)
}

func (o *CreateTaxItemsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
