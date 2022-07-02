package article

import (
	"github.com/gin-gonic/gin"
	"github.com/nsocrates/go-blog/middlewares"
)

func RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/articles")

	group.Use(middlewares.AuthLock(true))
	group.GET("/feed", GetFeed)
	group.POST("/", Create)

	group.Use(middlewares.AuthLock(false))
	group.GET("/", List)
	group.GET("/:slug", Show)
}
