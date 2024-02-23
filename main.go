package main

import (
	"PolicySearchEngine/config"
	"PolicySearchEngine/dao/database"
	"PolicySearchEngine/service"
	"PolicySearchEngine/service/industryInformatization-center"
)

func main() {
	// 配置初始化
	config.Init()
	database.Init()
	database.InitTable()

	var crawler service.Crawlers

	//var scienceColly science_center.ScienceColly
	//scienceColly.Register(&crawler)
	//
	//var educationColly education_center.EducationColly
	//educationColly.Register(&crawler)

	var industryInformatizationColly industryInformatization_center.IndustryInformatizationColly
	industryInformatizationColly.Register(&crawler)

	crawler.Run()
}
