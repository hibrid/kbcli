// Code generated by go-swagger; DO NOT EDIT.

package bundle

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

// NewPauseBundleParams creates a new PauseBundleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPauseBundleParams() *PauseBundleParams {
	return &PauseBundleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPauseBundleParamsWithTimeout creates a new PauseBundleParams object
// with the ability to set a timeout on a request.
func NewPauseBundleParamsWithTimeout(timeout time.Duration) *PauseBundleParams {
	return &PauseBundleParams{
		timeout: timeout,
	}
}

// NewPauseBundleParamsWithContext creates a new PauseBundleParams object
// with the ability to set a context for a request.
func NewPauseBundleParamsWithContext(ctx context.Context) *PauseBundleParams {
	return &PauseBundleParams{
		Context: ctx,
	}
}

// NewPauseBundleParamsWithHTTPClient creates a new PauseBundleParams object
// with the ability to set a custom HTTPClient for a request.
func NewPauseBundleParamsWithHTTPClient(client *http.Client) *PauseBundleParams {
	return &PauseBundleParams{
		HTTPClient: client,
	}
}

/* PauseBundleParams contains all the parameters to send to the API endpoint
   for the pause bundle operation.

   Typically these are written to a http.Request.
*/
type PauseBundleParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// BundleID.
	//
	// Format: uuid
	BundleID strfmt.UUID

	// PluginProperty.
	PluginProperty []string

	// RequestedDate.
	//
	// Format: date
	RequestedDate *strfmt.Date

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the pause bundle params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PauseBundleParams) WithDefaults() *PauseBundleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pause bundle params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PauseBundleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pause bundle params
func (o *PauseBundleParams) WithTimeout(timeout time.Duration) *PauseBundleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pause bundle params
func (o *PauseBundleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pause bundle params
func (o *PauseBundleParams) WithContext(ctx context.Context) *PauseBundleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pause bundle params
func (o *PauseBundleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pause bundle params
func (o *PauseBundleParams) WithHTTPClient(client *http.Client) *PauseBundleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pause bundle params
func (o *PauseBundleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the pause bundle params
func (o *PauseBundleParams) WithXKillbillComment(xKillbillComment *string) *PauseBundleParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the pause bundle params
func (o *PauseBundleParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the pause bundle params
func (o *PauseBundleParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *PauseBundleParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the pause bundle params
func (o *PauseBundleParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the pause bundle params
func (o *PauseBundleParams) WithXKillbillReason(xKillbillReason *string) *PauseBundleParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the pause bundle params
func (o *PauseBundleParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBundleID adds the bundleID to the pause bundle params
func (o *PauseBundleParams) WithBundleID(bundleID strfmt.UUID) *PauseBundleParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the pause bundle params
func (o *PauseBundleParams) SetBundleID(bundleID strfmt.UUID) {
	o.BundleID = bundleID
}

// WithPluginProperty adds the pluginProperty to the pause bundle params
func (o *PauseBundleParams) WithPluginProperty(pluginProperty []string) *PauseBundleParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the pause bundle params
func (o *PauseBundleParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithRequestedDate adds the requestedDate to the pause bundle params
func (o *PauseBundleParams) WithRequestedDate(requestedDate *strfmt.Date) *PauseBundleParams {
	o.SetRequestedDate(requestedDate)
	return o
}

// SetRequestedDate adds the requestedDate to the pause bundle params
func (o *PauseBundleParams) SetRequestedDate(requestedDate *strfmt.Date) {
	o.RequestedDate = requestedDate
}

// WriteToRequest writes these params to a swagger request
func (o *PauseBundleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param bundleId
	if err := r.SetPathParam("bundleId", o.BundleID.String()); err != nil {
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

	if o.RequestedDate != nil {

		// query param requestedDate
		var qrRequestedDate strfmt.Date

		if o.RequestedDate != nil {
			qrRequestedDate = *o.RequestedDate
		}
		qRequestedDate := qrRequestedDate.String()
		if qRequestedDate != "" {

			if err := r.SetQueryParam("requestedDate", qRequestedDate); err != nil {
				return err
			}
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

// bindParamPauseBundle binds the parameter pluginProperty
func (o *PauseBundleParams) bindParamPluginProperty(formats strfmt.Registry) []string {
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
