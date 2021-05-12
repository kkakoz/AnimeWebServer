package usecase

import (
	"red-bean-anime-server/internal/app/user/domain"
	"context"
)

type UserUsecase struct {

}

func (u *UserUsecase) Login(ctx context.Context, phone string) (domain.UserInfo, error) {
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

