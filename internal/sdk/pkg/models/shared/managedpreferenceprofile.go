// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ManagedPreferenceProfileGeneral struct {
	Enabled *bool   `json:"enabled,omitempty"`
	ID      *int64  `json:"id,omitempty"`
	Name    string  `json:"name"`
	Plist   *string `json:"plist,omitempty"`
}

// ManagedPreferenceProfile - OK
type ManagedPreferenceProfile struct {
	General *ManagedPreferenceProfileGeneral `json:"general,omitempty"`
}
