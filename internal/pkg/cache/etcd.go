package cache

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func NewEtcd(viper *viper.Viper) (*clientv3.Client, error) {
	config := &clientv3.Config{}
	err := viper.UnmarshalKey("etcd", config)
	if err != nil {
		return nil, err
	}
	client, err := clientv3.New(*config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

var ProverSet = wire.NewSet(NewEtcd)
