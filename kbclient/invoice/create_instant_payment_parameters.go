// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

// NewCreateInstantPaymentParams creates a new CreateInstantPaymentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateInstantPaymentParams() *CreateInstantPaymentParams {
	return &CreateInstantPaymentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInstantPaymentParamsWithTimeout creates a new CreateInstantPaymentParams object
// with the ability to set a timeout on a request.
func NewCreateInstantPaymentParamsWithTimeout(timeout time.Duration) *CreateInstantPaymentParams {
	return &CreateInstantPaymentParams{
		timeout: timeout,
	}
}

// NewCreateInstantPaymentParamsWithContext creates a new CreateInstantPaymentParams object
// with the ability to set a context for a request.
func NewCreateInstantPaymentParamsWithContext(ctx context.Context) *CreateInstantPaymentParams {
	return &CreateInstantPaymentParams{
		Context: ctx,
	}
}

// NewCreateInstantPaymentParamsWithHTTPClient creates a new CreateInstantPaymentParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateInstantPaymentParamsWithHTTPClient(client *http.Client) *CreateInstantPaymentParams {
	return &CreateInstantPaymentParams{
		HTTPClient: client,
	}
}

/* CreateInstantPaymentParams contains all the parameters to send to the API endpoint
   for the create instant payment operation.

   Typically these are written to a http.Request.
*/
type CreateInstantPaymentParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body *kbmodel.InvoicePayment

	// ControlPluginName.
	ControlPluginName []string

	// ExternalPayment.
	ExternalPayment *bool

	// InvoiceID.
	//
	// Format: uuid
	InvoiceID strfmt.UUID

	// PluginProperty.
	PluginProperty []string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the create instant payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInstantPaymentParams) WithDefaults() *CreateInstantPaymentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create instant payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInstantPaymentParams) SetDefaults() {
	var (
		externalPaymentDefault = bool(false)
	)

	val := CreateInstantPaymentParams{
		ExternalPayment: &externalPaymentDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create instant payment params
func (o *CreateInstantPaymentParams) WithTimeout(timeout time.Duration) *CreateInstantPaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create instant payment params
func (o *CreateInstantPaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create instant payment params
func (o *CreateInstantPaymentParams) WithContext(ctx context.Context) *CreateInstantPaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create instant payment params
func (o *CreateInstantPaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create instant payment params
func (o *CreateInstantPaymentParams) WithHTTPClient(client *http.Client) *CreateInstantPaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create instant payment params
func (o *CreateInstantPaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the create instant payment params
func (o *CreateInstantPaymentParams) WithXKillbillComment(xKillbillComment *string) *CreateInstantPaymentParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create instant payment params
func (o *CreateInstantPaymentParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create instant payment params
func (o *CreateInstantPaymentParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateInstantPaymentParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create instant payment params
func (o *CreateInstantPaymentParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create instant payment params
func (o *CreateInstantPaymentParams) WithXKillbillReason(xKillbillReason *string) *CreateInstantPaymentParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create instant payment params
func (o *CreateInstantPaymentParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the create instant payment params
func (o *CreateInstantPaymentParams) WithBody(body *kbmodel.InvoicePayment) *CreateInstantPaymentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create instant payment params
func (o *CreateInstantPaymentParams) SetBody(body *kbmodel.InvoicePayment) {
	o.Body = body
}

// WithControlPluginName adds the controlPluginName to the create instant payment params
func (o *CreateInstantPaymentParams) WithControlPluginName(controlPluginName []string) *CreateInstantPaymentParams {
	o.SetControlPluginName(controlPluginName)
	return o
}

// SetControlPluginName adds the controlPluginName to the create instant payment params
func (o *CreateInstantPaymentParams) SetControlPluginName(controlPluginName []string) {
	o.ControlPluginName = controlPluginName
}

// WithExternalPayment adds the externalPayment to the create instant payment params
func (o *CreateInstantPaymentParams) WithExternalPayment(externalPayment *bool) *CreateInstantPaymentParams {
	o.SetExternalPayment(externalPayment)
	return o
}

// SetExternalPayment adds the externalPayment to the create instant payment params
func (o *CreateInstantPaymentParams) SetExternalPayment(externalPayment *bool) {
	o.ExternalPayment = externalPayment
}

// WithInvoiceID adds the invoiceID to the create instant payment params
func (o *CreateInstantPaymentParams) WithInvoiceID(invoiceID strfmt.UUID) *CreateInstantPaymentParams {
	o.SetInvoiceID(invoiceID)
	return o
}

// SetInvoiceID adds the invoiceId to the create instant payment params
func (o *CreateInstantPaymentParams) SetInvoiceID(invoiceID strfmt.UUID) {
	o.InvoiceID = invoiceID
}

// WithPluginProperty adds the pluginProperty to the create instant payment params
func (o *CreateInstantPaymentParams) WithPluginProperty(pluginProperty []string) *CreateInstantPaymentParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the create instant payment params
func (o *CreateInstantPaymentParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInstantPaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ControlPluginName != nil {

		// binding items for controlPluginName
		joinedControlPluginName := o.bindParamControlPluginName(reg)

		// query array param controlPluginName
		if err := r.SetQueryParam("controlPluginName", joinedControlPluginName...); err != nil {
			return err
		}
	}

	if o.ExternalPayment != nil {

		// query param externalPayment
		var qrExternalPayment bool

		if o.ExternalPayment != nil {
			qrExternalPayment = *o.ExternalPayment
		}
		qExternalPayment := swag.FormatBool(qrExternalPayment)
		if qExternalPayment != "" {

			if err := r.SetQueryParam("externalPayment", qExternalPayment); err != nil {
				return err
			}
		}
	}

	// path param invoiceId
	if err := r.SetPathParam("invoiceId", o.InvoiceID.String()); err != nil {
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

// bindParamCreateInstantPayment binds the parameter controlPluginName
func (o *CreateInstantPaymentParams) bindParamControlPluginName(formats strfmt.Registry) []string {
	controlPluginNameIR := o.ControlPluginName

	var controlPluginNameIC []string
	for _, controlPluginNameIIR := range controlPluginNameIR { // explode []string

		controlPluginNameIIV := controlPluginNameIIR // string as string
		controlPluginNameIC = append(controlPluginNameIC, controlPluginNameIIV)
	}

	// items.CollectionFormat: "multi"
	controlPluginNameIS := swag.JoinByFormat(controlPluginNameIC, "multi")

	return controlPluginNameIS
}

// bindParamCreateInstantPayment binds the parameter pluginProperty
func (o *CreateInstantPaymentParams) bindParamPluginProperty(formats strfmt.Registry) []string {
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
