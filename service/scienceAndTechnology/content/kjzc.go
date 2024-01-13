package content

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
)

func (s *ScienceContentColly) kjzcCollector() *Rule {

	rule1 := regexp.MustCompile("https?://www\\.most\\.gov\\.cn/satp/kjzc/zh/.*\\.html?")
	rule2 := regexp.MustCompile("https?://www\\.most\\.gov\\.cn/tztg/.*\\.html?")

	combinedRule := regexp.MustCompile(fmt.Sprintf(
		"(%s|%s)",
		rule1.String(),
		rule2.String(),
	))

	c := colly.NewCollector(
		colly.URLFilters(combinedRule),
		colly.MaxDepth(0),
	)

	c.OnHTML("#Title", func(e *colly.HTMLElement) {
		title := e.Text
		regex := regexp.MustCompile(`[\n\t]`)
		cleanedTitle := regex.ReplaceAllString(title, "")
		fmt.Println(cleanedTitle)
	})

	c.OnHTML(".Custom_UnionStyle", func(e *colly.HTMLElement) {
		//context := e.Text
		//fmt.Println(context)
	})

	// 可自行决定是否要上存储
	return &Rule{
		r: combinedRule,
		c: c,
	}
}
