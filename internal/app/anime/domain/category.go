package domain

import (
	"context"
	"red-bean-anime-server/pkg/db/mysqlx"
)

type Category struct {
	mysqlx.Model
	Name string `json:"name"`
}

type AnimeCategory struct {
	AnimeId    int64 `json:"anime_id"`
	CategoryId int64 `json:"category_id"`
}

type ICategoryUsecase interface {
	AddCategory(ctx context.Context, name string) error
	GetCategoryList(ctx context.Context, ) ([]Category, error)
}

type ICategoryRepo interface {
	AddCategory(ctx context.Context, category *Category) error
	GetCategoryList(ctx context.Context) ([]Category, error)
	GetByName(ctx context.Context, name string) (*Category, error)
	AddAnimeCategory(ctx context.Context, id int64, categoryIds []int64) error
}
