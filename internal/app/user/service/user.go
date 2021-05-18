package service

import (
	"context"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	userpb "red-bean-anime-server/api/user"
	"red-bean-anime-server/internal/app/user/repo"
	"red-bean-anime-server/internal/app/user/usecase"
	"red-bean-anime-server/pkg/app"
)

type UserService struct {
	userUsercase *usecase.UserUsecase
}

func (u *UserService) Login(ctx context.Context, info *userpb.LoginInfo) (*userpb.LoginResponse, error) {

	return &userpb.LoginResponse{
		Token:                "hsidoahif",
	}, nil
}

func (u *UserService) UserInfo(ctx context.Context, id *userpb.Id) (*userpb.LoginResponse, error) {
	return &userpb.LoginResponse{Token: "sda"}, status.Errorf(codes.Internal, "internal err")
}

func NewUserService(userUsercase *usecase.UserUsecase) app.RegisterService {
	userService := &UserService{
		userUsercase: userUsercase,
	}
	return func(server *grpc.Server) {
		userpb.RegisterUserServiceServer(server, userService)
	}
	//return &UserService{}
}

var ProviderSet = wire.NewSet(NewUserService, usecase.NewUserUsecase, repo.NewUserRepo)