// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ClassAppleTvsAppleTv struct {
	AirplayPassword *string `json:"airplay_password,omitempty"`
	DeviceID        *string `json:"device_id,omitempty"`
	Name            *string `json:"name,omitempty"`
	Udid            *string `json:"udid,omitempty"`
	WifiMacAddress  *string `json:"wifi_mac_address,omitempty"`
}

type ClassAppleTvs struct {
	AppleTv *ClassAppleTvsAppleTv `json:"apple_tv,omitempty"`
}

type ClassMeetingTimesMeetingTimeDays string

const (
	ClassMeetingTimesMeetingTimeDaysM  ClassMeetingTimesMeetingTimeDays = "M"
	ClassMeetingTimesMeetingTimeDaysT  ClassMeetingTimesMeetingTimeDays = "T"
	ClassMeetingTimesMeetingTimeDaysW  ClassMeetingTimesMeetingTimeDays = "W"
	ClassMeetingTimesMeetingTimeDaysTh ClassMeetingTimesMeetingTimeDays = "Th"
	ClassMeetingTimesMeetingTimeDaysF  ClassMeetingTimesMeetingTimeDays = "F"
	ClassMeetingTimesMeetingTimeDaysSa ClassMeetingTimesMeetingTimeDays = "Sa"
	ClassMeetingTimesMeetingTimeDaysSu ClassMeetingTimesMeetingTimeDays = "Su"
)

func (e ClassMeetingTimesMeetingTimeDays) ToPointer() *ClassMeetingTimesMeetingTimeDays {
	return &e
}

func (e *ClassMeetingTimesMeetingTimeDays) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "M":
		fallthrough
	case "T":
		fallthrough
	case "W":
		fallthrough
	case "Th":
		fallthrough
	case "F":
		fallthrough
	case "Sa":
		fallthrough
	case "Su":
		*e = ClassMeetingTimesMeetingTimeDays(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ClassMeetingTimesMeetingTimeDays: %v", v)
	}
}

type ClassMeetingTimesMeetingTime struct {
	Days      *ClassMeetingTimesMeetingTimeDays `json:"days,omitempty"`
	EndTime   *int64                            `json:"end_time,omitempty"`
	StartTime *int64                            `json:"start_time,omitempty"`
}

type ClassMeetingTimes struct {
	MeetingTime *ClassMeetingTimesMeetingTime `json:"meeting_time,omitempty"`
}

type ClassMobileDeviceGroupID struct {
	ID *int64 `json:"id,omitempty"`
}

type ClassMobileDevicesMobileDevice struct {
	// Name of the device
	Name           *string `json:"name,omitempty"`
	Udid           *string `json:"udid,omitempty"`
	WifiMacAddress *string `json:"wifi_mac_address,omitempty"`
}

type ClassMobileDevices struct {
	MobileDevice *ClassMobileDevicesMobileDevice `json:"mobile_device,omitempty"`
}

type ClassStudentGroupIds struct {
	ID *int64 `json:"id,omitempty"`
}

type ClassStudents struct {
	// Name of the student
	Student *string `json:"student,omitempty"`
}

type ClassTeacherGroupIds struct {
	ID *int64 `json:"id,omitempty"`
}

type ClassTeacherIds struct {
	ID *int64 `json:"id,omitempty"`
}

type ClassTeachers struct {
	// Name of the teacher
	Teacher *string `json:"teacher,omitempty"`
}

// Class - OK
type Class struct {
	AppleTvs            []ClassAppleTvs            `json:"apple_tvs,omitempty"`
	Description         *string                    `json:"description,omitempty"`
	ID                  *int64                     `json:"id,omitempty"`
	MeetingTimes        *ClassMeetingTimes         `json:"meeting_times,omitempty"`
	MobileDeviceGroup   *IDName                    `json:"mobile_device_group,omitempty"`
	MobileDeviceGroupID []ClassMobileDeviceGroupID `json:"mobile_device_group_id,omitempty"`
	MobileDevices       []ClassMobileDevices       `json:"mobile_devices,omitempty"`
	// Name of the class
	Name            string                 `json:"name"`
	Site            *SiteObject            `json:"site,omitempty"`
	Source          *string                `json:"source,omitempty"`
	StudentGroupIds []ClassStudentGroupIds `json:"student_group_ids,omitempty"`
	Students        []ClassStudents        `json:"students,omitempty"`
	TeacherGroupIds []ClassTeacherGroupIds `json:"teacher_group_ids,omitempty"`
	TeacherIds      []ClassTeacherIds      `json:"teacher_ids,omitempty"`
	Teachers        []ClassTeachers        `json:"teachers,omitempty"`
}