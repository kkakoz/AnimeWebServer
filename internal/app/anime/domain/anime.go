package domain

import (
	"context"
	"gorm.io/gorm"
	"red-bean-anime-server/internal/pkg/query"
)

type Video struct {
	gorm.Model
	AnimeId int64  `json:"anime_id"`
	Episode int32  `json:"episode"`
	Name    string `json:"name"`
	Url     string `json:"url"`
}

type VideoInfo struct {
}

type AddVideo struct {
	AnimeId int64  `json:"anime_id"`
	Episode int32  `json:"episode"`
	Name    string `json:"name"`
	Url     string `json:"url"`
}

type Anime struct {
	gorm.Model
	Name          string `json:"name"`
	Description   string `json:"description"`
	Year          int32  `json:"year"`
	ImageUrl      string `json:"image_url"`
	Episode       int32  `json:"episode"`
	Quarter       int32  `json:"quarter"`
	FirstPlayTime string `json:"first_play_time"`
}

type AnimeInfoRes struct {
	gorm.Model
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
	AddVideo(ctx context.Context, video *AddVideo) error
	GetAnimeList(ctx context.Context, page *query.Page, req *AnimeListReq) ([]Anime, error)
	GetAnimeInfo(ctx, animeId int64, videoId int64) ([]AnimeInfoRes, error)
}

type IAnimeRepo interface {
	AddAnime(context.Context, *Anime) error
	AddVideo(ctx context.Context, video *Video) error
	GetAnimeList(ctx context.Context, page *query.Page, req int32) ([]Anime, error)
	GetAnimeListByCategoryId(ctx context.Context, page *query.Page, req *AnimeListReq) ([]Anime, error)
}
