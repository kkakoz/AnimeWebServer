package app

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"red-bean-anime-server/pkg/loadbalancing"
)

type GrpcServer struct {
	ctx        context.Context
	etcdCli    *clientv3.Client
	grpcServer *grpc.Server
}

type RegisterService func(server *grpc.Server)

func NewGrpcServer(ctx context.Context, viper *viper.Viper, client *clientv3.Client, register RegisterService) (*GrpcServer, error) {
	g := &GrpcServer{}
	server := grpc.NewServer(
		//grpc.UnaryInterceptor(gormx.ServerErrorInterceptor),
	)
	if err := viper.UnmarshalKey("grpc-server", g); err != nil {
		return nil, err
	}
	register(server)
	g.grpcServer = server
	g.etcdCli = client
	g.ctx = ctx
	return g, nil
}

func (s *GrpcServer) run(servname, host, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
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
