// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"jamf/internal/sdk/pkg/models/shared"
	"jamf/internal/sdk/pkg/types"
	"net/http"
)

type FindComputerHardwareSoftwareReportsByNameRequest struct {
	// End date (e.g. yyyy-mm-dd)
	EndDate types.Date `pathParam:"style=simple,explode=false,name=end_date"`
	// Name to filter by
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Start date (e.g. yyyy-mm-dd)
	StartDate types.Date `pathParam:"style=simple,explode=false,name=start_date"`
}

type FindComputerHardwareSoftwareReportsByNameResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ComputerHardwareSoftwareReports *shared.ComputerHardwareSoftwareReports
}
