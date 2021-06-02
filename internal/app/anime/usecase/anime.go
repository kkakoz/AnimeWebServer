package usecase

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"red-bean-anime-server/internal/app/anime/domain"
	"red-bean-anime-server/internal/pkg/query"
	"red-bean-anime-server/pkg/db/mysqlx"
)

type AnimeUsecase struct {
	db           *gorm.DB
	animeRepo    domain.IAnimeRepo
	categoryRepo domain.ICategoryRepo
}

func (a *AnimeUsecase) GetAnimeInfo(ctx, animeId int64, videoId int64) ([]domain.AnimeInfoRes, error) {
	panic("implement me")
}

func (a *AnimeUsecase) GetAnimeList(ctx context.Context, page *query.Page, req *domain.AnimeListReq) ([]domain.Anime, error) {
	if req.CategoryId != 0 {
		animes, err := a.animeRepo.GetAnimeListByCategoryId(ctx, page, req)
		return animes, err
	}
	return a.animeRepo.GetAnimeList(ctx, page, req.Sort)
}

func NewAnimeUsecase(db *gorm.DB, animeRepo domain.IAnimeRepo, categoryRepo domain.ICategoryRepo) domain.IAnimeUsecase {
	return &AnimeUsecase{db: db, animeRepo: animeRepo, categoryRepo: categoryRepo}
}

func (a *AnimeUsecase) AddAnime(ctx context.Context, addAnime *domain.AddAnime) error {
	ctx, tx := mysqlx.Begin(ctx, a.db)
	anime := &domain.Anime{
		Name:        addAnime.Name,
		Description: addAnime.Description,
		Year:        addAnime.Year,
		Quarter:     addAnime.Quarter,
	}
	err := a.animeRepo.AddAnime(ctx, anime)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = a.categoryRepo.AddAnimeCategory(ctx, int64(anime.ID), addAnime.CategoryIds)
	if err != nil {
		tx.Rollback()
		return err
	}
	return errors.Wrap(tx.Commit().Error, "添加动漫失败")
}

func (a *AnimeUsecase) AddVideo(ctx context.Context, addVideo *domain.AddVideo) error {
	video := &domain.Video{
		AnimeId: addVideo.AnimeId,
		Episode: addVideo.Episode,
		Name:    addVideo.Name,
		Url:     addVideo.Url,
	}
	err := a.animeRepo.AddVideo(ctx, video)
	return err
}