package article

import (
	"strconv"

	"github.com/jinzhu/gorm"
	c "github.com/nsocrates/go-blog/api/common"
	"github.com/nsocrates/go-blog/common"
)

func AutoMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&c.Article{})
	DB.AutoMigrate(&c.ArticleUser{})
	DB.AutoMigrate((&c.Favorite{}))
}

func SaveOne(data interface{}) error {
	err := common.DB.Save(data).Error
	return err
}

func FindMany(tag, author, limit, offset, favorited string) ([]c.Article, int, error) {
	var articles []c.Article
	var count int

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		offsetInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 20
	}

	tx := common.DB.Begin()

	if tag != "" {
		var tagModel c.Tag
		tx.Where(c.Tag{Tag: tag}).First(&tagModel)

		if tagModel.ID != 0 {
			tx.Model(&tagModel).Offset(offsetInt).Limit(limitInt).Related(&articles, "Articles")
			count = tx.Model(&tagModel).Association("Articles").Count()
		}
	} else if author != "" {
		var userModel c.User
		tx.Where(c.User{Username: author}).First(&userModel)
		articleUserModel := c.GetArticleUser(userModel)

		if articleUserModel.ID != 0 {
			tx.Model(&articleUserModel).Offset(offsetInt).Limit(limitInt).Related(&articles, "Articles")
			count = tx.Model(&articleUserModel).Association("Articles").Count()
		}
	} else if favorited != "" {
		var userModel c.User
		tx.Where(c.User{Username: favorited}).First(&userModel)
		articleUserModel := c.GetArticleUser(userModel)

		if articleUserModel.ID != 0 {
			tx.Model(&articleUserModel).Offset(offsetInt).Limit(limitInt).Related(&articles, "Articles")
			count = tx.Model(&articleUserModel).Association("Articles").Count()
		}
	} else {
		common.DB.Model(&articles).Count(&count)
		common.DB.Offset(offsetInt).Limit(limitInt).Find(&articles)
	}

	for _, article := range articles {
		tx.Model(&article).Related(&article.Author, "Author")
		tx.Model(&article.Author).Related(&article.Author.User)
		tx.Model(&article).Related(&article.Tags, "Tags")
	}

	err = tx.Commit().Error

	return articles, count, err
}

func FindOne(condition interface{}) (c.Article, error) {
	var article c.Article
	tx := common.DB.Begin()
	tx.Where(condition).First(&article)
	tx.Model(&article).Related(&article.Author, "Author")
	tx.Model(&article.Author).Related(&article.Author.User)
	tx.Model(&article).Related(&article.Tags, "Tags")
	err := tx.Commit().Error
	return article, err
}

func Delete(condition interface{}) error {
	err := common.DB.Where(condition).Delete(c.Article{}).Error
	return err
}
