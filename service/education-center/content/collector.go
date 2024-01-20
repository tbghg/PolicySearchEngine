package content

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
)

func (s *EducationContentColly) moe418Collector() *Rule {

	rule1 := regexp.MustCompile("https?://www\\.moe\\.gov\\.cn/jyb_sjzl/sjzl_zcfg/moe_418/.*\\.html?")
	rule2 := regexp.MustCompile("https?://www\\.moe\\.gov\\.cn/jyb_sjzl/sjzl_zcfg/zcfg_jyfl/.*\\.html?")
	rule3 := regexp.MustCompile("https?://www\\.moe\\.gov\\.cn/jyb_sjzl/sjzl_zcfg/zcfg_jyxzfg/.*\\.html?")
	rule4 := regexp.MustCompile("https?://www\\.moe\\.gov\\.cn/jyb_sjzl/sjzl_zcfg/zcfg_qtxgfl/.*\\.html?")
	rule5 := regexp.MustCompile("https?://www\\.moe\\.gov\\.cn/jyb_sjzl/sjzl_zcfg/zcfg_qtxgfg/.*\\.html?")

	combinedRule := regexp.MustCompile(fmt.Sprintf(
		"(%s|%s|%s|%s|%s)",
		rule1.String(),
		rule2.String(),
		rule3.String(),
		rule4.String(),
		rule5.String(),
	))

	c := colly.NewCollector(
		colly.URLFilters(combinedRule),
		colly.MaxDepth(0),
	)

	c.OnHTML(".moe-detail-box h1", func(e *colly.HTMLElement) {
		title := e.Text
		regex := regexp.MustCompile(`[\n\t]`)
		cleanedTitle := regex.ReplaceAllString(title, "")
		fmt.Println(cleanedTitle)
	})

	c.OnHTML(".TRS_Editor", func(e *colly.HTMLElement) {
		//context := e.Text
		//fmt.Println(context)
	})

	// 可自行决定是否要上存储
	return &Rule{
		r: combinedRule,
		c: c,
	}
}

func (s *EducationContentColly) srcsiteCollector() *Rule {

	rule := regexp.MustCompile("https?://www\\.moe\\.gov\\.cn/srcsite/.*\\.html?")

	c := colly.NewCollector(
		colly.URLFilters(rule),
		colly.MaxDepth(0),
	)

	c.OnHTML(".details-policy-box h1", func(e *colly.HTMLElement) {
		title := e.Text
		regex := regexp.MustCompile(`[\n\t]`)
		cleanedTitle := regex.ReplaceAllString(title, "")
		fmt.Println(cleanedTitle)
	})

	c.OnHTML(".details-policy-box", func(e *colly.HTMLElement) {
		//context := e.Text
		//fmt.Println(context)
	})

	// 可自行决定是否要上存储
	return &Rule{
		r: rule,
		c: c,
	}
}
