// Code generated by go-swagger; DO NOT EDIT.

package application_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateApplicationV2UsingPUT1Reader is a Reader for the UpdateApplicationV2UsingPUT1 structure.
type UpdateApplicationV2UsingPUT1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateApplicationV2UsingPUT1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateApplicationV2UsingPUT1Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateApplicationV2UsingPUT1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateApplicationV2UsingPUT1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateApplicationV2UsingPUT1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateApplicationV2UsingPUT1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateApplicationV2UsingPUT1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateApplicationV2UsingPUT1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateApplicationV2UsingPUT1Created creates a UpdateApplicationV2UsingPUT1Created with default headers values
func NewUpdateApplicationV2UsingPUT1Created() *UpdateApplicationV2UsingPUT1Created {
	return &UpdateApplicationV2UsingPUT1Created{}
}

/* UpdateApplicationV2UsingPUT1Created describes a response with status code 201, with default header values.

Created
*/
type UpdateApplicationV2UsingPUT1Created struct {
}

func (o *UpdateApplicationV2UsingPUT1Created) Error() string {
	return fmt.Sprintf("[PUT /mgmtconfig/v1/admin/customers/{customerId}/application/{applicationId}][%d] updateApplicationV2UsingPUT1Created ", 201)
}

func (o *UpdateApplicationV2UsingPUT1Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationV2UsingPUT1NoContent creates a UpdateApplicationV2UsingPUT1NoContent with default headers values
func NewUpdateApplicationV2UsingPUT1NoContent() *UpdateApplicationV2UsingPUT1NoContent {
	return &UpdateApplicationV2UsingPUT1NoContent{}
}

/* UpdateApplicationV2UsingPUT1NoContent describes a response with status code 204, with default header values.

No Content
*/
type UpdateApplicationV2UsingPUT1NoContent struct {
}

func (o *UpdateApplicationV2UsingPUT1NoContent) Error() string {
	return fmt.Sprintf("[PUT /mgmtconfig/v1/admin/customers/{customerId}/application/{applicationId}][%d] updateApplicationV2UsingPUT1NoContent ", 204)
}

func (o *UpdateApplicationV2UsingPUT1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationV2UsingPUT1BadRequest creates a UpdateApplicationV2UsingPUT1BadRequest with default headers values
func NewUpdateApplicationV2UsingPUT1BadRequest() *UpdateApplicationV2UsingPUT1BadRequest {
	return &UpdateApplicationV2UsingPUT1BadRequest{}
}

/* UpdateApplicationV2UsingPUT1BadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type UpdateApplicationV2UsingPUT1BadRequest struct {
}

func (o *UpdateApplicationV2UsingPUT1BadRequest) Error() string {
	return fmt.Sprintf("[PUT /mgmtconfig/v1/admin/customers/{customerId}/application/{applicationId}][%d] updateApplicationV2UsingPUT1BadRequest ", 400)
}

func (o *UpdateApplicationV2UsingPUT1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationV2UsingPUT1Unauthorized creates a UpdateApplicationV2UsingPUT1Unauthorized with default headers values
func NewUpdateApplicationV2UsingPUT1Unauthorized() *UpdateApplicationV2UsingPUT1Unauthorized {
	return &UpdateApplicationV2UsingPUT1Unauthorized{}
}

/* UpdateApplicationV2UsingPUT1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateApplicationV2UsingPUT1Unauthorized struct {
}

func (o *UpdateApplicationV2UsingPUT1Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /mgmtconfig/v1/admin/customers/{customerId}/application/{applicationId}][%d] updateApplicationV2UsingPUT1Unauthorized ", 401)
}

func (o *UpdateApplicationV2UsingPUT1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationV2UsingPUT1Forbidden creates a UpdateApplicationV2UsingPUT1Forbidden with default headers values
func NewUpdateApplicationV2UsingPUT1Forbidden() *UpdateApplicationV2UsingPUT1Forbidden {
	return &UpdateApplicationV2UsingPUT1Forbidden{}
}

/* UpdateApplicationV2UsingPUT1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateApplicationV2UsingPUT1Forbidden struct {
}

func (o *UpdateApplicationV2UsingPUT1Forbidden) Error() string {
	return fmt.Sprintf("[PUT /mgmtconfig/v1/admin/customers/{customerId}/application/{applicationId}][%d] updateApplicationV2UsingPUT1Forbidden ", 403)
}

func (o *UpdateApplicationV2UsingPUT1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationV2UsingPUT1NotFound creates a UpdateApplicationV2UsingPUT1NotFound with default headers values
func NewUpdateApplicationV2UsingPUT1NotFound() *UpdateApplicationV2UsingPUT1NotFound {
	return &UpdateApplicationV2UsingPUT1NotFound{}
}

/* UpdateApplicationV2UsingPUT1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateApplicationV2UsingPUT1NotFound struct {
}

func (o *UpdateApplicationV2UsingPUT1NotFound) Error() string {
	return fmt.Sprintf("[PUT /mgmtconfig/v1/admin/customers/{customerId}/application/{applicationId}][%d] updateApplicationV2UsingPUT1NotFound ", 404)
}

func (o *UpdateApplicationV2UsingPUT1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationV2UsingPUT1TooManyRequests creates a UpdateApplicationV2UsingPUT1TooManyRequests with default headers values
func NewUpdateApplicationV2UsingPUT1TooManyRequests() *UpdateApplicationV2UsingPUT1TooManyRequests {
	return &UpdateApplicationV2UsingPUT1TooManyRequests{}
}

/* UpdateApplicationV2UsingPUT1TooManyRequests describes a response with status code 429, with default header values.

TooManyRequest
*/
type UpdateApplicationV2UsingPUT1TooManyRequests struct {
}

func (o *UpdateApplicationV2UsingPUT1TooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /mgmtconfig/v1/admin/customers/{customerId}/application/{applicationId}][%d] updateApplicationV2UsingPUT1TooManyRequests ", 429)
}

func (o *UpdateApplicationV2UsingPUT1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
