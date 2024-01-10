package database

import (
	"PolicySearchEngine/model"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type MetaDal struct{ Db *gorm.DB }

// UpdateMeta 更新或添加meta
func (m *MetaDal) UpdateMeta(date time.Time, title string, url string, departmentID uint, provinceID uint) {
	meta := model.Meta{
		Date:         date,
		Title:        title,
		Url:          url,
		DepartmentID: departmentID,
		ProvinceID:   provinceID,
	}
	// 存在则更新，不存在则插入
	result := m.Db.Where(model.Meta{Url: meta.Url}).Assign(meta).FirstOrCreate(&meta)
	if result.Error != nil {
		fmt.Printf("I'm UpdateMeta... %s, %v", date.String(), meta)
		log.Fatal(result.Error)
	}
}
