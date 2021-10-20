// Code generated by go-swagger; DO NOT EDIT.

package segment_group_controller

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
)

// NewDeleteSegmentGroupUsingDELETE1Params creates a new DeleteSegmentGroupUsingDELETE1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSegmentGroupUsingDELETE1Params() *DeleteSegmentGroupUsingDELETE1Params {
	return &DeleteSegmentGroupUsingDELETE1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSegmentGroupUsingDELETE1ParamsWithTimeout creates a new DeleteSegmentGroupUsingDELETE1Params object
// with the ability to set a timeout on a request.
func NewDeleteSegmentGroupUsingDELETE1ParamsWithTimeout(timeout time.Duration) *DeleteSegmentGroupUsingDELETE1Params {
	return &DeleteSegmentGroupUsingDELETE1Params{
		timeout: timeout,
	}
}

// NewDeleteSegmentGroupUsingDELETE1ParamsWithContext creates a new DeleteSegmentGroupUsingDELETE1Params object
// with the ability to set a context for a request.
func NewDeleteSegmentGroupUsingDELETE1ParamsWithContext(ctx context.Context) *DeleteSegmentGroupUsingDELETE1Params {
	return &DeleteSegmentGroupUsingDELETE1Params{
		Context: ctx,
	}
}

// NewDeleteSegmentGroupUsingDELETE1ParamsWithHTTPClient creates a new DeleteSegmentGroupUsingDELETE1Params object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSegmentGroupUsingDELETE1ParamsWithHTTPClient(client *http.Client) *DeleteSegmentGroupUsingDELETE1Params {
	return &DeleteSegmentGroupUsingDELETE1Params{
		HTTPClient: client,
	}
}

/* DeleteSegmentGroupUsingDELETE1Params contains all the parameters to send to the API endpoint
   for the delete segment group using d e l e t e 1 operation.

   Typically these are written to a http.Request.
*/
type DeleteSegmentGroupUsingDELETE1Params struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.

	   Format: int64
	*/
	CustomerID int64

	/* SegmentGroupID.

	   The unique identifier of the Segment Group.

	   Format: int64
	*/
	SegmentGroupID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete segment group using d e l e t e 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSegmentGroupUsingDELETE1Params) WithDefaults() *DeleteSegmentGroupUsingDELETE1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete segment group using d e l e t e 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSegmentGroupUsingDELETE1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete segment group using d e l e t e 1 params
func (o *DeleteSegmentGroupUsingDELETE1Params) WithTimeout(timeout time.Duration) *DeleteSegmentGroupUsingDELETE1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete segment group using d e l e t e 1 params
func (o *DeleteSegmentGroupUsingDELETE1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete segment group using d e l e t e 1 params
func (o *DeleteSegmentGroupUsingDELETE1Params) WithContext(ctx context.Context) *DeleteSegmentGroupUsingDELETE1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete segment group using d e l e t e 1 params
func (o *DeleteSegmentGroupUsingDELETE1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete segment group using d e l e t e 1 params
func (o *DeleteSegmentGroupUsingDELETE1Params) WithHTTPClient(client *http.Client) *DeleteSegmentGroupUsingDELETE1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete segment group using d e l e t e 1 params
func (o *DeleteSegmentGroupUsingDELETE1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the delete segment group using d e l e t e 1 params
func (o *DeleteSegmentGroupUsingDELETE1Params) WithCustomerID(customerID int64) *DeleteSegmentGroupUsingDELETE1Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the delete segment group using d e l e t e 1 params
func (o *DeleteSegmentGroupUsingDELETE1Params) SetCustomerID(customerID int64) {
	o.CustomerID = customerID
}

// WithSegmentGroupID adds the segmentGroupID to the delete segment group using d e l e t e 1 params
func (o *DeleteSegmentGroupUsingDELETE1Params) WithSegmentGroupID(segmentGroupID int64) *DeleteSegmentGroupUsingDELETE1Params {
	o.SetSegmentGroupID(segmentGroupID)
	return o
}

// SetSegmentGroupID adds the segmentGroupId to the delete segment group using d e l e t e 1 params
func (o *DeleteSegmentGroupUsingDELETE1Params) SetSegmentGroupID(segmentGroupID int64) {
	o.SegmentGroupID = segmentGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSegmentGroupUsingDELETE1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", swag.FormatInt64(o.CustomerID)); err != nil {
		return err
	}

	// path param segmentGroupId
	if err := r.SetPathParam("segmentGroupId", swag.FormatInt64(o.SegmentGroupID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
