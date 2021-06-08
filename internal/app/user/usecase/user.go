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
	"time"
)

type UserUsecase struct {
	userRepo    domain.IUserRepo
	jwtTokenGen *auth.JwtTokenGen
	db          *gorm.DB
}

func NewUserUsecase(userRepo domain.IUserRepo, jwtTokenGen *auth.JwtTokenGen, db *gorm.DB) domain.IUserUsecase {
	return &UserUsecase{userRepo: userRepo, jwtTokenGen: jwtTokenGen, db: db}
}

func (u *UserUsecase) Login(ctx context.Context, email, password string) (*domain.UserInfo, error) {
	user, err := u.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gerrors.NewBusErr("未找到用户")
		}
		return nil, err
	}
	token, err := u.jwtTokenGen.GenTokenExpire(user.ID, user.Name, user.State, user.Auth, time.Hour*24*3)
	if err != nil {
		return nil, err
	}
	if cryption.Md5Str(password+user.Salt) != user.Password {
		return nil, gerrors.NewBusErr("密码错误")
	}
	info := &domain.UserInfo{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Token:     token,
		CreatedAt: times.FormatYMD(user.CreatedAt),
	}
	return info, nil
}

func (u *UserUsecase) Register(ctx context.Context, email, name, password string) error {
	_, err := u.userRepo.GetUserByEmail(ctx, email)
	if err == nil { // 如果找到用户
		return gerrors.NewBusErr("该邮箱已经注册")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) { // 如果找到了用户
		return err // 其他err
	}
	salt := cryption.UUID()
	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: cryption.Md5Str(password + salt),
		Salt:     salt,
		Auth:     domain.UserAuthNromal,
		State:    domain.UserStateNormal,
	}
	err = u.userRepo.AddUser(ctx, user)
	return err
}

func (u *UserUsecase) GetUserInfo(ctx context.Context, id int) (*domain.UserInfo, error) {
	panic("implement me")
}

func (u *UserUsecase) GetUserList(ctx context.Context, id []int) ([]domain.UserInfo, error) {
	panic("implement me")
}
