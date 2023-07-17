// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type MobileDeviceApplicationAppConfiguration struct {
	Preferences *string `json:"preferences,omitempty"`
}

type MobileDeviceApplicationGeneralDeploymentType string

const (
	MobileDeviceApplicationGeneralDeploymentTypeMakeAvailableInSelfService               MobileDeviceApplicationGeneralDeploymentType = "Make Available in Self Service"
	MobileDeviceApplicationGeneralDeploymentTypeInstallAutomaticallyPromptUsersToInstall MobileDeviceApplicationGeneralDeploymentType = "Install Automatically/Prompt Users to Install"
)

func (e MobileDeviceApplicationGeneralDeploymentType) ToPointer() *MobileDeviceApplicationGeneralDeploymentType {
	return &e
}

func (e *MobileDeviceApplicationGeneralDeploymentType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Make Available in Self Service":
		fallthrough
	case "Install Automatically/Prompt Users to Install":
		*e = MobileDeviceApplicationGeneralDeploymentType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MobileDeviceApplicationGeneralDeploymentType: %v", v)
	}
}

type MobileDeviceApplicationGeneralIcon struct {
	// base64 encoded
	Data *string `json:"data,omitempty"`
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	URI  *string `json:"uri,omitempty"`
}

type MobileDeviceApplicationGeneralIpa struct {
	Data *string `json:"data,omitempty"`
	Name *string `json:"name,omitempty"`
	URI  *string `json:"uri,omitempty"`
}

type MobileDeviceApplicationGeneral struct {
	BundleID                         string                                        `json:"bundle_id"`
	Category                         *CategoryObject                               `json:"category,omitempty"`
	DeployAsManagedApp               *bool                                         `json:"deploy_as_managed_app,omitempty"`
	DeployAutomatically              *bool                                         `json:"deploy_automatically,omitempty"`
	DeploymentType                   *MobileDeviceApplicationGeneralDeploymentType `json:"deployment_type,omitempty"`
	Description                      *string                                       `json:"description,omitempty"`
	DisplayName                      *string                                       `json:"display_name,omitempty"`
	ExternalURL                      *string                                       `json:"external_url,omitempty"`
	Free                             *bool                                         `json:"free,omitempty"`
	HostExternally                   *bool                                         `json:"host_externally,omitempty"`
	Icon                             *MobileDeviceApplicationGeneralIcon           `json:"icon,omitempty"`
	ID                               *int64                                        `json:"id,omitempty"`
	InternalApp                      *bool                                         `json:"internal_app,omitempty"`
	Ipa                              *MobileDeviceApplicationGeneralIpa            `json:"ipa,omitempty"`
	ItunesCountryRegion              *string                                       `json:"itunes_country_region,omitempty"`
	ItunesStoreURL                   *string                                       `json:"itunes_store_url,omitempty"`
	ItunesSyncTime                   *int64                                        `json:"itunes_sync_time,omitempty"`
	KeepDescriptionAndIconUpToDate   *bool                                         `json:"keep_description_and_icon_up_to_date,omitempty"`
	MakeAvailableAfterInstall        *bool                                         `json:"make_available_after_install,omitempty"`
	MobileDeviceProvisioningProfile  *int64                                        `json:"mobile_device_provisioning_profile,omitempty"`
	Name                             string                                        `json:"name"`
	PreventBackupOfAppData           *bool                                         `json:"prevent_backup_of_app_data,omitempty"`
	RemoveAppWhenMdmProfileIsRemoved *bool                                         `json:"remove_app_when_mdm_profile_is_removed,omitempty"`
	Site                             *SiteObject                                   `json:"site,omitempty"`
	TakeOverManagement               *bool                                         `json:"take_over_management,omitempty"`
	Version                          string                                        `json:"version"`
}

type MobileDeviceApplicationScopeBuildings struct {
	Building *IDName `json:"building,omitempty"`
}

type MobileDeviceApplicationScopeDepartments struct {
	Department *IDName `json:"department,omitempty"`
}

type MobileDeviceApplicationScopeExclusionsBuildings struct {
	Building *IDName `json:"building,omitempty"`
}

