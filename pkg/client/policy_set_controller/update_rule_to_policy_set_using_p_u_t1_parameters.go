// Code generated by go-swagger; DO NOT EDIT.

package policy_set_controller

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

// NewUpdateRuleToPolicySetUsingPUT1Params creates a new UpdateRuleToPolicySetUsingPUT1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRuleToPolicySetUsingPUT1Params() *UpdateRuleToPolicySetUsingPUT1Params {
	return &UpdateRuleToPolicySetUsingPUT1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRuleToPolicySetUsingPUT1ParamsWithTimeout creates a new UpdateRuleToPolicySetUsingPUT1Params object
// with the ability to set a timeout on a request.
func NewUpdateRuleToPolicySetUsingPUT1ParamsWithTimeout(timeout time.Duration) *UpdateRuleToPolicySetUsingPUT1Params {
	return &UpdateRuleToPolicySetUsingPUT1Params{
		timeout: timeout,
	}
}

// NewUpdateRuleToPolicySetUsingPUT1ParamsWithContext creates a new UpdateRuleToPolicySetUsingPUT1Params object
// with the ability to set a context for a request.
func NewUpdateRuleToPolicySetUsingPUT1ParamsWithContext(ctx context.Context) *UpdateRuleToPolicySetUsingPUT1Params {
	return &UpdateRuleToPolicySetUsingPUT1Params{
		Context: ctx,
	}
}

// NewUpdateRuleToPolicySetUsingPUT1ParamsWithHTTPClient creates a new UpdateRuleToPolicySetUsingPUT1Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRuleToPolicySetUsingPUT1ParamsWithHTTPClient(client *http.Client) *UpdateRuleToPolicySetUsingPUT1Params {
	return &UpdateRuleToPolicySetUsingPUT1Params{
		HTTPClient: client,
	}
}

/* UpdateRuleToPolicySetUsingPUT1Params contains all the parameters to send to the API endpoint
   for the update rule to policy set using p u t 1 operation.

   Typically these are written to a http.Request.
*/
type UpdateRuleToPolicySetUsingPUT1Params struct {

	/* CustomerID.

	   The unique identifier of the ZPA tenant.
	*/
	CustomerID string

	/* PolicySetID.

	   The unique identifier of the policy.
	*/
	PolicySetID string

	/* Rule.

	   The object of the rule in a policy.
	*/
	Rule *models.PolicyRule

	/* RuleID.

	   The unique identifier of a rule in a policy.
	*/
	RuleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update rule to policy set using p u t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRuleToPolicySetUsingPUT1Params) WithDefaults() *UpdateRuleToPolicySetUsingPUT1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update rule to policy set using p u t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRuleToPolicySetUsingPUT1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update rule to policy set using p u t 1 params
func (o *UpdateRuleToPolicySetUsingPUT1Params) WithTimeout(timeout time.Duration) *UpdateRuleToPolicySetUsingPUT1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update rule to policy set using p u t 1 params
func (o *UpdateRuleToPolicySetUsingPUT1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update rule to policy set using p u t 1 params
func (o *UpdateRuleToPolicySetUsingPUT1Params) WithContext(ctx context.Context) *UpdateRuleToPolicySetUsingPUT1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update rule to policy set using p u t 1 params
func (o *UpdateRuleToPolicySetUsingPUT1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update rule to policy set using p u t 1 params
func (o *UpdateRuleToPolicySetUsingPUT1Params) WithHTTPClient(client *http.Client) *UpdateRuleToPolicySetUsingPUT1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update rule to policy set using p u t 1 params
func (o *UpdateRuleToPolicySetUsingPUT1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerID adds the customerID to the update rule to policy set using p u t 1 params
func (o *UpdateRuleToPolicySetUsingPUT1Params) WithCustomerID(customerID string) *UpdateRuleToPolicySetUsingPUT1Params {
	o.SetCustomerID(customerID)
	return o
}

// SetCustomerID adds the customerId to the update rule to policy set using p u t 1 params
func (o *UpdateRuleToPolicySetUsingPUT1Params) SetCustomerID(customerID string) {
	o.CustomerID = customerID
}

// WithPolicySetID adds the policySetID to the update rule to policy set using p u t 1 params
func (o *UpdateRuleToPolicySetUsingPUT1Params) WithPolicySetID(policySetID string) *UpdateRuleToPolicySetUsingPUT1Params {
	o.SetPolicySetID(policySetID)
	return o
}

// SetPolicySetID adds the policySetId to the update rule to policy set using p u t 1 params
func (o *UpdateRuleToPolicySetUsingPUT1Params) SetPolicySetID(policySetID string) {
	o.PolicySetID = policySetID
}

// WithRule adds the rule to the update rule to policy set using p u t 1 params
func (o *UpdateRuleToPolicySetUsingPUT1Params) WithRule(rule *models.PolicyRule) *UpdateRuleToPolicySetUsingPUT1Params {
	o.SetRule(rule)
	return o
}

// SetRule adds the rule to the update rule to policy set using p u t 1 params
func (o *UpdateRuleToPolicySetUsingPUT1Params) SetRule(rule *models.PolicyRule) {
	o.Rule = rule
}

// WithRuleID adds the ruleID to the update rule to policy set using p u t 1 params
func (o *UpdateRuleToPolicySetUsingPUT1Params) WithRuleID(ruleID string) *UpdateRuleToPolicySetUsingPUT1Params {
	o.SetRuleID(ruleID)
	return o
}

// SetRuleID adds the ruleId to the update rule to policy set using p u t 1 params
func (o *UpdateRuleToPolicySetUsingPUT1Params) SetRuleID(ruleID string) {
	o.RuleID = ruleID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRuleToPolicySetUsingPUT1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customerId
	if err := r.SetPathParam("customerId", o.CustomerID); err != nil {
		return err
	}

	// path param policySetId
	if err := r.SetPathParam("policySetId", o.PolicySetID); err != nil {
		return err
	}
	if o.Rule != nil {
		if err := r.SetBodyParam(o.Rule); err != nil {
			return err
		}
	}

	// path param ruleId
	if err := r.SetPathParam("ruleId", o.RuleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
