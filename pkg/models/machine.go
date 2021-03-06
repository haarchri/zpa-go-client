// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Machine machine
//
// swagger:model Machine
type Machine struct {

	// creation time
	CreationTime string `json:"creationTime,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// enrollment cert
	EnrollmentCert map[string]string `json:"enrollmentCert,omitempty"`

	// fingerprint
	Fingerprint string `json:"fingerprint,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// issued cert Id
	IssuedCertID string `json:"issuedCertId,omitempty"`

	// machine group Id
	MachineGroupID string `json:"machineGroupId,omitempty"`

	// machine group name
	MachineGroupName string `json:"machineGroupName,omitempty"`

	// machine token Id
	MachineTokenID string `json:"machineTokenId,omitempty"`

	// modified by
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// modified time
	ModifiedTime string `json:"modifiedTime,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this machine
func (m *Machine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Machine) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this machine based on context it is used
func (m *Machine) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Machine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Machine) UnmarshalBinary(b []byte) error {
	var res Machine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
