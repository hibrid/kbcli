// Code generated by go-swagger; DO NOT EDIT.

package nodes_info

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

// NewGetNodesInfoParams creates a new GetNodesInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNodesInfoParams() *GetNodesInfoParams {
	return &GetNodesInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNodesInfoParamsWithTimeout creates a new GetNodesInfoParams object
// with the ability to set a timeout on a request.
func NewGetNodesInfoParamsWithTimeout(timeout time.Duration) *GetNodesInfoParams {
	return &GetNodesInfoParams{
		timeout: timeout,
	}
}

// NewGetNodesInfoParamsWithContext creates a new GetNodesInfoParams object
// with the ability to set a context for a request.
func NewGetNodesInfoParamsWithContext(ctx context.Context) *GetNodesInfoParams {
	return &GetNodesInfoParams{
		Context: ctx,
	}
}

// NewGetNodesInfoParamsWithHTTPClient creates a new GetNodesInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNodesInfoParamsWithHTTPClient(client *http.Client) *GetNodesInfoParams {
	return &GetNodesInfoParams{
		HTTPClient: client,
	}
}

/* GetNodesInfoParams contains all the parameters to send to the API endpoint
   for the get nodes info operation.

   Typically these are written to a http.Request.
*/
type GetNodesInfoParams struct {
	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the get nodes info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNodesInfoParams) WithDefaults() *GetNodesInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nodes info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNodesInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get nodes info params
func (o *GetNodesInfoParams) WithTimeout(timeout time.Duration) *GetNodesInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nodes info params
func (o *GetNodesInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nodes info params
func (o *GetNodesInfoParams) WithContext(ctx context.Context) *GetNodesInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nodes info params
func (o *GetNodesInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nodes info params
func (o *GetNodesInfoParams) WithHTTPClient(client *http.Client) *GetNodesInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nodes info params
func (o *GetNodesInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodesInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
