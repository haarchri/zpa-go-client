// Code generated by go-swagger; DO NOT EDIT.

package segment_group_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new segment group controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for segment group controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddSegmentGroupUsingPOST1(params *AddSegmentGroupUsingPOST1Params, opts ...ClientOption) (*AddSegmentGroupUsingPOST1Created, error)

	DeleteSegmentGroupUsingDELETE1(params *DeleteSegmentGroupUsingDELETE1Params, opts ...ClientOption) (*DeleteSegmentGroupUsingDELETE1NoContent, error)

	GetAllSegmentGroupsUsingGET1(params *GetAllSegmentGroupsUsingGET1Params, opts ...ClientOption) (*GetAllSegmentGroupsUsingGET1OK, error)

	GetSegmentGroupUsingGET1(params *GetSegmentGroupUsingGET1Params, opts ...ClientOption) (*GetSegmentGroupUsingGET1OK, error)

	UpdateSegmentGroupUsingPUT1(params *UpdateSegmentGroupUsingPUT1Params, opts ...ClientOption) (*UpdateSegmentGroupUsingPUT1Created, *UpdateSegmentGroupUsingPUT1NoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddSegmentGroupUsingPOST1 adds a new segment group
*/
func (a *Client) AddSegmentGroupUsingPOST1(params *AddSegmentGroupUsingPOST1Params, opts ...ClientOption) (*AddSegmentGroupUsingPOST1Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddSegmentGroupUsingPOST1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "addSegmentGroupUsingPOST_1",
		Method:             "POST",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/segmentGroup",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddSegmentGroupUsingPOST1Reader{formats: a.formats},
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
	success, ok := result.(*AddSegmentGroupUsingPOST1Created)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addSegmentGroupUsingPOST_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteSegmentGroupUsingDELETE1 deletes a segment group
*/
func (a *Client) DeleteSegmentGroupUsingDELETE1(params *DeleteSegmentGroupUsingDELETE1Params, opts ...ClientOption) (*DeleteSegmentGroupUsingDELETE1NoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSegmentGroupUsingDELETE1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteSegmentGroupUsingDELETE_1",
		Method:             "DELETE",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/segmentGroup/{segmentGroupId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSegmentGroupUsingDELETE1Reader{formats: a.formats},
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
	success, ok := result.(*DeleteSegmentGroupUsingDELETE1NoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteSegmentGroupUsingDELETE_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllSegmentGroupsUsingGET1 gets all the configured segment groups
*/
func (a *Client) GetAllSegmentGroupsUsingGET1(params *GetAllSegmentGroupsUsingGET1Params, opts ...ClientOption) (*GetAllSegmentGroupsUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllSegmentGroupsUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllSegmentGroupsUsingGET_1",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/segmentGroup",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllSegmentGroupsUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*GetAllSegmentGroupsUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllSegmentGroupsUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSegmentGroupUsingGET1 gets the segment group details
*/
func (a *Client) GetSegmentGroupUsingGET1(params *GetSegmentGroupUsingGET1Params, opts ...ClientOption) (*GetSegmentGroupUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSegmentGroupUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSegmentGroupUsingGET_1",
		Method:             "GET",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/segmentGroup/{segmentGroupId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSegmentGroupUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*GetSegmentGroupUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSegmentGroupUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateSegmentGroupUsingPUT1 updates a segment group
*/
func (a *Client) UpdateSegmentGroupUsingPUT1(params *UpdateSegmentGroupUsingPUT1Params, opts ...ClientOption) (*UpdateSegmentGroupUsingPUT1Created, *UpdateSegmentGroupUsingPUT1NoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSegmentGroupUsingPUT1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateSegmentGroupUsingPUT_1",
		Method:             "PUT",
		PathPattern:        "/mgmtconfig/v1/admin/customers/{customerId}/segmentGroup/{segmentGroupId}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSegmentGroupUsingPUT1Reader{formats: a.formats},
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
	case *UpdateSegmentGroupUsingPUT1Created:
		return value, nil, nil
	case *UpdateSegmentGroupUsingPUT1NoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for segment_group_controller: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
