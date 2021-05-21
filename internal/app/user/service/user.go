package service

import (
	"context"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	userpb "red-bean-anime-server/api/user"
	"red-bean-anime-server/internal/app/user/domain"
	"red-bean-anime-server/internal/app/user/repo"
	"red-bean-anime-server/internal/app/user/usecase"
	"red-bean-anime-server/pkg/app"
	"red-bean-anime-server/pkg/db/mysqlx"
)

type UserService struct {
	userUsecase domain.IUserUsecase
}

func (u *UserService) Login(ctx context.Context, info *userpb.LoginInfo) (*userpb.LoginRes, error) {
	return &userpb.LoginRes{
		Token:                "hsidoahif",
	}, nil
}

func (u *UserService) Register(ctx context.Context, info *userpb.RegisterReq) (*emptypb.Empty, error) {
	err := u.userUsecase.Register(ctx, info.Phone, info.Name, info.Password)
	return &emptypb.Empty{}, err
}

func (u *UserService) UserInfo(ctx context.Context, id *userpb.Id) (*userpb.UserInfoRes, error) {
	panic("implement me")
}

func NewUserService(userUsercase domain.IUserUsecase) app.RegisterService {
	userService := &UserService{
		userUsecase: userUsercase,
	}
	return func(server *grpc.Server) {
		db, err := mysqlx.GetDB(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		err = db.AutoMigrate(&domain.User{})
		if err != nil {
			log.Fatal(err)
		}
		userpb.RegisterUserServiceServer(server, userService)

	}
	//return &UserService{}
}

var ProviderSet = wire.NewSet(NewUserService, usecase.NewUserUsecase, repo.NewUserRepo)