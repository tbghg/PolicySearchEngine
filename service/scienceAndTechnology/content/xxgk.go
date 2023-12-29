package content

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
)

func (s *ScienceContentColly) xxgkCollector() *Rule {
	rule := regexp.MustCompile("https?://www\\.most\\.gov\\.cn/xxgk/.*\\.html?")
	c := colly.NewCollector(
		colly.URLFilters(rule),
		colly.MaxDepth(0),
	)

	c.OnHTML(".xxgk_detail_content", func(e *colly.HTMLElement) {
		//fmt.Println(e.DOM.Html())
		fmt.Println(e.Text)
	})

	// 可自行决定是否要上存储
	return &Rule{
		r: rule,
		c: c,
	}
}
