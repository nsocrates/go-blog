package tag

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nsocrates/go-blog/common"
)

func List(c *gin.Context) {
	tags, err := FindMany()

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("articles", errors.New("invalid param")))
		return
	}

	serializer := TagsSerializer{c, tags}
	response := serializer.Response()
	c.JSON(http.StatusOK, gin.H{"tags": response})
}
