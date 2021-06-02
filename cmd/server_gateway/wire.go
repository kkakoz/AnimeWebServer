// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"red-bean-anime-server/internal/app/gateway"
	"red-bean-anime-server/pkg/config"
	"red-bean-anime-server/pkg/db/etcd"
	"red-bean-anime-server/pkg/log"
)

func New(ctx context.Context, filepath string) (*gateway.Gateway, error) {
	panic(wire.Build(
		config.ProviderSet,
		etcd.NewEtcd,
		gateway.NewGateway,
		log.NewLog,
	))
}
