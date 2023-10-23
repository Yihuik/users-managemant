package user

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf-demo-user/v2/internal/dao"
	"github.com/gogf/gf-demo-user/v2/internal/model"
	"github.com/gogf/gf-demo-user/v2/internal/model/do"
	"github.com/gogf/gf-demo-user/v2/internal/model/entity"
	"github.com/gogf/gf-demo-user/v2/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"math/big"
	"time"
)

type (
	sUser struct{}
	//sCheckCode struct{}
)

func (s *sUser) CheckCode(ctx context.Context, in model.CheckIdInput) (err error) {
	//TODO implement me
	var checkCode *entity.CheckCode
	//now := time.Now().UnixNano()
	//end := checkCode.EndTime.UnixNano()
	err = dao.CheckCode.Ctx(ctx).Where(do.CheckCode{
		PhoneNumber: in.PhoneNumber,
		CodeId:      in.CodeId,
	}).Scan(&checkCode)
	/*if err != nil && checkCode == nil {
			if now < end {

			}
		}
		return
	}*/

	//tx := checkCode.EndTime
	/*location, _ := time.LoadLocation("Asia/Shanghai")
	  layout := "2006-01-02 15:04:05"
	  var t1, _ = time.ParseInLocation(layout, nw, location)
	  t2, _ := time.ParseInLocation(layout,tx, location)
	  var d = t2.Sub(t1)*/

	/*if now.Before(tx) && err != nil {
		return err
	}*/

	if checkCode == nil {
		return gerror.New(`Phone or CodeId not correct`)
	}
	return nil
}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

// Create creates user account
func (s *sUser) Create(ctx context.Context, in model.UserCreateInput) (err error) {
	// If Nickname is not specified, it then uses Passport as its default Nickname.
	if in.Nickname == "" {
		in.Nickname = in.Passport
	}
	var (
		available bool
	)

	// Passport checks.
	available, err = s.IsPassportAvailable(ctx, in.Passport)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`Passport "%s" is already token by others`, in.Passport)
	}
	// Nickname checks.
	available, err = s.IsNicknameAvailable(ctx, in.Nickname)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`Nickname "%s" is already token by others`, in.Nickname)
	}
	available, err = s.IsPhoneNumAvailable(ctx, in.PhoneNum)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`PhoneNum "%s" is already token by others`, in.PhoneNum)
	}
	available, err = s.IsCodeIdAvailable(ctx, in.CodeId, in.PhoneNum)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`CodeId "%s" is already token by others`, in.CodeId)
	}
	//************************************************************************
	return dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.User.Ctx(ctx).Data(do.User{
			Passport: in.Passport,
			Password: in.Password,
			Nickname: in.Nickname,
			PhoneNum: in.PhoneNum,
		}).Insert()
		return err
	})
}

func (s *sUser) Delete(ctx context.Context, passport string) (err error) {
	return dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.User.Ctx(ctx).Where(dao.User.Columns().Passport, passport).Delete()
		return err
	})
}

func (s *sUser) Update(ctx context.Context, in model.UserUpdateInput) (err error) {
	return dao.User.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.User.Ctx(ctx).Data(do.User{
			Passport: in.Passport,
			PhoneNum: in.PhoneNum,
		}).Update()
		return err
	})
}

func (s *sUser) GetUserList(ctx context.Context, in model.PageGet) (out []*model.UserOut, err error) {
	err = dao.User.Ctx(ctx).Limit((in.Page-1)*in.RowsPerPage, in.RowsPerPage).Scan(&out)
	if err != nil {
		return
	}
	return
}

type JwtClaims struct {
	*jwt.StandardClaims
	//用户编号
	UID      string
	Username string
}

// CreateJwtToken 生成一个token
func (s *sUser) CreateJwtToken(id string, passport string) (string, error) {
	var secret = []byte("WonderSafeBox")
	// 定义过期时间,7天后过期
	expireToken := time.Now().Add(time.Hour * 24 * 7).Unix()

	claims := JwtClaims{
		&jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // token信息生效时间
			ExpiresAt: int64(expireToken),              // 过期时间
			Issuer:    "wonders",                       // 发布者
		},
		id,
		passport,
	} // 对自定义claims加密
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 给这个token盐加密
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (s *sUser) SavePhoneNumber(ctx context.Context, in model.UserCreateInput) (err error) {
	return dao.CheckCode.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.CheckCode.Ctx(ctx).Data(do.CheckCode{
			PhoneNumber: in.PhoneNum,
		}).Insert()
		return err
	})

}

func (s *sUser) AddAdministrator(ctx context.Context, in model.UserAdministratorInput) (err error) {
	return dao.Administrator.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.Administrator.Ctx(ctx).Data(do.Administrator{
			AdministratorPassport: in.Passport,
			AdministratorPassword: in.Password,
		}).Insert()
		return err
	})
}

func (s *sUser) VerifyCodeSend(ctx context.Context, PhoneNumber string) (err error) {
	vCode := CreateRandomNumber(6)
	now := time.Now()
	var later = now.Add(time.Hour)
	fmt.Printf("current time:%v\n", now)
	return dao.CheckCode.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.CheckCode.Ctx(ctx).Data(do.CheckCode{
			PhoneNumber: PhoneNumber,
			CodeId:      vCode,
			EndTime:     later,
		}).Update()
		return err
	})
}

func CreateRandomNumber(len int) string {
	var numbers = []byte{1, 2, 3, 4, 5, 7, 8, 9}
	var container string
	length := bytes.NewReader(numbers).Len()

	for i := 1; i <= len; i++ {
		random, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {

		}
		container += fmt.Sprintf("%d", numbers[random.Int64()])
	}
	return container
}

// IsSignedIn checks and returns whether current user is already signed-in.
func (s *sUser) IsSignedIn(ctx context.Context) bool {
	if v := service.BizCtx().Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// SignIn creates session for given user account.
func (s *sUser) SignIn(ctx context.Context, in model.UserSignInInput) (err error) {
	var user *entity.User
	err = dao.User.Ctx(ctx).Where(do.User{
		Passport: in.Passport,
		Password: in.Password,
	}).Scan(&user)
	if err != nil {
		return err
	}
	if user == nil {
		return gerror.New(`Passport or Password not correct`)
	}
	if err = service.Session().SetUser(ctx, user); err != nil {
		return err
	}
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:       user.Id,
		Passport: user.Passport,
		Nickname: user.Nickname,
	})
	return nil
}

// SignOut removes the session for current signed-in user.
func (s *sUser) SignOut(ctx context.Context) error {
	return service.Session().RemoveUser(ctx)
}

// IsPassportAvailable checks and returns given passport is available for signing up.
func (s *sUser) IsPassportAvailable(ctx context.Context, passport string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Passport: passport,
	}).Count()
	dao.User.Ctx(ctx).Where(do.User{
		Passport: passport,
	}).Scan(passport)
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// IsNicknameAvailable checks and returns given nickname is available for signing up.
func (s *sUser) IsNicknameAvailable(ctx context.Context, nickname string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Nickname: nickname,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (s *sUser) IsPhoneNumAvailable(ctx context.Context, phoneNum string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{
		PhoneNum: phoneNum,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (s *sUser) IsCodeIdAvailable(ctx context.Context, codeId string, phoneNumber string) (bool, error) {
	count, err := dao.CheckCode.Ctx(ctx).Where(do.CheckCode{
		CodeId:      codeId,
		PhoneNumber: phoneNumber,
	}).Count()

	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// GetProfile retrieves and returns current user info in session.
func (s *sUser) GetProfile(ctx context.Context) *entity.User {
	return service.Session().GetUser(ctx)
}
