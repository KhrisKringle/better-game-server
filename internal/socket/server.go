package socket

import "net/http"

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
	defer conn.Close()

	client := newClient(conn)

	s.clients[client] = true

	for {

		_, msg, err := conn.ReadMessage()
		if err != nil {
			delete(s.clients, client)
			break
		}

		for c := range s.clients {
			c.channel <- msg
		}

	}
}
