package service

import (
	"github.com/google/wire"
	userpb "red-bean-anime-server/api/user"
	"context"
	"google.golang.org/grpc"
	"red-bean-anime-server/internal/pkg/app"
)

type UserService struct {

}

func (u *UserService) Login(ctx context.Context, info *userpb.LoginInfo) (*userpb.LoginResponse, error) {
	panic("implement me")
}

func (u *UserService) UserInfo(ctx context.Context, id *userpb.Id) (*userpb.LoginResponse, error) {
	panic("implement me")
}

func NewUserService() app.RegisterService {
	userService := &UserService{}
	return func(server *grpc.Server) {
		userpb.RegisterUserServiceServer(server, userService)
	}
	//return &UserService{}
}

var ProviderSet = wire.NewSet(NewUserService)