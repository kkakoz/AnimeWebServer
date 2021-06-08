package mongo

import (
	"context"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/qiniu/qmgo"
	"github.com/spf13/viper"
)

type Options struct {
	Url string
}

func NewMongoClient(viper *viper.Viper) (*qmgo.Client, error) {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017"})
	if err != nil {
		return nil, errors.Wrap(err, "连接mongodb失败")
	}
	return client, nil
}

var MongoSet = wire.NewSet(NewMongoClient)