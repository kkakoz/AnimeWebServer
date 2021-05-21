package gerrors

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	basepb "red-bean-anime-server/api/base"
)

var (
	ErrUnauthorized  = newError(basepb.ErrorCode_EC_UNAUTHORIZED, "请登录后进行操作") // 未登录
	ErrServerUnknown = newError(basepb.ErrorCode_EC_SERVER_UNKNOWN, "未知错误")
)

func newError(code basepb.ErrorCode, message string) error {
	return status.New(codes.Code(code), message).Err()
}

// 业务逻辑错误
func NewBusErr(msg string) error {
	return newError(basepb.ErrorCode_EC_BUSINESSERR, msg)
}

var InterruptErr = errors.New("gateway interrupt")
