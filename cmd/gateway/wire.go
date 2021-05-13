// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"red-bean-anime-server/internal/app/gateway"
	"red-bean-anime-server/internal/pkg/cache"
	"red-bean-anime-server/internal/pkg/config"
)

func New(ctx context.Context, filepath string) (*gateway.Gateway, error) {
	panic(wire.Build(
		config.ProviderSet,
		cache.NewEtcd,
		gateway.NewGateway,
	))
}
