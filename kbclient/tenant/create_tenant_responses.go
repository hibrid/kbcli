// Code generated by go-swagger; DO NOT EDIT.

package tenant

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

// CreateTenantReader is a Reader for the CreateTenant structure.
type CreateTenantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTenantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateTenantCreated()
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

// NewCreateTenantCreated creates a CreateTenantCreated with default headers values
func NewCreateTenantCreated() *CreateTenantCreated {
	return &CreateTenantCreated{}
}

/*CreateTenantCreated handles this case with default header values.

Tenant created successfully
*/
type CreateTenantCreated struct {
	Payload *kbmodel.Tenant

	HttpResponse runtime.ClientResponse
}

func (o *CreateTenantCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants][%d] createTenantCreated  %+v", 201, o.Payload)
}

func (o *CreateTenantCreated) GetPayload() *kbmodel.Tenant {
	return o.Payload
}

func (o *CreateTenantCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTenantConflict creates a CreateTenantConflict with default headers values
func NewCreateTenantConflict() *CreateTenantConflict {
	return &CreateTenantConflict{}
}

/*CreateTenantConflict handles this case with default header values.

Tenant already exists
*/
type CreateTenantConflict struct {
	HttpResponse runtime.ClientResponse
}

func (o *CreateTenantConflict) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants][%d] createTenantConflict ", 409)
}

func (o *CreateTenantConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
