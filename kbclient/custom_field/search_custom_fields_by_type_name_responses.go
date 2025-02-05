// Code generated by go-swagger; DO NOT EDIT.

package custom_field

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

// SearchCustomFieldsByTypeNameReader is a Reader for the SearchCustomFieldsByTypeName structure.
type SearchCustomFieldsByTypeNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchCustomFieldsByTypeNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSearchCustomFieldsByTypeNameOK()
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

// NewSearchCustomFieldsByTypeNameOK creates a SearchCustomFieldsByTypeNameOK with default headers values
func NewSearchCustomFieldsByTypeNameOK() *SearchCustomFieldsByTypeNameOK {
	return &SearchCustomFieldsByTypeNameOK{}
}

/* SearchCustomFieldsByTypeNameOK describes a response with status code 200, with default header values.

successful operation
*/
type SearchCustomFieldsByTypeNameOK struct {
	Payload      []*kbmodel.CustomField
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the search custom fields by type name o k response
func (o *SearchCustomFieldsByTypeNameOK) Code() int {
	return 200
}

// IsSuccess returns true when this search custom fields by type name o k response has a 2xx status code
func (o *SearchCustomFieldsByTypeNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search custom fields by type name o k response has a 3xx status code
func (o *SearchCustomFieldsByTypeNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search custom fields by type name o k response has a 4xx status code
func (o *SearchCustomFieldsByTypeNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search custom fields by type name o k response has a 5xx status code
func (o *SearchCustomFieldsByTypeNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search custom fields by type name o k response a status code equal to that given
func (o *SearchCustomFieldsByTypeNameOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchCustomFieldsByTypeNameOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/customFields/search][%d] searchCustomFieldsByTypeNameOK  %+v", 200, o.Payload)
}

func (o *SearchCustomFieldsByTypeNameOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/customFields/search][%d] searchCustomFieldsByTypeNameOK  %+v", 200, o.Payload)
}

func (o *SearchCustomFieldsByTypeNameOK) GetPayload() []*kbmodel.CustomField {
	return o.Payload
}

func (o *SearchCustomFieldsByTypeNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
