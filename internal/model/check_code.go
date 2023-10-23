package model

import "github.com/gogf/gf/v2/os/gtime"

type CheckIdInput struct {
	PhoneNumber string `json:"phoneNumber" description:""`
	CodeId      string
	EndTime     *gtime.Time
}
