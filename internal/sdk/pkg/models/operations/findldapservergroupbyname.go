// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindLDAPServerGroupByNameRequest struct {
	// Group to filter by
	Group string `pathParam:"style=simple,explode=false,name=group"`
	// Server name to filter by
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

type FindLDAPServerGroupByNameResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
