// Code generated by go-swagger; DO NOT EDIT.

package payment_transaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPaymentByTransactionIDParams creates a new GetPaymentByTransactionIDParams object
// with the default values initialized.
func NewGetPaymentByTransactionIDParams() *GetPaymentByTransactionIDParams {
	var (
		auditDefault          = string("NONE")
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &GetPaymentByTransactionIDParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentByTransactionIDParamsWithTimeout creates a new GetPaymentByTransactionIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentByTransactionIDParamsWithTimeout(timeout time.Duration) *GetPaymentByTransactionIDParams {
	var (
		auditDefault          = string("NONE")
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &GetPaymentByTransactionIDParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,

		timeout: timeout,
	}
}

// NewGetPaymentByTransactionIDParamsWithContext creates a new GetPaymentByTransactionIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentByTransactionIDParamsWithContext(ctx context.Context) *GetPaymentByTransactionIDParams {
	var (
		auditDefault          = string("NONE")
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &GetPaymentByTransactionIDParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,

		Context: ctx,
	}
}

// NewGetPaymentByTransactionIDParamsWithHTTPClient creates a new GetPaymentByTransactionIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentByTransactionIDParamsWithHTTPClient(client *http.Client) *GetPaymentByTransactionIDParams {
	var (
		auditDefault          = string("NONE")
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &GetPaymentByTransactionIDParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,
		HTTPClient:     client,
	}
}

/*GetPaymentByTransactionIDParams contains all the parameters to send to the API endpoint
for the get payment by transaction Id operation typically these are written to a http.Request
*/
type GetPaymentByTransactionIDParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*Audit*/
	Audit *string
	/*PluginProperty*/
	PluginProperty []string
	/*TransactionID*/
	TransactionID strfmt.UUID
	/*WithAttempts*/
	WithAttempts *bool
	/*WithPluginInfo*/
	WithPluginInfo *bool

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get payment by transaction Id params
func (o *GetPaymentByTransactionIDParams) WithTimeout(timeout time.Duration) *GetPaymentByTransactionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payment by transaction Id params
func (o *GetPaymentByTransactionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payment by transaction Id params
func (o *GetPaymentByTransactionIDParams) WithContext(ctx context.Context) *GetPaymentByTransactionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payment by transaction Id params
func (o *GetPaymentByTransactionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payment by transaction Id params
func (o *GetPaymentByTransactionIDParams) WithHTTPClient(client *http.Client) *GetPaymentByTransactionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payment by transaction Id params
func (o *GetPaymentByTransactionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get payment by transaction Id params
func (o *GetPaymentByTransactionIDParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetPaymentByTransactionIDParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get payment by transaction Id params
func (o *GetPaymentByTransactionIDParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get payment by transaction Id params
func (o *GetPaymentByTransactionIDParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetPaymentByTransactionIDParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get payment by transaction Id params
func (o *GetPaymentByTransactionIDParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAudit adds the audit to the get payment by transaction Id params
func (o *GetPaymentByTransactionIDParams) WithAudit(audit *string) *GetPaymentByTransactionIDParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get payment by transaction Id params
func (o *GetPaymentByTransactionIDParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithPluginProperty adds the pluginProperty to the get payment by transaction Id params
func (o *GetPaymentByTransactionIDParams) WithPluginProperty(pluginProperty []string) *GetPaymentByTransactionIDParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the get payment by transaction Id params
func (o *GetPaymentByTransactionIDParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithTransactionID adds the transactionID to the get payment by transaction Id params
func (o *GetPaymentByTransactionIDParams) WithTransactionID(transactionID strfmt.UUID) *GetPaymentByTransactionIDParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get payment by transaction Id params
func (o *GetPaymentByTransactionIDParams) SetTransactionID(transactionID strfmt.UUID) {
	o.TransactionID = transactionID
}

// WithWithAttempts adds the withAttempts to the get payment by transaction Id params
func (o *GetPaymentByTransactionIDParams) WithWithAttempts(withAttempts *bool) *GetPaymentByTransactionIDParams {
	o.SetWithAttempts(withAttempts)
	return o
}

// SetWithAttempts adds the withAttempts to the get payment by transaction Id params
func (o *GetPaymentByTransactionIDParams) SetWithAttempts(withAttempts *bool) {
	o.WithAttempts = withAttempts
}

// WithWithPluginInfo adds the withPluginInfo to the get payment by transaction Id params
func (o *GetPaymentByTransactionIDParams) WithWithPluginInfo(withPluginInfo *bool) *GetPaymentByTransactionIDParams {
	o.SetWithPluginInfo(withPluginInfo)
	return o
}

// SetWithPluginInfo adds the withPluginInfo to the get payment by transaction Id params
func (o *GetPaymentByTransactionIDParams) SetWithPluginInfo(withPluginInfo *bool) {
	o.WithPluginInfo = withPluginInfo
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentByTransactionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Killbill-ApiKey
	if err := r.SetHeaderParam("X-Killbill-ApiKey", o.XKillbillAPIKey); err != nil {
		return err
	}

	// header param X-Killbill-ApiSecret
	if err := r.SetHeaderParam("X-Killbill-ApiSecret", o.XKillbillAPISecret); err != nil {
		return err
	}

	if o.Audit != nil {

		// query param audit
		var qrAudit string
		if o.Audit != nil {
			qrAudit = *o.Audit
		}
		qAudit := qrAudit
		if qAudit != "" {
			if err := r.SetQueryParam("audit", qAudit); err != nil {
				return err
			}
		}

	}

	valuesPluginProperty := o.PluginProperty

	joinedPluginProperty := swag.JoinByFormat(valuesPluginProperty, "multi")
	// query array param pluginProperty
	if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
		return err
	}

	// path param transactionId
	if err := r.SetPathParam("transactionId", o.TransactionID.String()); err != nil {
		return err
	}

	if o.WithAttempts != nil {

		// query param withAttempts
		var qrWithAttempts bool
		if o.WithAttempts != nil {
			qrWithAttempts = *o.WithAttempts
		}
		qWithAttempts := swag.FormatBool(qrWithAttempts)
		if qWithAttempts != "" {
			if err := r.SetQueryParam("withAttempts", qWithAttempts); err != nil {
				return err
			}
		}

	}

	if o.WithPluginInfo != nil {

		// query param withPluginInfo
		var qrWithPluginInfo bool
		if o.WithPluginInfo != nil {
			qrWithPluginInfo = *o.WithPluginInfo
		}
		qWithPluginInfo := swag.FormatBool(qrWithPluginInfo)
		if qWithPluginInfo != "" {
			if err := r.SetQueryParam("withPluginInfo", qWithPluginInfo); err != nil {
				return err
			}
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
