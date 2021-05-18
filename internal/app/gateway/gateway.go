package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	proto "github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"red-bean-anime-server/pkg/gerrors"
)

type Gateway struct {
	etcdCli *clientv3.Client
	mux     *runtime.ServeMux
	ctx     context.Context
	port    string
}

func NewGateway(ctx context.Context, etcdCli *clientv3.Client, viper *viper.Viper) *Gateway {
	serveMux := runtime.NewServeMux(
		runtime.WithProtoErrorHandler(handleErr),
		runtime.WithForwardResponseOption(func(ctx context.Context, writer http.ResponseWriter, message proto.Message) error {
			type Res struct {
				Code int         `json:"code"`
				Data interface{} `json:"data"`
			}
			s := Res{Code: 200}
			s.Data = message
			bytes, err := json.Marshal(s)
			if err != nil {
				return err
			}
			_, err = writer.Write(bytes)
			if err != nil {
				return err
			}

			fmt.Println(message)
			return gerrors.InterruptErr
		}),
	)
	port := viper.Sub("gateway").GetString("port")
	return &Gateway{etcdCli: etcdCli, mux: serveMux, ctx: ctx, port: port}
}

func (g *Gateway) Start() error {
	err := Register(g.ctx, g.mux, g.etcdCli)
	if err != nil {
		return err
	}
	log.Println("gateway start")
	err = http.ListenAndServe(":"+g.port, g.mux)
	if err != nil {
		return err
	}
	return nil
}

func handleErr(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
	if err == gerrors.InterruptErr {
		return
	}

	bytes, err := marshaler.Marshal(err)
	_, err = writer.Write(bytes)
	log.Println(err)
}
