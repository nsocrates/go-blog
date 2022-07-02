package user

import (
	"github.com/jinzhu/gorm"
	m "github.com/nsocrates/go-blog/api/common"
	"github.com/nsocrates/go-blog/common"
)

func AutoMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&m.User{})
	DB.AutoMigrate(&m.Follow{})
}

func FindOne(condition interface{}) (m.User, error) {
	var user m.User
	err := common.DB.Where(condition).First(&user).Error
	return user, err
}

func SaveOne(data interface{}) error {
	err := common.DB.Save(data).Error
	return err
}
