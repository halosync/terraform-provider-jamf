// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CommandflushMobileDevicesMobileDevice struct {
	ID int64
}

type CommandflushMobileDevices struct {
	MobileDevice *CommandflushMobileDevicesMobileDevice
}

type CommandflushStatus string

const (
	CommandflushStatusPending           CommandflushStatus = "Pending"
	CommandflushStatusFailed            CommandflushStatus = "Failed"
	CommandflushStatusPendingPlusFailed CommandflushStatus = "Pending+Failed"
)

func (e CommandflushStatus) ToPointer() *CommandflushStatus {
	return &e
}

func (e *CommandflushStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Pending":
		fallthrough
	case "Failed":
		fallthrough
	case "Pending+Failed":
		*e = CommandflushStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CommandflushStatus: %v", v)
	}
}

type Commandflush struct {
	MobileDevices *CommandflushMobileDevices
	Status        CommandflushStatus
}
