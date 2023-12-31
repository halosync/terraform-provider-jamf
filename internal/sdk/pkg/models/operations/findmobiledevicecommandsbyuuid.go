// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"jamf/internal/sdk/pkg/models/shared"
	"net/http"
)

type FindMobileDeviceCommandsByUUIDRequest struct {
	// UUID value to filter by
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type FindMobileDeviceCommandsByUUIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	MobileDeviceCommand *shared.MobileDeviceCommand
}
