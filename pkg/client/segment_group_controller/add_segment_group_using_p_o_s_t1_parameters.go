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

	"github.com/haarchri/zpa-go-client/pkg/models"
)

// NewAddSegmentGroupUsingPOST1Params creates a new AddSegmentGroupUsingPOST1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddSegmentGroupUsingPOST1Params() *AddSegmentGroupUsingPOST1Params {
	return &AddSegmentGroupUsingPOST1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddSegmentGroupUsingPOST1ParamsWithTimeout creates a new AddSegmentGroupUsingPOST1Params object
// with the ability to set a timeout on a request.
func NewAddSegmentGroupUsingPOST1ParamsWithTimeout(timeout time.Duration) *AddSegmentGroupUsingPOST1Params {
	return &AddSegmentGroupUsingPOST1Params{
		timeout: timeout,
	}
}

// NewAddSegmentGroupUsingPOST1ParamsWithContext creates a new AddSegmentGroupUsingPOST1Params object
// with the ability to set a context for a request.
func NewAddSegmentGroupUsingPOST1ParamsWithContext(ctx context.Context) *AddSegmentGroupUsingPOST1Params {
	return &AddSegmentGroupUsingPOST1Params{
		Context: ctx,
	}
}

// NewAddSegmentGroupUsingPOST1ParamsWithHTTPClient creates a new AddSegmentGroupUsingPOST1Params object
// with the ability to set a custom HTTPClient for a request.
func NewAddSegmentGroupUsingPOST1ParamsWithHTTPClient(client *http.Client) *AddSegmentGroupUsingPOST1Params {
	return &AddSegmentGroupUsingPOST1Params{
		HTTPClient: client,
	}
}

/* AddSegmentGroupUsingPOST1Params contains all the parameters to send to the API endpoint
   for the add segment group using p o s t 1 operation.

   Typically these are written to a http.Request.
*/
type AddSegmentGroupUsingPOST1Params struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.

	   Format: int64
	*/
	CustomerID int64

	/* SegmentGroup.

	   The object of the Segment Group.
	*/
	SegmentGroup *models.SegmentGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add segment group using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddSegmentGroupUsingPOST1Params) WithDefaults() *AddSegmentGroupUsingPOST1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add segment group using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddSegmentGroupUsingPOST1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add segment group using p o s t 1 params
func (o *AddSegmentGroupUsingPOST1Params) WithTimeout(timeout time.Duration) *AddSegmentGroupUsingPOST1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add segment group using p o s t 1 params
func (o *AddSegmentGroupUsingPOST1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add segment group using p o s t 1 params
func (o *AddSegmentGroupUsingPOST1Params) WithContext(ctx context.Context) *AddSegmentGroupUsingPOST1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add segment group using p o s t 1 params
func (o *AddSegmentGroupUsingPOST1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add segment group using p o s t 1 params
func (o *AddSegmentGroupUsingPOST1Params) WithHTTPClient(client *http.Client) *AddSegmentGroupUsingPOST1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add segment group using p o s t 1 params
func (o *AddSegmentGroupUsingPOST1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the add segment group using p o s t 1 params
func (o *AddSegmentGroupUsingPOST1Params) WithCustomerID(customerID int64) *AddSegmentGroupUsingPOST1Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the add segment group using p o s t 1 params
func (o *AddSegmentGroupUsingPOST1Params) SetCustomerID(customerID int64) {
	o.CustomerID = customerID
}

// WithSegmentGroup adds the segmentGroup to the add segment group using p o s t 1 params
func (o *AddSegmentGroupUsingPOST1Params) WithSegmentGroup(segmentGroup *models.SegmentGroup) *AddSegmentGroupUsingPOST1Params {
	o.SetSegmentGroup(segmentGroup)
	return o
}

// SetSegmentGroup adds the segmentGroup to the add segment group using p o s t 1 params
func (o *AddSegmentGroupUsingPOST1Params) SetSegmentGroup(segmentGroup *models.SegmentGroup) {
	o.SegmentGroup = segmentGroup
}

// WriteToRequest writes these params to a swagger request
func (o *AddSegmentGroupUsingPOST1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", swag.FormatInt64(o.CustomerID)); err != nil {
		return err
	}
	if o.SegmentGroup != nil {
		if err := r.SetBodyParam(o.SegmentGroup); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
