// Code generated by go-swagger; DO NOT EDIT.

package saml_attr_controller_v_2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new saml attr controller v 2 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for saml attr controller v 2 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAllAttributesByIdpIDAndPageUsingGET(params *GetAllAttributesByIdpIDAndPageUsingGETParams, opts ...ClientOption) (*GetAllAttributesByIdpIDAndPageUsingGETOK, error)

	GetAllAttributesByPageUsingGET(params *GetAllAttributesByPageUsingGETParams, opts ...ClientOption) (*GetAllAttributesByPageUsingGETOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAllAttributesByIdpIDAndPageUsingGET gets all attributes configured for a given customer
*/
func (a *Client) GetAllAttributesByIdpIDAndPageUsingGET(params *GetAllAttributesByIdpIDAndPageUsingGETParams, opts ...ClientOption) (*GetAllAttributesByIdpIDAndPageUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllAttributesByIdpIDAndPageUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllAttributesByIdpIdAndPageUsingGET",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v2/admin/customers/{customerId}/samlAttribute/idp/{idpId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllAttributesByIdpIDAndPageUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetAllAttributesByIdpIDAndPageUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllAttributesByIdpIdAndPageUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllAttributesByPageUsingGET gets all s a m l attributes by page
*/
func (a *Client) GetAllAttributesByPageUsingGET(params *GetAllAttributesByPageUsingGETParams, opts ...ClientOption) (*GetAllAttributesByPageUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllAttributesByPageUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllAttributesByPageUsingGET",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v2/admin/customers/{customerId}/samlAttribute",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllAttributesByPageUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetAllAttributesByPageUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllAttributesByPageUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
