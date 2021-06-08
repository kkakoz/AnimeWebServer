package count

import (
	"github.com/google/wire"
	"red-bean-anime-server/internal/app/count/repo"
	"red-bean-anime-server/internal/app/count/service"
	"red-bean-anime-server/internal/app/count/usecase"
)

var CountServiceSet = wire.NewSet(
	service.NewCountService,
	repo.NewCountRepo,
	usecase.NewCountUsecase,
	repo.NewLikeRepo,
	)
