package main

import (
	"PolicySearchEngine/config"
	"PolicySearchEngine/dao/database"
	"PolicySearchEngine/service"
	"PolicySearchEngine/service/scienceAndTechnology"
)

func main() {
	// 配置初始化
	config.Init()
	database.Init()
	database.InitTable()

	var crawler service.Crawler

	var scienceColly scienceAndTechnology.ScienceColly
	scienceColly.Register(&crawler)
	// 新部门加入示例：
	//var educationColly education.EducationColly
	//educationColly.Register(&crawler)

	crawler.Run()
}
