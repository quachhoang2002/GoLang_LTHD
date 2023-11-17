package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		fmt.Println(r.Host)
		// Allow all connections for this example
		return true
	},
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			http.Error(w, "Could not read message", http.StatusBadRequest)
			return
		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			http.Error(w, "Could not write message", http.StatusBadRequest)
			return
		}

	}

}

func main1() {
	http.HandleFunc("/ws", handleConnection)
	fmt.Println("WebSocket server is running on :8080/ws...")
	http.ListenAndServe(":8080", nil)
}
