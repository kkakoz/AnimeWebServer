package auth

import (
	"context"
	"google.golang.org/grpc/metadata"
)

type UserToken struct {

}

func (u UserToken) GetUser(ctx context.Context)  {
	md, _ := metadata.FromIncomingContext(ctx)
	md.Get("authorization")
}