// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ComputerExtensionAttributesComputerExtensionAttribute struct {
	Enabled *bool   `json:"enabled,omitempty"`
	ID      *int64  `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
}

type ComputerExtensionAttributes struct {
	ComputerExtensionAttribute *ComputerExtensionAttributesComputerExtensionAttribute `json:"computer_extension_attribute,omitempty"`
	Size                       *int64                                                 `json:"size,omitempty"`
}
