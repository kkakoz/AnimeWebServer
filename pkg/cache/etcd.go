package cache

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"time"
)

type EtcdOptions struct {
	Endpoints []string `json:"endpoints"`
}

func NewEtcd(viper *viper.Viper) (*clientv3.Client, error) {
	viper.SetDefault("etcd.endpoints", []string{"127.0.0.1:2379"})
	options := &EtcdOptions{}
	err := viper.UnmarshalKey("etcd", options)
	if err != nil {
		return nil, errors.Wrap(err, "viper unmarshal失败")
	}
	config := clientv3.Config{
		Endpoints: options.Endpoints,
		DialTimeout: 5 * time.Second,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	}
	client, err := clientv3.New(config)
	if err != nil {
		return nil, errors.Wrap(err, "etcd连接失败")
	}
	return client, nil
}

var ProverSet = wire.NewSet(NewEtcd, NewRedis)
