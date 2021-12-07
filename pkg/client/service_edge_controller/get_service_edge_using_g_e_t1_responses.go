// Code generated by go-swagger; DO NOT EDIT.

package service_edge_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haarchri/zpa-go-client/pkg/models"
)

// GetServiceEdgeUsingGET1Reader is a Reader for the GetServiceEdgeUsingGET1 structure.
type GetServiceEdgeUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceEdgeUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceEdgeUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetServiceEdgeUsingGET1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetServiceEdgeUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetServiceEdgeUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetServiceEdgeUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetServiceEdgeUsingGET1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetServiceEdgeUsingGET1OK creates a GetServiceEdgeUsingGET1OK with default headers values
func NewGetServiceEdgeUsingGET1OK() *GetServiceEdgeUsingGET1OK {
	return &GetServiceEdgeUsingGET1OK{}
}

/* GetServiceEdgeUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetServiceEdgeUsingGET1OK struct {
	Payload *models.ServiceEdge
}

func (o *GetServiceEdgeUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/{serviceEdgeId}][%d] getServiceEdgeUsingGET1OK  %+v", 200, o.Payload)
}
func (o *GetServiceEdgeUsingGET1OK) GetPayload() *models.ServiceEdge {
	return o.Payload
}

func (o *GetServiceEdgeUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceEdge)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceEdgeUsingGET1BadRequest creates a GetServiceEdgeUsingGET1BadRequest with default headers values
func NewGetServiceEdgeUsingGET1BadRequest() *GetServiceEdgeUsingGET1BadRequest {
	return &GetServiceEdgeUsingGET1BadRequest{}
}

/* GetServiceEdgeUsingGET1BadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type GetServiceEdgeUsingGET1BadRequest struct {
}

func (o *GetServiceEdgeUsingGET1BadRequest) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/{serviceEdgeId}][%d] getServiceEdgeUsingGET1BadRequest ", 400)
}

func (o *GetServiceEdgeUsingGET1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceEdgeUsingGET1Unauthorized creates a GetServiceEdgeUsingGET1Unauthorized with default headers values
func NewGetServiceEdgeUsingGET1Unauthorized() *GetServiceEdgeUsingGET1Unauthorized {
	return &GetServiceEdgeUsingGET1Unauthorized{}
}

/* GetServiceEdgeUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetServiceEdgeUsingGET1Unauthorized struct {
}

func (o *GetServiceEdgeUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/{serviceEdgeId}][%d] getServiceEdgeUsingGET1Unauthorized ", 401)
}

func (o *GetServiceEdgeUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceEdgeUsingGET1Forbidden creates a GetServiceEdgeUsingGET1Forbidden with default headers values
func NewGetServiceEdgeUsingGET1Forbidden() *GetServiceEdgeUsingGET1Forbidden {
	return &GetServiceEdgeUsingGET1Forbidden{}
}

/* GetServiceEdgeUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetServiceEdgeUsingGET1Forbidden struct {
}

func (o *GetServiceEdgeUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/{serviceEdgeId}][%d] getServiceEdgeUsingGET1Forbidden ", 403)
}

func (o *GetServiceEdgeUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceEdgeUsingGET1NotFound creates a GetServiceEdgeUsingGET1NotFound with default headers values
func NewGetServiceEdgeUsingGET1NotFound() *GetServiceEdgeUsingGET1NotFound {
	return &GetServiceEdgeUsingGET1NotFound{}
}

/* GetServiceEdgeUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetServiceEdgeUsingGET1NotFound struct {
}

func (o *GetServiceEdgeUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/{serviceEdgeId}][%d] getServiceEdgeUsingGET1NotFound ", 404)
}

func (o *GetServiceEdgeUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetServiceEdgeUsingGET1TooManyRequests creates a GetServiceEdgeUsingGET1TooManyRequests with default headers values
func NewGetServiceEdgeUsingGET1TooManyRequests() *GetServiceEdgeUsingGET1TooManyRequests {
	return &GetServiceEdgeUsingGET1TooManyRequests{}
}

/* GetServiceEdgeUsingGET1TooManyRequests describes a response with status code 429, with default header values.

TooManyRequest
*/
type GetServiceEdgeUsingGET1TooManyRequests struct {
}

func (o *GetServiceEdgeUsingGET1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/{serviceEdgeId}][%d] getServiceEdgeUsingGET1TooManyRequests ", 429)
}

func (o *GetServiceEdgeUsingGET1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
