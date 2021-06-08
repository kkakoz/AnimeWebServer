package echox

import (
	"github.com/labstack/echo"
	"google.golang.org/grpc/status"
	"red-bean-anime-server/pkg/log"
)

func HandlerErr() echo.HTTPErrorHandler {
	return func(err error, context echo.Context) {
		if err != nil {
			statusErr, ok := status.FromError(err)
			if ok {
				err := ToErr(context, statusErr.Err())
				if err != nil {
					log.L().Error(statusErr.Proto().Message)
				}
			}
		}
	}
}