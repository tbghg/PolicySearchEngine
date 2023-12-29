package database

import (
	"gorm.io/gorm"
)

type MetaDal struct{ Db *gorm.DB }

func (c *MetaDal) AddMeta() {

}
