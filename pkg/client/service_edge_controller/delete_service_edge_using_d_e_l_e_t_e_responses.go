// Code generated by go-swagger; DO NOT EDIT.

package service_edge_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteServiceEdgeUsingDELETEReader is a Reader for the DeleteServiceEdgeUsingDELETE structure.
type DeleteServiceEdgeUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServiceEdgeUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteServiceEdgeUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteServiceEdgeUsingDELETEBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteServiceEdgeUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteServiceEdgeUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteServiceEdgeUsingDELETETooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteServiceEdgeUsingDELETENoContent creates a DeleteServiceEdgeUsingDELETENoContent with default headers values
func NewDeleteServiceEdgeUsingDELETENoContent() *DeleteServiceEdgeUsingDELETENoContent {
	return &DeleteServiceEdgeUsingDELETENoContent{}
}

/* DeleteServiceEdgeUsingDELETENoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteServiceEdgeUsingDELETENoContent struct {
}

func (o *DeleteServiceEdgeUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/{serviceEdgeId}][%d] deleteServiceEdgeUsingDELETENoContent ", 204)
}

func (o *DeleteServiceEdgeUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServiceEdgeUsingDELETEBadRequest creates a DeleteServiceEdgeUsingDELETEBadRequest with default headers values
func NewDeleteServiceEdgeUsingDELETEBadRequest() *DeleteServiceEdgeUsingDELETEBadRequest {
	return &DeleteServiceEdgeUsingDELETEBadRequest{}
}

/* DeleteServiceEdgeUsingDELETEBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type DeleteServiceEdgeUsingDELETEBadRequest struct {
}

func (o *DeleteServiceEdgeUsingDELETEBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/{serviceEdgeId}][%d] deleteServiceEdgeUsingDELETEBadRequest ", 400)
}

func (o *DeleteServiceEdgeUsingDELETEBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServiceEdgeUsingDELETEUnauthorized creates a DeleteServiceEdgeUsingDELETEUnauthorized with default headers values
func NewDeleteServiceEdgeUsingDELETEUnauthorized() *DeleteServiceEdgeUsingDELETEUnauthorized {
	return &DeleteServiceEdgeUsingDELETEUnauthorized{}
}

/* DeleteServiceEdgeUsingDELETEUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteServiceEdgeUsingDELETEUnauthorized struct {
}

func (o *DeleteServiceEdgeUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/{serviceEdgeId}][%d] deleteServiceEdgeUsingDELETEUnauthorized ", 401)
}

func (o *DeleteServiceEdgeUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServiceEdgeUsingDELETEForbidden creates a DeleteServiceEdgeUsingDELETEForbidden with default headers values
func NewDeleteServiceEdgeUsingDELETEForbidden() *DeleteServiceEdgeUsingDELETEForbidden {
	return &DeleteServiceEdgeUsingDELETEForbidden{}
}

/* DeleteServiceEdgeUsingDELETEForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteServiceEdgeUsingDELETEForbidden struct {
}

func (o *DeleteServiceEdgeUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/{serviceEdgeId}][%d] deleteServiceEdgeUsingDELETEForbidden ", 403)
}

func (o *DeleteServiceEdgeUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServiceEdgeUsingDELETETooManyRequests creates a DeleteServiceEdgeUsingDELETETooManyRequests with default headers values
func NewDeleteServiceEdgeUsingDELETETooManyRequests() *DeleteServiceEdgeUsingDELETETooManyRequests {
	return &DeleteServiceEdgeUsingDELETETooManyRequests{}
}

/* DeleteServiceEdgeUsingDELETETooManyRequests describes a response with status code 429, with default header values.

TooManyRequest
*/
type DeleteServiceEdgeUsingDELETETooManyRequests struct {
}

func (o *DeleteServiceEdgeUsingDELETETooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/{serviceEdgeId}][%d] deleteServiceEdgeUsingDELETETooManyRequests ", 429)
}

func (o *DeleteServiceEdgeUsingDELETETooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
