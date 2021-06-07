package domain

import "context"

type ICountUsecase interface {
	AddAnimeView(animeId int64) error
	UserLikeAnime(userId int64, animeId int64, likeType bool) error
	UserUnLikeAnime(userId int64, animeId int64, unlikeType bool) error
	GetAnimeView(animeIds []int64) error
	UpdateIncr() error
	AddAnimeCount(ctx context.Context, animeId int64) error
	GetViewCount(ctx context.Context, id []int64) ([]AnimeViewCountRes, error)
}

type ICountRepo interface {
	// 添加到数据库
	AddAnimeCount(ctx context.Context, id int64) error
	// 添加增量到缓存
	AddView(ctx context.Context, animeId int64) error
	// 从缓存获取增量count
	GetAnimeIncrCountByCache(ctx context.Context, animeId int64) (*AnimeCount, bool)
	// 从缓存中取count
	GetAnimeCount(ctx context.Context, animeId int64) (*AnimeCount, error)
}

type AnimeCount struct {
	AnimeId      int64
	LikeCount    int
	ViewCount    int
	CollectCount int
}

type UserLike struct {
	UserId  int64
	AnimeId int64
}

type UserCollect struct {
	UserId  int64
	AnimeId int64
}

type AnimeViewCountRes struct {
	AnimeId   int64
	ViewCount int64
}

type CountCache struct {
	Key int64
	Val int
}