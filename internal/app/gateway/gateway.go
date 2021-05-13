package gateway

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type Gateway struct {
	etcdCli *clientv3.Client
	mux     *runtime.ServeMux
	ctx     context.Context
	port    string
}

func NewGateway(ctx context.Context, etcdCli *clientv3.Client, viper *viper.Viper) *Gateway {
	serveMux := runtime.NewServeMux()
	port := viper.Sub("gateway").GetString("port")
	return &Gateway{etcdCli: etcdCli, mux: serveMux, ctx: ctx, port: port}
}

func (g *Gateway) Start() error {
	err := Register(g.ctx, g.mux, g.etcdCli)
	if err != nil {
		return err
	}
	log.Println("gateway start")
	err = http.ListenAndServe(":" + g.port, g.mux)
	if err != nil {
		return err
	}
	return nil
}
