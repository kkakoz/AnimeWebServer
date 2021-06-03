package service

type CountUsecase struct {

}

func (c CountUsecase) AddAnimeView(animeId int64) error {
	panic("implement me")
}

func (c CountUsecase) UserLikeAnime(userId int64, animeId int64, likeType bool) error {
	panic("implement me")
}

func (c CountUsecase) UserUnLikeAnime(userId int64, animeId int64, unlikeType bool) error {
	panic("implement me")
}

func (c CountUsecase) GetAnimeView(animeIds []int64) error {
	panic("implement me")
}
