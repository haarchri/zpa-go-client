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

// GetSegmentGroupUsingGET1Reader is a Reader for the GetSegmentGroupUsingGET1 structure.
type GetSegmentGroupUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSegmentGroupUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSegmentGroupUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSegmentGroupUsingGET1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSegmentGroupUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSegmentGroupUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSegmentGroupUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSegmentGroupUsingGET1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSegmentGroupUsingGET1OK creates a GetSegmentGroupUsingGET1OK with default headers values
func NewGetSegmentGroupUsingGET1OK() *GetSegmentGroupUsingGET1OK {
	return &GetSegmentGroupUsingGET1OK{}
}

/* GetSegmentGroupUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetSegmentGroupUsingGET1OK struct {
	Payload *models.SegmentGroup
}

func (o *GetSegmentGroupUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup/{segmentGroupId}][%d] getSegmentGroupUsingGET1OK  %+v", 200, o.Payload)
}
func (o *GetSegmentGroupUsingGET1OK) GetPayload() *models.SegmentGroup {
	return o.Payload
}

func (o *GetSegmentGroupUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SegmentGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSegmentGroupUsingGET1BadRequest creates a GetSegmentGroupUsingGET1BadRequest with default headers values
func NewGetSegmentGroupUsingGET1BadRequest() *GetSegmentGroupUsingGET1BadRequest {
	return &GetSegmentGroupUsingGET1BadRequest{}
}

/* GetSegmentGroupUsingGET1BadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type GetSegmentGroupUsingGET1BadRequest struct {
}

func (o *GetSegmentGroupUsingGET1BadRequest) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup/{segmentGroupId}][%d] getSegmentGroupUsingGET1BadRequest ", 400)
}

func (o *GetSegmentGroupUsingGET1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSegmentGroupUsingGET1Unauthorized creates a GetSegmentGroupUsingGET1Unauthorized with default headers values
func NewGetSegmentGroupUsingGET1Unauthorized() *GetSegmentGroupUsingGET1Unauthorized {
	return &GetSegmentGroupUsingGET1Unauthorized{}
}

/* GetSegmentGroupUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetSegmentGroupUsingGET1Unauthorized struct {
}

func (o *GetSegmentGroupUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup/{segmentGroupId}][%d] getSegmentGroupUsingGET1Unauthorized ", 401)
}

func (o *GetSegmentGroupUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSegmentGroupUsingGET1Forbidden creates a GetSegmentGroupUsingGET1Forbidden with default headers values
func NewGetSegmentGroupUsingGET1Forbidden() *GetSegmentGroupUsingGET1Forbidden {
	return &GetSegmentGroupUsingGET1Forbidden{}
}

/* GetSegmentGroupUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSegmentGroupUsingGET1Forbidden struct {
}

func (o *GetSegmentGroupUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup/{segmentGroupId}][%d] getSegmentGroupUsingGET1Forbidden ", 403)
}

func (o *GetSegmentGroupUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSegmentGroupUsingGET1NotFound creates a GetSegmentGroupUsingGET1NotFound with default headers values
func NewGetSegmentGroupUsingGET1NotFound() *GetSegmentGroupUsingGET1NotFound {
	return &GetSegmentGroupUsingGET1NotFound{}
}

/* GetSegmentGroupUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetSegmentGroupUsingGET1NotFound struct {
}

func (o *GetSegmentGroupUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup/{segmentGroupId}][%d] getSegmentGroupUsingGET1NotFound ", 404)
}

func (o *GetSegmentGroupUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSegmentGroupUsingGET1TooManyRequests creates a GetSegmentGroupUsingGET1TooManyRequests with default headers values
func NewGetSegmentGroupUsingGET1TooManyRequests() *GetSegmentGroupUsingGET1TooManyRequests {
	return &GetSegmentGroupUsingGET1TooManyRequests{}
}

/* GetSegmentGroupUsingGET1TooManyRequests describes a response with status code 429, with default header values.

TooManyRequest
*/
type GetSegmentGroupUsingGET1TooManyRequests struct {
}

func (o *GetSegmentGroupUsingGET1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup/{segmentGroupId}][%d] getSegmentGroupUsingGET1TooManyRequests ", 429)
}

func (o *GetSegmentGroupUsingGET1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
