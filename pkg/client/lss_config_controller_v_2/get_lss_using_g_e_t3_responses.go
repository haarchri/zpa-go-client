// Code generated by go-swagger; DO NOT EDIT.

package lss_config_controller_v_2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/haarchri/zpa-go-client/pkg/models"
)

// GetLssUsingGET3Reader is a Reader for the GetLssUsingGET3 structure.
type GetLssUsingGET3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLssUsingGET3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLssUsingGET3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLssUsingGET3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLssUsingGET3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLssUsingGET3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLssUsingGET3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLssUsingGET3TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLssUsingGET3OK creates a GetLssUsingGET3OK with default headers values
func NewGetLssUsingGET3OK() *GetLssUsingGET3OK {
	return &GetLssUsingGET3OK{}
}

/* GetLssUsingGET3OK describes a response with status code 200, with default header values.

OK
*/
type GetLssUsingGET3OK struct {
	Payload *models.LssResource
}

func (o *GetLssUsingGET3OK) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v2/admin/customers/{customerId}/lssConfig/{lssId}][%d] getLssUsingGET3OK  %+v", 200, o.Payload)
}
func (o *GetLssUsingGET3OK) GetPayload() *models.LssResource {
	return o.Payload
}

func (o *GetLssUsingGET3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LssResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLssUsingGET3BadRequest creates a GetLssUsingGET3BadRequest with default headers values
func NewGetLssUsingGET3BadRequest() *GetLssUsingGET3BadRequest {
	return &GetLssUsingGET3BadRequest{}
}

/* GetLssUsingGET3BadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type GetLssUsingGET3BadRequest struct {
}

func (o *GetLssUsingGET3BadRequest) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v2/admin/customers/{customerId}/lssConfig/{lssId}][%d] getLssUsingGET3BadRequest ", 400)
}

func (o *GetLssUsingGET3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLssUsingGET3Unauthorized creates a GetLssUsingGET3Unauthorized with default headers values
func NewGetLssUsingGET3Unauthorized() *GetLssUsingGET3Unauthorized {
	return &GetLssUsingGET3Unauthorized{}
}

/* GetLssUsingGET3Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetLssUsingGET3Unauthorized struct {
}

func (o *GetLssUsingGET3Unauthorized) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v2/admin/customers/{customerId}/lssConfig/{lssId}][%d] getLssUsingGET3Unauthorized ", 401)
}

func (o *GetLssUsingGET3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLssUsingGET3Forbidden creates a GetLssUsingGET3Forbidden with default headers values
func NewGetLssUsingGET3Forbidden() *GetLssUsingGET3Forbidden {
	return &GetLssUsingGET3Forbidden{}
}

/* GetLssUsingGET3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetLssUsingGET3Forbidden struct {
}

func (o *GetLssUsingGET3Forbidden) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v2/admin/customers/{customerId}/lssConfig/{lssId}][%d] getLssUsingGET3Forbidden ", 403)
}

func (o *GetLssUsingGET3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLssUsingGET3NotFound creates a GetLssUsingGET3NotFound with default headers values
func NewGetLssUsingGET3NotFound() *GetLssUsingGET3NotFound {
	return &GetLssUsingGET3NotFound{}
}

/* GetLssUsingGET3NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetLssUsingGET3NotFound struct {
}

func (o *GetLssUsingGET3NotFound) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v2/admin/customers/{customerId}/lssConfig/{lssId}][%d] getLssUsingGET3NotFound ", 404)
}

func (o *GetLssUsingGET3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLssUsingGET3TooManyRequests creates a GetLssUsingGET3TooManyRequests with default headers values
func NewGetLssUsingGET3TooManyRequests() *GetLssUsingGET3TooManyRequests {
	return &GetLssUsingGET3TooManyRequests{}
}

/* GetLssUsingGET3TooManyRequests describes a response with status code 429, with default header values.

TooManyRequest
*/
type GetLssUsingGET3TooManyRequests struct {
}

func (o *GetLssUsingGET3TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v2/admin/customers/{customerId}/lssConfig/{lssId}][%d] getLssUsingGET3TooManyRequests ", 429)
}

func (o *GetLssUsingGET3TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}