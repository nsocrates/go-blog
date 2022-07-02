package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/nsocrates/go-blog/api/article"
	"github.com/nsocrates/go-blog/api/article/comment"
	"github.com/nsocrates/go-blog/api/tag"
	"github.com/nsocrates/go-blog/api/user"
	"github.com/nsocrates/go-blog/common"
)

func ConfigureDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	user.AutoMigrate(database)
	article.AutoMigrate(database)
	comment.AutoMigrate(database)
	tag.AutoMigrate(database)

	common.DB = database
}
