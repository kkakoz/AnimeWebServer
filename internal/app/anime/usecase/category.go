package usecase

import (
	"context"
	"gorm.io/gorm"
	"red-bean-anime-server/internal/app/anime/domain"
)

type CategoryUsecase struct {
	db           *gorm.DB
	categoryRepo domain.ICategoryRepo
}

func NewCategoryUsecase(db *gorm.DB, categoryRepo domain.ICategoryRepo) domain.ICategoryUsecase {
	return &CategoryUsecase{db: db, categoryRepo: categoryRepo}
}

func (c *CategoryUsecase) AddCategory(ctx context.Context, name string) error {
	category := &domain.Category{
		Name: name,
	}
	return c.categoryRepo.AddCategory(ctx, category)
}

func (c *CategoryUsecase) GetCategoryList(ctx context.Context) ([]domain.Category, error) {
	return c.categoryRepo.GetCategoryList(ctx)
}
