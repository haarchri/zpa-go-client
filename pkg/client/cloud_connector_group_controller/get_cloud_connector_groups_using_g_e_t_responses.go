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

// GetCloudConnectorGroupsUsingGETReader is a Reader for the GetCloudConnectorGroupsUsingGET structure.
type GetCloudConnectorGroupsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudConnectorGroupsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCloudConnectorGroupsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCloudConnectorGroupsUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCloudConnectorGroupsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCloudConnectorGroupsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCloudConnectorGroupsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCloudConnectorGroupsUsingGETTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCloudConnectorGroupsUsingGETOK creates a GetCloudConnectorGroupsUsingGETOK with default headers values
func NewGetCloudConnectorGroupsUsingGETOK() *GetCloudConnectorGroupsUsingGETOK {
	return &GetCloudConnectorGroupsUsingGETOK{}
}

/* GetCloudConnectorGroupsUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetCloudConnectorGroupsUsingGETOK struct {
	Payload *models.PageListOfCloudConnectorGroupResource
}

func (o *GetCloudConnectorGroupsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/cloudConnectorGroup][%d] getCloudConnectorGroupsUsingGETOK  %+v", 200, o.Payload)
}
func (o *GetCloudConnectorGroupsUsingGETOK) GetPayload() *models.PageListOfCloudConnectorGroupResource {
	return o.Payload
}

func (o *GetCloudConnectorGroupsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageListOfCloudConnectorGroupResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudConnectorGroupsUsingGETBadRequest creates a GetCloudConnectorGroupsUsingGETBadRequest with default headers values
func NewGetCloudConnectorGroupsUsingGETBadRequest() *GetCloudConnectorGroupsUsingGETBadRequest {
	return &GetCloudConnectorGroupsUsingGETBadRequest{}
}

/* GetCloudConnectorGroupsUsingGETBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type GetCloudConnectorGroupsUsingGETBadRequest struct {
}

func (o *GetCloudConnectorGroupsUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/cloudConnectorGroup][%d] getCloudConnectorGroupsUsingGETBadRequest ", 400)
}

func (o *GetCloudConnectorGroupsUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCloudConnectorGroupsUsingGETUnauthorized creates a GetCloudConnectorGroupsUsingGETUnauthorized with default headers values
func NewGetCloudConnectorGroupsUsingGETUnauthorized() *GetCloudConnectorGroupsUsingGETUnauthorized {
	return &GetCloudConnectorGroupsUsingGETUnauthorized{}
}

/* GetCloudConnectorGroupsUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetCloudConnectorGroupsUsingGETUnauthorized struct {
}

func (o *GetCloudConnectorGroupsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/cloudConnectorGroup][%d] getCloudConnectorGroupsUsingGETUnauthorized ", 401)
}

func (o *GetCloudConnectorGroupsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCloudConnectorGroupsUsingGETForbidden creates a GetCloudConnectorGroupsUsingGETForbidden with default headers values
func NewGetCloudConnectorGroupsUsingGETForbidden() *GetCloudConnectorGroupsUsingGETForbidden {
	return &GetCloudConnectorGroupsUsingGETForbidden{}
}

/* GetCloudConnectorGroupsUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCloudConnectorGroupsUsingGETForbidden struct {
}

func (o *GetCloudConnectorGroupsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/cloudConnectorGroup][%d] getCloudConnectorGroupsUsingGETForbidden ", 403)
}

func (o *GetCloudConnectorGroupsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCloudConnectorGroupsUsingGETNotFound creates a GetCloudConnectorGroupsUsingGETNotFound with default headers values
func NewGetCloudConnectorGroupsUsingGETNotFound() *GetCloudConnectorGroupsUsingGETNotFound {
	return &GetCloudConnectorGroupsUsingGETNotFound{}
}

/* GetCloudConnectorGroupsUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetCloudConnectorGroupsUsingGETNotFound struct {
}

func (o *GetCloudConnectorGroupsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/cloudConnectorGroup][%d] getCloudConnectorGroupsUsingGETNotFound ", 404)
}

func (o *GetCloudConnectorGroupsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCloudConnectorGroupsUsingGETTooManyRequests creates a GetCloudConnectorGroupsUsingGETTooManyRequests with default headers values
func NewGetCloudConnectorGroupsUsingGETTooManyRequests() *GetCloudConnectorGroupsUsingGETTooManyRequests {
	return &GetCloudConnectorGroupsUsingGETTooManyRequests{}
}

/* GetCloudConnectorGroupsUsingGETTooManyRequests describes a response with status code 429, with default header values.

TooManyRequest
*/
type GetCloudConnectorGroupsUsingGETTooManyRequests struct {
}

func (o *GetCloudConnectorGroupsUsingGETTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/cloudConnectorGroup][%d] getCloudConnectorGroupsUsingGETTooManyRequests ", 429)
}

func (o *GetCloudConnectorGroupsUsingGETTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
