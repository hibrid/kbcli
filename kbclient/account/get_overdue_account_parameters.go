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
)

// NewGetOverdueAccountParams creates a new GetOverdueAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOverdueAccountParams() *GetOverdueAccountParams {
	return &GetOverdueAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOverdueAccountParamsWithTimeout creates a new GetOverdueAccountParams object
// with the ability to set a timeout on a request.
func NewGetOverdueAccountParamsWithTimeout(timeout time.Duration) *GetOverdueAccountParams {
	return &GetOverdueAccountParams{
		timeout: timeout,
	}
}

// NewGetOverdueAccountParamsWithContext creates a new GetOverdueAccountParams object
// with the ability to set a context for a request.
func NewGetOverdueAccountParamsWithContext(ctx context.Context) *GetOverdueAccountParams {
	return &GetOverdueAccountParams{
		Context: ctx,
	}
}

// NewGetOverdueAccountParamsWithHTTPClient creates a new GetOverdueAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOverdueAccountParamsWithHTTPClient(client *http.Client) *GetOverdueAccountParams {
	return &GetOverdueAccountParams{
		HTTPClient: client,
	}
}

/* GetOverdueAccountParams contains all the parameters to send to the API endpoint
   for the get overdue account operation.

   Typically these are written to a http.Request.
*/
type GetOverdueAccountParams struct {

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

// WithDefaults hydrates default values in the get overdue account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOverdueAccountParams) WithDefaults() *GetOverdueAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get overdue account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOverdueAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get overdue account params
func (o *GetOverdueAccountParams) WithTimeout(timeout time.Duration) *GetOverdueAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get overdue account params
func (o *GetOverdueAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get overdue account params
func (o *GetOverdueAccountParams) WithContext(ctx context.Context) *GetOverdueAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get overdue account params
func (o *GetOverdueAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get overdue account params
func (o *GetOverdueAccountParams) WithHTTPClient(client *http.Client) *GetOverdueAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get overdue account params
func (o *GetOverdueAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get overdue account params
func (o *GetOverdueAccountParams) WithAccountID(accountID strfmt.UUID) *GetOverdueAccountParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get overdue account params
func (o *GetOverdueAccountParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOverdueAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
