package usecase

import (
	"context"
	"gorm.io/gorm"
	"red-bean-anime-server/internal/app/count/domain"
)

type CountUsecase struct {
	countRepo domain.ICountRepo
	db        *gorm.DB
}

func (c CountUsecase) GetViewCount(ctx context.Context, id []int64) ([]domain.AnimeViewCountRes, error) {
	panic("implement me")
}

func (c CountUsecase) AddAnimeCount(ctx  context.Context,animeId int64) error {
	return c.countRepo.AddAnimeCount(ctx, animeId)
}

func (c CountUsecase) UpdateIncr() error {
	panic("implement me")
}

func (c CountUsecase) AddAnimeView(animeId int64) error {
	panic("implement me")
}

func (c CountUsecase) UserLikeAnime(userId int64, animeId int64, likeType bool) error {
	panic("implement me")
}

func (c CountUsecase) UserUnLikeAnime(userId int64, animeId int64, unlikeType bool) error {
	panic("implement me")
}

func (c CountUsecase) GetAnimeView(animeIds []int64) error {
	panic("implement me")
}

func NewCountUsecase(countRepo domain.ICountRepo, db *gorm.DB) domain.ICountUsecase {
	return &CountUsecase{countRepo: countRepo, db: db}
}
