// Code generated by go-swagger; DO NOT EDIT.

package provisioning_key_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new provisioning key controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for provisioning key controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateProvisioningKeyUsingPOST1(params *CreateProvisioningKeyUsingPOST1Params, opts ...ClientOption) (*CreateProvisioningKeyUsingPOST1OK, *CreateProvisioningKeyUsingPOST1Created, error)

	DeleteProvisioningKeyUsingDELETE1(params *DeleteProvisioningKeyUsingDELETE1Params, opts ...ClientOption) (*DeleteProvisioningKeyUsingDELETE1NoContent, error)

	GetProvisioningKeyForAssociationTypeUsingGET1(params *GetProvisioningKeyForAssociationTypeUsingGET1Params, opts ...ClientOption) (*GetProvisioningKeyForAssociationTypeUsingGET1OK, error)

	GetProvisioningKeyUsingGET1(params *GetProvisioningKeyUsingGET1Params, opts ...ClientOption) (*GetProvisioningKeyUsingGET1OK, error)

	UpdateProvisioningKeyUsingPUT1(params *UpdateProvisioningKeyUsingPUT1Params, opts ...ClientOption) (*UpdateProvisioningKeyUsingPUT1Created, *UpdateProvisioningKeyUsingPUT1NoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateProvisioningKeyUsingPOST1 adds a new provisioning key
*/
func (a *Client) CreateProvisioningKeyUsingPOST1(params *CreateProvisioningKeyUsingPOST1Params, opts ...ClientOption) (*CreateProvisioningKeyUsingPOST1OK, *CreateProvisioningKeyUsingPOST1Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateProvisioningKeyUsingPOST1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "createProvisioningKeyUsingPOST_1",
		Method:             "POST",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/associationType/{associationType}/provisioningKey",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateProvisioningKeyUsingPOST1Reader{formats: a.formats},
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
	case *CreateProvisioningKeyUsingPOST1OK:
		return value, nil, nil
	case *CreateProvisioningKeyUsingPOST1Created:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for provisioning_key_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteProvisioningKeyUsingDELETE1 deletes the provisioning key
*/
func (a *Client) DeleteProvisioningKeyUsingDELETE1(params *DeleteProvisioningKeyUsingDELETE1Params, opts ...ClientOption) (*DeleteProvisioningKeyUsingDELETE1NoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProvisioningKeyUsingDELETE1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteProvisioningKeyUsingDELETE_1",
		Method:             "DELETE",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/associationType/{associationType}/provisioningKey/{provisioningKeyId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteProvisioningKeyUsingDELETE1Reader{formats: a.formats},
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
	success, ok := result.(*DeleteProvisioningKeyUsingDELETE1NoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteProvisioningKeyUsingDELETE_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetProvisioningKeyForAssociationTypeUsingGET1 gets all the configured provisioning key details
*/
func (a *Client) GetProvisioningKeyForAssociationTypeUsingGET1(params *GetProvisioningKeyForAssociationTypeUsingGET1Params, opts ...ClientOption) (*GetProvisioningKeyForAssociationTypeUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProvisioningKeyForAssociationTypeUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getProvisioningKeyForAssociationTypeUsingGET_1",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/associationType/{associationType}/provisioningKey",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProvisioningKeyForAssociationTypeUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*GetProvisioningKeyForAssociationTypeUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getProvisioningKeyForAssociationTypeUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetProvisioningKeyUsingGET1 gets the provisioning key details
*/
func (a *Client) GetProvisioningKeyUsingGET1(params *GetProvisioningKeyUsingGET1Params, opts ...ClientOption) (*GetProvisioningKeyUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProvisioningKeyUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getProvisioningKeyUsingGET_1",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/associationType/{associationType}/provisioningKey/{provisioningKeyId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProvisioningKeyUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*GetProvisioningKeyUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getProvisioningKeyUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateProvisioningKeyUsingPUT1 updates the provisioning key details
*/
func (a *Client) UpdateProvisioningKeyUsingPUT1(params *UpdateProvisioningKeyUsingPUT1Params, opts ...ClientOption) (*UpdateProvisioningKeyUsingPUT1Created, *UpdateProvisioningKeyUsingPUT1NoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProvisioningKeyUsingPUT1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateProvisioningKeyUsingPUT_1",
		Method:             "PUT",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/associationType/{associationType}/provisioningKey/{provisioningKeyId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateProvisioningKeyUsingPUT1Reader{formats: a.formats},
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
	case *UpdateProvisioningKeyUsingPUT1Created:
		return value, nil, nil
	case *UpdateProvisioningKeyUsingPUT1NoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for provisioning_key_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
