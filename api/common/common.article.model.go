package common

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/nsocrates/go-blog/common"
)

type Article struct {
	gorm.Model
	Slug        string `gorm:"unique_id"`
	Title       string
	Description string `gorm:"size:2048"`
	Body        string `gorm:"size:2048"`
	Author      ArticleUser
	AuthorID    uint
	Tags        []Tag
	Comments    []Comment
}

type ArticleUser struct {
	gorm.Model
	User      User
	UserID    uint
	Articles  []Article
	Favorites []Favorite
}

type Favorite struct {
	gorm.Model
	Favorite     Article
	FavoriteID   uint
	FavoriteBy   ArticleUser
	FavoriteByID uint
}

func (article Article) UpdateArticle(data interface{}) error {
	err := common.DB.Model(article).Update(data).Error
	return err
}

func GetArticleUser(user User) ArticleUser {
	var articleUser ArticleUser

	if user.ID == 0 {
		return articleUser
	}

	condition := ArticleUser{UserID: user.ID}
	common.DB.Where(&articleUser, condition).FirstOrCreate(&articleUser)
	articleUser.User = user
	return articleUser
}

func (self ArticleUser) GetArticleFeed(limit, offset string) ([]Article, int, error) {
	var articles []Article
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
	followings := self.User.GetUserFollowings()

	var articleUsers []uint
	for _, following := range followings {
		articleUser := GetArticleUser(following)
		articleUsers = append(articleUsers, articleUser.ID)
	}

	tx.Where("author_id in (?)", articleUsers).Order("updated_at desc").Offset(offsetInt).Limit(limitInt).Find(&articles)

	for _, article := range articles {
		tx.Model(article).Related(&article.Author, "Author")
		tx.Model(article.Author).Related(&article.Author.User)
		tx.Model(article).Related(&article.Tags, "Tags")
	}

	err = tx.Commit().Error

	return articles, count, err
}

func (article Article) ArticleFavoritesCount() uint {
	var count uint
	condition := Favorite{FavoriteID: article.ID}
	common.DB.Model(&Favorite{}).Where(condition).Count(&count)
	return count
}

func (article Article) ArticleIsFavoritedBy(user ArticleUser) bool {
	var favorite Favorite
	condition := Favorite{
		FavoriteID:   article.ID,
		FavoriteByID: user.ID,
	}
	common.DB.Where(condition).First(&favorite)

	return favorite.ID != 0
}

func (article Article) FavoriteArticleBy(user ArticleUser) error {
	var favorite Favorite
	err := common.DB.FirstOrCreate(&favorite, &Favorite{
		FavoriteID:   article.ID,
		FavoriteByID: user.ID,
	}).Error

	return err
}

func (article Article) UnFavoriteArticleBy(user ArticleUser) error {
	condition := Favorite{
		FavoriteID:   article.ID,
		FavoriteByID: user.ID,
	}
	err := common.DB.Where(condition).Delete(Favorite{}).Error

	return err
}
