package repo

import (
	"context"
	"github.com/spf13/viper"
	"math/rand"
	"red-bean-anime-server/pkg/cache"
	"red-bean-anime-server/pkg/db/mysqlx"
	"testing"
)

func InitCountRepo() (*CountRepo, error) {
	viper.SetConfigFile("configs/test.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	redis, err := cache.NewRedis(viper.GetViper())
	if err != nil {
		return nil, err
	}
	_, err = mysqlx.New(viper.GetViper())
	if err != nil {
		return nil, err
	}
	repo := &CountRepo{redisCli: redis}
	return repo, nil
}

func TestCountRepo_AddView(t *testing.T) {
	repo, err := InitCountRepo()
	if err != nil {
		t.Fatal("init repo err:", err)
	}
	animeId := rand.Int63()
	err = repo.AddView(context.TODO(), animeId)
	if err != nil {
		t.Fatal("add view err:", err)
	}
	animeCount, b := repo.GetAnimeIncrCountByCache(context.TODO(), animeId)
	if !b {
		t.Logf("cache nil")
	}
	if animeCount.ViewCount == 0 {
		t.Fatal("view count incr err")
	}
	t.Logf("animecount = %+v", animeCount)
}
