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

// GetPriceListForSubscriptionAndDateReader is a Reader for the GetPriceListForSubscriptionAndDate structure.
type GetPriceListForSubscriptionAndDateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPriceListForSubscriptionAndDateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPriceListForSubscriptionAndDateOK()
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

// NewGetPriceListForSubscriptionAndDateOK creates a GetPriceListForSubscriptionAndDateOK with default headers values
func NewGetPriceListForSubscriptionAndDateOK() *GetPriceListForSubscriptionAndDateOK {
	return &GetPriceListForSubscriptionAndDateOK{}
}

/* GetPriceListForSubscriptionAndDateOK describes a response with status code 200, with default header values.

successful operation
*/
type GetPriceListForSubscriptionAndDateOK struct {
	Payload      *kbmodel.PriceList
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get price list for subscription and date o k response
func (o *GetPriceListForSubscriptionAndDateOK) Code() int {
	return 200
}

// IsSuccess returns true when this get price list for subscription and date o k response has a 2xx status code
func (o *GetPriceListForSubscriptionAndDateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get price list for subscription and date o k response has a 3xx status code
func (o *GetPriceListForSubscriptionAndDateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get price list for subscription and date o k response has a 4xx status code
func (o *GetPriceListForSubscriptionAndDateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get price list for subscription and date o k response has a 5xx status code
func (o *GetPriceListForSubscriptionAndDateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get price list for subscription and date o k response a status code equal to that given
func (o *GetPriceListForSubscriptionAndDateOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPriceListForSubscriptionAndDateOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/catalog/priceList][%d] getPriceListForSubscriptionAndDateOK  %+v", 200, o.Payload)
}

func (o *GetPriceListForSubscriptionAndDateOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/catalog/priceList][%d] getPriceListForSubscriptionAndDateOK  %+v", 200, o.Payload)
}

func (o *GetPriceListForSubscriptionAndDateOK) GetPayload() *kbmodel.PriceList {
	return o.Payload
}

func (o *GetPriceListForSubscriptionAndDateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.PriceList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
