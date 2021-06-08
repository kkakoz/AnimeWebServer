package client

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	countpb "red-bean-anime-server/api/count"
	"red-bean-anime-server/pkg/grpcx"
	"red-bean-anime-server/pkg/loadbalancing"
)

func NewCountClient(ctx context.Context, etcdClient *clientv3.Client) (countpb.CountServiceClient, error) {
	r := loadbalancing.NewServiceDiscovery(ctx, etcdClient)
	resolver.Register(r)
	conn, err := grpc.Dial(
		r.Scheme()+":///"+loadbalancing.CountServName,
		grpc.WithBalancerName("round_robin"), grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpcx.NewClientErrInterceptor(loadbalancing.CountServName)),
	)
	if err != nil {
		return nil, err
	}
	client := countpb.NewCountServiceClient(conn)
	return client, nil
}
