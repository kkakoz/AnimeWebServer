package keys

import "fmt"

const (
	AnimeIncrCountKey = "anime:incr:"


	AnimeView = "anime:view:"
	AnimeLike = "anime:like"
	AnimeCollect = "anime:collect"
	
	AnimeBloomKey = "anime:bit:"
)

func GetAnimeCountIncrKey(animeId int64) string {
	return fmt.Sprintf("%s%d", AnimeIncrCountKey, animeId)
}

func GetAnimeViewKey() string {
	return AnimeView
}

func GetAnimeLikeKey() string {
	return AnimeLike
}

func GetAnimeCollectKey() string {
	return AnimeCollect
}

func GetAnimeBloomKey(animeId int64) string {
	return fmt.Sprintf("%s%d", AnimeBloomKey, animeId)
}