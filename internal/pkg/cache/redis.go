package cache

import (
	"github.com/go-redis/redis"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type Options struct {
	Host string
	Port string
	Password string
	DB int
	PoolSize int
}

var Redis *redis.Client

func NewRedis(viper *viper.Viper) error {
	options := &redis.Options{}
	err := viper.UnmarshalKey("redis", options)
	if err != nil {
		return err
	}
	redis.NewClient(options)

	//Redis = redis.NewClient(&redis.Options{
	//	Addr:     viper.GetString("host") + ":" + viper.GetString("port"),
	//	Password: viper.GetString("password"),
	//	DB:       viper.GetInt("mysql"),
	//	PoolSize: viper.GetInt("poolSize"),
	//})
	_, err = Redis.Ping().Result()
	return err
}


var ProviderSet = wire.NewSet(NewRedis)