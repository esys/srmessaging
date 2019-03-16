package messaging

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/segmentio/kafka-go"
)

type KafkaWriter interface {
	WriteMessages(context.Context, ...kafka.Message) error
	Close() error
}

type Producer struct {
	writer KafkaWriter
}

func NewProducer(broker string) *Producer {
	producer := Producer{
		writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers:  []string{broker},
			Topic:    "topic-A",
			Balancer: &kafka.LeastBytes{},
		}),
	}

	return &producer
}

func (p *Producer) WriteMessage(mess Message) {
	m, err := json.Marshal(mess)
	if err != nil {
		log.Errorf("failed to marshal json message: %s", err)
		return
	}
	err = p.writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(uuid.New().String()),
			Value: m,
		},
	)
	if err != nil {
		log.Errorf("failed to write message %+v to kafka: %s", m, err)
	}
}

func (p *Producer) close() {
	p.writer.Close()
}
