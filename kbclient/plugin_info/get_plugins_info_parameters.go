// Code generated by go-swagger; DO NOT EDIT.

package plugin_info

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

// NewGetPluginsInfoParams creates a new GetPluginsInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPluginsInfoParams() *GetPluginsInfoParams {
	return &GetPluginsInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPluginsInfoParamsWithTimeout creates a new GetPluginsInfoParams object
// with the ability to set a timeout on a request.
func NewGetPluginsInfoParamsWithTimeout(timeout time.Duration) *GetPluginsInfoParams {
	return &GetPluginsInfoParams{
		timeout: timeout,
	}
}

// NewGetPluginsInfoParamsWithContext creates a new GetPluginsInfoParams object
// with the ability to set a context for a request.
func NewGetPluginsInfoParamsWithContext(ctx context.Context) *GetPluginsInfoParams {
	return &GetPluginsInfoParams{
		Context: ctx,
	}
}

// NewGetPluginsInfoParamsWithHTTPClient creates a new GetPluginsInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPluginsInfoParamsWithHTTPClient(client *http.Client) *GetPluginsInfoParams {
	return &GetPluginsInfoParams{
		HTTPClient: client,
	}
}

/* GetPluginsInfoParams contains all the parameters to send to the API endpoint
   for the get plugins info operation.

   Typically these are written to a http.Request.
*/
type GetPluginsInfoParams struct {
	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the get plugins info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPluginsInfoParams) WithDefaults() *GetPluginsInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get plugins info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPluginsInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get plugins info params
func (o *GetPluginsInfoParams) WithTimeout(timeout time.Duration) *GetPluginsInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get plugins info params
func (o *GetPluginsInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get plugins info params
func (o *GetPluginsInfoParams) WithContext(ctx context.Context) *GetPluginsInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get plugins info params
func (o *GetPluginsInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get plugins info params
func (o *GetPluginsInfoParams) WithHTTPClient(client *http.Client) *GetPluginsInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get plugins info params
func (o *GetPluginsInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPluginsInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
