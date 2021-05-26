package usecase

import (
	"context"
	"gorm.io/gorm"
	"red-bean-anime-server/internal/app/anime/domain"
)

type AnimeUsecase struct {
	db        *gorm.DB
	animeRepo domain.IAnimeRepo
}

func NewAnimeUsecase(db *gorm.DB, animeRepo domain.IAnimeRepo) domain.IAnimeUsecase {
	return &AnimeUsecase{db: db, animeRepo: animeRepo}
}

func (a AnimeUsecase) AddAnime(ctx context.Context, req domain.AddAnimeReq) error {
	anime := &domain.Anime{
		Name:        req.Name,
		Description: req.Description,
		Year:        req.Year,
		Quarter:     req.Quarter,
	}
	return a.animeRepo.AddAnime(ctx, anime)
}


