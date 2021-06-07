package service

import (
	"context"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"red-bean-anime-server/api/count"
	"red-bean-anime-server/internal/app/count/domain"
	"red-bean-anime-server/internal/app/count/repo"
	"red-bean-anime-server/internal/app/count/usecase"
	"red-bean-anime-server/pkg/app"
	"red-bean-anime-server/pkg/db/mysqlx"
	"red-bean-anime-server/pkg/kafkax"
)

type CountService struct {
	countUsecase domain.ICountUsecase
}

func (c *CountService) AddAnimeCount(ctx context.Context, req *countpb.AnimeIdReq) (*emptypb.Empty, error) {
	err := c.countUsecase.AddAnimeCount(ctx, req.AnimeId)
	return &emptypb.Empty{}, err
}

func (c *CountService) GetViewCount(ctx context.Context, req *countpb.AnimeIdsReq) (*countpb.ViewCountRes, error) {
	counts, err := c.countUsecase.GetViewCount(ctx, req.AnimeId)
	if err != nil {
		return nil, err
	}
	viewCounts := make([]*countpb.CountRes, 0, len(counts))
	for _, count := range counts {
		viewCounts = append(viewCounts, &countpb.CountRes{
			AnimeId: count.AnimeId,
			Count: count.ViewCount,
		})
	}
	return &countpb.ViewCountRes{
		ViewCounts:           viewCounts,
	}, nil
}

func (c *CountService) GetAnimeCount(ctx context.Context, req *countpb.AnimeIdReq) (*countpb.AnimeCountRes, error) {
	panic("implement me")
}

func NewCountService(countUsecase domain.ICountUsecase, run kafkax.ConsumerRun) app.RegisterService {
	//userService := &CountService{
	//	countUsecase: countUsecase,
	//}
	return func(server *grpc.Server) {
		db, err := mysqlx.GetDB(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		err = db.AutoMigrate(&domain.UserLike{}, &domain.AnimeCount{})
		if err != nil {
			log.Fatal(err)
		}
		go run.Run()
	}
}

var ServiceSet = wire.NewSet(NewCountService, usecase.NewCountUsecase, repo.NewCountRepo)

