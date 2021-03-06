// Code generated by go-swagger; DO NOT EDIT.

package idp_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new idp controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for idp controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetIdpByIDUsingGET1(params *GetIdpByIDUsingGET1Params, opts ...ClientOption) (*GetIdpByIDUsingGET1OK, error)

	GetIdpUsingGET1(params *GetIdpUsingGET1Params, opts ...ClientOption) (*GetIdpUsingGET1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetIdpByIDUsingGET1 gets all the Id p details
*/
func (a *Client) GetIdpByIDUsingGET1(params *GetIdpByIDUsingGET1Params, opts ...ClientOption) (*GetIdpByIDUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIdpByIDUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getIdpByIdUsingGET_1",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/idp/{idpId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIdpByIDUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*GetIdpByIDUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getIdpByIdUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetIdpUsingGET1 gets the configured Id p details this API will be deprecated in a future release
*/
func (a *Client) GetIdpUsingGET1(params *GetIdpUsingGET1Params, opts ...ClientOption) (*GetIdpUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIdpUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getIdpUsingGET_1",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/idp",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIdpUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*GetIdpUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getIdpUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
