package usecase

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"red-bean-anime-server/internal/app/user/domain"
	"red-bean-anime-server/pkg/auth"
	"red-bean-anime-server/pkg/cryption"
	"red-bean-anime-server/pkg/gerrors"
	"red-bean-anime-server/pkg/times"
	"strconv"
	"time"
)

type UserUsecase struct {
	userRepo    domain.IUserRepo
	jwtTokenGen *auth.JwtTokenGen
	db          *gorm.DB
}

func NewUserUsecase(userRepo domain.IUserRepo, jwtTokenGen *auth.JwtTokenGen, db *gorm.DB) *UserUsecase {
	return &UserUsecase{userRepo: userRepo, jwtTokenGen: jwtTokenGen, db: db}
}

func (u *UserUsecase) Login(ctx context.Context, phone, password string) (*domain.UserInfo, error) {
	user, err := u.userRepo.GetUserByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	token, err := u.jwtTokenGen.GenTokenExpire(strconv.Itoa(int(user.ID)), time.Hour*24*3)
	if err != nil {
		return nil, err
	}
	if cryption.Md5Str(password+user.Salt) != user.Password {
		return nil, gerrors.NewBusErr("密码错误")
	}
	info := &domain.UserInfo{
		ID:        int(user.ID),
		Name:      user.Name,
		Phone:     user.Phone,
		Token:     token,
		CreatedAt: times.FormatYMD(user.CreatedAt),
		UpdatedAt: times.FormatYMD(user.UpdatedAt),
	}
	return info, nil
}

func (u *UserUsecase) Register(ctx context.Context, phone, name, password string) error {
	_, err := u.userRepo.GetUserByPhone(ctx, phone)
	if err == nil { // 如果找到用户
		return gerrors.NewBusErr("该手机号已经注册")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) { // 如果未找到了用户
		salt := cryption.UUID()
		user := &domain.User{
			Name:     name,
			Phone:    phone,
			Password: cryption.Md5Str(password + salt),
			Salt:     salt,
		}
		err := u.userRepo.AddUser(ctx, user)
		return err
	}
	return err   // 其他err
}

func (u *UserUsecase) GetUserInfo(ctx context.Context, id int) (*domain.UserInfo, error) {
	panic("implement me")
}

func (u *UserUsecase) GetUserList(ctx context.Context, id []int) ([]domain.UserInfo, error) {
	panic("implement me")
}
