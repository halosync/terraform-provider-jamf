// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"jamf/internal/sdk/pkg/models/shared"
	"net/http"
)

type FindComputerGroupsResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ComputerGroups []shared.ComputerGroups
}
