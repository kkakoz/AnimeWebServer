package client

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	userpb "red-bean-anime-server/api/user"
	"red-bean-anime-server/pkg/grpcx"
	"red-bean-anime-server/pkg/loadbalancing"
)

func NewUserClient(ctx context.Context, etcdClient *clientv3.Client) (userpb.UserServiceClient, error) {
	r := loadbalancing.NewServiceDiscovery(ctx, etcdClient)
	resolver.Register(r)
	conn, err := grpc.Dial(
		r.Scheme()+":///"+loadbalancing.UserServName,
		grpc.WithBalancerName("round_robin"), grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpcx.NewClientErrInterceptor(loadbalancing.UserServName)),
	)
	if err != nil {
		return nil, err
	}
	client := userpb.NewUserServiceClient(conn)
	return client, nil
}
