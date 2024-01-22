package service

import (
	"PolicySearchEngine/utils"
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
)

// Rule 每个正则对应一个收集器
type Rule struct {
	R *regexp.Regexp
	C *colly.Collector
}

type HtmlFunc struct {
	QuerySelect string
	F           func(e *colly.HTMLElement)
}

func NormalRule(rule *regexp.Regexp, hfs ...*HtmlFunc) *Rule {
	c := colly.NewCollector(
		colly.URLFilters(rule),
		colly.MaxDepth(0),
	)
	for _, hf := range hfs {
		c.OnHTML(hf.QuerySelect, hf.F)
	}
	return &Rule{R: rule, C: c}
}

func NormalTitle(e *colly.HTMLElement) {
	fmt.Println(utils.TidyString(e.Text))
}

func NormalContent(e *colly.HTMLElement) {
	//fmt.Println(e.Text)
}
