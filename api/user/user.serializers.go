package user

import (
	"github.com/gin-gonic/gin"
	u "github.com/nsocrates/go-blog/api/common"
	"github.com/nsocrates/go-blog/common"
)

type UserSerializer struct {
	C *gin.Context
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
	Token    string `json:"token"`
}

func (self *UserSerializer) Response() UserResponse {
	myUser := self.C.MustGet(common.MY_USER_MODEL).(u.User)
	response := UserResponse{
		Username: myUser.Username,
		Email:    myUser.Email,
		Bio:      myUser.Bio,
		Image:    myUser.Image,
		Token:    common.GetToken(myUser.ID),
	}

	return response
}

type ProfileSerializer struct {
	C *gin.Context
	u.User
}

type ProfileResponse struct {
	ID        uint   `json:"-"`
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Following bool   `json:"following"`
}

func (self *ProfileSerializer) Response() ProfileResponse {
	myUser := self.C.MustGet(common.MY_USER_MODEL).(u.User)
	response := ProfileResponse{
		ID:        self.ID,
		Username:  self.Username,
		Bio:       self.Bio,
		Image:     self.Image,
		Following: myUser.IsUserFollowing(self.User),
	}

	return response
}
