package socket

import (
	"fmt"
	"net/http"
	"sync"
)

type Server struct {
	clients map[string]*Client
	rooms   map[RoomName]*Room
	mut     sync.Mutex
}

func NewServer() *Server {
	return &Server{
		clients: make(map[string]*Client),
		rooms:   make(map[RoomName]*Room),
	}
}

func (s *Server) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	c := newClient(conn)

	s.mut.Lock()
	s.clients[c.id] = c
	s.mut.Unlock()

	s.joinRoom(c, LOBBY)

	go c.readMessages(s)
	go c.writeMessages()

	fmt.Println("Clients connected: ", len(s.clients))
}

func (s *Server) joinRoom(c *Client, name RoomName) {
	s.mut.Lock()
	defer s.mut.Unlock()

	for _, r := range s.rooms {
		if r.name == name {
			r.mut.Lock()
			fmt.Println("Joining room", r.name)
			r.clients[c.id] = c
			return
		}
	}

	fmt.Println("Creating room", name)
	room := NewRoom(name)
	room.clients[c.id] = c
	s.rooms[room.name] = room
}
