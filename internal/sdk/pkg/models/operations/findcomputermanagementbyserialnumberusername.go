// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"jamf/internal/sdk/pkg/models/shared"
	"net/http"
)

type FindComputerManagementBySerialNumberUsernameRequest struct {
	// Computer serial number to filter by
	Serialnumber string `pathParam:"style=simple,explode=false,name=serialnumber"`
	// Username to filter by
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

type FindComputerManagementBySerialNumberUsernameResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ComputerManagement *shared.ComputerManagement
}
