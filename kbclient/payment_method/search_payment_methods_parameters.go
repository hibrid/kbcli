// Code generated by go-swagger; DO NOT EDIT.

package payment_method

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

// NewSearchPaymentMethodsParams creates a new SearchPaymentMethodsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchPaymentMethodsParams() *SearchPaymentMethodsParams {
	return &SearchPaymentMethodsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchPaymentMethodsParamsWithTimeout creates a new SearchPaymentMethodsParams object
// with the ability to set a timeout on a request.
func NewSearchPaymentMethodsParamsWithTimeout(timeout time.Duration) *SearchPaymentMethodsParams {
	return &SearchPaymentMethodsParams{
		timeout: timeout,
	}
}

// NewSearchPaymentMethodsParamsWithContext creates a new SearchPaymentMethodsParams object
// with the ability to set a context for a request.
func NewSearchPaymentMethodsParamsWithContext(ctx context.Context) *SearchPaymentMethodsParams {
	return &SearchPaymentMethodsParams{
		Context: ctx,
	}
}

// NewSearchPaymentMethodsParamsWithHTTPClient creates a new SearchPaymentMethodsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchPaymentMethodsParamsWithHTTPClient(client *http.Client) *SearchPaymentMethodsParams {
	return &SearchPaymentMethodsParams{
		HTTPClient: client,
	}
}

/* SearchPaymentMethodsParams contains all the parameters to send to the API endpoint
   for the search payment methods operation.

   Typically these are written to a http.Request.
*/
type SearchPaymentMethodsParams struct {

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

	// PluginName.
	PluginName *string

	// PluginProperty.
	PluginProperty []string

	// SearchKey.
	SearchKey string

	// WithPluginInfo.
	WithPluginInfo *bool

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the search payment methods params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchPaymentMethodsParams) WithDefaults() *SearchPaymentMethodsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search payment methods params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchPaymentMethodsParams) SetDefaults() {
	var (
		auditDefault = string("NONE")

		limitDefault = int64(100)

		offsetDefault = int64(0)

		withPluginInfoDefault = bool(false)
	)

	val := SearchPaymentMethodsParams{
		Audit:          &auditDefault,
		Limit:          &limitDefault,
		Offset:         &offsetDefault,
		WithPluginInfo: &withPluginInfoDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the search payment methods params
func (o *SearchPaymentMethodsParams) WithTimeout(timeout time.Duration) *SearchPaymentMethodsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search payment methods params
func (o *SearchPaymentMethodsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search payment methods params
func (o *SearchPaymentMethodsParams) WithContext(ctx context.Context) *SearchPaymentMethodsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search payment methods params
func (o *SearchPaymentMethodsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search payment methods params
func (o *SearchPaymentMethodsParams) WithHTTPClient(client *http.Client) *SearchPaymentMethodsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search payment methods params
func (o *SearchPaymentMethodsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the search payment methods params
func (o *SearchPaymentMethodsParams) WithAudit(audit *string) *SearchPaymentMethodsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the search payment methods params
func (o *SearchPaymentMethodsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithLimit adds the limit to the search payment methods params
func (o *SearchPaymentMethodsParams) WithLimit(limit *int64) *SearchPaymentMethodsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search payment methods params
func (o *SearchPaymentMethodsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the search payment methods params
func (o *SearchPaymentMethodsParams) WithOffset(offset *int64) *SearchPaymentMethodsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search payment methods params
func (o *SearchPaymentMethodsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPluginName adds the pluginName to the search payment methods params
func (o *SearchPaymentMethodsParams) WithPluginName(pluginName *string) *SearchPaymentMethodsParams {
	o.SetPluginName(pluginName)
	return o
}

// SetPluginName adds the pluginName to the search payment methods params
func (o *SearchPaymentMethodsParams) SetPluginName(pluginName *string) {
	o.PluginName = pluginName
}

// WithPluginProperty adds the pluginProperty to the search payment methods params
func (o *SearchPaymentMethodsParams) WithPluginProperty(pluginProperty []string) *SearchPaymentMethodsParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the search payment methods params
func (o *SearchPaymentMethodsParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithSearchKey adds the searchKey to the search payment methods params
func (o *SearchPaymentMethodsParams) WithSearchKey(searchKey string) *SearchPaymentMethodsParams {
	o.SetSearchKey(searchKey)
	return o
}

// SetSearchKey adds the searchKey to the search payment methods params
func (o *SearchPaymentMethodsParams) SetSearchKey(searchKey string) {
	o.SearchKey = searchKey
}

// WithWithPluginInfo adds the withPluginInfo to the search payment methods params
func (o *SearchPaymentMethodsParams) WithWithPluginInfo(withPluginInfo *bool) *SearchPaymentMethodsParams {
	o.SetWithPluginInfo(withPluginInfo)
	return o
}

// SetWithPluginInfo adds the withPluginInfo to the search payment methods params
func (o *SearchPaymentMethodsParams) SetWithPluginInfo(withPluginInfo *bool) {
	o.WithPluginInfo = withPluginInfo
}

// WriteToRequest writes these params to a swagger request
func (o *SearchPaymentMethodsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.PluginName != nil {

		// query param pluginName
		var qrPluginName string

		if o.PluginName != nil {
			qrPluginName = *o.PluginName
		}
		qPluginName := qrPluginName
		if qPluginName != "" {

			if err := r.SetQueryParam("pluginName", qPluginName); err != nil {
				return err
			}
		}
	}

	if o.PluginProperty != nil {

		// binding items for pluginProperty
		joinedPluginProperty := o.bindParamPluginProperty(reg)

		// query array param pluginProperty
		if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
			return err
		}
	}

	// path param searchKey
	if err := r.SetPathParam("searchKey", o.SearchKey); err != nil {
		return err
	}

	if o.WithPluginInfo != nil {

		// query param withPluginInfo
		var qrWithPluginInfo bool

		if o.WithPluginInfo != nil {
			qrWithPluginInfo = *o.WithPluginInfo
		}
		qWithPluginInfo := swag.FormatBool(qrWithPluginInfo)
		if qWithPluginInfo != "" {

			if err := r.SetQueryParam("withPluginInfo", qWithPluginInfo); err != nil {
				return err
			}
		}
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

// bindParamSearchPaymentMethods binds the parameter pluginProperty
func (o *SearchPaymentMethodsParams) bindParamPluginProperty(formats strfmt.Registry) []string {
	pluginPropertyIR := o.PluginProperty

	var pluginPropertyIC []string
	for _, pluginPropertyIIR := range pluginPropertyIR { // explode []string

		pluginPropertyIIV := pluginPropertyIIR // string as string
		pluginPropertyIC = append(pluginPropertyIC, pluginPropertyIIV)
	}

	// items.CollectionFormat: "multi"
	pluginPropertyIS := swag.JoinByFormat(pluginPropertyIC, "multi")

	return pluginPropertyIS
}
