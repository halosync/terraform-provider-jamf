// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type MobileDeviceGroupCriteria struct {
	Criterion *Criterion `json:"criterion,omitempty"`
	Size      *int64     `json:"size,omitempty"`
}

type MobileDeviceGroupMobileDevicesMobileDevice struct {
	ID         *int64  `json:"id,omitempty"`
	MacAddress *string `json:"mac_address,omitempty"`
	// Name of the device
	Name           *string `json:"name,omitempty"`
	SerialNumber   *string `json:"serial_number,omitempty"`
	Udid           *string `json:"udid,omitempty"`
	WifiMacAddress *string `json:"wifi_mac_address,omitempty"`
}

type MobileDeviceGroupMobileDevices struct {
	MobileDevice *MobileDeviceGroupMobileDevicesMobileDevice `json:"mobile_device,omitempty"`
}

// MobileDeviceGroup - OK
type MobileDeviceGroup struct {
	Criteria      []MobileDeviceGroupCriteria      `json:"criteria,omitempty"`
	ID            *int64                           `json:"id,omitempty"`
	IsSmart       bool                             `json:"is_smart"`
	MobileDevices []MobileDeviceGroupMobileDevices `json:"mobile_devices,omitempty"`
	Name          string                           `json:"name"`
	Site          *SiteObject                      `json:"site,omitempty"`
}
