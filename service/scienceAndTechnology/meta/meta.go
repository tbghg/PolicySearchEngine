package meta

import (
	"PolicySearchEngine/dao/database"
	"PolicySearchEngine/dao/redis"
	"PolicySearchEngine/service"
	"PolicySearchEngine/utils"
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/robfig/cron/v3"
	"regexp"
	"time"
)

const (
	initPage     = "https://www.most.gov.cn/satp/kjzc/zh/index.html"
	departmentID = 1  // 科学技术部
	provinceID   = 35 // 中央
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

		s.metaDal.InsertMeta(dateTime, title, url, departmentID, provinceID)

		fmt.Printf("Link found: %s %q -> %s\n\n", date, title, url)
	})

}

func (s *ScienceMetaColly) Run() {
	for _, page := range s.startPages {
		err := s.c.Visit(page)
		if err != nil {
			fmt.Println(fmt.Sprintf("page:%s, err:%+v", page, err))
		}
	}
}

func (s *ScienceMetaColly) Destroy() {
	// 下次运行是在一天后了，指向nil，保证内存释放，让gc自动去回收
	s.c = nil
	s.metaDal = nil
	s.startPages = nil
}

func (s *ScienceMetaColly) ExecuteWorkflow() {
	s.Init()
	s.PageTraverse()
	s.Operate()
	s.Run()
	s.Destroy()
}

func (s *ScienceMetaColly) Watch() {
	c := cron.New()
	// 添加每天早上8点执行的定时任务
	_, err := c.AddFunc("57 12 * * *", func() {
		fmt.Printf("定时任务运行 time:%s departmentID:%d provinceID:%d \n", time.Now(), departmentID, provinceID)
		s.ExecuteWorkflow()
	})
	if err != nil {
		fmt.Printf("定时任务添加失败 err: %+v \n", err)
		return
	}
	// 启动定时任务
	c.Start()
}

var _ service.MetaCrawler = (*ScienceMetaColly)(nil)
