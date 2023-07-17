// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type CreateMobileDeviceProvisioningProfilesByUUIDRequest struct {
	// Uuid value to filter by
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type CreateMobileDeviceProvisioningProfilesByUUIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}