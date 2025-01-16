package socket

import (
	"sync"

	"github.com/google/uuid"
)

type Room struct {
	id      string
	clients map[*Client]bool
	mut     sync.Mutex
}

func NewRoom() *Room {
	return &Room{
		id:      uuid.New().String(),
		clients: make(map[*Client]bool),
	}
}

func (r *Room) AddClient(client *Client) {
	r.mut.Lock()
	defer r.mut.Unlock()
	r.clients[client] = true
}

func (r *Room) RemoveClient(client *Client) {
	r.mut.Lock()
	defer r.mut.Unlock()
	if _, ok := r.clients[client]; !ok {
		return
	}
	delete(r.clients, client)
}
