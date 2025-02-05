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

// NewGetBundlesParams creates a new GetBundlesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBundlesParams() *GetBundlesParams {
	return &GetBundlesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBundlesParamsWithTimeout creates a new GetBundlesParams object
// with the ability to set a timeout on a request.
func NewGetBundlesParamsWithTimeout(timeout time.Duration) *GetBundlesParams {
	return &GetBundlesParams{
		timeout: timeout,
	}
}

// NewGetBundlesParamsWithContext creates a new GetBundlesParams object
// with the ability to set a context for a request.
func NewGetBundlesParamsWithContext(ctx context.Context) *GetBundlesParams {
	return &GetBundlesParams{
		Context: ctx,
	}
}

// NewGetBundlesParamsWithHTTPClient creates a new GetBundlesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBundlesParamsWithHTTPClient(client *http.Client) *GetBundlesParams {
	return &GetBundlesParams{
		HTTPClient: client,
	}
}

/* GetBundlesParams contains all the parameters to send to the API endpoint
   for the get bundles operation.

   Typically these are written to a http.Request.
*/
type GetBundlesParams struct {

	// Audit.
	//
	// Default: "NONE"
	Audit *string

	// Limit.
	//
	// Format: int64
	// Default: 100
	Limit *int64

	// Offset.
	//
	// Format: int64
	Offset *int64

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the get bundles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBundlesParams) WithDefaults() *GetBundlesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get bundles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBundlesParams) SetDefaults() {
	var (
		auditDefault = string("NONE")

		limitDefault = int64(100)

		offsetDefault = int64(0)
	)

	val := GetBundlesParams{
		Audit:  &auditDefault,
		Limit:  &limitDefault,
		Offset: &offsetDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get bundles params
func (o *GetBundlesParams) WithTimeout(timeout time.Duration) *GetBundlesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bundles params
func (o *GetBundlesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bundles params
func (o *GetBundlesParams) WithContext(ctx context.Context) *GetBundlesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bundles params
func (o *GetBundlesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bundles params
func (o *GetBundlesParams) WithHTTPClient(client *http.Client) *GetBundlesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bundles params
func (o *GetBundlesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the get bundles params
func (o *GetBundlesParams) WithAudit(audit *string) *GetBundlesParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get bundles params
func (o *GetBundlesParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithLimit adds the limit to the get bundles params
func (o *GetBundlesParams) WithLimit(limit *int64) *GetBundlesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get bundles params
func (o *GetBundlesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get bundles params
func (o *GetBundlesParams) WithOffset(offset *int64) *GetBundlesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get bundles params
func (o *GetBundlesParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetBundlesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
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
