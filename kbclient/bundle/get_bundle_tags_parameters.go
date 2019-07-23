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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBundleTagsParams creates a new GetBundleTagsParams object
// with the default values initialized.
func NewGetBundleTagsParams() *GetBundleTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetBundleTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBundleTagsParamsWithTimeout creates a new GetBundleTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBundleTagsParamsWithTimeout(timeout time.Duration) *GetBundleTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetBundleTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,

		timeout: timeout,
	}
}

// NewGetBundleTagsParamsWithContext creates a new GetBundleTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBundleTagsParamsWithContext(ctx context.Context) *GetBundleTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetBundleTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,

		Context: ctx,
	}
}

// NewGetBundleTagsParamsWithHTTPClient creates a new GetBundleTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBundleTagsParamsWithHTTPClient(client *http.Client) *GetBundleTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetBundleTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,
		HTTPClient:      client,
	}
}

/*GetBundleTagsParams contains all the parameters to send to the API endpoint
for the get bundle tags operation typically these are written to a http.Request
*/
type GetBundleTagsParams struct {

	/*Audit*/
	Audit *string
	/*BundleID*/
	BundleID strfmt.UUID
	/*IncludedDeleted*/
	IncludedDeleted *bool

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get bundle tags params
func (o *GetBundleTagsParams) WithTimeout(timeout time.Duration) *GetBundleTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bundle tags params
func (o *GetBundleTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bundle tags params
func (o *GetBundleTagsParams) WithContext(ctx context.Context) *GetBundleTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bundle tags params
func (o *GetBundleTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bundle tags params
func (o *GetBundleTagsParams) WithHTTPClient(client *http.Client) *GetBundleTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bundle tags params
func (o *GetBundleTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the get bundle tags params
func (o *GetBundleTagsParams) WithAudit(audit *string) *GetBundleTagsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get bundle tags params
func (o *GetBundleTagsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithBundleID adds the bundleID to the get bundle tags params
func (o *GetBundleTagsParams) WithBundleID(bundleID strfmt.UUID) *GetBundleTagsParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the get bundle tags params
func (o *GetBundleTagsParams) SetBundleID(bundleID strfmt.UUID) {
	o.BundleID = bundleID
}

// WithIncludedDeleted adds the includedDeleted to the get bundle tags params
func (o *GetBundleTagsParams) WithIncludedDeleted(includedDeleted *bool) *GetBundleTagsParams {
	o.SetIncludedDeleted(includedDeleted)
	return o
}

// SetIncludedDeleted adds the includedDeleted to the get bundle tags params
func (o *GetBundleTagsParams) SetIncludedDeleted(includedDeleted *bool) {
	o.IncludedDeleted = includedDeleted
}

// WriteToRequest writes these params to a swagger request
func (o *GetBundleTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IncludedDeleted != nil {

		// query param includedDeleted
		var qrIncludedDeleted bool
		if o.IncludedDeleted != nil {
			qrIncludedDeleted = *o.IncludedDeleted
		}
		qIncludedDeleted := swag.FormatBool(qrIncludedDeleted)
		if qIncludedDeleted != "" {
			if err := r.SetQueryParam("includedDeleted", qIncludedDeleted); err != nil {
				return err
			}
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
