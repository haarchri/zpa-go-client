// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PageListOfSCIMAttributeHeader page list of s c i m attribute header
//
// swagger:model PageListOfSCIMAttributeHeader
type PageListOfSCIMAttributeHeader struct {

	// list
	List []*SCIMAttributeHeader `json:"list"`

	// total pages
	TotalPages string `json:"totalPages,omitempty"`
}

// Validate validates this page list of s c i m attribute header
func (m *PageListOfSCIMAttributeHeader) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PageListOfSCIMAttributeHeader) validateList(formats strfmt.Registry) error {
	if swag.IsZero(m.List) { // not required
		return nil
	}

	for i := 0; i < len(m.List); i++ {
		if swag.IsZero(m.List[i]) { // not required
			continue
		}

		if m.List[i] != nil {
			if err := m.List[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this page list of s c i m attribute header based on the context it is used
func (m *PageListOfSCIMAttributeHeader) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PageListOfSCIMAttributeHeader) contextValidateList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.List); i++ {

		if m.List[i] != nil {
			if err := m.List[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PageListOfSCIMAttributeHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PageListOfSCIMAttributeHeader) UnmarshalBinary(b []byte) error {
	var res PageListOfSCIMAttributeHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
