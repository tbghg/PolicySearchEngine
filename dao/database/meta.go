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
func (m *MetaDal) InsertMeta(date time.Time, title string, url string, departmentID uint, provinceID uint) uint {
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
	return meta.ID
}

func (m *MetaDal) UpdateMetaTitle(title string, url string) {
	meta := model.Meta{
		Title: title,
		Url:   url,
	}
	result := m.Db.Where(model.Meta{Url: meta.Url}).Updates(&meta)
	if result.Error != nil {
		fmt.Printf("UpdateMetaTitle... %v", meta)
		log.Fatal(result.Error)
	}
}

func (m *MetaDal) GetAllMeta(departmentID, provinceID uint) *[]model.Meta {
	var metaList []model.Meta
	result := m.Db.Where(model.Meta{
		DepartmentID: departmentID,
		ProvinceID:   provinceID,
	}).Find(&metaList)
	if result.Error != nil {
		fmt.Printf("读取数据失败: %v\n", result.Error)
		return nil
	}
	return &metaList
}

func (m *MetaDal) GetAllMetaByIDs(provinceID uint, id uint) *[]model.Meta {
	var metaList []model.Meta
	result := m.Db.Where("province_id = ? AND id > ?", provinceID, id).Find(&metaList)
	if result.Error != nil {
		fmt.Printf("读取数据失败: %v\n", result.Error)
		return nil
	}
	return &metaList
}

func (m *MetaDal) GetMetaByUrl(url string) *model.Meta {
	var meta model.Meta
	result := m.Db.Where(model.Meta{
		Url: url,
	}).First(&meta)
	if result.Error != nil {
		fmt.Printf("读取数据失败: %v\n", result.Error)
		return nil
	}
	return &meta
}
