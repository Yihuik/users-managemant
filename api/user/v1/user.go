package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
)

type ProfileReq struct {
	g.Meta `path:"/user/profile" method:"get" tags:"UserService" summary:"Get the profile of current user"`
}

type ProfileRes struct {
	*entity.User
}

type SignUpReq struct {
	g.Meta    `path:"/user/sign-up" method:"post" tags:"UserService" summary:"Sign up a new user account"`
	Passport  string `v:"required|length:6,16"`
	Password  string `v:"required|length:6,16"`
	Password2 string `v:"required|length:6,16|same:Password"`
	Nickname  string
	PhoneNum  string `v:"required|length:11,11"`
}
type SignUpRes struct{}

type GetUserListReq struct {
	g.Meta      `path:"/GetUserList" method:"get" summary:""`
	Page        int `v:"required"`
	RowsPerPage int `v:"required"`
}
type GetUserListRes struct {
	UserList []*UserRes `json:"commodityList"`
}

type UserRes struct {
	Passport string `json:"passport"     description:""`
	Password string `json:"password"     description:""`
	PhoneNum string `json:"phoneNum"     description:""`
	Nickname string `json:"Nickname"     description:""`
}

type CheckNumReq struct {
	g.Meta   `path:"/user/sign-up/check_num" method:"post" tags:"UserService" summary:"Sign up: send a message"`
	PhoneNum string `v:"required|length:11"`
}

type CheckNumRes struct{}

type CheckCodeIdReq struct {
	g.Meta   `path:"/user/sign-up/check_code" method:"post" tags:"UserService" summary:"Sign up:check a message"`
	CodeId   string `v:"required|length:4"`
	PhoneNum string `v:"required|length:11"`
}

type CheckCodeIdRes struct {
	UpdateTime string `json:"updateTime" description:""`
}

type SignInReq struct {
	g.Meta   `path:"/user/sign-in" method:"post" tags:"UserService" summary:"Sign in with exist account"`
	Passport string `v:"required"`
	Password string `v:"required"`
}
type SignInRes struct{}

type CheckPassportReq struct {
	g.Meta   `path:"/user/check-passport" method:"post" tags:"UserService" summary:"Check passport available"`
	Passport string `v:"required"`
}
type CheckPassportRes struct{}

type CheckNickNameReq struct {
	g.Meta   `path:"/user/check-nick-name" method:"post" tags:"UserService" summary:"Check nickname available"`
	Nickname string `v:"required"`
}
type CheckNickNameRes struct{}

type IsSignedInReq struct {
	g.Meta `path:"/user/is-signed-in" method:"post" tags:"UserService" summary:"Check current user is already signed-in"`
}
type IsSignedInRes struct {
	OK bool `dc:"True if current user is signed in; or else false"`
}

type SignOutReq struct {
	g.Meta `path:"/user/sign-out" method:"post" tags:"UserService" summary:"Sign out current user"`
}
type SignOutRes struct{}

type AddAdministratorReq struct {
	g.Meta   `path:"/administrator/Add" method:"post" tags:"UserService" summary:"Sign out current user"`
	Passport string `v:"required"`
	Password string `v:"required"`
}

type AddAdministratorRes struct {
}

type DeleteUserReq struct {
	g.Meta   `path:"/administrator/userDelete" method:"post" summary:""`
	Passport string `v:"required"`
}

type DeleteUserRes struct {
}

type UpdateUserReq struct {
	g.Meta   `path:"/administrator/userUpdate" method:"post" summary:""`
	Passport string `v:"required"`
	PhoneNum string `v:"required"`
}

type UpdateUserRes struct {
}
