// Code generated by go-swagger; DO NOT EDIT.

package server_group_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteAppServerGroupUsingDELETE1Params creates a new DeleteAppServerGroupUsingDELETE1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAppServerGroupUsingDELETE1Params() *DeleteAppServerGroupUsingDELETE1Params {
	return &DeleteAppServerGroupUsingDELETE1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAppServerGroupUsingDELETE1ParamsWithTimeout creates a new DeleteAppServerGroupUsingDELETE1Params object
// with the ability to set a timeout on a request.
func NewDeleteAppServerGroupUsingDELETE1ParamsWithTimeout(timeout time.Duration) *DeleteAppServerGroupUsingDELETE1Params {
	return &DeleteAppServerGroupUsingDELETE1Params{
		timeout: timeout,
	}
}

// NewDeleteAppServerGroupUsingDELETE1ParamsWithContext creates a new DeleteAppServerGroupUsingDELETE1Params object
// with the ability to set a context for a request.
func NewDeleteAppServerGroupUsingDELETE1ParamsWithContext(ctx context.Context) *DeleteAppServerGroupUsingDELETE1Params {
	return &DeleteAppServerGroupUsingDELETE1Params{
		Context: ctx,
	}
}

// NewDeleteAppServerGroupUsingDELETE1ParamsWithHTTPClient creates a new DeleteAppServerGroupUsingDELETE1Params object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAppServerGroupUsingDELETE1ParamsWithHTTPClient(client *http.Client) *DeleteAppServerGroupUsingDELETE1Params {
	return &DeleteAppServerGroupUsingDELETE1Params{
		HTTPClient: client,
	}
}

/* DeleteAppServerGroupUsingDELETE1Params contains all the parameters to send to the API endpoint
   for the delete app server group using d e l e t e 1 operation.

   Typically these are written to a http.Request.
*/
type DeleteAppServerGroupUsingDELETE1Params struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.
	*/
	CustomerID string

	/* GroupID.

	   The unique identifier of the Server Group.
	*/
	GroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete app server group using d e l e t e 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAppServerGroupUsingDELETE1Params) WithDefaults() *DeleteAppServerGroupUsingDELETE1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete app server group using d e l e t e 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAppServerGroupUsingDELETE1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete app server group using d e l e t e 1 params
func (o *DeleteAppServerGroupUsingDELETE1Params) WithTimeout(timeout time.Duration) *DeleteAppServerGroupUsingDELETE1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete app server group using d e l e t e 1 params
func (o *DeleteAppServerGroupUsingDELETE1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete app server group using d e l e t e 1 params
func (o *DeleteAppServerGroupUsingDELETE1Params) WithContext(ctx context.Context) *DeleteAppServerGroupUsingDELETE1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete app server group using d e l e t e 1 params
func (o *DeleteAppServerGroupUsingDELETE1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete app server group using d e l e t e 1 params
func (o *DeleteAppServerGroupUsingDELETE1Params) WithHTTPClient(client *http.Client) *DeleteAppServerGroupUsingDELETE1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete app server group using d e l e t e 1 params
func (o *DeleteAppServerGroupUsingDELETE1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the delete app server group using d e l e t e 1 params
func (o *DeleteAppServerGroupUsingDELETE1Params) WithCustomerID(customerID string) *DeleteAppServerGroupUsingDELETE1Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the delete app server group using d e l e t e 1 params
func (o *DeleteAppServerGroupUsingDELETE1Params) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithGroupID adds the groupID to the delete app server group using d e l e t e 1 params
func (o *DeleteAppServerGroupUsingDELETE1Params) WithGroupID(groupID string) *DeleteAppServerGroupUsingDELETE1Params {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the delete app server group using d e l e t e 1 params
func (o *DeleteAppServerGroupUsingDELETE1Params) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAppServerGroupUsingDELETE1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", o.CustomerID); err != nil {
		return err
	}

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
