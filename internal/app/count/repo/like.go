package repo

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"red-bean-anime-server/internal/app/count/domain"
	"red-bean-anime-server/internal/app/count/pkg/keys"
	"red-bean-anime-server/pkg/cache/bloom_filter"
	"red-bean-anime-server/pkg/db/mysqlx"
	"strconv"
)

type LikeRepo struct {
	redisCli    *redis.Client
	logger      *zap.Logger
	bloomFilter bloom_filter.BloomFilter
}

func (l *LikeRepo) AddUserLike(ctx context.Context, like *domain.UserLike) error {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return err
	}
	err = db.Create(like).Error
	return errors.Wrap(err, "创建失败")
}

func (l *LikeRepo) AddBloomFilter(userId int64, animeId int64) error {
	err := l.bloomFilter.Add(keys.GetAnimeLikeBloomKey(animeId), strconv.FormatInt(userId, 10))
	return errors.Wrap(err, "添加布隆过滤器失败")
}

func (l *LikeRepo) GetLikeByBloomFilter(userId int64, animeId int64) bool {
	return l.bloomFilter.Contains(keys.GetAnimeLikeBloomKey(animeId), strconv.FormatInt(userId, 10))
}

func (l *LikeRepo) DelUserLike(ctx context.Context, like *domain.UserLike) error {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return err
	}
	err = db.Where("anime_id = ? and user_id = ?", like.AnimeId, like.UserId).Delete(&domain.UserLike{}).Error
	return errors.Wrap(err, "删除用户点赞失败")
}

func (l *LikeRepo) GetUserLike(ctx context.Context, userId int64, animeId int64) (bool, error) {
	db, err := mysqlx.GetDB(ctx)
	if err != nil {
		return false, err
	}
	like := &domain.UserLike{}
	err = db.Where("user_id = ? and anime_id = ?", userId, animeId).First(&like).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 不存在的err
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func NewLikeRepo(redisCli *redis.Client, logger *zap.Logger) domain.ILikeRepo {
	return &LikeRepo{redisCli: redisCli, logger: logger, bloomFilter: bloom_filter.NewBloomFilter(redisCli)}
}

