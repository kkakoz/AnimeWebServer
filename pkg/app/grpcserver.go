package app

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"net"
	"red-bean-anime-server/pkg/grpcx"
	"red-bean-anime-server/pkg/loadbalancing"
)

type GrpcServer struct {
	ctx            context.Context
	etcdCli        *clientv3.Client
	grpcServer     *grpc.Server
	registerServer RegisterService
}

type RegisterService func(server *grpc.Server)

func NewGrpcServer(ctx context.Context, client *clientv3.Client, register RegisterService) *GrpcServer {
	g := &GrpcServer{}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,
			grpcx.ServerErrorInterceptor,
			grpc_recovery.UnaryServerInterceptor(grpcx.RecoveryInterceptor()),
		)),
	)

	g.grpcServer = server
	g.etcdCli = client
	g.ctx = ctx
	g.registerServer = register
	return g
}

func (s *GrpcServer) run(servname, host, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	s.registerServer(s.grpcServer)
	_, err = loadbalancing.NewServiceRegister(s.ctx, s.etcdCli, servname, host+":"+port)
	if err != nil {
		return err
	}
	err = s.grpcServer.Serve(listen)
	if err != nil {
		return err
	}
	return nil
}
