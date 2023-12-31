// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"jamf/internal/sdk/pkg/models/shared"
	"net/http"
)

// FindMobileDevicesByNameSubsetSubset - Subset to filter by
type FindMobileDevicesByNameSubsetSubset string

const (
	FindMobileDevicesByNameSubsetSubsetGeneral               FindMobileDevicesByNameSubsetSubset = "General"
	FindMobileDevicesByNameSubsetSubsetLocation              FindMobileDevicesByNameSubsetSubset = "Location"
	FindMobileDevicesByNameSubsetSubsetPurchasing            FindMobileDevicesByNameSubsetSubset = "Purchasing"
	FindMobileDevicesByNameSubsetSubsetApplications          FindMobileDevicesByNameSubsetSubset = "Applications"
	FindMobileDevicesByNameSubsetSubsetSecurity              FindMobileDevicesByNameSubsetSubset = "Security"
	FindMobileDevicesByNameSubsetSubsetNetwork               FindMobileDevicesByNameSubsetSubset = "Network"
	FindMobileDevicesByNameSubsetSubsetCertificates          FindMobileDevicesByNameSubsetSubset = "Certificates"
	FindMobileDevicesByNameSubsetSubsetConfigurationProfiles FindMobileDevicesByNameSubsetSubset = "ConfigurationProfiles"
	FindMobileDevicesByNameSubsetSubsetProvisioningProfiles  FindMobileDevicesByNameSubsetSubset = "ProvisioningProfiles"
	FindMobileDevicesByNameSubsetSubsetMobileDeviceGroups    FindMobileDevicesByNameSubsetSubset = "MobileDeviceGroups"
	FindMobileDevicesByNameSubsetSubsetExtensionAttributes   FindMobileDevicesByNameSubsetSubset = "ExtensionAttributes"
)

func (e FindMobileDevicesByNameSubsetSubset) ToPointer() *FindMobileDevicesByNameSubsetSubset {
	return &e
}

func (e *FindMobileDevicesByNameSubsetSubset) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "General":
		fallthrough
	case "Location":
		fallthrough
	case "Purchasing":
		fallthrough
	case "Applications":
		fallthrough
	case "Security":
		fallthrough
	case "Network":
		fallthrough
	case "Certificates":
		fallthrough
	case "ConfigurationProfiles":
		fallthrough
	case "ProvisioningProfiles":
		fallthrough
	case "MobileDeviceGroups":
		fallthrough
	case "ExtensionAttributes":
		*e = FindMobileDevicesByNameSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDevicesByNameSubsetSubset: %v", v)
	}
}

type FindMobileDevicesByNameSubsetRequest struct {
	// Name to filter by
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Subset to filter by
	Subset FindMobileDevicesByNameSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
}

type FindMobileDevicesByNameSubsetResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	MobileDevice *shared.MobileDevice
}
