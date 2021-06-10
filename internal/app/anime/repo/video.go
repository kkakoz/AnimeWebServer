package repo

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"red-bean-anime-server/internal/app/anime/domain"
	"red-bean-anime-server/pkg/db/mysqlx"
)

type VideoRepo struct {}

func (v *VideoRepo) GetExistByAnimeIdEspisode(ctx context.Context, animeId int64, episode int32) (bool, error) {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return false, err
	}
	video := &domain.Video{}
	err = db.Model(video).Where("anime_id = ? and episode = ?", animeId, episode).First(video).Error
	if err == nil {
		return true, nil
	}
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return false, errors.Wrap(err, "查找失败")
}

func (v *VideoRepo) AddVideo(ctx context.Context, video *domain.Video) error {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return err
	}
	err = db.Create(video).Error
	return errors.Wrap(err, "添加视频失败")
}

func (v *VideoRepo) GetVideoByAnimeId(ctx context.Context, animeId int64) ([]domain.Video, error) {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return nil, err
	}
	videos := make([]domain.Video, 0)
	err = db.Where("anime_id = ?", animeId).Find(&videos).Error
	return videos, errors.Wrap(err, "获取视频失败")
}

func NewVideoRepo() domain.IVideoRepo {
	return &VideoRepo{}
}

