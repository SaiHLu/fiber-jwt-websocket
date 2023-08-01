package handlers

import (
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type client struct {
	isClosing bool
	mu        sync.Mutex
}

var (
	clients    = make(map[*websocket.Conn]*client)
	register   = make(chan *websocket.Conn)
	broadcast  = make(chan string)
	unregister = make(chan *websocket.Conn)
)

func RunHub() {
	for {
		select {
		case connection := <-register:
			clients[connection] = &client{}
			log.Println("Connection registered")

		case message := <-broadcast:
			log.Println("Message received: ", message)

			for connection, c := range clients {
				go func(connection *websocket.Conn, client *client) {
					c.mu.Lock()
					defer c.mu.Unlock()

					if c.isClosing {
						return
					}

					if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
						c.isClosing = true
						log.Println("write error:", err)

						connection.WriteMessage(websocket.CloseMessage, []byte{})
						connection.Close()
						unregister <- connection
					}
				}(connection, c)
			}

		case connection := <-unregister:
			delete(clients, connection)
			log.Println("connection unregistered")

		}

	}
}
