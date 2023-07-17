// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"jamf/internal/sdk/pkg/models/shared"
	"net/http"
)

type FindComputerApplicationUsageByUDIDRequest struct {
	// End date (e.g. yyyy-mm-dd)
	EndDate string `pathParam:"style=simple,explode=false,name=end_date"`
	// Start date (e.g. yyyy-mm-dd)
	StartDate string `pathParam:"style=simple,explode=false,name=start_date"`
	// UDID to filter by
	Udid string `pathParam:"style=simple,explode=false,name=udid"`
}

type FindComputerApplicationUsageByUDIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ComputerApplicationUsage []shared.ComputerApplicationUsage
}