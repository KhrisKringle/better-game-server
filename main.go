package main

import (
	"log"
	"net/http"

	"github.com/brightsidedeveloper/better-game-server/internal/socket"
)

func main() {

	server := socket.NewServer()
	http.HandleFunc("/ws", server.HandleConnections)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

}
