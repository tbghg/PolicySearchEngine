package main

import (
	"PolicySearchEngine/service"
	"PolicySearchEngine/service/scienceAndTechnology"
	//"PolicySearchEngine/service/education"
)

func main() {
	var crawlerCollector service.MetaCrawlerCollector

	var scienceColly scienceAndTechnology.ScienceMetaColly
	//var educationColly education.EducationColly

	crawlerCollector.Crawlers = append(crawlerCollector.Crawlers,
		&scienceColly,
		//&educationColly,
	)

	crawlerCollector.Run()
}
