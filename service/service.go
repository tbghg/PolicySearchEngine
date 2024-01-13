package service

type Crawlers struct {
	Crawler []Crawler
}

func (c *Crawlers) Run() {
	for _, crawler := range c.Crawler {

		meta := crawler.Meta()
		meta.ExecuteWorkflow()
		go meta.Watch()
		select {}
		//content := crawler.Content()
		//content.Init()
		//for {
		//	status := content.Import()
		//	if status <= 0 {
		//		return
		//	}
		//	content.Run()
		//}

	}
}

// Register 新部门加入Crawler
func (c *Crawlers) Register(crawler Crawler) {
	c.Crawler = append(c.Crawler, crawler)
}
