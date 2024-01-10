package database

import (
	"PolicySearchEngine/config"
	"PolicySearchEngine/model"
	mysqlCfg "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var myDb *gorm.DB

func InitTable() {
	// 初始化数据表
	_ = myDb.AutoMigrate(&model.Meta{})
	_ = myDb.AutoMigrate(&model.Content{})
	_ = myDb.AutoMigrate(&model.Department{})
	_ = myDb.AutoMigrate(&model.Province{})
}

func Init() {
	// 数据库配置
	cfg := mysqlCfg.Config{
		User:      config.V.GetString("mysql.user"),
		Passwd:    config.V.GetString("mysql.password"),
		Net:       "tcp",
		Addr:      config.V.GetString("mysql.addr"),
		DBName:    config.V.GetString("mysql.dbname"),
		Loc:       time.Local,
		ParseTime: true,
		// 允许原生密码
		AllowNativePasswords: true,
	}

	// 连接数据库
	db, err := gorm.Open(mysql.Open(cfg.FormatDSN()))
	if err != nil {
		log.Fatal(err)
	}

	myDb = db
	return
}

func MyDb() *gorm.DB {
	return myDb
}
