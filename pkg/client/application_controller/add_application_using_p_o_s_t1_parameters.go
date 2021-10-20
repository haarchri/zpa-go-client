// Code generated by go-swagger; DO NOT EDIT.

package application_controller

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
	"github.com/go-openapi/swag"

	"github.com/haarchri/zpa-go-client/pkg/models"
)

// NewAddApplicationUsingPOST1Params creates a new AddApplicationUsingPOST1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddApplicationUsingPOST1Params() *AddApplicationUsingPOST1Params {
	return &AddApplicationUsingPOST1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddApplicationUsingPOST1ParamsWithTimeout creates a new AddApplicationUsingPOST1Params object
// with the ability to set a timeout on a request.
func NewAddApplicationUsingPOST1ParamsWithTimeout(timeout time.Duration) *AddApplicationUsingPOST1Params {
	return &AddApplicationUsingPOST1Params{
		timeout: timeout,
	}
}

// NewAddApplicationUsingPOST1ParamsWithContext creates a new AddApplicationUsingPOST1Params object
// with the ability to set a context for a request.
func NewAddApplicationUsingPOST1ParamsWithContext(ctx context.Context) *AddApplicationUsingPOST1Params {
	return &AddApplicationUsingPOST1Params{
		Context: ctx,
	}
}

// NewAddApplicationUsingPOST1ParamsWithHTTPClient creates a new AddApplicationUsingPOST1Params object
// with the ability to set a custom HTTPClient for a request.
func NewAddApplicationUsingPOST1ParamsWithHTTPClient(client *http.Client) *AddApplicationUsingPOST1Params {
	return &AddApplicationUsingPOST1Params{
		HTTPClient: client,
	}
}

/* AddApplicationUsingPOST1Params contains all the parameters to send to the API endpoint
   for the add application using p o s t 1 operation.

   Typically these are written to a http.Request.
*/
type AddApplicationUsingPOST1Params struct {

	/* Application.

	   The object of the Application Segment.
	*/
	Application *models.ApplicationResource

	/* CustomerID.

	   The unique identifier of the ZPA tenant.

	   Format: int64
	*/
	CustomerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add application using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddApplicationUsingPOST1Params) WithDefaults() *AddApplicationUsingPOST1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add application using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddApplicationUsingPOST1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add application using p o s t 1 params
func (o *AddApplicationUsingPOST1Params) WithTimeout(timeout time.Duration) *AddApplicationUsingPOST1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add application using p o s t 1 params
func (o *AddApplicationUsingPOST1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add application using p o s t 1 params
func (o *AddApplicationUsingPOST1Params) WithContext(ctx context.Context) *AddApplicationUsingPOST1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add application using p o s t 1 params
func (o *AddApplicationUsingPOST1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add application using p o s t 1 params
func (o *AddApplicationUsingPOST1Params) WithHTTPClient(client *http.Client) *AddApplicationUsingPOST1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add application using p o s t 1 params
func (o *AddApplicationUsingPOST1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplication adds the application to the add application using p o s t 1 params
func (o *AddApplicationUsingPOST1Params) WithApplication(application *models.ApplicationResource) *AddApplicationUsingPOST1Params {
	o.SetApplication(application)
	return o
}

// SetApplication adds the application to the add application using p o s t 1 params
func (o *AddApplicationUsingPOST1Params) SetApplication(application *models.ApplicationResource) {
	o.Application = application
}

// WithCustomerID adds the customerID to the add application using p o s t 1 params
func (o *AddApplicationUsingPOST1Params) WithCustomerID(customerID int64) *AddApplicationUsingPOST1Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the add application using p o s t 1 params
func (o *AddApplicationUsingPOST1Params) SetCustomerID(customerID int64) {
	o.CustomerID = customerID
}

// WriteToRequest writes these params to a swagger request
func (o *AddApplicationUsingPOST1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Application != nil {
		if err := r.SetBodyParam(o.Application); err != nil {
			return err
		}
	}

	// path param customerId
	if err := r.SetPathParam("customerId", swag.FormatInt64(o.CustomerID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}