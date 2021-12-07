// Code generated by go-swagger; DO NOT EDIT.

package idp_controller_v_2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new idp controller v 2 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for idp controller v 2 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAllIdpUsingGET(params *GetAllIdpUsingGETParams, opts ...ClientOption) (*GetAllIdpUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAllIdpUsingGET gets the configured Id p details
*/
func (a *Client) GetAllIdpUsingGET(params *GetAllIdpUsingGETParams, opts ...ClientOption) (*GetAllIdpUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllIdpUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllIdpUsingGET",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v2/admin/customers/{customerId}/idp",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllIdpUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetAllIdpUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllIdpUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}