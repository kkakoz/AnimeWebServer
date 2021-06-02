package usecase

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"red-bean-anime-server/internal/app/anime/domain"
	"red-bean-anime-server/pkg/gerrors"
)

type CategoryUsecase struct {
	db           *gorm.DB
	categoryRepo domain.ICategoryRepo
}

func NewCategoryUsecase(db *gorm.DB, categoryRepo domain.ICategoryRepo) domain.ICategoryUsecase {
	return &CategoryUsecase{db: db, categoryRepo: categoryRepo}
}

func (c *CategoryUsecase) AddCategory(ctx context.Context, name string) error {
	_, err := c.categoryRepo.GetByName(ctx, name)
	if err == nil { // 找到了
		return gerrors.NewBusErr("该分类已存在")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) { // 不是未找到的错误
		return err
	}
	err = nil
	category := &domain.Category{
		Name: name,
	}
	return c.categoryRepo.AddCategory(ctx, category)
}

func (c *CategoryUsecase) GetCategoryList(ctx context.Context) ([]domain.Category, error) {
	return c.categoryRepo.GetCategoryList(ctx)
}
