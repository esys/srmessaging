package messaging

var TypeHello = "HELLO"
var TypeChat = "CHAT"
var Hello = "HELLO"
var ServerID = "SERVER"

type Message struct {
	MsgType   string `json:"type"`
	Sender    string `json:"from"`
	Recipient string `json:"to"`
	Content   string `json:"content"`
}

func NewHelloMessage(user *User) Message {
	return Message{
		MsgType:   Hello,
		Sender:    ServerID,
		Recipient: user.ID,
		Content:   Hello,
	}
}
