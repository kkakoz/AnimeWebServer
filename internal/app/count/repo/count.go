package service

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"red-bean-anime-server/internal/app/count/domain"
	"red-bean-anime-server/internal/app/count/pkg/keys"
	"red-bean-anime-server/pkg/db/mysqlx"
	"strconv"
)

type CountRepo struct {
	client *redis.Client
}

func NewCountRepo(client *redis.Client) domain.ICountRepo {
	return &CountRepo{client: client}
}

func (c *CountRepo) AddView(ctx context.Context, animeId int64) error {
	intCmd := c.client.HIncrBy(keys.GetAnimeCountIncrKey(animeId), keys.GetAnimeViewKey(), 1)
	return errors.Wrap(intCmd.Err(), "增加播放量失败")
}

// 从缓存获取count
func (c *CountRepo) GetAnimeCountCacheInfo(ctx context.Context, animeId int64) (*domain.AnimeCount, bool) {
	mapCmd := c.client.HGetAll(keys.GetAnimeCountIncrKey(animeId))
	animeCount := &domain.AnimeCount{}
	result, err := mapCmd.Result()
	if err != nil {
		return nil, false
	}
	viewCount, _ := strconv.Atoi(result[keys.GetAnimeViewKey()])
	likeCount, _ := strconv.Atoi(result[keys.GetAnimeLikeKey()])
	collectCount, _ := strconv.Atoi(result[keys.GetAnimeCollectKey()])
	animeCount.ViewCount = int64(viewCount)
	animeCount.CollectCount = int64(likeCount)
	animeCount.CollectCount = int64(collectCount)
	return animeCount, true
}

func (c *CountRepo) GetAnimeCount(ctx context.Context, animeId int64) (*domain.AnimeCount, error){
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return nil, err
	}
	animeCount := &domain.AnimeCount{}
	err = db.Where("anime_id = ?").First(animeCount).Error
	return animeCount, errors.Wrap(err, "发送")
}
