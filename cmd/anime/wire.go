// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"red-bean-anime-server/internal/app/anime/service"
	"red-bean-anime-server/pkg/app"
	"red-bean-anime-server/pkg/config"
	"red-bean-anime-server/pkg/db/etcd"
	"red-bean-anime-server/pkg/db/mysqlx"
	"red-bean-anime-server/pkg/log"
)

func NewApp(ctx context.Context, confpath string) (*app.App, error) {
	panic(wire.Build(
		config.ProviderSet,
		app.ProviderSet,
		etcd.ProverSet,
		service.ProviderSet,
		log.ProviderSet,
		mysqlx.ProviderSet,
		//auth.ProviderSet,
	))
}


