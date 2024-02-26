package main

import (
	"PolicySearchEngine/config"
	"PolicySearchEngine/dao/database"
	"PolicySearchEngine/dao/es"
	"PolicySearchEngine/http"
	"PolicySearchEngine/service"
)

func main() {
	// 配置初始化
	config.Init()
	database.Init()
	database.InitTable()
	es.Init()

	var crawler service.Crawlers

	//var scienceColly science_center.ScienceColly
	//scienceColly.Register(&crawler)

	//var educationColly education_center.EducationColly
	//educationColly.Register(&crawler)

	//var industryInformatizationColly industryInformatization_center.IndustryInformatizationColly
	//industryInformatizationColly.Register(&crawler)

	crawler.Run()
	http.Router()
}
