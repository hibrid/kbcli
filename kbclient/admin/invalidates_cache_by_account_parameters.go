// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewInvalidatesCacheByAccountParams creates a new InvalidatesCacheByAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInvalidatesCacheByAccountParams() *InvalidatesCacheByAccountParams {
	return &InvalidatesCacheByAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInvalidatesCacheByAccountParamsWithTimeout creates a new InvalidatesCacheByAccountParams object
// with the ability to set a timeout on a request.
func NewInvalidatesCacheByAccountParamsWithTimeout(timeout time.Duration) *InvalidatesCacheByAccountParams {
	return &InvalidatesCacheByAccountParams{
		timeout: timeout,
	}
}

// NewInvalidatesCacheByAccountParamsWithContext creates a new InvalidatesCacheByAccountParams object
// with the ability to set a context for a request.
func NewInvalidatesCacheByAccountParamsWithContext(ctx context.Context) *InvalidatesCacheByAccountParams {
	return &InvalidatesCacheByAccountParams{
		Context: ctx,
	}
}

// NewInvalidatesCacheByAccountParamsWithHTTPClient creates a new InvalidatesCacheByAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewInvalidatesCacheByAccountParamsWithHTTPClient(client *http.Client) *InvalidatesCacheByAccountParams {
	return &InvalidatesCacheByAccountParams{
		HTTPClient: client,
	}
}

/* InvalidatesCacheByAccountParams contains all the parameters to send to the API endpoint
   for the invalidates cache by account operation.

   Typically these are written to a http.Request.
*/
type InvalidatesCacheByAccountParams struct {

	// AccountID.
	//
	// Format: uuid
	AccountID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the invalidates cache by account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InvalidatesCacheByAccountParams) WithDefaults() *InvalidatesCacheByAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the invalidates cache by account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InvalidatesCacheByAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the invalidates cache by account params
func (o *InvalidatesCacheByAccountParams) WithTimeout(timeout time.Duration) *InvalidatesCacheByAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invalidates cache by account params
func (o *InvalidatesCacheByAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invalidates cache by account params
func (o *InvalidatesCacheByAccountParams) WithContext(ctx context.Context) *InvalidatesCacheByAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invalidates cache by account params
func (o *InvalidatesCacheByAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invalidates cache by account params
func (o *InvalidatesCacheByAccountParams) WithHTTPClient(client *http.Client) *InvalidatesCacheByAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invalidates cache by account params
func (o *InvalidatesCacheByAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the invalidates cache by account params
func (o *InvalidatesCacheByAccountParams) WithAccountID(accountID strfmt.UUID) *InvalidatesCacheByAccountParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the invalidates cache by account params
func (o *InvalidatesCacheByAccountParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *InvalidatesCacheByAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
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
