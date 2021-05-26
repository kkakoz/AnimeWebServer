// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"red-bean-anime-server/internal/app/gateway"
	"red-bean-anime-server/pkg/cache"
	"red-bean-anime-server/pkg/config"
	"red-bean-anime-server/pkg/log"
)

func New(ctx context.Context, filepath string) (*gateway.Gateway, error) {
	panic(wire.Build(
		config.ProviderSet,
		cache.NewEtcd,
		gateway.NewGateway,
		log.NewLog,
	))
}
