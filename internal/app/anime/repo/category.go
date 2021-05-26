package repo

import (
	"context"
	"github.com/pkg/errors"
	"red-bean-anime-server/internal/app/anime/domain"
	"red-bean-anime-server/pkg/db/mysqlx"
)

type CategoryRepo struct {
	
}

func NewCategoryRepo() domain.ICategoryRepo {
	return &CategoryRepo{}
}

func (c CategoryRepo) AddCategory(ctx context.Context, category *domain.Category) error {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return err
	}
	err = db.Create(category).Error
	return errors.Wrap(err, "添加失败")
}

func (c CategoryRepo) GetCategoryList(ctx context.Context) ([]domain.Category, error) {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return nil, err
	}
	categorys := make([]domain.Category, 0)
	err = db.Find(&categorys).Error
	return nil, errors.Wrap(err, "添加失败")
}

