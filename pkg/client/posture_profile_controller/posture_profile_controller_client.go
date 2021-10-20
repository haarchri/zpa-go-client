// Code generated by go-swagger; DO NOT EDIT.

package posture_profile_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new posture profile controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for posture profile controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAllAttributesUsingGET1(params *GetAllAttributesUsingGET1Params, opts ...ClientOption) (*GetAllAttributesUsingGET1OK, error)

	GetPostureProfileUsingGET1(params *GetPostureProfileUsingGET1Params, opts ...ClientOption) (*GetPostureProfileUsingGET1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAllAttributesUsingGET1 gets all posture profiles
*/
func (a *Client) GetAllAttributesUsingGET1(params *GetAllAttributesUsingGET1Params, opts ...ClientOption) (*GetAllAttributesUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllAttributesUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllAttributesUsingGET_1",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/posture",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllAttributesUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*GetAllAttributesUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllAttributesUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPostureProfileUsingGET1 gets the configured profile posture
*/
func (a *Client) GetPostureProfileUsingGET1(params *GetPostureProfileUsingGET1Params, opts ...ClientOption) (*GetPostureProfileUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPostureProfileUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPostureProfileUsingGET_1",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/posture/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPostureProfileUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*GetPostureProfileUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPostureProfileUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}