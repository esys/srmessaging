package messaging

import (
	"net/http"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type WsHandler struct {
	store    *Store
	broker   string
	upgrader websocket.Upgrader
}

func NewWsHandler(s *Store, broker string) *WsHandler {
	h := WsHandler{
		store:  s,
		broker: broker,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	return &h
}

func (h *WsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrader error %s\n" + err.Error())
		return
	}

	u := NewUser(conn)
	h.store.RegisterUser(u)
	log.Infof("user %s joined\n", u.ID)
	producer := NewProducer(h.broker)
	defer producer.close()

	for {
		var m Message
		if err := u.conn.ReadJSON(&m); err != nil {
			if websocket.IsCloseError(err, 1001) {
				log.Debugf("client %s closed ws", u.ID)
				return
			}
			log.Errorf("error receiving message on ws: %s, received %v", err, m)
		}
		log.Debugf("sending message %+v", m)
		switch m.MsgType {
		case TypeHello:
			u.conn.WriteJSON(NewHelloMessage(u))
		case TypeChat:
			producer.WriteMessage(m)
		default:
			log.Warnf("unkown message type %+v", m)
		}
	}
}
