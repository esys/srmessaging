package messaging

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	assert := assert.New(t)

	h := NewWsHandler(NewStore(), "localhost:9092")
	s := httptest.NewServer(http.HandlerFunc(h.Handle))
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")

	senderWs, _, err := websocket.DefaultDialer.Dial(u, nil)
	assert.NoErrorf(err, "error when connecting to ws")
	defer senderWs.Close()

	err = senderWs.WriteJSON(Message{MsgType: TypeHello})
	assert.NoErrorf(err, "error when writing to ws")

	var m Message
	err = senderWs.ReadJSON(&m)
	assert.NoError(err, "error when reading from ws")

	assert.Equal(TypeHello, m.MsgType)
	assert.Equal(Hello, m.Content)
	assert.NotEmpty(m.Recipient)

}
