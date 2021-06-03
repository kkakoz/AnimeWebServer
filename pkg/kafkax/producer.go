package kafkax

import (
	"github.com/Shopify/sarama"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type producerOptions struct {
	address []string
}

func NewSyncProducer(viper *viper.Viper) (sarama.SyncProducer, error) {
	o := &producerOptions{}
	viper.SetDefault("kafka.address", []string{"127.0.0.1:9092"})
	err := viper.UnmarshalKey("kafka", o)
	if err != nil {
		return nil, err
	}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // ack确认机制
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 选择分区-随机分区
	config.Producer.Return.Successes = true // 确认

	// 连接kafka
	client, err := sarama.NewSyncProducer(o.address, config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func SendSyncMsgByte(client sarama.SyncProducer, topic string, data []byte) error {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.ByteEncoder(data)
	_, _, err := client.SendMessage(msg)
	return err
}

var KafkaProducerSet = wire.NewSet(NewSyncProducer)