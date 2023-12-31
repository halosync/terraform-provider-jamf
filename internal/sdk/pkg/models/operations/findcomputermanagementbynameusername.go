// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"jamf/internal/sdk/pkg/models/shared"
	"net/http"
)

type FindComputerManagementByNameUsernameRequest struct {
	// Computer name to filter by
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Username to filter by
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

type FindComputerManagementByNameUsernameResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ComputerManagement *shared.ComputerManagement
}
