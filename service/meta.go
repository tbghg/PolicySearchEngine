package service

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

type MetaCrawlerCollector struct {
	Crawlers []MetaCrawler
}

func (m *MetaCrawlerCollector) Run() {
	for _, crawler := range m.Crawlers {
		c := crawler
		// todo 代码完善后添加协程
		c.Init()
		c.PageTraverse()
		c.Operate()
		c.Run()
	}
}
