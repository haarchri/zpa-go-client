// Code generated by go-swagger; DO NOT EDIT.

package cloud_connector_group_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haarchri/zpa-go-client/pkg/models"
)

// GetCloudConnectorGroupUsingGETReader is a Reader for the GetCloudConnectorGroupUsingGET structure.
type GetCloudConnectorGroupUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudConnectorGroupUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCloudConnectorGroupUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCloudConnectorGroupUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCloudConnectorGroupUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCloudConnectorGroupUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCloudConnectorGroupUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCloudConnectorGroupUsingGETTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCloudConnectorGroupUsingGETOK creates a GetCloudConnectorGroupUsingGETOK with default headers values
func NewGetCloudConnectorGroupUsingGETOK() *GetCloudConnectorGroupUsingGETOK {
	return &GetCloudConnectorGroupUsingGETOK{}
}

/* GetCloudConnectorGroupUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetCloudConnectorGroupUsingGETOK struct {
	Payload *models.CloudConnectorGroupResource
}

func (o *GetCloudConnectorGroupUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/cloudConnectorGroup/{id}][%d] getCloudConnectorGroupUsingGETOK  %+v", 200, o.Payload)
}
func (o *GetCloudConnectorGroupUsingGETOK) GetPayload() *models.CloudConnectorGroupResource {
	return o.Payload
}

func (o *GetCloudConnectorGroupUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudConnectorGroupResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudConnectorGroupUsingGETBadRequest creates a GetCloudConnectorGroupUsingGETBadRequest with default headers values
func NewGetCloudConnectorGroupUsingGETBadRequest() *GetCloudConnectorGroupUsingGETBadRequest {
	return &GetCloudConnectorGroupUsingGETBadRequest{}
}

/* GetCloudConnectorGroupUsingGETBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type GetCloudConnectorGroupUsingGETBadRequest struct {
}

func (o *GetCloudConnectorGroupUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/cloudConnectorGroup/{id}][%d] getCloudConnectorGroupUsingGETBadRequest ", 400)
}

func (o *GetCloudConnectorGroupUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCloudConnectorGroupUsingGETUnauthorized creates a GetCloudConnectorGroupUsingGETUnauthorized with default headers values
func NewGetCloudConnectorGroupUsingGETUnauthorized() *GetCloudConnectorGroupUsingGETUnauthorized {
	return &GetCloudConnectorGroupUsingGETUnauthorized{}
}

/* GetCloudConnectorGroupUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCloudConnectorGroupUsingGETUnauthorized struct {
}

func (o *GetCloudConnectorGroupUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/cloudConnectorGroup/{id}][%d] getCloudConnectorGroupUsingGETUnauthorized ", 401)
}

func (o *GetCloudConnectorGroupUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCloudConnectorGroupUsingGETForbidden creates a GetCloudConnectorGroupUsingGETForbidden with default headers values
func NewGetCloudConnectorGroupUsingGETForbidden() *GetCloudConnectorGroupUsingGETForbidden {
	return &GetCloudConnectorGroupUsingGETForbidden{}
}

/* GetCloudConnectorGroupUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCloudConnectorGroupUsingGETForbidden struct {
}

func (o *GetCloudConnectorGroupUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/cloudConnectorGroup/{id}][%d] getCloudConnectorGroupUsingGETForbidden ", 403)
}

func (o *GetCloudConnectorGroupUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCloudConnectorGroupUsingGETNotFound creates a GetCloudConnectorGroupUsingGETNotFound with default headers values
func NewGetCloudConnectorGroupUsingGETNotFound() *GetCloudConnectorGroupUsingGETNotFound {
	return &GetCloudConnectorGroupUsingGETNotFound{}
}

/* GetCloudConnectorGroupUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetCloudConnectorGroupUsingGETNotFound struct {
}

func (o *GetCloudConnectorGroupUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/cloudConnectorGroup/{id}][%d] getCloudConnectorGroupUsingGETNotFound ", 404)
}

func (o *GetCloudConnectorGroupUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCloudConnectorGroupUsingGETTooManyRequests creates a GetCloudConnectorGroupUsingGETTooManyRequests with default headers values
func NewGetCloudConnectorGroupUsingGETTooManyRequests() *GetCloudConnectorGroupUsingGETTooManyRequests {
	return &GetCloudConnectorGroupUsingGETTooManyRequests{}
}

/* GetCloudConnectorGroupUsingGETTooManyRequests describes a response with status code 429, with default header values.

TooManyRequest
*/
type GetCloudConnectorGroupUsingGETTooManyRequests struct {
}

func (o *GetCloudConnectorGroupUsingGETTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/cloudConnectorGroup/{id}][%d] getCloudConnectorGroupUsingGETTooManyRequests ", 429)
}

func (o *GetCloudConnectorGroupUsingGETTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
