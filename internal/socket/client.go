package socket

import (
	"fmt"

	"github.com/brightsidedeveloper/better-game-server/internal/packets"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	id   string
	conn *websocket.Conn
	send chan *packets.WebSocketMessage
}

func newClient(conn *websocket.Conn) *Client {
	client := &Client{
		id:   uuid.New().String(),
		conn: conn,
		send: make(chan *packets.WebSocketMessage),
	}

	return client
}

func (c *Client) readMessages(s *Server) {
	defer func() {
		delete(s.clients, c)
		c.conn.Close()
	}()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}

		var message packets.WebSocketMessage
		if err := proto.Unmarshal(msg, &message); err != nil {
			fmt.Println(err)
			continue
		}

		switch payload := message.Payload.(type) {
		case *packets.WebSocketMessage_TextMessage:
			fmt.Println("TextMessage", payload.TextMessage)
		case *packets.WebSocketMessage_UserAction:
			fmt.Println("UserAction", payload.UserAction)
		case *packets.WebSocketMessage_GameInvite:
			fmt.Println("GameInvite", payload.GameInvite)
		default:
			fmt.Println("Unknown message type")
		}
	}
}

func (c *Client) writeMessages() {
	for msg := range c.send {
		c.write(msg)
	}
}

func (c *Client) write(msg *packets.WebSocketMessage) {
	data, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.conn.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		fmt.Println(err)
		return
	}
}
