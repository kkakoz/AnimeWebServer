package keys

import "fmt"

const (
	// incr hash pre key
	HAnimeIncrCountKey = "anime:incr:"
	// incr hash field pre key
	HFAnimeIncrView    = "view"
	HFAnimeIncrLike    = "like"
	HFAnimeIncrCollect = "collect"

	AnimeBloomKey = "anime:bit:"
	// view hash key
	AnimeViewKey = "anime:view"
	// like and collect pre key
	AnimeLikeCollectCountKey = "anime:count:"
)

var HFAnimeIncrs = []string{HFAnimeIncrView, HFAnimeIncrLike, HFAnimeIncrView}

func GetHAnimeCountIncrKey(animeId int64) string {
	return fmt.Sprintf("%s%d", HAnimeIncrCountKey, animeId)
}

func GetAnimeBloomKey(animeId int64) string {
	return fmt.Sprintf("%s%d", AnimeBloomKey, animeId)
}

func GetAnimeViewKey() string {
	return AnimeViewKey
}