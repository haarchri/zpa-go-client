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

// CommonApplicationDto common application dto
//
// swagger:model CommonApplicationDto
type CommonApplicationDto struct {

	// apps config
	AppsConfig []*CommonAppConfigDto `json:"appsConfig"`

	// deleted ba apps
	DeletedBaApps []int64 `json:"deletedBaApps"`

	// deleted inspect apps
	DeletedInspectApps []int64 `json:"deletedInspectApps"`
}

// Validate validates this common application dto
func (m *CommonApplicationDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppsConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonApplicationDto) validateAppsConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AppsConfig) { // not required
		return nil
	}

	for i := 0; i < len(m.AppsConfig); i++ {
		if swag.IsZero(m.AppsConfig[i]) { // not required
			continue
		}

		if m.AppsConfig[i] != nil {
			if err := m.AppsConfig[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appsConfig" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("appsConfig" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this common application dto based on the context it is used
func (m *CommonApplicationDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppsConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonApplicationDto) contextValidateAppsConfig(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AppsConfig); i++ {

		if m.AppsConfig[i] != nil {
			if err := m.AppsConfig[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appsConfig" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("appsConfig" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommonApplicationDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonApplicationDto) UnmarshalBinary(b []byte) error {
	var res CommonApplicationDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
