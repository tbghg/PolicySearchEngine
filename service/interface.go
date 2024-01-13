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
	// Destroy 销毁
	Destroy()
	// ExecuteWorkflow 执行一次工作流
	ExecuteWorkflow()
	// Watch 监控
	Watch()
}

type ServiceCrawler interface {
	Meta() MetaCrawler
	Content() ContentCrawler
	Register(crawler *Crawler)
}
