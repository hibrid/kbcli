// Code generated by go-swagger; DO NOT EDIT.

package invoice_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"
)

// New creates a new invoice item API client.
func New(transport runtime.ClientTransport,
	formats strfmt.Registry,
	authInfo runtime.ClientAuthInfoWriter,
	defaults KillbillDefaults) *Client {

	return &Client{transport: transport, formats: formats, authInfo: authInfo, defaults: defaults}
}

// killbill default values. When a call is made to an operation, these values are used
// if params doesn't specify them.
type KillbillDefaults interface {
	// Default CreatedBy. If not set explicitly in params, this will be used.
	XKillbillCreatedBy() *string
	// Default Comment. If not set explicitly in params, this will be used.
	XKillbillComment() *string
	// Default Reason. If not set explicitly in params, this will be used.
	XKillbillReason() *string
	// Default WithWithProfilingInfo. If not set explicitly in params, this will be used.
	KillbillWithProfilingInfo() *string
	// Default WithStackTrace. If not set explicitly in params, this will be used.
	KillbillWithStackTrace() *bool
}

/*
Client for invoice item API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
	defaults  KillbillDefaults
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateInvoiceItemCustomFields(ctx context.Context, params *CreateInvoiceItemCustomFieldsParams) (*CreateInvoiceItemCustomFieldsCreated, error)

	CreateInvoiceItemTags(ctx context.Context, params *CreateInvoiceItemTagsParams) (*CreateInvoiceItemTagsCreated, error)

	DeleteInvoiceItemCustomFields(ctx context.Context, params *DeleteInvoiceItemCustomFieldsParams) (*DeleteInvoiceItemCustomFieldsNoContent, error)

	DeleteInvoiceItemTags(ctx context.Context, params *DeleteInvoiceItemTagsParams) (*DeleteInvoiceItemTagsNoContent, error)

	GetInvoiceItemAuditLogsWithHistory(ctx context.Context, params *GetInvoiceItemAuditLogsWithHistoryParams) (*GetInvoiceItemAuditLogsWithHistoryOK, error)

	GetInvoiceItemCustomFields(ctx context.Context, params *GetInvoiceItemCustomFieldsParams) (*GetInvoiceItemCustomFieldsOK, error)

	GetInvoiceItemTags(ctx context.Context, params *GetInvoiceItemTagsParams) (*GetInvoiceItemTagsOK, error)

	ModifyInvoiceItemCustomFields(ctx context.Context, params *ModifyInvoiceItemCustomFieldsParams) (*ModifyInvoiceItemCustomFieldsNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateInvoiceItemCustomFields adds custom fields to invoice item
*/
func (a *Client) CreateInvoiceItemCustomFields(ctx context.Context, params *CreateInvoiceItemCustomFieldsParams) (*CreateInvoiceItemCustomFieldsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateInvoiceItemCustomFieldsParams()
	}
	getParams := NewCreateInvoiceItemCustomFieldsParams()
	getParams.Context = ctx
	params.Context = ctx
	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}
	getParams.XKillbillComment = params.XKillbillComment
	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}
	getParams.XKillbillCreatedBy = params.XKillbillCreatedBy
	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}
	getParams.XKillbillReason = params.XKillbillReason
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	getParams.WithStackTrace = params.WithStackTrace

	op := &runtime.ClientOperation{
		ID:                 "createInvoiceItemCustomFields",
		Method:             "POST",
		PathPattern:        "/1.0/kb/invoiceItems/{invoiceItemId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateInvoiceItemCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	createdResult := result.(*CreateInvoiceItemCustomFieldsCreated)
	location := kbcommon.ParseLocationHeader(createdResult.HttpResponse.GetHeader("Location"))
	if !params.ProcessLocationHeader || location == "" {
		return createdResult, nil
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createInvoiceItemCustomFields",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &CreateInvoiceItemCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return getResult.(*CreateInvoiceItemCustomFieldsCreated), nil

}

/*
  CreateInvoiceItemTags adds tags to invoice item
*/
func (a *Client) CreateInvoiceItemTags(ctx context.Context, params *CreateInvoiceItemTagsParams) (*CreateInvoiceItemTagsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateInvoiceItemTagsParams()
	}
	getParams := NewCreateInvoiceItemTagsParams()
	getParams.Context = ctx
	params.Context = ctx
	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}
	getParams.XKillbillComment = params.XKillbillComment
	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}
	getParams.XKillbillCreatedBy = params.XKillbillCreatedBy
	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}
	getParams.XKillbillReason = params.XKillbillReason
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	getParams.WithStackTrace = params.WithStackTrace

	op := &runtime.ClientOperation{
		ID:                 "createInvoiceItemTags",
		Method:             "POST",
		PathPattern:        "/1.0/kb/invoiceItems/{invoiceItemId}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateInvoiceItemTagsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	createdResult := result.(*CreateInvoiceItemTagsCreated)
	location := kbcommon.ParseLocationHeader(createdResult.HttpResponse.GetHeader("Location"))
	if !params.ProcessLocationHeader || location == "" {
		return createdResult, nil
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createInvoiceItemTags",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &CreateInvoiceItemTagsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return getResult.(*CreateInvoiceItemTagsCreated), nil

}

/*
  DeleteInvoiceItemCustomFields removes custom fields from invoice item
*/
func (a *Client) DeleteInvoiceItemCustomFields(ctx context.Context, params *DeleteInvoiceItemCustomFieldsParams) (*DeleteInvoiceItemCustomFieldsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInvoiceItemCustomFieldsParams()
	}
	params.Context = ctx
	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}

	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "deleteInvoiceItemCustomFields",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/invoiceItems/{invoiceItemId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteInvoiceItemCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteInvoiceItemCustomFieldsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteInvoiceItemCustomFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
  DeleteInvoiceItemTags removes tags from invoice item
*/
func (a *Client) DeleteInvoiceItemTags(ctx context.Context, params *DeleteInvoiceItemTagsParams) (*DeleteInvoiceItemTagsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInvoiceItemTagsParams()
	}
	params.Context = ctx
	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}

	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "deleteInvoiceItemTags",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/invoiceItems/{invoiceItemId}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteInvoiceItemTagsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteInvoiceItemTagsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteInvoiceItemTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
  GetInvoiceItemAuditLogsWithHistory retrieves invoice item audit logs with history by id
