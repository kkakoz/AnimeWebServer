package gateway

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	animepb "red-bean-anime-server/api/anime"
	userpb "red-bean-anime-server/api/user"
	"red-bean-anime-server/internal/pkg/client"
)

func Register(ctx context.Context, s *runtime.ServeMux, etcdCli *clientv3.Client) error {
	userClient, err := client.NewUserClient(ctx, etcdCli)
	if err != nil {
		return err
	}
	err = userpb.RegisterUserServiceHandlerClient(ctx, s, userClient)
	if err != nil {
		return err
	}
	animeClient, err := client.NewAnimeClient(ctx, etcdCli)
	if err != nil {
		return err
	}
	err = animepb.RegisterAnimeServiceHandlerClient(ctx, s, animeClient)
	if err != nil {
		return err
	}
	return nil
}

