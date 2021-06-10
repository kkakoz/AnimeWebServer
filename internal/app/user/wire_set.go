package user

import (
	"github.com/google/wire"
	"red-bean-anime-server/internal/app/user/repo"
	"red-bean-anime-server/internal/app/user/service"
	"red-bean-anime-server/internal/app/user/usecase"
)

var UserSet = wire.NewSet(service.NewUserService, usecase.NewUserUsecase, repo.NewUserRepo)
