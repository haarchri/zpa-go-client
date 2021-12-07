


# Zscaler Private Access API Portal
To access detailed ZPA API documentation, including references and use cases, refer to the [Zscaler Help Portal](https://help.zscaler.com/zpa/publicapi/about-zpa-api)
  

## Informations

### Version

1.0

## Tags

  ### <span id="tag-app-server-controller"></span>app-server-controller

Server resource endpoints

  ### <span id="tag-application-controller"></span>application-controller

Application Segment resource endpoints

  ### <span id="tag-segment-group-controller"></span>segment-group-controller

Segment Group resource endpoints

  ### <span id="tag-connector-controller"></span>connector-controller

Connector resource endpoints

  ### <span id="tag-connector-group-controller"></span>connector-group-controller

App Connector Group resource endpoints

  ### <span id="tag-ba-certificate-controller"></span>ba-certificate-controller

Browser Access certificate related endpoints

  ### <span id="tag-ba-certificate-controller-v-2"></span>ba-certificate-controller-v-2

Browser Access certificate related v2 endpoints

  ### <span id="tag-customer-version-profile-controller"></span>customer-version-profile-controller

Customer Version profile resource endpoints

  ### <span id="tag-cloud-connector-group-controller"></span>cloud-connector-group-controller

Cloud Connector Group resource endpoints

  ### <span id="tag-idp-controller"></span>idp-controller

IdP resource endpoints

  ### <span id="tag-idp-controller-v-2"></span>idp-controller-v-2

IdP resource v2 endpoints

  ### <span id="tag-machine-group-controller"></span>machine-group-controller

Machine Group resource endpoints

  ### <span id="tag-provisioning-key-controller"></span>provisioningKey-controller

Provisioning key endpoints

  ### <span id="tag-policy-set-controller"></span>policy-set-controller

ZPA policy endpoints

  ### <span id="tag-posture-profile-controller"></span>posture-profile-controller

Posture profile endpoints

  ### <span id="tag-posture-profile-controller-v-2"></span>posture-profile-controller-v-2

Posture profile v2 endpoints

  ### <span id="tag-service-edge-controller"></span>service-edge-controller

Service Edge resource endpoints

  ### <span id="tag-service-edge-group-controller"></span>service-edge-group-controller

Service Edge group resource endpoints

  ### <span id="tag-saml-attr-controller"></span>saml-attr-controller

SAML attribute endpoints

  ### <span id="tag-saml-attr-controller-v-2"></span>saml-attr-controller-v-2

SAML attribute v2 endpoints

  ### <span id="tag-scim-attribute-header-controller"></span>scim-attribute-header-controller

SCIM attribute endpoints

  ### <span id="tag-server-group-controller"></span>server-group-controller

Server Group endpoints

  ### <span id="tag-lss-config-controller-v-2"></span>lss-config-controller-v-2

LSS V2 endpoints

  ### <span id="tag-enrollment-cert-controller"></span>enrollment-cert-controller

Signing Certificate related endpoints

  ### <span id="tag-enrollment-cert-controller-v-2"></span>enrollment-cert-controller-v-2

Signing certificate related v2 endpoints

  ### <span id="tag-trusted-network-controller"></span>trusted-network-controller

Trusted Network endpoints

  ### <span id="tag-trusted-network-controller-v-2"></span>trusted-network-controller-v-2

Trusted Network v2 endpoints

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json

### Produces
  * */*

## All endpoints

###  app_server_controller

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /mgmtconfig/v1/admin/customers/{customerId}/server | [add app server using p o s t 1](#add-app-server-using-p-o-s-t-1) | Add a new Server. |
| DELETE | /mgmtconfig/v1/admin/customers/{customerId}/server/{serverId} | [delete app server using d e l e t e 1](#delete-app-server-using-d-e-l-e-t-e-1) | Delete a Server. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/server | [get all app servers using g e t 1](#get-all-app-servers-using-g-e-t-1) | Get all the configured Servers. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/server/{serverId} | [get app server using g e t 1](#get-app-server-using-g-e-t-1) | Get the Server details. |
| PUT | /mgmtconfig/v1/admin/customers/{customerId}/server/{serverId} | [update app server using p u t 1](#update-app-server-using-p-u-t-1) | Update the Server details. |
  


###  application_controller

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /mgmtconfig/v1/admin/customers/{customerId}/application | [add application using p o s t 1](#add-application-using-p-o-s-t-1) | Add a new Application Segment. |
| DELETE | /mgmtconfig/v1/admin/customers/{customerId}/application/{applicationId} | [delete application using d e l e t e 1](#delete-application-using-d-e-l-e-t-e-1) | Delete an Application Segment. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/application | [get all applications using g e t 3](#get-all-applications-using-g-e-t-3) | Get all configured Application Segments. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/application/{applicationId} | [get application using g e t 1](#get-application-using-g-e-t-1) | Get the Application Segment details. |
| PUT | /mgmtconfig/v1/admin/customers/{customerId}/application/{applicationId} | [update application v2 using p u t 1](#update-application-v2-using-p-u-t-1) | Update the Application Segment details. |
  


###  ba_certificate_controller

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /mgmtconfig/v1/admin/customers/{customerId}/clientlessCertificate/issued | [get all issued certs using g e t 1](#get-all-issued-certs-using-g-e-t-1) | Get all the issued certificates. This API will be deprecated in a future release. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/clientlessCertificate/{certificateId} | [get using g e t 1](#get-using-g-e-t-1) | Get the details of the Browser Access certificate. |
  


###  ba_certificate_controller_v_2

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /mgmtconfig/v2/admin/customers/{customerId}/clientlessCertificate/issued | [get all issued certs using g e t 2](#get-all-issued-certs-using-g-e-t-2) | Get all the issued certificates. |
  


###  cloud_connector_group_controller

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /mgmtconfig/v1/admin/customers/{customerId}/cloudConnectorGroup/{id} | [get cloud connector group using g e t](#get-cloud-connector-group-using-g-e-t) | Get the Cloud Connector Group details. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/cloudConnectorGroup | [get cloud connector groups using g e t](#get-cloud-connector-groups-using-g-e-t) | Get all configured Cloud Connector Groups. |
  


###  connector_controller

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /mgmtconfig/v1/admin/customers/{customerId}/connector/bulkDelete | [bulk delete connector using p o s t 1](#bulk-delete-connector-using-p-o-s-t-1) | Bulk delete the Connectors. |
| DELETE | /mgmtconfig/v1/admin/customers/{customerId}/connector/{connectorId} | [delete connector using d e l e t e 1](#delete-connector-using-d-e-l-e-t-e-1) | Delete the Connector. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/connector | [get all connectors using g e t 1](#get-all-connectors-using-g-e-t-1) | Get all the configured Connector details. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/connector/{connectorId} | [get connector using g e t 1](#get-connector-using-g-e-t-1) | Get the Connector details. |
| PUT | /mgmtconfig/v1/admin/customers/{customerId}/connector/{connectorId} | [update connector using p u t 1](#update-connector-using-p-u-t-1) | Update the Connector details. |
  


###  connector_group_controller

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /mgmtconfig/v1/admin/customers/{customerId}/appConnectorGroup | [add app connector group using p o s t 1](#add-app-connector-group-using-p-o-s-t-1) | Add a new Connector Group. |
| DELETE | /mgmtconfig/v1/admin/customers/{customerId}/appConnectorGroup/{appConnectorGroupId} | [delete app connector group using d e l e t e 1](#delete-app-connector-group-using-d-e-l-e-t-e-1) | Delete the Connector Group. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/appConnectorGroup/{appConnectorGroupId} | [get app connector group using g e t 1](#get-app-connector-group-using-g-e-t-1) | Get the App Connector Group details. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/appConnectorGroup | [get app connector groups using g e t 1](#get-app-connector-groups-using-g-e-t-1) | Get all configured App Connector Groups. |
| PUT | /mgmtconfig/v1/admin/customers/{customerId}/appConnectorGroup/{appConnectorGroupId} | [update app connector group using p u t 1](#update-app-connector-group-using-p-u-t-1) | Update the Connector Group details. |
  


###  customer_version_profile_controller

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /mgmtconfig/v1/admin/customers/{customerId}/visible/versionProfiles | [get all version profiles visibile by customer Id using g e t](#get-all-version-profiles-visibile-by-customer-id-using-g-e-t) | Get Version Profiles visible to a customer |
  


###  enrollment_cert_controller

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /mgmtconfig/v1/admin/customers/{customerId}/enrollmentCert/{enrollmentCertId} | [get enrollment cert using g e t 3](#get-enrollment-cert-using-g-e-t-3) | Get the Enrollment Cert details. |
  


###  enrollment_cert_controller_v_2

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /mgmtconfig/v2/admin/customers/{customerId}/enrollmentCert | [get all signing cert using g e t 2](#get-all-signing-cert-using-g-e-t-2) | Get all the configured Enrollment Cert details. |
  


###  idp_controller

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /mgmtconfig/v1/admin/customers/{customerId}/idp/{idpId} | [get idp by Id using g e t 1](#get-idp-by-id-using-g-e-t-1) | Get all the IdP details. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/idp | [get idp using g e t 1](#get-idp-using-g-e-t-1) | Get the configured IdP details. This API will be deprecated in a future release. |
  


###  idp_controller_v_2

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /mgmtconfig/v2/admin/customers/{customerId}/idp | [get all idp using g e t](#get-all-idp-using-g-e-t) | Get the configured IdP details. |
  


###  lss_config_controller_v_2

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /mgmtconfig/v2/admin/customers/{customerId}/lssConfig | [add lss using p o s t 3](#add-lss-using-p-o-s-t-3) | Add a new LSS for a given customer. |
| DELETE | /mgmtconfig/v2/admin/customers/{customerId}/lssConfig/{lssId} | [delete lss using d e l e t e 3](#delete-lss-using-d-e-l-e-t-e-3) | Delete a LSS attribute. |
| GET | /mgmtconfig/v2/admin/customers/{customerId}/lssConfig | [get all using g e t 8](#get-all-using-g-e-t-8) | Get all LSS configured for a given customer |
| GET | /mgmtconfig/v2/admin/lssConfig/clientTypes | [get list of lss client types using g e t 1](#get-list-of-lss-client-types-using-g-e-t-1) | Get list of all LSS client types. |
| GET | /mgmtconfig/v2/admin/lssConfig/logType/formats | [get lss log format using g e t 3](#get-lss-log-format-using-g-e-t-3) | Get all LSS log format. |
| GET | /mgmtconfig/v2/admin/customers/{customerId}/lssConfig/{lssId} | [get lss using g e t 3](#get-lss-using-g-e-t-3) | Get the LSS details |
| GET | /mgmtconfig/v2/admin/lssConfig/statusCodes | [get status codes using g e t 1](#get-status-codes-using-g-e-t-1) | Get list of LSS status codes. |
| PUT | /mgmtconfig/v2/admin/customers/{customerId}/lssConfig/{lssId} | [update lss using p u t 3](#update-lss-using-p-u-t-3) | Update a LSS Attribute. |
  


###  machine_group_controller

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /mgmtconfig/v1/admin/customers/{customerId}/machineGroup/{Id} | [get machine group using g e t](#get-machine-group-using-g-e-t) | Get the Machine Group details. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/machineGroup | [get machine groups using g e t](#get-machine-groups-using-g-e-t) | Get all the configured Machine Groups. |
  


###  policy_set_controller

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /mgmtconfig/v1/admin/customers/{customerId}/policySet/{policySetId}/rule | [add rule to policy set using p o s t 1](#add-rule-to-policy-set-using-p-o-s-t-1) | Add a new policy rule for a given policy. |
| DELETE | /mgmtconfig/v1/admin/customers/{customerId}/policySet/{policySetId}/rule/{ruleId} | [delete rule in policy set using d e l e t e 1](#delete-rule-in-policy-set-using-d-e-l-e-t-e-1) | Delete a rule in a policy. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/policySet/bypass | [get bypass policy set using g e t 1](#get-bypass-policy-set-using-g-e-t-1) | Get the bypass policy and all rules for a Client Forwarding policy rule. This API will be deprecated in a future release. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/policySet/global | [get global policy set using g e t 1](#get-global-policy-set-using-g-e-t-1) | Get the global policy. This API will be deprecated in a future release. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/policySet/rules/policyType/{policyType} | [get policy rules by page using g e t 1](#get-policy-rules-by-page-using-g-e-t-1) | Get paginated policy rules for a given policy type. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/policySet/policyType/{policyType} | [get policy set by policy type using g e t 1](#get-policy-set-by-policy-type-using-g-e-t-1) | For a customer, get policy set by policy type |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/policySet/reauth | [get reauth policy set using g e t 1](#get-reauth-policy-set-using-g-e-t-1) | Get the authentication policy and all rules for a Timeout policy rule. This API will be deprecated in a future release. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/policySet/{policySetId}/rule/{ruleId} | [get rule in policy set using g e t 1](#get-rule-in-policy-set-using-g-e-t-1) | Get a rule in a policy. |
| PUT | /mgmtconfig/v1/admin/customers/{customerId}/policySet/{policySetId}/rule/{ruleId}/reorder/{newOrder} | [re order policy rule using p u t 1](#re-order-policy-rule-using-p-u-t-1) | Update rule order. |
| PUT | /mgmtconfig/v1/admin/customers/{customerId}/policySet/{policySetId}/rule/{ruleId} | [update rule to policy set using p u t 1](#update-rule-to-policy-set-using-p-u-t-1) | Update a rule in a policy. |
  


###  posture_profile_controller

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /mgmtconfig/v1/admin/customers/{customerId}/posture | [get all attributes using g e t 1](#get-all-attributes-using-g-e-t-1) | Get all posture profiles. This API will be deprecated in a future release. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/posture/{id} | [get posture profile using g e t 1](#get-posture-profile-using-g-e-t-1) | Get the configured profile posture. |
  


###  posture_profile_controller_v_2

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /mgmtconfig/v2/admin/customers/{customerId}/posture | [get all attributes using g e t 2](#get-all-attributes-using-g-e-t-2) | Get all posture profiles. |
  


###  provisioning_key_controller

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /mgmtconfig/v1/admin/customers/{customerId}/associationType/{associationType}/provisioningKey | [create provisioning key using p o s t 1](#create-provisioning-key-using-p-o-s-t-1) | Add a new Provisioning Key. |
| DELETE | /mgmtconfig/v1/admin/customers/{customerId}/associationType/{associationType}/provisioningKey/{provisioningKeyId} | [delete provisioning key using d e l e t e 1](#delete-provisioning-key-using-d-e-l-e-t-e-1) | Delete the Provisioning Key. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/associationType/{associationType}/provisioningKey | [get provisioning key for association type using g e t 1](#get-provisioning-key-for-association-type-using-g-e-t-1) | Get all the configured Provisioning Key details. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/associationType/{associationType}/provisioningKey/{provisioningKeyId} | [get provisioning key using g e t 1](#get-provisioning-key-using-g-e-t-1) | Get the Provisioning Key details. |
| PUT | /mgmtconfig/v1/admin/customers/{customerId}/associationType/{associationType}/provisioningKey/{provisioningKeyId} | [update provisioning key using p u t 1](#update-provisioning-key-using-p-u-t-1) | Update the Provisioning Key details. |
  


###  saml_attr_controller

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /mgmtconfig/v1/admin/customers/{customerId}/samlAttribute | [get all attributes using g e t 4](#get-all-attributes-using-g-e-t-4) | Get all SAML attributes. This API will be deprecated in a future release. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/samlAttribute/{attrId} | [get saml attribute using g e t 1](#get-saml-attribute-using-g-e-t-1) | Get the SAML attribute details. |
  


###  saml_attr_controller_v_2

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /mgmtconfig/v2/admin/customers/{customerId}/samlAttribute/idp/{idpId} | [get all attributes by idp Id and page using g e t](#get-all-attributes-by-idp-id-and-page-using-g-e-t) | Get all attributes configured for a given customer |
| GET | /mgmtconfig/v2/admin/customers/{customerId}/samlAttribute | [get all attributes by page using g e t](#get-all-attributes-by-page-using-g-e-t) | Get all SAML attributes by page. |
  


###  scim_attribute_header_controller

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /mgmtconfig/v1/admin/customers/{customerId}/idp/{idpId}/scimattribute | [get all s c i m attributes using g e t 1](#get-all-s-c-i-m-attributes-using-g-e-t-1) | Get all the SCIM attributes. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/idp/{idpId}/scimattribute/{scimAttributeId} | [get s c i m attribute using g e t 1](#get-s-c-i-m-attribute-using-g-e-t-1) | Get the SCIM attribute details. |
  


###  segment_group_controller

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup | [add segment group using p o s t 1](#add-segment-group-using-p-o-s-t-1) | Add a new Segment Group. |
| DELETE | /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup/{segmentGroupId} | [delete segment group using d e l e t e 1](#delete-segment-group-using-d-e-l-e-t-e-1) | Delete a Segment Group. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup | [get all segment groups using g e t 1](#get-all-segment-groups-using-g-e-t-1) | Get all the configured Segment Groups. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup/{segmentGroupId} | [get segment group using g e t 1](#get-segment-group-using-g-e-t-1) | Get the Segment Group details. |
| PUT | /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup/{segmentGroupId} | [update segment group using p u t 1](#update-segment-group-using-p-u-t-1) | Update a Segment Group. |
  


###  server_group_controller

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /mgmtconfig/v1/admin/customers/{customerId}/serverGroup | [add app server group using p o s t 1](#add-app-server-group-using-p-o-s-t-1) | Add a new Server Group. |
| DELETE | /mgmtconfig/v1/admin/customers/{customerId}/serverGroup/{groupId} | [delete app server group using d e l e t e 1](#delete-app-server-group-using-d-e-l-e-t-e-1) | Delete a Server Group. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/serverGroup | [get all server groups using g e t 1](#get-all-server-groups-using-g-e-t-1) | Get all configured Server Groups. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/serverGroup/{groupId} | [get server group using g e t 1](#get-server-group-using-g-e-t-1) | Get the Server Group details. |
| PUT | /mgmtconfig/v1/admin/customers/{customerId}/serverGroup/{groupId} | [update app server group using p u t 1](#update-app-server-group-using-p-u-t-1) | Update the Server Group. |
  


###  service_edge_controller

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/bulkDelete | [bulk delete service edge using p o s t](#bulk-delete-service-edge-using-p-o-s-t) | Bulk delete the Service Edges. |
| DELETE | /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/{serviceEdgeId} | [delete service edge using d e l e t e](#delete-service-edge-using-d-e-l-e-t-e) | Delete the Service Edge. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge | [get all service edges using g e t](#get-all-service-edges-using-g-e-t) | Get all the configured Service Edge details. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/{serviceEdgeId} | [get service edge using g e t 1](#get-service-edge-using-g-e-t-1) | Get the Service Edge details. |
| PUT | /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/{serviceEdgeId} | [update service edge using p u t](#update-service-edge-using-p-u-t) | Update the Service Edge details. |
  


###  service_edge_group_controller

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /mgmtconfig/v1/admin/customers/{customerId}/serviceEdgeGroup | [add service edge group using p o s t](#add-service-edge-group-using-p-o-s-t) | Add a new Service Edge Group. |
| DELETE | /mgmtconfig/v1/admin/customers/{customerId}/serviceEdgeGroup/{serviceEdgeGroupId} | [delete service edge group using d e l e t e](#delete-service-edge-group-using-d-e-l-e-t-e) | Delete the Service Edge Group. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/serviceEdgeGroup/{serviceEdgeGroupId} | [get service edge group using g e t](#get-service-edge-group-using-g-e-t) | Get the Service Edge Group details. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/serviceEdgeGroup | [get service edge groups using g e t](#get-service-edge-groups-using-g-e-t) | Get all the configured Service Edge Group details. |
| PUT | /mgmtconfig/v1/admin/customers/{customerId}/serviceEdgeGroup/{serviceEdgeGroupId} | [update service edge group using p u t](#update-service-edge-group-using-p-u-t) | Update the Service Edge Group details. |
  


###  trusted_network_controller

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /mgmtconfig/v1/admin/customers/{customerId}/network | [get all trusted networks using g e t 1](#get-all-trusted-networks-using-g-e-t-1) | Get all the Trusted Networks. This API will be deprecated in a future release. |
| GET | /mgmtconfig/v1/admin/customers/{customerId}/network/{id} | [get trusted network using g e t 1](#get-trusted-network-using-g-e-t-1) | Get the Trusted Networks. |
  


###  trusted_network_controller_v_2

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /mgmtconfig/v2/admin/customers/{customerId}/network | [get all trusted networks using g e t 2](#get-all-trusted-networks-using-g-e-t-2) | Get all the Trusted Networks. |
  


## Paths

### <span id="add-app-connector-group-using-p-o-s-t-1"></span> Add a new Connector Group. (*addAppConnectorGroupUsingPOST_1*)

```
POST /mgmtconfig/v1/admin/customers/{customerId}/appConnectorGroup
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| appConnectorGroup | `body` | [AppConnectorGroup](#app-connector-group) | `models.AppConnectorGroup` | | ✓ | | The object of the Connector Group. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#add-app-connector-group-using-p-o-s-t-1-201) | Created | Created |  | [schema](#add-app-connector-group-using-p-o-s-t-1-201-schema) |
| [400](#add-app-connector-group-using-p-o-s-t-1-400) | Bad Request | BadRequest |  | [schema](#add-app-connector-group-using-p-o-s-t-1-400-schema) |
| [401](#add-app-connector-group-using-p-o-s-t-1-401) | Unauthorized | Unauthorized |  | [schema](#add-app-connector-group-using-p-o-s-t-1-401-schema) |
| [403](#add-app-connector-group-using-p-o-s-t-1-403) | Forbidden | Forbidden |  | [schema](#add-app-connector-group-using-p-o-s-t-1-403-schema) |
| [404](#add-app-connector-group-using-p-o-s-t-1-404) | Not Found | Not Found |  | [schema](#add-app-connector-group-using-p-o-s-t-1-404-schema) |
| [429](#add-app-connector-group-using-p-o-s-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#add-app-connector-group-using-p-o-s-t-1-429-schema) |

#### Responses


##### <span id="add-app-connector-group-using-p-o-s-t-1-201"></span> 201 - Created
Status: Created

###### <span id="add-app-connector-group-using-p-o-s-t-1-201-schema"></span> Schema
   
  

[AppConnectorGroup](#app-connector-group)

##### <span id="add-app-connector-group-using-p-o-s-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="add-app-connector-group-using-p-o-s-t-1-400-schema"></span> Schema

##### <span id="add-app-connector-group-using-p-o-s-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="add-app-connector-group-using-p-o-s-t-1-401-schema"></span> Schema

##### <span id="add-app-connector-group-using-p-o-s-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="add-app-connector-group-using-p-o-s-t-1-403-schema"></span> Schema

##### <span id="add-app-connector-group-using-p-o-s-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="add-app-connector-group-using-p-o-s-t-1-404-schema"></span> Schema

##### <span id="add-app-connector-group-using-p-o-s-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="add-app-connector-group-using-p-o-s-t-1-429-schema"></span> Schema

### <span id="add-app-server-group-using-p-o-s-t-1"></span> Add a new Server Group. (*addAppServerGroupUsingPOST_1*)

```
POST /mgmtconfig/v1/admin/customers/{customerId}/serverGroup
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| group | `body` | [ServerGroupDTO](#server-group-d-t-o) | `models.ServerGroupDTO` | |  | | The object of the Server Group. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#add-app-server-group-using-p-o-s-t-1-201) | Created | Created |  | [schema](#add-app-server-group-using-p-o-s-t-1-201-schema) |
| [400](#add-app-server-group-using-p-o-s-t-1-400) | Bad Request | BadRequest |  | [schema](#add-app-server-group-using-p-o-s-t-1-400-schema) |
| [401](#add-app-server-group-using-p-o-s-t-1-401) | Unauthorized | Unauthorized |  | [schema](#add-app-server-group-using-p-o-s-t-1-401-schema) |
| [403](#add-app-server-group-using-p-o-s-t-1-403) | Forbidden | Forbidden |  | [schema](#add-app-server-group-using-p-o-s-t-1-403-schema) |
| [404](#add-app-server-group-using-p-o-s-t-1-404) | Not Found | Not Found |  | [schema](#add-app-server-group-using-p-o-s-t-1-404-schema) |
| [429](#add-app-server-group-using-p-o-s-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#add-app-server-group-using-p-o-s-t-1-429-schema) |

#### Responses


##### <span id="add-app-server-group-using-p-o-s-t-1-201"></span> 201 - Created
Status: Created

###### <span id="add-app-server-group-using-p-o-s-t-1-201-schema"></span> Schema
   
  

[ServerGroupDTO](#server-group-d-t-o)

##### <span id="add-app-server-group-using-p-o-s-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="add-app-server-group-using-p-o-s-t-1-400-schema"></span> Schema

##### <span id="add-app-server-group-using-p-o-s-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="add-app-server-group-using-p-o-s-t-1-401-schema"></span> Schema

##### <span id="add-app-server-group-using-p-o-s-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="add-app-server-group-using-p-o-s-t-1-403-schema"></span> Schema

##### <span id="add-app-server-group-using-p-o-s-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="add-app-server-group-using-p-o-s-t-1-404-schema"></span> Schema

##### <span id="add-app-server-group-using-p-o-s-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="add-app-server-group-using-p-o-s-t-1-429-schema"></span> Schema

### <span id="add-app-server-using-p-o-s-t-1"></span> Add a new Server. (*addAppServerUsingPOST_1*)

```
POST /mgmtconfig/v1/admin/customers/{customerId}/server
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| server | `body` | [ApplicationServer](#application-server) | `models.ApplicationServer` | |  | | The object of the Server. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#add-app-server-using-p-o-s-t-1-201) | Created | Created |  | [schema](#add-app-server-using-p-o-s-t-1-201-schema) |
| [400](#add-app-server-using-p-o-s-t-1-400) | Bad Request | BadRequest |  | [schema](#add-app-server-using-p-o-s-t-1-400-schema) |
| [401](#add-app-server-using-p-o-s-t-1-401) | Unauthorized | Unauthorized |  | [schema](#add-app-server-using-p-o-s-t-1-401-schema) |
| [403](#add-app-server-using-p-o-s-t-1-403) | Forbidden | Forbidden |  | [schema](#add-app-server-using-p-o-s-t-1-403-schema) |
| [404](#add-app-server-using-p-o-s-t-1-404) | Not Found | Not Found |  | [schema](#add-app-server-using-p-o-s-t-1-404-schema) |
| [429](#add-app-server-using-p-o-s-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#add-app-server-using-p-o-s-t-1-429-schema) |

#### Responses


##### <span id="add-app-server-using-p-o-s-t-1-201"></span> 201 - Created
Status: Created

###### <span id="add-app-server-using-p-o-s-t-1-201-schema"></span> Schema
   
  

[ApplicationServer](#application-server)

##### <span id="add-app-server-using-p-o-s-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="add-app-server-using-p-o-s-t-1-400-schema"></span> Schema

##### <span id="add-app-server-using-p-o-s-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="add-app-server-using-p-o-s-t-1-401-schema"></span> Schema

##### <span id="add-app-server-using-p-o-s-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="add-app-server-using-p-o-s-t-1-403-schema"></span> Schema

##### <span id="add-app-server-using-p-o-s-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="add-app-server-using-p-o-s-t-1-404-schema"></span> Schema

##### <span id="add-app-server-using-p-o-s-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="add-app-server-using-p-o-s-t-1-429-schema"></span> Schema

### <span id="add-application-using-p-o-s-t-1"></span> Add a new Application Segment. (*addApplicationUsingPOST_1*)

```
POST /mgmtconfig/v1/admin/customers/{customerId}/application
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| application | `body` | [ApplicationResource](#application-resource) | `models.ApplicationResource` | |  | | The object of the Application Segment. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#add-application-using-p-o-s-t-1-201) | Created | Created |  | [schema](#add-application-using-p-o-s-t-1-201-schema) |
| [400](#add-application-using-p-o-s-t-1-400) | Bad Request | BadRequest |  | [schema](#add-application-using-p-o-s-t-1-400-schema) |
| [401](#add-application-using-p-o-s-t-1-401) | Unauthorized | Unauthorized |  | [schema](#add-application-using-p-o-s-t-1-401-schema) |
| [403](#add-application-using-p-o-s-t-1-403) | Forbidden | Forbidden |  | [schema](#add-application-using-p-o-s-t-1-403-schema) |
| [404](#add-application-using-p-o-s-t-1-404) | Not Found | Not Found |  | [schema](#add-application-using-p-o-s-t-1-404-schema) |
| [429](#add-application-using-p-o-s-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#add-application-using-p-o-s-t-1-429-schema) |

#### Responses


##### <span id="add-application-using-p-o-s-t-1-201"></span> 201 - Created
Status: Created

###### <span id="add-application-using-p-o-s-t-1-201-schema"></span> Schema
   
  

[ApplicationResource](#application-resource)

##### <span id="add-application-using-p-o-s-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="add-application-using-p-o-s-t-1-400-schema"></span> Schema

##### <span id="add-application-using-p-o-s-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="add-application-using-p-o-s-t-1-401-schema"></span> Schema

##### <span id="add-application-using-p-o-s-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="add-application-using-p-o-s-t-1-403-schema"></span> Schema

##### <span id="add-application-using-p-o-s-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="add-application-using-p-o-s-t-1-404-schema"></span> Schema

##### <span id="add-application-using-p-o-s-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="add-application-using-p-o-s-t-1-429-schema"></span> Schema

### <span id="add-lss-using-p-o-s-t-3"></span> Add a new LSS for a given customer. (*addLssUsingPOST_3*)

```
POST /mgmtconfig/v2/admin/customers/{customerId}/lssConfig
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| lssResource | `body` | [LssResource](#lss-resource) | `models.LssResource` | |  | | The object of the LSS resource. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#add-lss-using-p-o-s-t-3-201) | Created | Created |  | [schema](#add-lss-using-p-o-s-t-3-201-schema) |
| [400](#add-lss-using-p-o-s-t-3-400) | Bad Request | BadRequest |  | [schema](#add-lss-using-p-o-s-t-3-400-schema) |
| [401](#add-lss-using-p-o-s-t-3-401) | Unauthorized | Unauthorized |  | [schema](#add-lss-using-p-o-s-t-3-401-schema) |
| [403](#add-lss-using-p-o-s-t-3-403) | Forbidden | Forbidden |  | [schema](#add-lss-using-p-o-s-t-3-403-schema) |
| [404](#add-lss-using-p-o-s-t-3-404) | Not Found | Not Found |  | [schema](#add-lss-using-p-o-s-t-3-404-schema) |
| [429](#add-lss-using-p-o-s-t-3-429) | Too Many Requests | TooManyRequest |  | [schema](#add-lss-using-p-o-s-t-3-429-schema) |

#### Responses


##### <span id="add-lss-using-p-o-s-t-3-201"></span> 201 - Created
Status: Created

###### <span id="add-lss-using-p-o-s-t-3-201-schema"></span> Schema
   
  

[LssResource](#lss-resource)

##### <span id="add-lss-using-p-o-s-t-3-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="add-lss-using-p-o-s-t-3-400-schema"></span> Schema

##### <span id="add-lss-using-p-o-s-t-3-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="add-lss-using-p-o-s-t-3-401-schema"></span> Schema

##### <span id="add-lss-using-p-o-s-t-3-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="add-lss-using-p-o-s-t-3-403-schema"></span> Schema

##### <span id="add-lss-using-p-o-s-t-3-404"></span> 404 - Not Found
Status: Not Found

###### <span id="add-lss-using-p-o-s-t-3-404-schema"></span> Schema

##### <span id="add-lss-using-p-o-s-t-3-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="add-lss-using-p-o-s-t-3-429-schema"></span> Schema

### <span id="add-rule-to-policy-set-using-p-o-s-t-1"></span> Add a new policy rule for a given policy. (*addRuleToPolicySetUsingPOST_1*)

```
POST /mgmtconfig/v1/admin/customers/{customerId}/policySet/{policySetId}/rule
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| policySetId | `path` | string | `string` |  | ✓ |  | The unique identifier of the policy. |
| rule | `body` | [PolicyRule](#policy-rule) | `models.PolicyRule` | | ✓ | | The object of the policy rule. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#add-rule-to-policy-set-using-p-o-s-t-1-201) | Created | Created |  | [schema](#add-rule-to-policy-set-using-p-o-s-t-1-201-schema) |
| [400](#add-rule-to-policy-set-using-p-o-s-t-1-400) | Bad Request | BadRequest |  | [schema](#add-rule-to-policy-set-using-p-o-s-t-1-400-schema) |
| [401](#add-rule-to-policy-set-using-p-o-s-t-1-401) | Unauthorized | Unauthorized |  | [schema](#add-rule-to-policy-set-using-p-o-s-t-1-401-schema) |
| [403](#add-rule-to-policy-set-using-p-o-s-t-1-403) | Forbidden | Forbidden |  | [schema](#add-rule-to-policy-set-using-p-o-s-t-1-403-schema) |
| [404](#add-rule-to-policy-set-using-p-o-s-t-1-404) | Not Found | Not Found |  | [schema](#add-rule-to-policy-set-using-p-o-s-t-1-404-schema) |
| [429](#add-rule-to-policy-set-using-p-o-s-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#add-rule-to-policy-set-using-p-o-s-t-1-429-schema) |

#### Responses


##### <span id="add-rule-to-policy-set-using-p-o-s-t-1-201"></span> 201 - Created
Status: Created

###### <span id="add-rule-to-policy-set-using-p-o-s-t-1-201-schema"></span> Schema
   
  

[PolicyRule](#policy-rule)

##### <span id="add-rule-to-policy-set-using-p-o-s-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="add-rule-to-policy-set-using-p-o-s-t-1-400-schema"></span> Schema

##### <span id="add-rule-to-policy-set-using-p-o-s-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="add-rule-to-policy-set-using-p-o-s-t-1-401-schema"></span> Schema

##### <span id="add-rule-to-policy-set-using-p-o-s-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="add-rule-to-policy-set-using-p-o-s-t-1-403-schema"></span> Schema

##### <span id="add-rule-to-policy-set-using-p-o-s-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="add-rule-to-policy-set-using-p-o-s-t-1-404-schema"></span> Schema

##### <span id="add-rule-to-policy-set-using-p-o-s-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="add-rule-to-policy-set-using-p-o-s-t-1-429-schema"></span> Schema

### <span id="add-segment-group-using-p-o-s-t-1"></span> Add a new Segment Group. (*addSegmentGroupUsingPOST_1*)

```
POST /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| segmentGroup | `body` | [SegmentGroup](#segment-group) | `models.SegmentGroup` | |  | | The object of the Segment Group. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#add-segment-group-using-p-o-s-t-1-201) | Created | Created |  | [schema](#add-segment-group-using-p-o-s-t-1-201-schema) |
| [400](#add-segment-group-using-p-o-s-t-1-400) | Bad Request | BadRequest |  | [schema](#add-segment-group-using-p-o-s-t-1-400-schema) |
| [401](#add-segment-group-using-p-o-s-t-1-401) | Unauthorized | Unauthorized |  | [schema](#add-segment-group-using-p-o-s-t-1-401-schema) |
| [403](#add-segment-group-using-p-o-s-t-1-403) | Forbidden | Forbidden |  | [schema](#add-segment-group-using-p-o-s-t-1-403-schema) |
| [404](#add-segment-group-using-p-o-s-t-1-404) | Not Found | Not Found |  | [schema](#add-segment-group-using-p-o-s-t-1-404-schema) |
| [429](#add-segment-group-using-p-o-s-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#add-segment-group-using-p-o-s-t-1-429-schema) |

#### Responses


##### <span id="add-segment-group-using-p-o-s-t-1-201"></span> 201 - Created
Status: Created

###### <span id="add-segment-group-using-p-o-s-t-1-201-schema"></span> Schema
   
  

[SegmentGroup](#segment-group)

##### <span id="add-segment-group-using-p-o-s-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="add-segment-group-using-p-o-s-t-1-400-schema"></span> Schema

##### <span id="add-segment-group-using-p-o-s-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="add-segment-group-using-p-o-s-t-1-401-schema"></span> Schema

##### <span id="add-segment-group-using-p-o-s-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="add-segment-group-using-p-o-s-t-1-403-schema"></span> Schema

##### <span id="add-segment-group-using-p-o-s-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="add-segment-group-using-p-o-s-t-1-404-schema"></span> Schema

##### <span id="add-segment-group-using-p-o-s-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="add-segment-group-using-p-o-s-t-1-429-schema"></span> Schema

### <span id="add-service-edge-group-using-p-o-s-t"></span> Add a new Service Edge Group. (*addServiceEdgeGroupUsingPOST*)

```
POST /mgmtconfig/v1/admin/customers/{customerId}/serviceEdgeGroup
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| serviceEdgeGroup | `body` | [ServiceEdgeGroup](#service-edge-group) | `models.ServiceEdgeGroup` | | ✓ | | The object of the Service Edge Group. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#add-service-edge-group-using-p-o-s-t-201) | Created | Created |  | [schema](#add-service-edge-group-using-p-o-s-t-201-schema) |
| [400](#add-service-edge-group-using-p-o-s-t-400) | Bad Request | BadRequest |  | [schema](#add-service-edge-group-using-p-o-s-t-400-schema) |
| [401](#add-service-edge-group-using-p-o-s-t-401) | Unauthorized | Unauthorized |  | [schema](#add-service-edge-group-using-p-o-s-t-401-schema) |
| [403](#add-service-edge-group-using-p-o-s-t-403) | Forbidden | Forbidden |  | [schema](#add-service-edge-group-using-p-o-s-t-403-schema) |
| [404](#add-service-edge-group-using-p-o-s-t-404) | Not Found | Not Found |  | [schema](#add-service-edge-group-using-p-o-s-t-404-schema) |
| [429](#add-service-edge-group-using-p-o-s-t-429) | Too Many Requests | TooManyRequest |  | [schema](#add-service-edge-group-using-p-o-s-t-429-schema) |

#### Responses


##### <span id="add-service-edge-group-using-p-o-s-t-201"></span> 201 - Created
Status: Created

###### <span id="add-service-edge-group-using-p-o-s-t-201-schema"></span> Schema
   
  

[ServiceEdgeGroup](#service-edge-group)

##### <span id="add-service-edge-group-using-p-o-s-t-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="add-service-edge-group-using-p-o-s-t-400-schema"></span> Schema

##### <span id="add-service-edge-group-using-p-o-s-t-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="add-service-edge-group-using-p-o-s-t-401-schema"></span> Schema

##### <span id="add-service-edge-group-using-p-o-s-t-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="add-service-edge-group-using-p-o-s-t-403-schema"></span> Schema

##### <span id="add-service-edge-group-using-p-o-s-t-404"></span> 404 - Not Found
Status: Not Found

###### <span id="add-service-edge-group-using-p-o-s-t-404-schema"></span> Schema

##### <span id="add-service-edge-group-using-p-o-s-t-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="add-service-edge-group-using-p-o-s-t-429-schema"></span> Schema

### <span id="bulk-delete-connector-using-p-o-s-t-1"></span> Bulk delete the Connectors. (*bulkDeleteConnectorUsingPOST_1*)

```
POST /mgmtconfig/v1/admin/customers/{customerId}/connector/bulkDelete
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| bulkDeleteResource | `body` | [BulkDeleteResource](#bulk-delete-resource) | `models.BulkDeleteResource` | | ✓ | | The resource for bulk deleting the Connectors. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#bulk-delete-connector-using-p-o-s-t-1-200) | OK | OK |  | [schema](#bulk-delete-connector-using-p-o-s-t-1-200-schema) |
| [201](#bulk-delete-connector-using-p-o-s-t-1-201) | Created | Created |  | [schema](#bulk-delete-connector-using-p-o-s-t-1-201-schema) |
| [400](#bulk-delete-connector-using-p-o-s-t-1-400) | Bad Request | BadRequest |  | [schema](#bulk-delete-connector-using-p-o-s-t-1-400-schema) |
| [401](#bulk-delete-connector-using-p-o-s-t-1-401) | Unauthorized | Unauthorized |  | [schema](#bulk-delete-connector-using-p-o-s-t-1-401-schema) |
| [403](#bulk-delete-connector-using-p-o-s-t-1-403) | Forbidden | Forbidden |  | [schema](#bulk-delete-connector-using-p-o-s-t-1-403-schema) |
| [404](#bulk-delete-connector-using-p-o-s-t-1-404) | Not Found | Not Found |  | [schema](#bulk-delete-connector-using-p-o-s-t-1-404-schema) |
| [429](#bulk-delete-connector-using-p-o-s-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#bulk-delete-connector-using-p-o-s-t-1-429-schema) |

#### Responses


##### <span id="bulk-delete-connector-using-p-o-s-t-1-200"></span> 200 - OK
Status: OK

###### <span id="bulk-delete-connector-using-p-o-s-t-1-200-schema"></span> Schema

##### <span id="bulk-delete-connector-using-p-o-s-t-1-201"></span> 201 - Created
Status: Created

###### <span id="bulk-delete-connector-using-p-o-s-t-1-201-schema"></span> Schema

##### <span id="bulk-delete-connector-using-p-o-s-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="bulk-delete-connector-using-p-o-s-t-1-400-schema"></span> Schema

##### <span id="bulk-delete-connector-using-p-o-s-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="bulk-delete-connector-using-p-o-s-t-1-401-schema"></span> Schema

##### <span id="bulk-delete-connector-using-p-o-s-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="bulk-delete-connector-using-p-o-s-t-1-403-schema"></span> Schema

##### <span id="bulk-delete-connector-using-p-o-s-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="bulk-delete-connector-using-p-o-s-t-1-404-schema"></span> Schema

##### <span id="bulk-delete-connector-using-p-o-s-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="bulk-delete-connector-using-p-o-s-t-1-429-schema"></span> Schema

### <span id="bulk-delete-service-edge-using-p-o-s-t"></span> Bulk delete the Service Edges. (*bulkDeleteServiceEdgeUsingPOST*)

```
POST /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/bulkDelete
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| bulkDeleteResource | `body` | [BulkDeleteResource](#bulk-delete-resource) | `models.BulkDeleteResource` | | ✓ | | The resource for bulk deleting the Service Edges. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#bulk-delete-service-edge-using-p-o-s-t-200) | OK | OK |  | [schema](#bulk-delete-service-edge-using-p-o-s-t-200-schema) |
| [201](#bulk-delete-service-edge-using-p-o-s-t-201) | Created | Created |  | [schema](#bulk-delete-service-edge-using-p-o-s-t-201-schema) |
| [400](#bulk-delete-service-edge-using-p-o-s-t-400) | Bad Request | BadRequest |  | [schema](#bulk-delete-service-edge-using-p-o-s-t-400-schema) |
| [401](#bulk-delete-service-edge-using-p-o-s-t-401) | Unauthorized | Unauthorized |  | [schema](#bulk-delete-service-edge-using-p-o-s-t-401-schema) |
| [403](#bulk-delete-service-edge-using-p-o-s-t-403) | Forbidden | Forbidden |  | [schema](#bulk-delete-service-edge-using-p-o-s-t-403-schema) |
| [404](#bulk-delete-service-edge-using-p-o-s-t-404) | Not Found | Not Found |  | [schema](#bulk-delete-service-edge-using-p-o-s-t-404-schema) |
| [429](#bulk-delete-service-edge-using-p-o-s-t-429) | Too Many Requests | TooManyRequest |  | [schema](#bulk-delete-service-edge-using-p-o-s-t-429-schema) |

#### Responses


##### <span id="bulk-delete-service-edge-using-p-o-s-t-200"></span> 200 - OK
Status: OK

###### <span id="bulk-delete-service-edge-using-p-o-s-t-200-schema"></span> Schema

##### <span id="bulk-delete-service-edge-using-p-o-s-t-201"></span> 201 - Created
Status: Created

###### <span id="bulk-delete-service-edge-using-p-o-s-t-201-schema"></span> Schema

##### <span id="bulk-delete-service-edge-using-p-o-s-t-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="bulk-delete-service-edge-using-p-o-s-t-400-schema"></span> Schema

##### <span id="bulk-delete-service-edge-using-p-o-s-t-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="bulk-delete-service-edge-using-p-o-s-t-401-schema"></span> Schema

##### <span id="bulk-delete-service-edge-using-p-o-s-t-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="bulk-delete-service-edge-using-p-o-s-t-403-schema"></span> Schema

##### <span id="bulk-delete-service-edge-using-p-o-s-t-404"></span> 404 - Not Found
Status: Not Found

###### <span id="bulk-delete-service-edge-using-p-o-s-t-404-schema"></span> Schema

##### <span id="bulk-delete-service-edge-using-p-o-s-t-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="bulk-delete-service-edge-using-p-o-s-t-429-schema"></span> Schema

### <span id="create-provisioning-key-using-p-o-s-t-1"></span> Add a new Provisioning Key. (*createProvisioningKeyUsingPOST_1*)

```
POST /mgmtconfig/v1/admin/customers/{customerId}/associationType/{associationType}/provisioningKey
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| associationType | `path` | string | `string` |  | ✓ |  | Specifies the provisioning key type for App Connectors or ZPA Private Service Edges. The supported values are CONNECTOR_GRP and SERVICE_EDGE_GRP. |
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| provisioningKey | `body` | [ProvisioningKey](#provisioning-key) | `models.ProvisioningKey` | | ✓ | | The object of the Provisioning Key. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#create-provisioning-key-using-p-o-s-t-1-200) | OK | OK |  | [schema](#create-provisioning-key-using-p-o-s-t-1-200-schema) |
| [201](#create-provisioning-key-using-p-o-s-t-1-201) | Created | Created |  | [schema](#create-provisioning-key-using-p-o-s-t-1-201-schema) |
| [400](#create-provisioning-key-using-p-o-s-t-1-400) | Bad Request | BadRequest |  | [schema](#create-provisioning-key-using-p-o-s-t-1-400-schema) |
| [401](#create-provisioning-key-using-p-o-s-t-1-401) | Unauthorized | Unauthorized |  | [schema](#create-provisioning-key-using-p-o-s-t-1-401-schema) |
| [403](#create-provisioning-key-using-p-o-s-t-1-403) | Forbidden | Forbidden |  | [schema](#create-provisioning-key-using-p-o-s-t-1-403-schema) |
| [404](#create-provisioning-key-using-p-o-s-t-1-404) | Not Found | Not Found |  | [schema](#create-provisioning-key-using-p-o-s-t-1-404-schema) |
| [429](#create-provisioning-key-using-p-o-s-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#create-provisioning-key-using-p-o-s-t-1-429-schema) |

#### Responses


##### <span id="create-provisioning-key-using-p-o-s-t-1-200"></span> 200 - OK
Status: OK

###### <span id="create-provisioning-key-using-p-o-s-t-1-200-schema"></span> Schema
   
  

[ProvisioningKey](#provisioning-key)

##### <span id="create-provisioning-key-using-p-o-s-t-1-201"></span> 201 - Created
Status: Created

###### <span id="create-provisioning-key-using-p-o-s-t-1-201-schema"></span> Schema

##### <span id="create-provisioning-key-using-p-o-s-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="create-provisioning-key-using-p-o-s-t-1-400-schema"></span> Schema

##### <span id="create-provisioning-key-using-p-o-s-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="create-provisioning-key-using-p-o-s-t-1-401-schema"></span> Schema

##### <span id="create-provisioning-key-using-p-o-s-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="create-provisioning-key-using-p-o-s-t-1-403-schema"></span> Schema

##### <span id="create-provisioning-key-using-p-o-s-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="create-provisioning-key-using-p-o-s-t-1-404-schema"></span> Schema

##### <span id="create-provisioning-key-using-p-o-s-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="create-provisioning-key-using-p-o-s-t-1-429-schema"></span> Schema

### <span id="delete-app-connector-group-using-d-e-l-e-t-e-1"></span> Delete the Connector Group. (*deleteAppConnectorGroupUsingDELETE_1*)

```
DELETE /mgmtconfig/v1/admin/customers/{customerId}/appConnectorGroup/{appConnectorGroupId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| appConnectorGroupId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Connector Group. |
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-app-connector-group-using-d-e-l-e-t-e-1-204) | No Content | No Content |  | [schema](#delete-app-connector-group-using-d-e-l-e-t-e-1-204-schema) |
| [400](#delete-app-connector-group-using-d-e-l-e-t-e-1-400) | Bad Request | BadRequest |  | [schema](#delete-app-connector-group-using-d-e-l-e-t-e-1-400-schema) |
| [401](#delete-app-connector-group-using-d-e-l-e-t-e-1-401) | Unauthorized | Unauthorized |  | [schema](#delete-app-connector-group-using-d-e-l-e-t-e-1-401-schema) |
| [403](#delete-app-connector-group-using-d-e-l-e-t-e-1-403) | Forbidden | Forbidden |  | [schema](#delete-app-connector-group-using-d-e-l-e-t-e-1-403-schema) |
| [429](#delete-app-connector-group-using-d-e-l-e-t-e-1-429) | Too Many Requests | TooManyRequest |  | [schema](#delete-app-connector-group-using-d-e-l-e-t-e-1-429-schema) |

#### Responses


##### <span id="delete-app-connector-group-using-d-e-l-e-t-e-1-204"></span> 204 - No Content
Status: No Content

###### <span id="delete-app-connector-group-using-d-e-l-e-t-e-1-204-schema"></span> Schema

##### <span id="delete-app-connector-group-using-d-e-l-e-t-e-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="delete-app-connector-group-using-d-e-l-e-t-e-1-400-schema"></span> Schema

##### <span id="delete-app-connector-group-using-d-e-l-e-t-e-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="delete-app-connector-group-using-d-e-l-e-t-e-1-401-schema"></span> Schema

##### <span id="delete-app-connector-group-using-d-e-l-e-t-e-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="delete-app-connector-group-using-d-e-l-e-t-e-1-403-schema"></span> Schema

##### <span id="delete-app-connector-group-using-d-e-l-e-t-e-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="delete-app-connector-group-using-d-e-l-e-t-e-1-429-schema"></span> Schema

### <span id="delete-app-server-group-using-d-e-l-e-t-e-1"></span> Delete a Server Group. (*deleteAppServerGroupUsingDELETE_1*)

```
DELETE /mgmtconfig/v1/admin/customers/{customerId}/serverGroup/{groupId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| groupId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Server Group. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-app-server-group-using-d-e-l-e-t-e-1-204) | No Content | No Content |  | [schema](#delete-app-server-group-using-d-e-l-e-t-e-1-204-schema) |
| [400](#delete-app-server-group-using-d-e-l-e-t-e-1-400) | Bad Request | BadRequest |  | [schema](#delete-app-server-group-using-d-e-l-e-t-e-1-400-schema) |
| [401](#delete-app-server-group-using-d-e-l-e-t-e-1-401) | Unauthorized | Unauthorized |  | [schema](#delete-app-server-group-using-d-e-l-e-t-e-1-401-schema) |
| [403](#delete-app-server-group-using-d-e-l-e-t-e-1-403) | Forbidden | Forbidden |  | [schema](#delete-app-server-group-using-d-e-l-e-t-e-1-403-schema) |
| [429](#delete-app-server-group-using-d-e-l-e-t-e-1-429) | Too Many Requests | TooManyRequest |  | [schema](#delete-app-server-group-using-d-e-l-e-t-e-1-429-schema) |

#### Responses


##### <span id="delete-app-server-group-using-d-e-l-e-t-e-1-204"></span> 204 - No Content
Status: No Content

###### <span id="delete-app-server-group-using-d-e-l-e-t-e-1-204-schema"></span> Schema

##### <span id="delete-app-server-group-using-d-e-l-e-t-e-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="delete-app-server-group-using-d-e-l-e-t-e-1-400-schema"></span> Schema

##### <span id="delete-app-server-group-using-d-e-l-e-t-e-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="delete-app-server-group-using-d-e-l-e-t-e-1-401-schema"></span> Schema

##### <span id="delete-app-server-group-using-d-e-l-e-t-e-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="delete-app-server-group-using-d-e-l-e-t-e-1-403-schema"></span> Schema

##### <span id="delete-app-server-group-using-d-e-l-e-t-e-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="delete-app-server-group-using-d-e-l-e-t-e-1-429-schema"></span> Schema

### <span id="delete-app-server-using-d-e-l-e-t-e-1"></span> Delete a Server. (*deleteAppServerUsingDELETE_1*)

```
DELETE /mgmtconfig/v1/admin/customers/{customerId}/server/{serverId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| serverId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Server. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-app-server-using-d-e-l-e-t-e-1-204) | No Content | No Content |  | [schema](#delete-app-server-using-d-e-l-e-t-e-1-204-schema) |
| [400](#delete-app-server-using-d-e-l-e-t-e-1-400) | Bad Request | BadRequest |  | [schema](#delete-app-server-using-d-e-l-e-t-e-1-400-schema) |
| [401](#delete-app-server-using-d-e-l-e-t-e-1-401) | Unauthorized | Unauthorized |  | [schema](#delete-app-server-using-d-e-l-e-t-e-1-401-schema) |
| [403](#delete-app-server-using-d-e-l-e-t-e-1-403) | Forbidden | Forbidden |  | [schema](#delete-app-server-using-d-e-l-e-t-e-1-403-schema) |
| [429](#delete-app-server-using-d-e-l-e-t-e-1-429) | Too Many Requests | TooManyRequest |  | [schema](#delete-app-server-using-d-e-l-e-t-e-1-429-schema) |

#### Responses


##### <span id="delete-app-server-using-d-e-l-e-t-e-1-204"></span> 204 - No Content
Status: No Content

###### <span id="delete-app-server-using-d-e-l-e-t-e-1-204-schema"></span> Schema

##### <span id="delete-app-server-using-d-e-l-e-t-e-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="delete-app-server-using-d-e-l-e-t-e-1-400-schema"></span> Schema

##### <span id="delete-app-server-using-d-e-l-e-t-e-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="delete-app-server-using-d-e-l-e-t-e-1-401-schema"></span> Schema

##### <span id="delete-app-server-using-d-e-l-e-t-e-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="delete-app-server-using-d-e-l-e-t-e-1-403-schema"></span> Schema

##### <span id="delete-app-server-using-d-e-l-e-t-e-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="delete-app-server-using-d-e-l-e-t-e-1-429-schema"></span> Schema

### <span id="delete-application-using-d-e-l-e-t-e-1"></span> Delete an Application Segment. (*deleteApplicationUsingDELETE_1*)

```
DELETE /mgmtconfig/v1/admin/customers/{customerId}/application/{applicationId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| applicationId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Application Segment. |
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| forceDelete | `query` | boolean | `bool` |  |  |  | Setting this field to true will delete the mapping between Application Segment and Segment Group |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-application-using-d-e-l-e-t-e-1-204) | No Content | No Content |  | [schema](#delete-application-using-d-e-l-e-t-e-1-204-schema) |
| [400](#delete-application-using-d-e-l-e-t-e-1-400) | Bad Request | BadRequest |  | [schema](#delete-application-using-d-e-l-e-t-e-1-400-schema) |
| [401](#delete-application-using-d-e-l-e-t-e-1-401) | Unauthorized | Unauthorized |  | [schema](#delete-application-using-d-e-l-e-t-e-1-401-schema) |
| [403](#delete-application-using-d-e-l-e-t-e-1-403) | Forbidden | Forbidden |  | [schema](#delete-application-using-d-e-l-e-t-e-1-403-schema) |
| [429](#delete-application-using-d-e-l-e-t-e-1-429) | Too Many Requests | TooManyRequest |  | [schema](#delete-application-using-d-e-l-e-t-e-1-429-schema) |

#### Responses


##### <span id="delete-application-using-d-e-l-e-t-e-1-204"></span> 204 - No Content
Status: No Content

###### <span id="delete-application-using-d-e-l-e-t-e-1-204-schema"></span> Schema

##### <span id="delete-application-using-d-e-l-e-t-e-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="delete-application-using-d-e-l-e-t-e-1-400-schema"></span> Schema

##### <span id="delete-application-using-d-e-l-e-t-e-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="delete-application-using-d-e-l-e-t-e-1-401-schema"></span> Schema

##### <span id="delete-application-using-d-e-l-e-t-e-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="delete-application-using-d-e-l-e-t-e-1-403-schema"></span> Schema

##### <span id="delete-application-using-d-e-l-e-t-e-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="delete-application-using-d-e-l-e-t-e-1-429-schema"></span> Schema

### <span id="delete-connector-using-d-e-l-e-t-e-1"></span> Delete the Connector. (*deleteConnectorUsingDELETE_1*)

```
DELETE /mgmtconfig/v1/admin/customers/{customerId}/connector/{connectorId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| connectorId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Connector. |
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-connector-using-d-e-l-e-t-e-1-204) | No Content | No Content |  | [schema](#delete-connector-using-d-e-l-e-t-e-1-204-schema) |
| [400](#delete-connector-using-d-e-l-e-t-e-1-400) | Bad Request | BadRequest |  | [schema](#delete-connector-using-d-e-l-e-t-e-1-400-schema) |
| [401](#delete-connector-using-d-e-l-e-t-e-1-401) | Unauthorized | Unauthorized |  | [schema](#delete-connector-using-d-e-l-e-t-e-1-401-schema) |
| [403](#delete-connector-using-d-e-l-e-t-e-1-403) | Forbidden | Forbidden |  | [schema](#delete-connector-using-d-e-l-e-t-e-1-403-schema) |
| [429](#delete-connector-using-d-e-l-e-t-e-1-429) | Too Many Requests | TooManyRequest |  | [schema](#delete-connector-using-d-e-l-e-t-e-1-429-schema) |

#### Responses


##### <span id="delete-connector-using-d-e-l-e-t-e-1-204"></span> 204 - No Content
Status: No Content

###### <span id="delete-connector-using-d-e-l-e-t-e-1-204-schema"></span> Schema

##### <span id="delete-connector-using-d-e-l-e-t-e-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="delete-connector-using-d-e-l-e-t-e-1-400-schema"></span> Schema

##### <span id="delete-connector-using-d-e-l-e-t-e-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="delete-connector-using-d-e-l-e-t-e-1-401-schema"></span> Schema

##### <span id="delete-connector-using-d-e-l-e-t-e-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="delete-connector-using-d-e-l-e-t-e-1-403-schema"></span> Schema

##### <span id="delete-connector-using-d-e-l-e-t-e-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="delete-connector-using-d-e-l-e-t-e-1-429-schema"></span> Schema

### <span id="delete-lss-using-d-e-l-e-t-e-3"></span> Delete a LSS attribute. (*deleteLssUsingDELETE_3*)

```
DELETE /mgmtconfig/v2/admin/customers/{customerId}/lssConfig/{lssId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| lssId | `path` | string | `string` |  | ✓ |  | The unique identifier of the LSS resource. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-lss-using-d-e-l-e-t-e-3-204) | No Content | No Content |  | [schema](#delete-lss-using-d-e-l-e-t-e-3-204-schema) |
| [400](#delete-lss-using-d-e-l-e-t-e-3-400) | Bad Request | BadRequest |  | [schema](#delete-lss-using-d-e-l-e-t-e-3-400-schema) |
| [401](#delete-lss-using-d-e-l-e-t-e-3-401) | Unauthorized | Unauthorized |  | [schema](#delete-lss-using-d-e-l-e-t-e-3-401-schema) |
| [403](#delete-lss-using-d-e-l-e-t-e-3-403) | Forbidden | Forbidden |  | [schema](#delete-lss-using-d-e-l-e-t-e-3-403-schema) |
| [429](#delete-lss-using-d-e-l-e-t-e-3-429) | Too Many Requests | TooManyRequest |  | [schema](#delete-lss-using-d-e-l-e-t-e-3-429-schema) |

#### Responses


##### <span id="delete-lss-using-d-e-l-e-t-e-3-204"></span> 204 - No Content
Status: No Content

###### <span id="delete-lss-using-d-e-l-e-t-e-3-204-schema"></span> Schema

##### <span id="delete-lss-using-d-e-l-e-t-e-3-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="delete-lss-using-d-e-l-e-t-e-3-400-schema"></span> Schema

##### <span id="delete-lss-using-d-e-l-e-t-e-3-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="delete-lss-using-d-e-l-e-t-e-3-401-schema"></span> Schema

##### <span id="delete-lss-using-d-e-l-e-t-e-3-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="delete-lss-using-d-e-l-e-t-e-3-403-schema"></span> Schema

##### <span id="delete-lss-using-d-e-l-e-t-e-3-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="delete-lss-using-d-e-l-e-t-e-3-429-schema"></span> Schema

### <span id="delete-provisioning-key-using-d-e-l-e-t-e-1"></span> Delete the Provisioning Key. (*deleteProvisioningKeyUsingDELETE_1*)

```
DELETE /mgmtconfig/v1/admin/customers/{customerId}/associationType/{associationType}/provisioningKey/{provisioningKeyId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| associationType | `path` | string | `string` |  | ✓ |  | Specifies the provisioning key type for App Connectors or ZPA Private Service Edges. The supported values are CONNECTOR_GRP and SERVICE_EDGE_GRP. |
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| provisioningKeyId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Provisioning Key. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-provisioning-key-using-d-e-l-e-t-e-1-204) | No Content | No Content |  | [schema](#delete-provisioning-key-using-d-e-l-e-t-e-1-204-schema) |
| [400](#delete-provisioning-key-using-d-e-l-e-t-e-1-400) | Bad Request | BadRequest |  | [schema](#delete-provisioning-key-using-d-e-l-e-t-e-1-400-schema) |
| [401](#delete-provisioning-key-using-d-e-l-e-t-e-1-401) | Unauthorized | Unauthorized |  | [schema](#delete-provisioning-key-using-d-e-l-e-t-e-1-401-schema) |
| [403](#delete-provisioning-key-using-d-e-l-e-t-e-1-403) | Forbidden | Forbidden |  | [schema](#delete-provisioning-key-using-d-e-l-e-t-e-1-403-schema) |
| [429](#delete-provisioning-key-using-d-e-l-e-t-e-1-429) | Too Many Requests | TooManyRequest |  | [schema](#delete-provisioning-key-using-d-e-l-e-t-e-1-429-schema) |

#### Responses


##### <span id="delete-provisioning-key-using-d-e-l-e-t-e-1-204"></span> 204 - No Content
Status: No Content

###### <span id="delete-provisioning-key-using-d-e-l-e-t-e-1-204-schema"></span> Schema

##### <span id="delete-provisioning-key-using-d-e-l-e-t-e-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="delete-provisioning-key-using-d-e-l-e-t-e-1-400-schema"></span> Schema

##### <span id="delete-provisioning-key-using-d-e-l-e-t-e-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="delete-provisioning-key-using-d-e-l-e-t-e-1-401-schema"></span> Schema

##### <span id="delete-provisioning-key-using-d-e-l-e-t-e-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="delete-provisioning-key-using-d-e-l-e-t-e-1-403-schema"></span> Schema

##### <span id="delete-provisioning-key-using-d-e-l-e-t-e-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="delete-provisioning-key-using-d-e-l-e-t-e-1-429-schema"></span> Schema

### <span id="delete-rule-in-policy-set-using-d-e-l-e-t-e-1"></span> Delete a rule in a policy. (*deleteRuleInPolicySetUsingDELETE_1*)

```
DELETE /mgmtconfig/v1/admin/customers/{customerId}/policySet/{policySetId}/rule/{ruleId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| policySetId | `path` | string | `string` |  | ✓ |  | The unique identifier of the policy. |
| ruleId | `path` | string | `string` |  | ✓ |  | The unique identifier of a rule in a policy. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-rule-in-policy-set-using-d-e-l-e-t-e-1-204) | No Content | No Content |  | [schema](#delete-rule-in-policy-set-using-d-e-l-e-t-e-1-204-schema) |
| [400](#delete-rule-in-policy-set-using-d-e-l-e-t-e-1-400) | Bad Request | BadRequest |  | [schema](#delete-rule-in-policy-set-using-d-e-l-e-t-e-1-400-schema) |
| [401](#delete-rule-in-policy-set-using-d-e-l-e-t-e-1-401) | Unauthorized | Unauthorized |  | [schema](#delete-rule-in-policy-set-using-d-e-l-e-t-e-1-401-schema) |
| [403](#delete-rule-in-policy-set-using-d-e-l-e-t-e-1-403) | Forbidden | Forbidden |  | [schema](#delete-rule-in-policy-set-using-d-e-l-e-t-e-1-403-schema) |
| [429](#delete-rule-in-policy-set-using-d-e-l-e-t-e-1-429) | Too Many Requests | TooManyRequest |  | [schema](#delete-rule-in-policy-set-using-d-e-l-e-t-e-1-429-schema) |

#### Responses


##### <span id="delete-rule-in-policy-set-using-d-e-l-e-t-e-1-204"></span> 204 - No Content
Status: No Content

###### <span id="delete-rule-in-policy-set-using-d-e-l-e-t-e-1-204-schema"></span> Schema

##### <span id="delete-rule-in-policy-set-using-d-e-l-e-t-e-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="delete-rule-in-policy-set-using-d-e-l-e-t-e-1-400-schema"></span> Schema

##### <span id="delete-rule-in-policy-set-using-d-e-l-e-t-e-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="delete-rule-in-policy-set-using-d-e-l-e-t-e-1-401-schema"></span> Schema

##### <span id="delete-rule-in-policy-set-using-d-e-l-e-t-e-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="delete-rule-in-policy-set-using-d-e-l-e-t-e-1-403-schema"></span> Schema

##### <span id="delete-rule-in-policy-set-using-d-e-l-e-t-e-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="delete-rule-in-policy-set-using-d-e-l-e-t-e-1-429-schema"></span> Schema

### <span id="delete-segment-group-using-d-e-l-e-t-e-1"></span> Delete a Segment Group. (*deleteSegmentGroupUsingDELETE_1*)

```
DELETE /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup/{segmentGroupId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| segmentGroupId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Segment Group. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-segment-group-using-d-e-l-e-t-e-1-204) | No Content | No Content |  | [schema](#delete-segment-group-using-d-e-l-e-t-e-1-204-schema) |
| [400](#delete-segment-group-using-d-e-l-e-t-e-1-400) | Bad Request | BadRequest |  | [schema](#delete-segment-group-using-d-e-l-e-t-e-1-400-schema) |
| [401](#delete-segment-group-using-d-e-l-e-t-e-1-401) | Unauthorized | Unauthorized |  | [schema](#delete-segment-group-using-d-e-l-e-t-e-1-401-schema) |
| [403](#delete-segment-group-using-d-e-l-e-t-e-1-403) | Forbidden | Forbidden |  | [schema](#delete-segment-group-using-d-e-l-e-t-e-1-403-schema) |
| [429](#delete-segment-group-using-d-e-l-e-t-e-1-429) | Too Many Requests | TooManyRequest |  | [schema](#delete-segment-group-using-d-e-l-e-t-e-1-429-schema) |

#### Responses


##### <span id="delete-segment-group-using-d-e-l-e-t-e-1-204"></span> 204 - No Content
Status: No Content

###### <span id="delete-segment-group-using-d-e-l-e-t-e-1-204-schema"></span> Schema

##### <span id="delete-segment-group-using-d-e-l-e-t-e-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="delete-segment-group-using-d-e-l-e-t-e-1-400-schema"></span> Schema

##### <span id="delete-segment-group-using-d-e-l-e-t-e-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="delete-segment-group-using-d-e-l-e-t-e-1-401-schema"></span> Schema

##### <span id="delete-segment-group-using-d-e-l-e-t-e-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="delete-segment-group-using-d-e-l-e-t-e-1-403-schema"></span> Schema

##### <span id="delete-segment-group-using-d-e-l-e-t-e-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="delete-segment-group-using-d-e-l-e-t-e-1-429-schema"></span> Schema

### <span id="delete-service-edge-group-using-d-e-l-e-t-e"></span> Delete the Service Edge Group. (*deleteServiceEdgeGroupUsingDELETE*)

```
DELETE /mgmtconfig/v1/admin/customers/{customerId}/serviceEdgeGroup/{serviceEdgeGroupId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| serviceEdgeGroupId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Service Edge Group. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-service-edge-group-using-d-e-l-e-t-e-204) | No Content | No Content |  | [schema](#delete-service-edge-group-using-d-e-l-e-t-e-204-schema) |
| [400](#delete-service-edge-group-using-d-e-l-e-t-e-400) | Bad Request | BadRequest |  | [schema](#delete-service-edge-group-using-d-e-l-e-t-e-400-schema) |
| [401](#delete-service-edge-group-using-d-e-l-e-t-e-401) | Unauthorized | Unauthorized |  | [schema](#delete-service-edge-group-using-d-e-l-e-t-e-401-schema) |
| [403](#delete-service-edge-group-using-d-e-l-e-t-e-403) | Forbidden | Forbidden |  | [schema](#delete-service-edge-group-using-d-e-l-e-t-e-403-schema) |
| [429](#delete-service-edge-group-using-d-e-l-e-t-e-429) | Too Many Requests | TooManyRequest |  | [schema](#delete-service-edge-group-using-d-e-l-e-t-e-429-schema) |

#### Responses


##### <span id="delete-service-edge-group-using-d-e-l-e-t-e-204"></span> 204 - No Content
Status: No Content

###### <span id="delete-service-edge-group-using-d-e-l-e-t-e-204-schema"></span> Schema

##### <span id="delete-service-edge-group-using-d-e-l-e-t-e-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="delete-service-edge-group-using-d-e-l-e-t-e-400-schema"></span> Schema

##### <span id="delete-service-edge-group-using-d-e-l-e-t-e-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="delete-service-edge-group-using-d-e-l-e-t-e-401-schema"></span> Schema

##### <span id="delete-service-edge-group-using-d-e-l-e-t-e-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="delete-service-edge-group-using-d-e-l-e-t-e-403-schema"></span> Schema

##### <span id="delete-service-edge-group-using-d-e-l-e-t-e-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="delete-service-edge-group-using-d-e-l-e-t-e-429-schema"></span> Schema

### <span id="delete-service-edge-using-d-e-l-e-t-e"></span> Delete the Service Edge. (*deleteServiceEdgeUsingDELETE*)

```
DELETE /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/{serviceEdgeId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| serviceEdgeId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Service Edge. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-service-edge-using-d-e-l-e-t-e-204) | No Content | No Content |  | [schema](#delete-service-edge-using-d-e-l-e-t-e-204-schema) |
| [400](#delete-service-edge-using-d-e-l-e-t-e-400) | Bad Request | BadRequest |  | [schema](#delete-service-edge-using-d-e-l-e-t-e-400-schema) |
| [401](#delete-service-edge-using-d-e-l-e-t-e-401) | Unauthorized | Unauthorized |  | [schema](#delete-service-edge-using-d-e-l-e-t-e-401-schema) |
| [403](#delete-service-edge-using-d-e-l-e-t-e-403) | Forbidden | Forbidden |  | [schema](#delete-service-edge-using-d-e-l-e-t-e-403-schema) |
| [429](#delete-service-edge-using-d-e-l-e-t-e-429) | Too Many Requests | TooManyRequest |  | [schema](#delete-service-edge-using-d-e-l-e-t-e-429-schema) |

#### Responses


##### <span id="delete-service-edge-using-d-e-l-e-t-e-204"></span> 204 - No Content
Status: No Content

###### <span id="delete-service-edge-using-d-e-l-e-t-e-204-schema"></span> Schema

##### <span id="delete-service-edge-using-d-e-l-e-t-e-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="delete-service-edge-using-d-e-l-e-t-e-400-schema"></span> Schema

##### <span id="delete-service-edge-using-d-e-l-e-t-e-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="delete-service-edge-using-d-e-l-e-t-e-401-schema"></span> Schema

##### <span id="delete-service-edge-using-d-e-l-e-t-e-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="delete-service-edge-using-d-e-l-e-t-e-403-schema"></span> Schema

##### <span id="delete-service-edge-using-d-e-l-e-t-e-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="delete-service-edge-using-d-e-l-e-t-e-429-schema"></span> Schema

### <span id="get-all-app-servers-using-g-e-t-1"></span> Get all the configured Servers. (*getAllAppServersUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/server
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-app-servers-using-g-e-t-1-200) | OK | OK |  | [schema](#get-all-app-servers-using-g-e-t-1-200-schema) |
| [400](#get-all-app-servers-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-all-app-servers-using-g-e-t-1-400-schema) |
| [401](#get-all-app-servers-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-all-app-servers-using-g-e-t-1-401-schema) |
| [403](#get-all-app-servers-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-all-app-servers-using-g-e-t-1-403-schema) |
| [404](#get-all-app-servers-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-all-app-servers-using-g-e-t-1-404-schema) |
| [429](#get-all-app-servers-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-all-app-servers-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-all-app-servers-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-all-app-servers-using-g-e-t-1-200-schema"></span> Schema
   
  

[PageListOfApplicationServer](#page-list-of-application-server)

##### <span id="get-all-app-servers-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-all-app-servers-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-all-app-servers-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-all-app-servers-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-all-app-servers-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-all-app-servers-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-all-app-servers-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-all-app-servers-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-all-app-servers-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-all-app-servers-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-all-applications-using-g-e-t-3"></span> Get all configured Application Segments. (*getAllApplicationsUsingGET_3*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/application
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-applications-using-g-e-t-3-200) | OK | OK |  | [schema](#get-all-applications-using-g-e-t-3-200-schema) |
| [400](#get-all-applications-using-g-e-t-3-400) | Bad Request | BadRequest |  | [schema](#get-all-applications-using-g-e-t-3-400-schema) |
| [401](#get-all-applications-using-g-e-t-3-401) | Unauthorized | Unauthorized |  | [schema](#get-all-applications-using-g-e-t-3-401-schema) |
| [403](#get-all-applications-using-g-e-t-3-403) | Forbidden | Forbidden |  | [schema](#get-all-applications-using-g-e-t-3-403-schema) |
| [404](#get-all-applications-using-g-e-t-3-404) | Not Found | Not Found |  | [schema](#get-all-applications-using-g-e-t-3-404-schema) |
| [429](#get-all-applications-using-g-e-t-3-429) | Too Many Requests | TooManyRequest |  | [schema](#get-all-applications-using-g-e-t-3-429-schema) |

#### Responses


##### <span id="get-all-applications-using-g-e-t-3-200"></span> 200 - OK
Status: OK

###### <span id="get-all-applications-using-g-e-t-3-200-schema"></span> Schema
   
  

[PageListOfApplicationResource](#page-list-of-application-resource)

##### <span id="get-all-applications-using-g-e-t-3-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-all-applications-using-g-e-t-3-400-schema"></span> Schema

##### <span id="get-all-applications-using-g-e-t-3-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-all-applications-using-g-e-t-3-401-schema"></span> Schema

##### <span id="get-all-applications-using-g-e-t-3-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-all-applications-using-g-e-t-3-403-schema"></span> Schema

##### <span id="get-all-applications-using-g-e-t-3-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-all-applications-using-g-e-t-3-404-schema"></span> Schema

##### <span id="get-all-applications-using-g-e-t-3-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-all-applications-using-g-e-t-3-429-schema"></span> Schema

### <span id="get-all-attributes-by-idp-id-and-page-using-g-e-t"></span> Get all attributes configured for a given customer (*getAllAttributesByIdpIdAndPageUsingGET*)

```
GET /mgmtconfig/v2/admin/customers/{customerId}/samlAttribute/idp/{idpId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| idpId | `path` | string | `string` |  | ✓ |  | The unique identifier of the IdP. |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-attributes-by-idp-id-and-page-using-g-e-t-200) | OK | OK |  | [schema](#get-all-attributes-by-idp-id-and-page-using-g-e-t-200-schema) |
| [400](#get-all-attributes-by-idp-id-and-page-using-g-e-t-400) | Bad Request | BadRequest |  | [schema](#get-all-attributes-by-idp-id-and-page-using-g-e-t-400-schema) |
| [401](#get-all-attributes-by-idp-id-and-page-using-g-e-t-401) | Unauthorized | Unauthorized |  | [schema](#get-all-attributes-by-idp-id-and-page-using-g-e-t-401-schema) |
| [403](#get-all-attributes-by-idp-id-and-page-using-g-e-t-403) | Forbidden | Forbidden |  | [schema](#get-all-attributes-by-idp-id-and-page-using-g-e-t-403-schema) |
| [404](#get-all-attributes-by-idp-id-and-page-using-g-e-t-404) | Not Found | Not Found |  | [schema](#get-all-attributes-by-idp-id-and-page-using-g-e-t-404-schema) |
| [429](#get-all-attributes-by-idp-id-and-page-using-g-e-t-429) | Too Many Requests | TooManyRequest |  | [schema](#get-all-attributes-by-idp-id-and-page-using-g-e-t-429-schema) |

#### Responses


##### <span id="get-all-attributes-by-idp-id-and-page-using-g-e-t-200"></span> 200 - OK
Status: OK

###### <span id="get-all-attributes-by-idp-id-and-page-using-g-e-t-200-schema"></span> Schema
   
  

[PageListOfSamlAttribute](#page-list-of-saml-attribute)

##### <span id="get-all-attributes-by-idp-id-and-page-using-g-e-t-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-all-attributes-by-idp-id-and-page-using-g-e-t-400-schema"></span> Schema

##### <span id="get-all-attributes-by-idp-id-and-page-using-g-e-t-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-all-attributes-by-idp-id-and-page-using-g-e-t-401-schema"></span> Schema

##### <span id="get-all-attributes-by-idp-id-and-page-using-g-e-t-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-all-attributes-by-idp-id-and-page-using-g-e-t-403-schema"></span> Schema

##### <span id="get-all-attributes-by-idp-id-and-page-using-g-e-t-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-all-attributes-by-idp-id-and-page-using-g-e-t-404-schema"></span> Schema

##### <span id="get-all-attributes-by-idp-id-and-page-using-g-e-t-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-all-attributes-by-idp-id-and-page-using-g-e-t-429-schema"></span> Schema

### <span id="get-all-attributes-by-page-using-g-e-t"></span> Get all SAML attributes by page. (*getAllAttributesByPageUsingGET*)

```
GET /mgmtconfig/v2/admin/customers/{customerId}/samlAttribute
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-attributes-by-page-using-g-e-t-200) | OK | OK |  | [schema](#get-all-attributes-by-page-using-g-e-t-200-schema) |
| [400](#get-all-attributes-by-page-using-g-e-t-400) | Bad Request | BadRequest |  | [schema](#get-all-attributes-by-page-using-g-e-t-400-schema) |
| [401](#get-all-attributes-by-page-using-g-e-t-401) | Unauthorized | Unauthorized |  | [schema](#get-all-attributes-by-page-using-g-e-t-401-schema) |
| [403](#get-all-attributes-by-page-using-g-e-t-403) | Forbidden | Forbidden |  | [schema](#get-all-attributes-by-page-using-g-e-t-403-schema) |
| [404](#get-all-attributes-by-page-using-g-e-t-404) | Not Found | Not Found |  | [schema](#get-all-attributes-by-page-using-g-e-t-404-schema) |
| [429](#get-all-attributes-by-page-using-g-e-t-429) | Too Many Requests | TooManyRequest |  | [schema](#get-all-attributes-by-page-using-g-e-t-429-schema) |

#### Responses


##### <span id="get-all-attributes-by-page-using-g-e-t-200"></span> 200 - OK
Status: OK

###### <span id="get-all-attributes-by-page-using-g-e-t-200-schema"></span> Schema
   
  

[PageListOfSamlAttribute](#page-list-of-saml-attribute)

##### <span id="get-all-attributes-by-page-using-g-e-t-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-all-attributes-by-page-using-g-e-t-400-schema"></span> Schema

##### <span id="get-all-attributes-by-page-using-g-e-t-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-all-attributes-by-page-using-g-e-t-401-schema"></span> Schema

##### <span id="get-all-attributes-by-page-using-g-e-t-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-all-attributes-by-page-using-g-e-t-403-schema"></span> Schema

##### <span id="get-all-attributes-by-page-using-g-e-t-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-all-attributes-by-page-using-g-e-t-404-schema"></span> Schema

##### <span id="get-all-attributes-by-page-using-g-e-t-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-all-attributes-by-page-using-g-e-t-429-schema"></span> Schema

### <span id="get-all-attributes-using-g-e-t-1"></span> Get all posture profiles. This API will be deprecated in a future release. (*getAllAttributesUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/posture
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-attributes-using-g-e-t-1-200) | OK | OK |  | [schema](#get-all-attributes-using-g-e-t-1-200-schema) |
| [400](#get-all-attributes-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-all-attributes-using-g-e-t-1-400-schema) |
| [401](#get-all-attributes-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-all-attributes-using-g-e-t-1-401-schema) |
| [403](#get-all-attributes-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-all-attributes-using-g-e-t-1-403-schema) |
| [404](#get-all-attributes-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-all-attributes-using-g-e-t-1-404-schema) |
| [429](#get-all-attributes-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-all-attributes-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-all-attributes-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-all-attributes-using-g-e-t-1-200-schema"></span> Schema
   
  

[][PostureProfile](#posture-profile)

##### <span id="get-all-attributes-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-all-attributes-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-all-attributes-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-all-attributes-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-all-attributes-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-all-attributes-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-all-attributes-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-all-attributes-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-all-attributes-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-all-attributes-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-all-attributes-using-g-e-t-2"></span> Get all posture profiles. (*getAllAttributesUsingGET_2*)

```
GET /mgmtconfig/v2/admin/customers/{customerId}/posture
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-attributes-using-g-e-t-2-200) | OK | OK |  | [schema](#get-all-attributes-using-g-e-t-2-200-schema) |
| [400](#get-all-attributes-using-g-e-t-2-400) | Bad Request | BadRequest |  | [schema](#get-all-attributes-using-g-e-t-2-400-schema) |
| [401](#get-all-attributes-using-g-e-t-2-401) | Unauthorized | Unauthorized |  | [schema](#get-all-attributes-using-g-e-t-2-401-schema) |
| [403](#get-all-attributes-using-g-e-t-2-403) | Forbidden | Forbidden |  | [schema](#get-all-attributes-using-g-e-t-2-403-schema) |
| [404](#get-all-attributes-using-g-e-t-2-404) | Not Found | Not Found |  | [schema](#get-all-attributes-using-g-e-t-2-404-schema) |
| [429](#get-all-attributes-using-g-e-t-2-429) | Too Many Requests | TooManyRequest |  | [schema](#get-all-attributes-using-g-e-t-2-429-schema) |

#### Responses


##### <span id="get-all-attributes-using-g-e-t-2-200"></span> 200 - OK
Status: OK

###### <span id="get-all-attributes-using-g-e-t-2-200-schema"></span> Schema
   
  

[PageListOfPostureProfile](#page-list-of-posture-profile)

##### <span id="get-all-attributes-using-g-e-t-2-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-all-attributes-using-g-e-t-2-400-schema"></span> Schema

##### <span id="get-all-attributes-using-g-e-t-2-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-all-attributes-using-g-e-t-2-401-schema"></span> Schema

##### <span id="get-all-attributes-using-g-e-t-2-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-all-attributes-using-g-e-t-2-403-schema"></span> Schema

##### <span id="get-all-attributes-using-g-e-t-2-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-all-attributes-using-g-e-t-2-404-schema"></span> Schema

##### <span id="get-all-attributes-using-g-e-t-2-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-all-attributes-using-g-e-t-2-429-schema"></span> Schema

### <span id="get-all-attributes-using-g-e-t-4"></span> Get all SAML attributes. This API will be deprecated in a future release. (*getAllAttributesUsingGET_4*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/samlAttribute
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-attributes-using-g-e-t-4-200) | OK | OK |  | [schema](#get-all-attributes-using-g-e-t-4-200-schema) |
| [400](#get-all-attributes-using-g-e-t-4-400) | Bad Request | BadRequest |  | [schema](#get-all-attributes-using-g-e-t-4-400-schema) |
| [401](#get-all-attributes-using-g-e-t-4-401) | Unauthorized | Unauthorized |  | [schema](#get-all-attributes-using-g-e-t-4-401-schema) |
| [403](#get-all-attributes-using-g-e-t-4-403) | Forbidden | Forbidden |  | [schema](#get-all-attributes-using-g-e-t-4-403-schema) |
| [404](#get-all-attributes-using-g-e-t-4-404) | Not Found | Not Found |  | [schema](#get-all-attributes-using-g-e-t-4-404-schema) |
| [429](#get-all-attributes-using-g-e-t-4-429) | Too Many Requests | TooManyRequest |  | [schema](#get-all-attributes-using-g-e-t-4-429-schema) |

#### Responses


##### <span id="get-all-attributes-using-g-e-t-4-200"></span> 200 - OK
Status: OK

###### <span id="get-all-attributes-using-g-e-t-4-200-schema"></span> Schema
   
  

[][SamlAttribute](#saml-attribute)

##### <span id="get-all-attributes-using-g-e-t-4-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-all-attributes-using-g-e-t-4-400-schema"></span> Schema

##### <span id="get-all-attributes-using-g-e-t-4-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-all-attributes-using-g-e-t-4-401-schema"></span> Schema

##### <span id="get-all-attributes-using-g-e-t-4-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-all-attributes-using-g-e-t-4-403-schema"></span> Schema

##### <span id="get-all-attributes-using-g-e-t-4-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-all-attributes-using-g-e-t-4-404-schema"></span> Schema

##### <span id="get-all-attributes-using-g-e-t-4-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-all-attributes-using-g-e-t-4-429-schema"></span> Schema

### <span id="get-all-connectors-using-g-e-t-1"></span> Get all the configured Connector details. (*getAllConnectorsUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/connector
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-connectors-using-g-e-t-1-200) | OK | OK |  | [schema](#get-all-connectors-using-g-e-t-1-200-schema) |
| [400](#get-all-connectors-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-all-connectors-using-g-e-t-1-400-schema) |
| [401](#get-all-connectors-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-all-connectors-using-g-e-t-1-401-schema) |
| [403](#get-all-connectors-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-all-connectors-using-g-e-t-1-403-schema) |
| [404](#get-all-connectors-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-all-connectors-using-g-e-t-1-404-schema) |
| [429](#get-all-connectors-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-all-connectors-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-all-connectors-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-all-connectors-using-g-e-t-1-200-schema"></span> Schema
   
  

[PageListOfConnector](#page-list-of-connector)

##### <span id="get-all-connectors-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-all-connectors-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-all-connectors-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-all-connectors-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-all-connectors-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-all-connectors-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-all-connectors-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-all-connectors-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-all-connectors-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-all-connectors-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-all-idp-using-g-e-t"></span> Get the configured IdP details. (*getAllIdpUsingGET*)

```
GET /mgmtconfig/v2/admin/customers/{customerId}/idp
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| scimEnabled | `query` | boolean | `bool` |  |  |  | Returns all SCIM IdPs if set to true. Returns all non SCIM IdPs if set to false. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-idp-using-g-e-t-200) | OK | OK |  | [schema](#get-all-idp-using-g-e-t-200-schema) |
| [400](#get-all-idp-using-g-e-t-400) | Bad Request | BadRequest |  | [schema](#get-all-idp-using-g-e-t-400-schema) |
| [401](#get-all-idp-using-g-e-t-401) | Unauthorized | Unauthorized |  | [schema](#get-all-idp-using-g-e-t-401-schema) |
| [403](#get-all-idp-using-g-e-t-403) | Forbidden | Forbidden |  | [schema](#get-all-idp-using-g-e-t-403-schema) |
| [404](#get-all-idp-using-g-e-t-404) | Not Found | Not Found |  | [schema](#get-all-idp-using-g-e-t-404-schema) |
| [429](#get-all-idp-using-g-e-t-429) | Too Many Requests | TooManyRequest |  | [schema](#get-all-idp-using-g-e-t-429-schema) |

#### Responses


##### <span id="get-all-idp-using-g-e-t-200"></span> 200 - OK
Status: OK

###### <span id="get-all-idp-using-g-e-t-200-schema"></span> Schema
   
  

[PageListOfIdp](#page-list-of-idp)

##### <span id="get-all-idp-using-g-e-t-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-all-idp-using-g-e-t-400-schema"></span> Schema

##### <span id="get-all-idp-using-g-e-t-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-all-idp-using-g-e-t-401-schema"></span> Schema

##### <span id="get-all-idp-using-g-e-t-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-all-idp-using-g-e-t-403-schema"></span> Schema

##### <span id="get-all-idp-using-g-e-t-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-all-idp-using-g-e-t-404-schema"></span> Schema

##### <span id="get-all-idp-using-g-e-t-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-all-idp-using-g-e-t-429-schema"></span> Schema

### <span id="get-all-issued-certs-using-g-e-t-1"></span> Get all the issued certificates. This API will be deprecated in a future release. (*getAllIssuedCertsUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/clientlessCertificate/issued
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-issued-certs-using-g-e-t-1-200) | OK | OK |  | [schema](#get-all-issued-certs-using-g-e-t-1-200-schema) |
| [400](#get-all-issued-certs-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-all-issued-certs-using-g-e-t-1-400-schema) |
| [401](#get-all-issued-certs-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-all-issued-certs-using-g-e-t-1-401-schema) |
| [403](#get-all-issued-certs-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-all-issued-certs-using-g-e-t-1-403-schema) |
| [404](#get-all-issued-certs-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-all-issued-certs-using-g-e-t-1-404-schema) |
| [429](#get-all-issued-certs-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-all-issued-certs-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-all-issued-certs-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-all-issued-certs-using-g-e-t-1-200-schema"></span> Schema
   
  

[][BACertificate](#b-a-certificate)

##### <span id="get-all-issued-certs-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-all-issued-certs-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-all-issued-certs-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-all-issued-certs-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-all-issued-certs-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-all-issued-certs-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-all-issued-certs-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-all-issued-certs-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-all-issued-certs-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-all-issued-certs-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-all-issued-certs-using-g-e-t-2"></span> Get all the issued certificates. (*getAllIssuedCertsUsingGET_2*)

```
GET /mgmtconfig/v2/admin/customers/{customerId}/clientlessCertificate/issued
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-issued-certs-using-g-e-t-2-200) | OK | OK |  | [schema](#get-all-issued-certs-using-g-e-t-2-200-schema) |
| [400](#get-all-issued-certs-using-g-e-t-2-400) | Bad Request | BadRequest |  | [schema](#get-all-issued-certs-using-g-e-t-2-400-schema) |
| [401](#get-all-issued-certs-using-g-e-t-2-401) | Unauthorized | Unauthorized |  | [schema](#get-all-issued-certs-using-g-e-t-2-401-schema) |
| [403](#get-all-issued-certs-using-g-e-t-2-403) | Forbidden | Forbidden |  | [schema](#get-all-issued-certs-using-g-e-t-2-403-schema) |
| [404](#get-all-issued-certs-using-g-e-t-2-404) | Not Found | Not Found |  | [schema](#get-all-issued-certs-using-g-e-t-2-404-schema) |
| [429](#get-all-issued-certs-using-g-e-t-2-429) | Too Many Requests | TooManyRequest |  | [schema](#get-all-issued-certs-using-g-e-t-2-429-schema) |

#### Responses


##### <span id="get-all-issued-certs-using-g-e-t-2-200"></span> 200 - OK
Status: OK

###### <span id="get-all-issued-certs-using-g-e-t-2-200-schema"></span> Schema
   
  

[PageListOfBACertificate](#page-list-of-b-a-certificate)

##### <span id="get-all-issued-certs-using-g-e-t-2-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-all-issued-certs-using-g-e-t-2-400-schema"></span> Schema

##### <span id="get-all-issued-certs-using-g-e-t-2-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-all-issued-certs-using-g-e-t-2-401-schema"></span> Schema

##### <span id="get-all-issued-certs-using-g-e-t-2-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-all-issued-certs-using-g-e-t-2-403-schema"></span> Schema

##### <span id="get-all-issued-certs-using-g-e-t-2-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-all-issued-certs-using-g-e-t-2-404-schema"></span> Schema

##### <span id="get-all-issued-certs-using-g-e-t-2-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-all-issued-certs-using-g-e-t-2-429-schema"></span> Schema

### <span id="get-all-s-c-i-m-attributes-using-g-e-t-1"></span> Get all the SCIM attributes. (*getAllSCIMAttributesUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/idp/{idpId}/scimattribute
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| idpId | `path` | string | `string` |  | ✓ |  | The unique identifier of the IdP. |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-s-c-i-m-attributes-using-g-e-t-1-200) | OK | OK |  | [schema](#get-all-s-c-i-m-attributes-using-g-e-t-1-200-schema) |
| [400](#get-all-s-c-i-m-attributes-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-all-s-c-i-m-attributes-using-g-e-t-1-400-schema) |
| [401](#get-all-s-c-i-m-attributes-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-all-s-c-i-m-attributes-using-g-e-t-1-401-schema) |
| [403](#get-all-s-c-i-m-attributes-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-all-s-c-i-m-attributes-using-g-e-t-1-403-schema) |
| [404](#get-all-s-c-i-m-attributes-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-all-s-c-i-m-attributes-using-g-e-t-1-404-schema) |
| [429](#get-all-s-c-i-m-attributes-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-all-s-c-i-m-attributes-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-all-s-c-i-m-attributes-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-all-s-c-i-m-attributes-using-g-e-t-1-200-schema"></span> Schema
   
  

[PageListOfSCIMAttributeHeader](#page-list-of-s-c-i-m-attribute-header)

##### <span id="get-all-s-c-i-m-attributes-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-all-s-c-i-m-attributes-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-all-s-c-i-m-attributes-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-all-s-c-i-m-attributes-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-all-s-c-i-m-attributes-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-all-s-c-i-m-attributes-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-all-s-c-i-m-attributes-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-all-s-c-i-m-attributes-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-all-s-c-i-m-attributes-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-all-s-c-i-m-attributes-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-all-segment-groups-using-g-e-t-1"></span> Get all the configured Segment Groups. (*getAllSegmentGroupsUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-segment-groups-using-g-e-t-1-200) | OK | OK |  | [schema](#get-all-segment-groups-using-g-e-t-1-200-schema) |
| [400](#get-all-segment-groups-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-all-segment-groups-using-g-e-t-1-400-schema) |
| [401](#get-all-segment-groups-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-all-segment-groups-using-g-e-t-1-401-schema) |
| [403](#get-all-segment-groups-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-all-segment-groups-using-g-e-t-1-403-schema) |
| [404](#get-all-segment-groups-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-all-segment-groups-using-g-e-t-1-404-schema) |
| [429](#get-all-segment-groups-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-all-segment-groups-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-all-segment-groups-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-all-segment-groups-using-g-e-t-1-200-schema"></span> Schema
   
  

[PageListOfSegmentGroup](#page-list-of-segment-group)

##### <span id="get-all-segment-groups-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-all-segment-groups-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-all-segment-groups-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-all-segment-groups-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-all-segment-groups-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-all-segment-groups-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-all-segment-groups-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-all-segment-groups-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-all-segment-groups-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-all-segment-groups-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-all-server-groups-using-g-e-t-1"></span> Get all configured Server Groups. (*getAllServerGroupsUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/serverGroup
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-server-groups-using-g-e-t-1-200) | OK | OK |  | [schema](#get-all-server-groups-using-g-e-t-1-200-schema) |
| [400](#get-all-server-groups-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-all-server-groups-using-g-e-t-1-400-schema) |
| [401](#get-all-server-groups-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-all-server-groups-using-g-e-t-1-401-schema) |
| [403](#get-all-server-groups-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-all-server-groups-using-g-e-t-1-403-schema) |
| [404](#get-all-server-groups-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-all-server-groups-using-g-e-t-1-404-schema) |
| [429](#get-all-server-groups-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-all-server-groups-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-all-server-groups-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-all-server-groups-using-g-e-t-1-200-schema"></span> Schema
   
  

[PageListOfServerGroupDTO](#page-list-of-server-group-d-t-o)

##### <span id="get-all-server-groups-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-all-server-groups-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-all-server-groups-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-all-server-groups-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-all-server-groups-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-all-server-groups-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-all-server-groups-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-all-server-groups-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-all-server-groups-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-all-server-groups-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-all-service-edges-using-g-e-t"></span> Get all the configured Service Edge details. (*getAllServiceEdgesUsingGET*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-service-edges-using-g-e-t-200) | OK | OK |  | [schema](#get-all-service-edges-using-g-e-t-200-schema) |
| [400](#get-all-service-edges-using-g-e-t-400) | Bad Request | BadRequest |  | [schema](#get-all-service-edges-using-g-e-t-400-schema) |
| [401](#get-all-service-edges-using-g-e-t-401) | Unauthorized | Unauthorized |  | [schema](#get-all-service-edges-using-g-e-t-401-schema) |
| [403](#get-all-service-edges-using-g-e-t-403) | Forbidden | Forbidden |  | [schema](#get-all-service-edges-using-g-e-t-403-schema) |
| [404](#get-all-service-edges-using-g-e-t-404) | Not Found | Not Found |  | [schema](#get-all-service-edges-using-g-e-t-404-schema) |
| [429](#get-all-service-edges-using-g-e-t-429) | Too Many Requests | TooManyRequest |  | [schema](#get-all-service-edges-using-g-e-t-429-schema) |

#### Responses


##### <span id="get-all-service-edges-using-g-e-t-200"></span> 200 - OK
Status: OK

###### <span id="get-all-service-edges-using-g-e-t-200-schema"></span> Schema
   
  

[PageListOfServiceEdge](#page-list-of-service-edge)

##### <span id="get-all-service-edges-using-g-e-t-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-all-service-edges-using-g-e-t-400-schema"></span> Schema

##### <span id="get-all-service-edges-using-g-e-t-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-all-service-edges-using-g-e-t-401-schema"></span> Schema

##### <span id="get-all-service-edges-using-g-e-t-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-all-service-edges-using-g-e-t-403-schema"></span> Schema

##### <span id="get-all-service-edges-using-g-e-t-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-all-service-edges-using-g-e-t-404-schema"></span> Schema

##### <span id="get-all-service-edges-using-g-e-t-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-all-service-edges-using-g-e-t-429-schema"></span> Schema

### <span id="get-all-signing-cert-using-g-e-t-2"></span> Get all the configured Enrollment Cert details. (*getAllSigningCertUsingGET_2*)

```
GET /mgmtconfig/v2/admin/customers/{customerId}/enrollmentCert
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-signing-cert-using-g-e-t-2-200) | OK | OK |  | [schema](#get-all-signing-cert-using-g-e-t-2-200-schema) |
| [400](#get-all-signing-cert-using-g-e-t-2-400) | Bad Request | BadRequest |  | [schema](#get-all-signing-cert-using-g-e-t-2-400-schema) |
| [401](#get-all-signing-cert-using-g-e-t-2-401) | Unauthorized | Unauthorized |  | [schema](#get-all-signing-cert-using-g-e-t-2-401-schema) |
| [403](#get-all-signing-cert-using-g-e-t-2-403) | Forbidden | Forbidden |  | [schema](#get-all-signing-cert-using-g-e-t-2-403-schema) |
| [404](#get-all-signing-cert-using-g-e-t-2-404) | Not Found | Not Found |  | [schema](#get-all-signing-cert-using-g-e-t-2-404-schema) |
| [429](#get-all-signing-cert-using-g-e-t-2-429) | Too Many Requests | TooManyRequest |  | [schema](#get-all-signing-cert-using-g-e-t-2-429-schema) |

#### Responses


##### <span id="get-all-signing-cert-using-g-e-t-2-200"></span> 200 - OK
Status: OK

###### <span id="get-all-signing-cert-using-g-e-t-2-200-schema"></span> Schema
   
  

[PageListOfEnrollmentCert](#page-list-of-enrollment-cert)

##### <span id="get-all-signing-cert-using-g-e-t-2-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-all-signing-cert-using-g-e-t-2-400-schema"></span> Schema

##### <span id="get-all-signing-cert-using-g-e-t-2-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-all-signing-cert-using-g-e-t-2-401-schema"></span> Schema

##### <span id="get-all-signing-cert-using-g-e-t-2-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-all-signing-cert-using-g-e-t-2-403-schema"></span> Schema

##### <span id="get-all-signing-cert-using-g-e-t-2-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-all-signing-cert-using-g-e-t-2-404-schema"></span> Schema

##### <span id="get-all-signing-cert-using-g-e-t-2-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-all-signing-cert-using-g-e-t-2-429-schema"></span> Schema

### <span id="get-all-trusted-networks-using-g-e-t-1"></span> Get all the Trusted Networks. This API will be deprecated in a future release. (*getAllTrustedNetworksUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/network
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-trusted-networks-using-g-e-t-1-200) | OK | OK |  | [schema](#get-all-trusted-networks-using-g-e-t-1-200-schema) |
| [400](#get-all-trusted-networks-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-all-trusted-networks-using-g-e-t-1-400-schema) |
| [401](#get-all-trusted-networks-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-all-trusted-networks-using-g-e-t-1-401-schema) |
| [403](#get-all-trusted-networks-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-all-trusted-networks-using-g-e-t-1-403-schema) |
| [404](#get-all-trusted-networks-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-all-trusted-networks-using-g-e-t-1-404-schema) |
| [429](#get-all-trusted-networks-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-all-trusted-networks-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-all-trusted-networks-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-all-trusted-networks-using-g-e-t-1-200-schema"></span> Schema
   
  

[][TrustedNetwork](#trusted-network)

##### <span id="get-all-trusted-networks-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-all-trusted-networks-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-all-trusted-networks-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-all-trusted-networks-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-all-trusted-networks-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-all-trusted-networks-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-all-trusted-networks-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-all-trusted-networks-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-all-trusted-networks-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-all-trusted-networks-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-all-trusted-networks-using-g-e-t-2"></span> Get all the Trusted Networks. (*getAllTrustedNetworksUsingGET_2*)

```
GET /mgmtconfig/v2/admin/customers/{customerId}/network
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-trusted-networks-using-g-e-t-2-200) | OK | OK |  | [schema](#get-all-trusted-networks-using-g-e-t-2-200-schema) |
| [400](#get-all-trusted-networks-using-g-e-t-2-400) | Bad Request | BadRequest |  | [schema](#get-all-trusted-networks-using-g-e-t-2-400-schema) |
| [401](#get-all-trusted-networks-using-g-e-t-2-401) | Unauthorized | Unauthorized |  | [schema](#get-all-trusted-networks-using-g-e-t-2-401-schema) |
| [403](#get-all-trusted-networks-using-g-e-t-2-403) | Forbidden | Forbidden |  | [schema](#get-all-trusted-networks-using-g-e-t-2-403-schema) |
| [404](#get-all-trusted-networks-using-g-e-t-2-404) | Not Found | Not Found |  | [schema](#get-all-trusted-networks-using-g-e-t-2-404-schema) |
| [429](#get-all-trusted-networks-using-g-e-t-2-429) | Too Many Requests | TooManyRequest |  | [schema](#get-all-trusted-networks-using-g-e-t-2-429-schema) |

#### Responses


##### <span id="get-all-trusted-networks-using-g-e-t-2-200"></span> 200 - OK
Status: OK

###### <span id="get-all-trusted-networks-using-g-e-t-2-200-schema"></span> Schema
   
  

[PageListOfTrustedNetwork](#page-list-of-trusted-network)

##### <span id="get-all-trusted-networks-using-g-e-t-2-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-all-trusted-networks-using-g-e-t-2-400-schema"></span> Schema

##### <span id="get-all-trusted-networks-using-g-e-t-2-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-all-trusted-networks-using-g-e-t-2-401-schema"></span> Schema

##### <span id="get-all-trusted-networks-using-g-e-t-2-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-all-trusted-networks-using-g-e-t-2-403-schema"></span> Schema

##### <span id="get-all-trusted-networks-using-g-e-t-2-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-all-trusted-networks-using-g-e-t-2-404-schema"></span> Schema

##### <span id="get-all-trusted-networks-using-g-e-t-2-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-all-trusted-networks-using-g-e-t-2-429-schema"></span> Schema

### <span id="get-all-using-g-e-t-8"></span> Get all LSS configured for a given customer (*getAllUsingGET_8*)

```
GET /mgmtconfig/v2/admin/customers/{customerId}/lssConfig
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-using-g-e-t-8-200) | OK | OK |  | [schema](#get-all-using-g-e-t-8-200-schema) |
| [400](#get-all-using-g-e-t-8-400) | Bad Request | BadRequest |  | [schema](#get-all-using-g-e-t-8-400-schema) |
| [401](#get-all-using-g-e-t-8-401) | Unauthorized | Unauthorized |  | [schema](#get-all-using-g-e-t-8-401-schema) |
| [403](#get-all-using-g-e-t-8-403) | Forbidden | Forbidden |  | [schema](#get-all-using-g-e-t-8-403-schema) |
| [404](#get-all-using-g-e-t-8-404) | Not Found | Not Found |  | [schema](#get-all-using-g-e-t-8-404-schema) |
| [429](#get-all-using-g-e-t-8-429) | Too Many Requests | TooManyRequest |  | [schema](#get-all-using-g-e-t-8-429-schema) |

#### Responses


##### <span id="get-all-using-g-e-t-8-200"></span> 200 - OK
Status: OK

###### <span id="get-all-using-g-e-t-8-200-schema"></span> Schema
   
  

[PageListOfLssResource](#page-list-of-lss-resource)

##### <span id="get-all-using-g-e-t-8-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-all-using-g-e-t-8-400-schema"></span> Schema

##### <span id="get-all-using-g-e-t-8-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-all-using-g-e-t-8-401-schema"></span> Schema

##### <span id="get-all-using-g-e-t-8-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-all-using-g-e-t-8-403-schema"></span> Schema

##### <span id="get-all-using-g-e-t-8-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-all-using-g-e-t-8-404-schema"></span> Schema

##### <span id="get-all-using-g-e-t-8-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-all-using-g-e-t-8-429-schema"></span> Schema

### <span id="get-all-version-profiles-visibile-by-customer-id-using-g-e-t"></span> Get Version Profiles visible to a customer (*getAllVersionProfilesVisibileByCustomerIdUsingGET*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/visible/versionProfiles
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-all-version-profiles-visibile-by-customer-id-using-g-e-t-200) | OK | OK |  | [schema](#get-all-version-profiles-visibile-by-customer-id-using-g-e-t-200-schema) |
| [400](#get-all-version-profiles-visibile-by-customer-id-using-g-e-t-400) | Bad Request | BadRequest |  | [schema](#get-all-version-profiles-visibile-by-customer-id-using-g-e-t-400-schema) |
| [401](#get-all-version-profiles-visibile-by-customer-id-using-g-e-t-401) | Unauthorized | Unauthorized |  | [schema](#get-all-version-profiles-visibile-by-customer-id-using-g-e-t-401-schema) |
| [403](#get-all-version-profiles-visibile-by-customer-id-using-g-e-t-403) | Forbidden | Forbidden |  | [schema](#get-all-version-profiles-visibile-by-customer-id-using-g-e-t-403-schema) |
| [404](#get-all-version-profiles-visibile-by-customer-id-using-g-e-t-404) | Not Found | Not Found |  | [schema](#get-all-version-profiles-visibile-by-customer-id-using-g-e-t-404-schema) |
| [429](#get-all-version-profiles-visibile-by-customer-id-using-g-e-t-429) | Too Many Requests | TooManyRequest |  | [schema](#get-all-version-profiles-visibile-by-customer-id-using-g-e-t-429-schema) |

#### Responses


##### <span id="get-all-version-profiles-visibile-by-customer-id-using-g-e-t-200"></span> 200 - OK
Status: OK

###### <span id="get-all-version-profiles-visibile-by-customer-id-using-g-e-t-200-schema"></span> Schema
   
  

[PageListOfVersionProfile](#page-list-of-version-profile)

##### <span id="get-all-version-profiles-visibile-by-customer-id-using-g-e-t-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-all-version-profiles-visibile-by-customer-id-using-g-e-t-400-schema"></span> Schema

##### <span id="get-all-version-profiles-visibile-by-customer-id-using-g-e-t-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-all-version-profiles-visibile-by-customer-id-using-g-e-t-401-schema"></span> Schema

##### <span id="get-all-version-profiles-visibile-by-customer-id-using-g-e-t-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-all-version-profiles-visibile-by-customer-id-using-g-e-t-403-schema"></span> Schema

##### <span id="get-all-version-profiles-visibile-by-customer-id-using-g-e-t-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-all-version-profiles-visibile-by-customer-id-using-g-e-t-404-schema"></span> Schema

##### <span id="get-all-version-profiles-visibile-by-customer-id-using-g-e-t-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-all-version-profiles-visibile-by-customer-id-using-g-e-t-429-schema"></span> Schema

### <span id="get-app-connector-group-using-g-e-t-1"></span> Get the App Connector Group details. (*getAppConnectorGroupUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/appConnectorGroup/{appConnectorGroupId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| appConnectorGroupId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Connector Group. |
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-app-connector-group-using-g-e-t-1-200) | OK | OK |  | [schema](#get-app-connector-group-using-g-e-t-1-200-schema) |
| [400](#get-app-connector-group-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-app-connector-group-using-g-e-t-1-400-schema) |
| [401](#get-app-connector-group-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-app-connector-group-using-g-e-t-1-401-schema) |
| [403](#get-app-connector-group-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-app-connector-group-using-g-e-t-1-403-schema) |
| [404](#get-app-connector-group-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-app-connector-group-using-g-e-t-1-404-schema) |
| [429](#get-app-connector-group-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-app-connector-group-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-app-connector-group-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-app-connector-group-using-g-e-t-1-200-schema"></span> Schema
   
  

[AppConnectorGroup](#app-connector-group)

##### <span id="get-app-connector-group-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-app-connector-group-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-app-connector-group-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-app-connector-group-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-app-connector-group-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-app-connector-group-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-app-connector-group-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-app-connector-group-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-app-connector-group-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-app-connector-group-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-app-connector-groups-using-g-e-t-1"></span> Get all configured App Connector Groups. (*getAppConnectorGroupsUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/appConnectorGroup
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-app-connector-groups-using-g-e-t-1-200) | OK | OK |  | [schema](#get-app-connector-groups-using-g-e-t-1-200-schema) |
| [400](#get-app-connector-groups-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-app-connector-groups-using-g-e-t-1-400-schema) |
| [401](#get-app-connector-groups-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-app-connector-groups-using-g-e-t-1-401-schema) |
| [403](#get-app-connector-groups-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-app-connector-groups-using-g-e-t-1-403-schema) |
| [404](#get-app-connector-groups-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-app-connector-groups-using-g-e-t-1-404-schema) |
| [429](#get-app-connector-groups-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-app-connector-groups-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-app-connector-groups-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-app-connector-groups-using-g-e-t-1-200-schema"></span> Schema
   
  

[PageListOfAppConnectorGroup](#page-list-of-app-connector-group)

##### <span id="get-app-connector-groups-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-app-connector-groups-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-app-connector-groups-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-app-connector-groups-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-app-connector-groups-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-app-connector-groups-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-app-connector-groups-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-app-connector-groups-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-app-connector-groups-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-app-connector-groups-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-app-server-using-g-e-t-1"></span> Get the Server details. (*getAppServerUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/server/{serverId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| serverId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Server. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-app-server-using-g-e-t-1-200) | OK | OK |  | [schema](#get-app-server-using-g-e-t-1-200-schema) |
| [400](#get-app-server-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-app-server-using-g-e-t-1-400-schema) |
| [401](#get-app-server-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-app-server-using-g-e-t-1-401-schema) |
| [403](#get-app-server-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-app-server-using-g-e-t-1-403-schema) |
| [404](#get-app-server-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-app-server-using-g-e-t-1-404-schema) |
| [429](#get-app-server-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-app-server-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-app-server-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-app-server-using-g-e-t-1-200-schema"></span> Schema
   
  

[ApplicationServer](#application-server)

##### <span id="get-app-server-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-app-server-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-app-server-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-app-server-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-app-server-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-app-server-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-app-server-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-app-server-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-app-server-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-app-server-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-application-using-g-e-t-1"></span> Get the Application Segment details. (*getApplicationUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/application/{applicationId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| applicationId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Application Segment. |
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-application-using-g-e-t-1-200) | OK | OK |  | [schema](#get-application-using-g-e-t-1-200-schema) |
| [400](#get-application-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-application-using-g-e-t-1-400-schema) |
| [401](#get-application-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-application-using-g-e-t-1-401-schema) |
| [403](#get-application-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-application-using-g-e-t-1-403-schema) |
| [404](#get-application-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-application-using-g-e-t-1-404-schema) |
| [429](#get-application-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-application-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-application-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-application-using-g-e-t-1-200-schema"></span> Schema
   
  

[ApplicationResource](#application-resource)

##### <span id="get-application-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-application-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-application-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-application-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-application-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-application-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-application-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-application-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-application-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-application-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-bypass-policy-set-using-g-e-t-1"></span> Get the bypass policy and all rules for a Client Forwarding policy rule. This API will be deprecated in a future release. (*getBypassPolicySetUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/bypass
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-bypass-policy-set-using-g-e-t-1-200) | OK | OK |  | [schema](#get-bypass-policy-set-using-g-e-t-1-200-schema) |
| [400](#get-bypass-policy-set-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-bypass-policy-set-using-g-e-t-1-400-schema) |
| [401](#get-bypass-policy-set-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-bypass-policy-set-using-g-e-t-1-401-schema) |
| [403](#get-bypass-policy-set-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-bypass-policy-set-using-g-e-t-1-403-schema) |
| [404](#get-bypass-policy-set-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-bypass-policy-set-using-g-e-t-1-404-schema) |
| [429](#get-bypass-policy-set-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-bypass-policy-set-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-bypass-policy-set-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-bypass-policy-set-using-g-e-t-1-200-schema"></span> Schema
   
  

[PolicySet](#policy-set)

##### <span id="get-bypass-policy-set-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-bypass-policy-set-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-bypass-policy-set-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-bypass-policy-set-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-bypass-policy-set-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-bypass-policy-set-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-bypass-policy-set-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-bypass-policy-set-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-bypass-policy-set-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-bypass-policy-set-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-cloud-connector-group-using-g-e-t"></span> Get the Cloud Connector Group details. (*getCloudConnectorGroupUsingGET*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/cloudConnectorGroup/{id}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| id | `path` | string | `string` |  | ✓ |  | The unique identifier of the Cloud Connector Group. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cloud-connector-group-using-g-e-t-200) | OK | OK |  | [schema](#get-cloud-connector-group-using-g-e-t-200-schema) |
| [400](#get-cloud-connector-group-using-g-e-t-400) | Bad Request | BadRequest |  | [schema](#get-cloud-connector-group-using-g-e-t-400-schema) |
| [401](#get-cloud-connector-group-using-g-e-t-401) | Unauthorized | Unauthorized |  | [schema](#get-cloud-connector-group-using-g-e-t-401-schema) |
| [403](#get-cloud-connector-group-using-g-e-t-403) | Forbidden | Forbidden |  | [schema](#get-cloud-connector-group-using-g-e-t-403-schema) |
| [404](#get-cloud-connector-group-using-g-e-t-404) | Not Found | Not Found |  | [schema](#get-cloud-connector-group-using-g-e-t-404-schema) |
| [429](#get-cloud-connector-group-using-g-e-t-429) | Too Many Requests | TooManyRequest |  | [schema](#get-cloud-connector-group-using-g-e-t-429-schema) |

#### Responses


##### <span id="get-cloud-connector-group-using-g-e-t-200"></span> 200 - OK
Status: OK

###### <span id="get-cloud-connector-group-using-g-e-t-200-schema"></span> Schema
   
  

[CloudConnectorGroupResource](#cloud-connector-group-resource)

##### <span id="get-cloud-connector-group-using-g-e-t-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-cloud-connector-group-using-g-e-t-400-schema"></span> Schema

##### <span id="get-cloud-connector-group-using-g-e-t-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-cloud-connector-group-using-g-e-t-401-schema"></span> Schema

##### <span id="get-cloud-connector-group-using-g-e-t-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-cloud-connector-group-using-g-e-t-403-schema"></span> Schema

##### <span id="get-cloud-connector-group-using-g-e-t-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-cloud-connector-group-using-g-e-t-404-schema"></span> Schema

##### <span id="get-cloud-connector-group-using-g-e-t-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-cloud-connector-group-using-g-e-t-429-schema"></span> Schema

### <span id="get-cloud-connector-groups-using-g-e-t"></span> Get all configured Cloud Connector Groups. (*getCloudConnectorGroupsUsingGET*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/cloudConnectorGroup
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-cloud-connector-groups-using-g-e-t-200) | OK | OK |  | [schema](#get-cloud-connector-groups-using-g-e-t-200-schema) |
| [400](#get-cloud-connector-groups-using-g-e-t-400) | Bad Request | BadRequest |  | [schema](#get-cloud-connector-groups-using-g-e-t-400-schema) |
| [401](#get-cloud-connector-groups-using-g-e-t-401) | Unauthorized | Unauthorized |  | [schema](#get-cloud-connector-groups-using-g-e-t-401-schema) |
| [403](#get-cloud-connector-groups-using-g-e-t-403) | Forbidden | Forbidden |  | [schema](#get-cloud-connector-groups-using-g-e-t-403-schema) |
| [404](#get-cloud-connector-groups-using-g-e-t-404) | Not Found | Not Found |  | [schema](#get-cloud-connector-groups-using-g-e-t-404-schema) |
| [429](#get-cloud-connector-groups-using-g-e-t-429) | Too Many Requests | TooManyRequest |  | [schema](#get-cloud-connector-groups-using-g-e-t-429-schema) |

#### Responses


##### <span id="get-cloud-connector-groups-using-g-e-t-200"></span> 200 - OK
Status: OK

###### <span id="get-cloud-connector-groups-using-g-e-t-200-schema"></span> Schema
   
  

[PageListOfCloudConnectorGroupResource](#page-list-of-cloud-connector-group-resource)

##### <span id="get-cloud-connector-groups-using-g-e-t-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-cloud-connector-groups-using-g-e-t-400-schema"></span> Schema

##### <span id="get-cloud-connector-groups-using-g-e-t-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-cloud-connector-groups-using-g-e-t-401-schema"></span> Schema

##### <span id="get-cloud-connector-groups-using-g-e-t-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-cloud-connector-groups-using-g-e-t-403-schema"></span> Schema

##### <span id="get-cloud-connector-groups-using-g-e-t-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-cloud-connector-groups-using-g-e-t-404-schema"></span> Schema

##### <span id="get-cloud-connector-groups-using-g-e-t-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-cloud-connector-groups-using-g-e-t-429-schema"></span> Schema

### <span id="get-connector-using-g-e-t-1"></span> Get the Connector details. (*getConnectorUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/connector/{connectorId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| connectorId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Connector. |
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-connector-using-g-e-t-1-200) | OK | OK |  | [schema](#get-connector-using-g-e-t-1-200-schema) |
| [400](#get-connector-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-connector-using-g-e-t-1-400-schema) |
| [401](#get-connector-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-connector-using-g-e-t-1-401-schema) |
| [403](#get-connector-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-connector-using-g-e-t-1-403-schema) |
| [404](#get-connector-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-connector-using-g-e-t-1-404-schema) |
| [429](#get-connector-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-connector-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-connector-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-connector-using-g-e-t-1-200-schema"></span> Schema
   
  

[Connector](#connector)

##### <span id="get-connector-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-connector-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-connector-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-connector-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-connector-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-connector-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-connector-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-connector-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-connector-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-connector-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-enrollment-cert-using-g-e-t-3"></span> Get the Enrollment Cert details. (*getEnrollmentCertUsingGET_3*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/enrollmentCert/{enrollmentCertId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| enrollmentCertId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Enrollment Cert. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-enrollment-cert-using-g-e-t-3-200) | OK | OK |  | [schema](#get-enrollment-cert-using-g-e-t-3-200-schema) |
| [400](#get-enrollment-cert-using-g-e-t-3-400) | Bad Request | BadRequest |  | [schema](#get-enrollment-cert-using-g-e-t-3-400-schema) |
| [401](#get-enrollment-cert-using-g-e-t-3-401) | Unauthorized | Unauthorized |  | [schema](#get-enrollment-cert-using-g-e-t-3-401-schema) |
| [403](#get-enrollment-cert-using-g-e-t-3-403) | Forbidden | Forbidden |  | [schema](#get-enrollment-cert-using-g-e-t-3-403-schema) |
| [404](#get-enrollment-cert-using-g-e-t-3-404) | Not Found | Not Found |  | [schema](#get-enrollment-cert-using-g-e-t-3-404-schema) |
| [429](#get-enrollment-cert-using-g-e-t-3-429) | Too Many Requests | TooManyRequest |  | [schema](#get-enrollment-cert-using-g-e-t-3-429-schema) |

#### Responses


##### <span id="get-enrollment-cert-using-g-e-t-3-200"></span> 200 - OK
Status: OK

###### <span id="get-enrollment-cert-using-g-e-t-3-200-schema"></span> Schema
   
  

[EnrollmentCert](#enrollment-cert)

##### <span id="get-enrollment-cert-using-g-e-t-3-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-enrollment-cert-using-g-e-t-3-400-schema"></span> Schema

##### <span id="get-enrollment-cert-using-g-e-t-3-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-enrollment-cert-using-g-e-t-3-401-schema"></span> Schema

##### <span id="get-enrollment-cert-using-g-e-t-3-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-enrollment-cert-using-g-e-t-3-403-schema"></span> Schema

##### <span id="get-enrollment-cert-using-g-e-t-3-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-enrollment-cert-using-g-e-t-3-404-schema"></span> Schema

##### <span id="get-enrollment-cert-using-g-e-t-3-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-enrollment-cert-using-g-e-t-3-429-schema"></span> Schema

### <span id="get-global-policy-set-using-g-e-t-1"></span> Get the global policy. This API will be deprecated in a future release. (*getGlobalPolicySetUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/global
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-global-policy-set-using-g-e-t-1-200) | OK | OK |  | [schema](#get-global-policy-set-using-g-e-t-1-200-schema) |
| [400](#get-global-policy-set-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-global-policy-set-using-g-e-t-1-400-schema) |
| [401](#get-global-policy-set-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-global-policy-set-using-g-e-t-1-401-schema) |
| [403](#get-global-policy-set-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-global-policy-set-using-g-e-t-1-403-schema) |
| [404](#get-global-policy-set-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-global-policy-set-using-g-e-t-1-404-schema) |
| [429](#get-global-policy-set-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-global-policy-set-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-global-policy-set-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-global-policy-set-using-g-e-t-1-200-schema"></span> Schema
   
  

[PolicySet](#policy-set)

##### <span id="get-global-policy-set-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-global-policy-set-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-global-policy-set-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-global-policy-set-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-global-policy-set-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-global-policy-set-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-global-policy-set-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-global-policy-set-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-global-policy-set-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-global-policy-set-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-idp-by-id-using-g-e-t-1"></span> Get all the IdP details. (*getIdpByIdUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/idp/{idpId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| idpId | `path` | string | `string` |  | ✓ |  | The unique identifier of the IdP. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-idp-by-id-using-g-e-t-1-200) | OK | OK |  | [schema](#get-idp-by-id-using-g-e-t-1-200-schema) |
| [400](#get-idp-by-id-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-idp-by-id-using-g-e-t-1-400-schema) |
| [401](#get-idp-by-id-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-idp-by-id-using-g-e-t-1-401-schema) |
| [403](#get-idp-by-id-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-idp-by-id-using-g-e-t-1-403-schema) |
| [404](#get-idp-by-id-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-idp-by-id-using-g-e-t-1-404-schema) |
| [429](#get-idp-by-id-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-idp-by-id-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-idp-by-id-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-idp-by-id-using-g-e-t-1-200-schema"></span> Schema
   
  

[Idp](#idp)

##### <span id="get-idp-by-id-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-idp-by-id-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-idp-by-id-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-idp-by-id-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-idp-by-id-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-idp-by-id-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-idp-by-id-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-idp-by-id-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-idp-by-id-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-idp-by-id-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-idp-using-g-e-t-1"></span> Get the configured IdP details. This API will be deprecated in a future release. (*getIdpUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/idp
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| scimEnabled | `query` | boolean | `bool` |  |  |  | Returns all SCIM IdPs if set to true. Returns all non SCIM IdPs if set to false. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-idp-using-g-e-t-1-200) | OK | OK |  | [schema](#get-idp-using-g-e-t-1-200-schema) |
| [400](#get-idp-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-idp-using-g-e-t-1-400-schema) |
| [401](#get-idp-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-idp-using-g-e-t-1-401-schema) |
| [403](#get-idp-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-idp-using-g-e-t-1-403-schema) |
| [404](#get-idp-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-idp-using-g-e-t-1-404-schema) |
| [429](#get-idp-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-idp-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-idp-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-idp-using-g-e-t-1-200-schema"></span> Schema
   
  

[][Idp](#idp)

##### <span id="get-idp-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-idp-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-idp-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-idp-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-idp-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-idp-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-idp-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-idp-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-idp-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-idp-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-list-of-lss-client-types-using-g-e-t-1"></span> Get list of all LSS client types. (*getListOfLssClientTypesUsingGET_1*)

```
GET /mgmtconfig/v2/admin/lssConfig/clientTypes
```

#### Produces
  * */*

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-list-of-lss-client-types-using-g-e-t-1-200) | OK | OK |  | [schema](#get-list-of-lss-client-types-using-g-e-t-1-200-schema) |
| [400](#get-list-of-lss-client-types-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-list-of-lss-client-types-using-g-e-t-1-400-schema) |
| [401](#get-list-of-lss-client-types-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-list-of-lss-client-types-using-g-e-t-1-401-schema) |
| [403](#get-list-of-lss-client-types-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-list-of-lss-client-types-using-g-e-t-1-403-schema) |
| [404](#get-list-of-lss-client-types-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-list-of-lss-client-types-using-g-e-t-1-404-schema) |
| [429](#get-list-of-lss-client-types-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-list-of-lss-client-types-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-list-of-lss-client-types-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-list-of-lss-client-types-using-g-e-t-1-200-schema"></span> Schema
   
  

map of string

##### <span id="get-list-of-lss-client-types-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-list-of-lss-client-types-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-list-of-lss-client-types-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-list-of-lss-client-types-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-list-of-lss-client-types-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-list-of-lss-client-types-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-list-of-lss-client-types-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-list-of-lss-client-types-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-list-of-lss-client-types-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-list-of-lss-client-types-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-lss-log-format-using-g-e-t-3"></span> Get all LSS log format. (*getLssLogFormatUsingGET_3*)

```
GET /mgmtconfig/v2/admin/lssConfig/logType/formats
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| logType | `query` | string | `string` |  |  |  | Supported Log type values are : zpn_trans_log, zpn_auth_log, zpn_ast_auth_log, zpn_http_trans_log, zpn_audit_log. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-lss-log-format-using-g-e-t-3-200) | OK | OK |  | [schema](#get-lss-log-format-using-g-e-t-3-200-schema) |
| [400](#get-lss-log-format-using-g-e-t-3-400) | Bad Request | BadRequest |  | [schema](#get-lss-log-format-using-g-e-t-3-400-schema) |
| [401](#get-lss-log-format-using-g-e-t-3-401) | Unauthorized | Unauthorized |  | [schema](#get-lss-log-format-using-g-e-t-3-401-schema) |
| [403](#get-lss-log-format-using-g-e-t-3-403) | Forbidden | Forbidden |  | [schema](#get-lss-log-format-using-g-e-t-3-403-schema) |
| [404](#get-lss-log-format-using-g-e-t-3-404) | Not Found | Not Found |  | [schema](#get-lss-log-format-using-g-e-t-3-404-schema) |
| [429](#get-lss-log-format-using-g-e-t-3-429) | Too Many Requests | TooManyRequest |  | [schema](#get-lss-log-format-using-g-e-t-3-429-schema) |

#### Responses


##### <span id="get-lss-log-format-using-g-e-t-3-200"></span> 200 - OK
Status: OK

###### <span id="get-lss-log-format-using-g-e-t-3-200-schema"></span> Schema
   
  

map of any 

##### <span id="get-lss-log-format-using-g-e-t-3-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-lss-log-format-using-g-e-t-3-400-schema"></span> Schema

##### <span id="get-lss-log-format-using-g-e-t-3-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-lss-log-format-using-g-e-t-3-401-schema"></span> Schema

##### <span id="get-lss-log-format-using-g-e-t-3-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-lss-log-format-using-g-e-t-3-403-schema"></span> Schema

##### <span id="get-lss-log-format-using-g-e-t-3-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-lss-log-format-using-g-e-t-3-404-schema"></span> Schema

##### <span id="get-lss-log-format-using-g-e-t-3-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-lss-log-format-using-g-e-t-3-429-schema"></span> Schema

### <span id="get-lss-using-g-e-t-3"></span> Get the LSS details (*getLssUsingGET_3*)

```
GET /mgmtconfig/v2/admin/customers/{customerId}/lssConfig/{lssId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| lssId | `path` | string | `string` |  | ✓ |  | The unique identifier of the LSS resource. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-lss-using-g-e-t-3-200) | OK | OK |  | [schema](#get-lss-using-g-e-t-3-200-schema) |
| [400](#get-lss-using-g-e-t-3-400) | Bad Request | BadRequest |  | [schema](#get-lss-using-g-e-t-3-400-schema) |
| [401](#get-lss-using-g-e-t-3-401) | Unauthorized | Unauthorized |  | [schema](#get-lss-using-g-e-t-3-401-schema) |
| [403](#get-lss-using-g-e-t-3-403) | Forbidden | Forbidden |  | [schema](#get-lss-using-g-e-t-3-403-schema) |
| [404](#get-lss-using-g-e-t-3-404) | Not Found | Not Found |  | [schema](#get-lss-using-g-e-t-3-404-schema) |
| [429](#get-lss-using-g-e-t-3-429) | Too Many Requests | TooManyRequest |  | [schema](#get-lss-using-g-e-t-3-429-schema) |

#### Responses


##### <span id="get-lss-using-g-e-t-3-200"></span> 200 - OK
Status: OK

###### <span id="get-lss-using-g-e-t-3-200-schema"></span> Schema
   
  

[LssResource](#lss-resource)

##### <span id="get-lss-using-g-e-t-3-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-lss-using-g-e-t-3-400-schema"></span> Schema

##### <span id="get-lss-using-g-e-t-3-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-lss-using-g-e-t-3-401-schema"></span> Schema

##### <span id="get-lss-using-g-e-t-3-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-lss-using-g-e-t-3-403-schema"></span> Schema

##### <span id="get-lss-using-g-e-t-3-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-lss-using-g-e-t-3-404-schema"></span> Schema

##### <span id="get-lss-using-g-e-t-3-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-lss-using-g-e-t-3-429-schema"></span> Schema

### <span id="get-machine-group-using-g-e-t"></span> Get the Machine Group details. (*getMachineGroupUsingGET*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/machineGroup/{Id}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Id | `path` | string | `string` |  | ✓ |  | The unique identifier of the Machine Group. |
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-machine-group-using-g-e-t-200) | OK | OK |  | [schema](#get-machine-group-using-g-e-t-200-schema) |
| [400](#get-machine-group-using-g-e-t-400) | Bad Request | BadRequest |  | [schema](#get-machine-group-using-g-e-t-400-schema) |
| [401](#get-machine-group-using-g-e-t-401) | Unauthorized | Unauthorized |  | [schema](#get-machine-group-using-g-e-t-401-schema) |
| [403](#get-machine-group-using-g-e-t-403) | Forbidden | Forbidden |  | [schema](#get-machine-group-using-g-e-t-403-schema) |
| [404](#get-machine-group-using-g-e-t-404) | Not Found | Not Found |  | [schema](#get-machine-group-using-g-e-t-404-schema) |
| [429](#get-machine-group-using-g-e-t-429) | Too Many Requests | TooManyRequest |  | [schema](#get-machine-group-using-g-e-t-429-schema) |

#### Responses


##### <span id="get-machine-group-using-g-e-t-200"></span> 200 - OK
Status: OK

###### <span id="get-machine-group-using-g-e-t-200-schema"></span> Schema
   
  

[MachineGroup](#machine-group)

##### <span id="get-machine-group-using-g-e-t-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-machine-group-using-g-e-t-400-schema"></span> Schema

##### <span id="get-machine-group-using-g-e-t-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-machine-group-using-g-e-t-401-schema"></span> Schema

##### <span id="get-machine-group-using-g-e-t-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-machine-group-using-g-e-t-403-schema"></span> Schema

##### <span id="get-machine-group-using-g-e-t-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-machine-group-using-g-e-t-404-schema"></span> Schema

##### <span id="get-machine-group-using-g-e-t-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-machine-group-using-g-e-t-429-schema"></span> Schema

### <span id="get-machine-groups-using-g-e-t"></span> Get all the configured Machine Groups. (*getMachineGroupsUsingGET*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/machineGroup
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-machine-groups-using-g-e-t-200) | OK | OK |  | [schema](#get-machine-groups-using-g-e-t-200-schema) |
| [400](#get-machine-groups-using-g-e-t-400) | Bad Request | BadRequest |  | [schema](#get-machine-groups-using-g-e-t-400-schema) |
| [401](#get-machine-groups-using-g-e-t-401) | Unauthorized | Unauthorized |  | [schema](#get-machine-groups-using-g-e-t-401-schema) |
| [403](#get-machine-groups-using-g-e-t-403) | Forbidden | Forbidden |  | [schema](#get-machine-groups-using-g-e-t-403-schema) |
| [404](#get-machine-groups-using-g-e-t-404) | Not Found | Not Found |  | [schema](#get-machine-groups-using-g-e-t-404-schema) |
| [429](#get-machine-groups-using-g-e-t-429) | Too Many Requests | TooManyRequest |  | [schema](#get-machine-groups-using-g-e-t-429-schema) |

#### Responses


##### <span id="get-machine-groups-using-g-e-t-200"></span> 200 - OK
Status: OK

###### <span id="get-machine-groups-using-g-e-t-200-schema"></span> Schema
   
  

[PageListOfMachineGroup](#page-list-of-machine-group)

##### <span id="get-machine-groups-using-g-e-t-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-machine-groups-using-g-e-t-400-schema"></span> Schema

##### <span id="get-machine-groups-using-g-e-t-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-machine-groups-using-g-e-t-401-schema"></span> Schema

##### <span id="get-machine-groups-using-g-e-t-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-machine-groups-using-g-e-t-403-schema"></span> Schema

##### <span id="get-machine-groups-using-g-e-t-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-machine-groups-using-g-e-t-404-schema"></span> Schema

##### <span id="get-machine-groups-using-g-e-t-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-machine-groups-using-g-e-t-429-schema"></span> Schema

### <span id="get-policy-rules-by-page-using-g-e-t-1"></span> Get paginated policy rules for a given policy type. (*getPolicyRulesByPageUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/rules/policyType/{policyType}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| policyType | `path` | string | `string` |  | ✓ |  | Type for differentiating policy types. Supported values : ACCESS_POLICY/GLOBAL_POLICY, TIMEOUT_POLICY/REAUTH_POLICY, SIEM_POLICY, CLIENT_FORWARDING_POLICY/BYPASS_POLICY |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-policy-rules-by-page-using-g-e-t-1-200) | OK | OK |  | [schema](#get-policy-rules-by-page-using-g-e-t-1-200-schema) |
| [400](#get-policy-rules-by-page-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-policy-rules-by-page-using-g-e-t-1-400-schema) |
| [401](#get-policy-rules-by-page-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-policy-rules-by-page-using-g-e-t-1-401-schema) |
| [403](#get-policy-rules-by-page-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-policy-rules-by-page-using-g-e-t-1-403-schema) |
| [404](#get-policy-rules-by-page-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-policy-rules-by-page-using-g-e-t-1-404-schema) |
| [429](#get-policy-rules-by-page-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-policy-rules-by-page-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-policy-rules-by-page-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-policy-rules-by-page-using-g-e-t-1-200-schema"></span> Schema
   
  

[PageListOfPolicyRule](#page-list-of-policy-rule)

##### <span id="get-policy-rules-by-page-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-policy-rules-by-page-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-policy-rules-by-page-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-policy-rules-by-page-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-policy-rules-by-page-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-policy-rules-by-page-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-policy-rules-by-page-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-policy-rules-by-page-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-policy-rules-by-page-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-policy-rules-by-page-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-policy-set-by-policy-type-using-g-e-t-1"></span> For a customer, get policy set by policy type (*getPolicySetByPolicyTypeUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/policyType/{policyType}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | customerId |
| policyType | `path` | string | `string` |  | ✓ |  | Type for differentiating policy types. Supported values : ACCESS_POLICY/GLOBAL_POLICY, TIMEOUT_POLICY/REAUTH_POLICY, SIEM_POLICY, CLIENT_FORWARDING_POLICY/BYPASS_POLICY |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-policy-set-by-policy-type-using-g-e-t-1-200) | OK | OK |  | [schema](#get-policy-set-by-policy-type-using-g-e-t-1-200-schema) |
| [400](#get-policy-set-by-policy-type-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-policy-set-by-policy-type-using-g-e-t-1-400-schema) |
| [401](#get-policy-set-by-policy-type-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-policy-set-by-policy-type-using-g-e-t-1-401-schema) |
| [403](#get-policy-set-by-policy-type-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-policy-set-by-policy-type-using-g-e-t-1-403-schema) |
| [404](#get-policy-set-by-policy-type-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-policy-set-by-policy-type-using-g-e-t-1-404-schema) |
| [429](#get-policy-set-by-policy-type-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-policy-set-by-policy-type-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-policy-set-by-policy-type-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-policy-set-by-policy-type-using-g-e-t-1-200-schema"></span> Schema
   
  

[PolicySet](#policy-set)

##### <span id="get-policy-set-by-policy-type-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-policy-set-by-policy-type-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-policy-set-by-policy-type-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-policy-set-by-policy-type-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-policy-set-by-policy-type-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-policy-set-by-policy-type-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-policy-set-by-policy-type-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-policy-set-by-policy-type-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-policy-set-by-policy-type-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-policy-set-by-policy-type-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-posture-profile-using-g-e-t-1"></span> Get the configured profile posture. (*getPostureProfileUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/posture/{id}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| id | `path` | string | `string` |  | ✓ |  | The unique identifier of the posture profile. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-posture-profile-using-g-e-t-1-200) | OK | OK |  | [schema](#get-posture-profile-using-g-e-t-1-200-schema) |
| [400](#get-posture-profile-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-posture-profile-using-g-e-t-1-400-schema) |
| [401](#get-posture-profile-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-posture-profile-using-g-e-t-1-401-schema) |
| [403](#get-posture-profile-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-posture-profile-using-g-e-t-1-403-schema) |
| [404](#get-posture-profile-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-posture-profile-using-g-e-t-1-404-schema) |
| [429](#get-posture-profile-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-posture-profile-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-posture-profile-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-posture-profile-using-g-e-t-1-200-schema"></span> Schema
   
  

[PostureProfile](#posture-profile)

##### <span id="get-posture-profile-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-posture-profile-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-posture-profile-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-posture-profile-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-posture-profile-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-posture-profile-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-posture-profile-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-posture-profile-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-posture-profile-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-posture-profile-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-provisioning-key-for-association-type-using-g-e-t-1"></span> Get all the configured Provisioning Key details. (*getProvisioningKeyForAssociationTypeUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/associationType/{associationType}/provisioningKey
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| associationType | `path` | string | `string` |  | ✓ |  | Specifies the provisioning key type for App Connectors or ZPA Private Service Edges. The supported values are CONNECTOR_GRP and SERVICE_EDGE_GRP. |
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-provisioning-key-for-association-type-using-g-e-t-1-200) | OK | OK |  | [schema](#get-provisioning-key-for-association-type-using-g-e-t-1-200-schema) |
| [400](#get-provisioning-key-for-association-type-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-provisioning-key-for-association-type-using-g-e-t-1-400-schema) |
| [401](#get-provisioning-key-for-association-type-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-provisioning-key-for-association-type-using-g-e-t-1-401-schema) |
| [403](#get-provisioning-key-for-association-type-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-provisioning-key-for-association-type-using-g-e-t-1-403-schema) |
| [404](#get-provisioning-key-for-association-type-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-provisioning-key-for-association-type-using-g-e-t-1-404-schema) |
| [429](#get-provisioning-key-for-association-type-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-provisioning-key-for-association-type-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-provisioning-key-for-association-type-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-provisioning-key-for-association-type-using-g-e-t-1-200-schema"></span> Schema
   
  

[PageListOfProvisioningKey](#page-list-of-provisioning-key)

##### <span id="get-provisioning-key-for-association-type-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-provisioning-key-for-association-type-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-provisioning-key-for-association-type-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-provisioning-key-for-association-type-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-provisioning-key-for-association-type-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-provisioning-key-for-association-type-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-provisioning-key-for-association-type-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-provisioning-key-for-association-type-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-provisioning-key-for-association-type-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-provisioning-key-for-association-type-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-provisioning-key-using-g-e-t-1"></span> Get the Provisioning Key details. (*getProvisioningKeyUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/associationType/{associationType}/provisioningKey/{provisioningKeyId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| associationType | `path` | string | `string` |  | ✓ |  | Specifies the provisioning key type for App Connectors or ZPA Private Service Edges. The supported values are CONNECTOR_GRP and SERVICE_EDGE_GRP. |
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| provisioningKeyId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Provisioning Key. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-provisioning-key-using-g-e-t-1-200) | OK | OK |  | [schema](#get-provisioning-key-using-g-e-t-1-200-schema) |
| [400](#get-provisioning-key-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-provisioning-key-using-g-e-t-1-400-schema) |
| [401](#get-provisioning-key-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-provisioning-key-using-g-e-t-1-401-schema) |
| [403](#get-provisioning-key-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-provisioning-key-using-g-e-t-1-403-schema) |
| [404](#get-provisioning-key-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-provisioning-key-using-g-e-t-1-404-schema) |
| [429](#get-provisioning-key-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-provisioning-key-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-provisioning-key-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-provisioning-key-using-g-e-t-1-200-schema"></span> Schema
   
  

[ProvisioningKey](#provisioning-key)

##### <span id="get-provisioning-key-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-provisioning-key-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-provisioning-key-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-provisioning-key-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-provisioning-key-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-provisioning-key-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-provisioning-key-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-provisioning-key-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-provisioning-key-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-provisioning-key-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-reauth-policy-set-using-g-e-t-1"></span> Get the authentication policy and all rules for a Timeout policy rule. This API will be deprecated in a future release. (*getReauthPolicySetUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/reauth
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-reauth-policy-set-using-g-e-t-1-200) | OK | OK |  | [schema](#get-reauth-policy-set-using-g-e-t-1-200-schema) |
| [400](#get-reauth-policy-set-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-reauth-policy-set-using-g-e-t-1-400-schema) |
| [401](#get-reauth-policy-set-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-reauth-policy-set-using-g-e-t-1-401-schema) |
| [403](#get-reauth-policy-set-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-reauth-policy-set-using-g-e-t-1-403-schema) |
| [404](#get-reauth-policy-set-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-reauth-policy-set-using-g-e-t-1-404-schema) |
| [429](#get-reauth-policy-set-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-reauth-policy-set-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-reauth-policy-set-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-reauth-policy-set-using-g-e-t-1-200-schema"></span> Schema
   
  

[PolicySet](#policy-set)

##### <span id="get-reauth-policy-set-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-reauth-policy-set-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-reauth-policy-set-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-reauth-policy-set-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-reauth-policy-set-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-reauth-policy-set-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-reauth-policy-set-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-reauth-policy-set-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-reauth-policy-set-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-reauth-policy-set-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-rule-in-policy-set-using-g-e-t-1"></span> Get a rule in a policy. (*getRuleInPolicySetUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/policySet/{policySetId}/rule/{ruleId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| policySetId | `path` | string | `string` |  | ✓ |  | The unique identifier of the policy. |
| ruleId | `path` | string | `string` |  | ✓ |  | The unique identifier of a rule in a policy. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-rule-in-policy-set-using-g-e-t-1-200) | OK | OK |  | [schema](#get-rule-in-policy-set-using-g-e-t-1-200-schema) |
| [400](#get-rule-in-policy-set-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-rule-in-policy-set-using-g-e-t-1-400-schema) |
| [401](#get-rule-in-policy-set-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-rule-in-policy-set-using-g-e-t-1-401-schema) |
| [403](#get-rule-in-policy-set-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-rule-in-policy-set-using-g-e-t-1-403-schema) |
| [404](#get-rule-in-policy-set-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-rule-in-policy-set-using-g-e-t-1-404-schema) |
| [429](#get-rule-in-policy-set-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-rule-in-policy-set-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-rule-in-policy-set-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-rule-in-policy-set-using-g-e-t-1-200-schema"></span> Schema
   
  

[PolicyRule](#policy-rule)

##### <span id="get-rule-in-policy-set-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-rule-in-policy-set-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-rule-in-policy-set-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-rule-in-policy-set-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-rule-in-policy-set-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-rule-in-policy-set-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-rule-in-policy-set-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-rule-in-policy-set-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-rule-in-policy-set-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-rule-in-policy-set-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-s-c-i-m-attribute-using-g-e-t-1"></span> Get the SCIM attribute details. (*getSCIMAttributeUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/idp/{idpId}/scimattribute/{scimAttributeId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| idpId | `path` | string | `string` |  | ✓ |  | The unique identifier of the IdP. |
| scimAttributeId | `path` | string | `string` |  | ✓ |  | The unique identifier of the SCIM attribute. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-s-c-i-m-attribute-using-g-e-t-1-200) | OK | OK |  | [schema](#get-s-c-i-m-attribute-using-g-e-t-1-200-schema) |
| [400](#get-s-c-i-m-attribute-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-s-c-i-m-attribute-using-g-e-t-1-400-schema) |
| [401](#get-s-c-i-m-attribute-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-s-c-i-m-attribute-using-g-e-t-1-401-schema) |
| [403](#get-s-c-i-m-attribute-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-s-c-i-m-attribute-using-g-e-t-1-403-schema) |
| [404](#get-s-c-i-m-attribute-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-s-c-i-m-attribute-using-g-e-t-1-404-schema) |
| [429](#get-s-c-i-m-attribute-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-s-c-i-m-attribute-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-s-c-i-m-attribute-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-s-c-i-m-attribute-using-g-e-t-1-200-schema"></span> Schema
   
  

[SCIMAttributeHeader](#s-c-i-m-attribute-header)

##### <span id="get-s-c-i-m-attribute-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-s-c-i-m-attribute-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-s-c-i-m-attribute-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-s-c-i-m-attribute-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-s-c-i-m-attribute-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-s-c-i-m-attribute-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-s-c-i-m-attribute-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-s-c-i-m-attribute-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-s-c-i-m-attribute-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-s-c-i-m-attribute-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-saml-attribute-using-g-e-t-1"></span> Get the SAML attribute details. (*getSamlAttributeUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/samlAttribute/{attrId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| attrId | `path` | string | `string` |  | ✓ |  | The unique identifier of the SAML attribute. |
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-saml-attribute-using-g-e-t-1-200) | OK | OK |  | [schema](#get-saml-attribute-using-g-e-t-1-200-schema) |
| [400](#get-saml-attribute-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-saml-attribute-using-g-e-t-1-400-schema) |
| [401](#get-saml-attribute-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-saml-attribute-using-g-e-t-1-401-schema) |
| [403](#get-saml-attribute-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-saml-attribute-using-g-e-t-1-403-schema) |
| [404](#get-saml-attribute-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-saml-attribute-using-g-e-t-1-404-schema) |
| [429](#get-saml-attribute-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-saml-attribute-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-saml-attribute-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-saml-attribute-using-g-e-t-1-200-schema"></span> Schema
   
  

[SamlAttribute](#saml-attribute)

##### <span id="get-saml-attribute-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-saml-attribute-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-saml-attribute-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-saml-attribute-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-saml-attribute-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-saml-attribute-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-saml-attribute-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-saml-attribute-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-saml-attribute-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-saml-attribute-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-segment-group-using-g-e-t-1"></span> Get the Segment Group details. (*getSegmentGroupUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup/{segmentGroupId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| segmentGroupId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Segment Group. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-segment-group-using-g-e-t-1-200) | OK | OK |  | [schema](#get-segment-group-using-g-e-t-1-200-schema) |
| [400](#get-segment-group-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-segment-group-using-g-e-t-1-400-schema) |
| [401](#get-segment-group-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-segment-group-using-g-e-t-1-401-schema) |
| [403](#get-segment-group-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-segment-group-using-g-e-t-1-403-schema) |
| [404](#get-segment-group-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-segment-group-using-g-e-t-1-404-schema) |
| [429](#get-segment-group-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-segment-group-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-segment-group-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-segment-group-using-g-e-t-1-200-schema"></span> Schema
   
  

[SegmentGroup](#segment-group)

##### <span id="get-segment-group-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-segment-group-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-segment-group-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-segment-group-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-segment-group-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-segment-group-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-segment-group-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-segment-group-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-segment-group-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-segment-group-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-server-group-using-g-e-t-1"></span> Get the Server Group details. (*getServerGroupUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/serverGroup/{groupId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| groupId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Segment Group. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-server-group-using-g-e-t-1-200) | OK | OK |  | [schema](#get-server-group-using-g-e-t-1-200-schema) |
| [400](#get-server-group-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-server-group-using-g-e-t-1-400-schema) |
| [401](#get-server-group-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-server-group-using-g-e-t-1-401-schema) |
| [403](#get-server-group-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-server-group-using-g-e-t-1-403-schema) |
| [404](#get-server-group-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-server-group-using-g-e-t-1-404-schema) |
| [429](#get-server-group-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-server-group-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-server-group-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-server-group-using-g-e-t-1-200-schema"></span> Schema
   
  

[ServerGroupDTO](#server-group-d-t-o)

##### <span id="get-server-group-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-server-group-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-server-group-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-server-group-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-server-group-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-server-group-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-server-group-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-server-group-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-server-group-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-server-group-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-service-edge-group-using-g-e-t"></span> Get the Service Edge Group details. (*getServiceEdgeGroupUsingGET*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/serviceEdgeGroup/{serviceEdgeGroupId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| serviceEdgeGroupId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Service Edge Group. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-service-edge-group-using-g-e-t-200) | OK | OK |  | [schema](#get-service-edge-group-using-g-e-t-200-schema) |
| [400](#get-service-edge-group-using-g-e-t-400) | Bad Request | BadRequest |  | [schema](#get-service-edge-group-using-g-e-t-400-schema) |
| [401](#get-service-edge-group-using-g-e-t-401) | Unauthorized | Unauthorized |  | [schema](#get-service-edge-group-using-g-e-t-401-schema) |
| [403](#get-service-edge-group-using-g-e-t-403) | Forbidden | Forbidden |  | [schema](#get-service-edge-group-using-g-e-t-403-schema) |
| [404](#get-service-edge-group-using-g-e-t-404) | Not Found | Not Found |  | [schema](#get-service-edge-group-using-g-e-t-404-schema) |
| [429](#get-service-edge-group-using-g-e-t-429) | Too Many Requests | TooManyRequest |  | [schema](#get-service-edge-group-using-g-e-t-429-schema) |

#### Responses


##### <span id="get-service-edge-group-using-g-e-t-200"></span> 200 - OK
Status: OK

###### <span id="get-service-edge-group-using-g-e-t-200-schema"></span> Schema
   
  

[ServiceEdgeGroup](#service-edge-group)

##### <span id="get-service-edge-group-using-g-e-t-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-service-edge-group-using-g-e-t-400-schema"></span> Schema

##### <span id="get-service-edge-group-using-g-e-t-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-service-edge-group-using-g-e-t-401-schema"></span> Schema

##### <span id="get-service-edge-group-using-g-e-t-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-service-edge-group-using-g-e-t-403-schema"></span> Schema

##### <span id="get-service-edge-group-using-g-e-t-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-service-edge-group-using-g-e-t-404-schema"></span> Schema

##### <span id="get-service-edge-group-using-g-e-t-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-service-edge-group-using-g-e-t-429-schema"></span> Schema

### <span id="get-service-edge-groups-using-g-e-t"></span> Get all the configured Service Edge Group details. (*getServiceEdgeGroupsUsingGET*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/serviceEdgeGroup
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| page | `query` | string | `string` |  |  |  | Specifies the page number. |
| pagesize | `query` | string | `string` |  |  |  | Specifies the page size. If not provided, the default page size is 20. The max page size is 500. |
| search | `query` | string | `string` |  |  |  | The search string used to support search by features and fields for the API. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-service-edge-groups-using-g-e-t-200) | OK | OK |  | [schema](#get-service-edge-groups-using-g-e-t-200-schema) |
| [400](#get-service-edge-groups-using-g-e-t-400) | Bad Request | BadRequest |  | [schema](#get-service-edge-groups-using-g-e-t-400-schema) |
| [401](#get-service-edge-groups-using-g-e-t-401) | Unauthorized | Unauthorized |  | [schema](#get-service-edge-groups-using-g-e-t-401-schema) |
| [403](#get-service-edge-groups-using-g-e-t-403) | Forbidden | Forbidden |  | [schema](#get-service-edge-groups-using-g-e-t-403-schema) |
| [404](#get-service-edge-groups-using-g-e-t-404) | Not Found | Not Found |  | [schema](#get-service-edge-groups-using-g-e-t-404-schema) |
| [429](#get-service-edge-groups-using-g-e-t-429) | Too Many Requests | TooManyRequest |  | [schema](#get-service-edge-groups-using-g-e-t-429-schema) |

#### Responses


##### <span id="get-service-edge-groups-using-g-e-t-200"></span> 200 - OK
Status: OK

###### <span id="get-service-edge-groups-using-g-e-t-200-schema"></span> Schema
   
  

[PageListOfServiceEdgeGroup](#page-list-of-service-edge-group)

##### <span id="get-service-edge-groups-using-g-e-t-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-service-edge-groups-using-g-e-t-400-schema"></span> Schema

##### <span id="get-service-edge-groups-using-g-e-t-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-service-edge-groups-using-g-e-t-401-schema"></span> Schema

##### <span id="get-service-edge-groups-using-g-e-t-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-service-edge-groups-using-g-e-t-403-schema"></span> Schema

##### <span id="get-service-edge-groups-using-g-e-t-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-service-edge-groups-using-g-e-t-404-schema"></span> Schema

##### <span id="get-service-edge-groups-using-g-e-t-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-service-edge-groups-using-g-e-t-429-schema"></span> Schema

### <span id="get-service-edge-using-g-e-t-1"></span> Get the Service Edge details. (*getServiceEdgeUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/{serviceEdgeId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| serviceEdgeId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Service Edge. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-service-edge-using-g-e-t-1-200) | OK | OK |  | [schema](#get-service-edge-using-g-e-t-1-200-schema) |
| [400](#get-service-edge-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-service-edge-using-g-e-t-1-400-schema) |
| [401](#get-service-edge-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-service-edge-using-g-e-t-1-401-schema) |
| [403](#get-service-edge-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-service-edge-using-g-e-t-1-403-schema) |
| [404](#get-service-edge-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-service-edge-using-g-e-t-1-404-schema) |
| [429](#get-service-edge-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-service-edge-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-service-edge-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-service-edge-using-g-e-t-1-200-schema"></span> Schema
   
  

[ServiceEdge](#service-edge)

##### <span id="get-service-edge-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-service-edge-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-service-edge-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-service-edge-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-service-edge-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-service-edge-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-service-edge-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-service-edge-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-service-edge-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-service-edge-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-status-codes-using-g-e-t-1"></span> Get list of LSS status codes. (*getStatusCodesUsingGET_1*)

```
GET /mgmtconfig/v2/admin/lssConfig/statusCodes
```

#### Produces
  * */*

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-status-codes-using-g-e-t-1-200) | OK | OK |  | [schema](#get-status-codes-using-g-e-t-1-200-schema) |
| [400](#get-status-codes-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-status-codes-using-g-e-t-1-400-schema) |
| [401](#get-status-codes-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-status-codes-using-g-e-t-1-401-schema) |
| [403](#get-status-codes-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-status-codes-using-g-e-t-1-403-schema) |
| [404](#get-status-codes-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-status-codes-using-g-e-t-1-404-schema) |
| [429](#get-status-codes-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-status-codes-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-status-codes-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-status-codes-using-g-e-t-1-200-schema"></span> Schema
   
  

map of any 

##### <span id="get-status-codes-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-status-codes-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-status-codes-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-status-codes-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-status-codes-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-status-codes-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-status-codes-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-status-codes-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-status-codes-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-status-codes-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-trusted-network-using-g-e-t-1"></span> Get the Trusted Networks. (*getTrustedNetworkUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/network/{id}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| id | `path` | string | `string` |  | ✓ |  | The unique identifier of the Trusted Network. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-trusted-network-using-g-e-t-1-200) | OK | OK |  | [schema](#get-trusted-network-using-g-e-t-1-200-schema) |
| [400](#get-trusted-network-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-trusted-network-using-g-e-t-1-400-schema) |
| [401](#get-trusted-network-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-trusted-network-using-g-e-t-1-401-schema) |
| [403](#get-trusted-network-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-trusted-network-using-g-e-t-1-403-schema) |
| [404](#get-trusted-network-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-trusted-network-using-g-e-t-1-404-schema) |
| [429](#get-trusted-network-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-trusted-network-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-trusted-network-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-trusted-network-using-g-e-t-1-200-schema"></span> Schema
   
  

[TrustedNetwork](#trusted-network)

##### <span id="get-trusted-network-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-trusted-network-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-trusted-network-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-trusted-network-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-trusted-network-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-trusted-network-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-trusted-network-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-trusted-network-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-trusted-network-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-trusted-network-using-g-e-t-1-429-schema"></span> Schema

### <span id="get-using-g-e-t-1"></span> Get the details of the Browser Access certificate. (*getUsingGET_1*)

```
GET /mgmtconfig/v1/admin/customers/{customerId}/clientlessCertificate/{certificateId}
```

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| certificateId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Browser Access certificate. |
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-using-g-e-t-1-200) | OK | OK |  | [schema](#get-using-g-e-t-1-200-schema) |
| [400](#get-using-g-e-t-1-400) | Bad Request | BadRequest |  | [schema](#get-using-g-e-t-1-400-schema) |
| [401](#get-using-g-e-t-1-401) | Unauthorized | Unauthorized |  | [schema](#get-using-g-e-t-1-401-schema) |
| [403](#get-using-g-e-t-1-403) | Forbidden | Forbidden |  | [schema](#get-using-g-e-t-1-403-schema) |
| [404](#get-using-g-e-t-1-404) | Not Found | Not Found |  | [schema](#get-using-g-e-t-1-404-schema) |
| [429](#get-using-g-e-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#get-using-g-e-t-1-429-schema) |

#### Responses


##### <span id="get-using-g-e-t-1-200"></span> 200 - OK
Status: OK

###### <span id="get-using-g-e-t-1-200-schema"></span> Schema
   
  

[BACertificate](#b-a-certificate)

##### <span id="get-using-g-e-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="get-using-g-e-t-1-400-schema"></span> Schema

##### <span id="get-using-g-e-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="get-using-g-e-t-1-401-schema"></span> Schema

##### <span id="get-using-g-e-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="get-using-g-e-t-1-403-schema"></span> Schema

##### <span id="get-using-g-e-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="get-using-g-e-t-1-404-schema"></span> Schema

##### <span id="get-using-g-e-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="get-using-g-e-t-1-429-schema"></span> Schema

### <span id="re-order-policy-rule-using-p-u-t-1"></span> Update rule order. (*reOrderPolicyRuleUsingPUT_1*)

```
PUT /mgmtconfig/v1/admin/customers/{customerId}/policySet/{policySetId}/rule/{ruleId}/reorder/{newOrder}
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| newOrder | `path` | string | `string` |  | ✓ |  | The new order of the rule. |
| policySetId | `path` | string | `string` |  | ✓ |  | The unique identifier of the policy. |
| ruleId | `path` | string | `string` |  | ✓ |  | The unique identifier of a rule in a policy. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#re-order-policy-rule-using-p-u-t-1-201) | Created | Created |  | [schema](#re-order-policy-rule-using-p-u-t-1-201-schema) |
| [204](#re-order-policy-rule-using-p-u-t-1-204) | No Content | No Content |  | [schema](#re-order-policy-rule-using-p-u-t-1-204-schema) |
| [400](#re-order-policy-rule-using-p-u-t-1-400) | Bad Request | BadRequest |  | [schema](#re-order-policy-rule-using-p-u-t-1-400-schema) |
| [401](#re-order-policy-rule-using-p-u-t-1-401) | Unauthorized | Unauthorized |  | [schema](#re-order-policy-rule-using-p-u-t-1-401-schema) |
| [403](#re-order-policy-rule-using-p-u-t-1-403) | Forbidden | Forbidden |  | [schema](#re-order-policy-rule-using-p-u-t-1-403-schema) |
| [404](#re-order-policy-rule-using-p-u-t-1-404) | Not Found | Not Found |  | [schema](#re-order-policy-rule-using-p-u-t-1-404-schema) |
| [429](#re-order-policy-rule-using-p-u-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#re-order-policy-rule-using-p-u-t-1-429-schema) |

#### Responses


##### <span id="re-order-policy-rule-using-p-u-t-1-201"></span> 201 - Created
Status: Created

###### <span id="re-order-policy-rule-using-p-u-t-1-201-schema"></span> Schema

##### <span id="re-order-policy-rule-using-p-u-t-1-204"></span> 204 - No Content
Status: No Content

###### <span id="re-order-policy-rule-using-p-u-t-1-204-schema"></span> Schema

##### <span id="re-order-policy-rule-using-p-u-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="re-order-policy-rule-using-p-u-t-1-400-schema"></span> Schema

##### <span id="re-order-policy-rule-using-p-u-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="re-order-policy-rule-using-p-u-t-1-401-schema"></span> Schema

##### <span id="re-order-policy-rule-using-p-u-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="re-order-policy-rule-using-p-u-t-1-403-schema"></span> Schema

##### <span id="re-order-policy-rule-using-p-u-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="re-order-policy-rule-using-p-u-t-1-404-schema"></span> Schema

##### <span id="re-order-policy-rule-using-p-u-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="re-order-policy-rule-using-p-u-t-1-429-schema"></span> Schema

### <span id="update-app-connector-group-using-p-u-t-1"></span> Update the Connector Group details. (*updateAppConnectorGroupUsingPUT_1*)

```
PUT /mgmtconfig/v1/admin/customers/{customerId}/appConnectorGroup/{appConnectorGroupId}
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| appConnectorGroupId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Connector Group. |
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| appConnectorGroup | `body` | [AppConnectorGroup](#app-connector-group) | `models.AppConnectorGroup` | | ✓ | | The object of the Connector Group. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#update-app-connector-group-using-p-u-t-1-201) | Created | Created |  | [schema](#update-app-connector-group-using-p-u-t-1-201-schema) |
| [204](#update-app-connector-group-using-p-u-t-1-204) | No Content | No Content |  | [schema](#update-app-connector-group-using-p-u-t-1-204-schema) |
| [400](#update-app-connector-group-using-p-u-t-1-400) | Bad Request | BadRequest |  | [schema](#update-app-connector-group-using-p-u-t-1-400-schema) |
| [401](#update-app-connector-group-using-p-u-t-1-401) | Unauthorized | Unauthorized |  | [schema](#update-app-connector-group-using-p-u-t-1-401-schema) |
| [403](#update-app-connector-group-using-p-u-t-1-403) | Forbidden | Forbidden |  | [schema](#update-app-connector-group-using-p-u-t-1-403-schema) |
| [404](#update-app-connector-group-using-p-u-t-1-404) | Not Found | Not Found |  | [schema](#update-app-connector-group-using-p-u-t-1-404-schema) |
| [429](#update-app-connector-group-using-p-u-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#update-app-connector-group-using-p-u-t-1-429-schema) |

#### Responses


##### <span id="update-app-connector-group-using-p-u-t-1-201"></span> 201 - Created
Status: Created

###### <span id="update-app-connector-group-using-p-u-t-1-201-schema"></span> Schema

##### <span id="update-app-connector-group-using-p-u-t-1-204"></span> 204 - No Content
Status: No Content

###### <span id="update-app-connector-group-using-p-u-t-1-204-schema"></span> Schema

##### <span id="update-app-connector-group-using-p-u-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="update-app-connector-group-using-p-u-t-1-400-schema"></span> Schema

##### <span id="update-app-connector-group-using-p-u-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="update-app-connector-group-using-p-u-t-1-401-schema"></span> Schema

##### <span id="update-app-connector-group-using-p-u-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="update-app-connector-group-using-p-u-t-1-403-schema"></span> Schema

##### <span id="update-app-connector-group-using-p-u-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="update-app-connector-group-using-p-u-t-1-404-schema"></span> Schema

##### <span id="update-app-connector-group-using-p-u-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="update-app-connector-group-using-p-u-t-1-429-schema"></span> Schema

### <span id="update-app-server-group-using-p-u-t-1"></span> Update the Server Group. (*updateAppServerGroupUsingPUT_1*)

```
PUT /mgmtconfig/v1/admin/customers/{customerId}/serverGroup/{groupId}
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| groupId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Server Group. |
| group | `body` | [ServerGroupDTO](#server-group-d-t-o) | `models.ServerGroupDTO` | |  | | The object of the Server Group. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#update-app-server-group-using-p-u-t-1-201) | Created | Created |  | [schema](#update-app-server-group-using-p-u-t-1-201-schema) |
| [204](#update-app-server-group-using-p-u-t-1-204) | No Content | No Content |  | [schema](#update-app-server-group-using-p-u-t-1-204-schema) |
| [400](#update-app-server-group-using-p-u-t-1-400) | Bad Request | BadRequest |  | [schema](#update-app-server-group-using-p-u-t-1-400-schema) |
| [401](#update-app-server-group-using-p-u-t-1-401) | Unauthorized | Unauthorized |  | [schema](#update-app-server-group-using-p-u-t-1-401-schema) |
| [403](#update-app-server-group-using-p-u-t-1-403) | Forbidden | Forbidden |  | [schema](#update-app-server-group-using-p-u-t-1-403-schema) |
| [404](#update-app-server-group-using-p-u-t-1-404) | Not Found | Not Found |  | [schema](#update-app-server-group-using-p-u-t-1-404-schema) |
| [429](#update-app-server-group-using-p-u-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#update-app-server-group-using-p-u-t-1-429-schema) |

#### Responses


##### <span id="update-app-server-group-using-p-u-t-1-201"></span> 201 - Created
Status: Created

###### <span id="update-app-server-group-using-p-u-t-1-201-schema"></span> Schema

##### <span id="update-app-server-group-using-p-u-t-1-204"></span> 204 - No Content
Status: No Content

###### <span id="update-app-server-group-using-p-u-t-1-204-schema"></span> Schema

##### <span id="update-app-server-group-using-p-u-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="update-app-server-group-using-p-u-t-1-400-schema"></span> Schema

##### <span id="update-app-server-group-using-p-u-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="update-app-server-group-using-p-u-t-1-401-schema"></span> Schema

##### <span id="update-app-server-group-using-p-u-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="update-app-server-group-using-p-u-t-1-403-schema"></span> Schema

##### <span id="update-app-server-group-using-p-u-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="update-app-server-group-using-p-u-t-1-404-schema"></span> Schema

##### <span id="update-app-server-group-using-p-u-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="update-app-server-group-using-p-u-t-1-429-schema"></span> Schema

### <span id="update-app-server-using-p-u-t-1"></span> Update the Server details. (*updateAppServerUsingPUT_1*)

```
PUT /mgmtconfig/v1/admin/customers/{customerId}/server/{serverId}
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| serverId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Server. |
| server | `body` | [ApplicationServer](#application-server) | `models.ApplicationServer` | |  | | The object of the Server. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#update-app-server-using-p-u-t-1-201) | Created | Created |  | [schema](#update-app-server-using-p-u-t-1-201-schema) |
| [204](#update-app-server-using-p-u-t-1-204) | No Content | No Content |  | [schema](#update-app-server-using-p-u-t-1-204-schema) |
| [400](#update-app-server-using-p-u-t-1-400) | Bad Request | BadRequest |  | [schema](#update-app-server-using-p-u-t-1-400-schema) |
| [401](#update-app-server-using-p-u-t-1-401) | Unauthorized | Unauthorized |  | [schema](#update-app-server-using-p-u-t-1-401-schema) |
| [403](#update-app-server-using-p-u-t-1-403) | Forbidden | Forbidden |  | [schema](#update-app-server-using-p-u-t-1-403-schema) |
| [404](#update-app-server-using-p-u-t-1-404) | Not Found | Not Found |  | [schema](#update-app-server-using-p-u-t-1-404-schema) |
| [429](#update-app-server-using-p-u-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#update-app-server-using-p-u-t-1-429-schema) |

#### Responses


##### <span id="update-app-server-using-p-u-t-1-201"></span> 201 - Created
Status: Created

###### <span id="update-app-server-using-p-u-t-1-201-schema"></span> Schema

##### <span id="update-app-server-using-p-u-t-1-204"></span> 204 - No Content
Status: No Content

###### <span id="update-app-server-using-p-u-t-1-204-schema"></span> Schema

##### <span id="update-app-server-using-p-u-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="update-app-server-using-p-u-t-1-400-schema"></span> Schema

##### <span id="update-app-server-using-p-u-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="update-app-server-using-p-u-t-1-401-schema"></span> Schema

##### <span id="update-app-server-using-p-u-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="update-app-server-using-p-u-t-1-403-schema"></span> Schema

##### <span id="update-app-server-using-p-u-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="update-app-server-using-p-u-t-1-404-schema"></span> Schema

##### <span id="update-app-server-using-p-u-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="update-app-server-using-p-u-t-1-429-schema"></span> Schema

### <span id="update-application-v2-using-p-u-t-1"></span> Update the Application Segment details. (*updateApplicationV2UsingPUT_1*)

```
PUT /mgmtconfig/v1/admin/customers/{customerId}/application/{applicationId}
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| applicationId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Application Segment. |
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| application | `body` | [ApplicationResource](#application-resource) | `models.ApplicationResource` | |  | | The object of the Application Segment. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#update-application-v2-using-p-u-t-1-201) | Created | Created |  | [schema](#update-application-v2-using-p-u-t-1-201-schema) |
| [204](#update-application-v2-using-p-u-t-1-204) | No Content | No Content |  | [schema](#update-application-v2-using-p-u-t-1-204-schema) |
| [400](#update-application-v2-using-p-u-t-1-400) | Bad Request | BadRequest |  | [schema](#update-application-v2-using-p-u-t-1-400-schema) |
| [401](#update-application-v2-using-p-u-t-1-401) | Unauthorized | Unauthorized |  | [schema](#update-application-v2-using-p-u-t-1-401-schema) |
| [403](#update-application-v2-using-p-u-t-1-403) | Forbidden | Forbidden |  | [schema](#update-application-v2-using-p-u-t-1-403-schema) |
| [404](#update-application-v2-using-p-u-t-1-404) | Not Found | Not Found |  | [schema](#update-application-v2-using-p-u-t-1-404-schema) |
| [429](#update-application-v2-using-p-u-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#update-application-v2-using-p-u-t-1-429-schema) |

#### Responses


##### <span id="update-application-v2-using-p-u-t-1-201"></span> 201 - Created
Status: Created

###### <span id="update-application-v2-using-p-u-t-1-201-schema"></span> Schema

##### <span id="update-application-v2-using-p-u-t-1-204"></span> 204 - No Content
Status: No Content

###### <span id="update-application-v2-using-p-u-t-1-204-schema"></span> Schema

##### <span id="update-application-v2-using-p-u-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="update-application-v2-using-p-u-t-1-400-schema"></span> Schema

##### <span id="update-application-v2-using-p-u-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="update-application-v2-using-p-u-t-1-401-schema"></span> Schema

##### <span id="update-application-v2-using-p-u-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="update-application-v2-using-p-u-t-1-403-schema"></span> Schema

##### <span id="update-application-v2-using-p-u-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="update-application-v2-using-p-u-t-1-404-schema"></span> Schema

##### <span id="update-application-v2-using-p-u-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="update-application-v2-using-p-u-t-1-429-schema"></span> Schema

### <span id="update-connector-using-p-u-t-1"></span> Update the Connector details. (*updateConnectorUsingPUT_1*)

```
PUT /mgmtconfig/v1/admin/customers/{customerId}/connector/{connectorId}
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| connectorId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Connector. |
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| connector | `body` | [Connector](#connector) | `models.Connector` | | ✓ | | The object of the Connector. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#update-connector-using-p-u-t-1-201) | Created | Created |  | [schema](#update-connector-using-p-u-t-1-201-schema) |
| [204](#update-connector-using-p-u-t-1-204) | No Content | No Content |  | [schema](#update-connector-using-p-u-t-1-204-schema) |
| [400](#update-connector-using-p-u-t-1-400) | Bad Request | BadRequest |  | [schema](#update-connector-using-p-u-t-1-400-schema) |
| [401](#update-connector-using-p-u-t-1-401) | Unauthorized | Unauthorized |  | [schema](#update-connector-using-p-u-t-1-401-schema) |
| [403](#update-connector-using-p-u-t-1-403) | Forbidden | Forbidden |  | [schema](#update-connector-using-p-u-t-1-403-schema) |
| [404](#update-connector-using-p-u-t-1-404) | Not Found | Not Found |  | [schema](#update-connector-using-p-u-t-1-404-schema) |
| [429](#update-connector-using-p-u-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#update-connector-using-p-u-t-1-429-schema) |

#### Responses


##### <span id="update-connector-using-p-u-t-1-201"></span> 201 - Created
Status: Created

###### <span id="update-connector-using-p-u-t-1-201-schema"></span> Schema

##### <span id="update-connector-using-p-u-t-1-204"></span> 204 - No Content
Status: No Content

###### <span id="update-connector-using-p-u-t-1-204-schema"></span> Schema

##### <span id="update-connector-using-p-u-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="update-connector-using-p-u-t-1-400-schema"></span> Schema

##### <span id="update-connector-using-p-u-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="update-connector-using-p-u-t-1-401-schema"></span> Schema

##### <span id="update-connector-using-p-u-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="update-connector-using-p-u-t-1-403-schema"></span> Schema

##### <span id="update-connector-using-p-u-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="update-connector-using-p-u-t-1-404-schema"></span> Schema

##### <span id="update-connector-using-p-u-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="update-connector-using-p-u-t-1-429-schema"></span> Schema

### <span id="update-lss-using-p-u-t-3"></span> Update a LSS Attribute. (*updateLssUsingPUT_3*)

```
PUT /mgmtconfig/v2/admin/customers/{customerId}/lssConfig/{lssId}
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| lssId | `path` | string | `string` |  | ✓ |  | The unique identifier of the LSS resource. |
| lssResource | `body` | [LssResource](#lss-resource) | `models.LssResource` | |  | | The object of the LSS Resource. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#update-lss-using-p-u-t-3-201) | Created | Created |  | [schema](#update-lss-using-p-u-t-3-201-schema) |
| [204](#update-lss-using-p-u-t-3-204) | No Content | No Content |  | [schema](#update-lss-using-p-u-t-3-204-schema) |
| [400](#update-lss-using-p-u-t-3-400) | Bad Request | BadRequest |  | [schema](#update-lss-using-p-u-t-3-400-schema) |
| [401](#update-lss-using-p-u-t-3-401) | Unauthorized | Unauthorized |  | [schema](#update-lss-using-p-u-t-3-401-schema) |
| [403](#update-lss-using-p-u-t-3-403) | Forbidden | Forbidden |  | [schema](#update-lss-using-p-u-t-3-403-schema) |
| [404](#update-lss-using-p-u-t-3-404) | Not Found | Not Found |  | [schema](#update-lss-using-p-u-t-3-404-schema) |
| [429](#update-lss-using-p-u-t-3-429) | Too Many Requests | TooManyRequest |  | [schema](#update-lss-using-p-u-t-3-429-schema) |

#### Responses


##### <span id="update-lss-using-p-u-t-3-201"></span> 201 - Created
Status: Created

###### <span id="update-lss-using-p-u-t-3-201-schema"></span> Schema

##### <span id="update-lss-using-p-u-t-3-204"></span> 204 - No Content
Status: No Content

###### <span id="update-lss-using-p-u-t-3-204-schema"></span> Schema

##### <span id="update-lss-using-p-u-t-3-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="update-lss-using-p-u-t-3-400-schema"></span> Schema

##### <span id="update-lss-using-p-u-t-3-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="update-lss-using-p-u-t-3-401-schema"></span> Schema

##### <span id="update-lss-using-p-u-t-3-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="update-lss-using-p-u-t-3-403-schema"></span> Schema

##### <span id="update-lss-using-p-u-t-3-404"></span> 404 - Not Found
Status: Not Found

###### <span id="update-lss-using-p-u-t-3-404-schema"></span> Schema

##### <span id="update-lss-using-p-u-t-3-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="update-lss-using-p-u-t-3-429-schema"></span> Schema

### <span id="update-provisioning-key-using-p-u-t-1"></span> Update the Provisioning Key details. (*updateProvisioningKeyUsingPUT_1*)

```
PUT /mgmtconfig/v1/admin/customers/{customerId}/associationType/{associationType}/provisioningKey/{provisioningKeyId}
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| associationType | `path` | string | `string` |  | ✓ |  | Specifies the provisioning key type for App Connectors or ZPA Private Service Edges. The supported values are CONNECTOR_GRP and SERVICE_EDGE_GRP. |
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| provisioningKeyId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Provisioning Key. |
| provisioningKey | `body` | [ProvisioningKey](#provisioning-key) | `models.ProvisioningKey` | | ✓ | | The object of the Provisioning Key. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#update-provisioning-key-using-p-u-t-1-201) | Created | Created |  | [schema](#update-provisioning-key-using-p-u-t-1-201-schema) |
| [204](#update-provisioning-key-using-p-u-t-1-204) | No Content | No Content |  | [schema](#update-provisioning-key-using-p-u-t-1-204-schema) |
| [400](#update-provisioning-key-using-p-u-t-1-400) | Bad Request | BadRequest |  | [schema](#update-provisioning-key-using-p-u-t-1-400-schema) |
| [401](#update-provisioning-key-using-p-u-t-1-401) | Unauthorized | Unauthorized |  | [schema](#update-provisioning-key-using-p-u-t-1-401-schema) |
| [403](#update-provisioning-key-using-p-u-t-1-403) | Forbidden | Forbidden |  | [schema](#update-provisioning-key-using-p-u-t-1-403-schema) |
| [404](#update-provisioning-key-using-p-u-t-1-404) | Not Found | Not Found |  | [schema](#update-provisioning-key-using-p-u-t-1-404-schema) |
| [429](#update-provisioning-key-using-p-u-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#update-provisioning-key-using-p-u-t-1-429-schema) |

#### Responses


##### <span id="update-provisioning-key-using-p-u-t-1-201"></span> 201 - Created
Status: Created

###### <span id="update-provisioning-key-using-p-u-t-1-201-schema"></span> Schema

##### <span id="update-provisioning-key-using-p-u-t-1-204"></span> 204 - No Content
Status: No Content

###### <span id="update-provisioning-key-using-p-u-t-1-204-schema"></span> Schema

##### <span id="update-provisioning-key-using-p-u-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="update-provisioning-key-using-p-u-t-1-400-schema"></span> Schema

##### <span id="update-provisioning-key-using-p-u-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="update-provisioning-key-using-p-u-t-1-401-schema"></span> Schema

##### <span id="update-provisioning-key-using-p-u-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="update-provisioning-key-using-p-u-t-1-403-schema"></span> Schema

##### <span id="update-provisioning-key-using-p-u-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="update-provisioning-key-using-p-u-t-1-404-schema"></span> Schema

##### <span id="update-provisioning-key-using-p-u-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="update-provisioning-key-using-p-u-t-1-429-schema"></span> Schema

### <span id="update-rule-to-policy-set-using-p-u-t-1"></span> Update a rule in a policy. (*updateRuleToPolicySetUsingPUT_1*)

```
PUT /mgmtconfig/v1/admin/customers/{customerId}/policySet/{policySetId}/rule/{ruleId}
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| policySetId | `path` | string | `string` |  | ✓ |  | The unique identifier of the policy. |
| ruleId | `path` | string | `string` |  | ✓ |  | The unique identifier of a rule in a policy. |
| rule | `body` | [PolicyRule](#policy-rule) | `models.PolicyRule` | | ✓ | | The object of the rule in a policy. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#update-rule-to-policy-set-using-p-u-t-1-201) | Created | Created |  | [schema](#update-rule-to-policy-set-using-p-u-t-1-201-schema) |
| [204](#update-rule-to-policy-set-using-p-u-t-1-204) | No Content | No Content |  | [schema](#update-rule-to-policy-set-using-p-u-t-1-204-schema) |
| [400](#update-rule-to-policy-set-using-p-u-t-1-400) | Bad Request | BadRequest |  | [schema](#update-rule-to-policy-set-using-p-u-t-1-400-schema) |
| [401](#update-rule-to-policy-set-using-p-u-t-1-401) | Unauthorized | Unauthorized |  | [schema](#update-rule-to-policy-set-using-p-u-t-1-401-schema) |
| [403](#update-rule-to-policy-set-using-p-u-t-1-403) | Forbidden | Forbidden |  | [schema](#update-rule-to-policy-set-using-p-u-t-1-403-schema) |
| [404](#update-rule-to-policy-set-using-p-u-t-1-404) | Not Found | Not Found |  | [schema](#update-rule-to-policy-set-using-p-u-t-1-404-schema) |
| [429](#update-rule-to-policy-set-using-p-u-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#update-rule-to-policy-set-using-p-u-t-1-429-schema) |

#### Responses


##### <span id="update-rule-to-policy-set-using-p-u-t-1-201"></span> 201 - Created
Status: Created

###### <span id="update-rule-to-policy-set-using-p-u-t-1-201-schema"></span> Schema

##### <span id="update-rule-to-policy-set-using-p-u-t-1-204"></span> 204 - No Content
Status: No Content

###### <span id="update-rule-to-policy-set-using-p-u-t-1-204-schema"></span> Schema

##### <span id="update-rule-to-policy-set-using-p-u-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="update-rule-to-policy-set-using-p-u-t-1-400-schema"></span> Schema

##### <span id="update-rule-to-policy-set-using-p-u-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="update-rule-to-policy-set-using-p-u-t-1-401-schema"></span> Schema

##### <span id="update-rule-to-policy-set-using-p-u-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="update-rule-to-policy-set-using-p-u-t-1-403-schema"></span> Schema

##### <span id="update-rule-to-policy-set-using-p-u-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="update-rule-to-policy-set-using-p-u-t-1-404-schema"></span> Schema

##### <span id="update-rule-to-policy-set-using-p-u-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="update-rule-to-policy-set-using-p-u-t-1-429-schema"></span> Schema

### <span id="update-segment-group-using-p-u-t-1"></span> Update a Segment Group. (*updateSegmentGroupUsingPUT_1*)

```
PUT /mgmtconfig/v1/admin/customers/{customerId}/segmentGroup/{segmentGroupId}
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| segmentGroupId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Segment Group. |
| segmentGroup | `body` | [SegmentGroup](#segment-group) | `models.SegmentGroup` | |  | | The object of the Segment Group. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#update-segment-group-using-p-u-t-1-201) | Created | Created |  | [schema](#update-segment-group-using-p-u-t-1-201-schema) |
| [204](#update-segment-group-using-p-u-t-1-204) | No Content | No Content |  | [schema](#update-segment-group-using-p-u-t-1-204-schema) |
| [400](#update-segment-group-using-p-u-t-1-400) | Bad Request | BadRequest |  | [schema](#update-segment-group-using-p-u-t-1-400-schema) |
| [401](#update-segment-group-using-p-u-t-1-401) | Unauthorized | Unauthorized |  | [schema](#update-segment-group-using-p-u-t-1-401-schema) |
| [403](#update-segment-group-using-p-u-t-1-403) | Forbidden | Forbidden |  | [schema](#update-segment-group-using-p-u-t-1-403-schema) |
| [404](#update-segment-group-using-p-u-t-1-404) | Not Found | Not Found |  | [schema](#update-segment-group-using-p-u-t-1-404-schema) |
| [429](#update-segment-group-using-p-u-t-1-429) | Too Many Requests | TooManyRequest |  | [schema](#update-segment-group-using-p-u-t-1-429-schema) |

#### Responses


##### <span id="update-segment-group-using-p-u-t-1-201"></span> 201 - Created
Status: Created

###### <span id="update-segment-group-using-p-u-t-1-201-schema"></span> Schema

##### <span id="update-segment-group-using-p-u-t-1-204"></span> 204 - No Content
Status: No Content

###### <span id="update-segment-group-using-p-u-t-1-204-schema"></span> Schema

##### <span id="update-segment-group-using-p-u-t-1-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="update-segment-group-using-p-u-t-1-400-schema"></span> Schema

##### <span id="update-segment-group-using-p-u-t-1-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="update-segment-group-using-p-u-t-1-401-schema"></span> Schema

##### <span id="update-segment-group-using-p-u-t-1-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="update-segment-group-using-p-u-t-1-403-schema"></span> Schema

##### <span id="update-segment-group-using-p-u-t-1-404"></span> 404 - Not Found
Status: Not Found

###### <span id="update-segment-group-using-p-u-t-1-404-schema"></span> Schema

##### <span id="update-segment-group-using-p-u-t-1-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="update-segment-group-using-p-u-t-1-429-schema"></span> Schema

### <span id="update-service-edge-group-using-p-u-t"></span> Update the Service Edge Group details. (*updateServiceEdgeGroupUsingPUT*)

```
PUT /mgmtconfig/v1/admin/customers/{customerId}/serviceEdgeGroup/{serviceEdgeGroupId}
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| serviceEdgeGroupId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Service Edge Group. |
| serviceEdgeGroup | `body` | [ServiceEdgeGroup](#service-edge-group) | `models.ServiceEdgeGroup` | | ✓ | | The object of the Service Edge Group. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#update-service-edge-group-using-p-u-t-201) | Created | Created |  | [schema](#update-service-edge-group-using-p-u-t-201-schema) |
| [204](#update-service-edge-group-using-p-u-t-204) | No Content | No Content |  | [schema](#update-service-edge-group-using-p-u-t-204-schema) |
| [400](#update-service-edge-group-using-p-u-t-400) | Bad Request | BadRequest |  | [schema](#update-service-edge-group-using-p-u-t-400-schema) |
| [401](#update-service-edge-group-using-p-u-t-401) | Unauthorized | Unauthorized |  | [schema](#update-service-edge-group-using-p-u-t-401-schema) |
| [403](#update-service-edge-group-using-p-u-t-403) | Forbidden | Forbidden |  | [schema](#update-service-edge-group-using-p-u-t-403-schema) |
| [404](#update-service-edge-group-using-p-u-t-404) | Not Found | Not Found |  | [schema](#update-service-edge-group-using-p-u-t-404-schema) |
| [429](#update-service-edge-group-using-p-u-t-429) | Too Many Requests | TooManyRequest |  | [schema](#update-service-edge-group-using-p-u-t-429-schema) |

#### Responses


##### <span id="update-service-edge-group-using-p-u-t-201"></span> 201 - Created
Status: Created

###### <span id="update-service-edge-group-using-p-u-t-201-schema"></span> Schema

##### <span id="update-service-edge-group-using-p-u-t-204"></span> 204 - No Content
Status: No Content

###### <span id="update-service-edge-group-using-p-u-t-204-schema"></span> Schema

##### <span id="update-service-edge-group-using-p-u-t-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="update-service-edge-group-using-p-u-t-400-schema"></span> Schema

##### <span id="update-service-edge-group-using-p-u-t-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="update-service-edge-group-using-p-u-t-401-schema"></span> Schema

##### <span id="update-service-edge-group-using-p-u-t-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="update-service-edge-group-using-p-u-t-403-schema"></span> Schema

##### <span id="update-service-edge-group-using-p-u-t-404"></span> 404 - Not Found
Status: Not Found

###### <span id="update-service-edge-group-using-p-u-t-404-schema"></span> Schema

##### <span id="update-service-edge-group-using-p-u-t-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="update-service-edge-group-using-p-u-t-429-schema"></span> Schema

### <span id="update-service-edge-using-p-u-t"></span> Update the Service Edge details. (*updateServiceEdgeUsingPUT*)

```
PUT /mgmtconfig/v1/admin/customers/{customerId}/serviceEdge/{serviceEdgeId}
```

#### Consumes
  * application/json

#### Produces
  * */*

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| customerId | `path` | string | `string` |  | ✓ |  | The unique identifier of the ZPA tenant. |
| serviceEdgeId | `path` | string | `string` |  | ✓ |  | The unique identifier of the Service Edge. |
| serviceEdge | `body` | [ServiceEdge](#service-edge) | `models.ServiceEdge` | | ✓ | | The object of the Service Edge. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#update-service-edge-using-p-u-t-201) | Created | Created |  | [schema](#update-service-edge-using-p-u-t-201-schema) |
| [204](#update-service-edge-using-p-u-t-204) | No Content | No Content |  | [schema](#update-service-edge-using-p-u-t-204-schema) |
| [400](#update-service-edge-using-p-u-t-400) | Bad Request | BadRequest |  | [schema](#update-service-edge-using-p-u-t-400-schema) |
| [401](#update-service-edge-using-p-u-t-401) | Unauthorized | Unauthorized |  | [schema](#update-service-edge-using-p-u-t-401-schema) |
| [403](#update-service-edge-using-p-u-t-403) | Forbidden | Forbidden |  | [schema](#update-service-edge-using-p-u-t-403-schema) |
| [404](#update-service-edge-using-p-u-t-404) | Not Found | Not Found |  | [schema](#update-service-edge-using-p-u-t-404-schema) |
| [429](#update-service-edge-using-p-u-t-429) | Too Many Requests | TooManyRequest |  | [schema](#update-service-edge-using-p-u-t-429-schema) |

#### Responses


##### <span id="update-service-edge-using-p-u-t-201"></span> 201 - Created
Status: Created

###### <span id="update-service-edge-using-p-u-t-201-schema"></span> Schema

##### <span id="update-service-edge-using-p-u-t-204"></span> 204 - No Content
Status: No Content

###### <span id="update-service-edge-using-p-u-t-204-schema"></span> Schema

##### <span id="update-service-edge-using-p-u-t-400"></span> 400 - BadRequest
Status: Bad Request

###### <span id="update-service-edge-using-p-u-t-400-schema"></span> Schema

##### <span id="update-service-edge-using-p-u-t-401"></span> 401 - Unauthorized
Status: Unauthorized

###### <span id="update-service-edge-using-p-u-t-401-schema"></span> Schema

##### <span id="update-service-edge-using-p-u-t-403"></span> 403 - Forbidden
Status: Forbidden

###### <span id="update-service-edge-using-p-u-t-403-schema"></span> Schema

##### <span id="update-service-edge-using-p-u-t-404"></span> 404 - Not Found
Status: Not Found

###### <span id="update-service-edge-using-p-u-t-404-schema"></span> Schema

##### <span id="update-service-edge-using-p-u-t-429"></span> 429 - TooManyRequest
Status: Too Many Requests

###### <span id="update-service-edge-using-p-u-t-429-schema"></span> Schema

## Models

### <span id="app-connector-group"></span> AppConnectorGroup


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| cityCountry | string| `string` |  | |  |  |
| connectors | [][Connector](#connector)| `[]*Connector` |  | |  |  |
| countryCode | string| `string` |  | |  |  |
| creationTime | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| dnsQueryType | string| `string` |  | |  |  |
| enabled | boolean| `bool` |  | |  |  |
| geoLocationId | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| latitude | string| `string` |  | |  |  |
| location | string| `string` |  | |  |  |
| longitude | string| `string` |  | |  |  |
| lssAppConnectorGroup | boolean| `bool` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |
| overrideVersionProfile | boolean| `bool` |  | |  |  |
| serverGroups | [][AppServerGroup](#app-server-group)| `[]*AppServerGroup` |  | |  |  |
| upgradeDay | string| `string` |  | |  |  |
| upgradeTimeInSecs | string| `string` |  | |  |  |
| versionProfileId | string| `string` |  | |  |  |
| versionProfileName | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| versionProfileVisibilityScope | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |



### <span id="app-connector-group-resource"></span> AppConnectorGroupResource


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| id | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |



### <span id="app-port-range"></span> AppPortRange


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| from | string| `string` |  | |  |  |
| to | string| `string` |  | |  |  |



### <span id="app-server-group"></span> AppServerGroup


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| configSpace | string| `string` |  | |  |  |
| creationTime | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| dynamicDiscovery | boolean| `bool` |  | |  |  |
| enabled | boolean| `bool` |  | |  |  |
| id | string| `string` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |



### <span id="app-server-group-resource"></span> AppServerGroupResource


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| id | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |



### <span id="application"></span> Application


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| bypassType | string| `string` |  | |  |  |
| configSpace | string| `string` |  | |  |  |
| creationTime | string| `string` |  | |  |  |
| defaultIdleTimeout | string| `string` |  | |  |  |
| defaultMaxAge | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| domainName | string| `string` |  | |  |  |
| domainNames | []string| `[]string` |  | |  |  |
| doubleEncrypt | boolean| `bool` |  | |  |  |
| enabled | boolean| `bool` |  | |  |  |
| healthCheckType | string| `string` |  | |  |  |
| icmpAccessType | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| ipAnchored | boolean| `bool` |  | |  |  |
| logFeatures | []string| `[]string` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |
| passiveHealthEnabled | boolean| `bool` |  | |  |  |
| serverGroups | [][AppServerGroup](#app-server-group)| `[]*AppServerGroup` |  | |  |  |
| tcpPortRange | [][AppPortRange](#app-port-range)| `[]*AppPortRange` |  | |  |  |
| tcpPortRanges | []string| `[]string` |  | |  |  |
| tcpPortsIn | []string| `[]string` |  | |  |  |
| tcpPortsOut | []string| `[]string` |  | |  |  |
| udpPortRange | [][AppPortRange](#app-port-range)| `[]*AppPortRange` |  | |  |  |
| udpPortRanges | []string| `[]string` |  | |  |  |



### <span id="application-resource"></span> ApplicationResource


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| bypassType | string| `string` |  | |  |  |
| clientlessApps | [][BAAppDto](#b-a-app-dto)| `[]*BAAppDto` |  | |  |  |
| commonAppsDto | [CommonApplicationDto](#common-application-dto)| `CommonApplicationDto` |  | |  |  |
| configSpace | string| `string` |  | |  |  |
| creationTime | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| defaultIdleTimeout | string| `string` |  | |  |  |
| defaultMaxAge | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| domainNames | []string| `[]string` |  | |  |  |
| doubleEncrypt | boolean| `bool` |  | |  |  |
| enabled | boolean| `bool` |  | |  |  |
| healthCheckType | string| `string` |  | |  |  |
| healthReporting | string| `string` |  | |  |  |
| icmpAccessType | string| `string` |  | |  |  |
| id | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| inspectionApps | [][InspectAppDto](#inspect-app-dto)| `[]*InspectAppDto` |  | |  |  |
| ipAnchored | boolean| `bool` |  | |  |  |
| isCnameEnabled | boolean| `bool` |  | |  |  |
| modifiedBy | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| modifiedTime | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| name | string| `string` |  | |  |  |
| passiveHealthEnabled | boolean| `bool` |  | |  |  |
| segmentGroupId | string| `string` |  | |  |  |
| segmentGroupName | string| `string` |  | |  |  |
| serverGroups | [][AppServerGroup](#app-server-group)| `[]*AppServerGroup` |  | |  |  |
| tcpPortRange | [][AppPortRange](#app-port-range)| `[]*AppPortRange` |  | |  |  |
| tcpPortRanges | []string| `[]string` |  | |  |  |
| udpPortRange | [][AppPortRange](#app-port-range)| `[]*AppPortRange` |  | |  |  |
| udpPortRanges | []string| `[]string` |  | |  |  |



### <span id="application-server"></span> ApplicationServer


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| address | string| `string` |  | |  |  |
| appServerGroupIds | []string| `[]string` |  | |  |  |
| configSpace | string| `string` |  | |  |  |
| creationTime | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| enabled | boolean| `bool` |  | |  |  |
| id | string| `string` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |



### <span id="b-a-app-dto"></span> BAAppDto


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| allowOptions | boolean| `bool` |  | |  |  |
| appId | string| `string` |  | |  |  |
| applicationPort | string| `string` |  | |  |  |
| applicationProtocol | string| `string` |  | |  |  |
| certificateId | string| `string` |  | |  |  |
| certificateName | string| `string` |  | |  |  |
| cname | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| domain | string| `string` |  | |  |  |
| enabled | boolean| `bool` |  | |  |  |
| hidden | boolean| `bool` |  | |  |  |
| id | string| `string` |  | |  |  |
| localDomain | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |
| path | string| `string` |  | |  |  |
| portal | boolean| `bool` |  | |  |  |
| trustUntrustedCert | boolean| `bool` |  | |  |  |



### <span id="b-a-certificate"></span> BACertificate


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| cName | string| `string` |  | |  |  |
| certChain | string| `string` |  | |  |  |
| certificate | string| `string` | ✓ | | The certificate text is in PEM format. |  |
| creationTime | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| issuedBy | string| `string` |  | |  |  |
| issuedTo | string| `string` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |
| publicKey | string| `string` |  | |  |  |
| san | []string| `[]string` |  | |  |  |
| serialNo | string| `string` |  | |  |  |
| status | string| `string` |  | |  |  |
| validFromInEpochSec | string| `string` |  | |  |  |
| validToInEpochSec | string| `string` |  | |  |  |



### <span id="bulk-delete-resource"></span> BulkDeleteResource


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ids | []int64 (formatted integer)| `[]int64` |  | |  |  |



### <span id="cloud-connector-group-resource"></span> CloudConnectorGroupResource


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| cloudConnectors | [][Znf](#znf)| `[]*Znf` |  | |  |  |
| creationTime | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| description | string| `string` |  | |  |  |
| enabled | boolean| `bool` |  | |  |  |
| geoLocationId | string| `string` |  | |  |  |
| id | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| modifiedBy | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| modifiedTime | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| name | string| `string` |  | |  |  |
| ziaCloud | string| `string` |  | |  |  |
| ziaOrgId | string| `string` |  | |  |  |



### <span id="common-app-config-dto"></span> CommonAppConfigDto


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| allowOptions | boolean| `bool` |  | |  |  |
| appId | string| `string` |  | |  |  |
| appTypes | []string| `[]string` |  | |  |  |
| applicationPort | string| `string` |  | |  |  |
| applicationProtocol | string| `string` |  | |  |  |
| baAppId | string| `string` |  | |  |  |
| certificateId | string| `string` |  | |  |  |
| certificateName | string| `string` |  | |  |  |
| cname | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| domain | string| `string` |  | |  |  |
| enabled | boolean| `bool` |  | |  |  |
| hidden | boolean| `bool` |  | |  |  |
| inspectAppId | string| `string` |  | |  |  |
| localDomain | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |
| path | string| `string` |  | |  |  |
| portal | boolean| `bool` |  | |  |  |
| trustUntrustedCert | boolean| `bool` |  | |  |  |



### <span id="common-application-dto"></span> CommonApplicationDto


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| appsConfig | [][CommonAppConfigDto](#common-app-config-dto)| `[]*CommonAppConfigDto` |  | |  |  |
| deletedBaApps | []int64 (formatted integer)| `[]int64` |  | |  |  |
| deletedInspectApps | []int64 (formatted integer)| `[]int64` |  | |  |  |



### <span id="condition-set"></span> ConditionSet


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationTime | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| negated | boolean| `bool` |  | |  |  |
| operands | [][Operand](#operand)| `[]*Operand` |  | |  |  |
| operator | string| `string` |  | |  |  |



### <span id="condition-set-resource"></span> ConditionSetResource


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationTime | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| id | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| modifiedBy | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| modifiedTime | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| negated | boolean| `bool` |  | |  |  |
| operands | [][OperandResource](#operand-resource)| `[]*OperandResource` |  | |  |  |
| operator | string| `string` |  | |  |  |
| setIds | []int64 (formatted integer)| `[]int64` |  | |  |  |



### <span id="connector"></span> Connector


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| appConnectorGroupId | string| `string` |  | |  |  |
| appConnectorGroupName | string| `string` |  | |  |  |
| applicationStartTime | string| `string` |  | |  |  |
| controlChannelStatus | string| `string` |  | | Read only. Ignored in PUT/POST calls. Expected values: UNKNOWN/ZPN_STATUS_AUTHENTICATED(1)/ZPN_STATUS_DISCONNECTED |  |
| creationTime | string| `string` |  | |  |  |
| ctrlBrokerName | string| `string` |  | | Read only. Ignored in PUT/POST calls. |  |
| currentVersion | string| `string` |  | | Read only. Ignored in PUT/POST calls. |  |
| description | string| `string` |  | |  |  |
| enabled | boolean| `bool` |  | |  |  |
| enrollmentCert | map of string| `map[string]string` |  | |  |  |
| expectedUpgradeTime | string| `string` |  | | Read only. Ignored in PUT/POST calls. |  |
| expectedVersion | string| `string` |  | | Read only. Ignored in PUT/POST calls. |  |
| fingerprint | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| ipAcl | []string| `[]string` |  | |  |  |
| issuedCertId | string| `string` |  | |  |  |
| lastBrokerConnectTime | string| `string` |  | | Read only. Ignored in PUT/POST calls. |  |
| lastBrokerConnectTimeDuration | string| `string` |  | | Read only. Ignored in PUT/POST calls. |  |
| lastBrokerDisconnectTime | string| `string` |  | | Read only. Ignored in PUT/POST calls. |  |
| lastBrokerDisconnectTimeDuration | string| `string` |  | | Read only. Ignored in PUT/POST calls. |  |
| lastUpgradeTime | string| `string` |  | | Read only. Ignored in PUT/POST calls. |  |
| latitude | double (formatted number)| `float64` |  | |  |  |
| location | string| `string` |  | |  |  |
| longitude | double (formatted number)| `float64` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |
| platform | string| `string` |  | |  |  |
| previousVersion | string| `string` |  | | Read only. Ignored in PUT/POST calls. |  |
| privateIp | string| `string` |  | |  |  |
| provisioningKeyId | string| `string` |  | |  |  |
| provisioningKeyName | string| `string` |  | |  |  |
| publicIp | string| `string` |  | |  |  |
| sargeVersion | string| `string` |  | |  |  |
| upgradeAttempt | string| `string` |  | | Read only. Ignored in PUT/POST calls. |  |
| upgradeStatus | string| `string` |  | | Read only. Ignored in PUT/POST calls. |  |



### <span id="custom-scope-request-customer-ids"></span> CustomScopeRequestCustomerIds


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| addCustomerIds | []int64 (formatted integer)| `[]int64` |  | |  |  |
| deleteCustomerIds | []int64 (formatted integer)| `[]int64` |  | |  |  |



### <span id="customer-id-name-d-t-o"></span> CustomerIdNameDTO


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| customerId | string| `string` |  | |  |  |
| excludeConstellation | boolean| `bool` |  | |  |  |
| name | string| `string` |  | |  |  |



### <span id="enrollment-cert"></span> EnrollmentCert


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| allowSigning | boolean| `bool` |  | | Not applicable for put and post calls. This field will be extracted from certificate. |  |
| cName | string| `string` |  | | Not applicable for put and post calls. This field will be extracted from certificate. |  |
| certificate | string| `string` | ✓ | | certificate text in pem format. |  |
| clientCertType | string| `string` |  | | Read only property. Applicable only for GET calls, ignored in PUT/POST calls. |  |
| creationTime | string| `string` |  | |  |  |
| csr | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| issuedBy | string| `string` |  | |  |  |
| issuedTo | string| `string` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |
| parentCertId | string| `string` |  | |  |  |
| parentCertName | string| `string` |  | | Read only property. Applicable only for GET calls, ignored in PUT/POST calls. |  |
| privateKey | string| `string` |  | |  |  |
| privateKeyPresent | boolean| `bool` |  | | Read only property. Applicable only for GET calls, ignored in PUT/POST calls. |  |
| serialNo | string| `string` |  | |  |  |
| validFromInEpochSec | string| `string` |  | | Not applicable for put and post calls. This field will be extracted from certificate. |  |
| validToInEpochSec | string| `string` |  | | Not applicable for put and post calls. This field will be extracted from certificate. |  |
| zrsaencryptedprivatekey | string| `string` |  | |  |  |
| zrsaencryptedsessionkey | string| `string` |  | |  |  |



### <span id="idp"></span> Idp


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| adminMetadata | [ServiceProvider](#service-provider)| `ServiceProvider` |  | |  |  |
| adminSpSigningCertId | string| `string` |  | |  |  |
| autoProvision | string| `string` |  | | defaults to 0. |  |
| certificates | [][IdpCertDto](#idp-cert-dto)| `[]*IdpCertDto` |  | |  |  |
| creationTime | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| disableSamlBasedPolicy | boolean| `bool` |  | |  |  |
| domainList | []string| `[]string` |  | |  |  |
| enableScimBasedPolicy | boolean| `bool` |  | |  |  |
| enabled | boolean| `bool` |  | | Default value if null is True. |  |
| id | string| `string` |  | |  |  |
| idpEntityId | string| `string` | ✓ | |  |  |
| loginNameAttribute | string| `string` |  | |  |  |
| loginUrl | string| `string` | ✓ | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |
| reauthOnUserUpdate | boolean| `bool` |  | |  |  |
| redirectBinding | boolean| `bool` |  | |  |  |
| scimEnabled | boolean| `bool` |  | |  |  |
| scimServiceProviderEndpoint | string| `string` |  | |  |  |
| scimSharedSecretExists | boolean| `bool` |  | |  |  |
| signSamlRequest | string| `string` |  | | Defaults to 1. |  |
| ssoType | []string| `[]string` |  | |  |  |
| useCustomSPMetadata | boolean| `bool` |  | |  |  |
| userMetadata | [ServiceProvider](#service-provider)| `ServiceProvider` |  | |  |  |
| userSpSigningCertId | string| `string` |  | |  |  |



### <span id="idp-cert-dto"></span> IdpCertDto


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| cName | string| `string` |  | |  |  |
| certificate | string| `string` |  | |  |  |
| serialNo | string| `string` |  | |  |  |
| validFromInSec | string| `string` |  | |  |  |
| validToInSec | string| `string` |  | |  |  |



### <span id="inspect-app-dto"></span> InspectAppDto


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| appId | string| `string` |  | |  |  |
| applicationPort | string| `string` |  | |  |  |
| applicationProtocol | string| `string` |  | |  |  |
| certificateId | string| `string` |  | |  |  |
| certificateName | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| domain | string| `string` |  | |  |  |
| enabled | boolean| `bool` |  | |  |  |
| id | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |



### <span id="lhs-rhs-value-resource"></span> LhsRhsValueResource


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| lhs | string| `string` |  | |  |  |
| rhs | string| `string` |  | |  |  |



### <span id="lss"></span> Lss


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| auditMessage | string| `string` |  | |  |  |
| creationTime | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| enabled | boolean| `bool` |  | |  |  |
| filter | []string| `[]string` |  | |  |  |
| format | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| lssHost | string| `string` |  | |  |  |
| lssPort | string| `string` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |
| sourceLogType | string| `string` |  | |  |  |
| useTls | boolean| `bool` |  | |  |  |



### <span id="lss-resource"></span> LssResource


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| config | [Lss](#lss)| `Lss` |  | |  |  |
| connectorGroups | [][AppConnectorGroup](#app-connector-group)| `[]*AppConnectorGroup` |  | |  |  |
| id | string| `string` |  | |  |  |
| policyRule | [PolicyRule](#policy-rule)| `PolicyRule` |  | |  |  |
| policyRuleResource | [PolicyRuleResource](#policy-rule-resource)| `PolicyRuleResource` |  | |  |  |



### <span id="machine"></span> Machine


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationTime | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| enrollmentCert | map of string| `map[string]string` |  | |  |  |
| fingerprint | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| issuedCertId | string| `string` |  | |  |  |
| machineGroupId | string| `string` |  | |  |  |
| machineGroupName | string| `string` |  | |  |  |
| machineTokenId | string| `string` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |



### <span id="machine-group"></span> MachineGroup


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationTime | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| enabled | boolean| `bool` |  | |  |  |
| id | string| `string` |  | |  |  |
| machines | [][Machine](#machine)| `[]*Machine` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |



### <span id="name-id-dto"></span> NameIdDto


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| id | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |



### <span id="operand"></span> Operand


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationTime | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| idpId | string| `string` |  | |  |  |
| lhs | string| `string` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |
| objectType | string| `string` |  | |  |  |
| rhs | string| `string` |  | |  |  |



### <span id="operand-resource"></span> OperandResource


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| commonProperties | [Operand](#operand)| `Operand` |  | |  |  |
| creationTime | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| entryValues | [][LHSRHSValueResource](#lhs-rhs-value-resource)| `[]*LHSRHSValueResource` |  | |  |  |
| id | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| idpId | string| `string` |  | |  |  |
| modifiedBy | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| modifiedTime | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| objectType | string| `string` |  | |  |  |
| values | []string| `[]string` |  | |  |  |



### <span id="page-list-of-app-connector-group"></span> PageListOfAppConnectorGroup


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| list | [][AppConnectorGroup](#app-connector-group)| `[]*AppConnectorGroup` |  | |  |  |
| totalPages | string| `string` |  | |  |  |



### <span id="page-list-of-application-resource"></span> PageListOfApplicationResource


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| list | [][ApplicationResource](#application-resource)| `[]*ApplicationResource` |  | |  |  |
| totalPages | string| `string` |  | |  |  |



### <span id="page-list-of-application-server"></span> PageListOfApplicationServer


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| list | [][ApplicationServer](#application-server)| `[]*ApplicationServer` |  | |  |  |
| totalPages | string| `string` |  | |  |  |



### <span id="page-list-of-b-a-certificate"></span> PageListOfBACertificate


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| list | [][BACertificate](#b-a-certificate)| `[]*BACertificate` |  | |  |  |
| totalPages | string| `string` |  | |  |  |



### <span id="page-list-of-cloud-connector-group-resource"></span> PageListOfCloudConnectorGroupResource


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| list | [][CloudConnectorGroupResource](#cloud-connector-group-resource)| `[]*CloudConnectorGroupResource` |  | |  |  |
| totalPages | string| `string` |  | |  |  |



### <span id="page-list-of-connector"></span> PageListOfConnector


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| list | [][Connector](#connector)| `[]*Connector` |  | |  |  |
| totalPages | string| `string` |  | |  |  |



### <span id="page-list-of-enrollment-cert"></span> PageListOfEnrollmentCert


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| list | [][EnrollmentCert](#enrollment-cert)| `[]*EnrollmentCert` |  | |  |  |
| totalPages | string| `string` |  | |  |  |



### <span id="page-list-of-idp"></span> PageListOfIdp


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| list | [][Idp](#idp)| `[]*Idp` |  | |  |  |
| totalPages | string| `string` |  | |  |  |



### <span id="page-list-of-lss-resource"></span> PageListOfLssResource


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| list | [][LssResource](#lss-resource)| `[]*LssResource` |  | |  |  |
| totalPages | string| `string` |  | |  |  |



### <span id="page-list-of-machine-group"></span> PageListOfMachineGroup


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| list | [][MachineGroup](#machine-group)| `[]*MachineGroup` |  | |  |  |
| totalPages | string| `string` |  | |  |  |



### <span id="page-list-of-policy-rule"></span> PageListOfPolicyRule


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| list | [][PolicyRule](#policy-rule)| `[]*PolicyRule` |  | |  |  |
| totalPages | string| `string` |  | |  |  |



### <span id="page-list-of-posture-profile"></span> PageListOfPostureProfile


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| list | [][PostureProfile](#posture-profile)| `[]*PostureProfile` |  | |  |  |
| totalPages | string| `string` |  | |  |  |



### <span id="page-list-of-provisioning-key"></span> PageListOfProvisioningKey


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| list | [][ProvisioningKey](#provisioning-key)| `[]*ProvisioningKey` |  | |  |  |
| totalPages | string| `string` |  | |  |  |



### <span id="page-list-of-s-c-i-m-attribute-header"></span> PageListOfSCIMAttributeHeader


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| list | [][SCIMAttributeHeader](#s-c-i-m-attribute-header)| `[]*SCIMAttributeHeader` |  | |  |  |
| totalPages | string| `string` |  | |  |  |



### <span id="page-list-of-saml-attribute"></span> PageListOfSamlAttribute


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| list | [][SamlAttribute](#saml-attribute)| `[]*SamlAttribute` |  | |  |  |
| totalPages | string| `string` |  | |  |  |



### <span id="page-list-of-segment-group"></span> PageListOfSegmentGroup


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| list | [][SegmentGroup](#segment-group)| `[]*SegmentGroup` |  | |  |  |
| totalPages | string| `string` |  | |  |  |



### <span id="page-list-of-server-group-d-t-o"></span> PageListOfServerGroupDTO


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| list | [][ServerGroupDTO](#server-group-d-t-o)| `[]*ServerGroupDTO` |  | |  |  |
| totalPages | string| `string` |  | |  |  |



### <span id="page-list-of-service-edge"></span> PageListOfServiceEdge


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| list | [][ServiceEdge](#service-edge)| `[]*ServiceEdge` |  | |  |  |
| totalPages | string| `string` |  | |  |  |



### <span id="page-list-of-service-edge-group"></span> PageListOfServiceEdgeGroup


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| list | [][ServiceEdgeGroup](#service-edge-group)| `[]*ServiceEdgeGroup` |  | |  |  |
| totalPages | string| `string` |  | |  |  |



### <span id="page-list-of-trusted-network"></span> PageListOfTrustedNetwork


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| list | [][TrustedNetwork](#trusted-network)| `[]*TrustedNetwork` |  | |  |  |
| totalPages | string| `string` |  | |  |  |



### <span id="page-list-of-version-profile"></span> PageListOfVersionProfile


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| list | [][VersionProfile](#version-profile)| `[]*VersionProfile` |  | |  |  |
| totalPages | string| `string` |  | |  |  |



### <span id="policy-rule"></span> PolicyRule


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| action | string| `string` |  | |  |  |
| actionId | string| `string` |  | |  |  |
| appConnectorGroups | [][AppConnectorGroup](#app-connector-group)| `[]*AppConnectorGroup` |  | |  |  |
| appServerGroups | [][AppServerGroup](#app-server-group)| `[]*AppServerGroup` |  | |  |  |
| conditions | [][ConditionSet](#condition-set)| `[]*ConditionSet` |  | |  |  |
| creationTime | string| `string` |  | |  |  |
| customMsg | string| `string` |  | |  |  |
| defaultRule | boolean| `bool` |  | |  |  |
| defaultRuleName | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |
| operator | string| `string` |  | |  |  |
| policySetId | string| `string` |  | |  |  |
| policyType | string| `string` |  | |  |  |
| priority | string| `string` |  | |  |  |
| reauthIdleTimeout | string| `string` |  | |  |  |
| reauthTimeout | string| `string` |  | |  |  |
| ruleOrder | string| `string` |  | |  |  |
| zpnCbiProfileId | string| `string` |  | |  |  |
| zpnInspectionProfileId | string| `string` |  | |  |  |
| zpnInspectionProfileName | string| `string` |  | |  |  |



### <span id="policy-rule-resource"></span> PolicyRuleResource


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| action | string| `string` |  | |  |  |
| actionId | string| `string` |  | |  |  |
| appConnectorGroups | [][AppConnectorGroupResource](#app-connector-group-resource)| `[]*AppConnectorGroupResource` |  | |  |  |
| appServerGroups | [][AppServerGroupResource](#app-server-group-resource)| `[]*AppServerGroupResource` |  | |  |  |
| auditMessage | string| `string` |  | |  |  |
| conditions | [][ConditionSetResource](#condition-set-resource)| `[]*ConditionSetResource` |  | |  |  |
| creationTime | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| customMsg | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| id | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| modifiedBy | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| modifiedTime | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| name | string| `string` |  | |  |  |
| operator | string| `string` |  | |  |  |
| policySetId | string| `string` |  | |  |  |
| policyType | string| `string` |  | |  |  |
| priority | string| `string` |  | |  |  |
| reauthIdleTimeout | string| `string` |  | |  |  |
| reauthTimeout | string| `string` |  | |  |  |
| ruleOrder | string| `string` |  | |  |  |
| version | string| `string` |  | |  |  |
| zpnCbiProfileId | string| `string` |  | |  |  |
| zpnInspectionProfileId | string| `string` |  | |  |  |
| zpnInspectionProfileName | string| `string` |  | |  |  |



### <span id="policy-set"></span> PolicySet


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationTime | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| enabled | boolean| `bool` |  | |  |  |
| id | string| `string` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |
| policyType | string| `string` |  | |  |  |
| rules | [][PolicyRule](#policy-rule)| `[]*PolicyRule` |  | |  |  |
| sorted | boolean| `bool` |  | |  |  |



### <span id="posture-profile"></span> PostureProfile


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationTime | string| `string` |  | |  |  |
| domain | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| masterCustomerId | string| `string` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |
| postureUdid | string| `string` |  | |  |  |
| zscalerCloud | string| `string` |  | |  |  |
| zscalerCustomerId | string| `string` |  | |  |  |



### <span id="provisioning-key"></span> Provisioning Key


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| appConnectorGroupId | string| `string` |  | |  |  |
| appConnectorGroupName | string| `string` |  | | Read only property. Applicable only for GET calls, ignored in PUT/POST calls. |  |
| creationTime | string| `string` |  | |  |  |
| enabled | boolean| `bool` |  | |  |  |
| enrollmentCertId | string| `string` | ✓ | |  |  |
| enrollmentCertName | string| `string` |  | | Read only property. Applicable only for GET calls, ignored in PUT/POST calls. |  |
| expirationInEpochSec | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| ipAcl | []string| `[]string` |  | |  |  |
| maxUsage | string| `string` | ✓ | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |
| provisioningKey | string| `string` |  | | read only field. Ignored in PUT/POST calls. |  |
| uiConfig | string| `string` |  | |  |  |
| usageCount | string| `string` |  | | defaults to 0 |  |
| zcomponentId | string| `string` |  | |  |  |
| zcomponentName | string| `string` |  | | Read only property. Applicable only for GET calls, ignored in PUT/POST calls. |  |



### <span id="s-c-i-m-attribute-header"></span> SCIMAttributeHeader


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| canonicalValues | []string| `[]string` |  | |  |  |
| caseSensitive | boolean| `bool` |  | |  |  |
| creationTime | string| `string` |  | |  |  |
| dataType | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| idpId | string| `string` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| multivalued | boolean| `bool` |  | |  |  |
| mutability | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |
| required | boolean| `bool` |  | |  |  |
| returned | string| `string` |  | |  |  |
| schemaURI | string| `string` |  | |  |  |
| uniqueness | boolean| `bool` |  | |  |  |



### <span id="saml-attribute"></span> SamlAttribute


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationTime | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| idpId | string| `string` |  | |  |  |
| idpName | string| `string` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |
| samlName | string| `string` |  | |  |  |
| userAttribute | boolean| `bool` |  | |  |  |



### <span id="segment-group"></span> SegmentGroup


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| applications | [][Application](#application)| `[]*Application` |  | |  |  |
| configSpace | string| `string` |  | |  |  |
| creationTime | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| enabled | boolean| `bool` |  | |  |  |
| id | string| `string` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |
| policyMigrated | boolean| `bool` |  | |  |  |
| tcpKeepAliveEnabled | string| `string` |  | |  |  |



### <span id="server-group-d-t-o"></span> ServerGroupDTO


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| appConnectorGroups | [][AppConnectorGroup](#app-connector-group)| `[]*AppConnectorGroup` |  | |  |  |
| applications | [][NameIDDto](#name-id-dto)| `[]*NameIDDto` |  | |  |  |
| configSpace | string| `string` |  | |  |  |
| creationTime | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| description | string| `string` |  | |  |  |
| dynamicDiscovery | boolean| `bool` |  | | Defaults to false. |  |
| enabled | boolean| `bool` |  | |  |  |
| id | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| ipAnchored | boolean| `bool` |  | |  |  |
| modifiedBy | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| modifiedTime | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| name | string| `string` |  | |  |  |
| servers | [][ApplicationServer](#application-server)| `[]*ApplicationServer` |  | |  |  |



### <span id="service-edge"></span> ServiceEdge


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| applicationStartTime | string| `string` |  | |  |  |
| controlChannelStatus | string| `string` |  | | ReadOnly. Ignored in PUT/POST calls. Expected values: UNKNOWN/ZPN_STATUS_AUTHENTICATED(1)/ZPN_STATUS_DISCONNECTED |  |
| creationTime | string| `string` |  | |  |  |
| ctrlBrokerName | string| `string` |  | | ReadOnly. Ignored in PUT/POST calls. |  |
| currentVersion | string| `string` |  | | ReadOnly. Ignored in PUT/POST calls |  |
| description | string| `string` |  | |  |  |
| enabled | boolean| `bool` |  | |  |  |
| enrollmentCert | map of string| `map[string]string` |  | |  |  |
| expectedUpgradeTime | string| `string` |  | | ReadOnly. Ignored in PUT/POST calls |  |
| expectedVersion | string| `string` |  | | ReadOnly. Ignored in PUT/POST calls |  |
| fingerprint | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| ipAcl | []string| `[]string` |  | |  |  |
| issuedCertId | string| `string` |  | |  |  |
| lastBrokerConnectTime | string| `string` |  | | ReadOnly. Ignored in PUT/POST calls. |  |
| lastBrokerConnectTimeDuration | string| `string` |  | | Read only. Ignored in PUT/POST calls. |  |
| lastBrokerDisconnectTime | string| `string` |  | | ReadOnly. Ignored in PUT/POST calls. |  |
| lastBrokerDisconnectTimeDuration | string| `string` |  | | Read only. Ignored in PUT/POST calls. |  |
| lastUpgradeTime | string| `string` |  | | ReadOnly. Ignored in PUT/POST calls |  |
| latitude | string| `string` |  | |  |  |
| listenIps | []string| `[]string` |  | |  |  |
| location | string| `string` |  | |  |  |
| longitude | string| `string` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |
| platform | string| `string` |  | |  |  |
| previousVersion | string| `string` |  | | ReadOnly. Ignored in PUT/POST calls |  |
| privateIp | string| `string` |  | |  |  |
| provisioningKeyId | string| `string` |  | |  |  |
| provisioningKeyName | string| `string` |  | |  |  |
| publicIp | string| `string` |  | |  |  |
| publishIps | []string| `[]string` |  | |  |  |
| sargeVersion | string| `string` |  | |  |  |
| serviceEdgeGroupId | string| `string` |  | |  |  |
| serviceEdgeGroupName | string| `string` |  | |  |  |
| upgradeAttempt | string| `string` |  | | ReadOnly. Ignored in PUT/POST calls. |  |
| upgradeStatus | string| `string` |  | | ReadOnly. Ignored in PUT/POST calls |  |



### <span id="service-edge-group"></span> ServiceEdgeGroup


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| cityCountry | string| `string` |  | |  |  |
| countryCode | string| `string` |  | |  |  |
| creationTime | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| enabled | boolean| `bool` |  | |  |  |
| geoLocationId | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| isPublic | string| `string` |  | |  |  |
| latitude | string| `string` |  | |  |  |
| location | string| `string` |  | |  |  |
| longitude | string| `string` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |
| overrideVersionProfile | boolean| `bool` |  | |  |  |
| serviceEdges | [][ServiceEdge](#service-edge)| `[]*ServiceEdge` |  | |  |  |
| trustedNetworks | [][TrustedNetwork](#trusted-network)| `[]*TrustedNetwork` |  | |  |  |
| upgradeDay | string| `string` |  | |  |  |
| upgradeTimeInSecs | string| `string` |  | |  |  |
| versionProfileId | string| `string` |  | |  |  |
| versionProfileName | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |
| versionProfileVisibilityScope | string| `string` |  | | Only applicable for a GET request. Ignored in PUT/POST/DELETE requests. |  |



### <span id="service-provider"></span> ServiceProvider


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| certificateUrl | string| `string` |  | |  |  |
| spBaseUrl | string| `string` |  | |  |  |
| spEntityId | string| `string` |  | |  |  |
| spMetadataUrl | string| `string` |  | |  |  |
| spPostUrl | string| `string` |  | |  |  |



### <span id="trusted-network"></span> TrustedNetwork


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationTime | string| `string` |  | |  |  |
| domain | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| masterCustomerId | string| `string` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |
| networkId | string| `string` |  | |  |  |
| zscalerCloud | string| `string` |  | |  |  |



### <span id="version"></span> Version


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationTime | string| `string` |  | |  |  |
| customerId | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| platform | string| `string` |  | |  |  |
| restartAfterUptimeInDays | string| `string` |  | |  |  |
| role | string| `string` |  | |  |  |
| version | string| `string` |  | |  |  |
| version_profile_gid | string| `string` |  | |  |  |



### <span id="version-profile"></span> VersionProfile


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationTime | string| `string` |  | |  |  |
| customScopeCustomerIds | [][CustomerIDNameDTO](#customer-id-name-d-t-o)| `[]*CustomerIDNameDTO` |  | | Only applicable for a GET requests. Ignored in PUT/POST requests. |  |
| customScopeRequestCustomerIds | [CustomScopeRequestCustomerIds](#custom-scope-request-customer-ids)| `CustomScopeRequestCustomerIds` |  | | Not applicable for GET requests. Field is only applicable in PUT/POST requests. |  |
| customerId | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` |  | |  |  |
| upgradePriority | string| `string` |  | |  |  |
| versions | [][Version](#version)| `[]*Version` |  | |  |  |
| visibilityScope | string| `string` |  | |  |  |



### <span id="znf"></span> Znf


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| creationTime | string| `string` |  | |  |  |
| description | string| `string` |  | |  |  |
| enabled | boolean| `bool` |  | |  |  |
| enrollmentCert | map of string| `map[string]string` |  | |  |  |
| fingerprint | string| `string` |  | |  |  |
| id | string| `string` |  | |  |  |
| ipAcl | []string| `[]string` |  | |  |  |
| issuedCertId | string| `string` |  | |  |  |
| modifiedBy | string| `string` |  | |  |  |
| modifiedTime | string| `string` |  | |  |  |
| name | string| `string` | ✓ | |  |  |


