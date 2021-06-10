package gerrors

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/any"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"red-bean-anime-server/pkg/log"
)

const TypeUrlStack = "type_url_stack"

func WrapError(err error) error {
	if err == nil {
		return nil
	}
	log.L().Error(fmt.Sprintf("%+v", err))
	statusErr, ok := status.FromError(err)
	if ok {
		return statusErr.Err()
	}
	spbErr := &spb.Status{
		Code:    int32(codes.Unknown),
		Message: err.Error(),
		Details: []*any.Any{
			{
				TypeUrl: TypeUrlStack,
				Value:   []byte(fmt.Sprintf("%+v", err)),
			},
		},
	}

	return status.FromProto(spbErr).Err()
}

func WrapRPCError(err error, servName string) error {
	if err == nil {
		return nil
	}
	statusErr, ok := status.FromError(err)
	if ok {
		details := statusErr.Proto().GetDetails()
		detailStr := ""
		for _, v := range details {
			detailStr += string(v.Value)
		}
		spbErr := &spb.Status{
			Code:    int32(statusErr.Code()),
			Message: statusErr.Message(),
			Details: []*any.Any{
				{
					TypeUrl: TypeUrlStack,
					Value:   []byte(fmt.Sprintf("%+v\n---grpc %s call----%s\n", err, servName, detailStr)),
				},
			},
		}
		return status.FromProto(spbErr).Err()
	}
	spbErr := &spb.Status{
		Code:    int32(codes.Unknown),
		Message: err.Error(),
		Details: []*any.Any{
			{
				TypeUrl: TypeUrlStack,
				Value:   []byte(fmt.Sprintf("%+v", err)),
			},
		},
	}
	return status.FromProto(spbErr).Err()
}
