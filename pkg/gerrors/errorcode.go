package gerrors

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	basepb "red-bean-anime-server/api/base"
)

var (
	ErrUnauthorized  = NewError(basepb.ErrorCode_EC_UNAUTHORIZED, "请登录后进行操作") // 未登录
	ErrServerUnknown = NewError(basepb.ErrorCode_EC_SERVER_UNKNOWN, "未知错误")
	ErrParam         = NewError(basepb.ErrorCode_EC_PARAMERR, "参数错误")
)

func NewError(code basepb.ErrorCode, message string) error {
	return status.New(codes.Code(code), message).Err()
}

// 参数错误
func NewParamErr(err error) error {
	return NewError(basepb.ErrorCode_EC_PARAMERR, "参数错误:" + err.Error())
}

// 业务逻辑错误
func NewBusErr(msg string) error {
	return NewError(basepb.ErrorCode_EC_BUSINESSERR, msg)
}

var InterruptErr = errors.New("gateway interrupt")
