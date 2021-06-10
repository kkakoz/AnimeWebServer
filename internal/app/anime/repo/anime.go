package repo

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"red-bean-anime-server/internal/app/anime/domain"
	"red-bean-anime-server/internal/pkg/query"
	"red-bean-anime-server/pkg/db/mysqlx"
)

type AnimeRepo struct {}

func (a *AnimeRepo) GetAnimeById(ctx context.Context, id int64) (*domain.AnimeInfoRes, error) {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return nil, err
	}
	res := &domain.AnimeInfoRes{}
	err = db.Model(&domain.Anime{}).Select("animes.id, animes.name, animes.description, animes.year, " +
		"animes.image_url, animes.episode, animes.quarter, animes.first_play_time, DATE_FORMAT(animes.updated_at,'%Y %h:%i') as updated_at").
		Where("id = ?", id).First(res).Error
	return res, errors.Wrap(err, "查找失败")
}

func (a *AnimeRepo) GetAnimeExistByName(ctx context.Context, name string) (bool, error) {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return false, err
	}
	anime := &domain.Anime{}
	err = db.Where("name = ?", name).First(anime).Error
	if err == nil {
		return true, nil
	}
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return false, errors.Wrap(err, "查找动漫失败")
}

func (a *AnimeRepo) GetAnimeListByCategoryId(ctx context.Context, page *query.Page, req *domain.AnimeListReq)  ([]domain.Anime, error) {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return nil, err
	}
	db = page.LimitPage(db)
	animes := make([]domain.Anime, 0)
	if req.Sort == domain.SortByTime {
		db = db.Order("order by animes.update_time desc")
	}
	err = db.Select("animes.*").Table("(select * from categories where category_id = ? and deleted_at is null ) as ac", req.CategoryId).
		Joins("left join animes on anime.id id = ac.anime_id").Find(&animes).Error
	return animes, errors.Wrap(err, "查找分类下的动漫失败")
}

func (a *AnimeRepo) GetAnimeList(ctx context.Context, page *query.Page, sort int32) ([]domain.Anime, error) {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return nil, err
	}
	db = page.LimitPage(db)
	if sort == domain.SortByTime {
		db = db.Order("order by animes.update_time desc")
	}
	animes := make([]domain.Anime, 0)
	err = db.Find(&animes).Error
	return animes, errors.Wrap(err, "查找动漫失败")
}

func (a *AnimeRepo) AddVideo(ctx context.Context, video *domain.Video) error {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return err
	}
	err = db.Create(video).Error
	return errors.Wrap(err, "添加视频失败")
}

func (a *AnimeRepo) AddAnime(ctx context.Context, anime *domain.Anime) error {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return err
	}
	err = db.Create(anime).Error
	return errors.Wrap(err, "添加动漫失败")
}

func NewAnimeRepo() domain.IAnimeRepo {
	return &AnimeRepo{}
}