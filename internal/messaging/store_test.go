package messaging

import (
	"testing"

	"github.com/esys/srmessaging/internal/messaging/mocks"
	"github.com/stretchr/testify/assert"
)

func createStore() Store {
	return Store{users: make(map[string]*User)}
}
func TestRegisterUser(t *testing.T) {
	assert := assert.New(t)
	store := createStore()

	want := User{"userid", nil}
	store.RegisterUser(&want)
	got := store.users["userid"]

	assert.EqualValues(want.ID, got.ID)
}

func TestUnregisterUser(t *testing.T) {
	assert := assert.New(t)
	store := createStore()

	mockedConn := new(mocks.WritableConnection)
	mockedConn.On("Close").Return(nil)

	u := User{"userid", mockedConn}
	store.users[u.ID] = &u
	store.UnregisterUser(&u)

	_, present := store.users[u.ID]
	assert.False(present)
	mockedConn.AssertExpectations(t)
}

func TestDeliver(t *testing.T) {
	assert := assert.New(t)
	store := createStore()

	mockedConn := new(mocks.WritableConnection)
	mockedConn.On("Close").Return(nil)

	u := User{"userid", mockedConn}
	store.users[u.ID] = &u
	store.UnregisterUser(&u)

	_, present := store.users[u.ID]
	assert.False(present)
	mockedConn.AssertExpectations(t)

}
