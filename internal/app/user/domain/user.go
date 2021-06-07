package domain

import (
	"context"
	"red-bean-anime-server/pkg/db/mysqlx"
)

type User struct {
	mysqlx.Model
	Name     string `json:"name" gorm:"size:24"`
	Email    string `json:"email" gorm:"size:24"`
	Password string `json:"password" gorm:"size:64"`
	Salt     string `json:"salt" gorm:"size:64"`
}

type UserInfo struct {
	ID        int64  `json:"id"`
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
