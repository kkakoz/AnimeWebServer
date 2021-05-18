package loadbalancing

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"log"
	"time"
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

const leaseTime int64 = 20

func NewServiceRegister(ctx context.Context, cli *clientv3.Client, serName, addr string) (*ServiceRegister, error) {
	s := &ServiceRegister{
		ctx: ctx,
		cli: cli,
		key: "/" + schema + "/" + serName + "/" + addr,
		val: addr,
	}
	log.Println("key = ", s.key)
	err := s.putKeyWithLease(leaseTime)
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
	go func() {
		i := 0
		for range s.keepAliveChan {
			i++
			fmt.Println("kepp alive time = ", time.Now(), " i = ", i)
		}
		fmt.Println("for loop is intercepter")
	}()
	return nil
}

func (s *ServiceRegister) Close() error {
	log.Println("invoke close")
	if _, err := s.cli.Revoke(s.ctx, s.leaseId); err != nil {
		return err
	}
	return nil
}
