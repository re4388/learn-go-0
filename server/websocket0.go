package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan Message)
)

// Message represents a chat message
type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

func RunWebSocketServer() {

	// setup protocol, and cb
	http.HandleFunc("/ws", handleWebSocket)

	go handleMessages()

	log.Println("Server started on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func handleWebSocket(writer http.ResponseWriter, req *http.Request) {

	// this is not safe
	// just a hacky way to let me to test locally
	// https://github.com/gorilla/websocket/issues/367
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	// setup connection
	conn, err := upgrader.Upgrade(writer, req, nil)
	if err != nil {
		log.Println("Upgrade:", err)
		return
	}

	// save this info into clients map
	clients[conn] = true

	// each each msg, send to broadcast channel
	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("ReadJSON:", err)
			delete(clients, conn)
			break
		}

		broadcast <- msg
	}

	conn.Close()
}

// send msg to connected clients
func handleMessages() {
	for {
		// get msg from broadcast channel
		msg := <-broadcast

		// write json to each client
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println("WriteJSON:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
