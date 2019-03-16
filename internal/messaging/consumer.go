package messaging

import (
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/segmentio/kafka-go"
)

type KafkaReader interface {
	ReadMessage(context.Context) (kafka.Message, error)
	Close() error
}

type Consumer struct {
	reader KafkaReader
	store  *Store
}

func NewConsumer(s *Store, broker string) *Consumer {
	consumer := Consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:   []string{broker},
			Topic:     "topic-A",
			Partition: 0,
			MinBytes:  10e3, // 10KB
			MaxBytes:  10e6, // 10MB
		}),
		store: s,
	}
	return &consumer
}

func (c *Consumer) readOneMessage() {
	kafkaMess, err := c.reader.ReadMessage(context.Background())
	if err != nil {
		log.Errorf("error when reading message: %s", err)
	}
	log.Debugf("message at topic/partition/offset %v/%v/%v: %s = %s\n", kafkaMess.Topic, kafkaMess.Partition, kafkaMess.Offset, string(kafkaMess.Key), string(kafkaMess.Value))
	var js Message
	json.Unmarshal(kafkaMess.Value, &js)
	log.Debugf("read json message: %s", js)
	c.store.Deliver(js)
}

func (c *Consumer) ReadMessages() {
	for {
		c.readOneMessage()
	}
}

func (c *Consumer) Close() {
	c.reader.Close()
}
