// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"jamf/internal/sdk/pkg/models/shared"
	"net/http"
)

type FindComputerManagementByUDIDRequest struct {
	// Computer UDID to filter by
	Udid string `pathParam:"style=simple,explode=false,name=udid"`
}

type FindComputerManagementByUDIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ComputerManagement *shared.ComputerManagement
}
