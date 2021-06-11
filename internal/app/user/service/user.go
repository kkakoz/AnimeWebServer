package service

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	userpb "red-bean-anime-server/api/user"
	"red-bean-anime-server/internal/app/user/domain"
	"red-bean-anime-server/pkg/app"
	"red-bean-anime-server/pkg/db/mysqlx"
	"red-bean-anime-server/pkg/gerrors"
)

type UserService struct {
	userUsecase domain.IUserUsecase
}

func (u *UserService) Register(ctx context.Context, req *userpb.RegisterReq) (*emptypb.Empty, error) {
	err := req.Validate()
	if err != nil {
		return nil, gerrors.NewParamErr(err)
	}
	err = u.userUsecase.Register(ctx, req.Email, req.Name, req.Password)
	return &emptypb.Empty{}, err
}

func (u *UserService) Login(ctx context.Context, req *userpb.LoginReq) (*userpb.LoginRes, error) {
	err := req.Validate()
	if err != nil {
		return nil, gerrors.NewParamErr(err)
	}
	userInfo, err := u.userUsecase.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &userpb.LoginRes{
		Id:       userInfo.ID,
		Name:     userInfo.Name,
		Email:    userInfo.Email,
		Token:    userInfo.Token,
		CreateAt: userInfo.CreatedAt,
	}, nil
}

func (u *UserService) UserInfo(ctx context.Context, id *userpb.Id) (*userpb.UserInfoRes, error) {
	return &userpb.UserInfoRes{}, nil
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
}

