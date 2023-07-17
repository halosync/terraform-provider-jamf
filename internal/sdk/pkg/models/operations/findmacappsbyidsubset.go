// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FindMacappsByIDSubsetSubset - Subset to filter by
type FindMacappsByIDSubsetSubset string

const (
	FindMacappsByIDSubsetSubsetGeneral     FindMacappsByIDSubsetSubset = "General"
	FindMacappsByIDSubsetSubsetScope       FindMacappsByIDSubsetSubset = "Scope"
	FindMacappsByIDSubsetSubsetSelfService FindMacappsByIDSubsetSubset = "SelfService"
	FindMacappsByIDSubsetSubsetVppCodes    FindMacappsByIDSubsetSubset = "VPPCodes"
	FindMacappsByIDSubsetSubsetVpp         FindMacappsByIDSubsetSubset = "VPP"
)

func (e FindMacappsByIDSubsetSubset) ToPointer() *FindMacappsByIDSubsetSubset {
	return &e
}

func (e *FindMacappsByIDSubsetSubset) UnmarshalJSON(data []byte) error {
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
		fallthrough
	case "VPPCodes":
		fallthrough
	case "VPP":
		*e = FindMacappsByIDSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMacappsByIDSubsetSubset: %v", v)
	}
}

type FindMacappsByIDSubsetRequest struct {
	// ID to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
	// Subset to filter by
	Subset FindMacappsByIDSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
}

type FindMacappsByIDSubsetResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
