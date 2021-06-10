package gateway

import (
	"context"
	"encoding/json"
	"github.com/coreos/etcd/clientv3"
	proto "github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	basepb "red-bean-anime-server/api/base"
	"red-bean-anime-server/pkg/gerrors"
)

type Gateway struct {
	etcdCli *clientv3.Client
	mux     *runtime.ServeMux
	ctx     context.Context
	port    string
	logger  *zap.Logger
}

type Res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewGateway(ctx context.Context, etcdCli *clientv3.Client, viper *viper.Viper, logger *zap.Logger) *Gateway {
	serveMux := runtime.NewServeMux(
		runtime.WithProtoErrorHandler(NewHandleErr(logger)),
		runtime.WithForwardResponseOption(NewResHandler()),
	)
	port := viper.Sub("gateway").GetString("port")
	return &Gateway{etcdCli: etcdCli, mux: serveMux, ctx: ctx, port: port, logger: logger}
}

func (g *Gateway) Start() error {
	err := Register(g.ctx, g.mux, g.etcdCli)
	if err != nil {
		return err
	}
	g.logger.Info("gateway start")
	err = http.ListenAndServe(":"+g.port, g.mux)
	if err != nil {
		return err
	}
	return nil
}

func NewHandleErr(logger *zap.Logger) runtime.ProtoErrorHandlerFunc {
	return func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
		log.Println("http url = ", request.RequestURI)
		if err == gerrors.InterruptErr {
			return
		}
		s := Res{
			Code: 200,
		}
		statusErr, ok := status.FromError(err)
		s.Msg = err.Error()
		if ok && int32(statusErr.Code()) == int32(basepb.ErrorCode_EC_BUSINESSERR) { //
			s.Code = int(basepb.ErrorCode_EC_BUSINESSERR)
			s.Msg = statusErr.Message()
			errs := make([]zap.Field, 0, len(statusErr.Proto().GetDetails()))
			for _, detail := range statusErr.Proto().GetDetails() {
				errs = append(errs, zap.Error(errors.New(string(detail.Value))))

			}
			logger.Error(s.Msg, errs...)
		}
		bytes, err := marshaler.Marshal(s)
		if err != nil {
			logger.Error("err = ", zap.Error(err))
		}
		_, err = writer.Write(bytes)
		if err != nil {
			logger.Error("err = ", zap.Error(err))
		}
	}
}

func NewResHandler() func(context.Context, http.ResponseWriter, proto.Message) error {
	return func(ctx context.Context, writer http.ResponseWriter, message proto.Message) error {
		s := Res{Code: 200}
		s.Data = message
		if message != nil {
			bytes, err := json.Marshal(s)
			if err != nil {
				return err
			}
			_, err = writer.Write(bytes)
			if err != nil {
				return err
			}
		}
		return gerrors.InterruptErr
	}
}
