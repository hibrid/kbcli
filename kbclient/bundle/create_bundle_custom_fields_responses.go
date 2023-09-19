// Code generated by go-swagger; DO NOT EDIT.

package bundle

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

// CreateBundleCustomFieldsReader is a Reader for the CreateBundleCustomFields structure.
type CreateBundleCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBundleCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateBundleCustomFieldsCreated()
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

// NewCreateBundleCustomFieldsCreated creates a CreateBundleCustomFieldsCreated with default headers values
func NewCreateBundleCustomFieldsCreated() *CreateBundleCustomFieldsCreated {
	return &CreateBundleCustomFieldsCreated{}
}

/* CreateBundleCustomFieldsCreated describes a response with status code 201, with default header values.

Custom field created successfully
*/
type CreateBundleCustomFieldsCreated struct {
	Payload      []*kbmodel.CustomField
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the create bundle custom fields created response
func (o *CreateBundleCustomFieldsCreated) Code() int {
	return 201
}

// IsSuccess returns true when this create bundle custom fields created response has a 2xx status code
func (o *CreateBundleCustomFieldsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create bundle custom fields created response has a 3xx status code
func (o *CreateBundleCustomFieldsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create bundle custom fields created response has a 4xx status code
func (o *CreateBundleCustomFieldsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create bundle custom fields created response has a 5xx status code
func (o *CreateBundleCustomFieldsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create bundle custom fields created response a status code equal to that given
func (o *CreateBundleCustomFieldsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateBundleCustomFieldsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/bundles/{bundleId}/customFields][%d] createBundleCustomFieldsCreated  %+v", 201, o.Payload)
}

func (o *CreateBundleCustomFieldsCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/bundles/{bundleId}/customFields][%d] createBundleCustomFieldsCreated  %+v", 201, o.Payload)
}

func (o *CreateBundleCustomFieldsCreated) GetPayload() []*kbmodel.CustomField {
	return o.Payload
}

func (o *CreateBundleCustomFieldsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBundleCustomFieldsBadRequest creates a CreateBundleCustomFieldsBadRequest with default headers values
func NewCreateBundleCustomFieldsBadRequest() *CreateBundleCustomFieldsBadRequest {
	return &CreateBundleCustomFieldsBadRequest{}
}

/* CreateBundleCustomFieldsBadRequest describes a response with status code 400, with default header values.

Invalid bundle id supplied
*/
type CreateBundleCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the create bundle custom fields bad request response
func (o *CreateBundleCustomFieldsBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this create bundle custom fields bad request response has a 2xx status code
func (o *CreateBundleCustomFieldsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create bundle custom fields bad request response has a 3xx status code
func (o *CreateBundleCustomFieldsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create bundle custom fields bad request response has a 4xx status code
func (o *CreateBundleCustomFieldsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create bundle custom fields bad request response has a 5xx status code
func (o *CreateBundleCustomFieldsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create bundle custom fields bad request response a status code equal to that given
func (o *CreateBundleCustomFieldsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateBundleCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/bundles/{bundleId}/customFields][%d] createBundleCustomFieldsBadRequest ", 400)
}

func (o *CreateBundleCustomFieldsBadRequest) String() string {
	return fmt.Sprintf("[POST /1.0/kb/bundles/{bundleId}/customFields][%d] createBundleCustomFieldsBadRequest ", 400)
}

func (o *CreateBundleCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
