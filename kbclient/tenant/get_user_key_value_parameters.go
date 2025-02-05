// Code generated by go-swagger; DO NOT EDIT.

package tenant

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

// NewGetUserKeyValueParams creates a new GetUserKeyValueParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserKeyValueParams() *GetUserKeyValueParams {
	return &GetUserKeyValueParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserKeyValueParamsWithTimeout creates a new GetUserKeyValueParams object
// with the ability to set a timeout on a request.
func NewGetUserKeyValueParamsWithTimeout(timeout time.Duration) *GetUserKeyValueParams {
	return &GetUserKeyValueParams{
		timeout: timeout,
	}
}

// NewGetUserKeyValueParamsWithContext creates a new GetUserKeyValueParams object
// with the ability to set a context for a request.
func NewGetUserKeyValueParamsWithContext(ctx context.Context) *GetUserKeyValueParams {
	return &GetUserKeyValueParams{
		Context: ctx,
	}
}

// NewGetUserKeyValueParamsWithHTTPClient creates a new GetUserKeyValueParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserKeyValueParamsWithHTTPClient(client *http.Client) *GetUserKeyValueParams {
	return &GetUserKeyValueParams{
		HTTPClient: client,
	}
}

/* GetUserKeyValueParams contains all the parameters to send to the API endpoint
   for the get user key value operation.

   Typically these are written to a http.Request.
*/
type GetUserKeyValueParams struct {

	// KeyName.
	KeyName string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the get user key value params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserKeyValueParams) WithDefaults() *GetUserKeyValueParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user key value params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserKeyValueParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get user key value params
func (o *GetUserKeyValueParams) WithTimeout(timeout time.Duration) *GetUserKeyValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user key value params
func (o *GetUserKeyValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user key value params
func (o *GetUserKeyValueParams) WithContext(ctx context.Context) *GetUserKeyValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user key value params
func (o *GetUserKeyValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user key value params
func (o *GetUserKeyValueParams) WithHTTPClient(client *http.Client) *GetUserKeyValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user key value params
func (o *GetUserKeyValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeyName adds the keyName to the get user key value params
func (o *GetUserKeyValueParams) WithKeyName(keyName string) *GetUserKeyValueParams {
	o.SetKeyName(keyName)
	return o
}

// SetKeyName adds the keyName to the get user key value params
func (o *GetUserKeyValueParams) SetKeyName(keyName string) {
	o.KeyName = keyName
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserKeyValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param keyName
	if err := r.SetPathParam("keyName", o.KeyName); err != nil {
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
