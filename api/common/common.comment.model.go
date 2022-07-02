package common

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Article   Article
	ArticleID uint
	Author    ArticleUser
	AuthorID  uint
	Body      string `gorm:"size:2048"`
}
