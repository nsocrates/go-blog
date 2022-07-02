package article

import (
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	c "github.com/nsocrates/go-blog/api/common"
	"github.com/nsocrates/go-blog/api/tag"
	"github.com/nsocrates/go-blog/api/user"
	"github.com/nsocrates/go-blog/common"
)

type ArticleUserSerializer struct {
	C *gin.Context
	c.ArticleUser
}

type ArticleSerializer struct {
	C *gin.Context
	c.Article
}

type ArticleResponse struct {
	ID             uint                 `json:"-"`
	Title          string               `json:"title"`
	Slug           string               `json:"slug"`
	Description    string               `json:"description"`
	Body           string               `json:"body"`
	CreatedAt      string               `json:"createdAt"`
	UpdatedAt      string               `json:"updatedAt"`
	Author         user.ProfileResponse `json:"author"`
	Tags           []string             `json:"tags"`
	Favorite       bool                 `json:"favorited"`
	FavoritesCount uint                 `json:"favoritesCount"`
}

type ArticlesSerializer struct {
	C        *gin.Context
	Articles []c.Article
}

func (self *ArticleUserSerializer) Response() user.ProfileResponse {
	response := user.ProfileSerializer{self.C, self.ArticleUser.User}
	return response.Response()
}

func (self *ArticleSerializer) Response() ArticleResponse {
	myUser := self.C.MustGet(common.MY_USER_MODEL).(c.User)
	authorSerializer := ArticleUserSerializer{self.C, self.Author}
	myArticleUser := c.GetArticleUser(myUser)
	response := ArticleResponse{
		ID:             self.ID,
		Slug:           slug.Make(self.Title),
		Title:          self.Title,
		Description:    self.Description,
		Body:           self.Body,
		CreatedAt:      self.CreatedAt.UTC().Format(common.DATE_LAYOUT_ISO),
		UpdatedAt:      self.UpdatedAt.UTC().Format(common.DATE_LAYOUT_ISO),
		Author:         authorSerializer.Response(),
		Favorite:       self.ArticleIsFavoritedBy(myArticleUser),
		FavoritesCount: self.ArticleFavoritesCount(),
	}
	response.Tags = make([]string, 0)
	for _, t := range self.Tags {
		serializer := tag.TagSerializer{self.C, t}
		response.Tags = append(response.Tags, serializer.Response())
	}

	return response
}

func (self *ArticlesSerializer) Response() []ArticleResponse {
	response := []ArticleResponse{}

	for _, article := range self.Articles {
		serializer := ArticleSerializer{self.C, article}
		response = append(response, serializer.Response())
	}

	return response
}
