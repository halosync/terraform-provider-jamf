// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateComputerCommandByCommandIDAndPasscodeCommand - Command to send
type CreateComputerCommandByCommandIDAndPasscodeCommand string

const (
	CreateComputerCommandByCommandIDAndPasscodeCommandDeviceLock  CreateComputerCommandByCommandIDAndPasscodeCommand = "DeviceLock"
	CreateComputerCommandByCommandIDAndPasscodeCommandEraseDevice CreateComputerCommandByCommandIDAndPasscodeCommand = "EraseDevice"
)

func (e CreateComputerCommandByCommandIDAndPasscodeCommand) ToPointer() *CreateComputerCommandByCommandIDAndPasscodeCommand {
	return &e
}

func (e *CreateComputerCommandByCommandIDAndPasscodeCommand) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DeviceLock":
		fallthrough
	case "EraseDevice":
		*e = CreateComputerCommandByCommandIDAndPasscodeCommand(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateComputerCommandByCommandIDAndPasscodeCommand: %v", v)
	}
}

type CreateComputerCommandByCommandIDAndPasscodeRequest struct {
	// Command to send
	Command CreateComputerCommandByCommandIDAndPasscodeCommand `pathParam:"style=simple,explode=false,name=command"`
	// Computer ID - supports comma separated values (e.g id/8,10,55)
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Passcode to apply to device
	Passcode string `pathParam:"style=simple,explode=false,name=passcode"`
}

type CreateComputerCommandByCommandIDAndPasscodeResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}