// Code generated by go-swagger; DO NOT EDIT.

package tag

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

// NewSearchTagsParams creates a new SearchTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchTagsParams() *SearchTagsParams {
	return &SearchTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchTagsParamsWithTimeout creates a new SearchTagsParams object
// with the ability to set a timeout on a request.
func NewSearchTagsParamsWithTimeout(timeout time.Duration) *SearchTagsParams {
	return &SearchTagsParams{
		timeout: timeout,
	}
}

// NewSearchTagsParamsWithContext creates a new SearchTagsParams object
// with the ability to set a context for a request.
func NewSearchTagsParamsWithContext(ctx context.Context) *SearchTagsParams {
	return &SearchTagsParams{
		Context: ctx,
	}
}

// NewSearchTagsParamsWithHTTPClient creates a new SearchTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchTagsParamsWithHTTPClient(client *http.Client) *SearchTagsParams {
	return &SearchTagsParams{
		HTTPClient: client,
	}
}

/* SearchTagsParams contains all the parameters to send to the API endpoint
   for the search tags operation.

   Typically these are written to a http.Request.
*/
type SearchTagsParams struct {

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

	// SearchKey.
	SearchKey string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the search tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchTagsParams) WithDefaults() *SearchTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchTagsParams) SetDefaults() {
	var (
		auditDefault = string("NONE")

		limitDefault = int64(100)

		offsetDefault = int64(0)
	)

	val := SearchTagsParams{
		Audit:  &auditDefault,
		Limit:  &limitDefault,
		Offset: &offsetDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the search tags params
func (o *SearchTagsParams) WithTimeout(timeout time.Duration) *SearchTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search tags params
func (o *SearchTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search tags params
func (o *SearchTagsParams) WithContext(ctx context.Context) *SearchTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search tags params
func (o *SearchTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search tags params
func (o *SearchTagsParams) WithHTTPClient(client *http.Client) *SearchTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search tags params
func (o *SearchTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the search tags params
func (o *SearchTagsParams) WithAudit(audit *string) *SearchTagsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the search tags params
func (o *SearchTagsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithLimit adds the limit to the search tags params
func (o *SearchTagsParams) WithLimit(limit *int64) *SearchTagsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search tags params
func (o *SearchTagsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the search tags params
func (o *SearchTagsParams) WithOffset(offset *int64) *SearchTagsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search tags params
func (o *SearchTagsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSearchKey adds the searchKey to the search tags params
func (o *SearchTagsParams) WithSearchKey(searchKey string) *SearchTagsParams {
	o.SetSearchKey(searchKey)
	return o
}

// SetSearchKey adds the searchKey to the search tags params
func (o *SearchTagsParams) SetSearchKey(searchKey string) {
	o.SearchKey = searchKey
}

// WriteToRequest writes these params to a swagger request
func (o *SearchTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param searchKey
	if err := r.SetPathParam("searchKey", o.SearchKey); err != nil {
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
