package service

import (
	"context"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"red-bean-anime-server/api/anime"
	"red-bean-anime-server/internal/app/anime/domain"
	"red-bean-anime-server/internal/app/anime/repo"
	"red-bean-anime-server/internal/app/anime/usecase"
	"red-bean-anime-server/pkg/app"
	"red-bean-anime-server/pkg/db/mysqlx"
)

type AnimeService struct {
	categoryUsecase domain.ICategoryUsecase
	animeUsecase    domain.IAnimeUsecase
}

func (a *AnimeService) AddCategory(ctx context.Context, req *animepb.AddCategoryReq) (*emptypb.Empty, error) {
	err := a.categoryUsecase.AddCategory(ctx, req.Name)
	return &emptypb.Empty{}, err
}

func (a *AnimeService) CategoryList(ctx context.Context, empty *emptypb.Empty) (*animepb.CategoryListRes, error) {
	categoryList, err := a.categoryUsecase.GetCategoryList(ctx)
	if err != nil {
		return nil, err
	}
	length := len(categoryList)
	res := &animepb.CategoryListRes{}
	categories := make([]*animepb.Category, 0, length)
	for _, c := range categoryList {
		category := &animepb.Category{
			Id:   int64(c.ID),
			Name: c.Name,
		}
		categories = append(categories, category)
	}
	res.CategoryList = categories
	res.Count = int64(length)
	return res, nil
}

func NewAnimeService(categoryUsecase domain.ICategoryUsecase, animeUsecase domain.IAnimeUsecase) app.RegisterService {
	animeService := &AnimeService{categoryUsecase: categoryUsecase, animeUsecase: animeUsecase}
	return func(server *grpc.Server) {
		db, err := mysqlx.GetDB(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		err = db.AutoMigrate(&domain.Anime{}, &domain.Category{})
		if err != nil {
			log.Fatal(err)
		}
		animepb.RegisterAnimeServiceServer(server, animeService)
	}
}

var ProviderSet = wire.NewSet(NewAnimeService, usecase.NewCategoryUsecase,
	repo.NewCategoryRepo, usecase.NewAnimeUsecase, repo.NewAnimeRepo)
