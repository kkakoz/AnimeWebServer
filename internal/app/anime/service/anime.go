package service

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"red-bean-anime-server/api/anime"
	"red-bean-anime-server/internal/app/anime/domain"
	"red-bean-anime-server/internal/pkg/query"
	"red-bean-anime-server/pkg/app"
	"red-bean-anime-server/pkg/db/mysqlx"
)

type AnimeService struct {
	categoryUsecase domain.ICategoryUsecase
	animeUsecase    domain.IAnimeUsecase
}

func (a *AnimeService) AnimeLike(ctx context.Context, req *animepb.AnimeLikeReq) (*emptypb.Empty, error) {
	err := a.animeUsecase.UserLikeAnime(ctx, req.AnimeId, req.LikeType)
	return &emptypb.Empty{}, err
}

func (a *AnimeService) AnimeUnLike(ctx context.Context, req *animepb.AnimeLikeReq) (*emptypb.Empty, error) {
	err := a.animeUsecase.UserUnLikeAnime(ctx, req.AnimeId, req.LikeType)
	return &emptypb.Empty{}, err
}

func (a *AnimeService) GetAnimeList(ctx context.Context, req *animepb.GetAnimeListReq) (*animepb.AnimeListRes, error) {
	page := &query.Page{
		Page:     int(req.Page),
		PageSize: int(req.PageSize),
	}
	listReq := &domain.AnimeListReq{
		CategoryId: req.CategoryId,
		Sort:       req.Sort,
	}
	list, err := a.animeUsecase.GetAnimeList(ctx, page, listReq)
	if err != nil {
		return nil, err
	}
	animeListRes := make([]*animepb.AnimeRes, 0)
	for _, anime := range list {
		animeListRes = append(animeListRes, &animepb.AnimeRes{
			Id:          int64(anime.ID),
			Name:        anime.Name,
			Description: anime.Description,
			ImageUrl:    anime.ImageUrl,
			Year:        anime.Year,
			Quarter:     anime.Quarter,
		})
	}
	res := &animepb.AnimeListRes{
		Animeinfo: animeListRes,
	}
	return res, nil
}

func (a *AnimeService) GetAnimeInfo(ctx context.Context, req *animepb.AnimeInfoReq) (*animepb.AnimeInfoRes, error) {
	info, err := a.animeUsecase.GetAnimeInfo(ctx, req.AnimeId)
	if err != nil {
		return nil, err
	}
	videos := make([]*animepb.VideoInfo, 0)
	for _, v := range info.Videos {
		videos = append(videos, &animepb.VideoInfo{
			Id:      v.ID,
			Episode: v.Episode,
			Name:    v.Name,
			Url:     v.Url,
		})
	}
	res := &animepb.AnimeInfoRes{
		Id:                   info.ID,
		Name:                 info.Name,
		Description:          info.Description,
		Year:                 info.Year,
		Quarter:              info.Quarter,
		LikeCount:            info.LikeCount,
		CollectCount:         info.CollectCount,
		VideoInfos:           videos,
		Like:                 info.Like,
		Collect:              info.Collect,
	}
	return res, nil
}

func (a *AnimeService) AddAnime(ctx context.Context, req *animepb.AddAnimeReq) (*emptypb.Empty, error) {
	addAnime := &domain.AddAnime{
		Name:          req.Name,
		Description:   req.Description,
		Year:          req.Year,
		Quarter:       req.Quarter,
		FirstPlayTime: req.FirstPlayTime,
		CategoryIds:   req.CategoryId,
	}
	err := a.animeUsecase.AddAnime(ctx, addAnime)
	return &emptypb.Empty{}, err
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

func (a *AnimeService) AddVideo(ctx context.Context, req *animepb.AddVideoReq) (*emptypb.Empty, error) {
	addVideo := &domain.Video{
		AnimeId: req.AnimeId,
		Episode: req.Episode,
		Name:    req.Name,
		Url:     req.Url,
	}
	err := a.animeUsecase.AddVideo(ctx, addVideo)
	return &emptypb.Empty{}, err
}

func NewAnimeService(categoryUsecase domain.ICategoryUsecase, animeUsecase domain.IAnimeUsecase) app.RegisterService {
	animeService := &AnimeService{categoryUsecase: categoryUsecase, animeUsecase: animeUsecase}
	return func(server *grpc.Server) {
		db, err := mysqlx.GetDB(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		err = db.AutoMigrate(&domain.Anime{}, &domain.Category{}, &domain.AnimeCategory{}, &domain.Video{})
		if err != nil {
			log.Fatal(err)
		}
		animepb.RegisterAnimeServiceServer(server, animeService)
	}
}
