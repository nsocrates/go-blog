package tag

import (
	"github.com/gin-gonic/gin"
	c "github.com/nsocrates/go-blog/api/common"
)

type TagSerializer struct {
	C *gin.Context
	c.Tag
}

func (self *TagSerializer) Response() string {
	return self.Tag.Tag
}

type TagsSerializer struct {
	C    *gin.Context
	Tags []c.Tag
}

func (self *TagsSerializer) Response() []string {
	response := []string{}

	for _, t := range self.Tags {
		serializer := TagSerializer{self.C, t}
		response = append(response, serializer.Response())
	}

	return response
}
