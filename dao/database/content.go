package database

import (
	"gorm.io/gorm"
)

type ContentDal struct{ Db *gorm.DB }

func (c *ContentDal) AddContent() {

}
