package socket

import (
	"sync"

	"github.com/brightsidedeveloper/better-game-server/internal/packets"
)

type RoomName string

const (
	LOBBY  RoomName = "lobby"
	BATTLE RoomName = "battle"
)

type Room struct {
	name    RoomName
	clients map[string]*Client
	mut     sync.Mutex
}

func NewRoom(name RoomName) *Room {
	return &Room{
		name:    name,
		clients: make(map[string]*Client),
	}
}

func (r *Room) broadcast(msg *packets.WSMessage) {
	r.mut.Lock()
	defer r.mut.Unlock()

	for _, c := range r.clients {
		c.send <- msg
	}
}
