package repo

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"red-bean-anime-server/internal/app/count/domain"
	"red-bean-anime-server/internal/app/count/pkg/keys"
	"red-bean-anime-server/internal/pkg/redisx"
	"red-bean-anime-server/pkg/db/mysqlx"
	"strconv"
)

type CountRepo struct {
	redisCli *redis.Client
	logger   *zap.Logger
}

func (c *CountRepo) AddAnimeCount(ctx context.Context, id int64) error {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return err
	}
	animeCount := &domain.AnimeCount{
		AnimeId: id,
	}
	err = db.Create(animeCount).Error
	return errors.Wrap(err, "添加视频计数失败")
}

func (c *CountRepo) AddView(ctx context.Context, animeId int64) error {
	intCmd := c.redisCli.HIncrBy(keys.GetHAnimeCountIncrKey(animeId), keys.HFAnimeIncrView, 1)
	return errors.Wrap(intCmd.Err(), "增加播放量失败")
}

func (c *CountRepo) GetAnimeIncrCountByCache(ctx context.Context, animeId int64) (*domain.AnimeCount, bool) {
	res, err := c.redisCli.HMGet(keys.GetHAnimeCountIncrKey(animeId), keys.HFAnimeIncrs...).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) { //
			return nil, false
		}
		c.logger.Error("get anime incr count cache err:", zap.Error(err))
		return nil, false
	}
	animeCount := &domain.AnimeCount{}
	for i, v := range res {
		switch keys.HFAnimeIncrs[i] {
		case keys.HFAnimeIncrView:
			animeCount.ViewCount = redisx.InterToIntF(v)
		case keys.HFAnimeIncrCollect:
			animeCount.CollectCount = redisx.InterToIntF(v)
		case keys.HFAnimeIncrLike:
			animeCount.LikeCount = redisx.InterToIntF(v)
		}
	}
	return animeCount, true
}

func (c *CountRepo) GetAnimeCount(ctx context.Context, animeId int64) (*domain.AnimeCount, error) {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return nil, err
	}
	animeCount := &domain.AnimeCount{}
	err = db.Where("anime_id = ?").First(animeCount).Error
	return animeCount, errors.Wrap(err, "发送")
}

func (c *CountRepo) GetAnimeIncrKey(ctx context.Context, cur uint64) ([]string, uint64, error) {
	scanCmd := c.redisCli.Scan(cur, keys.HAnimeIncrCountKey+"*", 100)
	result, cursor, err := scanCmd.Result()
	if err != nil {
		return nil, 0, errors.Wrap(err, "获取incr keys失败")
	}
	return result, cursor, nil
}

func (c *CountRepo) GetAnimeViewByCache(animeIds []int64) ([]domain.CountCache, []int64) {
	animeIdStrs := make([]string, 0, len(animeIds))
	for _, animeId := range animeIds {
		animeIdStrs = append(animeIdStrs, strconv.FormatInt(animeId, 10))
	}
	slice, err := c.redisCli.HMGet(keys.GetAnimeViewKey(), animeIdStrs...).Result() // 从缓存中获取
	if err != nil {                                                                 // key不存在或者redis出错
		if errors.Is(err, redis.Nil) { //
			return nil, animeIds
		}
		c.logger.Error("get anime view cache err:", zap.Error(err))
		return nil, animeIds
	}
	var existSlice = make([]domain.CountCache, 0) // 存在的cache
	var noexistSlice = make([]int64, 0)           // 不存在的cache
	for i, v := range slice {
		count, b := redisx.InterToInt(v)
		if b {
			existSlice = append(existSlice, domain.CountCache{
				Key: animeIds[i],
				Val: count,
			})
		} else {
			noexistSlice = append(noexistSlice, animeIds[i])
		}
	}
	return existSlice, noexistSlice
}

func (c *CountRepo) SetAnimeViewCache(animeId int64, viewCount int) error {
	result, err := c.redisCli.HSet(keys.GetAnimeViewKey(), strconv.FormatInt(animeId, 10), viewCount).Result()
	if err != nil {
		return errors.Wrap(err, "添加播放量缓存失败")
	}
	if !result {
		return errors.New("添加播放量缓存失败")
	}
	return nil
}

func (c *CountRepo) SetAnimeLikeCollectCache(animeId int64, likeCount, collectCount int) {
	c.redisCli.HSet(keys.GetAnimeViewKey())
}

func NewCountRepo(client *redis.Client) domain.ICountRepo {
	return &CountRepo{redisCli: client}
}
