// Code generated by go-swagger; DO NOT EDIT.

package lss_config_controller_v_2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetLssLogFormatUsingGET3Reader is a Reader for the GetLssLogFormatUsingGET3 structure.
type GetLssLogFormatUsingGET3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLssLogFormatUsingGET3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLssLogFormatUsingGET3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLssLogFormatUsingGET3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLssLogFormatUsingGET3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLssLogFormatUsingGET3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLssLogFormatUsingGET3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLssLogFormatUsingGET3TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLssLogFormatUsingGET3OK creates a GetLssLogFormatUsingGET3OK with default headers values
func NewGetLssLogFormatUsingGET3OK() *GetLssLogFormatUsingGET3OK {
	return &GetLssLogFormatUsingGET3OK{}
}

/* GetLssLogFormatUsingGET3OK describes a response with status code 200, with default header values.

OK
*/
type GetLssLogFormatUsingGET3OK struct {
	Payload map[string]interface{}
}

func (o *GetLssLogFormatUsingGET3OK) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v2/admin/lssConfig/logType/formats][%d] getLssLogFormatUsingGET3OK  %+v", 200, o.Payload)
}
func (o *GetLssLogFormatUsingGET3OK) GetPayload() map[string]interface{} {
	return o.Payload
}

func (o *GetLssLogFormatUsingGET3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLssLogFormatUsingGET3BadRequest creates a GetLssLogFormatUsingGET3BadRequest with default headers values
func NewGetLssLogFormatUsingGET3BadRequest() *GetLssLogFormatUsingGET3BadRequest {
	return &GetLssLogFormatUsingGET3BadRequest{}
}

/* GetLssLogFormatUsingGET3BadRequest describes a response with status code 400, with default header values.

BadRequest
*/
type GetLssLogFormatUsingGET3BadRequest struct {
}

func (o *GetLssLogFormatUsingGET3BadRequest) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v2/admin/lssConfig/logType/formats][%d] getLssLogFormatUsingGET3BadRequest ", 400)
}

func (o *GetLssLogFormatUsingGET3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLssLogFormatUsingGET3Unauthorized creates a GetLssLogFormatUsingGET3Unauthorized with default headers values
func NewGetLssLogFormatUsingGET3Unauthorized() *GetLssLogFormatUsingGET3Unauthorized {
	return &GetLssLogFormatUsingGET3Unauthorized{}
}

/* GetLssLogFormatUsingGET3Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetLssLogFormatUsingGET3Unauthorized struct {
}

func (o *GetLssLogFormatUsingGET3Unauthorized) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v2/admin/lssConfig/logType/formats][%d] getLssLogFormatUsingGET3Unauthorized ", 401)
}

func (o *GetLssLogFormatUsingGET3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLssLogFormatUsingGET3Forbidden creates a GetLssLogFormatUsingGET3Forbidden with default headers values
func NewGetLssLogFormatUsingGET3Forbidden() *GetLssLogFormatUsingGET3Forbidden {
	return &GetLssLogFormatUsingGET3Forbidden{}
}

/* GetLssLogFormatUsingGET3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetLssLogFormatUsingGET3Forbidden struct {
}

func (o *GetLssLogFormatUsingGET3Forbidden) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v2/admin/lssConfig/logType/formats][%d] getLssLogFormatUsingGET3Forbidden ", 403)
}

func (o *GetLssLogFormatUsingGET3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLssLogFormatUsingGET3NotFound creates a GetLssLogFormatUsingGET3NotFound with default headers values
func NewGetLssLogFormatUsingGET3NotFound() *GetLssLogFormatUsingGET3NotFound {
	return &GetLssLogFormatUsingGET3NotFound{}
}

/* GetLssLogFormatUsingGET3NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetLssLogFormatUsingGET3NotFound struct {
}

func (o *GetLssLogFormatUsingGET3NotFound) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v2/admin/lssConfig/logType/formats][%d] getLssLogFormatUsingGET3NotFound ", 404)
}

func (o *GetLssLogFormatUsingGET3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLssLogFormatUsingGET3TooManyRequests creates a GetLssLogFormatUsingGET3TooManyRequests with default headers values
func NewGetLssLogFormatUsingGET3TooManyRequests() *GetLssLogFormatUsingGET3TooManyRequests {
	return &GetLssLogFormatUsingGET3TooManyRequests{}
}

/* GetLssLogFormatUsingGET3TooManyRequests describes a response with status code 429, with default header values.

TooManyRequest
*/
type GetLssLogFormatUsingGET3TooManyRequests struct {
}

func (o *GetLssLogFormatUsingGET3TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mgmtconfig/v2/admin/lssConfig/logType/formats][%d] getLssLogFormatUsingGET3TooManyRequests ", 429)
}

func (o *GetLssLogFormatUsingGET3TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
