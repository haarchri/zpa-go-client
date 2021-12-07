// Code generated by go-swagger; DO NOT EDIT.

package service_edge_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// BulkDeleteServiceEdgeUsingPOSTReader is a Reader for the BulkDeleteServiceEdgeUsingPOST structure.
type BulkDeleteServiceEdgeUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkDeleteServiceEdgeUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkDeleteServiceEdgeUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewBulkDeleteServiceEdgeUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBulkDeleteServiceEdgeUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewBulkDeleteServiceEdgeUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBulkDeleteServiceEdgeUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewBulkDeleteServiceEdgeUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewBulkDeleteServiceEdgeUsingPOSTTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewBulkDeleteServiceEdgeUsingPOSTOK creates a BulkDeleteServiceEdgeUsingPOSTOK with default headers values
func NewBulkDeleteServiceEdgeUsingPOSTOK() *BulkDeleteServiceEdgeUsingPOSTOK {
	return &BulkDeleteServiceEdgeUsingPOSTOK{}
}

/* BulkDeleteServiceEdgeUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type BulkDeleteServiceEdgeUsingPOSTOK struct {
}

func (o *BulkDeleteServiceEdgeUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/bulkDelete][%d] bulkDeleteServiceEdgeUsingPOSTOK ", 200)
}

func (o *BulkDeleteServiceEdgeUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBulkDeleteServiceEdgeUsingPOSTCreated creates a BulkDeleteServiceEdgeUsingPOSTCreated with default headers values
func NewBulkDeleteServiceEdgeUsingPOSTCreated() *BulkDeleteServiceEdgeUsingPOSTCreated {
	return &BulkDeleteServiceEdgeUsingPOSTCreated{}
}

/* BulkDeleteServiceEdgeUsingPOSTCreated describes a response with status code 201, with default header values.

Created
*/
type BulkDeleteServiceEdgeUsingPOSTCreated struct {
}

func (o *BulkDeleteServiceEdgeUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/bulkDelete][%d] bulkDeleteServiceEdgeUsingPOSTCreated ", 201)
}

func (o *BulkDeleteServiceEdgeUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBulkDeleteServiceEdgeUsingPOSTBadRequest creates a BulkDeleteServiceEdgeUsingPOSTBadRequest with default headers values
func NewBulkDeleteServiceEdgeUsingPOSTBadRequest() *BulkDeleteServiceEdgeUsingPOSTBadRequest {
	return &BulkDeleteServiceEdgeUsingPOSTBadRequest{}
}

/* BulkDeleteServiceEdgeUsingPOSTBadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type BulkDeleteServiceEdgeUsingPOSTBadRequest struct {
}

func (o *BulkDeleteServiceEdgeUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/bulkDelete][%d] bulkDeleteServiceEdgeUsingPOSTBadRequest ", 400)
}

func (o *BulkDeleteServiceEdgeUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBulkDeleteServiceEdgeUsingPOSTUnauthorized creates a BulkDeleteServiceEdgeUsingPOSTUnauthorized with default headers values
func NewBulkDeleteServiceEdgeUsingPOSTUnauthorized() *BulkDeleteServiceEdgeUsingPOSTUnauthorized {
	return &BulkDeleteServiceEdgeUsingPOSTUnauthorized{}
}

/* BulkDeleteServiceEdgeUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type BulkDeleteServiceEdgeUsingPOSTUnauthorized struct {
}

func (o *BulkDeleteServiceEdgeUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/bulkDelete][%d] bulkDeleteServiceEdgeUsingPOSTUnauthorized ", 401)
}

func (o *BulkDeleteServiceEdgeUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBulkDeleteServiceEdgeUsingPOSTForbidden creates a BulkDeleteServiceEdgeUsingPOSTForbidden with default headers values
func NewBulkDeleteServiceEdgeUsingPOSTForbidden() *BulkDeleteServiceEdgeUsingPOSTForbidden {
	return &BulkDeleteServiceEdgeUsingPOSTForbidden{}
}

/* BulkDeleteServiceEdgeUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type BulkDeleteServiceEdgeUsingPOSTForbidden struct {
}

func (o *BulkDeleteServiceEdgeUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/bulkDelete][%d] bulkDeleteServiceEdgeUsingPOSTForbidden ", 403)
}

func (o *BulkDeleteServiceEdgeUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBulkDeleteServiceEdgeUsingPOSTNotFound creates a BulkDeleteServiceEdgeUsingPOSTNotFound with default headers values
func NewBulkDeleteServiceEdgeUsingPOSTNotFound() *BulkDeleteServiceEdgeUsingPOSTNotFound {
	return &BulkDeleteServiceEdgeUsingPOSTNotFound{}
}

/* BulkDeleteServiceEdgeUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type BulkDeleteServiceEdgeUsingPOSTNotFound struct {
}

func (o *BulkDeleteServiceEdgeUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/bulkDelete][%d] bulkDeleteServiceEdgeUsingPOSTNotFound ", 404)
}

func (o *BulkDeleteServiceEdgeUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBulkDeleteServiceEdgeUsingPOSTTooManyRequests creates a BulkDeleteServiceEdgeUsingPOSTTooManyRequests with default headers values
func NewBulkDeleteServiceEdgeUsingPOSTTooManyRequests() *BulkDeleteServiceEdgeUsingPOSTTooManyRequests {
	return &BulkDeleteServiceEdgeUsingPOSTTooManyRequests{}
}

/* BulkDeleteServiceEdgeUsingPOSTTooManyRequests describes a response with status code 429, with default header values.

TooManyRequest
*/
type BulkDeleteServiceEdgeUsingPOSTTooManyRequests struct {
}

func (o *BulkDeleteServiceEdgeUsingPOSTTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/bulkDelete][%d] bulkDeleteServiceEdgeUsingPOSTTooManyRequests ", 429)
}

func (o *BulkDeleteServiceEdgeUsingPOSTTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}