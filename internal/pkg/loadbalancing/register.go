package loadbalancing

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"log"
)

const (
	UserServName = "user-serv"
)

type ServiceRegister struct {
	ctx           context.Context
	cli           *clientv3.Client
	leaseId       clientv3.LeaseID
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	key           string
	val           string
}

func NewServiceRegister(ctx context.Context, cli *clientv3.Client, serName, addr string, lease int64) (*ServiceRegister, error) {
	s := &ServiceRegister{
		ctx: ctx,
		cli: cli,
		key: "/" + schema + "/" + serName + "/" + addr,
		val: addr,
	}
	err := s.putKeyWithLease(lease)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *ServiceRegister) putKeyWithLease(lease int64) error {
	res, err := s.cli.Grant(s.ctx, lease)
	if err != nil {
		return err
	}
	s.leaseId = res.ID
	_, err = s.cli.Put(s.ctx, s.key, s.val, clientv3.WithLease(s.leaseId))
	if err != nil {
		return err
	}
	resChan, err := s.cli.KeepAlive(s.ctx, s.leaseId)
	if err != nil {
		return err
	}
	s.keepAliveChan = resChan
	log.Println("put key:%s val")
	return nil
}

func (s *ServiceRegister) Close() error {
	if _, err := s.cli.Revoke(s.ctx, s.leaseId); err != nil {
		return err
	}
	return nil
}
