// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type UpdateUserByEmailAddressRequest struct {
	// Email address value to filter by
	Email string `pathParam:"style=simple,explode=false,name=email"`
}

type UpdateUserByEmailAddressResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}