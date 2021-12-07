// Code generated by go-swagger; DO NOT EDIT.

package ba_certificate_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ba certificate controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ba certificate controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAllIssuedCertsUsingGET1(params *GetAllIssuedCertsUsingGET1Params, opts ...ClientOption) (*GetAllIssuedCertsUsingGET1OK, error)

	GetUsingGET1(params *GetUsingGET1Params, opts ...ClientOption) (*GetUsingGET1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAllIssuedCertsUsingGET1 gets all the issued certificates this API will be deprecated in a future release
*/
func (a *Client) GetAllIssuedCertsUsingGET1(params *GetAllIssuedCertsUsingGET1Params, opts ...ClientOption) (*GetAllIssuedCertsUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllIssuedCertsUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllIssuedCertsUsingGET_1",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/clientlessCertificate/issued",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllIssuedCertsUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*GetAllIssuedCertsUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllIssuedCertsUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsingGET1 gets the details of the browser access certificate
*/
func (a *Client) GetUsingGET1(params *GetUsingGET1Params, opts ...ClientOption) (*GetUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUsingGET_1",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/clientlessCertificate/{certificateId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*GetUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
