package domain

type ICountUsecase interface {
	AddAnimeView(animeId int64) error
	UserLikeAnime(userId int64, animeId int64, likeType bool) error
	UserUnLikeAnime(userId int64, animeId int64, unlikeType bool) error
	GetAnimeView(animeIds []int64) error
}

type ICountRepo interface {

}

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

