package service

type ContentCrawler interface {
	// Init 初始化
	Init()
	// Import 分批读取要爬取的网站 来自DB
	Import() (status int)
	// Run 启动
	Run()
}

type ContentCrawlerCollector struct {
	Crawlers []ContentCrawler
}

func (c *ContentCrawlerCollector) Run() {
	for _, crawler := range c.Crawlers {
		crawler.Init()
		for {
			status := crawler.Import()
			if status <= 0 {
				return
			}
			crawler.Run()
		}
	}
}
