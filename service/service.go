package service

import (
	"PolicySearchEngine/config"
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

type Crawlers struct {
	Name    string
	Crawler []Crawler
}

func (c *Crawlers) Run() {
	cr := cron.New()
	for _, crawler := range c.Crawler {

		meta := crawler.Meta()
		content := crawler.Content()

		// todo 先运行一遍，防止本来就有问题，代码稳定后可删除
		//meta.ExecuteWorkflow()
		content.ExecuteWorkflow()

		spec := config.V.GetString("cron." + c.Name)
		_, err := cr.AddFunc(spec, func() {
			fmt.Printf("定时任务运行 time:%s name:%s \n", time.Now(), c.Name)
			meta.ExecuteWorkflow()
			content.ExecuteWorkflow()
		})
		if err != nil {
			fmt.Printf("定时任务添加失败 err: %+v \n", err)
			return
		}
	}
	cr.Start()
	select {}
}

// Register 新部门加入Crawler
func (c *Crawlers) Register(name string, crawler Crawler) {
	c.Name = name
	c.Crawler = append(c.Crawler, crawler)
}
