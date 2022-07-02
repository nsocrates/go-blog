package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nsocrates/go-blog/api"
	"github.com/nsocrates/go-blog/middlewares"
	"github.com/nsocrates/go-blog/websockets"
)

func ConfigureRoutes() {
	router := gin.Default()

	router.Use(middlewares.CORS())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World!"})
	})

	api.RegisterRoutes(router)
	websockets.RegisterRoutes(router)

	router.Run(":8080")
}
