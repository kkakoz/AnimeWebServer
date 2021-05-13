// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"context"
	"red-bean-anime-server/internal/app/user/service"
	"red-bean-anime-server/internal/pkg/app"
	"red-bean-anime-server/internal/pkg/cache"
	"red-bean-anime-server/internal/pkg/config"
	"red-bean-anime-server/internal/pkg/log"
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
	registerService := service.NewUserService()
	grpcServer, err := app.NewGrpcServer(ctx, viper, client, registerService)
	if err != nil {
		return nil, err
	}
	appApp, err := app.NewApp(viper, logger, grpcServer)
	if err != nil {
		return nil, err
	}
	return appApp, nil
}
