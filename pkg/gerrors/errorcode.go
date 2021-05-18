package gerrors

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	basepb "red-bean-anime-server/api/base"
)

var (
	ErrUnknown           = status.New(codes.Unknown, "error unknown").Err()                           // 服务器未知错误
	ErrUnauthorized      = newError(basepb.ErrorCode_EC_UNAUTHORIZED, "error unauthorized") // 未登录

)

func newError(code basepb.ErrorCode, message string) error {
	return status.New(codes.Code(code), message).Err()
}

var InterruptErr = errors.New("gateway interrupt")