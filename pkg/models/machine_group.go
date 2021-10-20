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
	"github.com/go-openapi/validate"
)

// MachineGroup machine group
//
// swagger:model MachineGroup
type MachineGroup struct {

	// creation time
	CreationTime int32 `json:"creationTime,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// machines
	Machines []*Machine `json:"machines"`

	// modified by
	ModifiedBy int64 `json:"modifiedBy,omitempty"`

	// modified time
	ModifiedTime int32 `json:"modifiedTime,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this machine group
func (m *MachineGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMachines(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MachineGroup) validateMachines(formats strfmt.Registry) error {
	if swag.IsZero(m.Machines) { // not required
		return nil
	}

	for i := 0; i < len(m.Machines); i++ {
		if swag.IsZero(m.Machines[i]) { // not required
			continue
		}

		if m.Machines[i] != nil {
			if err := m.Machines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("machines" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("machines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MachineGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this machine group based on the context it is used
func (m *MachineGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMachines(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MachineGroup) contextValidateMachines(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Machines); i++ {

		if m.Machines[i] != nil {
			if err := m.Machines[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("machines" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("machines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MachineGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MachineGroup) UnmarshalBinary(b []byte) error {
	var res MachineGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}