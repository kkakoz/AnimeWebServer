package domain

import (
	"context"
	"red-bean-anime-server/internal/pkg/query"
	"red-bean-anime-server/pkg/db/mysqlx"
)

type Video struct {
	mysqlx.Model
	AnimeId int64  `json:"anime_id"`
	Episode int32  `json:"episode"` // 第几集
	Name    string `json:"name" gorm:"size:20;"`
	Url     string `json:"url" gorm:"size:255;"`
}

type Anime struct {
	mysqlx.Model
	Name          string `json:"name" gorm:"size:20;"`
	Description   string `json:"description" gorm:"size:255;"`
	Year          int32  `json:"year"`
	ImageUrl      string `json:"image_url" gorm:"size:255;"`
	Episode       int32  `json:"episode"`  // 上传了几集
	Quarter       int32  `json:"quarter"`  // 季度
	FirstPlayTime string `json:"first_play_time" gorm:"size:50;"`
}

type AnimeInfoRes struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Year          int32   `json:"year"`
	Quarter       int32   `json:"quarter"`
	FirstPlayTime string  `json:"first_play_time"`
	LikeCount     int32   `json:"like_count"`
	CollectCount  int32   `json:"collect_count"`
	UpdatedAt     string  `json:"updated_at"`
	Like          bool    `json:"like"`
	Collect       bool    `json:"collect"`
	Videos        []Video `json:"videos" gorm:"-"`
}

type AddAnime struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Year          int32   `json:"year"`
	Quarter       int32   `json:"quarter"`
	FirstPlayTime string  `json:"first_play_time"`
	CategoryIds   []int64 `json:"categories"`
	UpdateTime    string  `json:"update_time"`
}

type AnimeListReq struct {
	CategoryId int64
	Sort       int32
}

const (
	SortByTime = 0
	SortByHot  = 1
)

type IAnimeUsecase interface {
	AddAnime(ctx context.Context, req *AddAnime) error
	AddVideo(ctx context.Context, video *Video) error
	GetAnimeList(ctx context.Context, page *query.Page, req *AnimeListReq) ([]Anime, error)
	GetAnimeInfo(ctx context.Context, animeId int64) (*AnimeInfoRes, error)
	UserLikeAnime(ctx context.Context, animeId int64, likeType bool) error
	UserUnLikeAnime(ctx context.Context, id int64, likeType bool) error
}

type IAnimeRepo interface {
	AddAnime(context.Context, *Anime) error
	AddVideo(ctx context.Context, video *Video) error
	GetAnimeList(ctx context.Context, page *query.Page, req int32) ([]Anime, error)
	GetAnimeListByCategoryId(ctx context.Context, page *query.Page, req *AnimeListReq) ([]Anime, error)
	GetAnimeExistByName(ctx context.Context, name string) (bool, error)
	GetAnimeById(ctx context.Context, id int64) (*AnimeInfoRes, error)
}

type IVideoRepo interface {
	AddVideo(ctx context.Context, video *Video) error
	GetVideoByAnimeId(ctx context.Context, animeId int64) ([]Video, error)
	GetExistByAnimeIdEspisode(ctx context.Context, id int64, episode int32) (bool, error)
}
