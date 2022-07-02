package user

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	u "github.com/nsocrates/go-blog/api/common"
	"github.com/nsocrates/go-blog/common"
)

func Register(c *gin.Context) {
	validator := NewUserValidator()

	if err := validator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	_, errUsername := FindOne(u.User{Username: validator.user.Username})
	_, errEmail := FindOne(u.User{Email: validator.user.Email})

	if errUsername == nil || errEmail == nil {
		var errMsg []string

		if errUsername == nil {
			errMsg = append(errMsg, "username already exists")
		}

		if errEmail == nil {
			errMsg = append(errMsg, "email already exists")
		}

		c.JSON(http.StatusConflict, common.NewError("registration", errors.New(strings.Join(errMsg, "; "))))
		return
	}

	if err := SaveOne(&validator.user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	c.Set(common.MY_USER_MODEL, validator.user)
	serializer := UserSerializer{c}
	response := serializer.Response()
	c.JSON(http.StatusCreated, gin.H{"user": response})
}

func Login(c *gin.Context) {
	validator := NewLoginValidator()
	if err := validator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	user, err := FindOne(&u.User{Email: validator.user.Email})

	if err != nil || user.CheckUserPassword(validator.User.Password) != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("not registered email or invalid password")))
		return
	}

	u.UpdateContextUserModel(c, user.ID)
	serializer := UserSerializer{c}
	response := serializer.Response()
	c.JSON(http.StatusOK, gin.H{"user": response})
}

func Show(c *gin.Context) {
	username := strings.ToLower(c.Param("username"))
	user, err := FindOne(u.User{Username: username})

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("profile", errors.New("invalid username")))
		return
	}

	serializer := ProfileSerializer{c, user}
	response := serializer.Response()
	c.JSON(http.StatusOK, gin.H{"profile": response})
}

func Update(c *gin.Context) {
	myUser := c.MustGet(common.MY_USER_MODEL).(u.User)
	validator := NewUserValidatorFillWith(myUser)

	if err := validator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	validator.user.ID = myUser.ID

	if err := myUser.UpdateUser(validator.user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	u.UpdateContextUserModel(c, myUser.ID)
	serializer := UserSerializer{c}
	response := serializer.Response()

	c.JSON(http.StatusOK, gin.H{"user": response})
}

func Me(c *gin.Context) {
	serializer := UserSerializer{c}
	response := serializer.Response()
	c.JSON(http.StatusOK, gin.H{"user": response})
}

func ProfileFollow(c *gin.Context) {
	username := strings.ToLower(c.Param("username"))
	condition := u.User{Username: username}
	user, err := FindOne(condition)

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("profile", errors.New("invalid username")))
		return
	}

	myUser := c.MustGet(common.MY_USER_MODEL).(u.User)

	err = myUser.FollowUser(user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	serializer := ProfileSerializer{c, user}
	response := serializer.Response()
	c.JSON(http.StatusOK, gin.H{"profile": response})
}

func ProfileUnfollow(c *gin.Context) {
	username := strings.ToLower(c.Param("username"))
	condition := u.User{Username: username}
	user, err := FindOne(condition)

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("profile", errors.New("invalid username")))
		return
	}

	myUser := c.MustGet(common.MY_USER_MODEL).(u.User)

	err = myUser.UnFollowUser(user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	serializer := ProfileSerializer{c, user}
	response := serializer.Response()
	c.JSON(http.StatusOK, gin.H{"profile": response})
}
