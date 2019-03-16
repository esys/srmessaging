package messaging

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type WritableConnection interface {
	WriteJSON(interface{}) error
	ReadJSON(interface{}) error
	Close() error
}

type User struct {
	ID   string
	conn WritableConnection
}

func NewUser(conn WritableConnection) *User {
	return &User{
		ID:   uuid.New().String(),
		conn: conn,
	}
}

func (u *User) SendMessage(m Message) {
	if err := u.conn.WriteJSON(m); err != nil {
		log.Errorf("error on message delivery to %s: %s", u.ID, err)
	} else {
		log.Infof("message sent to %s from %s", m.Recipient, m.Sender)
	}
}
