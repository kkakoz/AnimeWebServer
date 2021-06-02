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

func (c *CategoryRepo) AddCategory(ctx context.Context, category *domain.Category) error {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return err
	}
	err = db.Create(category).Error
	return errors.Wrap(err, "添加失败")
}

func (c *CategoryRepo) GetCategoryList(ctx context.Context) ([]domain.Category, error) {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return nil, err
	}
	categorys := make([]domain.Category, 0)
	err = db.Find(&categorys).Error
	return categorys, errors.Wrap(err, "添加失败")
}

func (c *CategoryRepo) GetByName(ctx context.Context, name string) (*domain.Category, error) {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return nil, err
	}
	category := &domain.Category{}
	err = db.Where("name = ?", name).First(&category).Error
	return nil, errors.Wrap(err, "查找失败")
}

func (c *CategoryRepo) AddAnimeCategory(ctx context.Context, animeId int64, categoryIds []int64) error {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return err
	}
	animeCategories := make([]domain.AnimeCategory, 0, len(categoryIds))
	for _, categoryId := range categoryIds {
		ac := domain.AnimeCategory{
			AnimeId:    animeId,
			CategoryId: categoryId,
		}
		animeCategories = append(animeCategories, ac)
	}
	err = db.CreateInBatches(&animeCategories, 100).Error
	return errors.Wrap(err, "添加动漫分类失败")
}