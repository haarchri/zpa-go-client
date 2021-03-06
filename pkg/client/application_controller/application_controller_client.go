// Code generated by go-swagger; DO NOT EDIT.

package application_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new application controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for application controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddApplicationUsingPOST1(params *AddApplicationUsingPOST1Params, opts ...ClientOption) (*AddApplicationUsingPOST1Created, error)

	DeleteApplicationUsingDELETE1(params *DeleteApplicationUsingDELETE1Params, opts ...ClientOption) (*DeleteApplicationUsingDELETE1NoContent, error)

	GetAllApplicationsUsingGET3(params *GetAllApplicationsUsingGET3Params, opts ...ClientOption) (*GetAllApplicationsUsingGET3OK, error)

	GetApplicationUsingGET1(params *GetApplicationUsingGET1Params, opts ...ClientOption) (*GetApplicationUsingGET1OK, error)

	UpdateApplicationV2UsingPUT1(params *UpdateApplicationV2UsingPUT1Params, opts ...ClientOption) (*UpdateApplicationV2UsingPUT1Created, *UpdateApplicationV2UsingPUT1NoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddApplicationUsingPOST1 adds a new application segment
*/
func (a *Client) AddApplicationUsingPOST1(params *AddApplicationUsingPOST1Params, opts ...ClientOption) (*AddApplicationUsingPOST1Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddApplicationUsingPOST1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "addApplicationUsingPOST_1",
		Method:             "POST",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/application",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddApplicationUsingPOST1Reader{formats: a.formats},
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
	success, ok := result.(*AddApplicationUsingPOST1Created)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addApplicationUsingPOST_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteApplicationUsingDELETE1 deletes an application segment
*/
func (a *Client) DeleteApplicationUsingDELETE1(params *DeleteApplicationUsingDELETE1Params, opts ...ClientOption) (*DeleteApplicationUsingDELETE1NoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteApplicationUsingDELETE1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteApplicationUsingDELETE_1",
		Method:             "DELETE",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/application/{applicationId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteApplicationUsingDELETE1Reader{formats: a.formats},
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
	success, ok := result.(*DeleteApplicationUsingDELETE1NoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteApplicationUsingDELETE_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllApplicationsUsingGET3 gets all configured application segments
*/
func (a *Client) GetAllApplicationsUsingGET3(params *GetAllApplicationsUsingGET3Params, opts ...ClientOption) (*GetAllApplicationsUsingGET3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllApplicationsUsingGET3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllApplicationsUsingGET_3",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/application",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllApplicationsUsingGET3Reader{formats: a.formats},
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
	success, ok := result.(*GetAllApplicationsUsingGET3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllApplicationsUsingGET_3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetApplicationUsingGET1 gets the application segment details
*/
func (a *Client) GetApplicationUsingGET1(params *GetApplicationUsingGET1Params, opts ...ClientOption) (*GetApplicationUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetApplicationUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getApplicationUsingGET_1",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/application/{applicationId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetApplicationUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*GetApplicationUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getApplicationUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateApplicationV2UsingPUT1 updates the application segment details
*/
func (a *Client) UpdateApplicationV2UsingPUT1(params *UpdateApplicationV2UsingPUT1Params, opts ...ClientOption) (*UpdateApplicationV2UsingPUT1Created, *UpdateApplicationV2UsingPUT1NoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateApplicationV2UsingPUT1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateApplicationV2UsingPUT_1",
		Method:             "PUT",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/application/{applicationId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateApplicationV2UsingPUT1Reader{formats: a.formats},
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
	case *UpdateApplicationV2UsingPUT1Created:
		return value, nil, nil
	case *UpdateApplicationV2UsingPUT1NoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for application_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
