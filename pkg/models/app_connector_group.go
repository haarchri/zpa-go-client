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

// AppConnectorGroup app connector group
//
// swagger:model AppConnectorGroup
type AppConnectorGroup struct {

	// city country
	CityCountry string `json:"cityCountry,omitempty"`

	// connectors
	Connectors []*Connector `json:"connectors"`

	// country code
	CountryCode string `json:"countryCode,omitempty"`

	// creation time
	CreationTime string `json:"creationTime,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// dns query type
	// Enum: [IPV4_IPV6 IPV4 IPV6]
	DNSQueryType string `json:"dnsQueryType,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// geo location Id
	GeoLocationID string `json:"geoLocationId,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// latitude
	Latitude string `json:"latitude,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// longitude
	Longitude string `json:"longitude,omitempty"`

	// lss app connector group
	LssAppConnectorGroup bool `json:"lssAppConnectorGroup,omitempty"`

	// modified by
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// modified time
	ModifiedTime string `json:"modifiedTime,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// override version profile
	OverrideVersionProfile bool `json:"overrideVersionProfile,omitempty"`

	// server groups
	ServerGroups []*AppServerGroup `json:"serverGroups"`

	// upgrade day
	UpgradeDay string `json:"upgradeDay,omitempty"`

	// upgrade time in secs
	UpgradeTimeInSecs string `json:"upgradeTimeInSecs,omitempty"`

	// version profile Id
	VersionProfileID string `json:"versionProfileId,omitempty"`

	// Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	VersionProfileName string `json:"versionProfileName,omitempty"`

	// Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	// Enum: [ALL NONE CUSTOM]
	VersionProfileVisibilityScope string `json:"versionProfileVisibilityScope,omitempty"`
}

// Validate validates this app connector group
func (m *AppConnectorGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSQueryType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionProfileVisibilityScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppConnectorGroup) validateConnectors(formats strfmt.Registry) error {
	if swag.IsZero(m.Connectors) { // not required
		return nil
	}

	for i := 0; i < len(m.Connectors); i++ {
		if swag.IsZero(m.Connectors[i]) { // not required
			continue
		}

		if m.Connectors[i] != nil {
			if err := m.Connectors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connectors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var appConnectorGroupTypeDNSQueryTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IPV4_IPV6","IPV4","IPV6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		appConnectorGroupTypeDNSQueryTypePropEnum = append(appConnectorGroupTypeDNSQueryTypePropEnum, v)
	}
}

const (

	// AppConnectorGroupDNSQueryTypeIPV4IPV6 captures enum value "IPV4_IPV6"
	AppConnectorGroupDNSQueryTypeIPV4IPV6 string = "IPV4_IPV6"

	// AppConnectorGroupDNSQueryTypeIPV4 captures enum value "IPV4"
	AppConnectorGroupDNSQueryTypeIPV4 string = "IPV4"

	// AppConnectorGroupDNSQueryTypeIPV6 captures enum value "IPV6"
	AppConnectorGroupDNSQueryTypeIPV6 string = "IPV6"
)

// prop value enum
func (m *AppConnectorGroup) validateDNSQueryTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, appConnectorGroupTypeDNSQueryTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AppConnectorGroup) validateDNSQueryType(formats strfmt.Registry) error {
	if swag.IsZero(m.DNSQueryType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDNSQueryTypeEnum("dnsQueryType", "body", m.DNSQueryType); err != nil {
		return err
	}

	return nil
}

func (m *AppConnectorGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *AppConnectorGroup) validateServerGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.ServerGroups); i++ {
		if swag.IsZero(m.ServerGroups[i]) { // not required
			continue
		}

		if m.ServerGroups[i] != nil {
			if err := m.ServerGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serverGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serverGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var appConnectorGroupTypeVersionProfileVisibilityScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALL","NONE","CUSTOM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		appConnectorGroupTypeVersionProfileVisibilityScopePropEnum = append(appConnectorGroupTypeVersionProfileVisibilityScopePropEnum, v)
	}
}

const (

	// AppConnectorGroupVersionProfileVisibilityScopeALL captures enum value "ALL"
	AppConnectorGroupVersionProfileVisibilityScopeALL string = "ALL"

	// AppConnectorGroupVersionProfileVisibilityScopeNONE captures enum value "NONE"
	AppConnectorGroupVersionProfileVisibilityScopeNONE string = "NONE"

	// AppConnectorGroupVersionProfileVisibilityScopeCUSTOM captures enum value "CUSTOM"
	AppConnectorGroupVersionProfileVisibilityScopeCUSTOM string = "CUSTOM"
)

// prop value enum
func (m *AppConnectorGroup) validateVersionProfileVisibilityScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, appConnectorGroupTypeVersionProfileVisibilityScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AppConnectorGroup) validateVersionProfileVisibilityScope(formats strfmt.Registry) error {
	if swag.IsZero(m.VersionProfileVisibilityScope) { // not required
		return nil
	}

	// value enum
	if err := m.validateVersionProfileVisibilityScopeEnum("versionProfileVisibilityScope", "body", m.VersionProfileVisibilityScope); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this app connector group based on the context it is used
func (m *AppConnectorGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnectors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServerGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppConnectorGroup) contextValidateConnectors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Connectors); i++ {

		if m.Connectors[i] != nil {
			if err := m.Connectors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connectors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connectors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppConnectorGroup) contextValidateServerGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServerGroups); i++ {

		if m.ServerGroups[i] != nil {
			if err := m.ServerGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serverGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serverGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppConnectorGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppConnectorGroup) UnmarshalBinary(b []byte) error {
	var res AppConnectorGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
