package messaging

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"

	"github.com/esys/srmessaging/internal/messaging/mocks"
)

func TestNewProducer(t *testing.T) {
	assert := assert.New(t)

	p := NewProducer("localhost:9092")
	assert.NotNil(p)
}

func TestWriteMessage(t *testing.T) {
	mockedWriter := new(mocks.KafkaWriter)
	mockedWriter.On("WriteMessages", mock.Anything, mock.Anything).Return(nil)

	p := Producer{writer: mockedWriter}
	p.WriteMessage(Message{Sender: "sendId", Content: "hello"})
	mockedWriter.AssertExpectations(t)
}
