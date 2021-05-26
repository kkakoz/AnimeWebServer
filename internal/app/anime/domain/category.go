package domain

import (
	"context"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

type CategoryRes struct {

}

type ICategoryUsecase interface {
	AddCategory(ctx context.Context,name string) error
	GetCategoryList(ctx context.Context, ) ([]Category, error)
}

type ICategoryRepo interface {
	AddCategory(ctx context.Context, category *Category) error
	GetCategoryList(ctx context.Context,) ([]Category, error)
}