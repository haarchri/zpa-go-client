// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServerGroupDTO server group d t o
//
// swagger:model ServerGroupDTO
type ServerGroupDTO struct {

	// app connector groups
	AppConnectorGroups []*AppConnectorGroup `json:"appConnectorGroups"`

	// applications
	Applications []*NameIDDto `json:"applications"`

	// config space
	// Enum: [DEFAULT SIEM]
	ConfigSpace string `json:"configSpace,omitempty"`

	// Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	CreationTime string `json:"creationTime,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// Defaults to false.
	DynamicDiscovery bool `json:"dynamicDiscovery,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	ID string `json:"id,omitempty"`

	// ip anchored
	IPAnchored bool `json:"ipAnchored,omitempty"`

	// Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	ModifiedTime string `json:"modifiedTime,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// servers
	Servers []*ApplicationServer `json:"servers"`
}

// Validate validates this server group d t o
func (m *ServerGroupDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppConnectorGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigSpace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerGroupDTO) validateAppConnectorGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.AppConnectorGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.AppConnectorGroups); i++ {
		if swag.IsZero(m.AppConnectorGroups[i]) { // not required
			continue
		}

		if m.AppConnectorGroups[i] != nil {
			if err := m.AppConnectorGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appConnectorGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("appConnectorGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerGroupDTO) validateApplications(formats strfmt.Registry) error {
	if swag.IsZero(m.Applications) { // not required
		return nil
	}

	for i := 0; i < len(m.Applications); i++ {
		if swag.IsZero(m.Applications[i]) { // not required
			continue
		}

		if m.Applications[i] != nil {
			if err := m.Applications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("applications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("applications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var serverGroupDTOTypeConfigSpacePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","SIEM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serverGroupDTOTypeConfigSpacePropEnum = append(serverGroupDTOTypeConfigSpacePropEnum, v)
	}
}

const (

	// ServerGroupDTOConfigSpaceDEFAULT captures enum value "DEFAULT"
	ServerGroupDTOConfigSpaceDEFAULT string = "DEFAULT"

	// ServerGroupDTOConfigSpaceSIEM captures enum value "SIEM"
	ServerGroupDTOConfigSpaceSIEM string = "SIEM"
)

// prop value enum
func (m *ServerGroupDTO) validateConfigSpaceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serverGroupDTOTypeConfigSpacePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServerGroupDTO) validateConfigSpace(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigSpace) { // not required
		return nil
	}

	// value enum
	if err := m.validateConfigSpaceEnum("configSpace", "body", m.ConfigSpace); err != nil {
		return err
	}

	return nil
}

func (m *ServerGroupDTO) validateServers(formats strfmt.Registry) error {
	if swag.IsZero(m.Servers) { // not required
		return nil
	}

	for i := 0; i < len(m.Servers); i++ {
		if swag.IsZero(m.Servers[i]) { // not required
			continue
		}

		if m.Servers[i] != nil {
			if err := m.Servers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this server group d t o based on the context it is used
func (m *ServerGroupDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppConnectorGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateApplications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServerGroupDTO) contextValidateAppConnectorGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AppConnectorGroups); i++ {

		if m.AppConnectorGroups[i] != nil {
			if err := m.AppConnectorGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("appConnectorGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("appConnectorGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerGroupDTO) contextValidateApplications(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Applications); i++ {

		if m.Applications[i] != nil {
			if err := m.Applications[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("applications" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("applications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServerGroupDTO) contextValidateServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Servers); i++ {

		if m.Servers[i] != nil {
			if err := m.Servers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("servers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("servers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServerGroupDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServerGroupDTO) UnmarshalBinary(b []byte) error {
	var res ServerGroupDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
