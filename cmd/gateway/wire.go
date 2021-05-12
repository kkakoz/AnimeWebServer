// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/viper"
	"red-bean-anime-server/internal/app/gateway"
	"red-bean-anime-server/internal/pkg/cache"
)

func New(ctx context.Context, viper *viper.Viper, opts ...runtime.ServeMuxOption) (*gateway.Gateway, error) {
	panic(wire.Build(
		runtime.NewServeMux,
		cache.NewEtcd,
		gateway.NewGateway,
	))
}
