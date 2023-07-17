// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"jamf/internal/sdk/pkg/models/shared"
	"net/http"
)

// FindComputerManagementByUDIDSubsetSubset - Subset to filter by
type FindComputerManagementByUDIDSubsetSubset string

const (
	FindComputerManagementByUDIDSubsetSubsetGeneral                      FindComputerManagementByUDIDSubsetSubset = "General"
	FindComputerManagementByUDIDSubsetSubsetPolicies                     FindComputerManagementByUDIDSubsetSubset = "Policies"
	FindComputerManagementByUDIDSubsetSubsetEbooks                       FindComputerManagementByUDIDSubsetSubset = "Ebooks"
	FindComputerManagementByUDIDSubsetSubsetMacAppStoreApps              FindComputerManagementByUDIDSubsetSubset = "MacAppStoreApps"
	FindComputerManagementByUDIDSubsetSubsetOsxConfigurationProfiles     FindComputerManagementByUDIDSubsetSubset = "OSXConfigurationProfiles"
	FindComputerManagementByUDIDSubsetSubsetManagedPreferenceProfiles    FindComputerManagementByUDIDSubsetSubset = "ManagedPreferenceProfiles"
	FindComputerManagementByUDIDSubsetSubsetRestrictedSoftware           FindComputerManagementByUDIDSubsetSubset = "RestrictedSoftware"
	FindComputerManagementByUDIDSubsetSubsetSmartGroups                  FindComputerManagementByUDIDSubsetSubset = "SmartGroups"
	FindComputerManagementByUDIDSubsetSubsetStaticGroups                 FindComputerManagementByUDIDSubsetSubset = "StaticGroups"
	FindComputerManagementByUDIDSubsetSubsetPatchReportingSoftwareTitles FindComputerManagementByUDIDSubsetSubset = "PatchReportingSoftwareTitles"
)

func (e FindComputerManagementByUDIDSubsetSubset) ToPointer() *FindComputerManagementByUDIDSubsetSubset {
	return &e
}

func (e *FindComputerManagementByUDIDSubsetSubset) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "General":
		fallthrough
	case "Policies":
		fallthrough
	case "Ebooks":
		fallthrough
	case "MacAppStoreApps":
		fallthrough
	case "OSXConfigurationProfiles":
		fallthrough
	case "ManagedPreferenceProfiles":
		fallthrough
	case "RestrictedSoftware":
		fallthrough
	case "SmartGroups":
		fallthrough
	case "StaticGroups":
		fallthrough
	case "PatchReportingSoftwareTitles":
		*e = FindComputerManagementByUDIDSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerManagementByUDIDSubsetSubset: %v", v)
	}
}

type FindComputerManagementByUDIDSubsetRequest struct {
	// Subset to filter by
	Subset FindComputerManagementByUDIDSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
	// Computer UDID to filter by
	Udid string `pathParam:"style=simple,explode=false,name=udid"`
}

type FindComputerManagementByUDIDSubsetResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ComputerManagement *shared.ComputerManagement
}
