package user

import (
	"strings"

	"github.com/gin-gonic/gin"
	u "github.com/nsocrates/go-blog/api/common"
	"github.com/nsocrates/go-blog/common"
)

type UserValidator struct {
	User struct {
		Username string `json:"username" binding:"required,alphanum,min=4,max=255"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=3,max=255"`
		Bio      string `json:"bio" binding:"max=1024"`
		Image    string `json:"image" binding:"omitempty,url"`
	} `json:"user"`
	user u.User `json:"-"`
}

func (self *UserValidator) Bind(c *gin.Context) error {
	if err := common.Bind(c, self); err != nil {
		return err
	}

	self.user.Username = strings.ToLower(self.User.Username)
	self.user.Email = strings.ToLower(self.User.Email)
	self.user.Bio = self.User.Bio
	self.user.Image = self.User.Image
	self.user.SetUserPassword(self.User.Password)

	return nil
}

func NewUserValidator() UserValidator {
	userValidator := UserValidator{}
	return userValidator
}

func NewUserValidatorFillWith(user u.User) UserValidator {
	userValidator := NewUserValidator()
	userValidator.User.Username = strings.ToLower(user.Username)
	userValidator.User.Email = strings.ToLower(user.Email)
	userValidator.User.Bio = user.Bio
	return userValidator
}

type LoginValidator struct {
	User struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=3,max=255"`
	} `json:"user"`
	user u.User `json:"-"`
}

func (self *LoginValidator) Bind(c *gin.Context) error {
	if err := common.Bind(c, self); err != nil {
		return err
	}

	self.user.Email = strings.ToLower(self.User.Email)
	return nil
}

func NewLoginValidator() LoginValidator {
	loginValidator := LoginValidator{}
	return loginValidator
}
