package anime

import (
	"github.com/google/wire"
	"red-bean-anime-server/internal/app/anime/repo"
	"red-bean-anime-server/internal/app/anime/service"
	"red-bean-anime-server/internal/app/anime/usecase"
)

var AnimeSet = wire.NewSet(service.NewAnimeService, usecase.NewCategoryUsecase,
	repo.NewCategoryRepo, usecase.NewAnimeUsecase, repo.NewAnimeRepo, repo.NewVideoRepo)