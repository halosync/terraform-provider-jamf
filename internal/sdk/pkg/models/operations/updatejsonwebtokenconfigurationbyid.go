// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type UpdateJSONWebTokenConfigurationByIDRequest struct {
	// ID value to filter by
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateJSONWebTokenConfigurationByIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
