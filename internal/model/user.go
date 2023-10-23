package model

type UserCreateInput struct {
	Passport string
	Password string
	Nickname string
	PhoneNum string
	CodeId   string
}

type UserSignInInput struct {
	Passport string
	Password string
}

type UserOut struct {
	Passport string `json:"passport"     description:""`
	Password string `json:"password"     description:""`
	PhoneNum string `json:"phoneNum"     description:""`
	Nickname string `json:"Nickname"     description:""`
}
type UserAdministratorInput struct {
	Passport string `json:"passport"     description:""`
	Password string `json:"password"     description:""`
}

type UserDeleteInput struct {
	Passport string `json:"passport"     description:""`
	PhoneNum string `json:"phoneNum"     description:""`
}

type UserUpdateInput struct {
	Passport string `json:"passport"     description:""`
	PhoneNum string `json:"phoneNum"     description:""`
}
