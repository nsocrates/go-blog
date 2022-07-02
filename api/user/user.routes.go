package user

import (
	"github.com/gin-gonic/gin"
	"github.com/nsocrates/go-blog/middlewares"
)

func RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/users")

	group.Use(middlewares.AuthLock(false))
	group.POST("/", Register)
	group.POST("/login", Login)

	group.Use(middlewares.AuthLock(true))
	group.GET("/", Me)
	group.PUT("/", Update)
	group.GET("/:username", Show)
	group.POST("/:username/follow", ProfileFollow)
	group.DELETE("/:username/unfollow", ProfileUnfollow)
}
