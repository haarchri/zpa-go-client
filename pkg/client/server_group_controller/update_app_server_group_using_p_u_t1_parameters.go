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

	"github.com/haarchri/zpa-go-client/pkg/models"
)

// NewUpdateAppServerGroupUsingPUT1Params creates a new UpdateAppServerGroupUsingPUT1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAppServerGroupUsingPUT1Params() *UpdateAppServerGroupUsingPUT1Params {
	return &UpdateAppServerGroupUsingPUT1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAppServerGroupUsingPUT1ParamsWithTimeout creates a new UpdateAppServerGroupUsingPUT1Params object
// with the ability to set a timeout on a request.
func NewUpdateAppServerGroupUsingPUT1ParamsWithTimeout(timeout time.Duration) *UpdateAppServerGroupUsingPUT1Params {
	return &UpdateAppServerGroupUsingPUT1Params{
		timeout: timeout,
	}
}

// NewUpdateAppServerGroupUsingPUT1ParamsWithContext creates a new UpdateAppServerGroupUsingPUT1Params object
// with the ability to set a context for a request.
func NewUpdateAppServerGroupUsingPUT1ParamsWithContext(ctx context.Context) *UpdateAppServerGroupUsingPUT1Params {
	return &UpdateAppServerGroupUsingPUT1Params{
		Context: ctx,
	}
}

// NewUpdateAppServerGroupUsingPUT1ParamsWithHTTPClient creates a new UpdateAppServerGroupUsingPUT1Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAppServerGroupUsingPUT1ParamsWithHTTPClient(client *http.Client) *UpdateAppServerGroupUsingPUT1Params {
	return &UpdateAppServerGroupUsingPUT1Params{
		HTTPClient: client,
	}
}

/* UpdateAppServerGroupUsingPUT1Params contains all the parameters to send to the API endpoint
   for the update app server group using p u t 1 operation.

   Typically these are written to a http.Request.
*/
type UpdateAppServerGroupUsingPUT1Params struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.
	*/
	CustomerID string

	/* Group.

	   The object of the Server Group.
	*/
	Group *models.ServerGroupDTO

	/* GroupID.

	   The unique identifier of the Server Group.
	*/
	GroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update app server group using p u t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAppServerGroupUsingPUT1Params) WithDefaults() *UpdateAppServerGroupUsingPUT1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update app server group using p u t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAppServerGroupUsingPUT1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update app server group using p u t 1 params
func (o *UpdateAppServerGroupUsingPUT1Params) WithTimeout(timeout time.Duration) *UpdateAppServerGroupUsingPUT1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update app server group using p u t 1 params
func (o *UpdateAppServerGroupUsingPUT1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update app server group using p u t 1 params
func (o *UpdateAppServerGroupUsingPUT1Params) WithContext(ctx context.Context) *UpdateAppServerGroupUsingPUT1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update app server group using p u t 1 params
func (o *UpdateAppServerGroupUsingPUT1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update app server group using p u t 1 params
func (o *UpdateAppServerGroupUsingPUT1Params) WithHTTPClient(client *http.Client) *UpdateAppServerGroupUsingPUT1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update app server group using p u t 1 params
func (o *UpdateAppServerGroupUsingPUT1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the update app server group using p u t 1 params
func (o *UpdateAppServerGroupUsingPUT1Params) WithCustomerID(customerID string) *UpdateAppServerGroupUsingPUT1Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the update app server group using p u t 1 params
func (o *UpdateAppServerGroupUsingPUT1Params) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithGroup adds the group to the update app server group using p u t 1 params
func (o *UpdateAppServerGroupUsingPUT1Params) WithGroup(group *models.ServerGroupDTO) *UpdateAppServerGroupUsingPUT1Params {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the update app server group using p u t 1 params
func (o *UpdateAppServerGroupUsingPUT1Params) SetGroup(group *models.ServerGroupDTO) {
	o.Group = group
}

// WithGroupID adds the groupID to the update app server group using p u t 1 params
func (o *UpdateAppServerGroupUsingPUT1Params) WithGroupID(groupID string) *UpdateAppServerGroupUsingPUT1Params {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the update app server group using p u t 1 params
func (o *UpdateAppServerGroupUsingPUT1Params) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAppServerGroupUsingPUT1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", o.CustomerID); err != nil {
		return err
	}
	if o.Group != nil {
		if err := r.SetBodyParam(o.Group); err != nil {
			return err
		}
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
