package main

import (
	"github.com/nsocrates/go-blog/config"
)

func main() {
	config.ConfigureDatabase()
	config.ConfigureWebSockets()
	config.ConfigureRoutes()
}
