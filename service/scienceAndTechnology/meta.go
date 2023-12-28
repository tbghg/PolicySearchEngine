package scienceAndTechnology

import (
	"PolicySearchEngine/service"
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
)

const initPage = "https://www.most.gov.cn/satp/kjzc/zh/index.html"

type ScienceMetaColly struct {
	c *colly.Collector
	// 遍历起始页
	startPages []string
}

func (s *ScienceMetaColly) Init() {

	s.c = colly.NewCollector(
		colly.AllowedDomains(
			"www.most.gov.cn",
			"www.gov.cn",
			"szs.mof.gov.cn",
			"www.chinatax.gov.cn",
		),
		colly.URLFilters(

			regexp.MustCompile("https?://www\\.most\\.gov\\.cn/xxgk/.*\\.html?"),
			regexp.MustCompile("https?://www\\.most\\.gov\\.cn/tztg/.*\\.html?"),
			regexp.MustCompile("https?://www\\.most\\.gov\\.cn/satp/kjzc/zh/.*\\.html?"),
			regexp.MustCompile("https?://www\\.most\\.gov\\.cn/kjbgz/.*\\.html?"),

			regexp.MustCompile("https?://szs\\.mof\\.gov\\.cn/zhengwuxinxi/zhengcefabu/.*\\.html?"),
			regexp.MustCompile("https?://www\\.chinatax\\.gov\\.cn/.*\\.html?"),
			regexp.MustCompile("https?://www\\.gov\\.cn/zhengce/content/.*\\.html?"),
			regexp.MustCompile("https?://www\\.gov\\.cn/xinwen/.*\\.html?"),

			regexp.MustCompile("https?://www\\.chinatax\\.gov\\.cn/.*\\.html?"),
			regexp.MustCompile("https?://www\\.gov\\.cn/gongbao/content/.*\\.htm"),
		),
		colly.DisallowedURLFilters(
			// 去除 404 页面
			regexp.MustCompile("http://www\\.mof\\.gov\\.cn/404\\.htm"),
		),
		colly.MaxDepth(0),
	)

}

func (s *ScienceMetaColly) PageTraverse() {
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

func (s *ScienceMetaColly) Operate() {

	// todo 存储应该写在 OnHTML 内部

	s.c.OnHTML(".list-main ul li", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		date := e.ChildText("span")
		err := s.c.Visit(e.Request.AbsoluteURL(link))
		if errors.Is(err, colly.ErrAlreadyVisited) {
			return
		}

		if err != nil {
			fmt.Println(err.Error() + fmt.Sprintf(" %q -> %s\n", e.Text, e.Request.AbsoluteURL(link)))
			return
		}
		fmt.Printf("Link found: %s %q -> %s\n\n", date, e.ChildText("a"), e.Request.AbsoluteURL(link))
	})

}

func (s *ScienceMetaColly) Run() {
	for _, page := range s.startPages {
		err := s.c.Visit(page)
		if err != nil {
			fmt.Println(err)
		}
	}
}

var _ service.MetaCrawler = (*ScienceMetaColly)(nil)
