// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PackageRequiredProcessor string

const (
	PackageRequiredProcessorNone PackageRequiredProcessor = "None"
	PackageRequiredProcessorPpc  PackageRequiredProcessor = "ppc"
	PackageRequiredProcessorX86  PackageRequiredProcessor = "x86"
)

func (e PackageRequiredProcessor) ToPointer() *PackageRequiredProcessor {
	return &e
}

func (e *PackageRequiredProcessor) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "None":
		fallthrough
	case "ppc":
		fallthrough
	case "x86":
		*e = PackageRequiredProcessor(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PackageRequiredProcessor: %v", v)
	}
}

// Package - OK
type Package struct {
	AllowUninstalled           *bool   `json:"allow_uninstalled,omitempty"`
	Category                   *string `json:"category,omitempty"`
	Filename                   string  `json:"filename"`
	FillExistingUsers          *bool   `json:"fill_existing_users,omitempty"`
	FillUserTemplate           *bool   `json:"fill_user_template,omitempty"`
	ID                         *int64  `json:"id,omitempty"`
	Info                       *string `json:"info,omitempty"`
	InstallIfReportedAvailable *bool   `json:"install_if_reported_available,omitempty"`
	// Name of the package
	Name              string                    `json:"name"`
	Notes             *string                   `json:"notes,omitempty"`
	OsRequirements    *string                   `json:"os_requirements,omitempty"`
	Priority          *int64                    `json:"priority,omitempty"`
	RebootRequired    *bool                     `json:"reboot_required,omitempty"`
	ReinstallOption   *string                   `json:"reinstall_option,omitempty"`
	RequiredProcessor *PackageRequiredProcessor `json:"required_processor,omitempty"`
	SendNotification  *bool                     `json:"send_notification,omitempty"`
	SwitchWithPackage *string                   `json:"switch_with_package,omitempty"`
	TriggeringFiles   *string                   `json:"triggering_files,omitempty"`
}
