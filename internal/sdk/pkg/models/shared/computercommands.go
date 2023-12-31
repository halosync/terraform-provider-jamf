// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ComputerCommandsComputerCommandComputersComputer struct {
	ID        *int64  `json:"id,omitempty"`
	IPAddress *string `json:"ip_address,omitempty"`
	Udid      *string `json:"udid,omitempty"`
	Users     *string `json:"users,omitempty"`
}

type ComputerCommandsComputerCommandComputers struct {
	Computer *ComputerCommandsComputerCommandComputersComputer `json:"computer,omitempty"`
	Size     *int64                                            `json:"size,omitempty"`
}

type ComputerCommandsComputerCommand struct {
	// Command type
	Command   *string                                    `json:"command,omitempty"`
	Computers []ComputerCommandsComputerCommandComputers `json:"computers,omitempty"`
	ID        *int64                                     `json:"id,omitempty"`
	ProfileID *int64                                     `json:"profile_id,omitempty"`
	Udid      *string                                    `json:"udid,omitempty"`
	UUID      *string                                    `json:"uuid,omitempty"`
}

type ComputerCommands struct {
	ComputerCommand *ComputerCommandsComputerCommand `json:"computer_command,omitempty"`
	Size            *int64                           `json:"size,omitempty"`
}
