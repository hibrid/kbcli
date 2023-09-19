// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

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

	"github.com/killbill/kbcli/v3/kbmodel"
)

// NewCreateChargebackReversalParams creates a new CreateChargebackReversalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateChargebackReversalParams() *CreateChargebackReversalParams {
	return &CreateChargebackReversalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateChargebackReversalParamsWithTimeout creates a new CreateChargebackReversalParams object
// with the ability to set a timeout on a request.
func NewCreateChargebackReversalParamsWithTimeout(timeout time.Duration) *CreateChargebackReversalParams {
	return &CreateChargebackReversalParams{
		timeout: timeout,
	}
}

// NewCreateChargebackReversalParamsWithContext creates a new CreateChargebackReversalParams object
// with the ability to set a context for a request.
func NewCreateChargebackReversalParamsWithContext(ctx context.Context) *CreateChargebackReversalParams {
	return &CreateChargebackReversalParams{
		Context: ctx,
	}
}

// NewCreateChargebackReversalParamsWithHTTPClient creates a new CreateChargebackReversalParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateChargebackReversalParamsWithHTTPClient(client *http.Client) *CreateChargebackReversalParams {
	return &CreateChargebackReversalParams{
		HTTPClient: client,
	}
}

/* CreateChargebackReversalParams contains all the parameters to send to the API endpoint
   for the create chargeback reversal operation.

   Typically these are written to a http.Request.
*/
type CreateChargebackReversalParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body *kbmodel.InvoicePaymentTransaction

	// PaymentID.
	//
	// Format: uuid
	PaymentID strfmt.UUID

	// PluginProperty.
	PluginProperty []string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the create chargeback reversal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateChargebackReversalParams) WithDefaults() *CreateChargebackReversalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create chargeback reversal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateChargebackReversalParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create chargeback reversal params
func (o *CreateChargebackReversalParams) WithTimeout(timeout time.Duration) *CreateChargebackReversalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create chargeback reversal params
func (o *CreateChargebackReversalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create chargeback reversal params
func (o *CreateChargebackReversalParams) WithContext(ctx context.Context) *CreateChargebackReversalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create chargeback reversal params
func (o *CreateChargebackReversalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create chargeback reversal params
func (o *CreateChargebackReversalParams) WithHTTPClient(client *http.Client) *CreateChargebackReversalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create chargeback reversal params
func (o *CreateChargebackReversalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the create chargeback reversal params
func (o *CreateChargebackReversalParams) WithXKillbillComment(xKillbillComment *string) *CreateChargebackReversalParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create chargeback reversal params
func (o *CreateChargebackReversalParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create chargeback reversal params
func (o *CreateChargebackReversalParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateChargebackReversalParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create chargeback reversal params
func (o *CreateChargebackReversalParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create chargeback reversal params
func (o *CreateChargebackReversalParams) WithXKillbillReason(xKillbillReason *string) *CreateChargebackReversalParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create chargeback reversal params
func (o *CreateChargebackReversalParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the create chargeback reversal params
func (o *CreateChargebackReversalParams) WithBody(body *kbmodel.InvoicePaymentTransaction) *CreateChargebackReversalParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create chargeback reversal params
func (o *CreateChargebackReversalParams) SetBody(body *kbmodel.InvoicePaymentTransaction) {
	o.Body = body
}

// WithPaymentID adds the paymentID to the create chargeback reversal params
func (o *CreateChargebackReversalParams) WithPaymentID(paymentID strfmt.UUID) *CreateChargebackReversalParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the create chargeback reversal params
func (o *CreateChargebackReversalParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WithPluginProperty adds the pluginProperty to the create chargeback reversal params
func (o *CreateChargebackReversalParams) WithPluginProperty(pluginProperty []string) *CreateChargebackReversalParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the create chargeback reversal params
func (o *CreateChargebackReversalParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *CreateChargebackReversalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID.String()); err != nil {
		return err
	}

	if o.PluginProperty != nil {

		// binding items for pluginProperty
		joinedPluginProperty := o.bindParamPluginProperty(reg)

		// query array param pluginProperty
		if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
			return err
		}
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

// bindParamCreateChargebackReversal binds the parameter pluginProperty
func (o *CreateChargebackReversalParams) bindParamPluginProperty(formats strfmt.Registry) []string {
	pluginPropertyIR := o.PluginProperty

	var pluginPropertyIC []string
	for _, pluginPropertyIIR := range pluginPropertyIR { // explode []string

		pluginPropertyIIV := pluginPropertyIIR // string as string
		pluginPropertyIC = append(pluginPropertyIC, pluginPropertyIIV)
	}

	// items.CollectionFormat: "multi"
	pluginPropertyIS := swag.JoinByFormat(pluginPropertyIC, "multi")

	return pluginPropertyIS
}
