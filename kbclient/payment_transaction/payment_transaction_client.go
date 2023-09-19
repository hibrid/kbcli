// Code generated by go-swagger; DO NOT EDIT.

package payment_transaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"
)

// New creates a new payment transaction API client.
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
Client for payment transaction API
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
	CreateTransactionCustomFields(ctx context.Context, params *CreateTransactionCustomFieldsParams) (*CreateTransactionCustomFieldsCreated, error)

	CreateTransactionTags(ctx context.Context, params *CreateTransactionTagsParams) (*CreateTransactionTagsCreated, error)

	DeleteTransactionCustomFields(ctx context.Context, params *DeleteTransactionCustomFieldsParams) (*DeleteTransactionCustomFieldsNoContent, error)

	DeleteTransactionTags(ctx context.Context, params *DeleteTransactionTagsParams) (*DeleteTransactionTagsNoContent, error)

	GetPaymentByTransactionExternalKey(ctx context.Context, params *GetPaymentByTransactionExternalKeyParams) (*GetPaymentByTransactionExternalKeyOK, error)

	GetPaymentByTransactionID(ctx context.Context, params *GetPaymentByTransactionIDParams) (*GetPaymentByTransactionIDOK, error)

	GetTransactionAuditLogsWithHistory(ctx context.Context, params *GetTransactionAuditLogsWithHistoryParams) (*GetTransactionAuditLogsWithHistoryOK, error)

	GetTransactionCustomFields(ctx context.Context, params *GetTransactionCustomFieldsParams) (*GetTransactionCustomFieldsOK, error)

	GetTransactionTags(ctx context.Context, params *GetTransactionTagsParams) (*GetTransactionTagsOK, error)

	ModifyTransactionCustomFields(ctx context.Context, params *ModifyTransactionCustomFieldsParams) (*ModifyTransactionCustomFieldsNoContent, error)

	NotifyStateChanged(ctx context.Context, params *NotifyStateChangedParams) (*NotifyStateChangedCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateTransactionCustomFields adds custom fields to payment transaction
*/
func (a *Client) CreateTransactionCustomFields(ctx context.Context, params *CreateTransactionCustomFieldsParams) (*CreateTransactionCustomFieldsCreated, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewCreateTransactionCustomFieldsParams()
	}
	getParams := NewCreateTransactionCustomFieldsParams()
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
		ID:                 "createTransactionCustomFields",
		Method:             "POST",
		PathPattern:        "/1.0/kb/paymentTransactions/{transactionId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateTransactionCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	var location string
	switch value := result.(type) {
	case *CreateTransactionCustomFieldsCreated:
		location = kbcommon.ParseLocationHeader(value.HttpResponse.GetHeader("Location"))
		if !params.ProcessLocationHeader || location == "" {
			return value, nil
		}
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createTransactionCustomFields",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &CreateTransactionCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch value := getResult.(type) {
	case *CreateTransactionCustomFieldsCreated:
		return value, nil
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

}

/*
  CreateTransactionTags adds tags to payment transaction
*/
func (a *Client) CreateTransactionTags(ctx context.Context, params *CreateTransactionTagsParams) (*CreateTransactionTagsCreated, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewCreateTransactionTagsParams()
	}
	getParams := NewCreateTransactionTagsParams()
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
		ID:                 "createTransactionTags",
		Method:             "POST",
		PathPattern:        "/1.0/kb/paymentTransactions/{transactionId}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateTransactionTagsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	var location string
	switch value := result.(type) {
	case *CreateTransactionTagsCreated:
		location = kbcommon.ParseLocationHeader(value.HttpResponse.GetHeader("Location"))
		if !params.ProcessLocationHeader || location == "" {
			return value, nil
		}
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createTransactionTags",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &CreateTransactionTagsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch value := getResult.(type) {
	case *CreateTransactionTagsCreated:
		return value, nil
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

}

/*
  DeleteTransactionCustomFields removes custom fields from payment transaction
*/
func (a *Client) DeleteTransactionCustomFields(ctx context.Context, params *DeleteTransactionCustomFieldsParams) (*DeleteTransactionCustomFieldsNoContent, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewDeleteTransactionCustomFieldsParams()
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
		ID:                 "deleteTransactionCustomFields",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/paymentTransactions/{transactionId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTransactionCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteTransactionCustomFieldsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteTransactionCustomFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
  DeleteTransactionTags removes tags from payment transaction
*/
func (a *Client) DeleteTransactionTags(ctx context.Context, params *DeleteTransactionTagsParams) (*DeleteTransactionTagsNoContent, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewDeleteTransactionTagsParams()
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
		ID:                 "deleteTransactionTags",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/paymentTransactions/{transactionId}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTransactionTagsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteTransactionTagsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteTransactionTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
  GetPaymentByTransactionExternalKey retrieves a payment by transaction external key
*/
func (a *Client) GetPaymentByTransactionExternalKey(ctx context.Context, params *GetPaymentByTransactionExternalKeyParams) (*GetPaymentByTransactionExternalKeyOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewGetPaymentByTransactionExternalKeyParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getPaymentByTransactionExternalKey",
		Method:             "GET",
		PathPattern:        "/1.0/kb/paymentTransactions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPaymentByTransactionExternalKeyReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPaymentByTransactionExternalKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPaymentByTransactionExternalKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
  GetPaymentByTransactionID retrieves a payment by transaction id
*/
func (a *Client) GetPaymentByTransactionID(ctx context.Context, params *GetPaymentByTransactionIDParams) (*GetPaymentByTransactionIDOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewGetPaymentByTransactionIDParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getPaymentByTransactionId",
		Method:             "GET",
		PathPattern:        "/1.0/kb/paymentTransactions/{transactionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPaymentByTransactionIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPaymentByTransactionIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPaymentByTransactionId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
  GetTransactionAuditLogsWithHistory retrieves payment transaction audit logs with history by id
*/
func (a *Client) GetTransactionAuditLogsWithHistory(ctx context.Context, params *GetTransactionAuditLogsWithHistoryParams) (*GetTransactionAuditLogsWithHistoryOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewGetTransactionAuditLogsWithHistoryParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getTransactionAuditLogsWithHistory",
		Method:             "GET",
		PathPattern:        "/1.0/kb/paymentTransactions/{transactionId}/auditLogsWithHistory",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTransactionAuditLogsWithHistoryReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTransactionAuditLogsWithHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTransactionAuditLogsWithHistory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
  GetTransactionCustomFields retrieves payment transaction custom fields
*/
func (a *Client) GetTransactionCustomFields(ctx context.Context, params *GetTransactionCustomFieldsParams) (*GetTransactionCustomFieldsOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewGetTransactionCustomFieldsParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getTransactionCustomFields",
		Method:             "GET",
		PathPattern:        "/1.0/kb/paymentTransactions/{transactionId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTransactionCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTransactionCustomFieldsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTransactionCustomFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
  GetTransactionTags retrieves payment transaction tags
*/
func (a *Client) GetTransactionTags(ctx context.Context, params *GetTransactionTagsParams) (*GetTransactionTagsOK, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewGetTransactionTagsParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	op := &runtime.ClientOperation{
		ID:                 "getTransactionTags",
		Method:             "GET",
		PathPattern:        "/1.0/kb/paymentTransactions/{transactionId}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTransactionTagsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTransactionTagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTransactionTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
  ModifyTransactionCustomFields modifies custom fields to payment transaction
*/
func (a *Client) ModifyTransactionCustomFields(ctx context.Context, params *ModifyTransactionCustomFieldsParams) (*ModifyTransactionCustomFieldsNoContent, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewModifyTransactionCustomFieldsParams()
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
		ID:                 "modifyTransactionCustomFields",
		Method:             "PUT",
		PathPattern:        "/1.0/kb/paymentTransactions/{transactionId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ModifyTransactionCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ModifyTransactionCustomFieldsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for modifyTransactionCustomFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

/*
  NotifyStateChanged marks a pending payment transaction as succeeded or failed
*/
func (a *Client) NotifyStateChanged(ctx context.Context, params *NotifyStateChangedParams) (*NotifyStateChangedCreated, error) {
	// TODO: Validate the params before sending

	if params == nil {
		params = NewNotifyStateChangedParams()
	}
	getParams := NewNotifyStateChangedParams()
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
		ID:                 "notifyStateChanged",
		Method:             "POST",
		PathPattern:        "/1.0/kb/paymentTransactions/{transactionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &NotifyStateChangedReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	var location string
	switch value := result.(type) {
	case *NotifyStateChangedCreated:
		location = kbcommon.ParseLocationHeader(value.HttpResponse.GetHeader("Location"))
		if !params.ProcessLocationHeader || location == "" {
			return value, nil
		}
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "notifyStateChanged",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &NotifyStateChangedReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}

	switch value := getResult.(type) {
	case *NotifyStateChangedCreated:
		return value, nil
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
