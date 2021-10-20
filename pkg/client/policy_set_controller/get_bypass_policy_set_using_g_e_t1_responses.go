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

// GetBypassPolicySetUsingGET1Reader is a Reader for the GetBypassPolicySetUsingGET1 structure.
type GetBypassPolicySetUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBypassPolicySetUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBypassPolicySetUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetBypassPolicySetUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBypassPolicySetUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBypassPolicySetUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBypassPolicySetUsingGET1OK creates a GetBypassPolicySetUsingGET1OK with default headers values
func NewGetBypassPolicySetUsingGET1OK() *GetBypassPolicySetUsingGET1OK {
	return &GetBypassPolicySetUsingGET1OK{}
}

/* GetBypassPolicySetUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetBypassPolicySetUsingGET1OK struct {
	Payload *models.PolicySet
}

func (o *GetBypassPolicySetUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/bypass][%d] getBypassPolicySetUsingGET1OK  %+v", 200, o.Payload)
}
func (o *GetBypassPolicySetUsingGET1OK) GetPayload() *models.PolicySet {
	return o.Payload
}

func (o *GetBypassPolicySetUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicySet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBypassPolicySetUsingGET1Unauthorized creates a GetBypassPolicySetUsingGET1Unauthorized with default headers values
func NewGetBypassPolicySetUsingGET1Unauthorized() *GetBypassPolicySetUsingGET1Unauthorized {
	return &GetBypassPolicySetUsingGET1Unauthorized{}
}

/* GetBypassPolicySetUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetBypassPolicySetUsingGET1Unauthorized struct {
}

func (o *GetBypassPolicySetUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/bypass][%d] getBypassPolicySetUsingGET1Unauthorized ", 401)
}

func (o *GetBypassPolicySetUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBypassPolicySetUsingGET1Forbidden creates a GetBypassPolicySetUsingGET1Forbidden with default headers values
func NewGetBypassPolicySetUsingGET1Forbidden() *GetBypassPolicySetUsingGET1Forbidden {
	return &GetBypassPolicySetUsingGET1Forbidden{}
}

/* GetBypassPolicySetUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetBypassPolicySetUsingGET1Forbidden struct {
}

func (o *GetBypassPolicySetUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/bypass][%d] getBypassPolicySetUsingGET1Forbidden ", 403)
}

func (o *GetBypassPolicySetUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBypassPolicySetUsingGET1NotFound creates a GetBypassPolicySetUsingGET1NotFound with default headers values
func NewGetBypassPolicySetUsingGET1NotFound() *GetBypassPolicySetUsingGET1NotFound {
	return &GetBypassPolicySetUsingGET1NotFound{}
}

/* GetBypassPolicySetUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetBypassPolicySetUsingGET1NotFound struct {
}

func (o *GetBypassPolicySetUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/bypass][%d] getBypassPolicySetUsingGET1NotFound ", 404)
}

func (o *GetBypassPolicySetUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
