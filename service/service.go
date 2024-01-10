package service

type Crawler struct {
	ServiceCrawler []ServiceCrawler
}

func (c *Crawler) Run() {
	for _, crawler := range c.ServiceCrawler {

		meta := crawler.Meta()
		meta.Init()
		meta.PageTraverse()
		meta.Operate()
		meta.Run()

		//content := crawler.Content()
		//content.Init()
		//for {
		//	status := content.Import()
		//	if status <= 0 {
		//		return
		//	}
		//	content.Run()
		//}

		// todo 什么时候启动监控比较好？
	}
}

// Register 新部门加入Crawler
func (c *Crawler) Register(crawler ServiceCrawler) {
	c.ServiceCrawler = append(c.ServiceCrawler, crawler)
}
