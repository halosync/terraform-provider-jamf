// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PatchAvailableTitlesAvailableTitlesAvailableTitle struct {
	AppName        *string `json:"app_name,omitempty"`
	CurrentVersion *string `json:"current_version,omitempty"`
	LastModified   *string `json:"last_modified,omitempty"`
	NameID         *string `json:"name_id,omitempty"`
	Publisher      *string `json:"publisher,omitempty"`
}

type PatchAvailableTitlesAvailableTitles struct {
	AvailableTitle *PatchAvailableTitlesAvailableTitlesAvailableTitle `json:"available_title,omitempty"`
}

// PatchAvailableTitles - OK
type PatchAvailableTitles struct {
	AvailableTitles []PatchAvailableTitlesAvailableTitles `json:"available_titles,omitempty"`
	Size            *int64                                `json:"size,omitempty"`
}