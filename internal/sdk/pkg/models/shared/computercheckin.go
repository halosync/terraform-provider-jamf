// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ComputerCheckIn - OK
type ComputerCheckIn struct {
	ApplyComputerLevelManagedPreferences *bool `json:"apply_computer_level_managed_preferences,omitempty"`
	ApplyUserLevelManagedPreferences     *bool `json:"apply_user_level_managed_preferences,omitempty"`
	CheckForPoliciesAtLoginLogout        *bool `json:"check_for_policies_at_login_logout,omitempty"`
	CheckForPoliciesAtStartup            *bool `json:"check_for_policies_at_startup,omitempty"`
	// Measured in minutes
	CheckInFrequency                *int64 `json:"check_in_frequency,omitempty"`
	CreateLoginLogoutHooks          *bool  `json:"create_login_logout_hooks,omitempty"`
	CreateStartupScript             *bool  `json:"create_startup_script,omitempty"`
	DisplayStatusToUser             *bool  `json:"display_status_to_user,omitempty"`
	EnsureSSHIsEnabled              *bool  `json:"ensure_ssh_is_enabled,omitempty"`
	HideRestorePartition            *bool  `json:"hide_restore_partition,omitempty"`
	LogStartupEvent                 *bool  `json:"log_startup_event,omitempty"`
	LogUsername                     *bool  `json:"log_username,omitempty"`
	PerformLoginActionsInBackground *bool  `json:"perform_login_actions_in_background,omitempty"`
}