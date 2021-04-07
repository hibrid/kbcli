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

// NewGetTenantByAPIKeyParams creates a new GetTenantByAPIKeyParams object
// with the default values initialized.
func NewGetTenantByAPIKeyParams() *GetTenantByAPIKeyParams {
	var ()
	return &GetTenantByAPIKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTenantByAPIKeyParamsWithTimeout creates a new GetTenantByAPIKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTenantByAPIKeyParamsWithTimeout(timeout time.Duration) *GetTenantByAPIKeyParams {
	var ()
	return &GetTenantByAPIKeyParams{

		timeout: timeout,
	}
}

// NewGetTenantByAPIKeyParamsWithContext creates a new GetTenantByAPIKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTenantByAPIKeyParamsWithContext(ctx context.Context) *GetTenantByAPIKeyParams {
	var ()
	return &GetTenantByAPIKeyParams{

		Context: ctx,
	}
}

// NewGetTenantByAPIKeyParamsWithHTTPClient creates a new GetTenantByAPIKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTenantByAPIKeyParamsWithHTTPClient(client *http.Client) *GetTenantByAPIKeyParams {
	var ()
	return &GetTenantByAPIKeyParams{
		HTTPClient: client,
	}
}

/*GetTenantByAPIKeyParams contains all the parameters to send to the API endpoint
for the get tenant by Api key operation typically these are written to a http.Request
*/
type GetTenantByAPIKeyParams struct {

	/*APIKey*/
	APIKey *string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get tenant by Api key params
func (o *GetTenantByAPIKeyParams) WithTimeout(timeout time.Duration) *GetTenantByAPIKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tenant by Api key params
func (o *GetTenantByAPIKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tenant by Api key params
func (o *GetTenantByAPIKeyParams) WithContext(ctx context.Context) *GetTenantByAPIKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tenant by Api key params
func (o *GetTenantByAPIKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tenant by Api key params
func (o *GetTenantByAPIKeyParams) WithHTTPClient(client *http.Client) *GetTenantByAPIKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tenant by Api key params
func (o *GetTenantByAPIKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIKey adds the aPIKey to the get tenant by Api key params
func (o *GetTenantByAPIKeyParams) WithAPIKey(aPIKey *string) *GetTenantByAPIKeyParams {
	o.SetAPIKey(aPIKey)
	return o
}

// SetAPIKey adds the apiKey to the get tenant by Api key params
func (o *GetTenantByAPIKeyParams) SetAPIKey(aPIKey *string) {
	o.APIKey = aPIKey
}

// WriteToRequest writes these params to a swagger request
func (o *GetTenantByAPIKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIKey != nil {

		// query param apiKey
		var qrAPIKey string
		if o.APIKey != nil {
			qrAPIKey = *o.APIKey
		}
		qAPIKey := qrAPIKey
		if qAPIKey != "" {
			if err := r.SetQueryParam("apiKey", qAPIKey); err != nil {
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
