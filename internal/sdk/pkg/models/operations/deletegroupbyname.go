// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteGroupByNameRequest struct {
	// Name value to filter by
	Name int64 `pathParam:"style=simple,explode=false,name=name"`
}

type DeleteGroupByNameResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
