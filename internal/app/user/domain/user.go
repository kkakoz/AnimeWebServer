package domain

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

type UserInfo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Token     string `json:"token"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type IUserUsecase interface {
	Login(ctx context.Context, phone, password string) (*UserInfo, error)
	Register(ctx context.Context, phone, name, password string) error
	GetUserInfo(ctx context.Context, id int) (*UserInfo, error)
	GetUserList(ctx context.Context, id []int) ([]UserInfo, error)
}

type IUserRepo interface {
	AddUser(ctx context.Context, user *User) error
	GetUserInfo(ctx context.Context, id int) (*User, error)
	GetUserList(ctx context.Context, id []int) ([]User, error)
	GetUserByPhone(ctx context.Context, phone string) (*User, error)
}
