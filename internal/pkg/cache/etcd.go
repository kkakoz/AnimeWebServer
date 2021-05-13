package cache

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"time"
)

type EtcdOptions struct {
	Endpoints []string `json:"endpoints"`
}

func NewEtcd(viper *viper.Viper) (*clientv3.Client, error) {
	options := &EtcdOptions{}
	err := viper.UnmarshalKey("etcd", options)
	if err != nil {
		return nil, err
	}
	config := clientv3.Config{
		Endpoints: options.Endpoints,
		DialTimeout: 5 * time.Second,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	}
	client, err := clientv3.New(config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

var ProverSet = wire.NewSet(NewEtcd)
