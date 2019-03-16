package messaging

import (
	"testing"

	"github.com/esys/srmessaging/internal/messaging/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewUserHasID(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(NewUser(nil).ID)
}

func TestSendMessage(t *testing.T) {
	mess := Message{Sender: "sender", Recipient: "recipient", Content: "hello"}
	mockedConn := new(mocks.WritableConnection)
	mockedConn.On("WriteJSON", mock.AnythingOfType("Message")).Return(nil)

	user := User{ID: "userid", conn: mockedConn}
	user.SendMessage(mess)
	mockedConn.AssertExpectations(t)
}
