// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type AdvancedUserSearchCriteriaCriterionAndOr string

const (
	AdvancedUserSearchCriteriaCriterionAndOrAnd AdvancedUserSearchCriteriaCriterionAndOr = "and"
	AdvancedUserSearchCriteriaCriterionAndOrOr  AdvancedUserSearchCriteriaCriterionAndOr = "or"
)

func (e AdvancedUserSearchCriteriaCriterionAndOr) ToPointer() *AdvancedUserSearchCriteriaCriterionAndOr {
	return &e
}

func (e *AdvancedUserSearchCriteriaCriterionAndOr) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "and":
		fallthrough
	case "or":
		*e = AdvancedUserSearchCriteriaCriterionAndOr(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AdvancedUserSearchCriteriaCriterionAndOr: %v", v)
	}
}

type AdvancedUserSearchCriteriaCriterion struct {
	AndOr        *AdvancedUserSearchCriteriaCriterionAndOr `json:"and_or,omitempty"`
	ClosingParen *bool                                     `json:"closing_paren,omitempty"`
	// Name of the criteria
	Name         *string `json:"name,omitempty"`
	OpeningParen *bool   `json:"opening_paren,omitempty"`
	Priority     *int64  `json:"priority,omitempty"`
	// Operator
	SearchType *string `json:"search_type,omitempty"`
	Value      *string `json:"value,omitempty"`
}

type AdvancedUserSearchCriteria struct {
	Criterion *AdvancedUserSearchCriteriaCriterion `json:"criterion,omitempty"`
	Size      *int64                               `json:"size,omitempty"`
}

type AdvancedUserSearchDisplayFieldsDisplayField struct {
	// Name of the display field
	Name *string `json:"name,omitempty"`
}

type AdvancedUserSearchDisplayFields struct {
	DisplayField *AdvancedUserSearchDisplayFieldsDisplayField `json:"display_field,omitempty"`
	Size         *int64                                       `json:"size,omitempty"`
}

type AdvancedUserSearchUsersUser struct {
	Username *string `json:"Username,omitempty"`
	ID       *int64  `json:"id,omitempty"`
	// Name of the user
	Name *string `json:"name,omitempty"`
}

type AdvancedUserSearchUsers struct {
	Size *int64                       `json:"size,omitempty"`
	User *AdvancedUserSearchUsersUser `json:"user,omitempty"`
}

// AdvancedUserSearch - OK
type AdvancedUserSearch struct {
	Criteria      []AdvancedUserSearchCriteria      `json:"criteria,omitempty"`
	DisplayFields []AdvancedUserSearchDisplayFields `json:"display_fields,omitempty"`
	ID            *int64                            `json:"id,omitempty"`
	// Name of the advanced user search
	Name  string                    `json:"name"`
	Site  *SiteObject               `json:"site,omitempty"`
	Users []AdvancedUserSearchUsers `json:"users,omitempty"`
}
