package socket

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	id      string
	conn    *websocket.Conn
	channel chan []byte
}

func newClient(conn *websocket.Conn) *Client {
	return &Client{
		id:      uuid.New().String(),
		conn:    conn,
		channel: make(chan []byte),
	}
}
