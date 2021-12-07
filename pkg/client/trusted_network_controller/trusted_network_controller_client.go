// Code generated by go-swagger; DO NOT EDIT.

package trusted_network_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new trusted network controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for trusted network controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAllTrustedNetworksUsingGET1(params *GetAllTrustedNetworksUsingGET1Params, opts ...ClientOption) (*GetAllTrustedNetworksUsingGET1OK, error)

	GetTrustedNetworkUsingGET1(params *GetTrustedNetworkUsingGET1Params, opts ...ClientOption) (*GetTrustedNetworkUsingGET1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAllTrustedNetworksUsingGET1 gets all the trusted networks this API will be deprecated in a future release
*/
func (a *Client) GetAllTrustedNetworksUsingGET1(params *GetAllTrustedNetworksUsingGET1Params, opts ...ClientOption) (*GetAllTrustedNetworksUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllTrustedNetworksUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllTrustedNetworksUsingGET_1",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/network",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllTrustedNetworksUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*GetAllTrustedNetworksUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllTrustedNetworksUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTrustedNetworkUsingGET1 gets the trusted networks
*/
func (a *Client) GetTrustedNetworkUsingGET1(params *GetTrustedNetworkUsingGET1Params, opts ...ClientOption) (*GetTrustedNetworkUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTrustedNetworkUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTrustedNetworkUsingGET_1",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/network/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTrustedNetworkUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*GetTrustedNetworkUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTrustedNetworkUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
