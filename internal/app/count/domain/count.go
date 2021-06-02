package service

type AnimeCount struct {
	AnimeId      int64
	LikeCount    int64
	ViewCount    int64
	CollectCount int64
}

type UserLike struct {
	UserId  int64
	AnimeId int64
}

type UserCollect struct {
	UserId  int64
	AnimeId int64
}

type ICountUsecase interface {
	AddAnimeView(animeId int) error
	UserLikeAnime(animeId int, likeType bool) error

}

type ICountRepo interface {
	ViewCacheIncr(animeId int) error
}
