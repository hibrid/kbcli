// Code generated by go-swagger; DO NOT EDIT.

package payment_method

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeletePaymentMethodCustomFieldsParams creates a new DeletePaymentMethodCustomFieldsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePaymentMethodCustomFieldsParams() *DeletePaymentMethodCustomFieldsParams {
	return &DeletePaymentMethodCustomFieldsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePaymentMethodCustomFieldsParamsWithTimeout creates a new DeletePaymentMethodCustomFieldsParams object
// with the ability to set a timeout on a request.
func NewDeletePaymentMethodCustomFieldsParamsWithTimeout(timeout time.Duration) *DeletePaymentMethodCustomFieldsParams {
	return &DeletePaymentMethodCustomFieldsParams{
		timeout: timeout,
	}
}

// NewDeletePaymentMethodCustomFieldsParamsWithContext creates a new DeletePaymentMethodCustomFieldsParams object
// with the ability to set a context for a request.
func NewDeletePaymentMethodCustomFieldsParamsWithContext(ctx context.Context) *DeletePaymentMethodCustomFieldsParams {
	return &DeletePaymentMethodCustomFieldsParams{
		Context: ctx,
	}
}

// NewDeletePaymentMethodCustomFieldsParamsWithHTTPClient creates a new DeletePaymentMethodCustomFieldsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePaymentMethodCustomFieldsParamsWithHTTPClient(client *http.Client) *DeletePaymentMethodCustomFieldsParams {
	return &DeletePaymentMethodCustomFieldsParams{
		HTTPClient: client,
	}
}

/* DeletePaymentMethodCustomFieldsParams contains all the parameters to send to the API endpoint
   for the delete payment method custom fields operation.

   Typically these are written to a http.Request.
*/
type DeletePaymentMethodCustomFieldsParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// CustomField.
	CustomField []strfmt.UUID

	// PaymentMethodID.
	//
	// Format: uuid
	PaymentMethodID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the delete payment method custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePaymentMethodCustomFieldsParams) WithDefaults() *DeletePaymentMethodCustomFieldsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete payment method custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePaymentMethodCustomFieldsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete payment method custom fields params
func (o *DeletePaymentMethodCustomFieldsParams) WithTimeout(timeout time.Duration) *DeletePaymentMethodCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete payment method custom fields params
func (o *DeletePaymentMethodCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete payment method custom fields params
func (o *DeletePaymentMethodCustomFieldsParams) WithContext(ctx context.Context) *DeletePaymentMethodCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete payment method custom fields params
func (o *DeletePaymentMethodCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete payment method custom fields params
func (o *DeletePaymentMethodCustomFieldsParams) WithHTTPClient(client *http.Client) *DeletePaymentMethodCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete payment method custom fields params
func (o *DeletePaymentMethodCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the delete payment method custom fields params
func (o *DeletePaymentMethodCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *DeletePaymentMethodCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the delete payment method custom fields params
func (o *DeletePaymentMethodCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the delete payment method custom fields params
func (o *DeletePaymentMethodCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *DeletePaymentMethodCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the delete payment method custom fields params
func (o *DeletePaymentMethodCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the delete payment method custom fields params
func (o *DeletePaymentMethodCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *DeletePaymentMethodCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the delete payment method custom fields params
func (o *DeletePaymentMethodCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithCustomField adds the customField to the delete payment method custom fields params
func (o *DeletePaymentMethodCustomFieldsParams) WithCustomField(customField []strfmt.UUID) *DeletePaymentMethodCustomFieldsParams {
	o.SetCustomField(customField)
	return o
}

// SetCustomField adds the customField to the delete payment method custom fields params
func (o *DeletePaymentMethodCustomFieldsParams) SetCustomField(customField []strfmt.UUID) {
	o.CustomField = customField
}

// WithPaymentMethodID adds the paymentMethodID to the delete payment method custom fields params
func (o *DeletePaymentMethodCustomFieldsParams) WithPaymentMethodID(paymentMethodID strfmt.UUID) *DeletePaymentMethodCustomFieldsParams {
	o.SetPaymentMethodID(paymentMethodID)
	return o
}

// SetPaymentMethodID adds the paymentMethodId to the delete payment method custom fields params
func (o *DeletePaymentMethodCustomFieldsParams) SetPaymentMethodID(paymentMethodID strfmt.UUID) {
	o.PaymentMethodID = paymentMethodID
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePaymentMethodCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XKillbillComment != nil {

		// header param X-Killbill-Comment
		if err := r.SetHeaderParam("X-Killbill-Comment", *o.XKillbillComment); err != nil {
			return err
		}
	}

	// header param X-Killbill-CreatedBy
	if err := r.SetHeaderParam("X-Killbill-CreatedBy", o.XKillbillCreatedBy); err != nil {
		return err
	}

	if o.XKillbillReason != nil {

		// header param X-Killbill-Reason
		if err := r.SetHeaderParam("X-Killbill-Reason", *o.XKillbillReason); err != nil {
			return err
		}
	}

	if o.CustomField != nil {

		// binding items for customField
		joinedCustomField := o.bindParamCustomField(reg)

		// query array param customField
		if err := r.SetQueryParam("customField", joinedCustomField...); err != nil {
			return err
		}
	}

	// path param paymentMethodId
	if err := r.SetPathParam("paymentMethodId", o.PaymentMethodID.String()); err != nil {
		return err
	}

	// header param WithProfilingInfo
	if o.WithProfilingInfo != nil && len(*o.WithProfilingInfo) > 0 {
		if err := r.SetHeaderParam("X-Killbill-Profiling-Req", *o.WithProfilingInfo); err != nil {
			return err
		}
	}

	// header param withStackTrace
	if o.WithStackTrace != nil && *o.WithStackTrace {
		if err := r.SetQueryParam("withStackTrace", "true"); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamDeletePaymentMethodCustomFields binds the parameter customField
func (o *DeletePaymentMethodCustomFieldsParams) bindParamCustomField(formats strfmt.Registry) []string {
	customFieldIR := o.CustomField

	var customFieldIC []string
	for _, customFieldIIR := range customFieldIR { // explode []strfmt.UUID

		customFieldIIV := customFieldIIR.String() // strfmt.UUID as string
		customFieldIC = append(customFieldIC, customFieldIIV)
	}

	// items.CollectionFormat: "multi"
	customFieldIS := swag.JoinByFormat(customFieldIC, "multi")

	return customFieldIS
}
