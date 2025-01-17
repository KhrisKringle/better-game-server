package socket

import (
	"net/http"

	"github.com/brightsidedeveloper/better-game-server/internal/packets"
)

type Server struct {
	clients map[*Client]bool
	rooms   map[string]*Room
}

func NewServer() *Server {
	return &Server{
		clients: make(map[*Client]bool),
		rooms:   make(map[string]*Room),
	}
}

func (s *Server) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	c := newClient(conn)

	s.clients[c] = true

	go c.readMessages(s)
	go c.writeMessages()

	c.send <- &packets.WebSocketMessage{
		Payload: &packets.WebSocketMessage_TextMessage{
			TextMessage: &packets.TextMessage{
				Content: "Woah, it works!",
			},
		},
	}
}
