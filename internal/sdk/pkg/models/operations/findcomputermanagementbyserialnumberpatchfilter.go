// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"jamf/internal/sdk/pkg/models/shared"
	"net/http"
)

type FindComputerManagementBySerialNumberPatchFilterRequest struct {
	// filter to apply
	Filter string `pathParam:"style=simple,explode=false,name=filter"`
	// Computer serial number to filter by
	Serialnumber string `pathParam:"style=simple,explode=false,name=serialnumber"`
}

type FindComputerManagementBySerialNumberPatchFilterResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ComputerManagement *shared.ComputerManagement
}
