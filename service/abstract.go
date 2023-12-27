package service

type CrawlerColly interface {
	// Init 初始化
	Init()
	// PageTraverse 页数遍历
	PageTraverse()
	// Operate 处理信息
	Operate()
	// Run 启动
	Run()
}

type CrawlerCollector struct {
	Crawlers []CrawlerColly
}

func (c *CrawlerCollector) Run() {
	for _, crawler := range c.Crawlers {
		crawler.Init()
		crawler.PageTraverse()
		crawler.Operate()
		crawler.Run()
	}
}
