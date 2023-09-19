// Code generated by go-swagger; DO NOT EDIT.

package account

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

// NewProcessPaymentByExternalKeyParams creates a new ProcessPaymentByExternalKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProcessPaymentByExternalKeyParams() *ProcessPaymentByExternalKeyParams {
	return &ProcessPaymentByExternalKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProcessPaymentByExternalKeyParamsWithTimeout creates a new ProcessPaymentByExternalKeyParams object
// with the ability to set a timeout on a request.
func NewProcessPaymentByExternalKeyParamsWithTimeout(timeout time.Duration) *ProcessPaymentByExternalKeyParams {
	return &ProcessPaymentByExternalKeyParams{
		timeout: timeout,
	}
}

// NewProcessPaymentByExternalKeyParamsWithContext creates a new ProcessPaymentByExternalKeyParams object
// with the ability to set a context for a request.
func NewProcessPaymentByExternalKeyParamsWithContext(ctx context.Context) *ProcessPaymentByExternalKeyParams {
	return &ProcessPaymentByExternalKeyParams{
		Context: ctx,
	}
}

// NewProcessPaymentByExternalKeyParamsWithHTTPClient creates a new ProcessPaymentByExternalKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewProcessPaymentByExternalKeyParamsWithHTTPClient(client *http.Client) *ProcessPaymentByExternalKeyParams {
	return &ProcessPaymentByExternalKeyParams{
		HTTPClient: client,
	}
}

/* ProcessPaymentByExternalKeyParams contains all the parameters to send to the API endpoint
   for the process payment by external key operation.

   Typically these are written to a http.Request.
*/
type ProcessPaymentByExternalKeyParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body *kbmodel.PaymentTransaction

	// ControlPluginName.
	ControlPluginName []string

	// ExternalKey.
	ExternalKey string

	// PaymentMethodID.
	//
	// Format: uuid
	PaymentMethodID *strfmt.UUID

	// PluginProperty.
	PluginProperty []string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the process payment by external key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProcessPaymentByExternalKeyParams) WithDefaults() *ProcessPaymentByExternalKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the process payment by external key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProcessPaymentByExternalKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) WithTimeout(timeout time.Duration) *ProcessPaymentByExternalKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) WithContext(ctx context.Context) *ProcessPaymentByExternalKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) WithHTTPClient(client *http.Client) *ProcessPaymentByExternalKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) WithXKillbillComment(xKillbillComment *string) *ProcessPaymentByExternalKeyParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *ProcessPaymentByExternalKeyParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) WithXKillbillReason(xKillbillReason *string) *ProcessPaymentByExternalKeyParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) WithBody(body *kbmodel.PaymentTransaction) *ProcessPaymentByExternalKeyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) SetBody(body *kbmodel.PaymentTransaction) {
	o.Body = body
}

// WithControlPluginName adds the controlPluginName to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) WithControlPluginName(controlPluginName []string) *ProcessPaymentByExternalKeyParams {
	o.SetControlPluginName(controlPluginName)
	return o
}

// SetControlPluginName adds the controlPluginName to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) SetControlPluginName(controlPluginName []string) {
	o.ControlPluginName = controlPluginName
}

// WithExternalKey adds the externalKey to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) WithExternalKey(externalKey string) *ProcessPaymentByExternalKeyParams {
	o.SetExternalKey(externalKey)
	return o
}

// SetExternalKey adds the externalKey to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) SetExternalKey(externalKey string) {
	o.ExternalKey = externalKey
}

// WithPaymentMethodID adds the paymentMethodID to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) WithPaymentMethodID(paymentMethodID *strfmt.UUID) *ProcessPaymentByExternalKeyParams {
	o.SetPaymentMethodID(paymentMethodID)
	return o
}

// SetPaymentMethodID adds the paymentMethodId to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) SetPaymentMethodID(paymentMethodID *strfmt.UUID) {
	o.PaymentMethodID = paymentMethodID
}

// WithPluginProperty adds the pluginProperty to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) WithPluginProperty(pluginProperty []string) *ProcessPaymentByExternalKeyParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the process payment by external key params
func (o *ProcessPaymentByExternalKeyParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *ProcessPaymentByExternalKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param externalKey
	qrExternalKey := o.ExternalKey
	qExternalKey := qrExternalKey
	if qExternalKey != "" {

		if err := r.SetQueryParam("externalKey", qExternalKey); err != nil {
			return err
		}
	}

	if o.PaymentMethodID != nil {

		// query param paymentMethodId
		var qrPaymentMethodID strfmt.UUID

		if o.PaymentMethodID != nil {
			qrPaymentMethodID = *o.PaymentMethodID
		}
		qPaymentMethodID := qrPaymentMethodID.String()
		if qPaymentMethodID != "" {

			if err := r.SetQueryParam("paymentMethodId", qPaymentMethodID); err != nil {
				return err
			}
		}
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

// bindParamProcessPaymentByExternalKey binds the parameter controlPluginName
func (o *ProcessPaymentByExternalKeyParams) bindParamControlPluginName(formats strfmt.Registry) []string {
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

// bindParamProcessPaymentByExternalKey binds the parameter pluginProperty
func (o *ProcessPaymentByExternalKeyParams) bindParamPluginProperty(formats strfmt.Registry) []string {
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
