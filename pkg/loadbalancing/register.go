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
	AnimeServName = "anime-serv"
	CountServName = "count-serv"
)

type ServiceRegister struct {
	ctx       context.Context
	cli       *clientv3.Client
	leaseId   clientv3.LeaseID
	retry     chan struct{} // 因为网络问题断开后的重试channel
	retryTime int64         // 重试间隔时间
	key       string
	val       string
	//keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
}

const leaseTime int64 = 20

func NewServiceRegister(ctx context.Context, cli *clientv3.Client, serName, addr string) (*ServiceRegister, error) {
	s := &ServiceRegister{
		ctx:       ctx,
		cli:       cli,
		retryTime: 20,
		retry: make(chan struct{}),
		key:       "/" + schema + "/" + serName + "/" + addr,
		val:       addr,
	}
	err := s.putKeyWithLease() // 第一次服务注册
	if err != nil {
		return nil, err
	}
	go s.RegisterAndKeeplive() // 服务注册重试
	return s, nil
}

func (s *ServiceRegister) putKeyWithLease() error {
	res, err := s.cli.Grant(s.ctx, leaseTime)
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
	//s.keepAliveChan = resChan
	go func() {
		for range resChan {

		}
		// keepalive关闭
		s.retry <- struct{}{}
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

func (s *ServiceRegister) RegisterAndKeeplive() {
	select {
	case <-s.ctx.Done():
		return
	case <-s.retry:
		fmt.Println("retry")
		time.Sleep(time.Duration(s.retryTime) * time.Second)
		err := s.putKeyWithLease() // 重新进行服务注册
		if err != nil {
			log.Println("注册服务失败,err:", err.Error())
			s.retry <- struct{}{}
		}
	}
}
