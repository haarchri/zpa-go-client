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

// ApplicationResource application resource
//
// swagger:model ApplicationResource
type ApplicationResource struct {

	// bypass type
	// Enum: [ALWAYS NEVER ON_NET]
	BypassType string `json:"bypassType,omitempty"`

	// clientless apps
	ClientlessApps []*BAAppDto `json:"clientlessApps"`

	// common apps dto
	CommonAppsDto *CommonApplicationDto `json:"commonAppsDto,omitempty"`

	// config space
	// Enum: [DEFAULT SIEM]
	ConfigSpace string `json:"configSpace,omitempty"`

	// Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	CreationTime string `json:"creationTime,omitempty"`

	// default idle timeout
	DefaultIdleTimeout string `json:"defaultIdleTimeout,omitempty"`

	// default max age
	DefaultMaxAge string `json:"defaultMaxAge,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// domain names
	DomainNames []string `json:"domainNames"`

	// double encrypt
	DoubleEncrypt bool `json:"doubleEncrypt,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// health check type
	// Enum: [DEFAULT NONE]
	HealthCheckType string `json:"healthCheckType,omitempty"`

	// health reporting
	// Enum: [NONE ON_ACCESS CONTINUOUS]
	HealthReporting string `json:"healthReporting,omitempty"`

	// icmp access type
	// Enum: [PING_TRACEROUTING PING NONE]
	IcmpAccessType string `json:"icmpAccessType,omitempty"`

	// Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	ID string `json:"id,omitempty"`

	// inspection apps
	InspectionApps []*InspectAppDto `json:"inspectionApps"`

	// ip anchored
	IPAnchored bool `json:"ipAnchored,omitempty"`

	// is cname enabled
	IsCnameEnabled bool `json:"isCnameEnabled,omitempty"`

	// Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// Only applicable for a GET request. Ignored in PUT/POST/DELETE requests.
	ModifiedTime string `json:"modifiedTime,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// passive health enabled
	PassiveHealthEnabled bool `json:"passiveHealthEnabled,omitempty"`

	// segment group Id
	SegmentGroupID string `json:"segmentGroupId,omitempty"`

	// segment group name
	SegmentGroupName string `json:"segmentGroupName,omitempty"`

	// server groups
	ServerGroups []*AppServerGroup `json:"serverGroups"`

	// tcp port range
	TCPPortRange []*AppPortRange `json:"tcpPortRange"`

	// tcp port ranges
	TCPPortRanges []string `json:"tcpPortRanges"`

	// udp port range
	UDPPortRange []*AppPortRange `json:"udpPortRange"`

	// udp port ranges
	UDPPortRanges []string `json:"udpPortRanges"`
}

// Validate validates this application resource
func (m *ApplicationResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBypassType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientlessApps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommonAppsDto(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigSpace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthCheckType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthReporting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIcmpAccessType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInspectionApps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTCPPortRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUDPPortRange(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var applicationResourceTypeBypassTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALWAYS","NEVER","ON_NET"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationResourceTypeBypassTypePropEnum = append(applicationResourceTypeBypassTypePropEnum, v)
	}
}

const (

	// ApplicationResourceBypassTypeALWAYS captures enum value "ALWAYS"
	ApplicationResourceBypassTypeALWAYS string = "ALWAYS"

	// ApplicationResourceBypassTypeNEVER captures enum value "NEVER"
	ApplicationResourceBypassTypeNEVER string = "NEVER"

	// ApplicationResourceBypassTypeONNET captures enum value "ON_NET"
	ApplicationResourceBypassTypeONNET string = "ON_NET"
)

// prop value enum
func (m *ApplicationResource) validateBypassTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, applicationResourceTypeBypassTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ApplicationResource) validateBypassType(formats strfmt.Registry) error {
	if swag.IsZero(m.BypassType) { // not required
		return nil
	}

	// value enum
	if err := m.validateBypassTypeEnum("bypassType", "body", m.BypassType); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationResource) validateClientlessApps(formats strfmt.Registry) error {
	if swag.IsZero(m.ClientlessApps) { // not required
		return nil
	}

	for i := 0; i < len(m.ClientlessApps); i++ {
		if swag.IsZero(m.ClientlessApps[i]) { // not required
			continue
		}

		if m.ClientlessApps[i] != nil {
			if err := m.ClientlessApps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clientlessApps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clientlessApps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationResource) validateCommonAppsDto(formats strfmt.Registry) error {
	if swag.IsZero(m.CommonAppsDto) { // not required
		return nil
	}

	if m.CommonAppsDto != nil {
		if err := m.CommonAppsDto.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commonAppsDto")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commonAppsDto")
			}
			return err
		}
	}

	return nil
}

var applicationResourceTypeConfigSpacePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","SIEM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationResourceTypeConfigSpacePropEnum = append(applicationResourceTypeConfigSpacePropEnum, v)
	}
}

const (

	// ApplicationResourceConfigSpaceDEFAULT captures enum value "DEFAULT"
	ApplicationResourceConfigSpaceDEFAULT string = "DEFAULT"

	// ApplicationResourceConfigSpaceSIEM captures enum value "SIEM"
	ApplicationResourceConfigSpaceSIEM string = "SIEM"
)

// prop value enum
func (m *ApplicationResource) validateConfigSpaceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, applicationResourceTypeConfigSpacePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ApplicationResource) validateConfigSpace(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigSpace) { // not required
		return nil
	}

	// value enum
	if err := m.validateConfigSpaceEnum("configSpace", "body", m.ConfigSpace); err != nil {
		return err
	}

	return nil
}

var applicationResourceTypeHealthCheckTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","NONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationResourceTypeHealthCheckTypePropEnum = append(applicationResourceTypeHealthCheckTypePropEnum, v)
	}
}

const (

	// ApplicationResourceHealthCheckTypeDEFAULT captures enum value "DEFAULT"
	ApplicationResourceHealthCheckTypeDEFAULT string = "DEFAULT"

	// ApplicationResourceHealthCheckTypeNONE captures enum value "NONE"
	ApplicationResourceHealthCheckTypeNONE string = "NONE"
)

// prop value enum
func (m *ApplicationResource) validateHealthCheckTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, applicationResourceTypeHealthCheckTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ApplicationResource) validateHealthCheckType(formats strfmt.Registry) error {
	if swag.IsZero(m.HealthCheckType) { // not required
		return nil
	}

	// value enum
	if err := m.validateHealthCheckTypeEnum("healthCheckType", "body", m.HealthCheckType); err != nil {
		return err
	}

	return nil
}

var applicationResourceTypeHealthReportingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","ON_ACCESS","CONTINUOUS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationResourceTypeHealthReportingPropEnum = append(applicationResourceTypeHealthReportingPropEnum, v)
	}
}

const (

	// ApplicationResourceHealthReportingNONE captures enum value "NONE"
	ApplicationResourceHealthReportingNONE string = "NONE"

	// ApplicationResourceHealthReportingONACCESS captures enum value "ON_ACCESS"
	ApplicationResourceHealthReportingONACCESS string = "ON_ACCESS"

	// ApplicationResourceHealthReportingCONTINUOUS captures enum value "CONTINUOUS"
	ApplicationResourceHealthReportingCONTINUOUS string = "CONTINUOUS"
)

// prop value enum
func (m *ApplicationResource) validateHealthReportingEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, applicationResourceTypeHealthReportingPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ApplicationResource) validateHealthReporting(formats strfmt.Registry) error {
	if swag.IsZero(m.HealthReporting) { // not required
		return nil
	}

	// value enum
	if err := m.validateHealthReportingEnum("healthReporting", "body", m.HealthReporting); err != nil {
		return err
	}

	return nil
}

var applicationResourceTypeIcmpAccessTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PING_TRACEROUTING","PING","NONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationResourceTypeIcmpAccessTypePropEnum = append(applicationResourceTypeIcmpAccessTypePropEnum, v)
	}
}

const (

	// ApplicationResourceIcmpAccessTypePINGTRACEROUTING captures enum value "PING_TRACEROUTING"
	ApplicationResourceIcmpAccessTypePINGTRACEROUTING string = "PING_TRACEROUTING"

	// ApplicationResourceIcmpAccessTypePING captures enum value "PING"
	ApplicationResourceIcmpAccessTypePING string = "PING"

	// ApplicationResourceIcmpAccessTypeNONE captures enum value "NONE"
	ApplicationResourceIcmpAccessTypeNONE string = "NONE"
)

// prop value enum
func (m *ApplicationResource) validateIcmpAccessTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, applicationResourceTypeIcmpAccessTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ApplicationResource) validateIcmpAccessType(formats strfmt.Registry) error {
	if swag.IsZero(m.IcmpAccessType) { // not required
		return nil
	}

	// value enum
	if err := m.validateIcmpAccessTypeEnum("icmpAccessType", "body", m.IcmpAccessType); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationResource) validateInspectionApps(formats strfmt.Registry) error {
	if swag.IsZero(m.InspectionApps) { // not required
		return nil
	}

	for i := 0; i < len(m.InspectionApps); i++ {
		if swag.IsZero(m.InspectionApps[i]) { // not required
			continue
		}

		if m.InspectionApps[i] != nil {
			if err := m.InspectionApps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inspectionApps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inspectionApps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationResource) validateServerGroups(formats strfmt.Registry) error {
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

func (m *ApplicationResource) validateTCPPortRange(formats strfmt.Registry) error {
	if swag.IsZero(m.TCPPortRange) { // not required
		return nil
	}

	for i := 0; i < len(m.TCPPortRange); i++ {
		if swag.IsZero(m.TCPPortRange[i]) { // not required
			continue
		}

		if m.TCPPortRange[i] != nil {
			if err := m.TCPPortRange[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tcpPortRange" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tcpPortRange" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationResource) validateUDPPortRange(formats strfmt.Registry) error {
	if swag.IsZero(m.UDPPortRange) { // not required
		return nil
	}

	for i := 0; i < len(m.UDPPortRange); i++ {
		if swag.IsZero(m.UDPPortRange[i]) { // not required
			continue
		}

		if m.UDPPortRange[i] != nil {
			if err := m.UDPPortRange[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("udpPortRange" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("udpPortRange" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this application resource based on the context it is used
func (m *ApplicationResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClientlessApps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCommonAppsDto(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInspectionApps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServerGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTCPPortRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUDPPortRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationResource) contextValidateClientlessApps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClientlessApps); i++ {

		if m.ClientlessApps[i] != nil {
			if err := m.ClientlessApps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clientlessApps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clientlessApps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationResource) contextValidateCommonAppsDto(ctx context.Context, formats strfmt.Registry) error {

	if m.CommonAppsDto != nil {
		if err := m.CommonAppsDto.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commonAppsDto")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commonAppsDto")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationResource) contextValidateInspectionApps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InspectionApps); i++ {

		if m.InspectionApps[i] != nil {
			if err := m.InspectionApps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("inspectionApps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("inspectionApps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationResource) contextValidateServerGroups(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ApplicationResource) contextValidateTCPPortRange(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TCPPortRange); i++ {

		if m.TCPPortRange[i] != nil {
			if err := m.TCPPortRange[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tcpPortRange" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tcpPortRange" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ApplicationResource) contextValidateUDPPortRange(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UDPPortRange); i++ {

		if m.UDPPortRange[i] != nil {
			if err := m.UDPPortRange[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("udpPortRange" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("udpPortRange" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationResource) UnmarshalBinary(b []byte) error {
	var res ApplicationResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
