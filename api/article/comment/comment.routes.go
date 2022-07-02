package comment

import (
	"github.com/gin-gonic/gin"
	"github.com/nsocrates/go-blog/middlewares"
)

func RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/articles/:slug/comments")

	group.Use(middlewares.AuthLock(true))
	group.POST("/", Create)
	group.DELETE("/:id", Destroy)

	group.Use(middlewares.AuthLock(false))
	group.GET("/", List)
}
