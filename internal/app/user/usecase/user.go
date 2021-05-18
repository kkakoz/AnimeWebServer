package usecase

import (
	"red-bean-anime-server/internal/app/user/domain"
	"context"
	"red-bean-anime-server/internal/app/user/repo"
)

type UserUsecase struct {
	userRepo *repo.UserRepo
}

func NewUserUsecase(userRepo *repo.UserRepo) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (u *UserUsecase) Login(ctx context.Context, phone, password string) (domain.UserInfo, error) {

	//u.userRepo.GetUserInfo()
	panic("implement me")
}

func (u *UserUsecase) CreateUser(ctx context.Context, phone string, name string) error {
	panic("implement me")
}

func (u *UserUsecase) GetUserInfo(ctx context.Context, id int) (domain.UserInfo, error) {
	panic("implement me")
}

func (u *UserUsecase) GetUserList(ctx context.Context, id []int) ([]domain.UserInfo, error) {
	panic("implement me")
}

