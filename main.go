package main

import (
	"PolicySearchEngine/service"
	"PolicySearchEngine/service/scienceAndTechnology"
	//"PolicySearchEngine/service/education"
)

func main() {
	var crawlerCollector service.CrawlerCollector

	var scienceColly scienceAndTechnology.ScienceColly
	//var educationColly education.EducationColly

	crawlerCollector.Crawlers = append(crawlerCollector.Crawlers,
		&scienceColly,
		//&educationColly,
	)

	crawlerCollector.Run()
}
