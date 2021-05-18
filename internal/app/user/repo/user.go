package repo

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"red-bean-anime-server/internal/app/user/domain"
	"red-bean-anime-server/pkg/db/mysqlx"
)

type UserRepo struct {
	redis *redis.Client
}

func NewUserRepo(redis *redis.Client) *UserRepo {
	return &UserRepo{redis: redis}
}

func (u *UserRepo) AddUser(ctx context.Context, user *domain.User) error {
	db := mysqlx.GetDB(ctx)
	err := db.Create(user).Error
	return errors.Wrap(err, "添加用户失败")
}

func (u *UserRepo) GetUserInfo(ctx context.Context, id int) (domain.UserInfo, error) {
	panic("implement me")
}

func (u *UserRepo) GetUserList(ctx context.Context, id []int) ([]domain.UserInfo, error) {
	panic("implement me")
}
