// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AppServerGroup app server group
//
// swagger:model AppServerGroup
type AppServerGroup struct {

	// config space
	// Enum: [DEFAULT SIEM]
	ConfigSpace string `json:"configSpace,omitempty"`

	// creation time
	CreationTime string `json:"creationTime,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// dynamic discovery
	DynamicDiscovery bool `json:"dynamicDiscovery,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// modified by
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// modified time
	ModifiedTime string `json:"modifiedTime,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this app server group
func (m *AppServerGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigSpace(formats); err != nil {
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

var appServerGroupTypeConfigSpacePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","SIEM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		appServerGroupTypeConfigSpacePropEnum = append(appServerGroupTypeConfigSpacePropEnum, v)
	}
}

const (

	// AppServerGroupConfigSpaceDEFAULT captures enum value "DEFAULT"
	AppServerGroupConfigSpaceDEFAULT string = "DEFAULT"

	// AppServerGroupConfigSpaceSIEM captures enum value "SIEM"
	AppServerGroupConfigSpaceSIEM string = "SIEM"
)

// prop value enum
func (m *AppServerGroup) validateConfigSpaceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, appServerGroupTypeConfigSpacePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AppServerGroup) validateConfigSpace(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigSpace) { // not required
		return nil
	}

	// value enum
	if err := m.validateConfigSpaceEnum("configSpace", "body", m.ConfigSpace); err != nil {
		return err
	}

	return nil
}

func (m *AppServerGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this app server group based on context it is used
func (m *AppServerGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppServerGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppServerGroup) UnmarshalBinary(b []byte) error {
	var res AppServerGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
