package domain

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"size:20"`
	Email    string `json:"email" gorm:"size:50"`
	Password string `json:"password" gorm:"size:50"`
	Salt     string `json:"salt" gorm:"size:50"`
}

type UserInfo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Token     string `json:"token"`
	CreatedAt string `json:"created_at"`
}

type IUserUsecase interface {
	Login(ctx context.Context, email, password string) (*UserInfo, error)
	Register(ctx context.Context, phone, name, password string) error
	GetUserInfo(ctx context.Context, id int) (*UserInfo, error)
	GetUserList(ctx context.Context, id []int) ([]UserInfo, error)
}

type IUserRepo interface {
	AddUser(ctx context.Context, user *User) error
	GetUserInfo(ctx context.Context, id int) (*User, error)
	GetUserList(ctx context.Context, id []int) ([]User, error)
	GetUserByEmail(ctx context.Context, phone string) (*User, error)
}
