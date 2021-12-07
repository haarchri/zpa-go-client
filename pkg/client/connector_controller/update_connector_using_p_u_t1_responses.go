// Code generated by go-swagger; DO NOT EDIT.

package connector_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateConnectorUsingPUT1Reader is a Reader for the UpdateConnectorUsingPUT1 structure.
type UpdateConnectorUsingPUT1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateConnectorUsingPUT1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateConnectorUsingPUT1Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateConnectorUsingPUT1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateConnectorUsingPUT1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateConnectorUsingPUT1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateConnectorUsingPUT1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateConnectorUsingPUT1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateConnectorUsingPUT1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateConnectorUsingPUT1Created creates a UpdateConnectorUsingPUT1Created with default headers values
func NewUpdateConnectorUsingPUT1Created() *UpdateConnectorUsingPUT1Created {
	return &UpdateConnectorUsingPUT1Created{}
}

/* UpdateConnectorUsingPUT1Created describes a response with status code 201, with default header values.

Created
*/
type UpdateConnectorUsingPUT1Created struct {
}

func (o *UpdateConnectorUsingPUT1Created) Error() string {
	return fmt.Sprintf("[PUT /mgmtconfig/v1/admin/customers/{customerId}/connector/{connectorId}][%d] updateConnectorUsingPUT1Created ", 201)
}

func (o *UpdateConnectorUsingPUT1Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateConnectorUsingPUT1NoContent creates a UpdateConnectorUsingPUT1NoContent with default headers values
func NewUpdateConnectorUsingPUT1NoContent() *UpdateConnectorUsingPUT1NoContent {
	return &UpdateConnectorUsingPUT1NoContent{}
}

/* UpdateConnectorUsingPUT1NoContent describes a response with status code 204, with default header values.

No Content
*/
type UpdateConnectorUsingPUT1NoContent struct {
}

func (o *UpdateConnectorUsingPUT1NoContent) Error() string {
	return fmt.Sprintf("[PUT /mgmtconfig/v1/admin/customers/{customerId}/connector/{connectorId}][%d] updateConnectorUsingPUT1NoContent ", 204)
}

func (o *UpdateConnectorUsingPUT1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateConnectorUsingPUT1BadRequest creates a UpdateConnectorUsingPUT1BadRequest with default headers values
func NewUpdateConnectorUsingPUT1BadRequest() *UpdateConnectorUsingPUT1BadRequest {
	return &UpdateConnectorUsingPUT1BadRequest{}
}

/* UpdateConnectorUsingPUT1BadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type UpdateConnectorUsingPUT1BadRequest struct {
}

func (o *UpdateConnectorUsingPUT1BadRequest) Error() string {
	return fmt.Sprintf("[PUT /mgmtconfig/v1/admin/customers/{customerId}/connector/{connectorId}][%d] updateConnectorUsingPUT1BadRequest ", 400)
}

func (o *UpdateConnectorUsingPUT1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateConnectorUsingPUT1Unauthorized creates a UpdateConnectorUsingPUT1Unauthorized with default headers values
func NewUpdateConnectorUsingPUT1Unauthorized() *UpdateConnectorUsingPUT1Unauthorized {
	return &UpdateConnectorUsingPUT1Unauthorized{}
}

/* UpdateConnectorUsingPUT1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateConnectorUsingPUT1Unauthorized struct {
}

func (o *UpdateConnectorUsingPUT1Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /mgmtconfig/v1/admin/customers/{customerId}/connector/{connectorId}][%d] updateConnectorUsingPUT1Unauthorized ", 401)
}

func (o *UpdateConnectorUsingPUT1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateConnectorUsingPUT1Forbidden creates a UpdateConnectorUsingPUT1Forbidden with default headers values
func NewUpdateConnectorUsingPUT1Forbidden() *UpdateConnectorUsingPUT1Forbidden {
	return &UpdateConnectorUsingPUT1Forbidden{}
}

/* UpdateConnectorUsingPUT1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateConnectorUsingPUT1Forbidden struct {
}

func (o *UpdateConnectorUsingPUT1Forbidden) Error() string {
	return fmt.Sprintf("[PUT /mgmtconfig/v1/admin/customers/{customerId}/connector/{connectorId}][%d] updateConnectorUsingPUT1Forbidden ", 403)
}

func (o *UpdateConnectorUsingPUT1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateConnectorUsingPUT1NotFound creates a UpdateConnectorUsingPUT1NotFound with default headers values
func NewUpdateConnectorUsingPUT1NotFound() *UpdateConnectorUsingPUT1NotFound {
	return &UpdateConnectorUsingPUT1NotFound{}
}

/* UpdateConnectorUsingPUT1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateConnectorUsingPUT1NotFound struct {
}

func (o *UpdateConnectorUsingPUT1NotFound) Error() string {
	return fmt.Sprintf("[PUT /mgmtconfig/v1/admin/customers/{customerId}/connector/{connectorId}][%d] updateConnectorUsingPUT1NotFound ", 404)
}

func (o *UpdateConnectorUsingPUT1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateConnectorUsingPUT1TooManyRequests creates a UpdateConnectorUsingPUT1TooManyRequests with default headers values
func NewUpdateConnectorUsingPUT1TooManyRequests() *UpdateConnectorUsingPUT1TooManyRequests {
	return &UpdateConnectorUsingPUT1TooManyRequests{}
}

/* UpdateConnectorUsingPUT1TooManyRequests describes a response with status code 429, with default header values.

TooManyRequest
*/
type UpdateConnectorUsingPUT1TooManyRequests struct {
}

func (o *UpdateConnectorUsingPUT1TooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /mgmtconfig/v1/admin/customers/{customerId}/connector/{connectorId}][%d] updateConnectorUsingPUT1TooManyRequests ", 429)
}

func (o *UpdateConnectorUsingPUT1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}