package domain

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type UserInfo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Token     string `json:"token"`
	CreatedAt string
	UpdatedAt string
}

type IUserUsecase interface {
	Login(ctx context.Context, phone string) (UserInfo, error)
	CreateUser(ctx context.Context, phone string, name string) error
	GetUserInfo(ctx context.Context, id int) (UserInfo, error)
	GetUserList(ctx context.Context, id []int) ([]UserInfo, error)
}

type IUserRepo interface {
	AddUser(ctx context.Context, user *User) error
	GetUserInfo(ctx context.Context, id int) (UserInfo, error)
	GetUserList(ctx context.Context, id []int) ([]UserInfo, error)
}
