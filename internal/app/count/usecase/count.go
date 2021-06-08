package usecase

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"red-bean-anime-server/internal/app/count/domain"
	"red-bean-anime-server/internal/app/count/pkg/keys"
)

type CountUsecase struct {
	countRepo domain.ICountRepo
	db        *gorm.DB
	logger    *zap.Logger
	likeRepo  domain.ILikeRepo
}

func (c *CountUsecase) GetAnimeLikeCollect(ctx context.Context, userId int64, animeId int64) (*domain.UserLikeCollectCountRes, error) {
	var (
		likeCollectCount *domain.AnimeLikeCollectCountRes
		b                bool
		err              error
		res              = &domain.UserLikeCollectCountRes{}
	)
	likeCollectCount, b = c.countRepo.GetAnimeLikeCollectCache(ctx, animeId)
	if !b {
		likeCollectCount, err = c.countRepo.GetAnimeLikeCollectCount(ctx, animeId)
		if err != nil {
			return nil, err
		}
	}
	res.AnimeId = animeId
	res.LikeCount = likeCollectCount.LikeCount
	res.CollectCount = likeCollectCount.CollectCount
	like := c.likeRepo.GetLikeByBloomFilter(userId, animeId)
	if like {
		likeAnime, err := c.likeRepo.GetUserLike(ctx, userId, animeId)
		if err != nil {
			return nil, err
		}
		res.UserLiked = likeAnime
	} else {
		res.UserLiked = false
	}
	return res, nil
}

func (c *CountUsecase) AddAnimeView(animeId int64) error {
	return c.countRepo.AddCount(animeId, keys.HFAnimeIncrView)
}

func (c *CountUsecase) UserLikeAnime(userId int64, animeId int64, likeType bool) error {
	userLike := &domain.UserLike{
		UserId:  userId,
		AnimeId: animeId,
	}
	if likeType { // 点赞
		err := c.countRepo.AddCount(animeId, keys.HFAnimeIncrLike)
		if err != nil {
			return err
		}
		err = c.likeRepo.AddBloomFilter(userId, animeId)
		if err != nil {
			return err
		}
		return c.likeRepo.AddUserLike(context.Background(), userLike)
	}
	// 取消点赞
	err := c.countRepo.DelCount(animeId, keys.HFAnimeIncrLike)
	if err != nil {
		return err
	}
	have := c.likeRepo.GetLikeByBloomFilter(userId, animeId)
	if have {
		err = c.likeRepo.DelUserLike(context.Background(), userLike) // 删除点赞记录
		return err
	}
	return nil
}

func (c *CountUsecase) UserUnLikeAnime(userId int64, animeId int64, unlikeType bool) error {
	panic("un implement")
}

func (c *CountUsecase) UpdateIncr(ctx context.Context) {
	var cur uint64 = 0
	var ids []int64
	var err error
	for {
		ids, cur, err = c.countRepo.GetAnimeIncrIds(ctx, cur)
		if err != nil {
			c.logger.Error("get incr ids err", zap.Error(err))
		}
		for _, id := range ids {
			cache, b := c.countRepo.GetAnimeIncrCountByCache(ctx, id)
			if !b {
				continue
			}
			err := c.countRepo.IncrAnimeCount(ctx, cache)
			if err != nil {
				c.logger.Error("incr count to db err", zap.Error(err))
				continue
			}
			c.countRepo.ReduceCache(cache) // 缓存减少对应的长度
			err = c.countRepo.RemoveCountCache(cache.AnimeId)
			if err != nil {
				c.logger.Error("incr count to db err", zap.Error(err))
				continue
			}
		}
		if cur == 0 {
			break
		}
	}
}

func (c CountUsecase) GetViewCount(ctx context.Context, id []int64) ([]domain.AnimeViewCountRes, error) {
	res := make([]domain.AnimeViewCountRes, 0, len(id))
	cache, noCacheIds := c.countRepo.GetAnimeViewByCache(ctx, id)
	for _, viewCount := range cache {
		res = append(res, domain.AnimeViewCountRes{
			AnimeId:   viewCount.Key,
			ViewCount: viewCount.Val,
		})
	}
	if len(noCacheIds) == 0 {
		return res, nil
	}
	viewCounts, err := c.countRepo.GetAnimeViewCount(ctx, noCacheIds)
	if err != nil {
		return res, err
	}
	res = append(res, viewCounts...)
	err = c.countRepo.SetAnimeViewCache(ctx, viewCounts)
	if err != nil {
		c.logger.Error("get view by cache err", zap.Error(err))
	}
	return res, nil
}

func (c *CountUsecase) AddAnimeCount(ctx context.Context, animeId int64) error {
	return c.countRepo.AddAnimeCount(ctx, animeId)
}

func NewCountUsecase(countRepo domain.ICountRepo, db *gorm.DB, logger *zap.Logger, likeRepo domain.ILikeRepo) domain.ICountUsecase {
	return &CountUsecase{countRepo: countRepo, db: db, logger: logger, likeRepo: likeRepo}
}
