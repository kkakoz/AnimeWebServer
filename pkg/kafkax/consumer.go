package kafkax

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type consumerOptions struct {
	Address []string
	Topic   []string
	GroupId string
}

type ConsumerRun struct {
	ctx      context.Context
	consumer *cluster.Consumer
	logger   *zap.Logger
	runFunc  ConsumerRunFunc
}

type ConsumerRunFunc func(*sarama.ConsumerMessage) error

func NewConsumer(ctx context.Context, viper *viper.Viper, logger *zap.Logger, runFunc ConsumerRunFunc) (*ConsumerRun, error) {
	o := &consumerOptions{}
	viper.SetDefault("kafka.address", []string{"127.0.0.1:9092"})
	err := viper.UnmarshalKey("kafka", o)
	if err != nil {
		return nil, err
	}

	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true

	consumer, err := cluster.NewConsumer(o.Address, o.GroupId, o.Topic, config)
	if err != nil {
		return nil, err
	}

	run := &ConsumerRun{
		ctx:      ctx,
		consumer: consumer,
		logger:   logger,
		runFunc:  runFunc,
	}

	return run, nil
}

func (c *ConsumerRun) Run() {
	// consume messages, watch signals
	for {
		select {
		case err := <- c.consumer.Errors():
			c.logger.Error(fmt.Sprintf("consumer err: %+v\n", err))
		case ntf := <- c.consumer.Notifications():
			c.logger.Info(fmt.Sprintf("Rebalanced: %+v\n", ntf))
		case msg, ok := <-c.consumer.Messages():
			if ok {
				err := c.runFunc(msg)
				if err != nil {
					c.logger.Error(fmt.Sprintf("consumer msg err: %+v\n", err), zap.String("msg", fmt.Sprintf("%+v", msg)))

				}
				c.consumer.MarkOffset(msg, "") // mark message as processed
			}
		case <-c.ctx.Done():
			return
		}
	}
}

var KafkaConsumerSet = wire.NewSet(NewConsumer)