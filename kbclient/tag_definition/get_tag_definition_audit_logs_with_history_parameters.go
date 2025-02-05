// Code generated by go-swagger; DO NOT EDIT.

package tag_definition

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

// NewGetTagDefinitionAuditLogsWithHistoryParams creates a new GetTagDefinitionAuditLogsWithHistoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTagDefinitionAuditLogsWithHistoryParams() *GetTagDefinitionAuditLogsWithHistoryParams {
	return &GetTagDefinitionAuditLogsWithHistoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTagDefinitionAuditLogsWithHistoryParamsWithTimeout creates a new GetTagDefinitionAuditLogsWithHistoryParams object
// with the ability to set a timeout on a request.
func NewGetTagDefinitionAuditLogsWithHistoryParamsWithTimeout(timeout time.Duration) *GetTagDefinitionAuditLogsWithHistoryParams {
	return &GetTagDefinitionAuditLogsWithHistoryParams{
		timeout: timeout,
	}
}

// NewGetTagDefinitionAuditLogsWithHistoryParamsWithContext creates a new GetTagDefinitionAuditLogsWithHistoryParams object
// with the ability to set a context for a request.
func NewGetTagDefinitionAuditLogsWithHistoryParamsWithContext(ctx context.Context) *GetTagDefinitionAuditLogsWithHistoryParams {
	return &GetTagDefinitionAuditLogsWithHistoryParams{
		Context: ctx,
	}
}

// NewGetTagDefinitionAuditLogsWithHistoryParamsWithHTTPClient creates a new GetTagDefinitionAuditLogsWithHistoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTagDefinitionAuditLogsWithHistoryParamsWithHTTPClient(client *http.Client) *GetTagDefinitionAuditLogsWithHistoryParams {
	return &GetTagDefinitionAuditLogsWithHistoryParams{
		HTTPClient: client,
	}
}

/* GetTagDefinitionAuditLogsWithHistoryParams contains all the parameters to send to the API endpoint
   for the get tag definition audit logs with history operation.

   Typically these are written to a http.Request.
*/
type GetTagDefinitionAuditLogsWithHistoryParams struct {

	// TagDefinitionID.
	//
	// Format: uuid
	TagDefinitionID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the get tag definition audit logs with history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTagDefinitionAuditLogsWithHistoryParams) WithDefaults() *GetTagDefinitionAuditLogsWithHistoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get tag definition audit logs with history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTagDefinitionAuditLogsWithHistoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get tag definition audit logs with history params
func (o *GetTagDefinitionAuditLogsWithHistoryParams) WithTimeout(timeout time.Duration) *GetTagDefinitionAuditLogsWithHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tag definition audit logs with history params
func (o *GetTagDefinitionAuditLogsWithHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tag definition audit logs with history params
func (o *GetTagDefinitionAuditLogsWithHistoryParams) WithContext(ctx context.Context) *GetTagDefinitionAuditLogsWithHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tag definition audit logs with history params
func (o *GetTagDefinitionAuditLogsWithHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tag definition audit logs with history params
func (o *GetTagDefinitionAuditLogsWithHistoryParams) WithHTTPClient(client *http.Client) *GetTagDefinitionAuditLogsWithHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tag definition audit logs with history params
func (o *GetTagDefinitionAuditLogsWithHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTagDefinitionID adds the tagDefinitionID to the get tag definition audit logs with history params
func (o *GetTagDefinitionAuditLogsWithHistoryParams) WithTagDefinitionID(tagDefinitionID strfmt.UUID) *GetTagDefinitionAuditLogsWithHistoryParams {
	o.SetTagDefinitionID(tagDefinitionID)
	return o
}

// SetTagDefinitionID adds the tagDefinitionId to the get tag definition audit logs with history params
func (o *GetTagDefinitionAuditLogsWithHistoryParams) SetTagDefinitionID(tagDefinitionID strfmt.UUID) {
	o.TagDefinitionID = tagDefinitionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTagDefinitionAuditLogsWithHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tagDefinitionId
	if err := r.SetPathParam("tagDefinitionId", o.TagDefinitionID.String()); err != nil {
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
