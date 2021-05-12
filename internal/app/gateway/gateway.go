package gateway

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"net/http"
)

type Gateway struct {
	etcdCli *clientv3.Client
	mux     *runtime.ServeMux
	ctx     context.Context
	port    string
}

func NewGateway(ctx context.Context, etcdCli *clientv3.Client, mux *runtime.ServeMux) *Gateway {
	return &Gateway{etcdCli: etcdCli, mux: mux, ctx: ctx}
}

func (g *Gateway) Start() error {
	err := Register(g.ctx, g.mux, g.etcdCli)
	if err != nil {
		return err
	}
	err = http.ListenAndServe(g.port, g.mux)
	if err != nil {
		return err
	}
	return nil
}
