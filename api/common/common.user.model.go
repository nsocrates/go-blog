package common

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/nsocrates/go-blog/common"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint   `gorm:"primary_key"`
	Username     string `gorm:"unique;uniqueIndex;not null"`
	Email        string `gorm:"unique;uniqueIndex;not null"`
	Bio          string `gorm:"size:1024"`
	Image        string
	PasswordHash string `gorm:"not null"`
}

type Follow struct {
	gorm.Model
	Following    User
	FollowingID  uint
	FollowedBy   User
	FollowedByID uint
}

func (user *User) SetUserPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}

	passwordByte := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	user.PasswordHash = string(passwordHash)
	return nil
}

func (user User) CheckUserPassword(password string) error {
	passwordByte := []byte(password)
	passwordHashByte := []byte(user.PasswordHash)
	err := bcrypt.CompareHashAndPassword(passwordHashByte, passwordByte)
	return err
}

func (user User) UpdateUser(data interface{}) error {
	err := common.DB.Model(user).Update(data).Error
	return err
}

func (user User) GetUserFollowings() []User {
	tx := common.DB.Begin()
	var follows []Follow
	var followings []User
	followModel := Follow{
		FollowedByID: user.ID,
	}

	tx.Where(followModel).Find(&follows)

	for _, follow := range follows {
		var userModel User
		tx.Model(&follow).Related(&userModel, "following")
		followings = append(followings, userModel)
	}

	tx.Commit()

	return followings
}

func (user User) IsUserFollowing(otherUser User) bool {
	var follow Follow
	followModel := Follow{
		FollowingID:  otherUser.ID,
		FollowedByID: user.ID,
	}
	common.DB.Where(followModel).First(&follow)
	return follow.ID != 0
}

func (user User) FollowUser(otherUser User) error {
	var follow Follow
	followModel := Follow{
		FollowingID:  otherUser.ID,
		FollowedByID: user.ID,
	}
	err := common.DB.FirstOrCreate(&follow, &followModel).Error
	return err
}

func (user User) UnFollowUser(otherUser User) error {
	followModel := Follow{
		FollowingID:  otherUser.ID,
		FollowedByID: user.ID,
	}

	err := common.DB.Where(followModel).Delete(Follow{}).Error
	return err
}
