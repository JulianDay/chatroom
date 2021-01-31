package logic

import (
	"sync"
)

// 房间
//1.管理用户
type room struct {
	players map[string]*Player
	mu      sync.Mutex
}

func NewRoom() *room {
	return &room{
		players: make(map[string]*Player),
	}
}

func (r *room) Broadcast(msg *Message) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, p := range r.players {
		p.SendMessage(msg)
	}
}

func (r *room) PlayerEnter(p *Player) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.players[p.Name] = p

	LastMessage.Send(p)
}

func (r *room) PlayerLeaving(p *Player) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.players, p.Name)
}

func (r *room) CheckName(name string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.players[name]; ok {
		return false
	} else {
		return true
	}
}

func (r *room) GetPlayer(name string) *Player {
	r.mu.Lock()
	defer r.mu.Unlock()
	if p, ok := r.players[name]; ok {
		return p
	} else {
		return nil
	}
}
