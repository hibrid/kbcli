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

// SearchCustomFieldsReader is a Reader for the SearchCustomFields structure.
type SearchCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSearchCustomFieldsOK()
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

// NewSearchCustomFieldsOK creates a SearchCustomFieldsOK with default headers values
func NewSearchCustomFieldsOK() *SearchCustomFieldsOK {
	return &SearchCustomFieldsOK{}
}

/* SearchCustomFieldsOK describes a response with status code 200, with default header values.

successful operation
*/
type SearchCustomFieldsOK struct {
	Payload      []*kbmodel.CustomField
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the search custom fields o k response
func (o *SearchCustomFieldsOK) Code() int {
	return 200
}

// IsSuccess returns true when this search custom fields o k response has a 2xx status code
func (o *SearchCustomFieldsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search custom fields o k response has a 3xx status code
func (o *SearchCustomFieldsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search custom fields o k response has a 4xx status code
func (o *SearchCustomFieldsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search custom fields o k response has a 5xx status code
func (o *SearchCustomFieldsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search custom fields o k response a status code equal to that given
func (o *SearchCustomFieldsOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchCustomFieldsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/customFields/search/{searchKey}][%d] searchCustomFieldsOK  %+v", 200, o.Payload)
}

func (o *SearchCustomFieldsOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/customFields/search/{searchKey}][%d] searchCustomFieldsOK  %+v", 200, o.Payload)
}

func (o *SearchCustomFieldsOK) GetPayload() []*kbmodel.CustomField {
	return o.Payload
}

func (o *SearchCustomFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
