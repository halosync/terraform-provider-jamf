// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"jamf/internal/sdk/pkg/models/shared"
	"net/http"
)

type FindMobileDeviceApplicationsByBundleIDRequest struct {
	// Bundle ID to filter by
	Bundleid string `pathParam:"style=simple,explode=false,name=bundleid"`
}

type FindMobileDeviceApplicationsByBundleIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	MobileDeviceApplication *shared.MobileDeviceApplication
}
