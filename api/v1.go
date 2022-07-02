package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nsocrates/go-blog/api/article"
	"github.com/nsocrates/go-blog/api/article/comment"
	"github.com/nsocrates/go-blog/api/tag"
	"github.com/nsocrates/go-blog/api/user"
)

func RegisterRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	user.RegisterRoutes(v1)
	tag.RegisterRoutes(v1)
	article.RegisterRoutes(v1)
	comment.RegisterRoutes(v1)
}
