package client

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	animepb "red-bean-anime-server/api/anime"
	"red-bean-anime-server/pkg/grpcx"
	"red-bean-anime-server/pkg/loadbalancing"
)

func NewAnimeClient(ctx context.Context, etcdClient *clientv3.Client) (animepb.AnimeServiceClient, error) {
	r := loadbalancing.NewServiceDiscovery(ctx, etcdClient)
	resolver.Register(r)
	conn, err := grpc.Dial(
		r.Scheme()+":///"+loadbalancing.AnimeServName,
		grpc.WithBalancerName("round_robin"), grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpcx.NewClientErrInterceptor(loadbalancing.AnimeServName)),
	)
	if err != nil {
		return nil, err
	}
	client := animepb.NewAnimeServiceClient(conn)
	return client, nil
}
