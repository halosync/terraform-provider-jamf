// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FindOsxConfigurationProfilesByIDSubsetSubset - Subset to filter by
type FindOsxConfigurationProfilesByIDSubsetSubset string

const (
	FindOsxConfigurationProfilesByIDSubsetSubsetGeneral     FindOsxConfigurationProfilesByIDSubsetSubset = "General"
	FindOsxConfigurationProfilesByIDSubsetSubsetScope       FindOsxConfigurationProfilesByIDSubsetSubset = "Scope"
	FindOsxConfigurationProfilesByIDSubsetSubsetSelfService FindOsxConfigurationProfilesByIDSubsetSubset = "SelfService"
)

func (e FindOsxConfigurationProfilesByIDSubsetSubset) ToPointer() *FindOsxConfigurationProfilesByIDSubsetSubset {
	return &e
}

func (e *FindOsxConfigurationProfilesByIDSubsetSubset) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "General":
		fallthrough
	case "Scope":
		fallthrough
	case "SelfService":
		*e = FindOsxConfigurationProfilesByIDSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindOsxConfigurationProfilesByIDSubsetSubset: %v", v)
	}
}

type FindOsxConfigurationProfilesByIDSubsetRequest struct {
	// ID to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
	// Subset to filter by
	Subset FindOsxConfigurationProfilesByIDSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
}

type FindOsxConfigurationProfilesByIDSubsetResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
