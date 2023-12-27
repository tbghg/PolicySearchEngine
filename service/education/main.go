package education

import (
	"PolicySearchEngine/service"
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
)

const initPage = "https://www.most.gov.cn/satp/kjzc/zh/index.html"

type EducationColly struct {
	c *colly.Collector
	// 遍历起始页
	startPages []string
}

func (s *EducationColly) Init() {

	s.c = colly.NewCollector(
		colly.AllowedDomains(
			"www.most.gov.cn",
		),
		colly.URLFilters(
			regexp.MustCompile("https://www\\.most\\.gov\\.cn/xxgk/xinxifenlei/fdzdgknr/fgzc/gfxwj/.*\\.html"),
			regexp.MustCompile("https://www\\.most\\.gov\\.cn/tztg/.*\\.html"),
			regexp.MustCompile("https://www\\.most\\.gov\\.cn/satp/kjzc/zh/.*\\.html"),
		),
		colly.MaxDepth(1),
	)

	// 请求链接时输出正在访问的链接
	s.c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

}

func (s *EducationColly) PageTraverse() {
	// todo 根据initPage起始页，确定要遍历的页数，暂时写死，等待后续优化
	s.startPages = append(s.startPages,
		initPage,
		"https://www.most.gov.cn/satp/kjzc/zh/index_1.html",
		"https://www.most.gov.cn/satp/kjzc/zh/index_2.html",
		"https://www.most.gov.cn/satp/kjzc/zh/index_3.html",
		"https://www.most.gov.cn/satp/kjzc/zh/index_4.html",
		"https://www.most.gov.cn/satp/kjzc/zh/index_5.html",
		"https://www.most.gov.cn/satp/kjzc/zh/index_6.html",
	)
}

func (s *EducationColly) Operate() {
	s.c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, e.Request.AbsoluteURL(link))
		err := s.c.Visit(e.Request.AbsoluteURL(link))

		if err != nil {
			fmt.Println(err)
		}
	})
}

func (s *EducationColly) Run() {
	for _, page := range s.startPages {
		err := s.c.Visit(page)
		if err != nil {
			fmt.Println(err)
		}
	}
}

var _ service.CrawlerColly = (*EducationColly)(nil)
