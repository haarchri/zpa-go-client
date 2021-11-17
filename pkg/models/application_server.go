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

// ApplicationServer application server
//
// swagger:model ApplicationServer
type ApplicationServer struct {

	// address
	Address string `json:"address,omitempty"`

	// app server group ids
	AppServerGroupIds []string `json:"appServerGroupIds"`

	// config space
	// Enum: [DEFAULT SIEM]
	ConfigSpace string `json:"configSpace,omitempty"`

	// creation time
	CreationTime string `json:"creationTime,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// modified by
	ModifiedBy int64 `json:"modifiedBy,omitempty"`

	// modified time
	ModifiedTime string `json:"modifiedTime,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this application server
func (m *ApplicationServer) Validate(formats strfmt.Registry) error {
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

var applicationServerTypeConfigSpacePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","SIEM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationServerTypeConfigSpacePropEnum = append(applicationServerTypeConfigSpacePropEnum, v)
	}
}

const (

	// ApplicationServerConfigSpaceDEFAULT captures enum value "DEFAULT"
	ApplicationServerConfigSpaceDEFAULT string = "DEFAULT"

	// ApplicationServerConfigSpaceSIEM captures enum value "SIEM"
	ApplicationServerConfigSpaceSIEM string = "SIEM"
)

// prop value enum
func (m *ApplicationServer) validateConfigSpaceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, applicationServerTypeConfigSpacePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ApplicationServer) validateConfigSpace(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigSpace) { // not required
		return nil
	}

	// value enum
	if err := m.validateConfigSpaceEnum("configSpace", "body", m.ConfigSpace); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationServer) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this application server based on context it is used
func (m *ApplicationServer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationServer) UnmarshalBinary(b []byte) error {
	var res ApplicationServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
