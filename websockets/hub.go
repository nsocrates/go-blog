package websockets

import (
	"log"
)

var HubInstance *Hub

type Message struct {
	Channel string
	Data    interface{}
}

type Hub struct {
	clients    map[string]map[*Client]bool
	broadcast  chan *Message
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]map[*Client]bool),
	}
}

func (h *Hub) Broadcast() chan *Message {
	return h.broadcast
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			channel := client.channel
			log.Printf("registered new client to '%s'", channel)

			if _, ok := h.clients[channel]; !ok {
				h.clients[channel] = make(map[*Client]bool)
			}

			h.clients[channel][client] = true
		case client := <-h.unregister:
			channel := client.channel
			log.Printf("unregistered client from '%s'", channel)

			if _, ok := h.clients[channel][client]; ok {
				delete(h.clients[channel], client)
				close(client.send)
			}
		case message := <-h.broadcast:
			channel := message.Channel

			if clients, ok := h.clients[channel]; ok {
				for client := range clients {
					select {
					case client.send <- message.Data:
					default:
						close(client.send)
						delete(h.clients[channel], client)
					}
				}
			}
		}
	}
}
