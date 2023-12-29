package content

import (
	"PolicySearchEngine/model"
	"PolicySearchEngine/service"
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
)

type ScienceContentColly struct {
	rules []*Rule
	// 等待处理的url队列
	waitQueue []model.Meta
}

type Rule struct {
	r *regexp.Regexp
	c *colly.Collector
}

func (s *ScienceContentColly) Init() {
	// 注册规则
	s.rules = append(s.rules,
		s.xxgkCollector(),
	)
}

func (s *ScienceContentColly) Import() (status int) {
	//TODO 从DB里面读取，目前可以先不接DB
	if len(s.waitQueue) != 0 {
		return 0
	}
	s.waitQueue = []model.Meta{
		{
			//Date      :time.Date(),
			Title: "中共中央办公厅 国务院办公厅印发《关于进一步加强青年科技人才培养和使用的若干措...",
			Url:   "https://www.most.gov.cn/xxgk/xinxifenlei/fdzdgknr/fgzc/gfxwj/gfxwj2023/202305/t20230517_186077.html",
		},
	}
	return 1
}

func (s *ScienceContentColly) Run() {
	// todo waitQueue中的url均处理完即可
	for _, meta := range s.waitQueue {
		fmt.Printf("I'm dealing %s...\n", meta.Url)

		// 依次匹配规则
		for _, rule := range s.rules {
			if rule.r.MatchString(meta.Url) {
				err := rule.c.Visit(meta.Url)
				if err != nil {
					fmt.Println(err)
				}
				break
			}
		}

	}
}

var _ service.ContentCrawler = (*ScienceContentColly)(nil)
