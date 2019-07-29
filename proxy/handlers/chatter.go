package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// HandleSocketRequest from another server
func HandleSocketRequest(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			log.Println("ERROR:", err)
			continue
		}

		log.Println("MESSAGE:", string(message))

		err = conn.WriteMessage(websocket.TextMessage, message)

		if err != nil {
			log.Println("ERROR:", err)
			continue
		}
	}
}
