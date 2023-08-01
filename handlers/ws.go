package handlers

import (
	"log"

	"github.com/gofiber/contrib/websocket"
)

func WSHandler(c *websocket.Conn) {
	go RunHub()

	register <- c

	defer func() {
		unregister <- c
		c.Close()
	}()

	for {
		messageType, message, err := c.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("read error:", err)
			}

			return
		}

		if messageType == websocket.TextMessage {
			broadcast <- string(message)
		} else {
			log.Println("websocket message received of type", messageType)
		}
	}
}
