package logic

import "container/ring"

// 保存最近N条消息
type lastMessage struct {
	n        int
	lastRing *ring.Ring
}

func newLastMessage(n int) *lastMessage {
	if n < 0 {
		panic("save last message len < 0")
	}
	return &lastMessage{
		n:        n,
		lastRing: ring.New(n),
	}
}

func (lm *lastMessage) Save(msg *Message) {
	lm.lastRing.Value = msg
	lm.lastRing = lm.lastRing.Next()
}

func (lm *lastMessage) Send(p *Player) {
	lm.lastRing.Do(func(value interface{}) {
		if value != nil {
			p.SendMessage(value.(*Message))
		}
	})
}
