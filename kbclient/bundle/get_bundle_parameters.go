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
)

// NewGetBundleParams creates a new GetBundleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBundleParams() *GetBundleParams {
	return &GetBundleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBundleParamsWithTimeout creates a new GetBundleParams object
// with the ability to set a timeout on a request.
func NewGetBundleParamsWithTimeout(timeout time.Duration) *GetBundleParams {
	return &GetBundleParams{
		timeout: timeout,
	}
}

// NewGetBundleParamsWithContext creates a new GetBundleParams object
// with the ability to set a context for a request.
func NewGetBundleParamsWithContext(ctx context.Context) *GetBundleParams {
	return &GetBundleParams{
		Context: ctx,
	}
}

// NewGetBundleParamsWithHTTPClient creates a new GetBundleParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBundleParamsWithHTTPClient(client *http.Client) *GetBundleParams {
	return &GetBundleParams{
		HTTPClient: client,
	}
}

/* GetBundleParams contains all the parameters to send to the API endpoint
   for the get bundle operation.

   Typically these are written to a http.Request.
*/
type GetBundleParams struct {

	// Audit.
	//
	// Default: "NONE"
	Audit *string

	// BundleID.
	//
	// Format: uuid
	BundleID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the get bundle params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBundleParams) WithDefaults() *GetBundleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get bundle params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBundleParams) SetDefaults() {
	var (
		auditDefault = string("NONE")
	)

	val := GetBundleParams{
		Audit: &auditDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get bundle params
func (o *GetBundleParams) WithTimeout(timeout time.Duration) *GetBundleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bundle params
func (o *GetBundleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bundle params
func (o *GetBundleParams) WithContext(ctx context.Context) *GetBundleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bundle params
func (o *GetBundleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bundle params
func (o *GetBundleParams) WithHTTPClient(client *http.Client) *GetBundleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bundle params
func (o *GetBundleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the get bundle params
func (o *GetBundleParams) WithAudit(audit *string) *GetBundleParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get bundle params
func (o *GetBundleParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithBundleID adds the bundleID to the get bundle params
func (o *GetBundleParams) WithBundleID(bundleID strfmt.UUID) *GetBundleParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the get bundle params
func (o *GetBundleParams) SetBundleID(bundleID strfmt.UUID) {
	o.BundleID = bundleID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBundleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param bundleId
	if err := r.SetPathParam("bundleId", o.BundleID.String()); err != nil {
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
