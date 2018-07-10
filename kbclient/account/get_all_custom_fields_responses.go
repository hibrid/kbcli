// Code generated by go-swagger; DO NOT EDIT.

package account

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

// GetAllCustomFieldsReader is a Reader for the GetAllCustomFields structure.
type GetAllCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllCustomFieldsOK()
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

// NewGetAllCustomFieldsOK creates a GetAllCustomFieldsOK with default headers values
func NewGetAllCustomFieldsOK() *GetAllCustomFieldsOK {
	return &GetAllCustomFieldsOK{}
}

/*GetAllCustomFieldsOK handles this case with default header values.

successful operation
*/
type GetAllCustomFieldsOK struct {
	Payload []*kbmodel.CustomField

	HttpResponse runtime.ClientResponse
}

func (o *GetAllCustomFieldsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/allCustomFields][%d] getAllCustomFieldsOK  %+v", 200, o.Payload)
}

func (o *GetAllCustomFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllCustomFieldsBadRequest creates a GetAllCustomFieldsBadRequest with default headers values
func NewGetAllCustomFieldsBadRequest() *GetAllCustomFieldsBadRequest {
	return &GetAllCustomFieldsBadRequest{}
}

/*GetAllCustomFieldsBadRequest handles this case with default header values.

Invalid account id supplied
*/
type GetAllCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetAllCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/allCustomFields][%d] getAllCustomFieldsBadRequest ", 400)
}

func (o *GetAllCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllCustomFieldsNotFound creates a GetAllCustomFieldsNotFound with default headers values
func NewGetAllCustomFieldsNotFound() *GetAllCustomFieldsNotFound {
	return &GetAllCustomFieldsNotFound{}
}

/*GetAllCustomFieldsNotFound handles this case with default header values.

Account not found
*/
type GetAllCustomFieldsNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetAllCustomFieldsNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/allCustomFields][%d] getAllCustomFieldsNotFound ", 404)
}

func (o *GetAllCustomFieldsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
