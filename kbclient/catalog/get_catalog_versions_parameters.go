// Code generated by go-swagger; DO NOT EDIT.

package catalog

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

// NewGetCatalogVersionsParams creates a new GetCatalogVersionsParams object
// with the default values initialized.
func NewGetCatalogVersionsParams() *GetCatalogVersionsParams {
	var ()
	return &GetCatalogVersionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCatalogVersionsParamsWithTimeout creates a new GetCatalogVersionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCatalogVersionsParamsWithTimeout(timeout time.Duration) *GetCatalogVersionsParams {
	var ()
	return &GetCatalogVersionsParams{

		timeout: timeout,
	}
}

// NewGetCatalogVersionsParamsWithContext creates a new GetCatalogVersionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCatalogVersionsParamsWithContext(ctx context.Context) *GetCatalogVersionsParams {
	var ()
	return &GetCatalogVersionsParams{

		Context: ctx,
	}
}

// NewGetCatalogVersionsParamsWithHTTPClient creates a new GetCatalogVersionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCatalogVersionsParamsWithHTTPClient(client *http.Client) *GetCatalogVersionsParams {
	var ()
	return &GetCatalogVersionsParams{
		HTTPClient: client,
	}
}

/*GetCatalogVersionsParams contains all the parameters to send to the API endpoint
for the get catalog versions operation typically these are written to a http.Request
*/
type GetCatalogVersionsParams struct {

	/*AccountID*/
	AccountID *strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get catalog versions params
func (o *GetCatalogVersionsParams) WithTimeout(timeout time.Duration) *GetCatalogVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get catalog versions params
func (o *GetCatalogVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get catalog versions params
func (o *GetCatalogVersionsParams) WithContext(ctx context.Context) *GetCatalogVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get catalog versions params
func (o *GetCatalogVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get catalog versions params
func (o *GetCatalogVersionsParams) WithHTTPClient(client *http.Client) *GetCatalogVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get catalog versions params
func (o *GetCatalogVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get catalog versions params
func (o *GetCatalogVersionsParams) WithAccountID(accountID *strfmt.UUID) *GetCatalogVersionsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get catalog versions params
func (o *GetCatalogVersionsParams) SetAccountID(accountID *strfmt.UUID) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCatalogVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountID != nil {

		// query param accountId
		var qrAccountID strfmt.UUID
		if o.AccountID != nil {
			qrAccountID = *o.AccountID
		}
		qAccountID := qrAccountID.String()
		if qAccountID != "" {
			if err := r.SetQueryParam("accountId", qAccountID); err != nil {
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
