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

// NewGetAccountAuditLogsWithHistoryParams creates a new GetAccountAuditLogsWithHistoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccountAuditLogsWithHistoryParams() *GetAccountAuditLogsWithHistoryParams {
	return &GetAccountAuditLogsWithHistoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountAuditLogsWithHistoryParamsWithTimeout creates a new GetAccountAuditLogsWithHistoryParams object
// with the ability to set a timeout on a request.
func NewGetAccountAuditLogsWithHistoryParamsWithTimeout(timeout time.Duration) *GetAccountAuditLogsWithHistoryParams {
	return &GetAccountAuditLogsWithHistoryParams{
		timeout: timeout,
	}
}

// NewGetAccountAuditLogsWithHistoryParamsWithContext creates a new GetAccountAuditLogsWithHistoryParams object
// with the ability to set a context for a request.
func NewGetAccountAuditLogsWithHistoryParamsWithContext(ctx context.Context) *GetAccountAuditLogsWithHistoryParams {
	return &GetAccountAuditLogsWithHistoryParams{
		Context: ctx,
	}
}

// NewGetAccountAuditLogsWithHistoryParamsWithHTTPClient creates a new GetAccountAuditLogsWithHistoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccountAuditLogsWithHistoryParamsWithHTTPClient(client *http.Client) *GetAccountAuditLogsWithHistoryParams {
	return &GetAccountAuditLogsWithHistoryParams{
		HTTPClient: client,
	}
}

/* GetAccountAuditLogsWithHistoryParams contains all the parameters to send to the API endpoint
   for the get account audit logs with history operation.

   Typically these are written to a http.Request.
*/
type GetAccountAuditLogsWithHistoryParams struct {

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

// WithDefaults hydrates default values in the get account audit logs with history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountAuditLogsWithHistoryParams) WithDefaults() *GetAccountAuditLogsWithHistoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get account audit logs with history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountAuditLogsWithHistoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get account audit logs with history params
func (o *GetAccountAuditLogsWithHistoryParams) WithTimeout(timeout time.Duration) *GetAccountAuditLogsWithHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account audit logs with history params
func (o *GetAccountAuditLogsWithHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account audit logs with history params
func (o *GetAccountAuditLogsWithHistoryParams) WithContext(ctx context.Context) *GetAccountAuditLogsWithHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account audit logs with history params
func (o *GetAccountAuditLogsWithHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account audit logs with history params
func (o *GetAccountAuditLogsWithHistoryParams) WithHTTPClient(client *http.Client) *GetAccountAuditLogsWithHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account audit logs with history params
func (o *GetAccountAuditLogsWithHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get account audit logs with history params
func (o *GetAccountAuditLogsWithHistoryParams) WithAccountID(accountID strfmt.UUID) *GetAccountAuditLogsWithHistoryParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get account audit logs with history params
func (o *GetAccountAuditLogsWithHistoryParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountAuditLogsWithHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
