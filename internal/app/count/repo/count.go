package repo

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"red-bean-anime-server/internal/app/count/domain"
	"red-bean-anime-server/internal/app/count/pkg/keys"
	"red-bean-anime-server/internal/pkg/redisx"
	"red-bean-anime-server/pkg/cache/bloom_filter"
	"red-bean-anime-server/pkg/db/mysqlx"
	"strconv"
	"strings"
)

type CountRepo struct {
	redisCli    *redis.Client
	logger      *zap.Logger
	bloomFilter bloom_filter.BloomFilter
}

func (c *CountRepo) RemoveCountCache(animeId int64) error {
	pipeline := c.redisCli.Pipeline()
	pipeline.HDel(keys.HAnimeViewKey, strconv.FormatInt(animeId, 10))
	pipeline.Del(keys.GetAnimeCountKey(animeId))
	_, err := pipeline.Exec()
	return errors.Wrap(err, "删除失败")
}

func (c *CountRepo) IncrAnimeCount(ctx context.Context, incrCount *domain.AnimeCount) error {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return err
	}
	updates := map[string]interface{}{}
	if incrCount.ViewCount == 0 && incrCount.CollectCount == 0 && incrCount.LikeCount == 0 {
		return nil
	}
	if incrCount.ViewCount != 0 {
		updates["view_count"] = gorm.Expr("view_count + ?", incrCount.ViewCount)
	}
	if incrCount.CollectCount != 0 {
		updates["collect_count"] = gorm.Expr("collect_count + ?", incrCount.CollectCount)
	}
	if incrCount.LikeCount != 0 {
		updates["like_count"] = gorm.Expr("like_count + ?", incrCount.LikeCount)
	}
	err = db.Table("anime_counts").Where("anime_id = ?", incrCount.AnimeId).
		Updates(updates).Error
	return errors.Wrap(err, "从缓存添加失败")
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

func (c *CountRepo) AddCount(animeId int64, fileKey string) error {
	intCmd := c.redisCli.HIncrBy(keys.GetHAnimeCountIncrKey(animeId), fileKey, 1)
	return errors.Wrap(intCmd.Err(), "增加播放量失败")
}

func (c *CountRepo) ReduceCache(cache *domain.AnimeCount) {
	pipeline := c.redisCli.Pipeline()
	if cache.CollectCount != 0 && cache.ViewCount != 0 && cache.LikeCount != 0 {
		return
	}
	if cache.CollectCount != 0 {
		pipeline.HIncrBy(keys.GetHAnimeCountIncrKey(cache.AnimeId), keys.HFAnimeIncrCollect, int64(-cache.CollectCount))
	}
	if cache.ViewCount != 0 {
		pipeline.HIncrBy(keys.GetHAnimeCountIncrKey(cache.AnimeId), keys.HFAnimeIncrView, int64(-cache.ViewCount))
	}
	if cache.LikeCount != 0 {
		pipeline.HIncrBy(keys.GetHAnimeCountIncrKey(cache.AnimeId), keys.HFAnimeIncrLike, int64(-cache.LikeCount))
	}
	_, err := pipeline.Exec()
	if err != nil {
		c.logger.Error("reduce count err", zap.Error(err))
	}
}

func (c *CountRepo) DelCount(animeId int64, fileKey string) error {
	intCmd := c.redisCli.HIncrBy(keys.GetHAnimeCountIncrKey(animeId), fileKey, -1)
	return errors.Wrap(intCmd.Err(), "取消失败")
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
	animeCount.AnimeId = animeId
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

func (c *CountRepo) GetAnimeViewCount(ctx context.Context, animeIds []int64) ([]domain.AnimeViewCountRes, error) {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return nil, err
	}
	counts := make([]domain.AnimeViewCountRes, 0, len(animeIds))
	err = db.Table("anime_counts").Select("anime_id, view_count").Where("anime_id in ?", animeIds).Find(&counts).Error
	return counts, errors.Wrap(err, "查询点击量失败")
}

func (c *CountRepo) GetAnimeLikeCollectCount(ctx context.Context, animeIds int64) (*domain.AnimeLikeCollectCountRes, error) {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return nil, err
	}
	count := &domain.AnimeLikeCollectCountRes{}
	err = db.Table("anime_counts").Select("anime_id, like_count, collect_count").Where("anime_id = ?", animeIds).First(&count).Error
	return count, errors.Wrap(err, "查询点赞和收藏量失败")
}

func (c *CountRepo) GetAnimeIncrIds(ctx context.Context, cur uint64) ([]int64, uint64, error) {
	incrKeys, cursor, err := c.redisCli.Scan(cur, keys.HAnimeIncrCountKey+"*", 100).Result()
	incrIds := make([]int64, 0, len(incrKeys))
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return incrIds, cursor, nil
		}
		return nil, 0, errors.Wrap(err, "获取incr keys失败")
	}
	for _, key := range incrKeys {
		split := strings.Split(key, keys.HAnimeIncrCountKey)
		idstr := split[len(split)-1]
		id, err := strconv.ParseInt(idstr, 10, 64)
		if err != nil {
			continue
		}
		incrIds = append(incrIds, int64(id))
	}
	return incrIds, cursor, nil
}

func (c *CountRepo) GetAnimeViewByCache(ctx context.Context, animeIds []int64) ([]domain.CountCache, []int64) {
	animeIdStrs := make([]string, 0, len(animeIds))
	for _, animeId := range animeIds {
		animeIdStrs = append(animeIdStrs, strconv.FormatInt(animeId, 10))
	}
	slice, err := c.redisCli.HMGet(keys.HAnimeViewKey, animeIdStrs...).Result() // 从缓存中获取
	if err != nil {                                                             // key不存在或者redis出错
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

func (c *CountRepo) SetAnimeViewCache(ctx context.Context, counts []domain.AnimeViewCountRes) error {
	pipeline := c.redisCli.Pipeline()
	for _, count := range counts {
		c.redisCli.HSet(keys.HAnimeViewKey, strconv.FormatInt(count.AnimeId, 10), count.ViewCount)
	}
	_, err := pipeline.Exec()
	return errors.Wrap(err, "添加到缓存失败")
}

func (c *CountRepo) SetAnimeLikeCollectCache(ctx context.Context, animeId int64, likeCount, collectCount int) error {
	likeCollect := &domain.AnimeLikeCollectCountRes{
		AnimeId:      animeId,
		LikeCount:    likeCount,
		CollectCount: collectCount,
	}
	_, err := c.redisCli.Set(keys.GetAnimeCountKey(animeId), likeCollect, 0).Result()
	return errors.Wrap(err, "添加到缓存失败")
}

func (c *CountRepo) GetAnimeLikeCollectCache(ctx context.Context, animeId int64) (*domain.AnimeLikeCollectCountRes, bool) {
	bytes, err := c.redisCli.Get(keys.GetAnimeCountKey(animeId)).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) { //
			return nil, false
		}
		c.logger.Error("get anime incr count cache err:", zap.Error(err))
		return nil, false
	}
	res := &domain.AnimeLikeCollectCountRes{}
	err = json.Unmarshal(bytes, res)
	if err != nil {
		return nil, false
	}
	return res, true
}

func NewCountRepo(redisCli *redis.Client, logger *zap.Logger) domain.ICountRepo {
	return &CountRepo{redisCli: redisCli, logger: logger, bloomFilter: bloom_filter.NewBloomFilter(redisCli)}
}
