// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

// NewUploadCatalogTranslationParams creates a new UploadCatalogTranslationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadCatalogTranslationParams() *UploadCatalogTranslationParams {
	return &UploadCatalogTranslationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadCatalogTranslationParamsWithTimeout creates a new UploadCatalogTranslationParams object
// with the ability to set a timeout on a request.
func NewUploadCatalogTranslationParamsWithTimeout(timeout time.Duration) *UploadCatalogTranslationParams {
	return &UploadCatalogTranslationParams{
		timeout: timeout,
	}
}

// NewUploadCatalogTranslationParamsWithContext creates a new UploadCatalogTranslationParams object
// with the ability to set a context for a request.
func NewUploadCatalogTranslationParamsWithContext(ctx context.Context) *UploadCatalogTranslationParams {
	return &UploadCatalogTranslationParams{
		Context: ctx,
	}
}

// NewUploadCatalogTranslationParamsWithHTTPClient creates a new UploadCatalogTranslationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUploadCatalogTranslationParamsWithHTTPClient(client *http.Client) *UploadCatalogTranslationParams {
	return &UploadCatalogTranslationParams{
		HTTPClient: client,
	}
}

/* UploadCatalogTranslationParams contains all the parameters to send to the API endpoint
   for the upload catalog translation operation.

   Typically these are written to a http.Request.
*/
type UploadCatalogTranslationParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body string

	// DeleteIfExists.
	DeleteIfExists *bool

	// Locale.
	Locale string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the upload catalog translation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadCatalogTranslationParams) WithDefaults() *UploadCatalogTranslationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload catalog translation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadCatalogTranslationParams) SetDefaults() {
	var (
		deleteIfExistsDefault = bool(false)
	)

	val := UploadCatalogTranslationParams{
		DeleteIfExists: &deleteIfExistsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the upload catalog translation params
func (o *UploadCatalogTranslationParams) WithTimeout(timeout time.Duration) *UploadCatalogTranslationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload catalog translation params
func (o *UploadCatalogTranslationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload catalog translation params
func (o *UploadCatalogTranslationParams) WithContext(ctx context.Context) *UploadCatalogTranslationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload catalog translation params
func (o *UploadCatalogTranslationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload catalog translation params
func (o *UploadCatalogTranslationParams) WithHTTPClient(client *http.Client) *UploadCatalogTranslationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload catalog translation params
func (o *UploadCatalogTranslationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the upload catalog translation params
func (o *UploadCatalogTranslationParams) WithXKillbillComment(xKillbillComment *string) *UploadCatalogTranslationParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the upload catalog translation params
func (o *UploadCatalogTranslationParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the upload catalog translation params
func (o *UploadCatalogTranslationParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *UploadCatalogTranslationParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the upload catalog translation params
func (o *UploadCatalogTranslationParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the upload catalog translation params
func (o *UploadCatalogTranslationParams) WithXKillbillReason(xKillbillReason *string) *UploadCatalogTranslationParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the upload catalog translation params
func (o *UploadCatalogTranslationParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the upload catalog translation params
func (o *UploadCatalogTranslationParams) WithBody(body string) *UploadCatalogTranslationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upload catalog translation params
func (o *UploadCatalogTranslationParams) SetBody(body string) {
	o.Body = body
}

// WithDeleteIfExists adds the deleteIfExists to the upload catalog translation params
func (o *UploadCatalogTranslationParams) WithDeleteIfExists(deleteIfExists *bool) *UploadCatalogTranslationParams {
	o.SetDeleteIfExists(deleteIfExists)
	return o
}

// SetDeleteIfExists adds the deleteIfExists to the upload catalog translation params
func (o *UploadCatalogTranslationParams) SetDeleteIfExists(deleteIfExists *bool) {
	o.DeleteIfExists = deleteIfExists
}

// WithLocale adds the locale to the upload catalog translation params
func (o *UploadCatalogTranslationParams) WithLocale(locale string) *UploadCatalogTranslationParams {
	o.SetLocale(locale)
	return o
}

// SetLocale adds the locale to the upload catalog translation params
func (o *UploadCatalogTranslationParams) SetLocale(locale string) {
	o.Locale = locale
}

// WriteToRequest writes these params to a swagger request
func (o *UploadCatalogTranslationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XKillbillComment != nil {

		// header param X-Killbill-Comment
		if err := r.SetHeaderParam("X-Killbill-Comment", *o.XKillbillComment); err != nil {
			return err
		}
	}

	// header param X-Killbill-CreatedBy
	if err := r.SetHeaderParam("X-Killbill-CreatedBy", o.XKillbillCreatedBy); err != nil {
		return err
	}

	if o.XKillbillReason != nil {

		// header param X-Killbill-Reason
		if err := r.SetHeaderParam("X-Killbill-Reason", *o.XKillbillReason); err != nil {
			return err
		}
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if o.DeleteIfExists != nil {

		// query param deleteIfExists
		var qrDeleteIfExists bool

		if o.DeleteIfExists != nil {
			qrDeleteIfExists = *o.DeleteIfExists
		}
		qDeleteIfExists := swag.FormatBool(qrDeleteIfExists)
		if qDeleteIfExists != "" {

			if err := r.SetQueryParam("deleteIfExists", qDeleteIfExists); err != nil {
				return err
			}
		}
	}

	// path param locale
	if err := r.SetPathParam("locale", o.Locale); err != nil {
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
