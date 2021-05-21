// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"context"
	"red-bean-anime-server/internal/app/user/repo"
	"red-bean-anime-server/internal/app/user/service"
	"red-bean-anime-server/internal/app/user/usecase"
	"red-bean-anime-server/pkg/app"
	"red-bean-anime-server/pkg/auth"
	"red-bean-anime-server/pkg/cache"
	"red-bean-anime-server/pkg/config"
	"red-bean-anime-server/pkg/db/mysqlx"
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
	client, err := cache.NewEtcd(viper)
	if err != nil {
		return nil, err
	}
	redisClient, err := cache.NewRedis(viper)
	if err != nil {
		return nil, err
	}
	userRepo := repo.NewUserRepo(redisClient)
	jwtTokenGen, err := auth.NewJwtTokenGen(viper)
	if err != nil {
		return nil, err
	}
	db, err := mysqlx.New(viper)
	if err != nil {
		return nil, err
	}
	userUsecase := usecase.NewUserUsecase(userRepo, jwtTokenGen, db)
	registerService := service.NewUserService(userUsecase)
	grpcServer := app.NewGrpcServer(ctx, client, registerService)
	appApp, err := app.NewApp(viper, logger, grpcServer)
	if err != nil {
		return nil, err
	}
	return appApp, nil
}
