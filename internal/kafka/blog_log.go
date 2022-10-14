package kafka

import (
	"errors"
	"github.com/Shopify/sarama"
	"github.com/go-programming-tour-book/blog-service/global"
)

type BlogLog struct {
	Topic   string
	Connect sarama.SyncProducer
}

func NewBlogLog() *BlogLog {
	topic := getTopic()
	connect, err := GetNewKafkaConnect()
	if err != nil {
		return &BlogLog{}
	}

	return &BlogLog{
		Topic:   topic,
		Connect: connect,
	}
}

func getTopic() string {
	return "blog_log"
}

func GetNewKafkaConnect() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //follow同步数据后返回
	config.Producer.Partitioner = sarama.NewRandomPartitioner //随机分配分区 partition
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer(global.KafkaSetting.Host, config)
	if err != nil {
		return nil, errors.New("kafaka connect fail")
	}

	return client, nil
}

func (b *BlogLog) SendMessage(message string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = b.Topic
	msg.Value = sarama.StringEncoder(message)
	b.Connect.SendMessage(msg)
}
