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
	room *Room
	send chan *packets.WSMessage
}

func newClient(conn *websocket.Conn) *Client {
	client := &Client{
		id:   uuid.New().String(),
		conn: conn,
		send: make(chan *packets.WSMessage),
	}

	return client
}

func (c *Client) readMessages(s *Server) {
	defer func() {
		delete(s.clients, c.id)
		c.conn.Close()
	}()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}

		var message packets.WSMessage
		if err := proto.Unmarshal(msg, &message); err != nil {
			fmt.Println(err)
			continue
		}

		switch payload := message.Payload.(type) {
		case *packets.WSMessage_TextMessage:
			msg := &packets.WSMessage{
				Payload: &packets.WSMessage_TextMessage{
					TextMessage: &packets.TextMessage{
						Content: payload.TextMessage.Content,
					},
				},
			}
			c.room.broadcast(msg)
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

func (c *Client) write(msg *packets.WSMessage) {
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
