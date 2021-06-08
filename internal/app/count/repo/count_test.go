package repo

import (
	"context"
	"github.com/spf13/viper"
	"math/rand"
	"os"
	"red-bean-anime-server/internal/app/count/domain"
	"red-bean-anime-server/internal/app/count/pkg/keys"
	"red-bean-anime-server/pkg/cache"
	"red-bean-anime-server/pkg/db/mysqlx"
	"red-bean-anime-server/pkg/log"
	"testing"
)

func InitCountRepo() (domain.ICountRepo, error) {
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
	logger, err := log.NewLog(viper.GetViper())
	countRepo := NewCountRepo(redis, logger)
	return countRepo, nil
}

func TestCountRepo_AddView(t *testing.T) {
	animeId := rand.Int63()
	err := repo.AddCount(animeId, keys.HFAnimeIncrView)
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

var repo domain.ICountRepo

func TestMain(m *testing.M) {
	var err error
	repo, err = InitCountRepo()
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}