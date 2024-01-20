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

		title := e.ChildText(".xxgk_title")
		regex := regexp.MustCompile(`[\n\t]`)
		cleanedTitle := regex.ReplaceAllString(title, "")
		fmt.Println(cleanedTitle)

		//content := e.ChildText("#Zoom")
		//fmt.Println(content)
	})

	// 可自行决定是否要上存储
	return &Rule{
		r: rule,
		c: c,
	}
}

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

func (s *ScienceContentColly) kjbgzCollector() *Rule {

	rule := regexp.MustCompile("https?://www\\.most\\.gov\\.cn/kjbgz/.*\\.html?")

	c := colly.NewCollector(
		colly.URLFilters(rule),
		colly.MaxDepth(0),
	)

	c.OnHTML("#Title", func(e *colly.HTMLElement) {
		title := e.Text
		regex := regexp.MustCompile(`[\n\t]`)
		cleanedTitle := regex.ReplaceAllString(title, "")
		fmt.Println(cleanedTitle)
	})

	c.OnHTML(".Zoom", func(e *colly.HTMLElement) {
		//context := e.Text
		//fmt.Println(context)
	})

	// 可自行决定是否要上存储
	return &Rule{
		r: rule,
		c: c,
	}
}

func (s *ScienceContentColly) zhengceCollector() *Rule {

	rule := regexp.MustCompile("https?://www\\.gov\\.cn/zhengce/content/.*\\.html?")

	c := colly.NewCollector(
		colly.URLFilters(rule),
		colly.MaxDepth(0),
	)

	c.OnHTML("td[colspan='3']", func(e *colly.HTMLElement) {
		title := e.Text
		regex := regexp.MustCompile(`[\n\t]`)
		cleanedTitle := regex.ReplaceAllString(title, "")
		fmt.Println(cleanedTitle)
	})

	c.OnHTML("#UCAP-CONTENT", func(e *colly.HTMLElement) {
		//context := e.Text
		//fmt.Println(context)
	})

	// 可自行决定是否要上存储
	return &Rule{
		r: rule,
		c: c,
	}
}

func (s *ScienceContentColly) gongbaoCollector() *Rule {

	rule := regexp.MustCompile("https?://www\\.gov\\.cn/gongbao/content/.*\\.html?")

	c := colly.NewCollector(
		colly.URLFilters(rule),
		colly.MaxDepth(0),
	)

	c.OnHTML(".share-title", func(e *colly.HTMLElement) {
		title := e.Text
		regex := regexp.MustCompile(`[\n\t]`)
		cleanedTitle := regex.ReplaceAllString(title, "")
		fmt.Println(cleanedTitle)
	})

	c.OnHTML(".pages_content", func(e *colly.HTMLElement) {
		//context := e.Text
		//fmt.Println(context)
	})

	// 可自行决定是否要上存储
	return &Rule{
		r: rule,
		c: c,
	}
}

func (s *ScienceContentColly) xinwenCollector() *Rule {

	rule := regexp.MustCompile("https?://www\\.gov\\.cn/xinwen/.*\\.html?")

	c := colly.NewCollector(
		colly.URLFilters(rule),
		colly.MaxDepth(0),
	)

	c.OnHTML("#ti", func(e *colly.HTMLElement) {
		title := e.Text
		regex := regexp.MustCompile(`[\n\t]`)
		cleanedTitle := regex.ReplaceAllString(title, "")
		fmt.Println(cleanedTitle)
	})

	c.OnHTML(".pages_content", func(e *colly.HTMLElement) {
		//context := e.Text
		//fmt.Println(context)
	})

	// 可自行决定是否要上存储
	return &Rule{
		r: rule,
		c: c,
	}
}

func (s *ScienceContentColly) chinataxCollector() *Rule {

	rule := regexp.MustCompile("https?://www\\.chinatax\\.gov\\.cn/.*\\.html?")

	c := colly.NewCollector(
		colly.URLFilters(rule),
		colly.MaxDepth(0),
	)

	c.OnHTML("title", func(e *colly.HTMLElement) {
		title := e.Text
		regex := regexp.MustCompile(`[\n\t]`)
		cleanedTitle := regex.ReplaceAllString(title, "")
		fmt.Println(cleanedTitle)
	})

	c.OnHTML("#fontzoom", func(e *colly.HTMLElement) {
		//context := e.Text
		//fmt.Println(context)
	})

	// 可自行决定是否要上存储
	return &Rule{
		r: rule,
		c: c,
	}
}
