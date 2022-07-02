package common

import (
	"github.com/gin-gonic/gin"
	"github.com/nsocrates/go-blog/common"
)

func UpdateContextUserModel(c *gin.Context, id uint) {
	var user User

	if id != 0 {
		common.DB.First(&user, id)
	}

	c.Set(common.MY_USER_ID, id)
	c.Set(common.MY_USER_MODEL, user)
}
