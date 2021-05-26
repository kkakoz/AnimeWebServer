// +build wireinject

package main

import (
	"github.com/google/wire"
	"red-bean-anime-server/internal/app/file/service"
	"red-bean-anime-server/internal/app/file/usecase"
)

func NewService() (*service.FileService) {
	panic(wire.Build(service.NewFileService, usecase.NewFileUsecase))
}
