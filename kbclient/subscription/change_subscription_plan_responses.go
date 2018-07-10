// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// ChangeSubscriptionPlanReader is a Reader for the ChangeSubscriptionPlan structure.
type ChangeSubscriptionPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeSubscriptionPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewChangeSubscriptionPlanNoContent()
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

// NewChangeSubscriptionPlanNoContent creates a ChangeSubscriptionPlanNoContent with default headers values
func NewChangeSubscriptionPlanNoContent() *ChangeSubscriptionPlanNoContent {
	return &ChangeSubscriptionPlanNoContent{}
}

/*ChangeSubscriptionPlanNoContent handles this case with default header values.

Successful operation
*/
type ChangeSubscriptionPlanNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChangeSubscriptionPlanNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/subscriptions/{subscriptionId}][%d] changeSubscriptionPlanNoContent ", 204)
}

func (o *ChangeSubscriptionPlanNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeSubscriptionPlanBadRequest creates a ChangeSubscriptionPlanBadRequest with default headers values
func NewChangeSubscriptionPlanBadRequest() *ChangeSubscriptionPlanBadRequest {
	return &ChangeSubscriptionPlanBadRequest{}
}

/*ChangeSubscriptionPlanBadRequest handles this case with default header values.

Invalid subscription id supplied
*/
type ChangeSubscriptionPlanBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChangeSubscriptionPlanBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/subscriptions/{subscriptionId}][%d] changeSubscriptionPlanBadRequest ", 400)
}

func (o *ChangeSubscriptionPlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeSubscriptionPlanNotFound creates a ChangeSubscriptionPlanNotFound with default headers values
func NewChangeSubscriptionPlanNotFound() *ChangeSubscriptionPlanNotFound {
	return &ChangeSubscriptionPlanNotFound{}
}

/*ChangeSubscriptionPlanNotFound handles this case with default header values.

Entitlement not found
*/
type ChangeSubscriptionPlanNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChangeSubscriptionPlanNotFound) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/subscriptions/{subscriptionId}][%d] changeSubscriptionPlanNotFound ", 404)
}

func (o *ChangeSubscriptionPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
