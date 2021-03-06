// Code generated by go-swagger; DO NOT EDIT.

package segment_group_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haarchri/zpa-go-client/pkg/models"
)

// AddSegmentGroupUsingPOST1Reader is a Reader for the AddSegmentGroupUsingPOST1 structure.
type AddSegmentGroupUsingPOST1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddSegmentGroupUsingPOST1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddSegmentGroupUsingPOST1Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddSegmentGroupUsingPOST1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddSegmentGroupUsingPOST1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddSegmentGroupUsingPOST1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddSegmentGroupUsingPOST1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAddSegmentGroupUsingPOST1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddSegmentGroupUsingPOST1Created creates a AddSegmentGroupUsingPOST1Created with default headers values
func NewAddSegmentGroupUsingPOST1Created() *AddSegmentGroupUsingPOST1Created {
	return &AddSegmentGroupUsingPOST1Created{}
}

/* AddSegmentGroupUsingPOST1Created describes a response with status code 201, with default header values.

Created
*/
type AddSegmentGroupUsingPOST1Created struct {
	Payload *models.SegmentGroup
}

func (o *AddSegmentGroupUsingPOST1Created) Error() string {
	return fmt.Sprintf("[POST /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup][%d] addSegmentGroupUsingPOST1Created  %+v", 201, o.Payload)
}
func (o *AddSegmentGroupUsingPOST1Created) GetPayload() *models.SegmentGroup {
	return o.Payload
}

func (o *AddSegmentGroupUsingPOST1Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SegmentGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSegmentGroupUsingPOST1BadRequest creates a AddSegmentGroupUsingPOST1BadRequest with default headers values
func NewAddSegmentGroupUsingPOST1BadRequest() *AddSegmentGroupUsingPOST1BadRequest {
	return &AddSegmentGroupUsingPOST1BadRequest{}
}

/* AddSegmentGroupUsingPOST1BadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type AddSegmentGroupUsingPOST1BadRequest struct {
}

func (o *AddSegmentGroupUsingPOST1BadRequest) Error() string {
	return fmt.Sprintf("[POST /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup][%d] addSegmentGroupUsingPOST1BadRequest ", 400)
}

func (o *AddSegmentGroupUsingPOST1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddSegmentGroupUsingPOST1Unauthorized creates a AddSegmentGroupUsingPOST1Unauthorized with default headers values
func NewAddSegmentGroupUsingPOST1Unauthorized() *AddSegmentGroupUsingPOST1Unauthorized {
	return &AddSegmentGroupUsingPOST1Unauthorized{}
}

/* AddSegmentGroupUsingPOST1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AddSegmentGroupUsingPOST1Unauthorized struct {
}

func (o *AddSegmentGroupUsingPOST1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup][%d] addSegmentGroupUsingPOST1Unauthorized ", 401)
}

func (o *AddSegmentGroupUsingPOST1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddSegmentGroupUsingPOST1Forbidden creates a AddSegmentGroupUsingPOST1Forbidden with default headers values
func NewAddSegmentGroupUsingPOST1Forbidden() *AddSegmentGroupUsingPOST1Forbidden {
	return &AddSegmentGroupUsingPOST1Forbidden{}
}

/* AddSegmentGroupUsingPOST1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AddSegmentGroupUsingPOST1Forbidden struct {
}

func (o *AddSegmentGroupUsingPOST1Forbidden) Error() string {
	return fmt.Sprintf("[POST /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup][%d] addSegmentGroupUsingPOST1Forbidden ", 403)
}

func (o *AddSegmentGroupUsingPOST1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddSegmentGroupUsingPOST1NotFound creates a AddSegmentGroupUsingPOST1NotFound with default headers values
func NewAddSegmentGroupUsingPOST1NotFound() *AddSegmentGroupUsingPOST1NotFound {
	return &AddSegmentGroupUsingPOST1NotFound{}
}

/* AddSegmentGroupUsingPOST1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type AddSegmentGroupUsingPOST1NotFound struct {
}

func (o *AddSegmentGroupUsingPOST1NotFound) Error() string {
	return fmt.Sprintf("[POST /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup][%d] addSegmentGroupUsingPOST1NotFound ", 404)
}

func (o *AddSegmentGroupUsingPOST1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddSegmentGroupUsingPOST1TooManyRequests creates a AddSegmentGroupUsingPOST1TooManyRequests with default headers values
func NewAddSegmentGroupUsingPOST1TooManyRequests() *AddSegmentGroupUsingPOST1TooManyRequests {
	return &AddSegmentGroupUsingPOST1TooManyRequests{}
}

/* AddSegmentGroupUsingPOST1TooManyRequests describes a response with status code 429, with default header values.

TooManyRequest
*/
type AddSegmentGroupUsingPOST1TooManyRequests struct {
}

func (o *AddSegmentGroupUsingPOST1TooManyRequests) Error() string {
	return fmt.Sprintf("[POST /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup][%d] addSegmentGroupUsingPOST1TooManyRequests ", 429)
}

func (o *AddSegmentGroupUsingPOST1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
