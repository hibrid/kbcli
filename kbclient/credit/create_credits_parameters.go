// Code generated by go-swagger; DO NOT EDIT.

package credit

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

	"github.com/killbill/kbcli/v3/kbmodel"
)

// NewCreateCreditsParams creates a new CreateCreditsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateCreditsParams() *CreateCreditsParams {
	return &CreateCreditsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCreditsParamsWithTimeout creates a new CreateCreditsParams object
// with the ability to set a timeout on a request.
func NewCreateCreditsParamsWithTimeout(timeout time.Duration) *CreateCreditsParams {
	return &CreateCreditsParams{
		timeout: timeout,
	}
}

// NewCreateCreditsParamsWithContext creates a new CreateCreditsParams object
// with the ability to set a context for a request.
func NewCreateCreditsParamsWithContext(ctx context.Context) *CreateCreditsParams {
	return &CreateCreditsParams{
		Context: ctx,
	}
}

// NewCreateCreditsParamsWithHTTPClient creates a new CreateCreditsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateCreditsParamsWithHTTPClient(client *http.Client) *CreateCreditsParams {
	return &CreateCreditsParams{
		HTTPClient: client,
	}
}

/* CreateCreditsParams contains all the parameters to send to the API endpoint
   for the create credits operation.

   Typically these are written to a http.Request.
*/
type CreateCreditsParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// AutoCommit.
	AutoCommit *bool

	// Body.
	Body []*kbmodel.InvoiceItem

	// PluginProperty.
	PluginProperty []string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the create credits params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCreditsParams) WithDefaults() *CreateCreditsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create credits params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateCreditsParams) SetDefaults() {
	var (
		autoCommitDefault = bool(false)
	)

	val := CreateCreditsParams{
		AutoCommit: &autoCommitDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create credits params
func (o *CreateCreditsParams) WithTimeout(timeout time.Duration) *CreateCreditsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create credits params
func (o *CreateCreditsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create credits params
func (o *CreateCreditsParams) WithContext(ctx context.Context) *CreateCreditsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create credits params
func (o *CreateCreditsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create credits params
func (o *CreateCreditsParams) WithHTTPClient(client *http.Client) *CreateCreditsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create credits params
func (o *CreateCreditsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the create credits params
func (o *CreateCreditsParams) WithXKillbillComment(xKillbillComment *string) *CreateCreditsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create credits params
func (o *CreateCreditsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create credits params
func (o *CreateCreditsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateCreditsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create credits params
func (o *CreateCreditsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create credits params
func (o *CreateCreditsParams) WithXKillbillReason(xKillbillReason *string) *CreateCreditsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create credits params
func (o *CreateCreditsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAutoCommit adds the autoCommit to the create credits params
func (o *CreateCreditsParams) WithAutoCommit(autoCommit *bool) *CreateCreditsParams {
	o.SetAutoCommit(autoCommit)
	return o
}

// SetAutoCommit adds the autoCommit to the create credits params
func (o *CreateCreditsParams) SetAutoCommit(autoCommit *bool) {
	o.AutoCommit = autoCommit
}

// WithBody adds the body to the create credits params
func (o *CreateCreditsParams) WithBody(body []*kbmodel.InvoiceItem) *CreateCreditsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create credits params
func (o *CreateCreditsParams) SetBody(body []*kbmodel.InvoiceItem) {
	o.Body = body
}

// WithPluginProperty adds the pluginProperty to the create credits params
func (o *CreateCreditsParams) WithPluginProperty(pluginProperty []string) *CreateCreditsParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the create credits params
func (o *CreateCreditsParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCreditsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.AutoCommit != nil {

		// query param autoCommit
		var qrAutoCommit bool

		if o.AutoCommit != nil {
			qrAutoCommit = *o.AutoCommit
		}
		qAutoCommit := swag.FormatBool(qrAutoCommit)
		if qAutoCommit != "" {

			if err := r.SetQueryParam("autoCommit", qAutoCommit); err != nil {
				return err
			}
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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

// bindParamCreateCredits binds the parameter pluginProperty
func (o *CreateCreditsParams) bindParamPluginProperty(formats strfmt.Registry) []string {
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
