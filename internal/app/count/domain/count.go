package domain

import "context"

type ICountUsecase interface {
	// 添加view incr缓存
	AddAnimeView(animeId int64) error
	// 添加点赞
	UserLikeAnime(userId int64, animeId int64, likeType bool) error
	// 取消点赞
	UserUnLikeAnime(userId int64, animeId int64, unlikeType bool) error
	// 更新incr到缓存
	UpdateIncr(ctx context.Context, )
	// 添加AnimeCount到数据库
	AddAnimeCount(ctx context.Context, animeId int64) error
	// 获取view count
	GetViewCount(ctx context.Context, id []int64) ([]AnimeViewCountRes, error)
	// 获取视频的点赞和收藏信息
	GetAnimeLikeCollect(ctx context.Context, userId int64, animeId int64) (*UserLikeCollectCountRes, error)
}

type ICountRepo interface {
	// 添加到数据库
	AddAnimeCount(ctx context.Context, id int64) error
	// 添加增量到缓存
	AddCount(animeId int64, fileKey string) error
	// 减少对应的点赞或者收藏
	DelCount(animeId int64, fileKey string) error
	// 从缓存获取增量count
	GetAnimeIncrCountByCache(ctx context.Context, animeId int64) (*AnimeCount, bool)
	// 从数据库中取view count
	GetAnimeViewCount(ctx context.Context, animeIds []int64) ([]AnimeViewCountRes, error)
	// 从数据库去like collect count
	GetAnimeLikeCollectCount(ctx context.Context, animeIds int64) (*AnimeLikeCollectCountRes, error)
	// 获取增量的的id
	GetAnimeIncrIds(ctx context.Context, cur uint64) ([]int64, uint64, error)
	// 从cache获取播放量
	GetAnimeViewByCache(ctx context.Context, animeIds []int64) ([]CountCache, []int64)
	// 设置cache播放量
	SetAnimeViewCache(ctx context.Context, counts []AnimeViewCountRes) error
	// 设置like和collect缓存
	SetAnimeLikeCollectCache(ctx context.Context, animeId int64, likeCount, collectCount int) error
	// 获取like collect 缓存
	GetAnimeLikeCollectCache(ctx context.Context, animeId int64) (*AnimeLikeCollectCountRes, bool)
	// 增量持久化到缓存
	IncrAnimeCount(ctx context.Context, incrCount *AnimeCount) error
	// 减少已经增加到缓存的计数
	ReduceCache(cache *AnimeCount)
	// 修改后删除对应的缓存
	RemoveCountCache(id int64) error
}

type ILikeRepo interface {
	// 保存用户的点赞
	AddUserLike(ctx context.Context, like *UserLike) error
	// 添加到布隆过滤器
	AddBloomFilter(userId int64, animeId int64) error
	// 通过布隆过滤器查找是否存在点赞信息
	GetLikeByBloomFilter(userId int64, animeId int64) bool
	// 删除用户点赞
	DelUserLike(ctx context.Context, like *UserLike) error
	// 获取用户是否点赞
	GetUserLike(ctx context.Context, userId int64, animeId int64) (bool, error)
}

type ICountCacheRepo interface {
}

type AnimeCount struct {
	AnimeId      int64 `gorm:"primarykey"`
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
	ViewCount int
}

type AnimeLikeCollectCountRes struct {
	AnimeId      int64
	LikeCount    int
	CollectCount int
}

type UserLikeCollectCountRes struct {
	AnimeId       int64
	LikeCount     int
	CollectCount  int
	UserLiked     bool
	UserCollected bool
}

type CountCache struct {
	Key int64
	Val int
}
