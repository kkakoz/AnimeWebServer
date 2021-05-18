package gormx

import (
	"context"
	"google.golang.org/grpc"
)

type Msg struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func ServerErrorInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	return resp, toStatusError(err)

}

func toStatusError(err error) error {
	//if err == nil {
	//	return nil
	//}
	//cause := errors.Cause(err)
	//
	//pbErr := &basepb.Error{
	//	Code:    500,
	//	Message: err.Error(),
	//	Details: []*any.Any {
	//		{
	//			TypeUrl:
	//		},
	//	},
	//}
	//st := status.New(codes.Internal, cause.Error())
	//st, e := st.WithDetails(pbErr)
	//if e != nil {
	//	// make sure pbErr implements proto.Message interface
	//	return errors.Cause()
	//}
	//return st.Err()
	return nil
}

func ClientErrorInterceptor(ctx context.Context, method string,
	req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//log.Println("req = ", req)
	//invoker
	return nil
}
