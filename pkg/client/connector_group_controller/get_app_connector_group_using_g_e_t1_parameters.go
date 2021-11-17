// Code generated by go-swagger; DO NOT EDIT.

package connector_group_controller

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

// NewGetAppConnectorGroupUsingGET1Params creates a new GetAppConnectorGroupUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAppConnectorGroupUsingGET1Params() *GetAppConnectorGroupUsingGET1Params {
	return &GetAppConnectorGroupUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppConnectorGroupUsingGET1ParamsWithTimeout creates a new GetAppConnectorGroupUsingGET1Params object
// with the ability to set a timeout on a request.
func NewGetAppConnectorGroupUsingGET1ParamsWithTimeout(timeout time.Duration) *GetAppConnectorGroupUsingGET1Params {
	return &GetAppConnectorGroupUsingGET1Params{
		timeout: timeout,
	}
}

// NewGetAppConnectorGroupUsingGET1ParamsWithContext creates a new GetAppConnectorGroupUsingGET1Params object
// with the ability to set a context for a request.
func NewGetAppConnectorGroupUsingGET1ParamsWithContext(ctx context.Context) *GetAppConnectorGroupUsingGET1Params {
	return &GetAppConnectorGroupUsingGET1Params{
		Context: ctx,
	}
}

// NewGetAppConnectorGroupUsingGET1ParamsWithHTTPClient creates a new GetAppConnectorGroupUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetAppConnectorGroupUsingGET1ParamsWithHTTPClient(client *http.Client) *GetAppConnectorGroupUsingGET1Params {
	return &GetAppConnectorGroupUsingGET1Params{
		HTTPClient: client,
	}
}

/* GetAppConnectorGroupUsingGET1Params contains all the parameters to send to the API endpoint
   for the get app connector group using g e t 1 operation.

   Typically these are written to a http.Request.
*/
type GetAppConnectorGroupUsingGET1Params struct {

	/* AppConnectorGroupID.

	   The unique identifier of the App Connector Group.
	*/
	AppConnectorGroupID string

	/* CustomerID.

	   The unique identifier of the ZPA tenant.
	*/
	CustomerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get app connector group using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppConnectorGroupUsingGET1Params) WithDefaults() *GetAppConnectorGroupUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get app connector group using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAppConnectorGroupUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get app connector group using g e t 1 params
func (o *GetAppConnectorGroupUsingGET1Params) WithTimeout(timeout time.Duration) *GetAppConnectorGroupUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get app connector group using g e t 1 params
func (o *GetAppConnectorGroupUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get app connector group using g e t 1 params
func (o *GetAppConnectorGroupUsingGET1Params) WithContext(ctx context.Context) *GetAppConnectorGroupUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get app connector group using g e t 1 params
func (o *GetAppConnectorGroupUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get app connector group using g e t 1 params
func (o *GetAppConnectorGroupUsingGET1Params) WithHTTPClient(client *http.Client) *GetAppConnectorGroupUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get app connector group using g e t 1 params
func (o *GetAppConnectorGroupUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppConnectorGroupID adds the appConnectorGroupID to the get app connector group using g e t 1 params
func (o *GetAppConnectorGroupUsingGET1Params) WithAppConnectorGroupID(appConnectorGroupID string) *GetAppConnectorGroupUsingGET1Params {
	o.SetAppConnectorGroupID(appConnectorGroupID)
	return o
}

// SetAppConnectorGroupID adds the appConnectorGroupId to the get app connector group using g e t 1 params
func (o *GetAppConnectorGroupUsingGET1Params) SetAppConnectorGroupID(appConnectorGroupID string) {
	o.AppConnectorGroupID = appConnectorGroupID
}

// WithCustomerID adds the customerID to the get app connector group using g e t 1 params
func (o *GetAppConnectorGroupUsingGET1Params) WithCustomerID(customerID string) *GetAppConnectorGroupUsingGET1Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the get app connector group using g e t 1 params
func (o *GetAppConnectorGroupUsingGET1Params) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppConnectorGroupUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appConnectorGroupId
	if err := r.SetPathParam("appConnectorGroupId", o.AppConnectorGroupID); err != nil {
		return err
	}

	// path param customerId
	if err := r.SetPathParam("customerId", o.CustomerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
