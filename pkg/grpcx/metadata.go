package grpcx

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func GetAuthorization(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	auths := md.Get("authorization")
	if len(auths) == 1 {
		return auths[0]
	}
	return ""
}
