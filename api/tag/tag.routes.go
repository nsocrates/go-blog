package tag

import (
	"github.com/gin-gonic/gin"
	"github.com/nsocrates/go-blog/middlewares"
)

func RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/tags")

	group.Use(middlewares.AuthLock(false))
	group.GET("/", List)
}
