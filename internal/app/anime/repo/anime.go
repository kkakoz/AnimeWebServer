package repo

import (
	"context"
	"github.com/pkg/errors"
	"red-bean-anime-server/internal/app/anime/domain"
	"red-bean-anime-server/pkg/db/mysqlx"
)

type AnimeRepo struct {}

func NewAnimeRepo() domain.IAnimeRepo {
	return &AnimeRepo{}
}

func (a AnimeRepo) AddAnime(ctx context.Context, anime *domain.Anime) error {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return err
	}
	err = db.Create(anime).Error
	return errors.Wrap(err, "添加动漫失败")
}


