// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id       uint        `json:"id"       description:""`
	Passport string      `json:"passport" description:""`
	Password string      `json:"password" description:""`
	Nickname string      `json:"nickname" description:""`
	CreateAt *gtime.Time `json:"createAt" description:""`
	UpdateAt *gtime.Time `json:"updateAt" description:""`
	PhoneNum string      `json:"phoneNum" description:""`
}
