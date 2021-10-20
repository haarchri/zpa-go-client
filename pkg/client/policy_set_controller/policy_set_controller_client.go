// Code generated by go-swagger; DO NOT EDIT.

package policy_set_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new policy set controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for policy set controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddRuleToPolicySetUsingPOST1(params *AddRuleToPolicySetUsingPOST1Params, opts ...ClientOption) (*AddRuleToPolicySetUsingPOST1Created, error)

	DeleteRuleInPolicySetUsingDELETE1(params *DeleteRuleInPolicySetUsingDELETE1Params, opts ...ClientOption) (*DeleteRuleInPolicySetUsingDELETE1NoContent, error)

	GetBypassPolicySetUsingGET1(params *GetBypassPolicySetUsingGET1Params, opts ...ClientOption) (*GetBypassPolicySetUsingGET1OK, error)

	GetGlobalPolicySetUsingGET1(params *GetGlobalPolicySetUsingGET1Params, opts ...ClientOption) (*GetGlobalPolicySetUsingGET1OK, error)

	GetPolicyRulesByPageUsingGET1(params *GetPolicyRulesByPageUsingGET1Params, opts ...ClientOption) (*GetPolicyRulesByPageUsingGET1OK, error)

	GetReauthPolicySetUsingGET1(params *GetReauthPolicySetUsingGET1Params, opts ...ClientOption) (*GetReauthPolicySetUsingGET1OK, error)

	GetRuleInPolicySetUsingGET1(params *GetRuleInPolicySetUsingGET1Params, opts ...ClientOption) (*GetRuleInPolicySetUsingGET1OK, error)

	UpdateRuleToPolicySetUsingPUT1(params *UpdateRuleToPolicySetUsingPUT1Params, opts ...ClientOption) (*UpdateRuleToPolicySetUsingPUT1Created, *UpdateRuleToPolicySetUsingPUT1NoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddRuleToPolicySetUsingPOST1 adds a new policy rule for a given policy
*/
func (a *Client) AddRuleToPolicySetUsingPOST1(params *AddRuleToPolicySetUsingPOST1Params, opts ...ClientOption) (*AddRuleToPolicySetUsingPOST1Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddRuleToPolicySetUsingPOST1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "addRuleToPolicySetUsingPOST_1",
		Method:             "POST",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/policySet/{policySetId}/rule",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddRuleToPolicySetUsingPOST1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddRuleToPolicySetUsingPOST1Created)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addRuleToPolicySetUsingPOST_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteRuleInPolicySetUsingDELETE1 deletes a rule in a policy
*/
func (a *Client) DeleteRuleInPolicySetUsingDELETE1(params *DeleteRuleInPolicySetUsingDELETE1Params, opts ...ClientOption) (*DeleteRuleInPolicySetUsingDELETE1NoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRuleInPolicySetUsingDELETE1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteRuleInPolicySetUsingDELETE_1",
		Method:             "DELETE",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/policySet/{policySetId}/rule/{ruleId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRuleInPolicySetUsingDELETE1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteRuleInPolicySetUsingDELETE1NoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteRuleInPolicySetUsingDELETE_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBypassPolicySetUsingGET1 gets the bypass policy and all rules for a client forwarding policy rule
*/
func (a *Client) GetBypassPolicySetUsingGET1(params *GetBypassPolicySetUsingGET1Params, opts ...ClientOption) (*GetBypassPolicySetUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBypassPolicySetUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBypassPolicySetUsingGET_1",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/policySet/bypass",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBypassPolicySetUsingGET1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBypassPolicySetUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBypassPolicySetUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetGlobalPolicySetUsingGET1 gets the global policy
*/
func (a *Client) GetGlobalPolicySetUsingGET1(params *GetGlobalPolicySetUsingGET1Params, opts ...ClientOption) (*GetGlobalPolicySetUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGlobalPolicySetUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGlobalPolicySetUsingGET_1",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/policySet/global",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGlobalPolicySetUsingGET1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGlobalPolicySetUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGlobalPolicySetUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPolicyRulesByPageUsingGET1 gets paginated policy rules for a given policy type
*/
func (a *Client) GetPolicyRulesByPageUsingGET1(params *GetPolicyRulesByPageUsingGET1Params, opts ...ClientOption) (*GetPolicyRulesByPageUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPolicyRulesByPageUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPolicyRulesByPageUsingGET_1",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/policySet/rules/policyType/{policyType}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPolicyRulesByPageUsingGET1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPolicyRulesByPageUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPolicyRulesByPageUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetReauthPolicySetUsingGET1 gets the authentication policy and all rules for a timeout policy rule
*/
func (a *Client) GetReauthPolicySetUsingGET1(params *GetReauthPolicySetUsingGET1Params, opts ...ClientOption) (*GetReauthPolicySetUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetReauthPolicySetUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getReauthPolicySetUsingGET_1",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/policySet/reauth",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetReauthPolicySetUsingGET1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetReauthPolicySetUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getReauthPolicySetUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRuleInPolicySetUsingGET1 gets a rule in a policy
*/
func (a *Client) GetRuleInPolicySetUsingGET1(params *GetRuleInPolicySetUsingGET1Params, opts ...ClientOption) (*GetRuleInPolicySetUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRuleInPolicySetUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getRuleInPolicySetUsingGET_1",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/policySet/{policySetId}/rule/{ruleId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRuleInPolicySetUsingGET1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRuleInPolicySetUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRuleInPolicySetUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateRuleToPolicySetUsingPUT1 updates a rule in a policy
*/
func (a *Client) UpdateRuleToPolicySetUsingPUT1(params *UpdateRuleToPolicySetUsingPUT1Params, opts ...ClientOption) (*UpdateRuleToPolicySetUsingPUT1Created, *UpdateRuleToPolicySetUsingPUT1NoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRuleToPolicySetUsingPUT1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateRuleToPolicySetUsingPUT_1",
		Method:             "PUT",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/policySet/{policySetId}/rule/{ruleId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateRuleToPolicySetUsingPUT1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateRuleToPolicySetUsingPUT1Created:
		return value, nil, nil
	case *UpdateRuleToPolicySetUsingPUT1NoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for policy_set_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}