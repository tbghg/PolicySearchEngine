package main

import (
	"PolicySearchEngine/config"
	"PolicySearchEngine/service"
	"PolicySearchEngine/service/scienceAndTechnology"
)

//func main() {
//	// 配置初始化
//	config.Init()
//
//	var crawlerCollector service.MetaCrawlerCollector
//
//	var scienceColly scienceAndTechnology.ScienceMetaColly
//	//var educationColly education.EducationColly
//
//	crawlerCollector.Crawlers = append(crawlerCollector.Crawlers,
//		&scienceColly,
//		//&educationColly,
//	)
//
//	crawlerCollector.Run()
//}

func main() {
	// 配置初始化
	config.Init()

	var crawlerCollector service.ContentCrawlerCollector

	var scienceColly scienceAndTechnology.ScienceContentColly

	crawlerCollector.Crawlers = append(crawlerCollector.Crawlers,
		&scienceColly,
	)

	crawlerCollector.Run()
}
