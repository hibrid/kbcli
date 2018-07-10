// Code generated by go-swagger; DO NOT EDIT.

package payment_method

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

// NewGetPaymentMethodByKeyParams creates a new GetPaymentMethodByKeyParams object
// with the default values initialized.
func NewGetPaymentMethodByKeyParams() *GetPaymentMethodByKeyParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
		withPluginInfoDefault  = bool(false)
	)
	return &GetPaymentMethodByKeyParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,
		WithPluginInfo:  &withPluginInfoDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentMethodByKeyParamsWithTimeout creates a new GetPaymentMethodByKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentMethodByKeyParamsWithTimeout(timeout time.Duration) *GetPaymentMethodByKeyParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
		withPluginInfoDefault  = bool(false)
	)
	return &GetPaymentMethodByKeyParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,
		WithPluginInfo:  &withPluginInfoDefault,

		timeout: timeout,
	}
}

// NewGetPaymentMethodByKeyParamsWithContext creates a new GetPaymentMethodByKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentMethodByKeyParamsWithContext(ctx context.Context) *GetPaymentMethodByKeyParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
		withPluginInfoDefault  = bool(false)
	)
	return &GetPaymentMethodByKeyParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,
		WithPluginInfo:  &withPluginInfoDefault,

		Context: ctx,
	}
}

// NewGetPaymentMethodByKeyParamsWithHTTPClient creates a new GetPaymentMethodByKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentMethodByKeyParamsWithHTTPClient(client *http.Client) *GetPaymentMethodByKeyParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
		withPluginInfoDefault  = bool(false)
	)
	return &GetPaymentMethodByKeyParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,
		WithPluginInfo:  &withPluginInfoDefault,
		HTTPClient:      client,
	}
}

/*GetPaymentMethodByKeyParams contains all the parameters to send to the API endpoint
for the get payment method by key operation typically these are written to a http.Request
*/
type GetPaymentMethodByKeyParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*Audit*/
	Audit *string
	/*ExternalKey*/
	ExternalKey string
	/*IncludedDeleted*/
	IncludedDeleted *bool
	/*PluginProperty*/
	PluginProperty []string
	/*WithPluginInfo*/
	WithPluginInfo *bool

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get payment method by key params
func (o *GetPaymentMethodByKeyParams) WithTimeout(timeout time.Duration) *GetPaymentMethodByKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payment method by key params
func (o *GetPaymentMethodByKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payment method by key params
func (o *GetPaymentMethodByKeyParams) WithContext(ctx context.Context) *GetPaymentMethodByKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payment method by key params
func (o *GetPaymentMethodByKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payment method by key params
func (o *GetPaymentMethodByKeyParams) WithHTTPClient(client *http.Client) *GetPaymentMethodByKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payment method by key params
func (o *GetPaymentMethodByKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get payment method by key params
func (o *GetPaymentMethodByKeyParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetPaymentMethodByKeyParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get payment method by key params
func (o *GetPaymentMethodByKeyParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get payment method by key params
func (o *GetPaymentMethodByKeyParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetPaymentMethodByKeyParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get payment method by key params
func (o *GetPaymentMethodByKeyParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAudit adds the audit to the get payment method by key params
func (o *GetPaymentMethodByKeyParams) WithAudit(audit *string) *GetPaymentMethodByKeyParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get payment method by key params
func (o *GetPaymentMethodByKeyParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithExternalKey adds the externalKey to the get payment method by key params
func (o *GetPaymentMethodByKeyParams) WithExternalKey(externalKey string) *GetPaymentMethodByKeyParams {
	o.SetExternalKey(externalKey)
	return o
}

// SetExternalKey adds the externalKey to the get payment method by key params
func (o *GetPaymentMethodByKeyParams) SetExternalKey(externalKey string) {
	o.ExternalKey = externalKey
}

// WithIncludedDeleted adds the includedDeleted to the get payment method by key params
func (o *GetPaymentMethodByKeyParams) WithIncludedDeleted(includedDeleted *bool) *GetPaymentMethodByKeyParams {
	o.SetIncludedDeleted(includedDeleted)
	return o
}

// SetIncludedDeleted adds the includedDeleted to the get payment method by key params
func (o *GetPaymentMethodByKeyParams) SetIncludedDeleted(includedDeleted *bool) {
	o.IncludedDeleted = includedDeleted
}

// WithPluginProperty adds the pluginProperty to the get payment method by key params
func (o *GetPaymentMethodByKeyParams) WithPluginProperty(pluginProperty []string) *GetPaymentMethodByKeyParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the get payment method by key params
func (o *GetPaymentMethodByKeyParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithWithPluginInfo adds the withPluginInfo to the get payment method by key params
func (o *GetPaymentMethodByKeyParams) WithWithPluginInfo(withPluginInfo *bool) *GetPaymentMethodByKeyParams {
	o.SetWithPluginInfo(withPluginInfo)
	return o
}

// SetWithPluginInfo adds the withPluginInfo to the get payment method by key params
func (o *GetPaymentMethodByKeyParams) SetWithPluginInfo(withPluginInfo *bool) {
	o.WithPluginInfo = withPluginInfo
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentMethodByKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param externalKey
	qrExternalKey := o.ExternalKey
	qExternalKey := qrExternalKey
	if qExternalKey != "" {
		if err := r.SetQueryParam("externalKey", qExternalKey); err != nil {
			return err
		}
	}

	if o.IncludedDeleted != nil {

		// query param includedDeleted
		var qrIncludedDeleted bool
		if o.IncludedDeleted != nil {
			qrIncludedDeleted = *o.IncludedDeleted
		}
		qIncludedDeleted := swag.FormatBool(qrIncludedDeleted)
		if qIncludedDeleted != "" {
			if err := r.SetQueryParam("includedDeleted", qIncludedDeleted); err != nil {
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
