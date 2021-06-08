package usecase

import (
	"context"
	"github.com/spf13/viper"
	"math/rand"
	"red-bean-anime-server/internal/app/count/domain"
	"red-bean-anime-server/internal/app/count/repo"
	"red-bean-anime-server/pkg/cache"
	"red-bean-anime-server/pkg/db/mysqlx"
	"red-bean-anime-server/pkg/log"
	"testing"
	"time"
)

func InitCountUsecase() (domain.ICountUsecase, error) {
	viper.SetConfigFile("configs/test.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	redis, err := cache.NewRedis(viper.GetViper())
	if err != nil {
		return nil, err
	}
	db, err := mysqlx.New(viper.GetViper())
	if err != nil {
		return nil, err
	}
	logger, err := log.NewLog(viper.GetViper())
	if err != nil {
		return nil, err
	}
	countRepo := repo.NewCountRepo(redis, logger)
	likeRepo := repo.NewLikeRepo(redis, logger)
	usecase := NewCountUsecase(countRepo, db, logger, likeRepo)
	return usecase, nil
}

func TestCountUsecase_AddAnimeViewCount(t *testing.T) {
	usecase, err := InitCountUsecase()
	if err != nil {
		t.Fatal("init count usecase", err)
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	animeId := r.Int63()
	err = usecase.AddAnimeCount(context.TODO(), animeId)
	if err != nil {
		t.Fatal("add count err:", err)
	}
	err = usecase.AddAnimeView(animeId)
	if err != nil {
		t.Fatal("add view err:", err)
	}
	t.Log("id = ", animeId)
	ids := make([]int64, 0)
	ids = append(ids, animeId)
	count, err := usecase.GetViewCount(context.TODO(), ids)
	if err != nil {
		t.Fatal("get view count err:", err)
	}
	t.Logf("count = %v", count)
	err = usecase.AddAnimeView(animeId)
	if err != nil {
		t.Fatal("add view err:", err)
	}
	usecase.UpdateIncr(context.TODO())
	viewCount, err := usecase.GetViewCount(context.TODO(), ids)
	if err != nil {
		t.Fatal("get view count err:", err)
	}
	t.Logf("get view count: %v", viewCount)
}

func TestCountUsecase_AddAnimeLikeCount(t *testing.T) {
	usecase, err := InitCountUsecase()
	if err != nil {
		t.Fatal("init count usecase", err)
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	animeId := r.Int63()
	userId := r.Int63()
	err = usecase.AddAnimeCount(context.TODO(), animeId)
	if err != nil {
		t.Fatal("add count err:", err)
	}
	t.Logf("animeId = %d, userId = %d", animeId, userId)
	err = usecase.UserLikeAnime(userId, animeId, true)
	if err != nil {
		t.Fatal("add view err:", err)
	}
	likeCollect, err := usecase.GetAnimeLikeCollect(context.TODO(), userId, animeId)
	if err != nil {
		t.Fatal("get like collect err:", err)
	}
	t.Logf("likeCollect = %v", likeCollect)
	usecase.UpdateIncr(context.TODO())
	t.Log("update incr")
	likeCollect, err = usecase.GetAnimeLikeCollect(context.TODO(), userId, animeId)
	if err != nil {
		t.Fatal("get like collect err:", err)
	}
	t.Logf("likeCollect = %v", likeCollect)
	if likeCollect.LikeCount != 1 || likeCollect.UserLiked != true {
		t.Error("like collect add err")
	}
}
