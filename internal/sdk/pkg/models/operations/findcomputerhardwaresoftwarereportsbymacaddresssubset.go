// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"jamf/internal/sdk/pkg/models/shared"
	"jamf/internal/sdk/pkg/types"
	"net/http"
)

// FindComputerHardwareSoftwareReportsByMacAddressSubsetSubset - Subset to filter by
type FindComputerHardwareSoftwareReportsByMacAddressSubsetSubset string

const (
	FindComputerHardwareSoftwareReportsByMacAddressSubsetSubsetSoftware FindComputerHardwareSoftwareReportsByMacAddressSubsetSubset = "Software"
	FindComputerHardwareSoftwareReportsByMacAddressSubsetSubsetHardwre  FindComputerHardwareSoftwareReportsByMacAddressSubsetSubset = "Hardwre"
	FindComputerHardwareSoftwareReportsByMacAddressSubsetSubsetFonts    FindComputerHardwareSoftwareReportsByMacAddressSubsetSubset = "Fonts"
	FindComputerHardwareSoftwareReportsByMacAddressSubsetSubsetPlugins  FindComputerHardwareSoftwareReportsByMacAddressSubsetSubset = "Plugins"
)

func (e FindComputerHardwareSoftwareReportsByMacAddressSubsetSubset) ToPointer() *FindComputerHardwareSoftwareReportsByMacAddressSubsetSubset {
	return &e
}

func (e *FindComputerHardwareSoftwareReportsByMacAddressSubsetSubset) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Software":
		fallthrough
	case "Hardwre":
		fallthrough
	case "Fonts":
		fallthrough
	case "Plugins":
		*e = FindComputerHardwareSoftwareReportsByMacAddressSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerHardwareSoftwareReportsByMacAddressSubsetSubset: %v", v)
	}
}

type FindComputerHardwareSoftwareReportsByMacAddressSubsetRequest struct {
	// End date (e.g. yyyy-mm-dd)
	EndDate types.Date `pathParam:"style=simple,explode=false,name=end_date"`
	// MAC address to filter by
	Macaddress string `pathParam:"style=simple,explode=false,name=macaddress"`
	// Start date (e.g. yyyy-mm-dd)
	StartDate types.Date `pathParam:"style=simple,explode=false,name=start_date"`
	// Subset to filter by
	Subset FindComputerHardwareSoftwareReportsByMacAddressSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
}

type FindComputerHardwareSoftwareReportsByMacAddressSubsetResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ComputerHardwareSoftwareReports *shared.ComputerHardwareSoftwareReports
}
