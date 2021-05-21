// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"red-bean-anime-server/internal/app/user/service"
	"red-bean-anime-server/pkg/app"
	"red-bean-anime-server/pkg/auth"
	"red-bean-anime-server/pkg/cache"
	"red-bean-anime-server/pkg/config"
	"red-bean-anime-server/pkg/db/mysqlx"
	"red-bean-anime-server/pkg/log"
)

func NewApp(ctx context.Context, confpath string) (*app.App, error) {
	panic(wire.Build(
		config.ProviderSet,
		app.ProviderSet,
		cache.ProverSet,
		service.ProviderSet,
		log.ProviderSet,
		mysqlx.ProviderSet,
		auth.ProviderSet,
	))
}
