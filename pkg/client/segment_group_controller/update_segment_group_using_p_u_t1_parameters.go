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

	"github.com/haarchri/zpa-go-client/pkg/models"
)

// NewUpdateSegmentGroupUsingPUT1Params creates a new UpdateSegmentGroupUsingPUT1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateSegmentGroupUsingPUT1Params() *UpdateSegmentGroupUsingPUT1Params {
	return &UpdateSegmentGroupUsingPUT1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSegmentGroupUsingPUT1ParamsWithTimeout creates a new UpdateSegmentGroupUsingPUT1Params object
// with the ability to set a timeout on a request.
func NewUpdateSegmentGroupUsingPUT1ParamsWithTimeout(timeout time.Duration) *UpdateSegmentGroupUsingPUT1Params {
	return &UpdateSegmentGroupUsingPUT1Params{
		timeout: timeout,
	}
}

// NewUpdateSegmentGroupUsingPUT1ParamsWithContext creates a new UpdateSegmentGroupUsingPUT1Params object
// with the ability to set a context for a request.
func NewUpdateSegmentGroupUsingPUT1ParamsWithContext(ctx context.Context) *UpdateSegmentGroupUsingPUT1Params {
	return &UpdateSegmentGroupUsingPUT1Params{
		Context: ctx,
	}
}

// NewUpdateSegmentGroupUsingPUT1ParamsWithHTTPClient creates a new UpdateSegmentGroupUsingPUT1Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateSegmentGroupUsingPUT1ParamsWithHTTPClient(client *http.Client) *UpdateSegmentGroupUsingPUT1Params {
	return &UpdateSegmentGroupUsingPUT1Params{
		HTTPClient: client,
	}
}

/* UpdateSegmentGroupUsingPUT1Params contains all the parameters to send to the API endpoint
   for the update segment group using p u t 1 operation.

   Typically these are written to a http.Request.
*/
type UpdateSegmentGroupUsingPUT1Params struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.
	*/
	CustomerID string

	/* SegmentGroup.

	   The object of the Segment Group.
	*/
	SegmentGroup *models.SegmentGroup

	/* SegmentGroupID.

	   The unique identifier of the Segment Group.
	*/
	SegmentGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update segment group using p u t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSegmentGroupUsingPUT1Params) WithDefaults() *UpdateSegmentGroupUsingPUT1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update segment group using p u t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSegmentGroupUsingPUT1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update segment group using p u t 1 params
func (o *UpdateSegmentGroupUsingPUT1Params) WithTimeout(timeout time.Duration) *UpdateSegmentGroupUsingPUT1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update segment group using p u t 1 params
func (o *UpdateSegmentGroupUsingPUT1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update segment group using p u t 1 params
func (o *UpdateSegmentGroupUsingPUT1Params) WithContext(ctx context.Context) *UpdateSegmentGroupUsingPUT1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update segment group using p u t 1 params
func (o *UpdateSegmentGroupUsingPUT1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update segment group using p u t 1 params
func (o *UpdateSegmentGroupUsingPUT1Params) WithHTTPClient(client *http.Client) *UpdateSegmentGroupUsingPUT1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update segment group using p u t 1 params
func (o *UpdateSegmentGroupUsingPUT1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the update segment group using p u t 1 params
func (o *UpdateSegmentGroupUsingPUT1Params) WithCustomerID(customerID string) *UpdateSegmentGroupUsingPUT1Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the update segment group using p u t 1 params
func (o *UpdateSegmentGroupUsingPUT1Params) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithSegmentGroup adds the segmentGroup to the update segment group using p u t 1 params
func (o *UpdateSegmentGroupUsingPUT1Params) WithSegmentGroup(segmentGroup *models.SegmentGroup) *UpdateSegmentGroupUsingPUT1Params {
	o.SetSegmentGroup(segmentGroup)
	return o
}

// SetSegmentGroup adds the segmentGroup to the update segment group using p u t 1 params
func (o *UpdateSegmentGroupUsingPUT1Params) SetSegmentGroup(segmentGroup *models.SegmentGroup) {
	o.SegmentGroup = segmentGroup
}

// WithSegmentGroupID adds the segmentGroupID to the update segment group using p u t 1 params
func (o *UpdateSegmentGroupUsingPUT1Params) WithSegmentGroupID(segmentGroupID string) *UpdateSegmentGroupUsingPUT1Params {
	o.SetSegmentGroupID(segmentGroupID)
	return o
}

// SetSegmentGroupID adds the segmentGroupId to the update segment group using p u t 1 params
func (o *UpdateSegmentGroupUsingPUT1Params) SetSegmentGroupID(segmentGroupID string) {
	o.SegmentGroupID = segmentGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSegmentGroupUsingPUT1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", o.CustomerID); err != nil {
		return err
	}
	if o.SegmentGroup != nil {
		if err := r.SetBodyParam(o.SegmentGroup); err != nil {
			return err
		}
	}

	// path param segmentGroupId
	if err := r.SetPathParam("segmentGroupId", o.SegmentGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
