package messaging

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/esys/srmessaging/internal/messaging/mocks"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/mock"
)

func TestNewConsumer(t *testing.T) {
	assert := assert.New(t)

	c := NewConsumer(NewStore(), "localhost:9092")
	assert.NotNil(c)
}

func TestReadOneMessage(t *testing.T) {
	km := kafka.Message{Value: []byte("hello")}

	mockedReader := new(mocks.KafkaReader)
	mockedReader.On("ReadMessage", mock.Anything).Return(km, nil)

	c := Consumer{mockedReader, NewStore()}
	c.readOneMessage()
	mockedReader.AssertExpectations(t)
}
