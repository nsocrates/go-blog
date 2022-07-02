package article

import (
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	a "github.com/nsocrates/go-blog/api/common"
	"github.com/nsocrates/go-blog/common"
)

type ArticleValidator struct {
	Article struct {
		Title       string   `json:"title" binding:"required,min=4"`
		Description string   `json:"description" binding:"max=2048"`
		Body        string   `json:"body" binding:"max=2084"`
		Tags        []string `json:"tags"`
	} `json:"article"`
	article a.Article
}

func NewArticleValidator() ArticleValidator {
	return ArticleValidator{}
}

func NewArticleValidatorFillWith(article a.Article) ArticleValidator {
	articleValidator := NewArticleValidator()
	articleValidator.Article.Title = article.Title
	articleValidator.Article.Description = article.Description
	articleValidator.Article.Body = article.Body

	for _, tag := range article.Tags {
		articleValidator.Article.Tags = append(articleValidator.Article.Tags, tag.Tag)
	}

	return articleValidator
}

func (self *ArticleValidator) Bind(c *gin.Context) error {
	myUser := c.MustGet(common.MY_USER_MODEL).(a.User)

	err := common.Bind(c, self)
	if err != nil {
		return err
	}

	self.article.Slug = slug.Make(self.Article.Title)
	self.article.Title = self.Article.Title
	self.article.Description = self.Article.Description
	self.article.Body = self.Article.Body
	self.article.Author = a.GetArticleUser(myUser)
	self.article.SetTags(self.Article.Tags)
	return nil
}
