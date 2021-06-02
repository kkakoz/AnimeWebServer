package kafkax

import "github.com/Shopify/sarama"

func SendSyncMsgByte(client sarama.SyncProducer, topic string, data []byte) error {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.ByteEncoder("第一个消息")
	_, _, err := client.SendMessage(msg)
	return err
}