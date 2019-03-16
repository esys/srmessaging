package messaging

import (
	log "github.com/sirupsen/logrus"

	"sync"
)

type Store struct {
	users map[string]*User
	sync.Mutex
}

func NewStore() *Store {
	return &Store{users: make(map[string]*User)}
}

func (s *Store) RegisterUser(user *User) {
	s.Lock()
	defer s.Unlock()
	s.users[user.ID] = user
}

func (s *Store) UnregisterUser(user *User) {
	s.Lock()
	defer s.Unlock()
	user.conn.Close()
	delete(s.users, user.ID)
}

func (s *Store) Deliver(m Message) {
	u, exist := s.users[m.Recipient]
	if !exist {
		log.Infof("user %s not found at our store\n", m.Recipient)
		return
	}
	u.SendMessage(m)
}
