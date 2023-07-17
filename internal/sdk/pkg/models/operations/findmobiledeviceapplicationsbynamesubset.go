// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"jamf/internal/sdk/pkg/models/shared"
	"net/http"
)

// FindMobileDeviceApplicationsByNameSubsetSubset - Subset to filter by
type FindMobileDeviceApplicationsByNameSubsetSubset string

const (
	FindMobileDeviceApplicationsByNameSubsetSubsetGeneral          FindMobileDeviceApplicationsByNameSubsetSubset = "General"
	FindMobileDeviceApplicationsByNameSubsetSubsetScope            FindMobileDeviceApplicationsByNameSubsetSubset = "Scope"
	FindMobileDeviceApplicationsByNameSubsetSubsetSelfService      FindMobileDeviceApplicationsByNameSubsetSubset = "SelfService"
	FindMobileDeviceApplicationsByNameSubsetSubsetVppCodes         FindMobileDeviceApplicationsByNameSubsetSubset = "VPPCodes"
	FindMobileDeviceApplicationsByNameSubsetSubsetVpp              FindMobileDeviceApplicationsByNameSubsetSubset = "VPP"
	FindMobileDeviceApplicationsByNameSubsetSubsetAppConfiguration FindMobileDeviceApplicationsByNameSubsetSubset = "AppConfiguration"
)

func (e FindMobileDeviceApplicationsByNameSubsetSubset) ToPointer() *FindMobileDeviceApplicationsByNameSubsetSubset {
	return &e
}

func (e *FindMobileDeviceApplicationsByNameSubsetSubset) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "General":
		fallthrough
	case "Scope":
		fallthrough
	case "SelfService":
		fallthrough
	case "VPPCodes":
		fallthrough
	case "VPP":
		fallthrough
	case "AppConfiguration":
		*e = FindMobileDeviceApplicationsByNameSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceApplicationsByNameSubsetSubset: %v", v)
	}
}

type FindMobileDeviceApplicationsByNameSubsetRequest struct {
	// Name to filter by
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Subset to filter by
	Subset FindMobileDeviceApplicationsByNameSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
}

type FindMobileDeviceApplicationsByNameSubsetResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	MobileDeviceApplication *shared.MobileDeviceApplication
}