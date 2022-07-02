package config

import (
	"github.com/nsocrates/go-blog/websockets"
)

func ConfigureWebSockets() {
	hub := websockets.NewHub()
	go hub.Run()

	websockets.HubInstance = hub
}
