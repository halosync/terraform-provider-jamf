// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FindPoliciesByTypeCreatedBy - Type to filter by
type FindPoliciesByTypeCreatedBy string

const (
	FindPoliciesByTypeCreatedByJss    FindPoliciesByTypeCreatedBy = "jss"
	FindPoliciesByTypeCreatedByCasper FindPoliciesByTypeCreatedBy = "casper"
)

func (e FindPoliciesByTypeCreatedBy) ToPointer() *FindPoliciesByTypeCreatedBy {
	return &e
}

func (e *FindPoliciesByTypeCreatedBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "jss":
		fallthrough
	case "casper":
		*e = FindPoliciesByTypeCreatedBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindPoliciesByTypeCreatedBy: %v", v)
	}
}

type FindPoliciesByTypeRequest struct {
	// Type to filter by
	CreatedBy FindPoliciesByTypeCreatedBy `pathParam:"style=simple,explode=false,name=createdBy"`
}

type FindPoliciesByTypeResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