*/
func (a *Client) GetInvoiceItemAuditLogsWithHistory(ctx context.Context, params *GetInvoiceItemAuditLogsWithHistoryParams) (*GetInvoiceItemAuditLogsWithHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInvoiceItemAuditLogsWithHistoryParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getInvoiceItemAuditLogsWithHistory",
		Method:             "GET",
		PathPattern:        "/1.0/kb/invoiceItems/{invoiceItemId}/auditLogsWithHistory",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInvoiceItemAuditLogsWithHistoryReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInvoiceItemAuditLogsWithHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInvoiceItemAuditLogsWithHistory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
  GetInvoiceItemCustomFields retrieves invoice item custom fields
*/
func (a *Client) GetInvoiceItemCustomFields(ctx context.Context, params *GetInvoiceItemCustomFieldsParams) (*GetInvoiceItemCustomFieldsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInvoiceItemCustomFieldsParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getInvoiceItemCustomFields",
		Method:             "GET",
		PathPattern:        "/1.0/kb/invoiceItems/{invoiceItemId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInvoiceItemCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInvoiceItemCustomFieldsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInvoiceItemCustomFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
  GetInvoiceItemTags retrieves invoice item tags
*/
func (a *Client) GetInvoiceItemTags(ctx context.Context, params *GetInvoiceItemTagsParams) (*GetInvoiceItemTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInvoiceItemTagsParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getInvoiceItemTags",
		Method:             "GET",
		PathPattern:        "/1.0/kb/invoiceItems/{invoiceItemId}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInvoiceItemTagsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInvoiceItemTagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInvoiceItemTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
  ModifyInvoiceItemCustomFields modifies custom fields to invoice item
*/
func (a *Client) ModifyInvoiceItemCustomFields(ctx context.Context, params *ModifyInvoiceItemCustomFieldsParams) (*ModifyInvoiceItemCustomFieldsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyInvoiceItemCustomFieldsParams()
	}
	params.Context = ctx
	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}

	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "modifyInvoiceItemCustomFields",
		Method:             "PUT",
		PathPattern:        "/1.0/kb/invoiceItems/{invoiceItemId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ModifyInvoiceItemCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ModifyInvoiceItemCustomFieldsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for modifyInvoiceItemCustomFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
