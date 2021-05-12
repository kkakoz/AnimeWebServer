package app

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type App struct {
	Name       string
	Host       string
	Port       string
	logger     *zap.Logger
	grpcServer *GrpcServer
}

func NewApp(viper *viper.Viper, logger *zap.Logger, grpcServer *GrpcServer) *App {
	return &App{Name: viper.Sub("app").GetString("name"), logger: logger, grpcServer: grpcServer}
}

func (a *App) Start() error {
	err := a.grpcServer.run(a.Name, a.Host, a.Port)
	if err != nil {
		return err
	}
	return nil
}

var ProviderSet = wire.NewSet(NewApp, NewGrpcServer)
