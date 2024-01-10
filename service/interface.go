package service

type ContentCrawler interface {
	// Init 初始化
	Init()
	// Import 分批读取要爬取的网站 来自DB
	Import() (status int)
	// Run 启动
	Run()
}

type MetaCrawler interface {
	// Init 初始化
	Init()
	// PageTraverse 页数遍历
	PageTraverse()
	// Operate 处理信息
	Operate()
	// Run 启动
	Run()
}

type ServiceCrawler interface {
	Meta() MetaCrawler
	Content() ContentCrawler
	Register(crawler *Crawler)
}
