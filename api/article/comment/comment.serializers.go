package comment

import (
	"github.com/gin-gonic/gin"
	"github.com/nsocrates/go-blog/api/article"
	c "github.com/nsocrates/go-blog/api/common"
	"github.com/nsocrates/go-blog/api/user"
	"github.com/nsocrates/go-blog/common"
)

type CommentSerializer struct {
	C *gin.Context
	c.Comment
}

type CommentsSerializer struct {
	C        *gin.Context
	Comments []c.Comment
}

type CommentResponse struct {
	ID        uint                 `json:"-"`
	Body      string               `json:"body"`
	CreatedAt string               `json:"createdAt"`
	UpdatedAt string               `json:"updatedAt"`
	Author    user.ProfileResponse `json:"author"`
}

func (self *CommentSerializer) Response() CommentResponse {
	authorSerializer := article.ArticleUserSerializer{self.C, self.Author}
	response := CommentResponse{
		ID:        self.ID,
		Body:      self.Body,
		CreatedAt: self.CreatedAt.UTC().Format(common.DATE_LAYOUT_ISO),
		UpdatedAt: self.UpdatedAt.UTC().Format(common.DATE_LAYOUT_ISO),
		Author:    authorSerializer.Response(),
	}

	return response
}

func (self *CommentsSerializer) Response() []CommentResponse {
	response := []CommentResponse{}

	for _, comment := range self.Comments {
		serializer := CommentSerializer{self.C, comment}
		response = append(response, serializer.Response())
	}

	return response
}
