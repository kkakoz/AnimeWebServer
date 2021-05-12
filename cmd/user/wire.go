// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"red-bean-anime-server/internal/app/user/service"
	"red-bean-anime-server/internal/pkg/app"
	"red-bean-anime-server/internal/pkg/cache"
	"red-bean-anime-server/internal/pkg/config"
	"red-bean-anime-server/internal/pkg/log"
)

func NewApp(ctx context.Context, confpath string) (*app.App, error) {
	panic(wire.Build(
		config.ProviderSet,
		app.ProviderSet,
		cache.ProverSet,
		service.ProviderSet,
		log.ProviderSet,
	))
}
