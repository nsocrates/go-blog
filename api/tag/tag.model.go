package tag

import (
	"github.com/jinzhu/gorm"
	m "github.com/nsocrates/go-blog/api/common"
	"github.com/nsocrates/go-blog/common"
)

func AutoMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&m.Tag{})
}

func FindMany() ([]m.Tag, error) {
	var tags []m.Tag
	err := common.DB.Find(&tags).Error
	return tags, err
}
