// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"context"
	"red-bean-anime-server/internal/app/count/pkg/kafka_runner"
	"red-bean-anime-server/internal/app/count/repo"
	"red-bean-anime-server/internal/app/count/service"
	"red-bean-anime-server/internal/app/count/usecase"
	"red-bean-anime-server/pkg/app"
	"red-bean-anime-server/pkg/cache"
	"red-bean-anime-server/pkg/config"
	"red-bean-anime-server/pkg/db/etcd"
	"red-bean-anime-server/pkg/db/mysqlx"
	"red-bean-anime-server/pkg/kafkax"
	"red-bean-anime-server/pkg/log"
)

// Injectors from wire.go:

func NewApp(ctx context.Context, confpath string) (*app.App, error) {
	viper, err := config.NewConfig(confpath)
	if err != nil {
		return nil, err
	}
	logger, err := log.NewLog(viper)
	if err != nil {
		return nil, err
	}
	client, err := etcd.NewEtcd(viper)
	if err != nil {
		return nil, err
	}
	redisClient, err := cache.NewRedis(viper)
	if err != nil {
		return nil, err
	}
	iCountRepo := repo.NewCountRepo(redisClient, logger)
	db, err := mysqlx.New(viper)
	if err != nil {
		return nil, err
	}
	iLikeRepo := repo.NewLikeRepo(redisClient, logger)
	iCountUsecase := usecase.NewCountUsecase(iCountRepo, db, logger, iLikeRepo)
	consumerRunFunc := kafka_runner.NewCountConsumerRunner(iCountUsecase)
	getTopic := kafka_runner.GetTopic()
	consumerRun, err := kafkax.NewConsumer(ctx, viper, logger, consumerRunFunc, getTopic)
	if err != nil {
		return nil, err
	}
	registerService := service.NewCountService(iCountUsecase, consumerRun)
	grpcServer := app.NewGrpcServer(ctx, client, registerService)
	appApp, err := app.NewApp(viper, logger, grpcServer)
	if err != nil {
		return nil, err
	}
	return appApp, nil
}