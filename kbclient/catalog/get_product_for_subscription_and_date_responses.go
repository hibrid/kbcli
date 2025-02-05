// Code generated by go-swagger; DO NOT EDIT.

package catalog

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

// GetProductForSubscriptionAndDateReader is a Reader for the GetProductForSubscriptionAndDate structure.
type GetProductForSubscriptionAndDateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductForSubscriptionAndDateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProductForSubscriptionAndDateOK()
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

// NewGetProductForSubscriptionAndDateOK creates a GetProductForSubscriptionAndDateOK with default headers values
func NewGetProductForSubscriptionAndDateOK() *GetProductForSubscriptionAndDateOK {
	return &GetProductForSubscriptionAndDateOK{}
}

/* GetProductForSubscriptionAndDateOK describes a response with status code 200, with default header values.

successful operation
*/
type GetProductForSubscriptionAndDateOK struct {
	Payload      *kbmodel.Product
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get product for subscription and date o k response
func (o *GetProductForSubscriptionAndDateOK) Code() int {
	return 200
}

// IsSuccess returns true when this get product for subscription and date o k response has a 2xx status code
func (o *GetProductForSubscriptionAndDateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get product for subscription and date o k response has a 3xx status code
func (o *GetProductForSubscriptionAndDateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get product for subscription and date o k response has a 4xx status code
func (o *GetProductForSubscriptionAndDateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get product for subscription and date o k response has a 5xx status code
func (o *GetProductForSubscriptionAndDateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get product for subscription and date o k response a status code equal to that given
func (o *GetProductForSubscriptionAndDateOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetProductForSubscriptionAndDateOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/catalog/product][%d] getProductForSubscriptionAndDateOK  %+v", 200, o.Payload)
}

func (o *GetProductForSubscriptionAndDateOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/catalog/product][%d] getProductForSubscriptionAndDateOK  %+v", 200, o.Payload)
}

func (o *GetProductForSubscriptionAndDateOK) GetPayload() *kbmodel.Product {
	return o.Payload
}

func (o *GetProductForSubscriptionAndDateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
