package socket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn   *websocket.Conn
	room   *Room
	id     string
	send   chan []byte
	server *Server
}

type Room struct {
	id      string
	clients map[*Client]bool
	lock    sync.Mutex
}

type Server struct {
	clients     map[*Client]bool
	rooms       map[string]*Room
	gameManager *GameManager
	lock        sync.Mutex
}

func NewServer() *Server {
	return &Server{
		clients:     make(map[*Client]bool),
		rooms:       make(map[string]*Room),
		gameManager: NewGameManager(),
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *Server) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := &Client{
		id:     fmt.Sprintf("%p", conn),
		conn:   conn,
		send:   make(chan []byte),
		server: s,
	}

	s.lock.Lock()
	s.clients[client] = true
	s.lock.Unlock()

	s.joinRoom(client, "lobby")

	go client.readMessages()
	go client.writeMessages()
}

func (s *Server) joinRoom(client *Client, roomID string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if client.room != nil {
		s.leaveRoom(client)
	}

	room, ok := s.rooms[roomID]
	if !ok {
		room = &Room{
			id:      roomID,
			clients: make(map[*Client]bool),
		}
		s.rooms[roomID] = room
	}

	room.lock.Lock()
	room.clients[client] = true
	room.lock.Unlock()

	client.room = room
	log.Printf("Client %s joined room %s", client.id, room.id)
}

func (s *Server) leaveRoom(client *Client) {
	if client.room == nil {
		return
	}

	client.room.lock.Lock()
	delete(client.room.clients, client)
	client.room.lock.Unlock()

	if len(client.room.clients) == 0 {
		s.lock.Lock()
		delete(s.rooms, client.room.id)
		s.lock.Unlock()
	}

	client.room = nil
	log.Printf("Client %s left room", client.id)
}

func (c *Client) readMessages() {
	defer func() {
		c.server.leaveRoom(c)
		c.conn.Close()
	}()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		var actionMsg GameActionMessage
		if err := json.Unmarshal(msg, &actionMsg); err != nil {
			log.Println("Invalid message format")
			continue
		}

		c.server.lock.Lock()
		game, ok := c.server.gameManager.Games[actionMsg.Player]
		c.server.lock.Unlock()

		log.Printf("Received message from client %s: %s", c.id, msg)

		if ok {
			game.Actions <- actionMsg
		}

	}
}

func (c *Client) writeMessages() {
	for msg := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}

func (c *Client) SendMessage(msg string) {
	c.send <- []byte(msg)
}

func (s *Server) sendMessage(client *Client, msgType, msg string) {
	client.SendMessage(fmt.Sprintf(`{"type":"%s","msg":"%s"}`, msgType, msg))
}

func (s *Server) InviteToBattle(sender *Client, receiverID string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	var receiver *Client
	for c := range s.clients {
		if c.id == receiverID {
			receiver = c
			break
		}
	}
	if receiver == nil {
		s.sendMessage(sender, "error", "User not found")
		return
	}

	// Create a new game
	gameID := fmt.Sprintf("battle-%s-%s", sender.id, receiver.id)
	newGame := &Game{
		ID:      gameID,
		Players: []*Client{sender, receiver},
		State:   Waiting,
		Actions: make(chan GameActionMessage, 10),
	}

	s.gameManager.Lock.Lock()
	s.gameManager.Games[gameID] = newGame
	s.gameManager.Lock.Unlock()

	// Move players into the new room
	s.joinRoom(sender, gameID)
	s.joinRoom(receiver, gameID)

	// Start processing game logic
	go newGame.ProcessActions()

	// Notify both players
	s.sendMessage(sender, "info", "Waiting for opponent to accept...")
	s.sendMessage(receiver, "invite", fmt.Sprintf(`{"game_id": "%s", "opponent": "%s"}`, gameID, sender.id))
}

type GameState string

const (
	Waiting    GameState = "waiting"
	InProgress GameState = "inProgress"
	Complete   GameState = "complete"
)

type GameAction string

const (
	Invite  GameAction = "invite"
	Accept  GameAction = "accept"
	EndGame GameAction = "endGame"
)

type Game struct {
	ID      string
	Players []*Client
	State   GameState
	Actions chan GameActionMessage
	Lock    sync.Mutex
}

type GameActionMessage struct {
	Action GameAction `json:"action"`
	Player string     `json:"player"`
	Target string     `json:"target"`
}

type GameManager struct {
	Games map[string]*Game
	Lock  sync.Mutex
}

func NewGameManager() *GameManager {
	return &GameManager{
		Games: make(map[string]*Game),
	}
}

func (g *Game) ProcessActions() {
	for a := range g.Actions {
		g.Lock.Lock()
		switch a.Action {
		case Invite:
			for _, p := range g.Players {
				if p.id == a.Target {
					p.SendMessage(fmt.Sprintf(`{"type":"invite","player":"%s"}`, a.Player))
				}
			}
		case Accept:
			log.Printf("Player %s accepted invite to game %s", a.Player, g.ID)
			g.State = InProgress
			g.broadcastState("Game Started")
		case EndGame:
			log.Printf("Player %s ended game %s", a.Player, g.ID)
			g.State = Complete
			g.broadcastState("Game Ended")
		}

		g.Lock.Unlock()
	}
}

func (g *Game) broadcastState(state string) {
	for _, p := range g.Players {
		p.SendMessage(fmt.Sprintf(`{"type": "game_state", "message": "%s"}`, state))
	}
}
