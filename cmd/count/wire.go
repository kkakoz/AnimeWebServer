package main

import (
	"context"
	"github.com/google/wire"
	"red-bean-anime-server/internal/app/count/pkg/kafka_runner"
	"red-bean-anime-server/internal/app/count/service"
	"red-bean-anime-server/pkg/app"
	"red-bean-anime-server/pkg/auth"
	"red-bean-anime-server/pkg/config"
	"red-bean-anime-server/pkg/db/etcd"
	"red-bean-anime-server/pkg/db/mysqlx"
	"red-bean-anime-server/pkg/kafkax"
	"red-bean-anime-server/pkg/log"
)

func NewApp(ctx context.Context, confpath string) (*app.App, error) {
	panic(wire.Build(
		config.ConfigSet,
		app.AppSet,
		etcd.EtcdSet,
		service.ServiceSet,
		log.LogSet,
		mysqlx.MysqlSet,
		auth.AuthSet,
		kafkax.KafkaConsumerSet,
		kafka_runner.CountConsumerSet,
	))
}
