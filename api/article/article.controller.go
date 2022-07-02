package article

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	a "github.com/nsocrates/go-blog/api/common"
	"github.com/nsocrates/go-blog/common"
)

func GetFeed(c *gin.Context) {}

func List(c *gin.Context) {
	tag := c.Query("tag")
	author := c.Query("author")
	limit := c.Query("limit")
	offset := c.Query("offset")
	favorited := c.Query("favorited")
	articles, articlesCount, err := FindMany(tag, author, limit, offset, favorited)

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("articles", errors.New("invalid param")))
		return
	}

	serializer := ArticlesSerializer{c, articles}
	c.JSON(http.StatusOK, gin.H{"articles": serializer.Response(), "articlesCount": articlesCount})
}

func Show(c *gin.Context) {
	slug := c.Param("slug")
	article, err := FindOne(&a.Article{Slug: slug})

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("articles", errors.New("invalid slug")))
		return
	}

	serializer := ArticleSerializer{c, article}
	response := serializer.Response()
	c.JSON(http.StatusOK, gin.H{"article": response})
}

func Create(c *gin.Context) {
	articleValidator := NewArticleValidator()
	if err := articleValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := SaveOne(&articleValidator.article); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	serializer := ArticleSerializer{c, articleValidator.article}
	response := serializer.Response()

	a.Broadcast("POST", "articles", "/", gin.H{"article": response})
	c.JSON(http.StatusCreated, gin.H{"article": response})
}
