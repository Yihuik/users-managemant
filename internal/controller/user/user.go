package user

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf-demo-user/v2/api/user/v1"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/service"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// SignUp is the API for user sign up.
func (c *Controller) SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error) {
	err = service.User().Create(ctx, model.UserCreateInput{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
		PhoneNum: req.PhoneNum,
	})
	return
}

// SignIn is the API for user sign in.
func (c *Controller) SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error) {
	err = service.User().SignIn(ctx, model.UserSignInInput{
		Passport: req.Passport,
		Password: req.Password,
	})
	return
}

// IsSignedIn checks and returns whether the user is signed in.
func (c *Controller) IsSignedIn(ctx context.Context, req *v1.IsSignedInReq) (res *v1.IsSignedInRes, err error) {
	res = &v1.IsSignedInRes{
		OK: service.User().IsSignedIn(ctx),
	}
	return
}

// 用户登录

// SignOut is the API for user sign out.
func (c *Controller) SignOut(ctx context.Context, req *v1.SignOutReq) (res *v1.SignOutRes, err error) {
	err = service.User().SignOut(ctx)
	return
}

// CheckPassport checks and returns whether the user passport is available.
func (c *Controller) CheckPassport(ctx context.Context, req *v1.CheckPassportReq) (res *v1.CheckPassportRes, err error) {
	available, err := service.User().IsPassportAvailable(ctx, req.Passport)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf(`Passport "%s" is already token by others`, req.Passport)
	}
	return
}

func (c *Controller) CheckNum(ctx context.Context, req *v1.CheckNumReq) (res *v1.CheckNumRes, err error) {
	available, err := service.User().IsPhoneNumAvailable(ctx, req.PhoneNum)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf(`PhoneNum  is already token by others`)
	}
	var _ = service.User().VerifyCodeSend(ctx, req.PhoneNum)
	return
}

func (c *Controller) CheckCodeId(ctx context.Context, req *v1.CheckCodeIdReq) (res *v1.CheckCodeIdRes, err error) {
	available, err := service.User().IsCodeIdAvailable(ctx, req.CodeId, req.PhoneNum)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf(`CodeId  is incorrect`)
	}

	return
}

func (c *Controller) CheckNickName(ctx context.Context, req *v1.CheckNickNameReq) (res *v1.CheckNickNameRes, err error) {
	available, err := service.User().IsNicknameAvailable(ctx, req.Nickname)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, gerror.Newf(`Nickname "%s" is already token by others`, req.Nickname)
	}
	return
}

// Profile returns the user profile.
func (c *Controller) Profile(ctx context.Context, req *v1.ProfileReq) (res *v1.ProfileRes, err error) {
	res = &v1.ProfileRes{
		User: service.User().GetProfile(ctx),
	}
	return
}

func (c *Controller) GetUserList(ctx context.Context, req *v1.GetUserListReq) (res v1.GetUserListRes, err error) {
	out, err := service.User().GetUserList(ctx, model.PageGet{
		Page:        req.Page,
		RowsPerPage: req.RowsPerPage,
	})
	if err != nil {
		return
	}
	fmt.Println(out)

	var userList []*v1.UserRes

	if err = gconv.Scan(out, &userList); err != nil {
		return
	}

	res.UserList = userList
	return
}
func (c *Controller) AddAdministrator(ctx context.Context, req *v1.AddAdministratorReq) (res v1.AddAdministratorRes, err error) {
	err = service.User().Create(ctx, model.UserCreateInput{
		Passport: req.Passport,
		Password: req.Password,
	})
	return
}

func (c *Controller) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (res v1.DeleteUserRes, err error) {
	err = service.User().Delete(ctx, req.Passport)
	return
}

func (c *Controller) UserUpdate(ctx context.Context, req *v1.UpdateUserReq) (res v1.UpdateUserRes, err error) {
	err = service.User().Update(ctx, model.UserUpdateInput{
		Passport: req.Passport,
		PhoneNum: req.PhoneNum,
	})
	return
}
