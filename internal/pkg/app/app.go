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

func NewApp(viper *viper.Viper, logger *zap.Logger, grpcServer *GrpcServer) (*App, error) {
	app := &App{}
	err := viper.UnmarshalKey("app", app)
	if err != nil {
		return nil, err
	}
	app.grpcServer = grpcServer
	return app, nil
}

func (a *App) Start() error {
	err := a.grpcServer.run(a.Name, a.Host, a.Port)
	if err != nil {
		return err
	}
	return nil
}

var ProviderSet = wire.NewSet(NewApp, NewGrpcServer)
