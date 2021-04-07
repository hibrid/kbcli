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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAllPluginConfigurationParams creates a new GetAllPluginConfigurationParams object
// with the default values initialized.
func NewGetAllPluginConfigurationParams() *GetAllPluginConfigurationParams {
	var ()
	return &GetAllPluginConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllPluginConfigurationParamsWithTimeout creates a new GetAllPluginConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllPluginConfigurationParamsWithTimeout(timeout time.Duration) *GetAllPluginConfigurationParams {
	var ()
	return &GetAllPluginConfigurationParams{

		timeout: timeout,
	}
}

// NewGetAllPluginConfigurationParamsWithContext creates a new GetAllPluginConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllPluginConfigurationParamsWithContext(ctx context.Context) *GetAllPluginConfigurationParams {
	var ()
	return &GetAllPluginConfigurationParams{

		Context: ctx,
	}
}

// NewGetAllPluginConfigurationParamsWithHTTPClient creates a new GetAllPluginConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllPluginConfigurationParamsWithHTTPClient(client *http.Client) *GetAllPluginConfigurationParams {
	var ()
	return &GetAllPluginConfigurationParams{
		HTTPClient: client,
	}
}

/*GetAllPluginConfigurationParams contains all the parameters to send to the API endpoint
for the get all plugin configuration operation typically these are written to a http.Request
*/
type GetAllPluginConfigurationParams struct {

	/*KeyPrefix*/
	KeyPrefix string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get all plugin configuration params
func (o *GetAllPluginConfigurationParams) WithTimeout(timeout time.Duration) *GetAllPluginConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all plugin configuration params
func (o *GetAllPluginConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all plugin configuration params
func (o *GetAllPluginConfigurationParams) WithContext(ctx context.Context) *GetAllPluginConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all plugin configuration params
func (o *GetAllPluginConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all plugin configuration params
func (o *GetAllPluginConfigurationParams) WithHTTPClient(client *http.Client) *GetAllPluginConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all plugin configuration params
func (o *GetAllPluginConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKeyPrefix adds the keyPrefix to the get all plugin configuration params
func (o *GetAllPluginConfigurationParams) WithKeyPrefix(keyPrefix string) *GetAllPluginConfigurationParams {
	o.SetKeyPrefix(keyPrefix)
	return o
}

// SetKeyPrefix adds the keyPrefix to the get all plugin configuration params
func (o *GetAllPluginConfigurationParams) SetKeyPrefix(keyPrefix string) {
	o.KeyPrefix = keyPrefix
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllPluginConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param keyPrefix
	if err := r.SetPathParam("keyPrefix", o.KeyPrefix); err != nil {
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
