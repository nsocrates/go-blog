package comment

import (
	"github.com/gin-gonic/gin"
	a "github.com/nsocrates/go-blog/api/common"
	"github.com/nsocrates/go-blog/common"
)

type CommentValidator struct {
	Comment struct {
		Body string `json:"body" binding:"max=2048"`
	} `json:"comment"`
	comment a.Comment `json:"-"`
}

func NewCommentValidator() CommentValidator {
	return CommentValidator{}
}

func (self *CommentValidator) Bind(c *gin.Context) error {
	myUser := c.MustGet(common.MY_USER_MODEL).(a.User)

	err := common.Bind(c, self)
	if err != nil {
		return err
	}

	self.comment.Body = self.Comment.Body
	self.comment.Author = a.GetArticleUser(myUser)
	return nil
}
