package websockets

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	ws := router.Group("/ws")
	ws.GET("/:channel", HandleWebSocket)
}
