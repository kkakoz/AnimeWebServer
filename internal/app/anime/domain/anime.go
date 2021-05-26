package domain

import (
	"context"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Episode int    `json:"episode"`
	Name    int    `json:"name"`
	Url     string `json:"url"`
}

type Anime struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Year        int32  `json:"year"`
	Quarter     int32  `json:"quarter"`
}

type AddAnimeReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Year        int32  `json:"year"`
	Quarter     int32  `json:"quarter"`
}

type IAnimeUsecase interface {
	AddAnime(ctx context.Context, req AddAnimeReq) error
}

type IAnimeRepo interface {
	AddAnime(context.Context, *Anime) error
}