type MobileDeviceApplicationScopeExclusionsDepartments struct {
	Department *IDName `json:"department,omitempty"`
}

type MobileDeviceApplicationScopeExclusionsJssUserGroups struct {
	UserGroup *IDName `json:"user_group,omitempty"`
}

type MobileDeviceApplicationScopeExclusionsJssUsers struct {
	User *IDName `json:"user,omitempty"`
}

type MobileDeviceApplicationScopeExclusionsMobileDeviceGroups struct {
	MobileDeviceGroup *IDName `json:"mobile_device_group,omitempty"`
}

type MobileDeviceApplicationScopeExclusionsMobileDevicesMobileDevice struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the device
	Name           *string `json:"name,omitempty"`
	Udid           *string `json:"udid,omitempty"`
	WifiMacAddress *string `json:"wifi_mac_address,omitempty"`
}

type MobileDeviceApplicationScopeExclusionsMobileDevices struct {
	MobileDevice *MobileDeviceApplicationScopeExclusionsMobileDevicesMobileDevice `json:"mobile_device,omitempty"`
}

type MobileDeviceApplicationScopeExclusionsNetworkSegmentsNetworkSegment struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the network segment
	Name *string `json:"name,omitempty"`
	UID  *string `json:"uid,omitempty"`
}

type MobileDeviceApplicationScopeExclusionsNetworkSegments struct {
	NetworkSegment *MobileDeviceApplicationScopeExclusionsNetworkSegmentsNetworkSegment `json:"network_segment,omitempty"`
}

type MobileDeviceApplicationScopeExclusionsUserGroups struct {
	UserGroup *IDName `json:"user_group,omitempty"`
}

type MobileDeviceApplicationScopeExclusionsUsersUser struct {
	Name *string `json:"name,omitempty"`
}

type MobileDeviceApplicationScopeExclusionsUsers struct {
	User *MobileDeviceApplicationScopeExclusionsUsersUser `json:"user,omitempty"`
}

type MobileDeviceApplicationScopeExclusions struct {
	Buildings          []MobileDeviceApplicationScopeExclusionsBuildings          `json:"buildings,omitempty"`
	Departments        []MobileDeviceApplicationScopeExclusionsDepartments        `json:"departments,omitempty"`
	JssUserGroups      []MobileDeviceApplicationScopeExclusionsJssUserGroups      `json:"jss_user_groups,omitempty"`
	JssUsers           []MobileDeviceApplicationScopeExclusionsJssUsers           `json:"jss_users,omitempty"`
	MobileDeviceGroups []MobileDeviceApplicationScopeExclusionsMobileDeviceGroups `json:"mobile_device_groups,omitempty"`
	MobileDevices      []MobileDeviceApplicationScopeExclusionsMobileDevices      `json:"mobile_devices,omitempty"`
	NetworkSegments    []MobileDeviceApplicationScopeExclusionsNetworkSegments    `json:"network_segments,omitempty"`
	UserGroups         []MobileDeviceApplicationScopeExclusionsUserGroups         `json:"user_groups,omitempty"`
	Users              []MobileDeviceApplicationScopeExclusionsUsers              `json:"users,omitempty"`
}

type MobileDeviceApplicationScopeJssUserGroups struct {
	UserGroup *IDName `json:"user_group,omitempty"`
}

type MobileDeviceApplicationScopeJssUsers struct {
	User *IDName `json:"user,omitempty"`
}

type MobileDeviceApplicationScopeLimitationsNetworkSegments struct {
	NetworkSegment *IDName `json:"network_segment,omitempty"`
}

type MobileDeviceApplicationScopeLimitationsUserGroups struct {
	UserGroup *IDName `json:"user_group,omitempty"`
}

type MobileDeviceApplicationScopeLimitationsUsers struct {
	User *IDName `json:"user,omitempty"`
}

type MobileDeviceApplicationScopeLimitations struct {
	NetworkSegments []MobileDeviceApplicationScopeLimitationsNetworkSegments `json:"network_segments,omitempty"`
	UserGroups      []MobileDeviceApplicationScopeLimitationsUserGroups      `json:"user_groups,omitempty"`
	Users           []MobileDeviceApplicationScopeLimitationsUsers           `json:"users,omitempty"`
}

