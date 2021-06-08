package keys

import "fmt"

const (
	// incr hash pre key
	HAnimeIncrCountKey = "anime:incr:"
	// incr hash field pre key
	HFAnimeIncrView    = "view"
	HFAnimeIncrLike    = "like"
	HFAnimeIncrCollect = "collect"


	// view hash key
	HAnimeViewKey = "anime:view"
	// like and collect pre key
	HAnimeLikeCollectCountKey = "anime:count:"
	// like and collect hash field
	HFLikeKey = "like"
	HFCollectKey = "collect"
	// anime bloom filter key
	AnimeLikeBloomKey = "anime:like:bit:"
	AnimeCollectBloomKey = "anime:collect:bit:"
)

var HFAnimeIncrs = []string{HFAnimeIncrView, HFAnimeIncrLike, HFAnimeIncrView}

func GetHAnimeCountIncrKey(animeId int64) string {
	return fmt.Sprintf("%s%d", HAnimeIncrCountKey, animeId)
}

func GetAnimeLikeBloomKey(animeId int64) string {
	return fmt.Sprintf("%s%d", AnimeLikeBloomKey, animeId)
}

func GetAnimeCollectBloomKey(animeId int64) string {
	return fmt.Sprintf("%s%d", AnimeCollectBloomKey, animeId)
}

func GetAnimeCountKey(animeId int64) string {
	return fmt.Sprintf("%s%d", HAnimeLikeCollectCountKey, animeId)
}