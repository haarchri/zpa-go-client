// Code generated by go-swagger; DO NOT EDIT.

package trusted_network_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haarchri/zpa-go-client/pkg/models"
)

// GetTrustedNetworkUsingGET1Reader is a Reader for the GetTrustedNetworkUsingGET1 structure.
type GetTrustedNetworkUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTrustedNetworkUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTrustedNetworkUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTrustedNetworkUsingGET1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTrustedNetworkUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTrustedNetworkUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTrustedNetworkUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTrustedNetworkUsingGET1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTrustedNetworkUsingGET1OK creates a GetTrustedNetworkUsingGET1OK with default headers values
func NewGetTrustedNetworkUsingGET1OK() *GetTrustedNetworkUsingGET1OK {
	return &GetTrustedNetworkUsingGET1OK{}
}

/* GetTrustedNetworkUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetTrustedNetworkUsingGET1OK struct {
	Payload *models.TrustedNetwork
}

func (o *GetTrustedNetworkUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/network/{id}][%d] getTrustedNetworkUsingGET1OK  %+v", 200, o.Payload)
}
func (o *GetTrustedNetworkUsingGET1OK) GetPayload() *models.TrustedNetwork {
	return o.Payload
}

func (o *GetTrustedNetworkUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TrustedNetwork)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTrustedNetworkUsingGET1BadRequest creates a GetTrustedNetworkUsingGET1BadRequest with default headers values
func NewGetTrustedNetworkUsingGET1BadRequest() *GetTrustedNetworkUsingGET1BadRequest {
	return &GetTrustedNetworkUsingGET1BadRequest{}
}

/* GetTrustedNetworkUsingGET1BadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type GetTrustedNetworkUsingGET1BadRequest struct {
}

func (o *GetTrustedNetworkUsingGET1BadRequest) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/network/{id}][%d] getTrustedNetworkUsingGET1BadRequest ", 400)
}

func (o *GetTrustedNetworkUsingGET1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTrustedNetworkUsingGET1Unauthorized creates a GetTrustedNetworkUsingGET1Unauthorized with default headers values
func NewGetTrustedNetworkUsingGET1Unauthorized() *GetTrustedNetworkUsingGET1Unauthorized {
	return &GetTrustedNetworkUsingGET1Unauthorized{}
}

/* GetTrustedNetworkUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetTrustedNetworkUsingGET1Unauthorized struct {
}

func (o *GetTrustedNetworkUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/network/{id}][%d] getTrustedNetworkUsingGET1Unauthorized ", 401)
}

func (o *GetTrustedNetworkUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTrustedNetworkUsingGET1Forbidden creates a GetTrustedNetworkUsingGET1Forbidden with default headers values
func NewGetTrustedNetworkUsingGET1Forbidden() *GetTrustedNetworkUsingGET1Forbidden {
	return &GetTrustedNetworkUsingGET1Forbidden{}
}

/* GetTrustedNetworkUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetTrustedNetworkUsingGET1Forbidden struct {
}

func (o *GetTrustedNetworkUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/network/{id}][%d] getTrustedNetworkUsingGET1Forbidden ", 403)
}

func (o *GetTrustedNetworkUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTrustedNetworkUsingGET1NotFound creates a GetTrustedNetworkUsingGET1NotFound with default headers values
func NewGetTrustedNetworkUsingGET1NotFound() *GetTrustedNetworkUsingGET1NotFound {
	return &GetTrustedNetworkUsingGET1NotFound{}
}

/* GetTrustedNetworkUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetTrustedNetworkUsingGET1NotFound struct {
}

func (o *GetTrustedNetworkUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/network/{id}][%d] getTrustedNetworkUsingGET1NotFound ", 404)
}

func (o *GetTrustedNetworkUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTrustedNetworkUsingGET1TooManyRequests creates a GetTrustedNetworkUsingGET1TooManyRequests with default headers values
func NewGetTrustedNetworkUsingGET1TooManyRequests() *GetTrustedNetworkUsingGET1TooManyRequests {
	return &GetTrustedNetworkUsingGET1TooManyRequests{}
}

/* GetTrustedNetworkUsingGET1TooManyRequests describes a response with status code 429, with default header values.

TooManyRequest
*/
type GetTrustedNetworkUsingGET1TooManyRequests struct {
}

func (o *GetTrustedNetworkUsingGET1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v1/admin/customers/{customerId}/network/{id}][%d] getTrustedNetworkUsingGET1TooManyRequests ", 429)
}

func (o *GetTrustedNetworkUsingGET1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
