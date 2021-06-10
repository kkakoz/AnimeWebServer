package grpcx

import (
	"context"
	"fmt"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
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

type validated interface {
	Validate() error
}

func NewValidateInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		validated, ok := req.(validated)
		if ok {
			err := validated.Validate()
			if err != nil {
				return nil, gerrors.NewParamErr(err)
			}
		}
		return handler(ctx, req)
	}
}

func RecoveryInterceptor() grpc_recovery.Option {
	return grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
		err, ok := p.(error)
		if ok {
			return errors.New(err.Error())
		}
		return errors.New(fmt.Sprintf("%+v", p))
	})
}