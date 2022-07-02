package common

import (
	"github.com/jinzhu/gorm"
	"github.com/nsocrates/go-blog/common"
)

type Tag struct {
	gorm.Model
	Tag      string    `gorm:"unique_index"`
	Articles []Article `gorm:"many2many:article_tags;"`
}

func (article *Article) SetTags(tags []string) error {
	var tagList []Tag

	for _, tag := range tags {
		var tagModel Tag
		condition := Tag{Tag: tag}
		err := common.DB.FirstOrCreate(&tagModel, condition).Error

		if err != nil {
			return err
		}

		tagList = append(tagList, tagModel)
	}

	article.Tags = tagList
	return nil
}
