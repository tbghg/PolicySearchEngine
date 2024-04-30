package database

import (
	"PolicySearchEngine/model"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type SmallDepartmentMapDal struct{ Db *gorm.DB }

func (d *SmallDepartmentMapDal) InsertDID(metaID uint, sdID uint) {

	var dMap model.SmallDepartmentMap
	err := d.Db.Where("meta_id = ? and small_department_id = ?", metaID, sdID).First(&dMap).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Printf("查找dMap记录失败 metaID:%d, dID:%d, err:%+v\n", metaID, sdID, err)
		return
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	dMap = model.SmallDepartmentMap{
		MetaID:            metaID,
		SmallDepartmentID: sdID,
	}

	result := d.Db.Create(&dMap)
	if result.Error != nil {
		fmt.Printf("插入dMap记录失败 err:%+v", result.Error)
	}
}

func (d *SmallDepartmentMapDal) GetDepartmentIDsByMetaID(id uint) (sdIDs []uint) {
	var DepartmentMaps []model.SmallDepartmentMap
	result := d.Db.Where(model.SmallDepartmentMap{MetaID: id}).Find(&DepartmentMaps)
	if result.Error != nil {
		fmt.Printf("读取DepartmentMap失败: %v\n", result.Error)
		return sdIDs
	}
	for _, departmentMap := range DepartmentMaps {
		sdIDs = append(sdIDs, departmentMap.SmallDepartmentID)
	}
	return sdIDs
}
