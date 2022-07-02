package comment

import (
	"github.com/jinzhu/gorm"
	c "github.com/nsocrates/go-blog/api/common"
	"github.com/nsocrates/go-blog/common"
)

func AutoMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&c.Comment{})
}

func GetArticleComments(article *c.Article) error {
	tx := common.DB.Begin()
	tx.Model(&article).Related(&article.Comments, "Comments")

	for _, comment := range article.Comments {
		tx.Model(&comment).Related(&comment.Author, "Author")
		tx.Model(&comment.Author).Related(&comment.Author.User)
	}

	err := tx.Commit().Error

	return err
}

func SaveOne(data interface{}) error {
	err := common.DB.Save(data).Error
	return err
}

func Delete(condition interface{}) error {
	err := common.DB.Where(condition).Delete(c.Comment{}).Error
	return err
}
