package database

import (
	"PolicySearchEngine/model"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type MetaDal struct{ Db *gorm.DB }

// InsertMeta 添加meta
func (m *MetaDal) InsertMeta(date time.Time, title string, url string, departmentID uint, provinceID uint) {
	meta := model.Meta{
		Date:         date,
		Title:        title,
		Url:          url,
		DepartmentID: departmentID,
		ProvinceID:   provinceID,
	}
	// 存在则忽略，不存在则插入
	result := m.Db.Where(model.Meta{Url: meta.Url}).FirstOrCreate(&meta)
	if result.Error != nil {
		fmt.Printf("InsertMeta... %s, %v", date.String(), meta)
		log.Fatal(result.Error)
	}
}

func (m *MetaDal) UpdateMeta(date time.Time, title string, url string) {
	meta := model.Meta{
		Date:  date,
		Title: title,
		Url:   url,
	}
	result := m.Db.Where(model.Meta{Url: meta.Url}).Updates(&meta)
	if result.Error != nil {
		fmt.Printf("UpdateMeta... %s, %v", date.String(), meta)
		log.Fatal(result.Error)
	}
}
