package main

import (
	"PolicySearchEngine/config"
	"PolicySearchEngine/dao/database"
	"PolicySearchEngine/dao/es"
	"PolicySearchEngine/http"
)

func main() {
	// 配置初始化
	config.Init()
	database.Init()
	database.InitTable()
	es.Init()

	//var crawler service.Crawlers

	//var scienceColly science_center.ScienceColly
	//scienceColly.Register(&crawler)

	//var industryInformatizationColly industryInformatization_center.IndustryInformatizationColly
	//industryInformatizationColly.Register(&crawler)

	//var educationColly education_center.EducationColly
	//educationColly.Register(&crawler)

	//var stateCouncilColly stateCouncil_center.StateCouncilColly
	//stateCouncilColly.Register(&crawler)

	//var externalSourcesColly externalSources.ExternalSourcesColly
	//externalSourcesColly.Register(&crawler)

	//crawler.Run()
	http.Router()
}
