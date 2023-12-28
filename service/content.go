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
		c := crawler
		go func() {
			c.Init()
			for {
				if status := c.Import(); status >= 0 {
					c.Run()
				} else {
					return
				}
			}
		}()
	}
}
