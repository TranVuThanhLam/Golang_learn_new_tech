package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// config update HTTP to websocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(request *http.Request) bool {
		return true
	},
}

type Client struct {
	Conn *websocket.Conn
}

var (
	clients   = make(map[*Client]bool)
	broadcast = make(chan Message)
	mutex     sync.Mutex
)

// define message structure
type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

func handlerConnections(context *gin.Context) {
	// Upgrade HTTP connection to websocket
	websocket, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		fmt.Println("error while upgrading to Websocket", err)
	}
	defer websocket.Close()

	client := &Client{Conn: websocket}

	mutex.Lock()
	clients[client] = true
	mutex.Unlock()

	// listen message from client
	for {
		var message Message
		err := websocket.ReadJSON(&message)
		if err != nil {
			mutex.Lock()
			delete(clients, client)
			mutex.Unlock()
			fmt.Println("Client disconnect: ", err)
			break
		}
		// send message to all client
		broadcast <- message
	}

}

func handleMessages() {
	for {
		message := <-broadcast
		mutex.Lock()
		for client := range clients {
			err := client.Conn.WriteJSON(message)
			if err != nil {
				fmt.Println("error when send message", err)
				client.Conn.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}

func main() {
	r := gin.Default()

	// router socket
	r.GET("/ws", handlerConnections)

	// gorotine to process message send
	go handleMessages()
	r.Run(":8080")
}
