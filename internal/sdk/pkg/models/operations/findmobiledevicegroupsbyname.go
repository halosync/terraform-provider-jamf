// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"jamf/internal/sdk/pkg/models/shared"
	"net/http"
)

type FindMobileDeviceGroupsByNameRequest struct {
	// Name to filter by
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

type FindMobileDeviceGroupsByNameResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	MobileDeviceGroup *shared.MobileDeviceGroup
}
