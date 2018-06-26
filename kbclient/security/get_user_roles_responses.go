// Code generated by go-swagger; DO NOT EDIT.

package security

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

// GetUserRolesReader is a Reader for the GetUserRoles structure.
type GetUserRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUserRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetUserRolesNotFound()
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

// NewGetUserRolesOK creates a GetUserRolesOK with default headers values
func NewGetUserRolesOK() *GetUserRolesOK {
	return &GetUserRolesOK{}
}

/*GetUserRolesOK handles this case with default header values.

successful operation
*/
type GetUserRolesOK struct {
	Payload *kbmodel.UserRoles
}

func (o *GetUserRolesOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/security/users/{username}/roles][%d] getUserRolesOK  %+v", 200, o.Payload)
}

func (o *GetUserRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.UserRoles)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRolesNotFound creates a GetUserRolesNotFound with default headers values
func NewGetUserRolesNotFound() *GetUserRolesNotFound {
	return &GetUserRolesNotFound{}
}

/*GetUserRolesNotFound handles this case with default header values.

The user does not exist or has been inactivated
*/
type GetUserRolesNotFound struct {
}

func (o *GetUserRolesNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/security/users/{username}/roles][%d] getUserRolesNotFound ", 404)
}

func (o *GetUserRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
