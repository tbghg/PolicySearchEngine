package meta

import (
	"PolicySearchEngine/dao/database"
	"PolicySearchEngine/dao/redis"
	"PolicySearchEngine/service"
	"PolicySearchEngine/utils"
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
)

const (
	initPage     = "https://www.most.gov.cn/satp/kjzc/zh/index.html"
	departmentID = 1
	provinceID   = 35
)

type ScienceMetaColly struct {
	c *colly.Collector
	// 遍历起始页
	startPages []string
	metaDal    *database.MetaDal
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
		colly.MaxDepth(1),
	)

	s.metaDal = &database.MetaDal{Db: database.MyDb()}
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

	redis.SetRedisStorage(s.c, "meta-sci", s.startPages)

	// 请求链接时输出正在访问的链接
	//s.c.OnRequest(func(r *colly.Request) {
	//	fmt.Println("Visiting ", r.URL)
	//})

	s.c.OnHTML(".list-main ul li", func(e *colly.HTMLElement) {

		url := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
		date := e.ChildText("span")
		title := e.ChildText("a")

		err := s.c.Visit(url)
		if errors.Is(err, colly.ErrAlreadyVisited) {
			return
		}
		if err != nil {
			fmt.Println(err.Error() + fmt.Sprintf(" %q -> %s\n", e.Text, url))
			return
		}

		dateTime, err := utils.StringToTime(date)
		if err != nil {
			fmt.Println(err.Error() + fmt.Sprintf("Time Falted %s %q -> %s\n", date, title, url))
			return
		}

		s.metaDal.UpdateMeta(dateTime, title, url, departmentID, provinceID)

		fmt.Printf("Link found: %s %q -> %s\n\n", date, title, url)
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
