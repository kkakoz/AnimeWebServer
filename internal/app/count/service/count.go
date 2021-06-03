package service

import (
	"context"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"log"
	"red-bean-anime-server/internal/app/count/domain"
	"red-bean-anime-server/pkg/app"
	"red-bean-anime-server/pkg/db/mysqlx"
	"red-bean-anime-server/pkg/kafkax"
)

type CountService struct {
	countUsecase domain.ICountUsecase
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
		//userpb.RegisterUserServiceServer(server, userService)
	}
	//return &UserService{}
}

var ServiceSet = wire.NewSet(NewCountService,)