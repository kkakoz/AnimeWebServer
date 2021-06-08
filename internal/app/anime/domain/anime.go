package domain

import (
	"context"
	"red-bean-anime-server/internal/pkg/query"
	"red-bean-anime-server/pkg/db/mysqlx"
)

type Video struct {
	mysqlx.Model
	AnimeId int64  `json:"anime_id"`
	Episode int32  `json:"episode"`
	Name    string `json:"name" gorm:"size:20;"`
	Url     string `json:"url" gorm:"size:255;"`
}

type Anime struct {
	mysqlx.Model
	Name          string `json:"name" gorm:"size:20;"`
	Description   string `json:"description" gorm:"size:255;"`
	Year          int32  `json:"year"`
	ImageUrl      string `json:"image_url" gorm:"size:255;"`
	Episode       int32  `json:"episode"`
	Quarter       int32  `json:"quarter"`
	FirstPlayTime string `json:"first_play_time" gorm:"size:50;"`
}

type AnimeInfoRes struct {
	mysqlx.Model
	Name          string `json:"name"`
	Description   string `json:"description"`
	Year          int32  `json:"year"`
	Quarter       int32  `json:"quarter"`
	FirstPlayTime string `json:"first_play_time"`
	VideoUrl      string `json:"video_url"`
	Videos        Video  `json:"videos"`
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
	GetAnimeInfo(ctx context.Context, animeId int64, videoId int64) ([]AnimeInfoRes, error)
	UserLikeAnime(ctx context.Context, animeId int64, likeType bool) error
}

type IAnimeRepo interface {
	AddAnime(context.Context, *Anime) error
	AddVideo(ctx context.Context, video *Video) error
	GetAnimeList(ctx context.Context, page *query.Page, req int32) ([]Anime, error)
	GetAnimeListByCategoryId(ctx context.Context, page *query.Page, req *AnimeListReq) ([]Anime, error)
}
