// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"jamf/internal/sdk/pkg/models/shared"
	"net/http"
)

type FindComputerManagementByIDUsernameRequest struct {
	// Computer ID to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
	// Username to filter by
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

type FindComputerManagementByIDUsernameResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ComputerManagement *shared.ComputerManagement
}
