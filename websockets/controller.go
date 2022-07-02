package websockets

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func HandleWebSocket(c *gin.Context) {
	channel := c.Param("channel")
	ServeWS(c.Writer, c.Request, channel)
}

func ServeWS(w http.ResponseWriter, r *http.Request, channel string) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	client := NewClient(HubInstance, conn, channel)
	HubInstance.register <- client

	go client.WritePump()
	go client.ReadPump()
}
