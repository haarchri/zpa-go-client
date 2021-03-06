// Code generated by go-swagger; DO NOT EDIT.

package connector_group_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateAppConnectorGroupUsingPUT1Reader is a Reader for the UpdateAppConnectorGroupUsingPUT1 structure.
type UpdateAppConnectorGroupUsingPUT1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAppConnectorGroupUsingPUT1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateAppConnectorGroupUsingPUT1Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateAppConnectorGroupUsingPUT1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAppConnectorGroupUsingPUT1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateAppConnectorGroupUsingPUT1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateAppConnectorGroupUsingPUT1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAppConnectorGroupUsingPUT1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateAppConnectorGroupUsingPUT1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateAppConnectorGroupUsingPUT1Created creates a UpdateAppConnectorGroupUsingPUT1Created with default headers values
func NewUpdateAppConnectorGroupUsingPUT1Created() *UpdateAppConnectorGroupUsingPUT1Created {
	return &UpdateAppConnectorGroupUsingPUT1Created{}
}

/* UpdateAppConnectorGroupUsingPUT1Created describes a response with status code 201, with default header values.

Created
*/
type UpdateAppConnectorGroupUsingPUT1Created struct {
}

func (o *UpdateAppConnectorGroupUsingPUT1Created) Error() string {
	return fmt.Sprintf("[PUT /mgmtconfig/v1/admin/customers/{customerId}/appConnectorGroup/{appConnectorGroupId}][%d] updateAppConnectorGroupUsingPUT1Created ", 201)
}

func (o *UpdateAppConnectorGroupUsingPUT1Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAppConnectorGroupUsingPUT1NoContent creates a UpdateAppConnectorGroupUsingPUT1NoContent with default headers values
func NewUpdateAppConnectorGroupUsingPUT1NoContent() *UpdateAppConnectorGroupUsingPUT1NoContent {
	return &UpdateAppConnectorGroupUsingPUT1NoContent{}
}

/* UpdateAppConnectorGroupUsingPUT1NoContent describes a response with status code 204, with default header values.

No Content
*/
type UpdateAppConnectorGroupUsingPUT1NoContent struct {
}

func (o *UpdateAppConnectorGroupUsingPUT1NoContent) Error() string {
	return fmt.Sprintf("[PUT /mgmtconfig/v1/admin/customers/{customerId}/appConnectorGroup/{appConnectorGroupId}][%d] updateAppConnectorGroupUsingPUT1NoContent ", 204)
}

func (o *UpdateAppConnectorGroupUsingPUT1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAppConnectorGroupUsingPUT1BadRequest creates a UpdateAppConnectorGroupUsingPUT1BadRequest with default headers values
func NewUpdateAppConnectorGroupUsingPUT1BadRequest() *UpdateAppConnectorGroupUsingPUT1BadRequest {
	return &UpdateAppConnectorGroupUsingPUT1BadRequest{}
}

/* UpdateAppConnectorGroupUsingPUT1BadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type UpdateAppConnectorGroupUsingPUT1BadRequest struct {
}

func (o *UpdateAppConnectorGroupUsingPUT1BadRequest) Error() string {
	return fmt.Sprintf("[PUT /mgmtconfig/v1/admin/customers/{customerId}/appConnectorGroup/{appConnectorGroupId}][%d] updateAppConnectorGroupUsingPUT1BadRequest ", 400)
}

func (o *UpdateAppConnectorGroupUsingPUT1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAppConnectorGroupUsingPUT1Unauthorized creates a UpdateAppConnectorGroupUsingPUT1Unauthorized with default headers values
func NewUpdateAppConnectorGroupUsingPUT1Unauthorized() *UpdateAppConnectorGroupUsingPUT1Unauthorized {
	return &UpdateAppConnectorGroupUsingPUT1Unauthorized{}
}

/* UpdateAppConnectorGroupUsingPUT1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateAppConnectorGroupUsingPUT1Unauthorized struct {
}

func (o *UpdateAppConnectorGroupUsingPUT1Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /mgmtconfig/v1/admin/customers/{customerId}/appConnectorGroup/{appConnectorGroupId}][%d] updateAppConnectorGroupUsingPUT1Unauthorized ", 401)
}

func (o *UpdateAppConnectorGroupUsingPUT1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAppConnectorGroupUsingPUT1Forbidden creates a UpdateAppConnectorGroupUsingPUT1Forbidden with default headers values
func NewUpdateAppConnectorGroupUsingPUT1Forbidden() *UpdateAppConnectorGroupUsingPUT1Forbidden {
	return &UpdateAppConnectorGroupUsingPUT1Forbidden{}
}

/* UpdateAppConnectorGroupUsingPUT1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateAppConnectorGroupUsingPUT1Forbidden struct {
}

func (o *UpdateAppConnectorGroupUsingPUT1Forbidden) Error() string {
	return fmt.Sprintf("[PUT /mgmtconfig/v1/admin/customers/{customerId}/appConnectorGroup/{appConnectorGroupId}][%d] updateAppConnectorGroupUsingPUT1Forbidden ", 403)
}

func (o *UpdateAppConnectorGroupUsingPUT1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAppConnectorGroupUsingPUT1NotFound creates a UpdateAppConnectorGroupUsingPUT1NotFound with default headers values
func NewUpdateAppConnectorGroupUsingPUT1NotFound() *UpdateAppConnectorGroupUsingPUT1NotFound {
	return &UpdateAppConnectorGroupUsingPUT1NotFound{}
}

/* UpdateAppConnectorGroupUsingPUT1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateAppConnectorGroupUsingPUT1NotFound struct {
}

func (o *UpdateAppConnectorGroupUsingPUT1NotFound) Error() string {
	return fmt.Sprintf("[PUT /mgmtconfig/v1/admin/customers/{customerId}/appConnectorGroup/{appConnectorGroupId}][%d] updateAppConnectorGroupUsingPUT1NotFound ", 404)
}

func (o *UpdateAppConnectorGroupUsingPUT1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAppConnectorGroupUsingPUT1TooManyRequests creates a UpdateAppConnectorGroupUsingPUT1TooManyRequests with default headers values
func NewUpdateAppConnectorGroupUsingPUT1TooManyRequests() *UpdateAppConnectorGroupUsingPUT1TooManyRequests {
	return &UpdateAppConnectorGroupUsingPUT1TooManyRequests{}
}

/* UpdateAppConnectorGroupUsingPUT1TooManyRequests describes a response with status code 429, with default header values.

TooManyRequest
*/
type UpdateAppConnectorGroupUsingPUT1TooManyRequests struct {
}

func (o *UpdateAppConnectorGroupUsingPUT1TooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /mgmtconfig/v1/admin/customers/{customerId}/appConnectorGroup/{appConnectorGroupId}][%d] updateAppConnectorGroupUsingPUT1TooManyRequests ", 429)
}

func (o *UpdateAppConnectorGroupUsingPUT1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
