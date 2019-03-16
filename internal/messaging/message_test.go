package messaging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHelloMessage(t *testing.T) {
	assert := assert.New(t)

	u := User{ID: "id"}
	m := NewHelloMessage(&u)

	assert.Equal(u.ID, m.Recipient)
	assert.Equal(Hello, m.Content)
	assert.Equal(ServerID, m.Sender)
}