type MobileDeviceApplicationScopeMobileDeviceGroups struct {
	MobileDeviceGroup *IDName `json:"mobile_device_group,omitempty"`
}

type MobileDeviceApplicationScopeMobileDevicesMobileDevice struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the device
	Name           *string `json:"name,omitempty"`
	Udid           *string `json:"udid,omitempty"`
	WifiMacAddress *string `json:"wifi_mac_address,omitempty"`
}

type MobileDeviceApplicationScopeMobileDevices struct {
	MobileDevice *MobileDeviceApplicationScopeMobileDevicesMobileDevice `json:"mobile_device,omitempty"`
}

type MobileDeviceApplicationScope struct {
	AllJssUsers        *bool                                            `json:"all_jss_users,omitempty"`
	AllMobileDevices   *bool                                            `json:"all_mobile_devices,omitempty"`
	Buildings          []MobileDeviceApplicationScopeBuildings          `json:"buildings,omitempty"`
	Departments        []MobileDeviceApplicationScopeDepartments        `json:"departments,omitempty"`
	Exclusions         *MobileDeviceApplicationScopeExclusions          `json:"exclusions,omitempty"`
	JssUserGroups      []MobileDeviceApplicationScopeJssUserGroups      `json:"jss_user_groups,omitempty"`
	JssUsers           []MobileDeviceApplicationScopeJssUsers           `json:"jss_users,omitempty"`
	Limitations        *MobileDeviceApplicationScopeLimitations         `json:"limitations,omitempty"`
	MobileDeviceGroups []MobileDeviceApplicationScopeMobileDeviceGroups `json:"mobile_device_groups,omitempty"`
	MobileDevices      []MobileDeviceApplicationScopeMobileDevices      `json:"mobile_devices,omitempty"`
}

type MobileDeviceApplicationSelfServiceSelfServiceCategoriesCategory struct {
	DisplayIn *bool   `json:"display_in,omitempty"`
	ID        *int64  `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
}

type MobileDeviceApplicationSelfServiceSelfServiceCategories struct {
	Category *MobileDeviceApplicationSelfServiceSelfServiceCategoriesCategory `json:"category,omitempty"`
}

type MobileDeviceApplicationSelfServiceSelfServiceIcon struct {
	// base64 encoded
	Data *string `json:"data,omitempty"`
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	URI  *string `json:"uri,omitempty"`
}

type MobileDeviceApplicationSelfService struct {
	FeatureOnMainPage      *bool                                                     `json:"feature_on_main_page,omitempty"`
	Notification           *bool                                                     `json:"notification,omitempty"`
	NotificationMessage    *string                                                   `json:"notification_message,omitempty"`
	NotificationSubject    *string                                                   `json:"notification_subject,omitempty"`
	SelfServiceCategories  []MobileDeviceApplicationSelfServiceSelfServiceCategories `json:"self_service_categories,omitempty"`
	SelfServiceDescription *string                                                   `json:"self_service_description,omitempty"`
	SelfServiceIcon        *MobileDeviceApplicationSelfServiceSelfServiceIcon        `json:"self_service_icon,omitempty"`
}

type MobileDeviceApplicationVpp struct {
	AssignVppDeviceBasedLicenses *bool  `json:"assign_vpp_device_based_licenses,omitempty"`
	VppAdminAccountID            *int64 `json:"vpp_admin_account_id,omitempty"`
}

// MobileDeviceApplication - OK
type MobileDeviceApplication struct {
	AppConfiguration *MobileDeviceApplicationAppConfiguration `json:"app_configuration,omitempty"`
	General          *MobileDeviceApplicationGeneral          `json:"general,omitempty"`
	Scope            *MobileDeviceApplicationScope            `json:"scope,omitempty"`
	SelfService      *MobileDeviceApplicationSelfService      `json:"self_service,omitempty"`
	Vpp              *MobileDeviceApplicationVpp              `json:"vpp,omitempty"`
}
