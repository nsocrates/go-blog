package comment

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nsocrates/go-blog/api/article"
	a "github.com/nsocrates/go-blog/api/common"
	"github.com/nsocrates/go-blog/common"
)

func List(c *gin.Context) {
	slug := c.Param("slug")
	article, err := article.FindOne(&a.Article{Slug: slug})

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comments", errors.New("invalid slug")))
		return
	}

	err = GetArticleComments(&article)

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comments", errors.New("database error")))
		return
	}

	serializer := CommentsSerializer{c, article.Comments}
	response := serializer.Response()
	c.JSON(http.StatusOK, gin.H{"comments": response})
}

func Create(c *gin.Context) {
	slug := c.Param("slug")
	article, err := article.FindOne(&a.Article{Slug: slug})

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comments", errors.New("invalid slug")))
		return
	}

	commentValidator := NewCommentValidator()
	if err := commentValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	commentValidator.comment.Article = article
	if err := SaveOne(&commentValidator.comment); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	serializer := CommentSerializer{c, commentValidator.comment}
	response := serializer.Response()

	path := fmt.Sprintf("/%v/comments", slug)
	a.Broadcast("POST", "articles", path, gin.H{"comment": response})
	c.JSON(http.StatusCreated, gin.H{"comment": response})
}

func Destroy(c *gin.Context) {
	slug := c.Param("slug")
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comment", errors.New("invalid id")))
		return
	}

	id := uint(idInt)
	err = Delete([]uint{id})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("comment", errors.New("invalid id")))
		return
	}

	path := fmt.Sprintf("/%v/comments/%d", slug, id)
	a.Broadcast("DELETE", "articles", path, gin.H{"id": id})
	c.JSON(http.StatusOK, gin.H{"comment": "delete success"})
}
