package logic

import "time"

//消息
type Message struct {
	Player         *Player   `json:"player"`
	Content        string    `json:"content"`
	MsgTime        time.Time `json:"msgtime"`
	ClientSendTime time.Time `json:"client_send_time"`
}

func NewMessage(p *Player, content string) *Message {
	message := &Message{
		Player:  p,
		Content: content,
		MsgTime: time.Now(),
	}
	return message
}

func NewWelcomeMessage(p *Player) *Message {
	return &Message{
		Player: &Player{
			Name: "[System]",
		},
		Content: p.Name + " hello, welcome chat-room！",
		MsgTime: time.Now(),
	}
}

func NewPlayerEnterMessage(p *Player) *Message {
	return &Message{
		Player:  p,
		Content: p.Name + " enter chat-room",
		MsgTime: time.Now(),
	}
}

func NewPlayerLeaveMessage(p *Player) *Message {
	return &Message{
		Player:  p,
		Content: p.Name + " leave chat-room",
		MsgTime: time.Now(),
	}
}

func NewErrorMessage(err string) *Message {
	return &Message{
		Player: &Player{
			Name: "[Error]",
		},
		Content:        err,
		MsgTime:        time.Time{},
		ClientSendTime: time.Time{},
	}
}

func NewSystemMessage(notify string) *Message {
	return &Message{
		Player: &Player{
			Name: "[System]",
		},
		Content:        notify,
		MsgTime:        time.Time{},
		ClientSendTime: time.Time{},
	}
}
