package database

import (
	"PolicySearchEngine/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type ContentDal struct{ Db *gorm.DB }

func (c *ContentDal) InsertContent(url string, article string) {

	var meta model.Meta
	if err := c.Db.Where("url = ?", url).First(&meta).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			url = strings.Replace(url, "https://", "http://", 1)
			err = c.Db.Where("url = ?", url).First(&meta).Error
		}
		if err != nil {
			// 处理查找失败的情况，例如返回错误或者进行其他逻辑
			fmt.Printf("查找Meta记录失败 url:%s err:%+v\n", url, err)
			return
		}
	}

	content := model.Content{
		MetaID:  meta.ID,
		Article: article,
	}

	result := c.Db.Where(model.Content{MetaID: meta.ID}).
		Assign(model.Content{Article: article}).
		FirstOrCreate(&content)
	if result.Error != nil {
		fmt.Printf("插入Content记录失败 err:%+v", result.Error)
	}
}

func (c *ContentDal) GetContentByMetaID(id uint) *model.Content {
	var content model.Content
	result := c.Db.Where(model.Content{MetaID: id}).First(&content)
	if result.Error != nil {
		fmt.Printf("读取文章内容失败: %v\n", result.Error)
		return nil
	}
	return &content
}
