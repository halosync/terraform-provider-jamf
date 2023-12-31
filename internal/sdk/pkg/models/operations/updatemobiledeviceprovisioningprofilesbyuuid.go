// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"jamf/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateMobileDeviceProvisioningProfilesByUUIDRequest struct {
	// Uuid value to filter by
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type UpdateMobileDeviceProvisioningProfilesByUUIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	MobileDeviceProvisioningProfile *shared.MobileDeviceProvisioningProfile
}
