// Code generated by go-swagger; DO NOT EDIT.

package policy_set_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haarchri/zpa-go-client/pkg/models"
)

// GetRuleInPolicySetUsingGET1Reader is a Reader for the GetRuleInPolicySetUsingGET1 structure.
type GetRuleInPolicySetUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuleInPolicySetUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRuleInPolicySetUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRuleInPolicySetUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRuleInPolicySetUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRuleInPolicySetUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRuleInPolicySetUsingGET1OK creates a GetRuleInPolicySetUsingGET1OK with default headers values
func NewGetRuleInPolicySetUsingGET1OK() *GetRuleInPolicySetUsingGET1OK {
	return &GetRuleInPolicySetUsingGET1OK{}
}

/* GetRuleInPolicySetUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetRuleInPolicySetUsingGET1OK struct {
	Payload *models.PolicyRule
}

func (o *GetRuleInPolicySetUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/{policySetId}/rule/{ruleId}][%d] getRuleInPolicySetUsingGET1OK  %+v", 200, o.Payload)
}
func (o *GetRuleInPolicySetUsingGET1OK) GetPayload() *models.PolicyRule {
	return o.Payload
}

func (o *GetRuleInPolicySetUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuleInPolicySetUsingGET1Unauthorized creates a GetRuleInPolicySetUsingGET1Unauthorized with default headers values
func NewGetRuleInPolicySetUsingGET1Unauthorized() *GetRuleInPolicySetUsingGET1Unauthorized {
	return &GetRuleInPolicySetUsingGET1Unauthorized{}
}

/* GetRuleInPolicySetUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRuleInPolicySetUsingGET1Unauthorized struct {
}

func (o *GetRuleInPolicySetUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/{policySetId}/rule/{ruleId}][%d] getRuleInPolicySetUsingGET1Unauthorized ", 401)
}

func (o *GetRuleInPolicySetUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRuleInPolicySetUsingGET1Forbidden creates a GetRuleInPolicySetUsingGET1Forbidden with default headers values
func NewGetRuleInPolicySetUsingGET1Forbidden() *GetRuleInPolicySetUsingGET1Forbidden {
	return &GetRuleInPolicySetUsingGET1Forbidden{}
}

/* GetRuleInPolicySetUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRuleInPolicySetUsingGET1Forbidden struct {
}

func (o *GetRuleInPolicySetUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/{policySetId}/rule/{ruleId}][%d] getRuleInPolicySetUsingGET1Forbidden ", 403)
}

func (o *GetRuleInPolicySetUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRuleInPolicySetUsingGET1NotFound creates a GetRuleInPolicySetUsingGET1NotFound with default headers values
func NewGetRuleInPolicySetUsingGET1NotFound() *GetRuleInPolicySetUsingGET1NotFound {
	return &GetRuleInPolicySetUsingGET1NotFound{}
}

/* GetRuleInPolicySetUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetRuleInPolicySetUsingGET1NotFound struct {
}

func (o *GetRuleInPolicySetUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/{policySetId}/rule/{ruleId}][%d] getRuleInPolicySetUsingGET1NotFound ", 404)
}

func (o *GetRuleInPolicySetUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
