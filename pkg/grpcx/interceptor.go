package grpcx

import (
	"context"
	"fmt"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"red-bean-anime-server/pkg/gerrors"
)

type Msg struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}


func ServerErrorInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	return resp, gerrors.WrapError(err)
}

func NewClientErrInterceptor(servName string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc)
		return gerrors.WrapRPCError(err, servName)
	}
}

func RecoveryInterceptor() grpc_recovery.Option {
	return grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
		err = errors.New(fmt.Sprintf("%v", p))
		return grpc.Errorf(codes.Unknown, fmt.Sprintf("%+v", err))
	})
}