package service

import (
	"context"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	userpb "red-bean-anime-server/api/user"
	"red-bean-anime-server/internal/pkg/app"
)

type UserService struct {

}

func (u *UserService) Login(ctx context.Context, info *userpb.LoginInfo) (*userpb.LoginResponse, error) {
	return &userpb.LoginResponse{
		Token:                "hsidoahif",
	}, nil
}

func (u *UserService) UserInfo(ctx context.Context, id *userpb.Id) (*userpb.LoginResponse, error) {
	return nil, errors.New("ttt")
}

func NewUserService() app.RegisterService {
	userService := &UserService{}
	return func(server *grpc.Server) {
		userpb.RegisterUserServiceServer(server, userService)
	}
	//return &UserService{}
}

var ProviderSet = wire.NewSet(NewUserService)