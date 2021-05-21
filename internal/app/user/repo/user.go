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
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return err
	}
	err = db.Create(user).Error
	return errors.Wrap(err, "添加用户失败")
}

func (u *UserRepo) GetUserInfo(ctx context.Context, id int) (*domain.User, error) {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return nil, err
	}
	user := &domain.User{}
	err = db.Where("id = ?", id).First(user).Error
	return user, errors.Wrap(err, "查找用户失败")
}

func (u *UserRepo) GetUserByPhone(ctx context.Context, phone string) (*domain.User, error) {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return nil, err
	}
	user := &domain.User{}
	err = db.Where("phone = ?", phone).First(user).Error
	return user, errors.Wrap(err, "查找用户失败")
}

func (u *UserRepo) GetUserList(ctx context.Context, id []int) ([]domain.User, error) {
	panic("implement me")
}
